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

	VoterRegistration VoteRegistrationConfig `toml:"voter_registration"`
	SigningPolicy     SigningPolicyConfig    `toml:"signing_policy"`

	Ftso ProtocolConfig `toml:"ftso"`
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type VoteRegistrationConfig struct {
	Address string `toml:"address"`
}

type SigningPolicyConfig struct {
	Topic0 string `toml:"topic0"`
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
