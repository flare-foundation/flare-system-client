package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGwei(t *testing.T) {
	tests := []struct {
		name string
		wei  *big.Int
		want string
	}{
		{"nil", nil, "unset"},
		{"zero", big.NewInt(0), "0"},
		{"integral", big.NewInt(100e9), "100"},
		{"default maximal cap", big.NewInt(5000e9), "5000"},
		{"fractional", big.NewInt(6160500000000), "6160.5"},
		{"two decimals", big.NewInt(123210000000), "123.21"},
		{"one wei", big.NewInt(1), "0.000000001"},
		{"negative", big.NewInt(-100e9), "-100"},
		// beyond float64 exactness: big.Rat must not round
		{"huge", new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil), "1000000000000000000000"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, Gwei(test.wei))
		})
	}
}
