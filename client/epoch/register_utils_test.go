package epoch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	clconfig "github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/epoch"
	"github.com/flare-foundation/flare-system-client/config"
	"github.com/stretchr/testify/require"
)

func TestSetGas(t *testing.T) {
	chainCfg := config.Chain{
		EthRPCURL: "https://flare-api.flare.network/ext/C/rpc",
	}
	cl, err := chainCfg.DialETH()
	require.NoError(t, err)

	tests := []struct {
		gasConfig clconfig.Gas
	}{
		{
			gasConfig: clconfig.Gas{
				TxType: 0,
			},
		},
		{
			gasConfig: clconfig.Gas{
				TxType: 3,
			},
		},
		{
			gasConfig: clconfig.Gas{
				TxType:                2,
				MaxPriorityMultiplier: 2,
				BaseFeeMultiplier:     3,
				// Caps must be set: SetGas enforces them via
				// EnforceMaxPriorityFeeCaps, which dereferences these.
				MinimalMaxPriorityFee: big.NewInt(100e9),
				MaximalMaxPriorityFee: big.NewInt(5000e9),
			},
		}}

	for _, test := range tests {
		address := common.HexToAddress("0xdead")
		txOptions := new(bind.TransactOpts)
		txOptions.From = address

		err = epoch.SetGas(context.Background(), txOptions, cl, &test.gasConfig)

		require.Equal(t, address, txOptions.From)

		switch test.gasConfig.TxType {
		case 0:
			require.NoError(t, err)
			require.True(t, txOptions.GasPrice.Cmp(big.NewInt(1)) >= 0)
		case 2:
			require.NoError(t, err)
			require.True(t, txOptions.GasFeeCap.Cmp(big.NewInt(1)) >= 0)
			// The tip cap must lie within the configured bounds.
			require.True(t, txOptions.GasTipCap.Cmp(test.gasConfig.MinimalMaxPriorityFee) >= 0)
			require.True(t, txOptions.GasTipCap.Cmp(test.gasConfig.MaximalMaxPriorityFee) <= 0)
		default:
			require.Error(t, err)
		}
	}
}

// newBaseFeeRPCStub starts an HTTP JSON-RPC server that answers eth_baseFee with
// baseFeeWei, so SetGas's type-2 path can be exercised deterministically without
// a live node. SetGas only issues eth_baseFee for type 2.
func newBaseFeeRPCStub(t *testing.T, baseFeeWei *big.Int) *ethclient.Client {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		body, _ := io.ReadAll(r.Body)
		require.NoError(t, json.Unmarshal(body, &req))

		result := "0x0"
		if req.Method == "eth_baseFee" {
			result = (*hexutil.Big)(baseFeeWei).String()
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, result)
	}))
	t.Cleanup(srv.Close)

	cl, err := (&config.Chain{EthRPCURL: srv.URL}).DialETH()
	require.NoError(t, err)
	t.Cleanup(cl.Close)

	return cl
}

// TestSetGasEnforcesPriorityFeeCap verifies the fix in 74e0d1e: the type-2 tip
// cap is clamped to [MinimalMaxPriorityFee, MaximalMaxPriorityFee] before being
// added to the fee cap.
func TestSetGasEnforcesPriorityFeeCap(t *testing.T) {
	maximal := big.NewInt(5000e9) // 5000 Gwei
	minimal := big.NewInt(100e9)  // 100 Gwei

	baseGas := func() clconfig.Gas {
		return clconfig.Gas{
			TxType:                2,
			MaxPriorityMultiplier: 2,
			BaseFeeMultiplier:     4,
			MaximalMaxPriorityFee: new(big.Int).Set(maximal),
			MinimalMaxPriorityFee: new(big.Int).Set(minimal),
		}
	}

	t.Run("tip capped to maximal when base fee is high", func(t *testing.T) {
		baseFee := big.NewInt(1e14) // 100000 Gwei; *2 far exceeds maximal
		cl := newBaseFeeRPCStub(t, baseFee)

		g := baseGas()
		txOpts := new(bind.TransactOpts)
		require.NoError(t, epoch.SetGas(context.Background(), txOpts, cl, &g))

		require.Zero(t, maximal.Cmp(txOpts.GasTipCap), "tip must be capped to maximal, got %s", txOpts.GasTipCap)
		// GasFeeCap = baseFee*BaseFeeMultiplier + cappedTip
		wantFeeCap := new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(4)), maximal)
		require.Zero(t, wantFeeCap.Cmp(txOpts.GasFeeCap), "fee cap must be base*mult + tip, got %s want %s", txOpts.GasFeeCap, wantFeeCap)
	})

	t.Run("tip raised to minimal when base fee is low", func(t *testing.T) {
		baseFee := big.NewInt(1) // *2 = 2 wei, well below minimal
		cl := newBaseFeeRPCStub(t, baseFee)

		g := baseGas()
		txOpts := new(bind.TransactOpts)
		require.NoError(t, epoch.SetGas(context.Background(), txOpts, cl, &g))

		require.Zero(t, minimal.Cmp(txOpts.GasTipCap), "tip must be floored to minimal, got %s", txOpts.GasTipCap)
	})

	t.Run("config caps not mutated and not aliased", func(t *testing.T) {
		cl := newBaseFeeRPCStub(t, big.NewInt(1e14))

		g := baseGas()
		txOpts := new(bind.TransactOpts)
		require.NoError(t, epoch.SetGas(context.Background(), txOpts, cl, &g))

		require.Zero(t, maximal.Cmp(g.MaximalMaxPriorityFee), "MaximalMaxPriorityFee must be unchanged")
		require.Zero(t, minimal.Cmp(g.MinimalMaxPriorityFee), "MinimalMaxPriorityFee must be unchanged")

		// Mutating the produced tip cap must not corrupt the config caps.
		txOpts.GasTipCap.Add(txOpts.GasTipCap, big.NewInt(1))
		require.Zero(t, maximal.Cmp(g.MaximalMaxPriorityFee), "config cap must not be aliased into TransactOpts")
	})
}
