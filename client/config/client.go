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

func Build(cfgFileName string) (*Client, error) {
	cfg := defaultConfig()
	err := config.ParseConfigFile(cfg, cfgFileName, false)
	if err != nil {
		return nil, err
	}
	err = config.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	err = cfg.validate()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// methods to satisfy config.Global interface

func (c Client) ChainConfig() config.Chain {
	return c.Chain
}

func (c Client) LoggerConfig() logger.Config {
	return c.Logger
}

func defaultConfig() *Client {
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
			Submit:   defaultSubmitConfig,
			Deadline: 60 * time.Second,
		},
		SubmitGas:   DefaultGas(),
		RelayGas:    DefaultGas(),
		RegisterGas: DefaultGas(),
		Rewards: RewardsConfig{
			Retries:       8,
			RetryInterval: 6 * time.Hour,
		},
	}
}

// validate checks consistency of configurations.
func (c *Client) validate() error {
	err := c.SubmitGas.validate()
	if err != nil {
		return fmt.Errorf("validating SubmitGas: %v", err)
	}
	err = c.RegisterGas.validate()
	if err != nil {
		return fmt.Errorf("validating RegisterGas: %v", err)
	}
	err = c.RelayGas.validate()
	if err != nil {
		return fmt.Errorf("validating RelayGas: %v", err)
	}
	err = c.validateContracts()
	if err != nil {
		return fmt.Errorf("validating contracts: %v", err)
	}
	return nil
}

func (cfg *Client) validateContracts() error {
	var zeroAddress common.Address

	if cfg.ContractAddresses.Submission == zeroAddress {
		return errors.New("submission contract address not set")
	}
	if cfg.ContractAddresses.SystemsManager == zeroAddress {
		return errors.New("systems_manager contract address not set")
	}

	if cfg.Clients.EnabledPreregistration && cfg.ContractAddresses.VoterPreRegistry == zeroAddress {
		return errors.New("pre-registration enabled but voter_preregistry contract address not set")
	}

	if cfg.Clients.EnabledRegistration && cfg.ContractAddresses.VoterRegistry == zeroAddress {
		return errors.New("registration enabled but voter_registry contract address not set")
	}

	if cfg.Clients.EnabledFinalizer && cfg.ContractAddresses.Relay == zeroAddress {
		return errors.New("finalizer enabled but relay contract address not set")
	}

	return nil
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

// Gas dictates how gas for the transaction is set.
//
// TxType decides the type of the transaction. The available options are 0 and 2.
// It is recommended to use type 2.
//
// For type 0, there are two options. If GasPriceFixed is set, it is used as it is.
// ONLY SET GasPriceFixed IF YOU KNOW WHAT YOU ARE DOING.
// If GasPriceFixed is not set, a gas price recommended by the node is multiplied by GasPriceMultiplier.
//
// For type 2, two values have to be set. GasFeeCap and GasTipCap.
// The actual gas is minimum of GasFeeCap and GasTipCap + baseFee. BaseFee is defined by the block.
// GasTipCap is set to MaxPriorityMultiplier times the estimation of the baseFee and is capped by
// MaximalMaxPriorityFee and MinimalMaxPriorityFee from above and below, respectively.
// GasFeeCap as a sum GasTipCap and baseFeeCap.
// If BaseFeePerGasCap is set (IT IS RECOMMENDED TO LEAVE IT UNSET), it is used as baseFeeCap, otherwise
// baseFeeCap is baseFee estimated by the node multiplied by BaseFeeMultiplier.
type Gas struct {
	TxType uint8 `toml:"tx_type"` // 0 for legacy, 2 for eip-1559

	GasLimit int `toml:"gas_limit"` // LEAVE UNSET OR SET TO 0 UNLESS YOU KNOW WHAT YOU ARE DOING.

	// type 0
	GasPriceMultiplier float32  `toml:"gas_price_multiplier"`
	GasPriceFixed      *big.Int `toml:"gas_price_fixed"`

	// type 2
	MaxPriorityMultiplier *big.Int `toml:"max_priority_fee_multiplier"`
	MaximalMaxPriorityFee *big.Int `toml:"maximal_max_priority_fee"`
	MinimalMaxPriorityFee *big.Int `toml:"minimal_max_priority_fee"`

	BaseFeeMultiplier *big.Int `toml:"base_fee_multiplier"`
	BaseFeePerGasCap  *big.Int `toml:"base_fee_per_gas_cap"` // LEAVE UNSET UNLESS YOU KNOW WHAT YOU ARE DOING.
}

// DefaultGas
func DefaultGas() Gas {
	return Gas{
		TxType: 2,

		GasLimit: 0,

		MaxPriorityMultiplier: big.NewInt(2),
		MinimalMaxPriorityFee: big.NewInt(100e9),
		MaximalMaxPriorityFee: big.NewInt(5000e9),

		BaseFeeMultiplier: big.NewInt(4),
	}
}

// EnforceMaxPriorityFeeCaps returns capped fee.
// A new value is returned and the underlying value of the fee is unchanged/
func (g *Gas) EnforceMaxPriorityFeeCaps(fee *big.Int) *big.Int {
	out := new(big.Int)

	if g.MaximalMaxPriorityFee.Cmp(fee) == -1 {
		out.Set(g.MaximalMaxPriorityFee)
	} else if g.MinimalMaxPriorityFee.Cmp(fee) == 1 {
		out.Set(g.MinimalMaxPriorityFee)
	} else {
		out.Set(fee)
	}

	return out
}

// validate checks viability of gas configurations.
func (g *Gas) validate() error {
	if g.GasPriceMultiplier != 0.0 && g.GasPriceMultiplier < 1 {
		return errors.New("if set, gas_price_multiplier value cannot be less than 1")
	}

	switch g.TxType {
	case 0:

		if g.GasPriceFixed.Cmp(common.Big0) != 0 && g.GasPriceMultiplier != 0.0 {
			return errors.New("only one of gas_price_fixed and gas_price_multiplier can be set to a non-zero value for type 0 transaction")
		}

	case 2:
		if g.BaseFeeMultiplier.Cmp(common.Big0) == -1 {
			return errors.New("negative base fee multiplier")
		}

		if g.MaximalMaxPriorityFee.Cmp(g.MinimalMaxPriorityFee) == -1 {
			return errors.New("MaximalMaxPriorityFee cannot be less than MinimalMaxPriorityFee")
		}

		if g.MinimalMaxPriorityFee.Cmp(common.Big0) == -1 {
			return errors.New("negative MinimalMaxPriorityFee")
		}
	default:
		return errors.New("unsupported tx_type")
	}

	return nil
}

type RewardsConfig struct {
	UrlPrefix string `toml:"url_prefix"`

	MinRewardWei *big.Int `toml:"min_reward"`
	MaxRewardWei *big.Int `toml:"max_reward"`

	Retries       int           `toml:"retries"`
	RetryInterval time.Duration `toml:"retry_interval"`
}
