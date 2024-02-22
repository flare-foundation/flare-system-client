package config

import (
	"flare-tlc/config"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type ClientConfig struct {
	DB      config.DBConfig     `toml:"db"`
	Logger  config.LoggerConfig `toml:"logger"`
	Chain   config.ChainConfig  `toml:"chain"`
	Metrics MetricsConfig       `toml:"metrics"`

	Clients VotingClientsConfig `toml:"clients"`

	ContractAddresses config.ContractAddresses `toml:"contract_addresses"`
	Identity          IdentityConfig           `toml:"identity"`
	Credentials       CredentialsConfig        `toml:"credentials"`

	Protocol map[string]ProtocolConfig `toml:"protocol"`

	Submit1          SubmitConfig           `toml:"submit1"`
	Submit2          SubmitConfig           `toml:"submit2"`
	SubmitSignatures SubmitSignaturesConfig `toml:"submit_signatures"`

	Finalizer FinalizerConfig `toml:"finalizer"`
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type IdentityConfig struct {
	Address common.Address `toml:"address"`
}

type CredentialsConfig struct {
	// Sign all data
	SigningPolicyPrivateKeyFile string `toml:"signing_policy_private_key_file"`
	SigningPolicyPrivateKey     string `envconfig:"SIGNING_POLICY_PRIVATE_KEY"`

	// Send RegisterVoter and SignNewSigningPolicy transactions
	SystemClientSenderPrivateKeyFile string `toml:"system_client_sender_private_key_file"`
	SystemClientSenderPrivateKey     string `envconfig:"SYSTEM_CLIENT_SENDER_PRIVATE_KEY"`

	// Submit protocol data (submit1, submit2, submit3)
	ProtocolManagerSubmitPrivateKeyFile string `toml:"protocol_manager_submit_private_key_file"`
	ProtocolManagerSubmitPrivateKey     string `envconfig:"PROTOCOL_MANAGER_SUBMIT_PRIVATE_KEY"`

	// Submit protocol signatures
	ProtocolManagerSubmitSignaturesPrivateKeyFile string `toml:"protocol_manager_submit_signatures_private_key_file"`
	ProtocolManagerSubmitSignaturesPrivateKey     string `envconfig:"PROTOCOL_MANAGER_SUBMIT_SIGNATURES_PRIVATE_KEY"`
}

type SubmitConfig struct {
	// offset from the start of the epoch
	StartOffset     time.Duration `toml:"start_offset"`
	TxSubmitRetries int           `toml:"tx_submit_retries"`
}

type SubmitSignaturesConfig struct {
	SubmitConfig

	DataFetchRetries int `toml:"data_fetch_retries"`
	MaxRounds        int `toml:"max_rounds"`
}

type VotingClientsConfig struct {
	EnabledRegistration   bool `toml:"enabled_registration"`
	EnabledProtocolVoting bool `toml:"enabled_protocol_voting"`
	EnabledFinalizer      bool `toml:"enabled_finalizer"`
}

type FinalizerConfig struct {
	StartingRewardEpoch int64  `toml:"starting_reward_epoch"`
	StartingVotingRound uint32 `toml:"starting_voting_round"`

	// how far in the past we start fetching reward epochs from the indexer at the start of the finalizer client
	// default is 7 days
	StartOffset time.Duration `toml:"start_offset"`

	VoterThresholdBIPS uint16 `toml:"voter_threshold_bips"`

	// Offset from the start of the voting round
	GracePeriodEndOffset time.Duration `toml:"grace_period_end_offset"`
}

func newConfig() *ClientConfig {
	return &ClientConfig{
		Chain: config.ChainConfig{
			EthRPCURL: "http://localhost:9650/ext/C/rpc",
		},
		Finalizer: FinalizerConfig{
			StartOffset:        7 * 24 * time.Hour,
			VoterThresholdBIPS: 500,
		},
	}
}

func (c ClientConfig) LoggerConfig() config.LoggerConfig {
	return c.Logger
}

func (c ClientConfig) ChainConfig() config.ChainConfig {
	return c.Chain
}

func BuildConfig(cfgFileName string) (*ClientConfig, error) {
	cfg := newConfig()
	err := config.ParseConfigFile(cfg, cfgFileName, false)
	if err != nil {
		return nil, err
	}
	err = config.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
