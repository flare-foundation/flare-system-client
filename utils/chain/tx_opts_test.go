package chain_test

import (
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare-system-client/utils/chain"
	"github.com/stretchr/testify/require"
)

func TestCopyTxOptsIndependence(t *testing.T) {
	orig := &bind.TransactOpts{
		From:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Nonce:     big.NewInt(7),
		Value:     big.NewInt(0),
		GasPrice:  big.NewInt(100),
		GasFeeCap: big.NewInt(200),
		GasTipCap: big.NewInt(50),
		GasLimit:  21000,
	}

	cp := chain.CopyTxOpts(orig)
	require.Equal(t, orig.From, cp.From)
	require.Equal(t, 0, orig.GasPrice.Cmp(cp.GasPrice))

	// mutating the copy, both by assignment and in place, must not affect the original
	cp.GasLimit = 500000
	cp.GasPrice.SetInt64(999)
	cp.GasFeeCap = big.NewInt(888)
	cp.GasTipCap.SetInt64(777)

	require.Equal(t, uint64(21000), orig.GasLimit)
	require.Equal(t, int64(100), orig.GasPrice.Int64())
	require.Equal(t, int64(200), orig.GasFeeCap.Int64())
	require.Equal(t, int64(50), orig.GasTipCap.Int64())
}

func TestCopyTxOptsNilFields(t *testing.T) {
	orig := &bind.TransactOpts{
		From: common.HexToAddress("0x2222222222222222222222222222222222222222"),
	}

	cp := chain.CopyTxOpts(orig)
	require.Equal(t, orig.From, cp.From)
	require.Nil(t, cp.Nonce)
	require.Nil(t, cp.Value)
	require.Nil(t, cp.GasPrice)
	require.Nil(t, cp.GasFeeCap)
	require.Nil(t, cp.GasTipCap)
}

// TestCopyTxOptsConcurrentUse simulates the epoch client send paths: multiple
// goroutines copying the shared opts and applying per-tx gas settings to the
// copy. Run with -race; mutating the shared instance instead would fail it.
func TestCopyTxOptsConcurrentUse(t *testing.T) {
	shared := &bind.TransactOpts{
		From:      common.HexToAddress("0x3333333333333333333333333333333333333333"),
		GasPrice:  big.NewInt(100),
		GasFeeCap: big.NewInt(200),
		GasTipCap: big.NewInt(50),
	}

	var wg sync.WaitGroup
	for i := range 16 {
		wg.Go(func() {
			cp := chain.CopyTxOpts(shared)
			cp.GasLimit = uint64(i + 1)
			cp.GasPrice.Add(cp.GasPrice, big.NewInt(int64(i)))
			cp.GasFeeCap = big.NewInt(int64(i))
			cp.GasTipCap.SetInt64(int64(i))
		})
	}
	wg.Wait()

	require.Equal(t, int64(100), shared.GasPrice.Int64())
	require.Equal(t, int64(200), shared.GasFeeCap.Int64())
	require.Equal(t, int64(50), shared.GasTipCap.Int64())
}
