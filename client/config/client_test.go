package config

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestGasValidate(t *testing.T) {
	tests := []struct {
		name    string
		gas     Gas
		wantErr bool
	}{
		{
			// regression: this previously panicked on a nil *big.Int
			name:    "type 0 with no gas price settings",
			gas:     Gas{TxType: 0},
			wantErr: false,
		},
		{
			name:    "type 0 with multiplier only",
			gas:     Gas{TxType: 0, GasPriceMultiplier: 1.5},
			wantErr: false,
		},
		{
			name:    "type 0 with fixed price only",
			gas:     Gas{TxType: 0, GasPriceFixed: big.NewInt(50_000_000_000)},
			wantErr: false,
		},
		{
			name:    "type 0 with both fixed price and multiplier",
			gas:     Gas{TxType: 0, GasPriceFixed: big.NewInt(50_000_000_000), GasPriceMultiplier: 1.5},
			wantErr: true,
		},
		{
			name:    "multiplier less than 1",
			gas:     Gas{TxType: 0, GasPriceMultiplier: 0.5},
			wantErr: true,
		},
		{
			name:    "type 2 defaults",
			gas:     DefaultGas(),
			wantErr: false,
		},
		{
			name: "type 2 with negative base fee multiplier",
			gas: Gas{
				TxType:                2,
				BaseFeeMultiplier:     -1,
				MinimalMaxPriorityFee: DefaultMinimalMaxPriorityFee,
				MaximalMaxPriorityFee: DefaultMaximalMaxPriorityFee,
			},
			wantErr: true,
		},
		{
			name: "type 2 with max priority fee below min",
			gas: Gas{
				TxType:                2,
				BaseFeeMultiplier:     DefaultBaseFeeMultiplier,
				MinimalMaxPriorityFee: big.NewInt(100),
				MaximalMaxPriorityFee: big.NewInt(50),
			},
			wantErr: true,
		},
		{
			name:    "unsupported tx type",
			gas:     Gas{TxType: 1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.gas.validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func validSubmit() Submit {
	return Submit{
		StartOffset:      20 * time.Second,
		TxSubmitRetries:  2,
		TxSubmitTimeout:  5 * time.Second,
		DataFetchRetries: 1,
		DataFetchTimeout: 5 * time.Second,
	}
}

func TestSubmitValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(s *Submit)
		wantErr bool
	}{
		{"valid", func(*Submit) {}, false},
		{"zero start offset is allowed", func(s *Submit) { s.StartOffset = 0 }, false},
		{"negative start offset", func(s *Submit) { s.StartOffset = -time.Second }, true},
		{"zero retries", func(s *Submit) { s.TxSubmitRetries = 0 }, true},
		{"negative retries", func(s *Submit) { s.TxSubmitRetries = -1 }, true},
		{"non-positive submit timeout", func(s *Submit) { s.TxSubmitTimeout = 0 }, true},
		{"zero data fetch retries", func(s *Submit) { s.DataFetchRetries = 0 }, true},
		{"non-positive data fetch timeout", func(s *Submit) { s.DataFetchTimeout = -time.Second }, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := validSubmit()
			tt.mutate(&s)
			err := s.validate("submit")
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func validSubmitSignatures() SubmitSignatures {
	return SubmitSignatures{
		Submit:        validSubmit(),
		Deadline:      58 * time.Second,
		MaxCycles:     5,
		CycleDuration: 2 * time.Second,
	}
}

func TestSubmitSignaturesValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(s *SubmitSignatures)
		wantErr bool
	}{
		{"valid", func(*SubmitSignatures) {}, false},
		{"inherits Submit checks", func(s *SubmitSignatures) { s.TxSubmitTimeout = 0 }, true},
		{"deadline equal to start offset", func(s *SubmitSignatures) { s.Deadline = s.StartOffset }, true},
		{"deadline before start offset", func(s *SubmitSignatures) { s.Deadline = s.StartOffset - time.Second }, true},
		{"negative max cycles", func(s *SubmitSignatures) { s.MaxCycles = -1 }, true},
		{"zero max cycles is allowed", func(s *SubmitSignatures) { s.MaxCycles = 0 }, false},
		{"negative cycle duration", func(s *SubmitSignatures) { s.CycleDuration = -time.Second }, true},
		{"zero cycle duration is allowed", func(s *SubmitSignatures) { s.CycleDuration = 0 }, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := validSubmitSignatures()
			tt.mutate(&s)
			err := s.validate()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateSubmitters(t *testing.T) {
	base := func() *Client {
		return &Client{
			Submit1:          validSubmit(),
			Submit2:          validSubmit(),
			SubmitSignatures: validSubmitSignatures(),
		}
	}

	t.Run("valid", func(t *testing.T) {
		require.NoError(t, base().validateSubmitters())
	})

	t.Run("enabled submitter is validated", func(t *testing.T) {
		c := base()
		c.Submit1.TxSubmitTimeout = 0
		require.Error(t, c.validateSubmitters())
	})

	t.Run("signatures must not run before reveal", func(t *testing.T) {
		c := base()
		c.Submit2.StartOffset = 45 * time.Second
		c.SubmitSignatures.StartOffset = 20 * time.Second // before submit2
		c.SubmitSignatures.Deadline = 30 * time.Second    // keep deadline > its start offset
		require.Error(t, c.validateSubmitters())
	})
}

func TestGasValidateType2(t *testing.T) {
	valid := DefaultGas()
	valid.MaxPriorityMultiplier = 1.5
	valid.BaseFeeMultiplier = 2.25
	require.NoError(t, valid.validate())

	zeroBase := DefaultGas()
	zeroBase.BaseFeeMultiplier = 0
	require.ErrorContains(t, zeroBase.validate(), "base_fee_multiplier")

	negativePriority := DefaultGas()
	negativePriority.MaxPriorityMultiplier = -1
	require.ErrorContains(t, negativePriority.validate(), "max_priority_fee_multiplier")

	infBase := DefaultGas()
	infBase.BaseFeeMultiplier = Multiplier(math.Inf(1))
	require.ErrorContains(t, infBase.validate(), "base_fee_multiplier")

	nanPriority := DefaultGas()
	nanPriority.MaxPriorityMultiplier = Multiplier(math.NaN())
	require.ErrorContains(t, nanPriority.validate(), "max_priority_fee_multiplier")

	nanPrice := DefaultGas()
	nanPrice.GasPriceMultiplier = float32(math.NaN())
	require.ErrorContains(t, nanPrice.validate(), "gas_price_multiplier")

	swappedCaps := DefaultGas()
	swappedCaps.MaximalMaxPriorityFee = big.NewInt(1)
	require.ErrorContains(t, swappedCaps.validate(), "maximal_max_priority_fee")
}

func TestGasCopyAndDefaultMultipliers(t *testing.T) {
	unset := Gas{TxType: 2}
	got := unset.CopyAndDefault()
	require.Equal(t, Multiplier(DefaultMaxPriorityMultiplier), got.MaxPriorityMultiplier)
	require.Equal(t, Multiplier(DefaultBaseFeeMultiplier), got.BaseFeeMultiplier)

	fractional := Gas{TxType: 2, MaxPriorityMultiplier: 1.5, BaseFeeMultiplier: 2.25}
	got = fractional.CopyAndDefault()
	require.Equal(t, Multiplier(1.5), got.MaxPriorityMultiplier)
	require.Equal(t, Multiplier(2.25), got.BaseFeeMultiplier)
}

func TestGasTOMLMultipliers(t *testing.T) {
	var g Gas
	_, err := toml.Decode(`
tx_type = 2
max_priority_fee_multiplier = 1.5
base_fee_multiplier = 2.25
maximal_max_priority_fee = "5000000000000"
`, &g)
	require.NoError(t, err)
	require.Equal(t, Multiplier(1.5), g.MaxPriorityMultiplier)
	require.Equal(t, Multiplier(2.25), g.BaseFeeMultiplier)
	require.Zero(t, g.MaximalMaxPriorityFee.Cmp(big.NewInt(5_000_000_000_000)))

	// Backward compatibility with the old *big.Int format: quoted strings and bare ints.
	var gOld Gas
	_, err = toml.Decode(`
max_priority_fee_multiplier = "2"
base_fee_multiplier = 4
`, &gOld)
	require.NoError(t, err)
	require.Equal(t, Multiplier(2), gOld.MaxPriorityMultiplier)
	require.Equal(t, Multiplier(4), gOld.BaseFeeMultiplier)

	_, err = toml.Decode(`base_fee_multiplier = "abc"`, &g)
	require.ErrorContains(t, err, "invalid multiplier")
}
