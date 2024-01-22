package protocol

import (
	clientContext "flare-tlc/client/context"
	"flare-tlc/client/shared"
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

	votingEpoch   *utils.Epoch
	systemManager *system.FlareSystemManager
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

	systemManager, err := system.NewFlareSystemManager(cfg.ContractAddresses.SystemManager, cl)
	if err != nil {
		return nil, errors.Wrap(err, "error creating system manager contract")
	}

	votingEpoch, err := shared.VotingEpochFromChain(systemManager)
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
		systemManager:   systemManager,
	}

	selectors := newContractSelectors()

	pc.submitter1 = newSubmitter(cl, protocolContext, votingEpoch,
		&cfg.Submit1, selectors.submit1, subProtocols, 0, "submit1")
	pc.submitter2 = newSubmitter(cl, protocolContext, votingEpoch,
		&cfg.Submit2, selectors.submit2, subProtocols, -1, "submit2")
	pc.signatureSubmitter = newSignatureSubmitter(cl, protocolContext, votingEpoch,
		&cfg.SubmitSignatures, selectors.submitSignatures, subProtocols)

	return pc, nil
}

func (c *ProtocolClient) Run() error {
	go Run(c.submitter1)
	go Run(c.submitter2)
	go Run(c.signatureSubmitter)

	return nil
}
