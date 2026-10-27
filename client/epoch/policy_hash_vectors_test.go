package epoch

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/policy"
)

// 5-voter epoch-1 policy via setSigningPolicy on Relay.sol (setter mode, sourceChainId 114);
// hash read back via toSigningPolicyHash, cast-checked.
const (
	vectorPolicyHex         = "000500000100000d20010400000000000000000000000000000000000000000000000000000000000012341eff47bc3a10a45d4b230b5d10e37751fe6aa71800642b5ad5c4795c026514f8317c7a215e218dccd6cf00646813eb9362372eef6200f3b1dbc3f819671cba6900647e5f4552091a69125d5dfcb7b8c2659029395bdf0064e1ab8145f7e55dc933d51a18c793f901a3a0b2760064"
	vectorChainID           = 114
	vectorChainBoundHashHex = "416149338e7fa025aeea3154a8f2c39f929b1e62ecfef6b0c0d27280c5464b6a"
)

func TestPolicyHashSolidityVector(t *testing.T) {
	encoded, err := hex.DecodeString(vectorPolicyHex)
	require.NoError(t, err)
	require.Len(t, encoded, 153) // 43 + 5*22, not a multiple of 32

	require.Equal(t, vectorChainBoundHashHex, hex.EncodeToString(policy.ChainBoundHash(vectorChainID, encoded)))
}
