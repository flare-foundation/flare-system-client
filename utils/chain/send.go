package chain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

// Node responses arrive as RPC text, not typed errors, so they are matched as substrings.
const (
	nonceTooLowMsg  = "nonce too low" // a tx with the same nonce was already accepted
	alreadyKnownMsg = "already known" // the node already has this exact tx
)

// ReconcileLookupTimeout bounds each per-hash receipt/revert lookup during the
// relay's nonce-too-low reconciliation, so a hung RPC endpoint cannot stall a
// whole retry cycle (which would otherwise wait up to a full tx timeout per lookup).
const ReconcileLookupTimeout = 5 * time.Second

// definitiveRejectMsgs are broadcast errors where the node rejected the tx
// synchronously during validation, so it definitely never entered the mempool.
// Any OTHER broadcast error (timeout, connection reset, EOF, proxy 5xx, ...) is
// ambiguous — the tx may already have been accepted — so we must treat it as
// broadcast to avoid resending a duplicate.
var definitiveRejectMsgs = []string{
	nonceTooLowMsg,
	"replacement transaction underpriced",
	"transaction underpriced",
	"insufficient funds",
	"intrinsic gas too low",
	"exceeds block gas limit",
	"gas limit reached",
	"negative value",
	"oversized data",
	"invalid sender",
}

// mayBeInMempool reports whether a non-nil SendTransaction error leaves the tx
// possibly in the mempool (true) rather than definitively rejected (false).
func mayBeInMempool(err error) bool {
	return !MatchesError(err, definitiveRejectMsgs)
}

// Acceptance is the result of reconciling broadcast hashes against the chain.
type Acceptance int

const (
	// NonceConsumed: conclusively none of our broadcast hashes are on chain (we
	// broadcast nothing, or a foreign tx took the nonce), so the caller must move
	// to a fresh nonce.
	NonceConsumed Acceptance = iota
	Accepted                 // one of the hashes was accepted on chain
	Undetermined             // could not determine (RPC error / receipt not found); outcome unknown
	// Reverted: one of our hashes mined but reverted with a non-allowed reason, so
	// our tx consumed the nonce and a resend would deterministically revert again —
	// terminal, do not retry.
	Reverted
)

// SendResult reports the outcome of a single send so the caller can retry.
// Broadcast is the pre-/post-broadcast discriminator: false means the tx never
// reached the node (retry with a fresh nonce); true means it may be in the
// mempool (keep Hash for nonce-too-low reconciliation and reuse the nonce).
type SendResult struct {
	Hash      common.Hash // zero only on a pre-broadcast failure (never signed)
	Broadcast bool        // tx handed to the node (or send timed out, so maybe in mempool)
	Mined     bool        // a receipt was observed; with a non-nil Err the tx mined but reverted (nonce consumed)
	Err       error       // nil only when mined with a successful receipt
}

// IsNonceTooLow reports whether a tx with the same nonce was already accepted.
func IsNonceTooLow(err error) bool {
	return err != nil && strings.Contains(err.Error(), nonceTooLowMsg)
}

// IsTimeout reports whether err is (or wraps) one of our own context deadlines.
func IsTimeout(err error) bool {
	return errors.Is(err, context.DeadlineExceeded)
}

// MatchesError reports whether err's message contains any of the substrings
// (used to recognize per-path non-fatal errors, e.g. "Already relayed").
func MatchesError(err error, substrings []string) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	for _, s := range substrings {
		if strings.Contains(msg, s) {
			return true
		}
	}
	return false
}

func isAlreadyKnown(err error) bool {
	return err != nil && strings.Contains(err.Error(), alreadyKnownMsg)
}

// BroadcastAndWait broadcasts an already-signed transaction and waits for it to
// be mined, classifying the outcome for retry decisions and logging the hash of
// every broadcast tx. It is the shared post-signing half of a send.
func BroadcastAndWait(ctx context.Context, client *ethclient.Client, from common.Address, signedTx *types.Transaction, timeout time.Duration) SendResult {
	hash := signedTx.Hash()

	sendCtx, cancel := context.WithTimeout(ctx, timeout)
	err := client.SendTransaction(sendCtx, signedTx)
	cancel()
	if err != nil {
		if !isAlreadyKnown(err) {
			// Broadcast=true unless the node definitively rejected the tx, so an
			// ambiguous failure (timeout, transport error) keeps the hash for
			// nonce-too-low reconciliation instead of risking a duplicate.
			return SendResult{Hash: hash, Broadcast: mayBeInMempool(err), Err: fmt.Errorf("broadcasting tx %s: %w", hash.Hex(), err)}
		}
		// The node already has this exact tx: treat as broadcast, wait for it.
		logger.Infof("Broadcast tx %s (nonce %d): already known, waiting to be mined", hash.Hex(), signedTx.Nonce())
	} else {
		logger.Infof("Broadcast tx %s (nonce %d)", hash.Hex(), signedTx.Nonce())
	}

	verifier := NewTxVerifier(client)
	if err := verifier.WaitUntilMined(ctx, from, signedTx, timeout); err != nil {
		// Mined marks a reverted receipt (nonce consumed), unlike a pending timeout.
		return SendResult{Hash: hash, Broadcast: true, Mined: errors.Is(err, errReverted), Err: err}
	}
	logger.Debugf("Mined tx %s (nonce %d)", hash.Hex(), signedTx.Nonce())
	return SendResult{Hash: hash, Broadcast: true, Mined: true, Err: nil}
}

// AnyAccepted reconciles a "nonce too low" against the txs already broadcast in
// this send: if one made it on chain the send already succeeded and must not be
// repeated. It returns the accepting hash and a three-state Acceptance:
//   - Accepted: a hash was mined successfully, or (preserving the caller's
//     non-fatal semantics) mined reverting with one of allowedErrors;
//   - Reverted: one of our hashes mined but reverted with a non-allowed reason, so
//     our tx consumed the nonce and a resend would revert again — terminal;
//   - Undetermined: a lookup failed or a receipt was not found (a behind RPC
//     backend may not have the mined tx yet); duplicate-tolerant callers refresh
//     the nonce and resend for liveness;
//   - NonceConsumed: none of our hashes are on chain (nothing was broadcast, or a
//     foreign tx took the nonce), so the caller must retry at a fresh nonce.
func AnyAccepted(ctx context.Context, c Client, from common.Address, hashes []common.Hash, allowedErrors []string, timeout time.Duration) (common.Hash, Acceptance) {
	conclusive := true
	var revertedHash common.Hash
	reverted := false
	for _, h := range hashes {
		receipt, err := c.Receipt(ctx, h, timeout)
		if err != nil {
			logger.Warnf("checking receipt for tx %s: %v", h.Hex(), err)
			conclusive = false
			continue
		}
		if receipt == nil {
			// Not found on this backend. "nonce too low" means some tx at this nonce
			// mined, so our tx may have mined on a node this (behind) backend hasn't
			// caught up to. Stay inconclusive rather than risk resending a duplicate.
			conclusive = false
			continue
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			return h, Accepted
		}
		// Mined but reverted: honor the path's non-fatal errors, matching the way
		// those errors are treated when they surface on the current attempt.
		if len(allowedErrors) == 0 {
			logger.Warnf("tx %s was mined but reverted", h.Hex())
			revertedHash, reverted = h, true
			continue
		}
		reason, rerr := c.RevertReason(ctx, from, h, timeout)
		if rerr != nil {
			if errors.Is(rerr, errRevertUndecodable) {
				// Deterministically reverted with a reason we cannot decode: it
				// cannot match allowedErrors, so this hash is conclusively reverted.
				logger.Warnf("tx %s mined but reverted with an undecodable reason: %v", h.Hex(), rerr)
				revertedHash, reverted = h, true
				continue
			}
			// Transient lookup failure: outcome unknown, so the caller must not resend.
			logger.Warnf("getting revert reason for mined tx %s: %v", h.Hex(), rerr)
			conclusive = false
			continue
		}
		if MatchesError(errors.New(reason), allowedErrors) {
			logger.Infof("tx %s mined but reverted with non-fatal error %q; treating as accepted", h.Hex(), reason)
			return h, Accepted
		}
		logger.Warnf("tx %s was mined but reverted: %s", h.Hex(), reason)
		revertedHash, reverted = h, true
	}
	// A confirmed own-tx revert is deterministic, so it wins over an inconclusive
	// lookup (all broadcast hashes share one nonce; only one can have mined).
	if reverted {
		return revertedHash, Reverted
	}
	if !conclusive {
		return common.Hash{}, Undetermined
	}
	return common.Hash{}, NonceConsumed
}
