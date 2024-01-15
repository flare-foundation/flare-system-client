package config

import (
	"flare-tlc/config"
	"time"
)

type ClientConfig struct {
	DB      config.DBConfig     `toml:"db"`
	Logger  config.LoggerConfig `toml:"logger"`
	Chain   config.ChainConfig  `toml:"chain"`
	Metrics MetricsConfig       `toml:"metrics"`

	Voting VotingConfig `toml:"voting"`

	ContractAddresses config.ContractAddresses `toml:"contract_addresses"`
	Identity          IdentityConfig           `toml:"identity"`
	Credentials       CredentialsConfig        `toml:"credentials"`

	Protocol map[string]ProtocolConfig `toml:"protocol"`

	Submit1          SubmitConfig           `toml:"submit1"`
	Submit2          SubmitConfig           `toml:"submit2"`
	SubmitSignatures SubmitSignaturesConfig `toml:"submit_signatures"`
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type IdentityConfig struct {
	Address string `toml:"address"`
}

type CredentialsConfig struct {
	// Sign all data
	SigningPolicyPrivateKeyFile string `toml:"signing_policy_private_key_file"`

	// Send RegisterVoter and SignNewSigningPolicy transactions
	SystemClientSenderPrivateKeyFile string `toml:"system_client_sender_private_key_file"`

	// Submit protocol data (submit1, submit2, submit3)
	ProtocolManagerSubmitPrivateKeyFile string `toml:"protocol_manager_submit_private_key_file"`

	// Submit protocol signatures
	ProtocolManagerSubmitSignaturesPrivateKeyFile string `toml:"protocol_manager_submit_signatures_private_key_file"`
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

type VotingConfig struct {
	EnabledRegistration   bool `toml:"enabled_registration"`
	EnabledProtocolVoting bool `toml:"enabled_protocol_voting"`
}

func newConfig() *ClientConfig {
	return &ClientConfig{
		Chain: config.ChainConfig{
			NodeURL: "http://localhost:9650/",
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
