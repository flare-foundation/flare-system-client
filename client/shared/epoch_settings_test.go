package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/stretchr/testify/require"
)

// relayStateData answers eth_call with a stateData() tuple carrying the given values.
func relayStateData(t *testing.T, randomNumberProtocolID uint8, thresholdIncreaseBIPS uint16) *relay.Relay {
	t.Helper()

	relayABI, err := relay.RelayMetaData.GetAbi()
	require.NoError(t, err)
	encoded, err := relayABI.Methods["stateData"].Outputs.Pack(
		randomNumberProtocolID,
		uint32(1e9),  // firstVotingRoundStartTs
		uint8(90),    // votingEpochDurationSeconds
		uint32(0),    // firstRewardEpochStartVotingRoundId
		uint16(3360), // rewardEpochDurationInVotingEpochs
		thresholdIncreaseBIPS,
		uint32(0),  // randomVotingRoundId
		false,      // isSecureRandom
		uint32(0),  // lastInitializedRewardEpoch
		false,      // noSigningPolicyRelay
		uint32(10), // messageFinalizationWindowInRewardEpochs
	)
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(body, &req))
		require.Equal(t, "eth_call", req.Method)

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, hexutil.Encode(encoded))
	}))
	t.Cleanup(srv.Close)

	eth, err := ethclient.Dial(srv.URL)
	require.NoError(t, err)
	t.Cleanup(eth.Close)

	r, err := relay.NewRelay(common.HexToAddress("0x00000000000000000000000000000000000000aa"), eth)
	require.NoError(t, err)
	return r
}

// Both scalars come off the wire, so a hardcoded value cannot pass for them: the
// trailer gate and the prolonged-epoch threshold are the Relay's to set.
func TestEpochsFromChainReadsTheRelaysScalars(t *testing.T) {
	timing, rewardEpoch, protocolID, bips, err := EpochsFromChain(relayStateData(t, 100, 12000))
	require.NoError(t, err)
	require.Equal(t, uint8(100), protocolID)
	require.Equal(t, uint16(12000), bips)
	require.NotNil(t, timing)
	require.NotNil(t, rewardEpoch)

	_, _, protocolID, bips, err = EpochsFromChain(relayStateData(t, 7, 15000))
	require.NoError(t, err)
	require.Equal(t, uint8(7), protocolID)
	require.Equal(t, uint16(15000), bips)
}
