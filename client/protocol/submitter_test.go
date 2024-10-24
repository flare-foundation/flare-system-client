package protocol

import (
	"math/big"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/config"
)

func TestGasConfigForAttempt(t *testing.T) {
	tests := []struct {
		name     string
		cfg      config.Gas
		ri       int
		expected config.Gas
	}{
		{
			name: "retry 0",
			cfg: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
			ri: 0,
			expected: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
		},
		{
			name: "retry 1",
			cfg: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
			ri: 1,
			expected: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.5,
			},
		},
		{
			name: "retry 1 - no config",
			cfg: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 0,
			},
			ri: 1,
			expected: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.5,
			},
		},
		{
			name: "retry 2",
			cfg: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
			ri: 2,
			expected: config.Gas{
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 2.25,
			},
		},
		{
			name: "retry 1 - fixed",
			cfg: config.Gas{
				GasPriceFixed:      big.NewInt(100),
				GasPriceMultiplier: 0,
			},
			ri: 1,
			expected: config.Gas{
				GasPriceFixed:      big.NewInt(100),
				GasPriceMultiplier: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gasConfigForAttempt(&tt.cfg, tt.ri)
			if got.GasPriceFixed.Cmp(tt.expected.GasPriceFixed) != 0 {
				t.Errorf("GasPriceFixed = %v, want %v", got.GasPriceFixed, tt.expected.GasPriceFixed)
			}
			if got.GasPriceMultiplier != tt.expected.GasPriceMultiplier {
				t.Errorf("GasPriceMultiplier = %v, want %v", got.GasPriceMultiplier, tt.expected.GasPriceMultiplier)
			}
		})
	}
}
