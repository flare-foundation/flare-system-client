package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
)

const (
	CONFIG_FILE string = "config.toml"
)

var (
	GlobalConfigCallback ConfigCallback[GlobalConfig] = ConfigCallback[GlobalConfig]{}
)

type GlobalConfig interface {
	LoggerConfig() LoggerConfig
	ChainConfig() ChainConfig
}

type LoggerLevel string

type LoggerConfig struct {
	Level       string `toml:"level"` // valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (zap)
	File        string `toml:"file"`
	MaxFileSize int    `toml:"max_file_size"` // In megabytes
	Console     bool   `toml:"console"`
}

type DBConfig struct {
	Host       string `toml:"host" envconfig:"DB_HOST"`
	Port       int    `toml:"port" envconfig:"DB_PORT"`
	Database   string `toml:"database" envconfig:"DB_DATABASE"`
	Username   string `toml:"username" envconfig:"DB_USERNAME"`
	Password   string `toml:"password" envconfig:"DB_PASSWORD"`
	LogQueries bool   `toml:"log_queries"`
}

type ChainConfig struct {
	NodeURL         string `toml:"node_url" envconfig:"CHAIN_NODE_URL"`
	ChainAddressHRP string `toml:"address_hrp" envconfig:"CHAIN_ADDRESS_HRP"`
	ChainID         int    `toml:"chain_id" envconfig:"CHAIN_ID"`
	EthRPCURL       string `toml:"eth_rpc_url" envconfig:"ETH_RPC_URL"`
	ApiKey          string `toml:"api_key" envconfig:"API_KEY"`
	PrivateKeyFile  string `toml:"private_key_file" envconfig:"PRIVATE_KEY_FILE"`
}

func (cfg ChainConfig) GetPrivateKey() (string, error) {
	content, err := os.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		return "", fmt.Errorf("error opening private key file: %w", err)
	}
	return strings.TrimSpace(string(content)), nil
}

// Dial the chain node and return an ethclient.Client.
func (chain *ChainConfig) DialETH() (*ethclient.Client, error) {
	rpcURL, err := chain.getRPCURL()
	if err != nil {
		return nil, err
	}

	return ethclient.Dial(rpcURL)
}

// Get the full RPC URL which may be passed to ethclient.Dial. Includes API key
// as query param if it is configured.
func (chain *ChainConfig) getRPCURL() (string, error) {
	u, err := url.Parse(chain.EthRPCURL)
	if err != nil {
		return "", err
	}

	if chain.ApiKey == "" {
		return u.String(), nil
	}

	q := u.Query()
	q.Set("x-apikey", chain.ApiKey)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

type EpochConfig struct {
	First int64 `toml:"first" envconfig:"EPOCH_FIRST"`
}

type ContractAddresses struct {
	Voting common.Address `toml:"voting" envconfig:"VOTING_CONTRACT_ADDRESS"`
}

func ParseConfigFile(cfg interface{}, fileName string, allowMissing bool) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		if allowMissing {
			return nil
		} else {
			return fmt.Errorf("error opening config file: %w", err)
		}
	}

	_, err = toml.Decode(string(content), cfg)
	if err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}
	return nil
}

func ReadEnv(cfg interface{}) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("error reading env config: %w", err)
	}
	return nil
}
