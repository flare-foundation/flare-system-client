package chain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const (
	// default timeout for waiting for a tx to be mined.
	DefaultTxTimeout = 60 * time.Second
	// timeout for waiting for a non-time-sensitive tx (uptime vote, reward signing) to be
	// mined. We wait much longer for a single modestly-priced tx to confirm rather than
	// timing out and re-sending duplicate transactions.
	LongTxTimeout   = 30 * time.Minute
	DefaultGasLimit = 2_500_000

	multiplierBumpTimes100 = 111
	normalizer             = 100

	// receipt polling: fast while a tx normally mines (~1-2s Flare blocks), backed
	// off for long waits (LongTxTimeout) so a stuck tx is not polled at 2.5 Hz for 30min.
	txPollInterval     = 400 * time.Millisecond
	txPollIntervalSlow = time.Second
	txPollBackoffAfter = 10 * time.Second
)

// CopyTxOpts returns a copy of opts that can be mutated per transaction
// (gas limit, gas price, fee caps) without racing with other goroutines
// sharing the original TransactOpts instance.
func CopyTxOpts(opts *bind.TransactOpts) *bind.TransactOpts {
	cp := *opts
	if opts.Nonce != nil {
		cp.Nonce = new(big.Int).Set(opts.Nonce)
	}
	if opts.Value != nil {
		cp.Value = new(big.Int).Set(opts.Value)
	}
	if opts.GasPrice != nil {
		cp.GasPrice = new(big.Int).Set(opts.GasPrice)
	}
	if opts.GasFeeCap != nil {
		cp.GasFeeCap = new(big.Int).Set(opts.GasFeeCap)
	}
	if opts.GasTipCap != nil {
		cp.GasTipCap = new(big.Int).Set(opts.GasTipCap)
	}
	return &cp
}

// txBackend is the ethclient subset TxVerifier uses: receipt polling and
// eth_call replay for revert reasons.
type txBackend interface {
	ethereum.ContractCaller
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type TxVerifier struct {
	eth txBackend

	pollInterval     time.Duration // receipt poll interval; tests shrink these
	pollIntervalSlow time.Duration // interval after pollBackoffAfter of polling
	pollBackoffAfter time.Duration
}

func NewTxVerifier(eth *ethclient.Client) *TxVerifier {
	return &TxVerifier{eth: eth, pollInterval: txPollInterval,
		pollIntervalSlow: txPollIntervalSlow, pollBackoffAfter: txPollBackoffAfter}
}

// waitMined polls for the tx receipt until found or ctx expires. Ported from
// go-ethereum v1.17.3 accounts/abi/bind/v2/util.go WaitMined (what bind.WaitMined
// delegates to) — apply upstream fixes there here too. Only the timing differs:
// upstream's fixed 1s ticker adds ~0.5s mean latency on ~1-2s Flare blocks. Like
// upstream it keeps polling on any error (transient RPC failures included) and
// checks once before the first wait.
func (t TxVerifier) waitMined(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	start := time.Now()
	timer := time.NewTimer(0)
	defer timer.Stop()
	for {
		receipt, err := t.eth.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		if !errors.Is(err, ethereum.NotFound) {
			logger.Debugf("Receipt retrieval for tx %s failed: %v", txHash.Hex(), err)
		}
		interval := t.pollInterval
		if time.Since(start) >= t.pollBackoffAfter {
			interval = t.pollIntervalSlow
		}
		timer.Reset(interval)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}

func (t TxVerifier) WaitUntilMined(ctx context.Context, from common.Address, tx *types.Transaction, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	receipt, err := t.waitMined(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("waitMined: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		// Wrap errReverted whether or not the reason decodes: the tx still mined.
		reason, rerr := errorReason(ctx, t.eth, from, tx, receipt.BlockNumber)
		if rerr != nil {
			return fmt.Errorf("%w: %v", errReverted, rerr)
		}
		return fmt.Errorf("%w: %s", errReverted, reason)
	}
	return nil
}

// errReverted marks a mined-but-reverted tx (nonce consumed; a resend is futile).
// Callers detect it via SendResult.Mined.
var errReverted = errors.New("tx mined but reverted")

// errRevertUndecodable marks a deterministically-reverted tx whose revert reason
// could not be decoded (e.g. a custom error, not Error(string)). Callers
// reconciling acceptance must treat it as conclusively reverted, not as a
// transient RPC failure worth retrying.
var errRevertUndecodable = errors.New("revert reason not decodable")

// Taken from: https://ethereum.stackexchange.com/questions/48383/how-to-retrieve-revert-reason-for-past-transactions
func errorReason(ctx context.Context, b ethereum.ContractCaller, from common.Address, tx *types.Transaction, blockNum *big.Int) (string, error) {
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	res, err := b.CallContract(ctx, msg, blockNum)
	if err != nil {
		// geth/coreth return a reverting eth_call as a JSON-RPC error carrying the
		// ABI-encoded revert data in ErrorData(), not as return bytes; recover it.
		if revert, ok := revertDataFromError(err); ok {
			return decodeRevert(revert)
		}
		return "", fmt.Errorf("CallContract: %w", err)
	}
	return decodeRevert(res)
}

// decodeRevert unpacks ABI-encoded revert data into its reason string: an
// Error(string) payload, or a known custom-error selector rendered as its
// signature. Anything else is undecodable but deterministic, so the error wraps
// errRevertUndecodable to separate it from a transient RPC failure.
func decodeRevert(data []byte) (string, error) {
	reason, err := unpackError(data)
	if err == nil {
		return reason, nil
	}
	if name, ok := relayErrorName(data); ok {
		return name, nil
	}
	return reason, fmt.Errorf("%w: %w", errRevertUndecodable, err)
}

// annotateRevert appends the decoded revert reason to a reverting eth_call /
// eth_estimateGas error, which otherwise reads "execution reverted" with the
// reason only in its data. It appends rather than replaces: callers match the
// node's own text (e.g. the non-fatal lists in client/epoch).
func annotateRevert(err error) error {
	data, ok := revertDataFromError(err)
	if !ok {
		return err
	}
	reason, decErr := decodeRevert(data)
	if decErr != nil || strings.Contains(err.Error(), reason) {
		return err
	}
	return fmt.Errorf("%w: %s", err, reason)
}

// revertDataFromError extracts ABI-encoded revert data from a JSON-RPC error
// (rpc.DataError) as returned by a reverting eth_call. ok is false for a plain
// transport/RPC error, which carries no revert data.
func revertDataFromError(err error) ([]byte, bool) {
	var de rpc.DataError
	if !errors.As(err, &de) {
		return nil, false
	}
	switch data := de.ErrorData().(type) {
	case string:
		raw, decErr := hexutil.Decode(data)
		if decErr != nil {
			return nil, false
		}
		return raw, true
	case hexutil.Bytes:
		return data, true
	case []byte:
		return data, true
	default:
		return nil, false
	}
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

func unpackError(result []byte) (string, error) {
	if len(result) < 4 || !bytes.Equal(result[:4], errorSig) {
		return "<tx result not Error(string)>", fmt.Errorf("unknown error signature: %x", result)
	}
	vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "<invalid tx result>", fmt.Errorf("unpacking revert reason: %w", err)
	}

	errStr, ok := vs[0].(string)
	if !ok {
		return "<invalid tx result>", errors.New("unexpected error type")
	}

	return errStr, nil
}

// BaseFee returns the BaseFee per gas if the block was mined immediately.
//
// WORKS ONLY ON AVALANCHE C-CHAIN LIKE CHAINS
func BaseFee(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	var result hexutil.Big
	err := client.Client().CallContext(ctx, &result, "eth_baseFee")
	return (*big.Int)(&result), err
}

// SendRawTx signs a transaction to toAddress with the prescribed nonce,
// gasConfig and EIP-155 chainID, broadcasts it and waits for it to be mined.
// SendResult classifies the outcome (pre-broadcast failure vs post-broadcast
// timeout) and carries the broadcast hash for nonce-too-low reconciliation.
func SendRawTx(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey, chainID *big.Int, nonce uint64, toAddress common.Address, data []byte, dryRun bool, gasConfig *config.Gas, timeout time.Duration) SendResult {
	signedTx, fromAddress, err := buildAndSignRawTx(ctx, client, privateKey, chainID, nonce, toAddress, data, dryRun, gasConfig, timeout)
	if err != nil {
		// Failed before broadcast: never reached the node, so Broadcast stays false.
		return SendResult{Err: fmt.Errorf("preparing tx: %w", err)}
	}
	return BroadcastAndWait(ctx, client, fromAddress, signedTx, timeout)
}

// buildAndSignRawTx does the pre-broadcast work (gas limit and fee reads,
// signing; dry-running when dryRun is set). Any error it returns is a
// pre-broadcast failure — nothing was sent to the network.
func buildAndSignRawTx(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey, chainID *big.Int, nonce uint64, toAddress common.Address, data []byte, dryRun bool, gasConfig *config.Gas, timeout time.Duration) (*types.Transaction, common.Address, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	value := big.NewInt(0)

	if gasConfig.TxType != 0 && gasConfig.TxType != 2 {
		return nil, fromAddress, errors.New("unsupported tx type: set TxType to 0 or 2")
	}

	// Gas limit and fee data are independent reads — fetch them concurrently so
	// a slow node costs one stage timeout, not two.
	type gasLimitResult struct {
		gasLimit uint64
		err      error
	}
	gasLimitCh := make(chan gasLimitResult, 1)
	go func() {
		gasLimit, err := resolveGasLimit(ctx, client, gasConfig, fromAddress, toAddress, value, data, dryRun, timeout)
		gasLimitCh <- gasLimitResult{gasLimit, err}
	}()

	var feePerGas *big.Int // gas price (type 0) or base fee (type 2)
	var feeErr error
	switch gasConfig.TxType {
	case 0:
		feePerGas, feeErr = GetGasPrice(ctx, gasConfig, client, timeout)
	case 2:
		feeCtx, cancelFunc := context.WithTimeout(ctx, timeout)
		feePerGas, feeErr = BaseFee(feeCtx, client)
		cancelFunc()
	}

	gl := <-gasLimitCh
	// dry-run errors first — they carry the tx's own revert reason, fee errors are transport
	if gl.err != nil {
		return nil, fromAddress, gl.err
	}
	if feeErr != nil {
		return nil, fromAddress, feeErr
	}

	var signedTx *types.Transaction
	var err error
	switch gasConfig.TxType {
	case 0:
		signedTx, err = prepareAndSignType0(privateKey, chainID, nonce, gl.gasLimit, feePerGas, toAddress, value, data)
	case 2:
		signedTx, err = prepareAndSignType2(gasConfig, privateKey, chainID, nonce, gl.gasLimit, feePerGas, toAddress, value, data)
	}
	if err != nil {
		return nil, fromAddress, err
	}

	return signedTx, fromAddress, nil
}

// resolveGasLimit returns the tx gas limit: with dryRun the estimate (or, with a
// configured limit, that limit after dry-run validation), otherwise the
// configured/estimated limit, which never fails (falls back to a default).
func resolveGasLimit(ctx context.Context, client *ethclient.Client, gasConfig *config.Gas, fromAddress, toAddress common.Address, value *big.Int, data []byte, dryRun bool, timeout time.Duration) (uint64, error) {
	if !dryRun {
		return getGasLimit(ctx, gasConfig, client, fromAddress, toAddress, value, data, timeout), nil
	}
	gasLimit, err := DryRunTx(ctx, client, fromAddress, toAddress, value, data, timeout)
	if err != nil {
		return 0, fmt.Errorf("dry run: %w", err)
	}
	if gasConfig.GasLimit > 0 {
		gasLimit = uint64(gasConfig.GasLimit)
	}
	return gasLimit, nil
}

// prepareAndSignType0 builds and signs a type 0 (legacy) transaction from the
// prefetched gasPrice.
func prepareAndSignType0(privateKey *ecdsa.PrivateKey, chainID *big.Int, nonce uint64, gasLimit uint64, gasPrice *big.Int, toAddress common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	logger.Debugf("built tx nonce=%d tx_type=0 gas_limit=%d gas_price_gwei=%s", nonce, gasLimit, utils.Gwei(gasPrice))

	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	}

	tx := types.NewTx(&txData)
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// tipClampNote reports which cap bound the tip, "" when it was not clamped. A tip
// pinned below market times out every attempt with nothing else saying why.
func tipClampNote(wanted, enforced *big.Int, cfg *config.Gas) string {
	if wanted.Cmp(enforced) == 0 {
		return ""
	}
	if wanted.Cmp(cfg.MinimalMaxPriorityFee) < 0 {
		return "(clamped:minimal_max_priority_fee)"
	}
	return "(clamped:maximal_max_priority_fee)"
}

// prepareAndSignType2 builds and signs a type 2 (eip 1559) transaction from the
// prefetched baseFeePerGas.
func prepareAndSignType2(gasConfig *config.Gas, privateKey *ecdsa.PrivateKey, chainID *big.Int, nonce uint64, gasLimit uint64, baseFeePerGas *big.Int, toAddress common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	// Default unset fields so EnforceMaxPriorityFeeCaps never sees a nil cap,
	// even if a caller passes a raw config.
	cfg := gasConfig.CopyAndDefault()

	gasFeeCap := new(big.Int)
	baseFeeSource := "multiplier"
	if cfg.BaseFeePerGasCap != nil && cfg.BaseFeePerGasCap.Sign() == 1 {
		// pinned: the per-attempt bump no longer moves the base-fee component
		gasFeeCap.Set(cfg.BaseFeePerGasCap)
		baseFeeSource = "base_fee_per_gas_cap"
	} else {
		gasFeeCap = MultiplyWithFloat(baseFeePerGas, float64(cfg.BaseFeeMultiplier), gasFeeCap)
	}

	wantedTipCap := MultiplyWithFloat(baseFeePerGas, float64(cfg.MaxPriorityMultiplier), nil)
	gasTipCap := cfg.EnforceMaxPriorityFeeCaps(wantedTipCap)

	gasFeeCap.Add(gasFeeCap, gasTipCap)

	logger.Debugf("built tx nonce=%d tx_type=2 gas_limit=%d base_fee_gwei=%s tip_cap_gwei=%s%s fee_cap_gwei=%s (base fee from %s)",
		nonce, gasLimit, utils.Gwei(baseFeePerGas), utils.Gwei(gasTipCap),
		tipClampNote(wantedTipCap, gasTipCap, cfg), utils.Gwei(gasFeeCap), baseFeeSource)

	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	}

	tx := types.NewTx(&txData)
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// DryRunTx locally executes a transaction with the current state of blockchain and returns estimated Gas multiplied with 1.5 and potential errors.
func DryRunTx(ctx context.Context, client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, data []byte, timeout time.Duration) (uint64, error) {
	estCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	estimatedGas, err := estimateGas(estCtx, client, fromAddress, toAddress, value, data)
	cancelFunc()
	return 3 * estimatedGas / 2, err
}

// DryRunTxAbi locally executes a transaction to method with arguments with the current state of blockchain and returns estimated Gas multiplied with 1.5 and potential errors.
func DryRunTxAbi(ctx context.Context, client *ethclient.Client, timeout time.Duration, fromAddress common.Address, toAddress common.Address, value *big.Int, abi *abi.ABI, method string, arguments ...any) (uint64, error) {
	data, err := abi.Pack(method, arguments...)
	if err != nil {
		return 0, fmt.Errorf("DryRunTxAbi packing: %w", err)
	}
	estCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	estimatedGas, err := estimateGas(estCtx, client, fromAddress, toAddress, value, data)
	cancelFunc()
	return 3 * estimatedGas / 2, err
}

func estimateGas(ctx context.Context, client *ethclient.Client, from, to common.Address, value *big.Int, data []byte) (uint64, error) {
	gas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  from,
		To:    &to,
		Value: value,
		Data:  data,
	})
	return gas, annotateRevert(err)
}

func getGasLimit(ctx context.Context, gasConfig *config.Gas, client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, data []byte, timeout time.Duration) uint64 {
	var gasLimit uint64
	if gasConfig.GasLimit == 0 {
		estCtx, cancelFunc := context.WithTimeout(ctx, timeout)
		estimatedGas, err := estimateGas(estCtx, client, fromAddress, toAddress, value, data)
		cancelFunc()
		if err != nil {
			// estimation usually fails because the call would revert: the tx still sends and burns the nonce
			logger.Warnf("Unable to estimate gas (the tx may revert): %v; using default gas limit: %d", err, DefaultGasLimit)
			gasLimit = DefaultGasLimit
		} else {
			gasLimit = 3 * estimatedGas / 2
		}
	} else {
		gasLimit = uint64(gasConfig.GasLimit)
	}
	return gasLimit
}

func GetGasPrice(ctx context.Context, gasConfig *config.Gas, client *ethclient.Client, timeout time.Duration) (*big.Int, error) {
	var gasPrice *big.Int
	if gasConfig.GasPriceFixed != nil && gasConfig.GasPriceFixed.Cmp(common.Big0) == 1 {
		gasPrice = gasConfig.GasPriceFixed
	} else {
		priceCtx, cancelFunc := context.WithTimeout(ctx, timeout)
		suggestedPrice, err := client.SuggestGasPrice(priceCtx)
		cancelFunc()
		if err != nil {
			return nil, fmt.Errorf("estimating gas price: %w", err)
		}
		if gasConfig.GasPriceMultiplier != 0 {
			gasPriceFloat := new(big.Float).SetInt(suggestedPrice)
			gasPriceMultiplierFloat := new(big.Float).SetFloat64(float64(gasConfig.GasPriceMultiplier))
			gasPriceFloat.Mul(gasPriceFloat, gasPriceMultiplierFloat)
			gasPrice, _ = gasPriceFloat.Int(nil)
		} else {
			gasPrice = suggestedPrice
		}
	}
	return gasPrice, nil
}

// GasConfigForAttempt sets gas config for a retry attempt.
//
// For type 0 transaction, it bumps up GasPriceMultiplier for each retry attempt by 50%,
// up to a maximum of 10x the original value.
// If GasPriceFixed is used, the retry multiplier will not be applied.
//
// For type 2 transaction on i-th attempt,
// the multipliers are increased by i, and caps are increased by 11% per attempt.
//
// Only the type 2 branch defaults unset values; type 0 returns the config as configured
// (the fixed-price branch returns the caller's own pointer).
func GasConfigForAttempt(cfg *config.Gas, attempt int) *config.Gas {
	switch cfg.TxType {
	case 0:
		if cfg.GasPriceFixed != nil && cfg.GasPriceFixed.Cmp(common.Big0) != 0 {
			return cfg
		}

		retryMultiplier := min(10.0, math.Pow(1.5, float64(attempt)))

		return &config.Gas{
			TxType:   0,
			GasLimit: cfg.GasLimit,

			GasPriceMultiplier: max(1.0, cfg.GasPriceMultiplier) * float32(retryMultiplier),
			GasPriceFixed:      cfg.GasPriceFixed,
		}
	default: // type 2 and invalid types
		attemptBig := big.NewInt(int64(attempt))

		c := cfg.CopyAndDefault()

		c.MaxPriorityMultiplier += config.Multiplier(attempt)
		c.BaseFeeMultiplier += config.Multiplier(attempt)

		// Increase caps by 11% per retry
		if attempt > 0 {
			multiplier := new(big.Int).Exp(big.NewInt(multiplierBumpTimes100), attemptBig, nil)
			normalizer := new(big.Int).Exp(big.NewInt(normalizer), attemptBig, nil)

			c.MaximalMaxPriorityFee.Mul(c.MaximalMaxPriorityFee, multiplier)
			c.MaximalMaxPriorityFee.Div(c.MaximalMaxPriorityFee, normalizer)

			c.MinimalMaxPriorityFee.Mul(c.MinimalMaxPriorityFee, multiplier)
			c.MinimalMaxPriorityFee.Div(c.MinimalMaxPriorityFee, normalizer)
		}

		return c
	}
}

// MultiplyWithFloat multiplies x by y and truncates the result to an integer.
// The result is stored in out (allocated if nil) and returned. x is not modified.
// y must be a finite number.
func MultiplyWithFloat(x *big.Int, y float64, out *big.Int) *big.Int {
	if out == nil {
		out = new(big.Int)
	}

	outFloat := new(big.Float).Mul(big.NewFloat(y), new(big.Float).SetInt(x))

	out, _ = outFloat.Int(out)

	return out
}
