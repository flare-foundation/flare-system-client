package epoch

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const (
	testSignerKey = "b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291"
	testChainID   = int64(14)
	testEpochID   = uint32(424)
)

var testIdentity = common.HexToAddress("0x1234567890AbcdEF1234567890aBcdef12345678")

func testSigner(t *testing.T) *registryContractClientImpl {
	t.Helper()

	key, err := crypto.HexToECDSA(testSignerKey)
	require.NoError(t, err)

	return &registryContractClientImpl{signerPrivateKey: key, chainID: testChainID}
}

// abi.encode(uint256, uint32, address) is three left-padded 32-byte words; built by hand
// so the assertion does not go through the same registratorArguments the code under test uses.
func expectedMessage() []byte {
	message := make([]byte, 0, 96)
	message = append(message, common.LeftPadBytes(big.NewInt(testChainID).Bytes(), 32)...)
	message = append(message, common.LeftPadBytes(new(big.Int).SetUint64(uint64(testEpochID)).Bytes(), 32)...)
	message = append(message, common.LeftPadBytes(testIdentity.Bytes(), 32)...)
	return message
}

func TestCreateSignatureMessageForm(t *testing.T) {
	r := testSigner(t)

	got, err := r.createSignature(testEpochID, testIdentity)
	require.NoError(t, err)

	message := expectedMessage()
	require.Len(t, message, 96)

	want, err := crypto.Sign(accounts.TextHash(crypto.Keccak256(message)), r.signerPrivateKey)
	require.NoError(t, err)
	require.Equal(t, want, got, "registration signature must cover (chainID, epochID, address)")
}

// crypto.Sign is deterministic (RFC 6979), so any change to the message form breaks this.
func TestCreateSignatureGoldenVector(t *testing.T) {
	const wantHex = "ff248683056be9426fcd072191c5c785f4b0bf42a560fe133311212609001f13" +
		"2ee3cdf7db75c91ed95ccce3deb522007f5a640ff2c462c8de1b53be678af463" + "01"

	got, err := testSigner(t).createSignature(testEpochID, testIdentity)
	require.NoError(t, err)
	require.Equal(t, wantHex, hex.EncodeToString(got))
}

// chainID is part of the signed message, so a different chain must yield a different signature.
func TestCreateSignatureDependsOnChainID(t *testing.T) {
	r := testSigner(t)

	onFlare, err := r.createSignature(testEpochID, testIdentity)
	require.NoError(t, err)

	r.chainID = 19
	onSongbird, err := r.createSignature(testEpochID, testIdentity)
	require.NoError(t, err)

	require.NotEqual(t, onFlare, onSongbird)
}
