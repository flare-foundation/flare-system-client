package protocol

import (
	clientContext "flare-tlc/client/context"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/system"

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

	var subProtocols []*SubProtocol
	for _, protocol := range cfg.Protocol {
		subProtocols = append(subProtocols, NewSubProtocol(protocol))
	}

	pc := &ProtocolClient{
		eth:             cl,
		protocolContext: protocolContext,
		subProtocols:    subProtocols,
		votingEpoch:     votingEpoch,
		systemsManager:  systemsManager,
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

func (c *ProtocolClient) Run() error {
	if c.submitter1 != nil {
		go Run(c.submitter1)
	}
	if c.submitter2 != nil {
		go Run(c.submitter2)
	}
	if c.signatureSubmitter != nil {
		go Run(c.signatureSubmitter)
	}

	return nil
}
