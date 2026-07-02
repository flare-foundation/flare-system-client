package protocol

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	clientContext "github.com/flare-foundation/flare-system-client/client/context"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/registry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// client manages fetching data from subProtocol providers and posting them on submission contract.
type client struct {
	subProtocols []*SubProtocol
	eth          *ethclient.Client

	protocolContext *protocolContext

	submitter1         *Submitter
	submitter2         *Submitter
	signatureSubmitter *SignatureSubmitter

	votingRoundTiming *utils.EpochTimingConfig
	systemsManager    *system.FlareSystemsManager

	rewardEpochTiming *utils.EpochTimingConfig
	registry          voterRegistry
	identityAddress   common.Address
}

type voterRegistry interface {
	IsVoterRegistered(context.Context, common.Address, int64) (bool, error)
}

type voterRegistryImpl struct {
	registry *registry.Registry
}

func (r voterRegistryImpl) IsVoterRegistered(ctx context.Context, address common.Address, epoch int64) (bool, error) {
	return r.registry.IsVoterRegistered(&bind.CallOpts{Context: ctx}, address, big.NewInt(epoch))
}

// NewClient creates new Client that manages fetching data from subProtocol providers and posting them on submission contract.
//
// messageChannel is used to provider messages from submitSignature to the finalizer.Client.
func NewClient(ctx clientContext.ClientContext, messageChannel chan<- shared.ProtocolMessage) (*client, error) {
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
		return nil, fmt.Errorf("creating system manager contract: %w", err)
	}

	votingRoundTiming, err := shared.VotingRoundTimingFromChain(systemsManager)
	if err != nil {
		return nil, fmt.Errorf("getting voting round timing: %w", err)
	}

	protocolContext, err := newProtocolContext(cfg)
	if err != nil {
		return nil, err
	}

	rewardEpochTiming, err := shared.RewardEpochTimingFromChain(systemsManager)
	if err != nil {
		return nil, err
	}

	subProtocols := make([]*SubProtocol, 0, len(cfg.Protocol))
	for _, protocol := range cfg.Protocol {
		subProtocols = append(subProtocols, NewSubProtocol(protocol))
	}

	registryClient, err := registry.NewRegistry(cfg.ContractAddresses.VoterRegistry, cl)
	if err != nil {
		return nil, err
	}

	pc := &client{
		eth:               cl,
		protocolContext:   protocolContext,
		subProtocols:      subProtocols,
		votingRoundTiming: votingRoundTiming,
		systemsManager:    systemsManager,
		rewardEpochTiming: rewardEpochTiming,
		registry:          voterRegistryImpl{registryClient},
		identityAddress:   cfg.Identity.Address,
	}

	selectors := ContractSelectors()

	pc.submitter1 = newSubmitter(cl, protocolContext, votingRoundTiming,
		&cfg.Submit1, &cfg.SubmitGas, selectors.submit1, subProtocols, 0, "submit1")

	pc.submitter2 = newSubmitter(cl, protocolContext, votingRoundTiming,
		&cfg.Submit2, &cfg.SubmitGas, selectors.submit2, subProtocols, -1, "submit2")

	pc.signatureSubmitter = newSignatureSubmitter(cl, protocolContext, votingRoundTiming,
		&cfg.SubmitSignatures, &cfg.SubmitGas, selectors.submitSignatures, subProtocols, messageChannel)

	return pc, nil
}

// Run runs the client. Should be called in a goroutine.
func (c *client) Run(ctx context.Context) error {
	if err := c.waitUntilRegistered(ctx); err != nil {
		return err
	}

	var wg sync.WaitGroup

	logger.Info("Starting submitters, waiting for next voting round start.")
	ticker := utils.NewEpochTicker(c.votingRoundTiming)
	for {
		select {
		case currentEpoch := <-ticker.C:
			// submit1 (commit) -> submit2 (reveal) -> submitSignatures is a
			// dependency chain: a submit1 with no submit2 is penalised (FTSO), as
			// is a submit2 with no submitSignatures (FDC). Run the enabled
			// submitters as one chain so these obligations survive shutdown (see
			// runChain). submit1 is never enabled without submit2, so the enabled
			// submitters are contiguous.
			if chain := c.submitterChain(ticker, currentEpoch); len(chain) > 0 {
				wg.Go(func() { runChain(ctx, chain) })
			}
		case <-ctx.Done():
			logger.Warn("Stopping submitters. Waiting for the in-flight submitter chain to complete for the voting round.")
			wg.Wait()
			return nil
		}
	}
}

// submitterStep is one submitter invocation within a chain: wait until offset
// (from the round tick), then run under the context runChain supplies.
type submitterStep struct {
	offset time.Duration
	run    func(ctx context.Context)
}

// submitterChain builds the ordered dependency chain of enabled submitters for
// the given voting round, in protocol order (submit1, submit2,
// submitSignatures), each with its offset from the round tick.
func (c *client) submitterChain(ticker *utils.EpochTicker, currentEpoch int64) []submitterStep {
	var chain []submitterStep
	if c.submitter1 != nil {
		chain = append(chain, submitterStep{
			offset: c.submitter1.startOffset,
			run:    func(ctx context.Context) { c.submitter1.RunEpoch(ctx, currentEpoch) },
		})
	}
	if c.submitter2 != nil {
		// Submit2 processes the current epoch data in the following epoch, so it
		// waits a full epoch period + offset before invoking.
		// TODO: this assumes c.submitter2.epochOffset is always -1
		chain = append(chain, submitterStep{
			offset: ticker.Epoch.Period + c.submitter2.startOffset,
			run:    func(ctx context.Context) { c.submitter2.RunEpoch(ctx, currentEpoch+1) },
		})
	}
	if c.signatureSubmitter != nil {
		chain = append(chain, submitterStep{
			offset: ticker.Epoch.Period + c.signatureSubmitter.startOffset,
			run:    func(ctx context.Context) { c.signatureSubmitter.RunEpoch(ctx, currentEpoch) },
		})
	}
	return chain
}

// runChain runs one voting round's submitters. The wait before the first (gate)
// step is the only cancellable part: if shutdown happens before the gate offset,
// nothing has run and the chain exits with nothing owed. Past the gate every
// step is an obligation, so each runs via runStep (detached from shutdown,
// panics recovered) in its own goroutine on its own offset, so a slow step
// delays no other. runChain returns only once every step completes, so the
// caller's WaitGroup drains the round on shutdown. Offsets are assumed
// non-decreasing in protocol order.
func runChain(ctx context.Context, steps []submitterStep) {
	if len(steps) == 0 {
		return
	}
	start := time.Now()

	// Cancellable gate: the only point shutdown can abandon the round unowed.
	if !sleepUnlessCancelled(ctx, time.Until(start.Add(steps[0].offset))) {
		return
	}

	var wg sync.WaitGroup
	for _, step := range steps {
		wg.Go(func() {
			if wait := time.Until(start.Add(step.offset)); wait > 0 {
				time.Sleep(wait)
			}
			runStep(ctx, step)
		})
	}
	wg.Wait()
}

// runStep runs one chain step under a context shutdown does not cancel (else
// the submit work would no-op and break the obligation), cancelled once the
// step returns. A panic is recovered so it can't skip the remaining steps.
func runStep(ctx context.Context, step submitterStep) {
	stepCtx, cancel := context.WithCancel(context.WithoutCancel(ctx))
	defer cancel()
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("submitter step panicked, continuing chain: %v", r)
		}
	}()
	step.run(stepCtx)
}

// sleepUnlessCancelled waits for the given duration and returns true,
// or returns false immediately if ctx is cancelled first.
func sleepUnlessCancelled(ctx context.Context, d time.Duration) bool {
	select {
	case <-time.After(d):
		return true
	case <-ctx.Done():
		return false
	}
}

func (c *client) waitUntilRegistered(ctx context.Context) error {
	for {
		currentEpoch := c.rewardEpochTiming.EpochIndex(time.Now())

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

func (c *client) isRegistered(ctx context.Context, epoch int64) (bool, error) {
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
			logger.Errorf(
				"unexpected error while checking submitter registration: %s, retrying after %s", err, d,
			)
		},
	)

	return registered, err
}

func (c *client) waitForNextRewardEpoch(ctx context.Context, currentEpoch int64) error {
	nextEpochStart := c.rewardEpochTiming.StartTime(currentEpoch + 1)
	now := time.Now()

	// Edge case if the time passed while checking the registration means
	// we are already in the next epoch - return immediately in that case.
	if nextEpochStart.Before(now) {
		logger.Infof("submitter is not registered for voting epoch %d, checking the next epoch", currentEpoch)
		return nil
	}

	sleepTime := nextEpochStart.Sub(now)

	logger.Infof(
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
