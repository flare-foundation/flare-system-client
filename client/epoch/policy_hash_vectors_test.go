package epoch

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

// Solidity-derived vectors: a 5-voter epoch-1 policy installed through setSigningPolicy on the real
// Relay.sol (setter mode, sourceChainId 114) and on RelayMainDeployed.sol; the hashes are the
// contracts' own stored values read back via toSigningPolicyHash. Cross-checked with cast.
const (
	vectorPolicyHex         = "000500000100000d20010400000000000000000000000000000000000000000000000000000000000012341eff47bc3a10a45d4b230b5d10e37751fe6aa71800642b5ad5c4795c026514f8317c7a215e218dccd6cf00646813eb9362372eef6200f3b1dbc3f819671cba6900647e5f4552091a69125d5dfcb7b8c2659029395bdf0064e1ab8145f7e55dc933d51a18c793f901a3a0b2760064"
	vectorChainID           = 114
	vectorChainBoundHashHex = "416149338e7fa025aeea3154a8f2c39f929b1e62ecfef6b0c0d27280c5464b6a"
	vectorLegacyHashHex     = "804446241e0bde9220bf0b8fdfdbf50a65a949039a940c8f956c6705b932a799"
)

// Production-bytecode vector: eth_call setSigningPolicy on the live Coston2 Relay
// (0xa10B672D1c62e5457b17af63d4302add6A99d7dE, epoch 5956) returns the hash the
// deployed contract stores for these 153 bytes.
const (
	livePolicyHex     = "00050017440015dbf0010400000000000000000000000000000000000000000000000000000000000012341eff47bc3a10a45d4b230b5d10e37751fe6aa71800642b5ad5c4795c026514f8317c7a215e218dccd6cf00646813eb9362372eef6200f3b1dbc3f819671cba6900647e5f4552091a69125d5dfcb7b8c2659029395bdf0064e1ab8145f7e55dc933d51a18c793f901a3a0b2760064"
	liveLegacyHashHex = "1971bd4b3dc21fc8f443434bb2cde6cc7ab1b3c508f8becb453a849e272ec443"
)

// The fold is pinned against the deployed bytecode, not against its own re-expression.
func TestSigningPolicyHashLiveRelayVector(t *testing.T) {
	policy, err := hex.DecodeString(livePolicyHex)
	require.NoError(t, err)
	require.Equal(t, liveLegacyHashHex, hex.EncodeToString(SigningPolicyHash(policy)))
}

func TestPolicyHashSolidityVectors(t *testing.T) {
	policy, err := hex.DecodeString(vectorPolicyHex)
	require.NoError(t, err)
	require.Len(t, policy, 153) // 43 + 5*22, not a multiple of 32

	require.Equal(t, vectorChainBoundHashHex, hex.EncodeToString(ChainBoundSigningPolicyHash(policy, vectorChainID)))
	require.Equal(t, vectorLegacyHashHex, hex.EncodeToString(SigningPolicyHash(policy)))
}
