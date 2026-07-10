package finalizer

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils/chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

const (
	listenerBufferSize = 10
)

// nonFatalRelayErrors are relay revert reasons that mean the finalization is
// already done on chain, so our tx failing that way is a success. "nonce too
// low" is handled separately (chain.IsNonceTooLow) via hash reconciliation.
var nonFatalRelayErrors = []string{
	"Already relayed",
}

var (
	RelayFlareOld          = common.HexToAddress("0x57a4c3676d08Aa5d15410b5A6A80fBcEF72f3F45")
	RelayFlareNew          = common.HexToAddress("0xCcF30790A93F15e24EB909548a2C58a9b0a7FBd4")
	RewardEpochChangeFlare = int64(374)

	RelayCoston2Old          = common.HexToAddress("0x97702e350CaEda540935d92aAf213307e9069784")
	RelayCoston2New          = common.HexToAddress("0xa10B672D1c62e5457b17af63d4302add6A99d7dE")
	RewardEpochChangeCoston2 = int64(5236)

	RelaySongbirdOld          = common.HexToAddress("0x67a916E175a2aF01369294739AA60dDdE1Fad189")
	RelaySongbirdNew          = common.HexToAddress("0xCB86E8Be709001e01897Bf59847406853da8f14b")
	RewardEpochChangeSongbird = int64(374)

	RelayCostonOld          = common.HexToAddress("0x92a6E1127262106611e1e129BB64B6D8654273F7")
	RelayCostonNew          = common.HexToAddress("0x051f214D346Cfd97B107BECb87E2B35D1b4287E9")
	RewardEpochChangeCoston = int64(5236)
)

type relayContractClient struct {
	address common.Address

	chainClient chain.Client
	gasConfig   *config.Gas

	relay         *relay.Relay
	privateKey    *ecdsa.PrivateKey
	senderAddress common.Address

	relaySelector []byte      // for relay method
	topic0SPI     common.Hash // for SigningPolicyInitialized event
	topic0PMR     common.Hash // for ProtocolMessageRelayed event
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
		chainClient:   chain.ClientImpl{EthClient: ethClient},
		address:       address,
		relay:         relayContract,
		privateKey:    privateKey,
		senderAddress: senderAddress,
		relaySelector: relaySelectorBytes,
		topic0SPI:     topic0SPI,
		topic0PMR:     topic0PMR,
		gasConfig:     gasConfig,
	}, nil
}

// FetchSigningPolicies fetches signing policies emitted by in SigningPolicyInitialized events from Relay smart contract with timestamps in the interval (from,to].
func (r *relayContractClient) FetchSigningPolicies(ctx context.Context, db finalizerDB, from, to int64) ([]signingPolicyListenerResponse, error) {
	logs, err := db.FetchLogsByAddressAndTopic0(ctx, r.address, r.topic0SPI, from, to)
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

			logs, err := db.FetchLogsByAddressAndTopic0(ctx, r.address, r.topic0SPI, eventRangeStart, now)
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

// SubmitPayloads sends a transaction with input to the Relay contract, retrying
// with the same pre-/post-broadcast and nonce-too-low reconciliation logic as
// the protocol submitter (see SubmitterBase.submit).
func (r *relayContractClient) SubmitPayloads(ctx context.Context, input []byte, dryRun bool, protocolID uint8) {
	if len(input) == 0 {
		return
	}

	// Fetch once (a stable nonce for the reconciliation below), with the full
	// send-retry budget so a transient RPC blip doesn't drop the finalization.
	nonceResult := <-shared.ExecuteWithRetryChan(ctx, func() (uint64, error) {
		return r.chainClient.Nonce(ctx, r.privateKey, 2*time.Second)
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)
	if !nonceResult.Success {
		logger.Warnf("Relaying failed for protocol %d: getting nonce: %v", protocolID, nonceResult.Message)
		return
	}
	nonce := nonceResult.Value

	// Fit boundedSendAttempts attempts into a bounded ctx (grace path) so gas
	// bumps still fire; without a deadline (delayed queue) keep the full timeout.
	perAttempt := chain.DefaultTxTimeout
	if dl, ok := ctx.Deadline(); ok {
		if v := time.Until(dl) / boundedSendAttempts; v < perAttempt {
			perAttempt = max(v, minAttemptTimeout)
		}
	}

	var broadcastHashes []common.Hash
	undetermined := false // an own broadcast's fate was never resolved

	sendResult := <-shared.ExecuteWithRetryAttempts(ctx, func(ri int) (string, error) {
		gasConfig := chain.GasConfigForAttempt(r.gasConfig, ri)

		res := r.chainClient.SendRawTx(ctx, r.privateKey, nonce, r.address, input, gasConfig, perAttempt, dryRun)
		if res.Broadcast {
			broadcastHashes = append(broadcastHashes, res.Hash)
		}

		switch {
		case res.Err == nil:
			return "confirmed " + res.Hash.Hex(), nil
		case chain.MatchesError(res.Err, nonFatalRelayErrors):
			logger.Debugf("Non fatal error sending relay tx for protocol %d: %v", protocolID, res.Err)
			return "non fatal error", nil
		case res.Mined:
			// Mined-reverted is deterministic ("Already relayed" handled above); a resend can't help.
			logger.Warnf("Relaying for protocol %d reverted, not retrying: %v", protocolID, res.Err)
			return "reverted " + res.Hash.Hex(), nil
		case chain.IsNonceTooLow(res.Err):
			h, acc := chain.AnyAccepted(ctx, r.chainClient, r.senderAddress, broadcastHashes, nonFatalRelayErrors, chain.ReconcileLookupTimeout)
			switch acc {
			case chain.Accepted:
				return "reconciled " + h.Hex(), nil
			case chain.Reverted:
				// Our prior broadcast mined but reverted — same terminal handling as
				// a direct mined-revert (deterministic; a resend can't help).
				logger.Warnf("Relaying for protocol %d reverted (reconciled), not retrying: %v", protocolID, res.Err)
				return "reverted " + h.Hex(), nil
			case chain.Undetermined:
				// Outcome unknown: refresh the nonce and resend (a duplicate reverts
				// non-fatally with "Already relayed").
				undetermined = true
				logger.Warnf("Relay protocol %d: nonce %d too low, prior tx status unknown; resending at a refreshed nonce", protocolID, nonce)
				nonce = r.refreshNonce(ctx, nonce)
				return "", res.Err
			default: // NonceConsumed
				nonce = r.refreshNonce(ctx, nonce)
				return "", res.Err
			}
		case res.Broadcast && chain.IsTimeout(res.Err):
			return "", res.Err // keep nonce, retry as replacement
		default:
			// Keep the nonce if a tx is already outstanding at it (retry replaces);
			// only refresh when nothing has been broadcast yet.
			if len(broadcastHashes) == 0 {
				nonce = r.refreshNonce(ctx, nonce)
			}
			return "", res.Err
		}
	}, shared.MaxTxSendRetries, shared.TxRetryInterval)

	if sendResult.Success {
		logger.Infof("Relaying finished for protocol %d with %s", protocolID, sendResult.Value)
	} else if undetermined {
		logger.Warnf("Relay protocol %d: outcome unknown, a broadcast tx may be on chain: %v", protocolID, sendResult.Message)
	} else {
		logger.Warnf("Relaying failed for protocol %d with: %v", protocolID, sendResult.Message)
	}
}

// refreshNonce best-effort re-fetches the sender nonce, keeping current on error.
func (r *relayContractClient) refreshNonce(ctx context.Context, current uint64) uint64 {
	nonce, err := r.chainClient.Nonce(ctx, r.privateKey, time.Second)
	if err != nil {
		logger.Warnf("relay failed to refresh nonce: %v", err)
		return current
	}
	return nonce
}

// relayedKey is the lookup key for ProtocolMessageRelayed events.
//
// It deliberately excludes the seed and msgHash fields of queueItem:
// ProtocolMessageRelayed events identify a finalization uniquely by
// (protocolID, votingRoundID), and using queueItem directly as a map
// key would compare *big.Int by pointer identity — guaranteeing the
// lookup in processDelayedQueue never matches.
type relayedKey struct {
	protocolID    uint8
	votingRoundID uint32
}

// ProtocolMessageRelayed returns a set of (protocolID, votingRoundID)
// pairs that have already been finalized on chain in the given time
// range.
func (r *relayContractClient) ProtocolMessageRelayed(ctx context.Context, db finalizerDB, from time.Time, to time.Time) (map[relayedKey]bool, error) {
	logs, err := db.FetchLogsByAddressAndTopic0(ctx, r.address, r.topic0PMR, from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}

	result := make(map[relayedKey]bool)
	for _, log := range logs {
		data, err := shared.ParseProtocolMessageRelayedEvent(r.relay, log)
		if err != nil {
			return nil, err
		}
		result[relayedKey{
			protocolID:    data.ProtocolId,
			votingRoundID: data.VotingRoundId,
		}] = true
	}
	return result, nil
}
