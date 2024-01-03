package config

import (
	"flare-tlc/config"
	"time"
)

type ClientConfig struct {
	DB                config.DBConfig          `toml:"db"`
	Logger            config.LoggerConfig      `toml:"logger"`
	Chain             config.ChainConfig       `toml:"chain"`
	Metrics           MetricsConfig            `toml:"metrics"`
	ContractAddresses config.ContractAddresses `toml:"contract_addresses"`

	VoterRegistration VoterRegistrationConfig `toml:"voter_registration"`

	Ftso ProtocolConfig `toml:"ftso"`
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type VoterRegistrationConfig struct {
	EpochStart  time.Time     `toml:"epoch_start"`  // Temporary, get from some contract
	EpochPeriod time.Duration `toml:"epoch_period"` // Temporary, get from some contract
	Topic0      string        `toml:"topic0"`
	Address     string        `toml:"address"`
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
