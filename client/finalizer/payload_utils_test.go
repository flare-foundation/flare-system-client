package finalizer

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare-system-client/client/protocol"
	"github.com/flare-foundation/go-flare-common/pkg/payload"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

const (
	type0PayloadLen = 1 + 38 + 65 // typeID + message + signature
	type1PayloadLen = 1 + 65      // typeID + signature
)

func TestFromSignedPayload(t *testing.T) {
	message := bytes.Repeat([]byte{0xAB}, 38)
	signature := bytes.Repeat([]byte{0xCD}, 65)

	type0Payload := append(append([]byte{0x00}, message...), signature...)
	type1Payload := append([]byte{0x01}, signature...)

	tests := []struct {
		name    string
		msg     payloadMessage
		wantErr bool
		check   func(t *testing.T, s *submitSignaturesPayload)
	}{
		{
			name:    "nil payload",
			msg:     payloadMessage{payload: nil},
			wantErr: true,
		},
		{
			name:    "empty payload",
			msg:     payloadMessage{payload: []byte{}},
			wantErr: true,
		},
		{
			name:    "type 0 only type byte",
			msg:     payloadMessage{payload: []byte{0x00}},
			wantErr: true,
		},
		{
			name:    "type 0 one byte short",
			msg:     payloadMessage{payload: type0Payload[:type0PayloadLen-1]},
			wantErr: true,
		},
		{
			name:    "type 1 one byte short",
			msg:     payloadMessage{payload: type1Payload[:type1PayloadLen-1]},
			wantErr: true,
		},
		{
			name:    "invalid typeID 2",
			msg:     payloadMessage{payload: append([]byte{0x02}, signature...)},
			wantErr: true,
		},
		{
			name:    "invalid typeID 0xFF",
			msg:     payloadMessage{payload: append([]byte{0xFF}, signature...)},
			wantErr: true,
		},
		{
			name: "type 0 exact length",
			msg: payloadMessage{
				protocolID:    7,
				votingRoundID: 42,
				payload:       type0Payload,
			},
			check: func(t *testing.T, s *submitSignaturesPayload) {
				t.Helper()
				require.Equal(t, uint8(7), s.protocolID)
				require.Equal(t, uint32(42), s.votingRoundID)
				require.Equal(t, uint8(0), s.typeID)
				require.Equal(t, message, []byte(s.message))
				require.Equal(t, signature, s.signature)
				require.Equal(t, -1, s.voterIndex)
			},
		},
		{
			name: "type 1 exact length",
			msg: payloadMessage{
				protocolID:    9,
				votingRoundID: 1234,
				payload:       type1Payload,
			},
			check: func(t *testing.T, s *submitSignaturesPayload) {
				t.Helper()
				require.Equal(t, uint8(9), s.protocolID)
				require.Equal(t, uint32(1234), s.votingRoundID)
				require.Equal(t, uint8(1), s.typeID)
				require.Nil(t, s.message)
				require.Equal(t, signature, s.signature)
				require.Equal(t, -1, s.voterIndex)
			},
		},
		{
			name: "type 0 with trailing bytes",
			msg: payloadMessage{
				protocolID:    3,
				votingRoundID: 99,
				payload:       append(append([]byte{}, type0Payload...), 0xAA, 0xBB, 0xCC),
			},
			check: func(t *testing.T, s *submitSignaturesPayload) {
				t.Helper()
				require.Equal(t, message, []byte(s.message))
				require.Equal(t, signature, s.signature)
				require.Len(t, s.signature, 65)
				require.Len(t, s.message, 38)
			},
		},
		{
			name: "type 1 with trailing bytes",
			msg: payloadMessage{
				protocolID:    3,
				votingRoundID: 99,
				payload:       append(append([]byte{}, type1Payload...), 0xAA, 0xBB),
			},
			check: func(t *testing.T, s *submitSignaturesPayload) {
				t.Helper()
				require.Nil(t, s.message)
				require.Equal(t, signature, s.signature)
				require.Len(t, s.signature, 65)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var s submitSignaturesPayload
			err := s.FromSignedPayload(tc.msg)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			tc.check(t, &s)
		})
	}
}

// TestRoundTripWithEncodePayload generates a tx-style input with protocol.EncodePayload
// for both protocol types, parses it back through ExtractPayloads and FromSignedPayload,
// and confirms the round trip preserves protocolID, votingRoundID, typeID, and message.
func TestRoundTripWithEncodePayload(t *testing.T) {
	privateKey, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)

	const votingRound int64 = 1234

	type0Message := bytes.Repeat([]byte{0x11}, 38)
	type1Message := bytes.Repeat([]byte{0x22}, 16)

	cases := []struct {
		protocolID   uint8
		protocolType uint8
		data         []byte
	}{
		{protocolID: 1, protocolType: 0, data: type0Message},
		{protocolID: 5, protocolType: 1, data: type1Message},
	}

	buf := new(bytes.Buffer)
	buf.Write([]byte{0xde, 0xad, 0xbe, 0xef}) // function selector

	for _, c := range cases {
		resp := &protocol.SubProtocolResponse{
			Status: payload.Ok,
			Data:   c.data,
		}
		err := protocol.EncodePayload(buf, votingRound, resp, c.protocolID, c.protocolType, privateKey)
		require.NoError(t, err)
	}

	payloads, err := ExtractPayloads(buf.Bytes())
	require.NoError(t, err)
	require.Len(t, payloads, len(cases))

	for i, c := range cases {
		var s submitSignaturesPayload
		require.NoError(t, s.FromSignedPayload(payloads[i]))
		require.Equal(t, c.protocolID, s.protocolID)
		require.Equal(t, uint32(votingRound), s.votingRoundID)
		require.Equal(t, c.protocolType, s.typeID)
		require.Len(t, s.signature, 65)
		require.Equal(t, -1, s.voterIndex)
		if c.protocolType == 0 {
			require.Equal(t, c.data, []byte(s.message))
		} else {
			require.Nil(t, s.message)
		}
	}
}

func TestExtractPayloadsZeroLength(t *testing.T) {
	// 4-byte selector + 1-byte protocol + 4-byte votingRound + 2-byte length=0
	data := make([]byte, 4+1+4+2)
	binary.BigEndian.PutUint32(data[5:9], 1)
	binary.BigEndian.PutUint16(data[9:11], 0)

	payloads, err := ExtractPayloads(data)
	require.NoError(t, err)
	require.Len(t, payloads, 1)

	var s submitSignaturesPayload
	require.Error(t, s.FromSignedPayload(payloads[0]))
}

func TestExtractUint16Overflow(t *testing.T) {
	// 4-byte selector + 1-byte protocol + 4-byte votingRound + 2-byte length=0
	data := make([]byte, 4+1+4+2)
	binary.BigEndian.PutUint32(data[5:9], 1)
	binary.BigEndian.PutUint16(data[9:11], 0xffff)

	_, err := ExtractPayloads(data)
	require.Error(t, err)

	dataTrue := make([]byte, 4+1+4+2+0xffff)
	binary.BigEndian.PutUint32(dataTrue[5:9], 1)
	binary.BigEndian.PutUint16(dataTrue[9:11], 0xffff)

	payloads, err := ExtractPayloads(dataTrue)
	require.NoError(t, err)
	require.Len(t, payloads, 1)
}

// Regression: inputs shorter than the 4-byte function selector used to panic
// on the unguarded data[4:] slice.
func TestExtractPayloadsShorterThanSelector(t *testing.T) {
	for _, length := range []int{0, 1, 2, 3} {
		_, err := ExtractPayloads(make([]byte, length))
		require.Error(t, err, "input of %d bytes must be rejected, not panic", length)
	}
}

// --- AddSigner --------------------------------------------------------------
//
// Recovers the signer from a VRS-formatted signature, looks them up in the
// active voter set, stores their index + weight. Three real branches matter:
// happy path, signer-not-in-set, signature-fails-to-recover.

// signVRS produces a 65-byte VRS-formatted signature in the wire format the
// system-client receives via submitSignatures calldata (V at byte 0 with the
// canonical +27 offset, then R||S).
func signVRS(t *testing.T, hash []byte, priv *ecdsa.PrivateKey) []byte {
	t.Helper()
	rsv, err := crypto.Sign(hash, priv)
	require.NoError(t, err)
	require.Len(t, rsv, 65)
	vrs := make([]byte, 65)
	vrs[0] = rsv[64] + 27
	copy(vrs[1:33], rsv[:32])
	copy(vrs[33:65], rsv[32:64])
	return vrs
}

func newKeyAndAddress(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	priv, err := crypto.GenerateKey()
	require.NoError(t, err)
	return priv, crypto.PubkeyToAddress(priv.PublicKey)
}

func TestAddSigner_HappyPath(t *testing.T) {
	priv, addr := newKeyAndAddress(t)
	hash := crypto.Keccak256([]byte("any-message"))

	pld := submitSignaturesPayload{
		signature:  signVRS(t, hash, priv),
		voterIndex: -1,
	}

	other := common.HexToAddress("0x1111111111111111111111111111111111111111")
	set := voters.NewSet(
		[]common.Address{addr, other},
		[]uint16{777, 222},
		nil,
	)

	require.NoError(t, pld.AddSigner(hash, set))
	require.Equal(t, 0, pld.voterIndex)
	require.Equal(t, uint16(777), pld.weight)
	require.Equal(t, addr, pld.signer)
}

func TestAddSigner_LooksUpCorrectIndexForNonZeroPosition(t *testing.T) {
	// Pins that the index isn't hard-coded.
	_, addr1 := newKeyAndAddress(t)
	priv2, addr2 := newKeyAndAddress(t)

	hash := crypto.Keccak256([]byte("msg"))
	pld := submitSignaturesPayload{
		signature:  signVRS(t, hash, priv2),
		voterIndex: -1,
	}
	set := voters.NewSet(
		[]common.Address{addr1, addr2},
		[]uint16{100, 250},
		nil,
	)

	require.NoError(t, pld.AddSigner(hash, set))
	require.Equal(t, 1, pld.voterIndex)
	require.Equal(t, uint16(250), pld.weight)
}

func TestAddSigner_RejectsSignerOutsideVoterSet(t *testing.T) {
	priv, _ := newKeyAndAddress(t)
	hash := crypto.Keccak256([]byte("msg"))

	pld := submitSignaturesPayload{
		signature:  signVRS(t, hash, priv),
		voterIndex: -1,
	}
	stranger := common.HexToAddress("0x2222222222222222222222222222222222222222")
	set := voters.NewSet([]common.Address{stranger}, []uint16{500}, nil)

	err := pld.AddSigner(hash, set)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not a registered voter")
	require.Equal(t, -1, pld.voterIndex)
	require.Equal(t, uint16(0), pld.weight)
}

func TestAddSigner_WrongMessageHashRecoversWrongAddress(t *testing.T) {
	// Pins that AddSigner uses the SUPPLIED hash, not a stored one.
	// Recovery against the wrong hash yields a different public key.
	priv, addr := newKeyAndAddress(t)
	hash1 := crypto.Keccak256([]byte("first"))
	hash2 := crypto.Keccak256([]byte("second"))

	pld := submitSignaturesPayload{
		signature:  signVRS(t, hash1, priv),
		voterIndex: -1,
	}
	set := voters.NewSet([]common.Address{addr}, []uint16{100}, nil)

	err := pld.AddSigner(hash2, set)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not a registered voter")
}

func TestAddSigner_RejectsSignatureWithInvalidVByte(t *testing.T) {
	// V outside {27, 28} produces an out-of-range recovery byte after the
	// VRS→RSV transform; crypto.SigToPub rejects.
	priv, _ := newKeyAndAddress(t)
	hash := crypto.Keccak256([]byte("msg"))
	sig := signVRS(t, hash, priv)
	sig[0] = 99 // garbage V

	pld := submitSignaturesPayload{signature: sig, voterIndex: -1}
	addr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	set := voters.NewSet([]common.Address{addr}, []uint16{1}, nil)

	err := pld.AddSigner(hash, set)
	require.Error(t, err)
}
