package finalizer

import (
	clientContext "flare-tlc/client/context"
	"flare-tlc/config"
	"flare-tlc/utils/chain"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type finalizerClient struct {
	db *gorm.DB

	relayClient *relayContractClient
}

func NewFinalizerClient(ctx clientContext.ClientContext) (*finalizerClient, error) {
	cfg := ctx.Config()
	if !cfg.Voting.EnabledFinalizer {
		return nil, nil
	}

	chainCfg := cfg.ChainConfig()
	ethClient, err := chainCfg.DialETH()
	if err != nil {
		return nil, err
	}

	senderPk, err := config.ReadFileToString(cfg.Credentials.SigningPolicyPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading sender private key")
	}
	senderTxOpts, _, err := chain.CredentialsFromPrivateKey(senderPk, chainCfg.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sender register tx opts")
	}
	relayClient, err := NewRelayContractClient(
		ethClient,
		cfg.ContractAddresses.Relay,
		senderTxOpts,
	)
	if err != nil {
		return nil, err
	}

	return &finalizerClient{
		db:          ctx.DB(),
		relayClient: relayClient,
	}, nil
}

func (c *finalizerClient) Run() error {
	return nil
}
