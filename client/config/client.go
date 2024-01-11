package config

import (
	"flare-tlc/config"
)

type ClientConfig struct {
	DB                config.DBConfig          `toml:"db"`
	Logger            config.LoggerConfig      `toml:"logger"`
	Chain             config.ChainConfig       `toml:"chain"`
	Metrics           MetricsConfig            `toml:"metrics"`
	ContractAddresses config.ContractAddresses `toml:"contract_addresses"`

	Credentials CredentialsConfig `toml:"credentials"`
	Voting      VotingConfig      `toml:"voting"`

	Protocol ProtocolConfig `toml:"protocol"`
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type CredentialsConfig struct {
	RunScheduler                      bool   `toml:"run_scheduler"`
	IdentityAddress                   string `toml:"identity_address"`
	SystemManagerSenderPrivateKeyFile string `toml:"system_manager_sender_private_key_file"`
	SigningPolicyPrivateKeyFile       string `toml:"signing_policy_private_key_file"`
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
