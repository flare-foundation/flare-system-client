package chain_test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"
	"time"

	clientConfig "github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

type fakeClient struct {
	receipts   map[common.Hash]*types.Receipt
	reverts    map[common.Hash]string
	receiptErr map[common.Hash]bool // hashes whose Receipt lookup fails (RPC error)
	revertErr  map[common.Hash]bool // hashes whose RevertReason lookup fails
}

func (f fakeClient) SendRawTx(context.Context, *ecdsa.PrivateKey, uint64, common.Address, []byte, *clientConfig.Gas, time.Duration, bool) chain.SendResult {
	return chain.SendResult{}
}
func (f fakeClient) Nonce(context.Context, *ecdsa.PrivateKey, time.Duration) (uint64, error) {
	return 0, nil
}
func (f fakeClient) Receipt(_ context.Context, h common.Hash, _ time.Duration) (*types.Receipt, error) {
	if f.receiptErr[h] {
		return nil, errors.New("rpc down")
	}
	return f.receipts[h], nil
}
func (f fakeClient) RevertReason(_ context.Context, _ common.Address, h common.Hash, _ time.Duration) (string, error) {
	if f.revertErr[h] {
		return "", errors.New("rpc down")
	}
	return f.reverts[h], nil
}

var (
	hA = common.HexToHash("0x0a")
	hB = common.HexToHash("0x0b")
)

func TestAnyAccepted(t *testing.T) {
	ok := &types.Receipt{Status: types.ReceiptStatusSuccessful}
	reverted := &types.Receipt{Status: types.ReceiptStatusFailed}

	tests := []struct {
		name     string
		client   fakeClient
		hashes   []common.Hash
		allowed  []string
		wantHash common.Hash
		want     chain.Acceptance
	}{
		{
			name:     "mined successful",
			client:   fakeClient{receipts: map[common.Hash]*types.Receipt{hA: ok}},
			hashes:   []common.Hash{hA},
			wantHash: hA, want: chain.Accepted,
		},
		{
			name:   "not mined -> conclusively not accepted",
			client: fakeClient{receipts: map[common.Hash]*types.Receipt{}},
			hashes: []common.Hash{hA},
			want:   chain.NotAccepted,
		},
		{
			name:     "reverted with allowed reason",
			client:   fakeClient{receipts: map[common.Hash]*types.Receipt{hA: reverted}, reverts: map[common.Hash]string{hA: "Already relayed"}},
			hashes:   []common.Hash{hA},
			allowed:  []string{"Already relayed"},
			wantHash: hA, want: chain.Accepted,
		},
		{
			name:    "reverted with other reason",
			client:  fakeClient{receipts: map[common.Hash]*types.Receipt{hA: reverted}, reverts: map[common.Hash]string{hA: "boom"}},
			hashes:  []common.Hash{hA},
			allowed: []string{"Already relayed"},
			want:    chain.NotAccepted,
		},
		{
			name:   "reverted, no allowed list",
			client: fakeClient{receipts: map[common.Hash]*types.Receipt{hA: reverted}},
			hashes: []common.Hash{hA},
			want:   chain.NotAccepted,
		},
		{
			name:     "first pending, second mined",
			client:   fakeClient{receipts: map[common.Hash]*types.Receipt{hB: ok}},
			hashes:   []common.Hash{hA, hB},
			wantHash: hB, want: chain.Accepted,
		},
		{
			name:   "receipt RPC error -> undetermined",
			client: fakeClient{receiptErr: map[common.Hash]bool{hA: true}},
			hashes: []common.Hash{hA},
			want:   chain.Undetermined,
		},
		{
			name:    "revert-reason RPC error -> undetermined",
			client:  fakeClient{receipts: map[common.Hash]*types.Receipt{hA: reverted}, revertErr: map[common.Hash]bool{hA: true}},
			hashes:  []common.Hash{hA},
			allowed: []string{"Already relayed"},
			want:    chain.Undetermined,
		},
		{
			name:     "one hash errors but another is accepted -> accepted",
			client:   fakeClient{receiptErr: map[common.Hash]bool{hA: true}, receipts: map[common.Hash]*types.Receipt{hB: ok}},
			hashes:   []common.Hash{hA, hB},
			wantHash: hB, want: chain.Accepted,
		},
		{
			name:   "no hashes -> conclusively not accepted",
			client: fakeClient{},
			hashes: nil,
			want:   chain.NotAccepted,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h, acc := chain.AnyAccepted(context.Background(), tc.client, common.Address{}, tc.hashes, tc.allowed, time.Second)
			require.Equal(t, tc.want, acc)
			if tc.want == chain.Accepted {
				require.Equal(t, tc.wantHash, h)
			}
		})
	}
}

func TestErrorClassifiers(t *testing.T) {
	require.True(t, chain.IsNonceTooLow(errors.New("rpc: nonce too low: ...")))
	require.False(t, chain.IsNonceTooLow(errors.New("replacement transaction underpriced")))
	require.False(t, chain.IsNonceTooLow(nil))

	require.True(t, chain.IsTimeout(context.DeadlineExceeded))
	require.False(t, chain.IsTimeout(errors.New("nonce too low")))
	require.False(t, chain.IsTimeout(nil))

	require.True(t, chain.MatchesError(errors.New("tx failed: Already relayed"), []string{"Already relayed"}))
	require.False(t, chain.MatchesError(errors.New("boom"), []string{"Already relayed"}))
	require.False(t, chain.MatchesError(nil, []string{"Already relayed"}))
}
