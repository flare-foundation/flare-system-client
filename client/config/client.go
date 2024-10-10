package config

import (
	"errors"
	"flare-fsc/config"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type ClientConfig struct {
	DB      config.DBConfig     `toml:"db"`
	Logger  config.LoggerConfig `toml:"logger"`
	Chain   config.ChainConfig  `toml:"chain"`
	Metrics MetricsConfig       `toml:"metrics"`

	Clients ClientsConfig `toml:"clients"`

	ContractAddresses config.ContractAddresses `toml:"contract_addresses"`
	Identity          IdentityConfig           `toml:"identity"`
	Credentials       CredentialsConfig        `toml:"credentials"`

	Protocol map[string]ProtocolConfig `toml:"protocol"`

	Submit1          SubmitConfig           `toml:"submit1"`
	Submit2          SubmitConfig           `toml:"submit2"`
	SubmitSignatures SubmitSignaturesConfig `toml:"submit_signatures"`

	Finalizer FinalizerConfig `toml:"finalizer"`

	SubmitGas   GasConfig `toml:"gas_submit"`
	RegisterGas GasConfig `toml:"gas_register"`

	Uptime  UptimeConfig  `toml:"uptime"`
	Rewards RewardsConfig `toml:"rewards"`
}

type MetricsConfig struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type IdentityConfig struct {
	Address common.Address `toml:"address"`
}

type CredentialsConfig struct {
	// Sign all data.
	SigningPolicyPrivateKeyFile string `toml:"signing_policy_private_key_file"`
	SigningPolicyPrivateKey     string `toml:"-" envconfig:"SIGNING_POLICY_PRIVATE_KEY"`

	// Send RegisterVoter and SignNewSigningPolicy transactions.
	SystemClientSenderPrivateKeyFile string `toml:"system_client_sender_private_key_file"`
	SystemClientSenderPrivateKey     string `toml:"-" envconfig:"SYSTEM_CLIENT_SENDER_PRIVATE_KEY"`

	// Submit protocol data (submit1, submit2, submit3).
	ProtocolManagerSubmitPrivateKeyFile string `toml:"protocol_manager_submit_private_key_file"`
	ProtocolManagerSubmitPrivateKey     string `toml:"-" envconfig:"PROTOCOL_MANAGER_SUBMIT_PRIVATE_KEY"`

	// Submit protocol signatures.
	ProtocolManagerSubmitSignaturesPrivateKeyFile string `toml:"protocol_manager_submit_signatures_private_key_file"`
	ProtocolManagerSubmitSignaturesPrivateKey     string `toml:"-" envconfig:"PROTOCOL_MANAGER_SUBMIT_SIGNATURES_PRIVATE_KEY"`
}

var defaultSubmitConfig = SubmitConfig{
	Enabled:          true,
	DataFetchRetries: 1,
	DataFetchTimeout: 5 * time.Second,
}

type SubmitConfig struct {
	Enabled          bool          `toml:"enabled"`
	StartOffset      time.Duration `toml:"start_offset"` // offset from the start of the epoch
	TxSubmitRetries  int           `toml:"tx_submit_retries"`
	DataFetchRetries int           `toml:"data_fetch_retries"`
	DataFetchTimeout time.Duration `toml:"data_fetch_timeout"`
}

type SubmitSignaturesConfig struct {
	SubmitConfig

	MaxRounds int `toml:"max_rounds"`
}

type ClientsConfig struct {
	EnabledRegistration   bool `toml:"enabled_registration"`
	EnabledUptimeVoting   bool `toml:"enabled_uptime_voting"`
	EnabledRewardSigning  bool `toml:"enabled_reward_signing"`
	EnabledProtocolVoting bool `toml:"enabled_protocol_voting"`
	EnabledFinalizer      bool `toml:"enabled_finalizer"`
}

func (c *ClientsConfig) EpochClientEnabled() bool {
	return c.EnabledRegistration || c.EnabledUptimeVoting || c.EnabledRewardSigning
}

type FinalizerConfig struct {
	StartingRewardEpoch int64  `toml:"starting_reward_epoch"`
	StartingVotingRound uint32 `toml:"starting_voting_round"`

	// How far in the past we start fetching reward epochs from the indexer at the start of the finalizer client.
	// Default is 7 days.
	StartOffset time.Duration `toml:"start_offset"`

	VoterThresholdBIPS uint16 `toml:"voter_threshold_bips"`

	// Offset from the start of the voting round.
	GracePeriodEndOffset time.Duration `toml:"grace_period_end_offset"`
}

type GasConfig struct {
	GasPriceMultiplier float32  `toml:"gas_price_multiplier"`
	GasPriceFixed      *big.Int `toml:"gas_price_fixed"`
	GasLimit           int      `toml:"gas_limit"`

	MaxPriorityFeePerGas *big.Int `toml:"max_priority_fee_per_gas"`
	BaseFeePerGasCap     *big.Int `toml:"base_fee_per_gas_cap"`
}

type UptimeConfig struct {
	SigningWindow int64 `toml:"signing_window"`
}

type RewardsConfig struct {
	PathPrefix    string `toml:"hash_path_prefix"`
	SigningWindow int64  `toml:"signing_window"`
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
		Submit1: defaultSubmitConfig,
		Submit2: defaultSubmitConfig,
		SubmitSignatures: SubmitSignaturesConfig{
			SubmitConfig: defaultSubmitConfig,
		},
		Uptime: UptimeConfig{
			SigningWindow: 2,
		},
		Rewards: RewardsConfig{
			SigningWindow: 2,
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
	err = validateConfig(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func validateConfig(cfg *ClientConfig) error {
	err := validateGasConfig(&cfg.SubmitGas)
	if err != nil {
		return err
	}
	err = validateGasConfig(&cfg.RegisterGas)
	if err != nil {
		return err
	}
	return nil
}

func validateGasConfig(cfg *GasConfig) error {
	if cfg.GasPriceFixed.Cmp(common.Big0) != 0 && cfg.GasPriceMultiplier != 0.0 {
		return errors.New("only one of gas_price_fixed and gas_price_multiplier can be set to a non-zero value")
	}
	return nil
}
