package chain_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	config2 "github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/config"
	"github.com/flare-foundation/flare-system-client/utils/chain"
	"github.com/flare-foundation/flare-system-client/utils/credentials"
	"github.com/stretchr/testify/require"
)

func TestSendTx(t *testing.T) {
	chainCfg := config.Chain{
		EthRPCURL: "https://coston2-api.flare.network/ext/C/rpc",
	}
	cl, err := chainCfg.DialETH()
	require.NoError(t, err)

	testPrivateKey := "38f9137948fd4779212fa53fcdb0e41cfe8fa6c249c0e3c50994743f444aaded"

	pk, err := credentials.PrivateKeyFromHex(testPrivateKey)

	require.NoError(t, err)

	testPrivateAddress := "0xf52413dD9D7dDB8b4c9DAF249BF79De7a7821577"

	addr := common.HexToAddress(testPrivateAddress)

	fmt.Printf("addr: %v\n", addr)

	deadAddress := "0x000000000000000000000000000000000000dead"
	toAddress := common.HexToAddress(deadAddress)

	// deadAddress2 := "0x00000000000000000000000000000000000dead22"
	// toAddress2 := common.HexToAddress(deadAddress2)

	gasConfig := config2.Gas{TxType: 2, MaxPriorityFeePerGas: big.NewInt(1)}

	gasConfig0 := config2.Gas{
		TxType:             0,
		GasPriceMultiplier: 2,
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	nonce, err := cl.NonceAt(ctx, addr, nil)
	require.NoError(t, err)
	cancelFunc()
	go func() {
		err = chain.SendRawTx(cl, pk, nonce+1, toAddress, []byte{1, 2}, true, &gasConfig0, 3*time.Second)

		fmt.Printf("err: %v\n", err)
	}()

	err = chain.SendRawTx(cl, pk, nonce-1, toAddress, []byte{1, 2}, true, &gasConfig, 3*time.Second)

	fmt.Printf("err: %v\n", err)
}
