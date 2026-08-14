package chain

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// encodeRevertReason builds revert data as returned by a contract: Error(string) selector + ABI-encoded string.
func encodeRevertReason(t *testing.T, reason string) []byte {
	t.Helper()

	encoded, err := abi.Arguments{{Type: abiString}}.Pack(reason)
	require.NoError(t, err)

	return append(append([]byte{}, errorSig...), encoded...)
}

func TestUnpackError(t *testing.T) {
	tests := []struct {
		name      string
		result    []byte
		want      string
		expectErr bool
	}{
		{
			name:      "valid revert reason",
			result:    encodeRevertReason(t, "nonce too low"),
			want:      "nonce too low",
			expectErr: false,
		},
		{
			name:      "empty result",
			result:    []byte{},
			want:      "<tx result not Error(string)>",
			expectErr: true,
		},
		{
			name:      "result shorter than selector",
			result:    []byte{0x08, 0xc3},
			want:      "<tx result not Error(string)>",
			expectErr: true,
		},
		{
			name:      "unknown selector",
			result:    []byte{0xde, 0xad, 0xbe, 0xef, 0x01, 0x02},
			want:      "<tx result not Error(string)>",
			expectErr: true,
		},
		{
			name:      "valid selector with malformed payload",
			result:    append(append([]byte{}, errorSig...), 0x01, 0x02, 0x03),
			want:      "<invalid tx result>",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := unpackError(tc.result)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.want, got)
		})
	}
}

// pollStub scripts TransactionReceipt by call count, never wall time.
type pollStub struct {
	stubCaller
	calls  int
	script func(call int) (*types.Receipt, error)
}

func (s *pollStub) TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error) {
	s.calls++
	return s.script(s.calls)
}

// newPollVerifier sets all three intervals explicitly — a zero interval busy-spins.
func newPollVerifier(eth txBackend, fast, slow, backoffAfter time.Duration) TxVerifier {
	return TxVerifier{eth: eth, pollInterval: fast, pollIntervalSlow: slow, pollBackoffAfter: backoffAfter}
}

func TestWaitUntilMined(t *testing.T) {
	ctx := context.Background()
	tx := types.NewTx(&types.LegacyTx{})
	success := &types.Receipt{Status: types.ReceiptStatusSuccessful}

	t.Run("immediate receipt", func(t *testing.T) {
		stub := &pollStub{script: func(int) (*types.Receipt, error) { return success, nil }}
		v := newPollVerifier(stub, time.Hour, time.Hour, time.Hour)

		err := v.WaitUntilMined(ctx, common.Address{}, tx, time.Second)
		require.NoError(t, err)
		require.Equal(t, 1, stub.calls) // pins check-before-first-wait
	})

	t.Run("found after fast polling", func(t *testing.T) {
		stub := &pollStub{script: func(call int) (*types.Receipt, error) {
			if call < 5 {
				return nil, ethereum.NotFound
			}
			return success, nil
		}}
		// slow branch would wait an hour — success within timeout pins the fast branch
		v := newPollVerifier(stub, time.Millisecond, time.Hour, time.Hour)

		err := v.WaitUntilMined(ctx, common.Address{}, tx, time.Second)
		require.NoError(t, err)
		require.Equal(t, 5, stub.calls)
	})

	t.Run("backoff selects slow interval", func(t *testing.T) {
		stub := &pollStub{script: func(call int) (*types.Receipt, error) {
			if call == 1 {
				return nil, ethereum.NotFound
			}
			return success, nil
		}}
		// fast branch would wait an hour — success within timeout pins the slow branch
		v := newPollVerifier(stub, time.Hour, time.Millisecond, 0)

		err := v.WaitUntilMined(ctx, common.Address{}, tx, time.Second)
		require.NoError(t, err)
		require.Equal(t, 2, stub.calls)
	})

	t.Run("transient RPC error tolerated", func(t *testing.T) {
		stub := &pollStub{script: func(call int) (*types.Receipt, error) {
			switch call {
			case 1:
				return nil, errors.New("502")
			case 2:
				return nil, ethereum.NotFound
			default:
				return success, nil
			}
		}}
		v := newPollVerifier(stub, time.Millisecond, time.Millisecond, time.Hour)

		err := v.WaitUntilMined(ctx, common.Address{}, tx, time.Second)
		require.NoError(t, err)
		require.Equal(t, 3, stub.calls)
	})

	t.Run("timeout while pending", func(t *testing.T) {
		stub := &pollStub{script: func(int) (*types.Receipt, error) { return nil, ethereum.NotFound }}
		v := newPollVerifier(stub, time.Millisecond, time.Millisecond, time.Hour)

		err := v.WaitUntilMined(ctx, common.Address{}, tx, 20*time.Millisecond)
		require.ErrorIs(t, err, context.DeadlineExceeded)
		require.NotErrorIs(t, err, errReverted)
	})

	t.Run("ctx already cancelled", func(t *testing.T) {
		stub := &pollStub{script: func(int) (*types.Receipt, error) { return nil, ethereum.NotFound }}
		v := newPollVerifier(stub, time.Hour, time.Hour, time.Hour)

		cancelled, cancel := context.WithCancel(ctx)
		cancel()

		err := v.WaitUntilMined(cancelled, common.Address{}, tx, time.Minute)
		require.ErrorIs(t, err, context.Canceled)
		require.Equal(t, 1, stub.calls)
	})

	t.Run("mined but reverted", func(t *testing.T) {
		stub := &pollStub{
			stubCaller: stubCaller{err: stubDataError{
				msg:  "execution reverted: some reason",
				data: hexutil.Encode(encodeRevertReason(t, "some reason")),
			}},
			script: func(int) (*types.Receipt, error) {
				return &types.Receipt{Status: types.ReceiptStatusFailed, BlockNumber: big.NewInt(1)}, nil
			},
		}
		v := newPollVerifier(stub, time.Hour, time.Hour, time.Hour)

		err := v.WaitUntilMined(ctx, common.Address{}, tx, time.Second)
		require.ErrorIs(t, err, errReverted)
		require.ErrorContains(t, err, "some reason")
	})
}
