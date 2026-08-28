package finalizer

import (
	"cmp"
	"context"
	"crypto/ecdsa"
	"fmt"
	"slices"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

const (
	listenerBufferSize = 10
)

// nonFatalRelayErrors are relay revert reasons that mean the finalization is
// already done on chain, so our tx failing that way is a success. Both relays
// are matched: only the old one can emit the string, only the new one the custom
// error, so the round's contract picks itself. "nonce too low" is handled
// separately (chain.IsNonceTooLow) via hash reconciliation.
var nonFatalRelayErrors = []string{
	"Already relayed",  // old Relay
	"AlreadyRelayed()", // new Relay
}

type relayContractClient struct {
	address      common.Address       // configured Relay, holds the signing policies before the cutover
	relayCutover *shared.RelayCutover // NewAddress serves the reward epochs from the cutover on

	chainClient chain.Client
	gasConfig   *config.Gas

	relay         *relay.Relay
	privateKey    *ecdsa.PrivateKey
	senderAddress common.Address

	relaySelector []byte      // for relay method
	topic0SPI     common.Hash // for SigningPolicyInitialized event
	topic0PMR     common.Hash // for ProtocolMessageRelayed event

	retryDelay time.Duration // backoff between send/nonce retries; tests shrink it
}

type signingPolicyListenerResponse struct {
	policyData *relay.RelaySigningPolicyInitialized
	timestamp  int64
}

func NewRelayContractClient(
	ethClient *ethclient.Client,
	address common.Address,
	privateKey *ecdsa.PrivateKey,
	senderAddress common.Address,
	gasConfig *config.Gas,
	chainID int64,
	relayCutover *shared.RelayCutover,
) (*relayContractClient, error) {
	relayContract, err := relay.NewRelay(address, ethClient)
	if err != nil {
		return nil, err
	}

	relayABI, err := relay.RelayMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	relaySelectorBytes := relayABI.Methods["relay"].ID

	topic0SPI, err := chain.EventIDFromMetadata(relay.RelayMetaData, "SigningPolicyInitialized")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	topic0PMR, err := chain.EventIDFromMetadata(relay.RelayMetaData, "ProtocolMessageRelayed")
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}

	return &relayContractClient{
		chainClient:   chain.NewClientImpl(ethClient, chainID),
		address:       address,
		relayCutover:  relayCutover,
		relay:         relayContract,
		privateKey:    privateKey,
		senderAddress: senderAddress,
		relaySelector: relaySelectorBytes,
		topic0SPI:     topic0SPI,
		topic0PMR:     topic0PMR,
		gasConfig:     gasConfig,
		retryDelay:    shared.TxRetryInterval,
	}, nil
}

// addressForRewardEpoch returns the Relay holding rewardEpochID's signing policy:
// only that one can verify a finalization signed under it.
func (r *relayContractClient) addressForRewardEpoch(rewardEpochID int64) common.Address {
	if r.relayCutover.NewRelayFromRewardEpoch(rewardEpochID) {
		return r.relayCutover.NewAddress
	}
	return r.address
}

// addresses lists the Relays to read events from — both across the cutover, since the old
// one emits only the events before it and the new one only those after.
func (r *relayContractClient) addresses() []common.Address {
	if !r.relayCutover.Scheduled() {
		return []common.Address{r.address}
	}
	return []common.Address{r.address, r.relayCutover.NewAddress}
}

// fetchLogs fetches topic0 logs of every relevant Relay in (from,to], in chain order.
func (r *relayContractClient) fetchLogs(ctx context.Context, db finalizerDB, topic0 common.Hash, from, to int64) ([]database.Log, error) {
	var all []database.Log
	for _, address := range r.addresses() {
		logs, err := db.FetchLogsByAddressAndTopic0(ctx, address, topic0, from, to)
		if err != nil {
			return nil, err
		}
		all = append(all, logs...)
	}
	sortLogs(all)
	return all, nil
}

// sortLogs restores chain order across merged per-address queries; callers rely on it to
// take the latest policy and to advance their event range.
func sortLogs(logs []database.Log) {
	slices.SortFunc(logs, func(a, b database.Log) int {
		if a.BlockNumber != b.BlockNumber {
			return cmp.Compare(a.BlockNumber, b.BlockNumber)
		}
		return cmp.Compare(a.LogIndex, b.LogIndex)
	})
}

// FetchSigningPolicies fetches signing policies emitted by in SigningPolicyInitialized events from Relay smart contract with timestamps in the interval (from,to].
func (r *relayContractClient) FetchSigningPolicies(ctx context.Context, db finalizerDB, from, to int64) ([]signingPolicyListenerResponse, error) {
	logs, err := r.fetchLogs(ctx, db, r.topic0SPI, from, to)
	if err != nil {
		return nil, err
	}

	result := make([]signingPolicyListenerResponse, 0, len(logs))
	for _, log := range logs {
		policyData, err := shared.ParseSigningPolicyInitializedEvent(r.relay, log)
		if err != nil {
			logger.Errorf("Error parsing SigningPolicyInitialized event %v", err)
			return nil, err
		}
		result = append(result, signingPolicyListenerResponse{policyData, int64(log.Timestamp)})
	}
	return result, nil
}

func (r *relayContractClient) SigningPolicyInitializedListener(ctx context.Context, db finalizerDB, startTime time.Time) <-chan signingPolicyListenerResponse {
	out := make(chan signingPolicyListenerResponse, listenerBufferSize)
	go func() {
		ticker := time.NewTicker(shared.EventListenerInterval)
		eventRangeStart := startTime.Unix()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			now := time.Now().Unix()

			logs, err := r.fetchLogs(ctx, db, r.topic0SPI, eventRangeStart, now)
			if err != nil {
				logger.Errorf("Error fetching logs %v", err)
				continue
			}

			for _, log := range logs {
				policyData, err := shared.ParseSigningPolicyInitializedEvent(r.relay, log)
				if err != nil {
					logger.Errorf("Error parsing SigningPolicyInitialized event %v", err)
					break
				}
				out <- signingPolicyListenerResponse{policyData, int64(log.Timestamp)}
				// continue with timestamps > log.Timestamp,
				// there should be only one such log per timestamp
				eventRangeStart = int64(log.Timestamp)
			}
		}
	}()
	return out
}

// SubmitPayloads sends a transaction with input to the Relay at address, with the same
// pre-/post-broadcast and nonce-too-low reconciliation logic as SubmitterBase.submit.
func (r *relayContractClient) SubmitPayloads(ctx context.Context, address common.Address, input []byte, dryRun bool, protocolID uint8, votingRoundID uint32) {
	if len(input) == 0 {
		logger.Warnf("Relay protocol %d round %d: empty tx input, nothing to send", protocolID, votingRoundID)
		return
	}
	if address == (common.Address{}) {
		logger.Errorf("Relay protocol %d round %d: no Relay address for this round, nothing to send", protocolID, votingRoundID)
		return
	}

	start := time.Now()

	// Fetch once (a stable nonce for the reconciliation below), with the full
	// send-retry budget so a transient RPC blip doesn't drop the finalization;
	// on a bounded ctx (grace path) the fetch stops sendPhaseReserve before the
	// deadline so the send attempts keep a usable window.
	nonceCtx := ctx
	if dl, ok := ctx.Deadline(); ok {
		fetchDeadline := dl.Add(-sendPhaseReserve)
		// never below one attempt's floor, or a short window couldn't fetch at all
		if earliest := time.Now().Add(minAttemptTimeout); fetchDeadline.Before(earliest) {
			fetchDeadline = earliest
		}
		var cancel context.CancelFunc
		nonceCtx, cancel = context.WithDeadline(ctx, fetchDeadline)
		defer cancel()
	}
	nonceResult := <-shared.ExecuteWithRetryChan(nonceCtx, func() (uint64, error) {
		return r.chainClient.Nonce(nonceCtx, r.privateKey, 2*time.Second)
	}, shared.MaxTxSendRetries, r.retryDelay)
	if !nonceResult.Success {
		logger.Errorf("Relay protocol %d round %d: getting nonce: %v", protocolID, votingRoundID, nonceResult.Message)
		return
	}
	nonce := nonceResult.Value

	// Fit boundedSendAttempts attempts into a bounded ctx (grace path) so gas
	// bumps still fire; without a deadline (delayed queue) keep the full timeout.
	perAttempt := chain.DefaultTxTimeout
	if dl, ok := ctx.Deadline(); ok {
		// budget the inter-attempt sleeps too, or the last attempt's wait is cut short
		budget := time.Until(dl) - (boundedSendAttempts-1)*r.retryDelay
		if v := budget / boundedSendAttempts; v < perAttempt {
			perAttempt = max(v, minAttemptTimeout)
		}
	}

	var broadcastHashes []common.Hash
	attempts := 0

	sendResult := <-shared.ExecuteWithRetryAttempts(ctx, func(ri int) (string, error) {
		attempts = ri + 1
		gasConfig := chain.GasConfigForAttempt(r.gasConfig, ri)
		p := relaySendPrefix(protocolID, votingRoundID, ri, nonce) // holds the nonce this attempt used, even after a refresh below
		logger.Debugf("%s: sending tx, timeout %s, dry_run=%t, gas: %s", p, perAttempt, dryRun, gasConfig)

		// one slice per attempt — SendRawTx spends timeout per phase, uncapped that's ~3 slices
		attemptCtx, cancelAttempt := context.WithTimeout(ctx, perAttempt)
		res := r.chainClient.SendRawTx(attemptCtx, r.privateKey, nonce, address, input, gasConfig, perAttempt, dryRun)
		cancelAttempt()
		if res.Broadcast {
			broadcastHashes = append(broadcastHashes, res.Hash)
		}

		switch {
		case res.Err == nil:
			return "confirmed " + res.Hash.Hex(), nil
		case chain.MatchesError(res.Err, nonFatalRelayErrors):
			return fmt.Sprintf("already relayed (non-fatal): %v", res.Err), nil
		case res.Mined:
			// Mined-reverted is deterministic ("Already relayed" handled above); a resend can't help.
			logger.Warnf("%s: tx %s mined but reverted, not retrying: %v", p, res.Hash.Hex(), res.Err)
			return "reverted " + res.Hash.Hex(), nil
		case chain.IsNonceTooLow(res.Err):
			h, acc := chain.AnyAccepted(ctx, r.chainClient, r.senderAddress, broadcastHashes, nonFatalRelayErrors, chain.ReconcileLookupTimeout)
			switch acc {
			case chain.Accepted:
				return fmt.Sprintf("reconciled, earlier broadcast %s accepted (nonce too low is non-fatal)", h.Hex()), nil
			case chain.Reverted:
				// Our prior broadcast mined but reverted — same terminal handling as
				// a direct mined-revert (deterministic; a resend can't help).
				logger.Warnf("%s: own tx %s mined but reverted (reconciled), not retrying", p, h.Hex())
				return "reverted " + h.Hex(), nil
			case chain.Undetermined:
				// Outcome unknown: refresh the nonce and resend (a duplicate reverts
				// non-fatally with "Already relayed").
				nonce = r.refreshNonce(ctx, p,
					fmt.Sprintf("rejected as nonce too low, fate of own broadcast(s) %s unresolved", chain.HashList(broadcastHashes)), nonce)
				return "", res.Err
			default: // NonceConsumed
				nonce = r.refreshNonce(ctx, p, "nonce consumed by another tx", nonce)
				return "", res.Err
			}
		case res.Broadcast && chain.IsTimeout(res.Err):
			// Keep the nonce so the gas-bumped retry replaces the pending tx.
			logger.Warnf("%s: no confirmation of tx %s within %s; retrying as a gas-bumped replacement at the same nonce",
				p, res.Hash.Hex(), perAttempt)
			return "", res.Err
		default:
			logger.Warnf("%s: send failed (broadcast=%t mined=%t): %v", p, res.Broadcast, res.Mined, res.Err)
			// Keep the nonce if a tx is already outstanding at it (retry replaces);
			// only refresh when nothing has been broadcast yet.
			if len(broadcastHashes) == 0 {
				nonce = r.refreshNonce(ctx, p, "nothing broadcast at this nonce", nonce)
			}
			return "", res.Err
		}
	}, shared.MaxTxSendRetries, r.retryDelay)

	elapsed := time.Since(start).Round(time.Millisecond)
	switch {
	case sendResult.Success:
		// nonce is the terminal attempt's — no terminal branch refreshes it
		logger.Infof("Relay protocol %d round %d: finished in %d attempt(s), %s, nonce %d, outcome: %s",
			protocolID, votingRoundID, attempts, elapsed, nonce, sendResult.Value)
	case len(broadcastHashes) > 0:
		// an outstanding broadcast can still mine: not a confirmed failure
		logger.Warnf("Relay protocol %d round %d: outcome unknown after %d attempt(s), %s; broadcast tx(s) %s may be on chain: %v",
			protocolID, votingRoundID, attempts, elapsed, chain.HashList(broadcastHashes), sendResult.Message)
	default:
		logger.Errorf("Relay protocol %d round %d: failed after %d attempt(s), %s: %v",
			protocolID, votingRoundID, attempts, elapsed, sendResult.Message)
	}
}

func relaySendPrefix(protocolID uint8, votingRoundID uint32, attempt int, nonce uint64) string {
	return fmt.Sprintf("Relay protocol %d round %d attempt %d/%d nonce %d",
		protocolID, votingRoundID, attempt+1, shared.MaxTxSendRetries, nonce)
}

// refreshNonce best-effort re-fetches the sender nonce, keeping current on error.
// It logs reason (why the resend happens) together with the outcome, so a refresh
// costs one line. Nonce reads the latest mined nonce, so an accepted-but-unmined tx
// (or a lagging backend) reads back the rejected nonce — hence the unchanged case.
func (r *relayContractClient) refreshNonce(ctx context.Context, prefix, reason string, current uint64) uint64 {
	nonce, err := r.chainClient.Nonce(ctx, r.privateKey, time.Second)
	switch {
	case err != nil:
		logger.Warnf("%s: %s; nonce refresh failed, resending at %d: %v", prefix, reason, current, err)
		return current
	case nonce == current:
		logger.Warnf("%s: %s; nonce refresh returned %d unchanged, resending at it", prefix, reason, current)
		return current
	default:
		logger.Warnf("%s: %s; resending at refreshed nonce %d -> %d", prefix, reason, current, nonce)
		return nonce
	}
}

// relayedKey is the lookup key for ProtocolMessageRelayed events.
//
// It deliberately excludes the seed and digest fields of queueItem:
// ProtocolMessageRelayed events identify a finalization uniquely by
// (protocolID, votingRoundID), and using queueItem directly as a map
// key would compare *big.Int by pointer identity — guaranteeing the
// lookup in processDelayedQueue never matches.
type relayedKey struct {
	protocolID    uint8
	votingRoundID uint32
}

// relayedSet holds the finalizations seen on chain, per Relay: a round relayed only on
// the old Relay still has to be sent to the new one.
type relayedSet map[common.Address]map[relayedKey]bool

func (s relayedSet) has(address common.Address, key relayedKey) bool {
	return s[address][key]
}

// ProtocolMessageRelayed returns, per Relay, the (protocolID, votingRoundID) pairs
// already finalized on chain in the given time range.
func (r *relayContractClient) ProtocolMessageRelayed(ctx context.Context, db finalizerDB, from time.Time, to time.Time) (relayedSet, error) {
	result := make(relayedSet)
	for _, address := range r.addresses() {
		logs, err := db.FetchLogsByAddressAndTopic0(ctx, address, r.topic0PMR, from.Unix(), to.Unix())
		if err != nil {
			return nil, err
		}

		relayed := make(map[relayedKey]bool)
		for _, log := range logs {
			data, err := shared.ParseProtocolMessageRelayedEvent(r.relay, log)
			if err != nil {
				return nil, err
			}
			relayed[relayedKey{
				protocolID:    data.ProtocolId,
				votingRoundID: data.VotingRoundId,
			}] = true
		}
		result[address] = relayed
	}
	return result, nil
}
