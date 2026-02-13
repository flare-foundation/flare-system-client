package epoch_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	clconfig "github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/epoch"
	"github.com/flare-foundation/flare-system-client/config"
	"github.com/stretchr/testify/require"
)

func TestSetGas(t *testing.T) {
	chainCfg := config.Chain{
		EthRPCURL: "https://coston2-api.flare.network/ext/C/rpc",
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
				MaxPriorityMultiplier: big.NewInt(2),
				BaseFeeMultiplier:     big.NewInt(3),
			},
		}}

	for _, test := range tests {
		address := common.HexToAddress("0xdead")
		txOptions := new(bind.TransactOpts)
		txOptions.From = address

		err = epoch.SetGas(txOptions, cl, &test.gasConfig)

		require.Equal(t, address, txOptions.From)

		switch test.gasConfig.TxType {
		case 0:
			require.NoError(t, err)
			require.True(t, txOptions.GasPrice.Cmp(big.NewInt(25*1e9)) >= 0)
		case 2:
			require.NoError(t, err)
			require.True(t, txOptions.GasFeeCap.Cmp(big.NewInt(25*1e9)) >= 0)
		default:
			require.Error(t, err)
		}
	}
}
