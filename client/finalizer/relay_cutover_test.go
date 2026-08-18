package finalizer

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/policy"
	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

var (
	oldRelayAddress = common.HexToAddress("0x00000000000000000000000000000000000000aa")
	newRelayAddress = common.HexToAddress("0x00000000000000000000000000000000000000bb")
)

const testBreakingEpoch = int64(5236)

func cutoverForTest() *shared.RelayCutover {
	return &shared.RelayCutover{
		ChainID:             testChainID,
		NewAddress:          newRelayAddress,
		BreakingRewardEpoch: testBreakingEpoch,
	}
}

// A finalization carries the policy bytes, so it only verifies on the Relay that
// stores that policy's hash: pre-cutover epochs stay on the old Relay even after
// the switch, or a late finalization would revert on a hash mismatch.
func TestAddressForRewardEpoch(t *testing.T) {
	cutover := cutoverForTest()
	r := &relayContractClient{address: oldRelayAddress, relayCutover: cutover}

	require.Equal(t, oldRelayAddress, r.addressForRewardEpoch(0))
	require.Equal(t, oldRelayAddress, r.addressForRewardEpoch(testBreakingEpoch-1))
	require.Equal(t, newRelayAddress, r.addressForRewardEpoch(testBreakingEpoch))
	require.Equal(t, newRelayAddress, r.addressForRewardEpoch(testBreakingEpoch+1))

	require.Equal(t, []common.Address{oldRelayAddress, newRelayAddress}, r.addresses())

	// with no switch scheduled nothing changes and only one Relay is read
	plain := &relayContractClient{address: oldRelayAddress, relayCutover: shared.NewRelayCutover(testChainID)}
	require.Equal(t, oldRelayAddress, plain.addressForRewardEpoch(1<<40))
	require.Equal(t, []common.Address{oldRelayAddress}, plain.addresses())
}

// The chosen Relay must be the one the tx is actually sent to, and an unset one
// (a scheduled cutover missing its address) must not fall back to the old Relay.
func TestSubmitPayloadsSendsToGivenAddress(t *testing.T) {
	cc := &scriptedRelayClient{nonces: []uint64{10}, results: []chain.SendResult{{}}}
	r := testRelayClient(t, cc)

	r.SubmitPayloads(context.Background(), newRelayAddress, make([]byte, 40), false, 1, 100)
	require.Equal(t, []common.Address{newRelayAddress}, cc.sentTo)

	r.SubmitPayloads(context.Background(), common.Address{}, make([]byte, 40), false, 1, 100)
	require.Len(t, cc.sentTo, 1, "a zero target must not be sent anywhere")
}

// logsDB returns per-address logs; everything else is unused by these tests.
type logsDB struct {
	finalizerDB
	logs map[common.Address][]database.Log
}

func (db logsDB) FetchLogsByAddressAndTopic0(
	_ context.Context, address common.Address, _ common.Hash, _, _ int64,
) ([]database.Log, error) {
	return db.logs[address], nil
}

func spiLogAt(block, index uint64) database.Log {
	return database.Log{BlockNumber: block, LogIndex: index, Timestamp: block}
}

// Policies before the cutover are only emitted by the old Relay and after it only
// by the new one, so both are read — and the merge must come back in chain order:
// the listener advances its event range from the last log it sees.
func TestFetchSigningPoliciesMergesBothRelaysInChainOrder(t *testing.T) {
	db := logsDB{logs: map[common.Address][]database.Log{
		oldRelayAddress: {spiLogAt(10, 0), spiLogAt(30, 1)},
		newRelayAddress: {spiLogAt(20, 0), spiLogAt(30, 0)},
	}}
	r := &relayContractClient{address: oldRelayAddress, relayCutover: cutoverForTest()}

	logs, err := r.fetchLogs(context.Background(), db, r.topic0SPI, 0, 100)
	require.NoError(t, err)
	require.Equal(t,
		[]database.Log{spiLogAt(10, 0), spiLogAt(20, 0), spiLogAt(30, 0), spiLogAt(30, 1)},
		logs)
}

// A round relayed only on the old Relay is not relayed for readers of the new one,
// so the already-relayed check must be per target rather than a merged set.
func TestRelayedSetIsPerAddress(t *testing.T) {
	key := relayedKey{protocolID: 100, votingRoundID: 42}
	set := relayedSet{oldRelayAddress: {key: true}, newRelayAddress: {}}

	require.True(t, set.has(oldRelayAddress, key))
	require.False(t, set.has(newRelayAddress, key))
	require.False(t, set.has(common.Address{}, key))
}

func signDigest(t *testing.T, digest []byte, key *ecdsa.PrivateKey) []byte {
	t.Helper()
	signature, err := crypto.Sign(digest, key)
	require.NoError(t, err)
	vrs, err := utils.TransformSignatureRSVtoVRS(signature)
	require.NoError(t, err)
	return vrs
}

// The finalizer recovers signers under the digest of the Relay holding the
// collection's policy. Getting this gate wrong rejects every peer signature from
// the breaking epoch on, so both sides of the boundary are pinned.
func TestFinalizerRecoversSignersUnderThePolicyEpochDigest(t *testing.T) {
	cutover := cutoverForTest()

	key, err := crypto.HexToECDSA(testPrivateKeyHex)
	require.NoError(t, err)
	signer := crypto.PubkeyToAddress(key.PublicKey)

	message := make(shared.Message, 38)
	message[0] = 1

	cases := []struct {
		name        string
		rewardEpoch int64
		chainBound  bool
	}{
		{"policy before the cutover", testBreakingEpoch - 1, false},
		{"policy of the breaking epoch", testBreakingEpoch, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sp := &policy.SigningPolicy{
				RewardEpochID: c.rewardEpoch,
				Voters:        voters.NewSet([]common.Address{signer}, []uint16{2}, nil),
			}

			matching := &submitSignaturesPayload{
				typeID:        0,
				sender:        signer,
				votingRoundID: 7,
				protocolID:    1,
				message:       message,
				signature:     signDigest(t, shared.MessageDigest(message, cutover.ChainID, c.chainBound), key),
			}
			s := newFinalizationStorage(cutover)
			ready, err := s.addPayload(matching, sp, 1)
			require.NoError(t, err)
			require.True(t, ready.thresholdReached)

			// the other digest form recovers a stranger, which is not in the policy
			wrong := &submitSignaturesPayload{
				typeID:        0,
				sender:        signer,
				votingRoundID: 7,
				protocolID:    1,
				message:       message,
				signature:     signDigest(t, shared.MessageDigest(message, cutover.ChainID, !c.chainBound), key),
			}
			s = newFinalizationStorage(cutover)
			_, err = s.addPayload(wrong, sp, 1)
			require.ErrorIs(t, err, errBadPayload)
		})
	}
}
