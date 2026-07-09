package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// mayBeInMempool must classify ambiguous broadcast errors as possibly-in-mempool
// (so their hash is tracked) and only known synchronous rejections as not.
func TestMayBeInMempool(t *testing.T) {
	// Definitive rejections: the node never admitted the tx.
	require.False(t, mayBeInMempool(errors.New("nonce too low")))
	require.False(t, mayBeInMempool(errors.New("replacement transaction underpriced")))
	require.False(t, mayBeInMempool(errors.New("insufficient funds for gas * price + value")))
	require.False(t, mayBeInMempool(errors.New("intrinsic gas too low")))

	// Ambiguous failures: the tx may already be in the mempool.
	require.True(t, mayBeInMempool(context.DeadlineExceeded))
	require.True(t, mayBeInMempool(errors.New("read tcp 1.2.3.4:443: connection reset by peer")))
	require.True(t, mayBeInMempool(errors.New("unexpected EOF")))
	require.True(t, mayBeInMempool(errors.New("http2: server sent GOAWAY and closed the connection")))
	require.True(t, mayBeInMempool(errors.New("502 Bad Gateway")))
}

func TestIsAlreadyKnown(t *testing.T) {
	require.True(t, isAlreadyKnown(errors.New("already known")))
	require.False(t, isAlreadyKnown(errors.New("nonce too low")))
	require.False(t, isAlreadyKnown(nil))
}

// stubDataError implements rpc.DataError, mimicking the JSON-RPC error a
// geth/coreth node returns for a reverting eth_call (revert bytes hex-encoded
// in the Data field).
type stubDataError struct {
	msg  string
	data any
}

func (e stubDataError) Error() string  { return e.msg }
func (e stubDataError) ErrorCode() int { return 3 }
func (e stubDataError) ErrorData() any { return e.data }

// stubCaller is an ethereum.ContractCaller whose CallContract returns a scripted
// result/error, so errorReason can be exercised without a real node.
type stubCaller struct {
	res []byte
	err error
}

func (c stubCaller) CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error) {
	return c.res, c.err
}

// TestErrorReasonRecoversRevert: on geth/coreth a reverting eth_call comes back as
// a JSON-RPC error carrying the revert data in ErrorData(), not as return bytes.
// errorReason must recover the reason from there, and must distinguish a
// deterministic (undecodable) revert from a
// transient RPC failure.
func TestErrorReasonRecoversRevert(t *testing.T) {
	ctx := context.Background()
	tx := types.NewTx(&types.LegacyTx{})
	from := common.Address{}
	reason := "Already relayed"
	revert := encodeRevertReason(t, reason)

	t.Run("json-rpc error with hex revert data (geth path)", func(t *testing.T) {
		got, err := errorReason(ctx, stubCaller{
			err: stubDataError{msg: "execution reverted: " + reason, data: hexutil.Encode(revert)},
		}, from, tx, nil)
		require.NoError(t, err)
		require.Equal(t, reason, got)
	})

	t.Run("revert data returned as bytes (legacy path)", func(t *testing.T) {
		got, err := errorReason(ctx, stubCaller{res: revert}, from, tx, nil)
		require.NoError(t, err)
		require.Equal(t, reason, got)
	})

	t.Run("undecodable custom-error revert is deterministic", func(t *testing.T) {
		custom := []byte{0xde, 0xad, 0xbe, 0xef}
		_, err := errorReason(ctx, stubCaller{
			err: stubDataError{msg: "execution reverted", data: hexutil.Encode(custom)},
		}, from, tx, nil)
		require.Error(t, err)
		require.ErrorIs(t, err, errRevertUndecodable)
	})

	t.Run("reasonless revert (\"0x\" data) is deterministic", func(t *testing.T) {
		// geth encodes an empty revert (require(false), panics) as "0x", never nil.
		_, err := errorReason(ctx, stubCaller{
			err: stubDataError{msg: "execution reverted", data: "0x"},
		}, from, tx, nil)
		require.Error(t, err)
		require.ErrorIs(t, err, errRevertUndecodable)
	})

	t.Run("plain transport error carries no revert data", func(t *testing.T) {
		_, err := errorReason(ctx, stubCaller{err: errors.New("connection reset by peer")}, from, tx, nil)
		require.Error(t, err)
		require.NotErrorIs(t, err, errRevertUndecodable)
	})
}

func TestRevertDataFromError(t *testing.T) {
	revert := encodeRevertReason(t, "boom")

	raw, ok := revertDataFromError(stubDataError{data: hexutil.Encode(revert)})
	require.True(t, ok)
	require.Equal(t, revert, raw)

	raw, ok = revertDataFromError(stubDataError{data: hexutil.Bytes(revert)})
	require.True(t, ok)
	require.Equal(t, revert, raw)

	raw, ok = revertDataFromError(stubDataError{data: revert}) // plain []byte arm
	require.True(t, ok)
	require.Equal(t, revert, raw)

	_, ok = revertDataFromError(errors.New("connection reset"))
	require.False(t, ok)

	// DataError with non-hex string data is not usable revert data.
	_, ok = revertDataFromError(stubDataError{data: "not hex"})
	require.False(t, ok)
}

// stubReconcileClient is a minimal chain.Client for exercising AnyAccepted's
// mined-but-reverted branch (only Receipt/RevertReason are used).
type stubReconcileClient struct {
	receipts  map[common.Hash]*types.Receipt
	revertErr map[common.Hash]error
}

func (stubReconcileClient) SendRawTx(context.Context, *ecdsa.PrivateKey, uint64, common.Address, []byte, *config.Gas, time.Duration, bool) SendResult {
	return SendResult{}
}
func (stubReconcileClient) Nonce(context.Context, *ecdsa.PrivateKey, time.Duration) (uint64, error) {
	return 0, nil
}
func (c stubReconcileClient) Receipt(_ context.Context, h common.Hash, _ time.Duration) (*types.Receipt, error) {
	return c.receipts[h], nil
}
func (c stubReconcileClient) RevertReason(_ context.Context, _ common.Address, h common.Hash, _ time.Duration) (string, error) {
	return "", c.revertErr[h]
}

// TestAnyAcceptedRevertClassification checks that a deterministically-reverted
// tx whose reason cannot be decoded is terminal (Reverted), while a transient
// revert-reason RPC failure stays Undetermined.
func TestAnyAcceptedRevertClassification(t *testing.T) {
	reverted := &types.Receipt{Status: types.ReceiptStatusFailed}
	h := common.HexToHash("0x01")

	deterministic := stubReconcileClient{
		receipts:  map[common.Hash]*types.Receipt{h: reverted},
		revertErr: map[common.Hash]error{h: fmt.Errorf("%w: unknown error signature", errRevertUndecodable)},
	}
	gotHash, acc := AnyAccepted(context.Background(), deterministic, common.Address{}, []common.Hash{h}, []string{"Already relayed"}, time.Second)
	require.Equal(t, Reverted, acc)
	require.Equal(t, h, gotHash)

	transient := stubReconcileClient{
		receipts:  map[common.Hash]*types.Receipt{h: reverted},
		revertErr: map[common.Hash]error{h: errors.New("rpc down")},
	}
	_, acc = AnyAccepted(context.Background(), transient, common.Address{}, []common.Hash{h}, []string{"Already relayed"}, time.Second)
	require.Equal(t, Undetermined, acc)
}
