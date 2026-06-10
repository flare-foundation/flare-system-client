package config

import (
	"math/big"
	"testing"

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
