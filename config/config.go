package config

import (
	"crypto/ecdsa"
	"errors"
	"flare-fsc/utils/credentials"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
)

const (
	CONFIG_FILE string = "config.toml"
)

var (
	GlobalConfigCallback ConfigCallback[GlobalConfig] = ConfigCallback[GlobalConfig]{}
)

type GlobalConfig interface {
	LoggerConfig() logger.Config
	ChainConfig() ChainConfig
}

type ChainConfig struct {
	ChainID   int    `toml:"chain_id" envconfig:"CHAIN_ID"`
	EthRPCURL string `toml:"eth_rpc_url" envconfig:"ETH_RPC_URL"`
	ApiKey    string `toml:"api_key" envconfig:"API_KEY"`
}

// Dial the chain node and return an ethclient.Client.
func (chain *ChainConfig) DialETH() (*ethclient.Client, error) {
	rpcURL, err := chain.getRPCURL()
	if err != nil {
		return nil, err
	}

	return ethclient.Dial(rpcURL)
}

// Get the full RPC URL which may be passed to ethclient.Dial.
// Includes API key as query param if it is configured.
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

type ContractAddresses struct {
	Submission     common.Address `toml:"submission" envconfig:"SUBMISSION_CONTRACT_ADDRESS"`
	SystemsManager common.Address `toml:"systems_manager" envconfig:"SYSTEMS_MANAGER_CONTRACT_ADDRESS"`
	VoterRegistry  common.Address `toml:"voter_registry" envconfig:"VOTER_REGISTRY_CONTRACT_ADDRESS"`
	Relay          common.Address `toml:"relay" envconfig:"RELAY_CONTRACT_ADDRESS"`
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

func ReadFileToString(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	return strings.TrimSpace(string(content)), nil
}

// Read private key from env variable or file if insecure private key handling
// is enabled (INSECURE_PRIVATE_KEYS).
func PrivateKeyFromConfig(fileName string, envString string) (pk *ecdsa.PrivateKey, err error) {
	envString = strings.TrimSpace(envString)
	fileName = strings.TrimSpace(fileName)

	var pkString string
	if len(envString) > 0 {
		pkString = envString
	} else if len(fileName) > 0 {
		allowInsecureEnv := strings.ToLower(os.Getenv("INSECURE_PRIVATE_KEYS"))
		if allowInsecureEnv == "true" {
			pkString, err = ReadFileToString(fileName)
			if err != nil {
				return nil, fmt.Errorf("error reading private key from file: %w", err)
			}
		} else {
			return nil, errors.New("private keys in files are disabled")
		}
	} else {
		return nil, errors.New("no private key specified")
	}
	pk, err = credentials.PrivateKeyFromHex(pkString)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %w", err)
	}
	return pk, nil
}
