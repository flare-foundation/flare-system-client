package chain_test

import (
	"context"
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

	// if out off gas use faucet for coston 2
	testPrivateKey := "38f9137948fd4779212fa53fcdb0e41cfe8fa6c249c0e3c50994743f444aaded"
	pk, err := credentials.PrivateKeyFromHex(testPrivateKey)
	require.NoError(t, err)
	testPrivateAddress := "0xf52413dD9D7dDB8b4c9DAF249BF79De7a7821577"
	addr := common.HexToAddress(testPrivateAddress)

	// fmt.Printf("addr: %v\n", addr)

	deadAddress := "0x000000000000000000000000000000000000dead"
	toAddress := common.HexToAddress(deadAddress)

	// deadAddress2 := "0x00000000000000000000000000000000000dead22"
	// toAddress2 := common.HexToAddress(deadAddress2)

	gasConfig := config2.Gas{TxType: 2, MaxPriorityFeePerGas: big.NewInt(1)}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	nonce, err := cl.NonceAt(ctx, addr, nil)
	require.NoError(t, err)
	cancelFunc()

	err = chain.SendRawTx(cl, pk, nonce, toAddress, []byte{1, 2}, true, &gasConfig, 10*time.Second)

	require.NoError(t, err)
}

func TestGasConfigForAttemptType0(t *testing.T) {
	tests := []struct {
		name     string
		cfg      config2.Gas
		ri       int
		expected config2.Gas
	}{
		{
			name: "retry 0",
			cfg: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
			ri: 0,
			expected: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
		},
		{
			name: "retry 1",
			cfg: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
			ri: 1,
			expected: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.5,
			},
		},
		{
			name: "retry 1 - no config",
			cfg: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 0,
			},
			ri: 1,
			expected: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.5,
			},
		},
		{
			name: "retry 2",
			cfg: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 1.0,
			},
			ri: 2,
			expected: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(0),
				GasPriceMultiplier: 2.25,
			},
		},
		{
			name: "retry 1 - fixed",
			cfg: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(100),
				GasPriceMultiplier: 0,
			},
			ri: 1,
			expected: config2.Gas{
				TxType:             0,
				GasPriceFixed:      big.NewInt(100),
				GasPriceMultiplier: 0,
			},
		},
		{
			name: "empty type 0",
			cfg:  config2.Gas{},
			ri:   0,
			expected: config2.Gas{
				TxType:             0,
				GasPriceFixed:      nil,
				GasPriceMultiplier: 1,
			},
		},
		{
			name: "empty type 2",
			cfg:  config2.Gas{TxType: 2},
			ri:   0,
			expected: config2.Gas{
				TxType:               2,
				GasLimit:             0,
				MaxPriorityFeePerGas: big.NewInt(20_000_000_000),
				BaseFeeMultiplier:    big.NewInt(3),
				BaseFeePerGasCap:     nil,
			},
		},
		{
			name: "zero type 2",
			cfg: config2.Gas{
				TxType:               2,
				MaxPriorityFeePerGas: big.NewInt(0),
				BaseFeeMultiplier:    big.NewInt(0),
				BaseFeePerGasCap:     big.NewInt(0),
			},
			ri: 0,
			expected: config2.Gas{
				TxType:               2,
				GasLimit:             0,
				MaxPriorityFeePerGas: big.NewInt(20_000_000_000),
				BaseFeeMultiplier:    big.NewInt(3),
				BaseFeePerGasCap:     big.NewInt(0),
			},
		},
		{
			name: "zero type 2",
			cfg: config2.Gas{
				TxType:               2,
				MaxPriorityFeePerGas: big.NewInt(10_000_000_000),
				BaseFeeMultiplier:    big.NewInt(2),
				BaseFeePerGasCap:     big.NewInt(0),
			},
			ri: 2,
			expected: config2.Gas{
				TxType:               2,
				GasLimit:             0,
				MaxPriorityFeePerGas: big.NewInt(14_400_000_000),
				BaseFeeMultiplier:    big.NewInt(4),
				BaseFeePerGasCap:     big.NewInt(0),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := chain.GasConfigForAttempt(&test.cfg, test.ri)

			require.Equal(t, test.expected.TxType, got.TxType)

			if got.TxType == 0 {
				if got.GasPriceFixed.Cmp(test.expected.GasPriceFixed) != 0 {
					t.Errorf("GasPriceFixed = %v, want %v", got.GasPriceFixed, test.expected.GasPriceFixed)
				}
				if got.GasPriceMultiplier != test.expected.GasPriceMultiplier {
					t.Errorf("GasPriceMultiplier = %v, want %v", got.GasPriceMultiplier, test.expected.GasPriceMultiplier)
				}
			} else if got.TxType == 2 {
				require.Equal(t, test.expected.BaseFeeMultiplier, got.BaseFeeMultiplier)
				require.Equal(t, test.expected.MaxPriorityFeePerGas, got.MaxPriorityFeePerGas)
				require.Equal(t, test.expected.BaseFeePerGasCap, got.BaseFeePerGasCap)
			}
		})
	}
}
