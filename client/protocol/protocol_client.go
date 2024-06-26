package protocol

import (
	"context"
	clientContext "flare-tlc/client/context"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/registry"
	"flare-tlc/utils/contracts/system"
	"math/big"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type ProtocolClient struct {
	subProtocols []*SubProtocol
	eth          *ethclient.Client

	protocolContext *protocolContext

	submitter1         *Submitter
	submitter2         *Submitter
	signatureSubmitter *SignatureSubmitter

	votingEpoch    *utils.Epoch
	systemsManager *system.FlareSystemsManager

	rewardEpoch     *utils.Epoch
	registry        voterRegistry
	identityAddress common.Address
}

type voterRegistry interface {
	IsVoterRegistered(context.Context, common.Address, int64) (bool, error)
}

type voterRegistryImpl struct {
	registry *registry.Registry
}

func (r voterRegistryImpl) IsVoterRegistered(
	ctx context.Context, address common.Address, epoch int64,
) (bool, error) {
	return r.registry.IsVoterRegistered(&bind.CallOpts{Context: ctx}, address, big.NewInt(epoch))
}

func NewProtocolClient(ctx clientContext.ClientContext) (*ProtocolClient, error) {
	cfg := ctx.Config()

	if !cfg.Clients.EnabledProtocolVoting {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	cl, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	systemsManager, err := system.NewFlareSystemsManager(cfg.ContractAddresses.SystemsManager, cl)
	if err != nil {
		return nil, errors.Wrap(err, "error creating system manager contract")
	}

	votingEpoch, err := shared.VotingEpochFromChain(systemsManager)
	if err != nil {
		return nil, errors.Wrap(err, "error getting voting epoch")
	}

	protocolContext, err := newProtocolContext(cfg)
	if err != nil {
		return nil, err
	}

	rewardEpoch, err := shared.RewardEpochFromChain(systemsManager)
	if err != nil {
		return nil, err
	}

	var subProtocols []*SubProtocol
	for _, protocol := range cfg.Protocol {
		subProtocols = append(subProtocols, NewSubProtocol(protocol))
	}

	registryClient, err := registry.NewRegistry(cfg.ContractAddresses.VoterRegistry, cl)
	if err != nil {
		return nil, err
	}

	pc := &ProtocolClient{
		eth:             cl,
		protocolContext: protocolContext,
		subProtocols:    subProtocols,
		votingEpoch:     votingEpoch,
		systemsManager:  systemsManager,
		rewardEpoch:     rewardEpoch,
		registry:        voterRegistryImpl{registryClient},
		identityAddress: cfg.Identity.Address,
	}

	selectors := newContractSelectors()

	if cfg.Submit1.Enabled {
		pc.submitter1 = newSubmitter(cl, protocolContext, votingEpoch,
			&cfg.Submit1, &cfg.SubmitGas, selectors.submit1, subProtocols, 0, "submit1")
	} else {
		logger.Warn("submit1 is disabled")
	}
	if cfg.Submit2.Enabled {
		pc.submitter2 = newSubmitter(cl, protocolContext, votingEpoch,
			&cfg.Submit2, &cfg.SubmitGas, selectors.submit2, subProtocols, -1, "submit2")
	} else {
		logger.Warn("submit2 is disabled")
	}
	if cfg.SubmitSignatures.Enabled {
		pc.signatureSubmitter = newSignatureSubmitter(cl, protocolContext, votingEpoch,
			&cfg.SubmitSignatures, &cfg.SubmitGas, selectors.submitSignatures, subProtocols)
	} else {
		logger.Warn("submitSignatures is disabled")
	}
	return pc, nil
}

func (c *ProtocolClient) Run(ctx context.Context) error {
	if err := c.waitUntilRegistered(ctx); err != nil {
		return err
	}

	done := make(chan bool, 1)
	var wg sync.WaitGroup

	logger.Info("Starting submitters, waiting for next voting round start.")
	ticker := utils.NewEpochTicker(c.votingEpoch)
L:
	for {
		select {
		case currentEpoch := <-ticker.C:
			if c.submitter1 != nil {
				go func() {
					time.Sleep(c.submitter1.startOffset)
					if c.submitter2 != nil {
						// if running submitter1, and submitter2 is enabled,
						// we need to wait for it to complete before shutdown.
						wg.Add(1)
					}
					c.submitter1.RunEpoch(currentEpoch)
				}()
			}

			if c.submitter2 != nil {
				go func() {
					// Submit2 processes the current epoch data in the following epoch
					// so we wait a full epoch duration + offset before invoking.
					// TODO: this assumes c.submitter2.epochOffset is always -1
					time.Sleep(ticker.Epoch.Period + c.submitter2.startOffset)
					c.submitter2.RunEpoch(currentEpoch + 1)
					if c.submitter1 != nil {
						wg.Done()
					}
				}()
			}
			if c.signatureSubmitter != nil {
				// signatureSubmitter is independent of submit1 and submit2
				go func() {
					time.Sleep(c.signatureSubmitter.startOffset)
					c.signatureSubmitter.RunEpoch(currentEpoch)
				}()
			}
		case <-ctx.Done():
			if c.submitter1 != nil && c.submitter2 != nil {
				logger.Warn("Stopping submitters. Making sure both submit1 & submit2 have completed for the voting round. Not running submit2 might result in reward penalties.")
				wg.Wait()
			}

			done <- true
			break L
		}
	}

	<-done
	return nil
}

func (c *ProtocolClient) waitUntilRegistered(ctx context.Context) error {
	for {
		currentEpoch := c.rewardEpoch.EpochIndex(time.Now())

		registered, err := c.isRegistered(ctx, currentEpoch)
		if err != nil {
			return err
		}

		if registered {
			return nil
		}

		if err := c.waitForNextRewardEpoch(ctx, currentEpoch); err != nil {
			return err
		}
	}
}

const registerCheckTimeout = 5 * time.Second

func (c *ProtocolClient) isRegistered(ctx context.Context, epoch int64) (bool, error) {
	bOff := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	var registered bool

	// Use an exponential backoff retry in case of temporary errors
	// in querying the registry contract.
	err := backoff.RetryNotify(
		func() (err error) {
			ctx, cancel := context.WithTimeout(ctx, registerCheckTimeout)
			defer cancel()

			registered, err = c.registry.IsVoterRegistered(ctx, c.identityAddress, epoch)

			return err
		},
		bOff,
		func(err error, d time.Duration) {
			logger.Error(
				"unexpected error while checking submitter registration: %s, retrying after %s", err, d,
			)
		},
	)

	return registered, err
}

func (c *ProtocolClient) waitForNextRewardEpoch(ctx context.Context, currentEpoch int64) error {
	nextEpochStart := c.rewardEpoch.StartTime(currentEpoch + 1)
	now := time.Now()

	// Edge case if the time passed while checking the registration means
	// we are already in the next epoch - return immediately in that case.
	if nextEpochStart.Before(now) {
		logger.Info("submitter is not registered for voting epoch %d, checking the next epoch", currentEpoch)
		return nil
	}

	sleepTime := nextEpochStart.Sub(now)

	logger.Info(
		"submitter is not registered for voting epoch %d, waiting for %s until the next epoch",
		currentEpoch,
		sleepTime,
	)

	select {
	case <-time.After(sleepTime):
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}
