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
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type CronjobConfig struct {
	Enabled   bool          `toml:"enabled"`
	Timeout   time.Duration `toml:"timeout"`
	BatchSize int64         `toml:"batch_size"`
	Delay     time.Duration `toml:"delay"`
}

type ContractAddresses struct {
	config.ContractAddresses
	Mirroring common.Address `toml:"mirroring" envconfig:"MIRRORING_CONTRACT_ADDRESS"`
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
