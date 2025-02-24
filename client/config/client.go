package config

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/flare-foundation/flare-system-client/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type Client struct {
	DB      database.Config `toml:"db"`
	Logger  logger.Config   `toml:"logger"`
	Chain   config.Chain    `toml:"chain"`
	Metrics Metrics         `toml:"metrics"`

	Clients Clients `toml:"clients"`

	ContractAddresses config.ContractAddresses `toml:"contract_addresses"`
	Identity          Identity                 `toml:"identity"`
	Credentials       Credentials              `toml:"credentials"`

	Protocol map[string]ProtocolConfig `toml:"protocol"`

	Submit1          Submit           `toml:"submit1"`
	Submit2          Submit           `toml:"submit2"`
	SubmitSignatures SubmitSignatures `toml:"submit_signatures"`

	Finalizer Finalizer `toml:"finalizer"`

	SubmitGas   Gas `toml:"gas_submit"`
	RegisterGas Gas `toml:"gas_register"`
	RelayGas    Gas `toml:"gas_relay"`

	Rewards RewardsConfig `toml:"rewards"`
}

type Metrics struct {
	PrometheusAddress string `toml:"prometheus_address" envconfig:"PROMETHEUS_ADDRESS"`
}

type Identity struct {
	Address common.Address `toml:"address"`
}

type Credentials struct {
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

var defaultSubmitConfig = Submit{
	Enabled:          true,
	TxSubmitRetries:  1,
	TxSubmitTimeout:  10 * time.Second,
	DataFetchRetries: 1,
	DataFetchTimeout: 5 * time.Second,
}

type Submit struct {
	Enabled          bool          `toml:"enabled"`
	StartOffset      time.Duration `toml:"start_offset"` // offset from the start of the epoch
	TxSubmitRetries  int           `toml:"tx_submit_retries"`
	TxSubmitTimeout  time.Duration `toml:"tx_submit_timeout"`
	DataFetchRetries int           `toml:"data_fetch_retries"`
	DataFetchTimeout time.Duration `toml:"data_fetch_timeout"`
}

type SubmitSignatures struct {
	Submit

	Deadline time.Duration `toml:"deadline"` // from the start of the epoch, recommended to be before the end of the grace period

	MaxCycles     int           `toml:"max_cycles"`     // maximal number of query cycles after the deadline
	CycleDuration time.Duration `toml:"cycle_duration"` // minimal duration of a cycle after the deadline
}

type Clients struct {
	EnabledRegistration    bool `toml:"enabled_registration"`
	EnabledPreregistration bool `toml:"enabled_pre_registration"`
	EnabledUptimeVoting    bool `toml:"enabled_uptime_voting"`
	EnabledRewardSigning   bool `toml:"enabled_reward_signing"`
	EnabledProtocolVoting  bool `toml:"enabled_protocol_voting"`
	EnabledFinalizer       bool `toml:"enabled_finalizer"`
}

func (c *Clients) EpochClientEnabled() bool {
	return c.EnabledRegistration || c.EnabledUptimeVoting || c.EnabledRewardSigning || c.EnabledPreregistration
}

type Finalizer struct {
	StartingRewardEpoch int64  `toml:"starting_reward_epoch"`
	StartingVotingRound uint32 `toml:"starting_voting_round"`

	// How far in the past we start fetching reward epochs from the indexer at the start of the finalizer client.
	// Default is 7 days.
	StartOffset time.Duration `toml:"start_offset"`

	VoterThresholdBIPS uint16 `toml:"voter_threshold_bips"`

	// Offset from the start of the voting round.
	GracePeriodEndOffset time.Duration `toml:"grace_period_end_offset"`
}

type Gas struct {
	TxType uint8 `toml:"tx_type"` // 0 for legacy, 2 for eip-1559

	GasLimit int `toml:"gas_limit"`

	// type 0
	GasPriceMultiplier float32  `toml:"gas_price_multiplier"`
	GasPriceFixed      *big.Int `toml:"gas_price_fixed"`

	// type 2
	MaxPriorityFeePerGas *big.Int `toml:"max_priority_fee_per_gas"`
	BaseFeeMultiplier    *big.Int `toml:"base_fee_multiplier"`
	BaseFeePerGasCap     *big.Int `toml:"base_fee_per_gas_cap"`
}

type RewardsConfig struct {
	UrlPrefix string `toml:"url_prefix"`

	MinRewardWei *big.Int `toml:"min_reward"`
	MaxRewardWei *big.Int `toml:"max_reward"`

	Retries       int           `toml:"retries"`
	RetryInterval time.Duration `toml:"retry_interval"`
}

func new() *Client {
	return &Client{
		Chain: config.Chain{
			EthRPCURL: "http://localhost:9650/ext/C/rpc",
		},
		Finalizer: Finalizer{
			StartOffset:        7 * 24 * time.Hour,
			VoterThresholdBIPS: 500,
		},
		Submit1: defaultSubmitConfig,
		Submit2: defaultSubmitConfig,
		SubmitSignatures: SubmitSignatures{
			Submit: defaultSubmitConfig,
		},
		SubmitGas:   Gas{GasPriceFixed: big.NewInt(0)},
		RegisterGas: Gas{GasPriceFixed: big.NewInt(0)},
		Rewards: RewardsConfig{
			Retries:       8,
			RetryInterval: 6 * time.Hour,
		},
	}
}

// methods to satisfy config.Global interface

func (c Client) ChainConfig() config.Chain {
	return c.Chain
}

func (c Client) LoggerConfig() logger.Config {
	return c.Logger
}

func Build(cfgFileName string) (*Client, error) {
	cfg := new()
	err := config.ParseConfigFile(cfg, cfgFileName, false)
	if err != nil {
		return nil, err
	}
	err = config.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	err = validate(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// validate checks consistency of configurations.
func validate(cfg *Client) error {
	err := validateGas(&cfg.SubmitGas)
	if err != nil {
		return fmt.Errorf("validating SubmitGas: %v", err)
	}
	err = validateGas(&cfg.RegisterGas)
	if err != nil {
		return fmt.Errorf("validating RegisterGas: %v", err)
	}
	err = validateGas(&cfg.RelayGas)
	if err != nil {
		return fmt.Errorf("validating RelayGas: %v", err)
	}
	return nil
}

// validateGas checks viability of gas configurations.
func validateGas(cfg *Gas) error {
	if cfg.TxType != 0 && cfg.TxType != 2 {
		return errors.New("unsupported tx_type")
	}

	if cfg.TxType == 2 && cfg.BaseFeePerGasCap.Cmp(common.Big0) == 1 {
		logger.Warnf("a fixed BaseFeePerGasCap %v is used", cfg.BaseFeePerGasCap)
	}
	if cfg.TxType == 2 && cfg.BaseFeeMultiplier.Cmp(common.Big0) == -1 {
		return errors.New("negative base fee multiplier")
	}

	if cfg.TxType == 0 && cfg.GasPriceFixed.Cmp(common.Big0) != 0 && cfg.GasPriceMultiplier != 0.0 {
		return errors.New("only one of gas_price_fixed and gas_price_multiplier can be set to a non-zero value for type 0 transaction")
	}
	if cfg.GasPriceMultiplier != 0.0 && cfg.GasPriceMultiplier < 1 {
		return errors.New("if set, gas_price_multiplier value cannot be less than 1")
	}
	return nil
}
