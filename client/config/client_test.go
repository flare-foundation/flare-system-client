package config

import (
	"math/big"
	"testing"
	"time"

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
				BaseFeeMultiplier:     big.NewInt(-1),
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
		Enabled:          true,
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

	t.Run("disabled submitter is not validated", func(t *testing.T) {
		c := base()
		c.Submit1.Enabled = false
		c.Submit1.TxSubmitTimeout = 0 // nonsense, but disabled so ignored
		require.NoError(t, c.validateSubmitters())
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

	t.Run("ordering check skipped when submit2 disabled", func(t *testing.T) {
		c := base()
		c.Submit2.Enabled = false
		c.SubmitSignatures.StartOffset = 1 * time.Second
		c.SubmitSignatures.Deadline = 30 * time.Second
		require.NoError(t, c.validateSubmitters())
	})
}
