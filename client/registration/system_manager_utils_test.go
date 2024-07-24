package registration

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func Test_abiEncode(t *testing.T) {
	t.Run("abiEncode packs arguments correctly", func(t *testing.T) {
		rewardHash := common.HexToHash("0x9d91c7d9595969e7d21783b904f8707316d6267c656b60fad0e070e9c698a672")
		encoded := abiEncode(
			big.NewInt(3),
			31337,
			&rewardHash,
			56,
		)
		encodedHashHex := hex.EncodeToString(crypto.Keccak256(encoded))
		require.Equal(t, "ac5a8c3adc6d9a499eb3bc1440a5e07b041d77b0a508bebe7429d189a41acc6a", encodedHashHex)
	})
}
