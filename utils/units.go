package utils

import (
	"math/big"
	"strings"
)

var weiPerGwei = big.NewInt(1e9)

// Gwei renders a wei amount in gwei for logs, trimming trailing zeros; "unset" if nil.
// Exact: big.Rat, not float division.
func Gwei(wei *big.Int) string {
	if wei == nil {
		return "unset"
	}
	s := new(big.Rat).SetFrac(wei, weiPerGwei).FloatString(9)
	return strings.TrimSuffix(strings.TrimRight(s, "0"), ".")
}
