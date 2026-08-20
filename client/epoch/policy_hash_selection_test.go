package epoch

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/stretchr/testify/require"
)

const (
	toSigningPolicyHashSelector = "0c85bf07"

	breakingEpoch = 5236
)

// policyHashNode answers manager.relay() with relayAddr and relay.toSigningPolicyHash(epoch)
// with stored, recording the epoch argument of every hash read. No production seam needed.
type policyHashNode struct {
	eth *ethclient.Client

	mu     sync.Mutex
	asked  []int64
	relays int
}

func newPolicyHashNode(t *testing.T, relayAddr common.Address, stored []byte) *policyHashNode {
	t.Helper()
	n := &policyHashNode{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
			Params []any           `json:"params"`
		}
		require.NoError(t, json.Unmarshal(body, &req))
		require.Equal(t, "eth_call", req.Method, "unexpected RPC method")

		// this go-ethereum version sends the call payload as "input" (older ones use "data")
		call, _ := req.Params[0].(map[string]any)
		data, _ := call["input"].(string)
		if data == "" {
			data, _ = call["data"].(string)
		}
		data = strings.TrimPrefix(data, "0x")

		var result string
		n.mu.Lock()
		if strings.HasPrefix(data, toSigningPolicyHashSelector) {
			arg, err := hex.DecodeString(data[len(toSigningPolicyHashSelector):])
			require.NoError(t, err)
			require.Len(t, arg, 32, "toSigningPolicyHash takes one word")
			n.asked = append(n.asked, new(big.Int).SetBytes(arg).Int64())
			result = "0x" + hex.EncodeToString(stored)
		} else {
			n.relays++
			result = "0x" + hex.EncodeToString(common.LeftPadBytes(relayAddr.Bytes(), 32))
		}
		n.mu.Unlock()
		_, _ = fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, result)
	}))
	t.Cleanup(srv.Close)

	eth, err := ethclient.Dial(srv.URL)
	require.NoError(t, err)
	t.Cleanup(eth.Close)
	n.eth = eth
	return n
}

func (n *policyHashNode) epochsAsked() []int64 {
	n.mu.Lock()
	defer n.mu.Unlock()
	return append([]int64(nil), n.asked...)
}

var (
	oldRelayAddress = common.HexToAddress("0x00000000000000000000000000000000000000aa")
	newRelayAddress = common.HexToAddress("0x00000000000000000000000000000000000000bb")
)

// pointsAt is the Relay address FlareSystemsManager reports — the one it verifies against.
func clientForStoredHash(t *testing.T, pointsAt common.Address, stored []byte, cutover *shared.RelayCutover) (*systemsManagerContractClientImpl, *policyHashNode) {
	t.Helper()
	managerAddr := common.HexToAddress("0x0000000000000000000000000000000000000043")
	node := newPolicyHashNode(t, pointsAt, stored)
	manager, err := system.NewFlareSystemsManager(managerAddr, node.eth)
	require.NoError(t, err)
	return &systemsManagerContractClientImpl{
		address:             managerAddr,
		flareSystemsManager: manager,
		chainID:             cutover.ChainID,
		ethClient:           node.eth,
		relayCutover:        cutover,
	}, node
}

// literal, not NewRelayCutover: these fixtures must not drift when the table is filled
func unscheduledCutover() *shared.RelayCutover {
	return &shared.RelayCutover{ChainID: vectorChainID}
}

func scheduledCutover() *shared.RelayCutover {
	return &shared.RelayCutover{
		ChainID:             vectorChainID,
		NewAddress:          newRelayAddress,
		BreakingRewardEpoch: breakingEpoch,
	}
}

func testPolicy(t *testing.T) []byte {
	t.Helper()
	policy, err := hex.DecodeString(vectorPolicyHex)
	require.NoError(t, err)
	return policy
}

func TestSigningPolicyHashOutcomes(t *testing.T) {
	policy := testPolicy(t)
	epoch := big.NewInt(5236)

	t.Run("chain-bound match", func(t *testing.T) {
		c, node := clientForStoredHash(t, newRelayAddress, ChainBoundSigningPolicyHash(policy, vectorChainID), scheduledCutover())
		got, err := c.signingPolicyHash(epoch, policy)
		require.NoError(t, err)
		require.Equal(t, vectorChainBoundHashHex, hex.EncodeToString(got))
		// the hash must be read for the epoch being signed, not for whatever the binding defaults to
		require.Equal(t, []int64{5236}, node.epochsAsked())
	})

	t.Run("legacy match", func(t *testing.T) {
		c, node := clientForStoredHash(t, oldRelayAddress, SigningPolicyHash(policy), unscheduledCutover())
		got, err := c.signingPolicyHash(epoch, policy)
		require.NoError(t, err)
		require.Equal(t, vectorLegacyHashHex, hex.EncodeToString(got))
		require.Equal(t, []int64{5236}, node.epochsAsked())
	})

	t.Run("no match", func(t *testing.T) {
		c, _ := clientForStoredHash(t, oldRelayAddress, common.LeftPadBytes([]byte{0xde, 0xad}, 32), unscheduledCutover())
		got, err := c.signingPolicyHash(epoch, policy)
		require.Nil(t, got)
		require.ErrorContains(t, err, "no supported hash of the signing policy of epoch 5236 matches relay")
	})

	// a policy hashed for the wrong source chain must never be signed
	t.Run("chain-bound for another chain does not match", func(t *testing.T) {
		c, _ := clientForStoredHash(t, newRelayAddress, ChainBoundSigningPolicyHash(policy, 14), scheduledCutover())
		_, err := c.signingPolicyHash(epoch, policy)
		require.Error(t, err)
	})
}

// The exact truth table of what each Relay stores: the old one folds every epoch, the
// new one folds below the breaking epoch (it delegates) and binds the chain from it on.
// Epoch B itself is signed before governance can repoint — the deploy cannot precede B's
// policy, and the repoint waits out a timelock on top — so B under the old pointer is the
// sanctioned case, not an anomaly.
func TestSigningPolicyHashFollowsTheBreakingEpochAndThePointer(t *testing.T) {
	policy := testPolicy(t)
	chainBound := ChainBoundSigningPolicyHash(policy, vectorChainID)
	legacy := SigningPolicyHash(policy)

	tests := []struct {
		name     string
		pointsAt common.Address
		epoch    int64
		cutover  *shared.RelayCutover
		want     []byte
	}{
		{"old relay, below the breaking epoch", oldRelayAddress, 5235, scheduledCutover(), legacy},
		{"old relay, at the breaking epoch", oldRelayAddress, 5236, scheduledCutover(), legacy},
		{"old relay, above the breaking epoch", oldRelayAddress, 5237, scheduledCutover(), legacy},
		{"new relay, below the breaking epoch", newRelayAddress, 5235, scheduledCutover(), legacy},
		{"new relay, at the breaking epoch", newRelayAddress, 5236, scheduledCutover(), chainBound},
		{"new relay, above the breaking epoch", newRelayAddress, 5237, scheduledCutover(), chainBound},
		{"unscheduled", oldRelayAddress, 5236, unscheduledCutover(), legacy},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			warnings := captureWarnings(t)
			c, _ := clientForStoredHash(t, tt.pointsAt, tt.want, tt.cutover)

			got, err := c.signingPolicyHash(big.NewInt(tt.epoch), policy)
			require.NoError(t, err)
			require.Equal(t, hex.EncodeToString(tt.want), hex.EncodeToString(got))
			require.NotContains(t, warnings(), "hashed", "the expected scheme must need no fallback")
		})
	}
}

// The fallback keeps the signature when the table disagrees with the chain, and says so.
func TestSigningPolicyHashFallsBackLoudly(t *testing.T) {
	policy := testPolicy(t)
	chainBound := ChainBoundSigningPolicyHash(policy, vectorChainID)
	legacy := SigningPolicyHash(policy)

	tests := []struct {
		name     string
		pointsAt common.Address
		epoch    int64
		stored   []byte
		warn     string
	}{
		{"new relay binds an epoch the table calls old", newRelayAddress, 5235, chainBound, "hashed chain-bound"},
		{"new relay folds an epoch the table calls new", newRelayAddress, 5236, legacy, "hashed legacy"},
		{"unrecognised pointer already binds", oldRelayAddress, 5236, chainBound, "hashed chain-bound"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			warnings := captureWarnings(t)
			c, _ := clientForStoredHash(t, tt.pointsAt, tt.stored, scheduledCutover())

			got, err := c.signingPolicyHash(big.NewInt(tt.epoch), policy)
			require.NoError(t, err)
			require.Equal(t, hex.EncodeToString(tt.stored), hex.EncodeToString(got))
			require.Contains(t, warnings(), tt.warn)
			require.Contains(t, warnings(), fmt.Sprintf("epoch %d", tt.epoch))
		})
	}
}

// The pointer is re-read per call, so a governance repoint mid-epoch is picked up.
func TestSigningPolicyHashReadsTheManagersRelayEveryCall(t *testing.T) {
	policy := testPolicy(t)
	c, node := clientForStoredHash(t, newRelayAddress, ChainBoundSigningPolicyHash(policy, vectorChainID), scheduledCutover())

	for range 3 {
		_, err := c.signingPolicyHash(big.NewInt(5236), policy)
		require.NoError(t, err)
	}

	node.mu.Lock()
	defer node.mu.Unlock()
	require.Equal(t, 3, node.relays, "manager.relay() must not be cached across attempts")
}

// captureWarnings redirects the process logger to a file for the duration of the test.
// Tests using it must not be parallel: the logger is global.
func captureWarnings(t *testing.T) func() string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "warn.log")
	logger.Set(logger.Config{Level: "WARN", File: path, MaxFileSize: 1})
	t.Cleanup(func() { logger.Set(logger.DefaultConfig()) })

	return func() string {
		b, err := os.ReadFile(path)
		if err != nil {
			return ""
		}
		return string(b)
	}
}
