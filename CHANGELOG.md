# Changelog

## [Unreleased]

### Added

- Support for the Relay contract cutover. A per-chain table (`client/shared/relay.go`) names the new Relay address and the first reward epoch it serves; up to that epoch everything is read from and sent to the configured Relay exactly as before, and from it on to the new one. The voting round the switch takes effect on is not configured — a reward epoch's start can be delayed, so it is only fixed once that epoch's signing policy exists, and it is learned at runtime from the policy's own `startVotingRoundId` (from the `SigningPolicyInitialized` event where a client listens for it, and read back from the systems manager otherwise, which stores it in the same transaction that emits the event). Both Relays are read for `SigningPolicyInitialized` and `ProtocolMessageRelayed` across the switch, since a reward epoch's policy is only emitted by the Relay that holds it. A finalization goes to the Relay storing its signing policy's hash, so a late pre-cutover round is still relayed to the old contract, and the already-relayed check is made per target: a round relayed only on the old Relay is not relayed for readers of the new one. The table ships empty — a chain without a complete entry keeps the current behaviour, and a half-filled entry aborts startup rather than silently disabling the switch.
- The random number and its Merkle proof are appended to a finalization from the breaking reward epoch on, as the new Relay's `relay()` requires for the random protocol: `randomNumber(32)` followed by the proof nodes, after the signatures. The old Relay derived the random from the signed Merkle root instead, so nothing was appended and no source was needed. The data comes from the random protocol's own data provider — only it can produce the number and its proof — which serves it with its `submitSignatures` message, as a new `finalizationData` field of that response: a `0x` hex string of the trailer bytes themselves, the value followed by the proof nodes. The submitter forwards the field to the finalizer with the message it already hands over, so the duty needs no new configuration and no new endpoint, and the data exists exactly when the message does — the proof is a leaf of the tree the message's root commits to. The finalizer knows which protocol must carry it (the one whose id the Relay's `stateData` reports) and from which round on (the breaking epoch's first), and checks it the moment the message arrives — a whole number of 32-byte words, folded against the message's own Merkle root with the voting round and the secure flag taken from the signed message, so only the value and the proof are trusted — logging an error for a round whose provider omitted or broke it. It is checked once more against the finalized message before the send, and a round whose proof is missing or does not fold is not sent at all, since the transaction could only revert. A finalization whose threshold is reached from peers before the local provider has answered waits a bounded time for the message, and the delayed queue retries later. Nothing is expected until a Relay switch is scheduled for the chain; from then on the finalizer refuses to start unless that protocol is configured, since the data can only arrive through it.
- The reward epoch client's signing-policy listener selects the policy it is waiting for by reward epoch, falling back to the highest one present when a delayed epoch leaves the anticipated index running ahead of the chain. It previously took the last log in the window by position, which across the Relay cutover — where the merged logs of two contracts have no meaningful order by emitter — could pick the wrong contract's policy.

### Changed

- Protocol-message signatures bind the source chain from the Relay cutover on: the digest becomes `keccak256(chainID ‖ message)` instead of `keccak256(message)`, so a signature minted for another network is rejected even under a fully overlapping voter set. Signing (submitter, gated on the voting round) and signer recovery (finalizer, gated on the collection's signing-policy reward epoch) share one digest function and one cutover instance, so the two gates cannot drift apart. While the breaking epoch's start round is still unknown the round cannot be past it, so the pre-switch digest is used.
- The signing-policy hash signed in `signNewSigningPolicy` follows the cutover: the old Relay's chained keccak fold before the breaking reward epoch, one keccak over the chain id and the raw policy bytes from it on. Which Relay the FlareSystemsManager points at is part of that decision rather than a substitute for it — the new Relay delegates epochs below the breaking one to the old Relay, and until governance repoints the manager the old Relay answers for every epoch, including the breaking one, whose policy is signed before a repoint is possible. The other scheme stays a checked fallback that logs a warning naming the epoch and the Relay, so a cutover table disagreeing with the chain costs a warning instead of the signature; the manager accumulates signing weight per epoch rather than per hash, so a mixed-scheme fleet still reaches one threshold. Only a hash derived from the policy bytes carried by the `SigningPolicyInitialized` event is ever signed.
- `APIUrl`, `XApiKey` and `ApiKey` identifiers follow Go's initialism convention (`APIURL`, `XAPIKey`, `APIKey`); TOML keys and environment variable names are unchanged. `SubProtocol.APIURL` is renamed `BaseURL`, which is what it holds — the provider's base, with the per-request path appended.

### Fixed

- Submitted signatures are canonicalized before they can count toward the finalization threshold. The new Relay rejects a non-canonical ECDSA record — `v` outside {27, 28}, or `s` above half the group order (EIP-2) — by reverting the *whole* `relay()` call, while the previously deployed Relay checked neither, so one peer submitting the equally valid high-`s` encoding of its own signature would have made every finalization of that round revert on every finalizer, unrecoverably. A high-`s` signature is normalized to `(r, n-s, v^1)`, which recovers the same signer, rather than dropped: dropping it would lose that voter's weight and could put the round below threshold, which is the same outcome. A signature that cannot be normalized is rejected before its weight is credited, so every signature reaching the finalization calldata satisfies the checks. Normalization logs at warning, naming the sender, since it signals a signer that does not normalize `s`.

## [v1.1.2](https://github.com/flare-foundation/flare-system-client/tree/v1.1.2) - 2026-8-18

### Added

- Startup warning when an enabled client's gas config sets `gas_price_fixed` (a backwards-compatibility option pinning the whole fee, so retries cannot replace a stuck transaction) or `base_fee_per_gas_cap` (pinning the base-fee component of the cap, which is not bumped on retry).
- Startup validation rejecting a negative `gas_limit`, which previously wrapped via `uint64()` into an unusable ~1.8e19 gas limit at transaction-build time, permanently rejecting every relay and voter-registration transaction.
- Startup verification of the configured `chain_id` against the node's `eth_chainId`: a mismatch (which would make every send fail with "invalid sender") aborts startup with a fatal log; an unreachable node only warns, so a node outage does not block a restart.
- Startup validation for an enabled finalizer: `grace_period_end_offset` must be set (it has no default, and unset silently disabled grace gating, relaying every round immediately) and `voter_threshold_bips` must be positive (0 silently made the node never-selected for grace finalization).
- Startup validation rejecting a negative `gas_price_fixed` or `base_fee_per_gas_cap`: both were silently ignored at send time, while the gas-override startup warning still described the pin as active.
- Startup verification that the configured `voter_registry` and `voter_preregistry` addresses hold contract code, logging a warning when they do not. `eth_estimateGas` succeeds against an address with no code, so a registry address pasted from another network produces transactions that mine successfully and register nothing. The check warns rather than aborts, since a configured address may legally predate its deployment; either address is skipped when unset, which config validation already permits when the matching client is disabled.

### Changed

- Every line of both send loops now carries the voting round, `attempt N/M`, the nonce and, on the terminal line, the elapsed time; the round is threaded into the submitter's and the relay's send functions, which previously logged nothing that could be tied to a round. A nonce refresh logs its result, including when it comes back unchanged — the nonce is read from the latest mined block, so a resend at an already-rejected nonce looked like a successful refresh.
- Gas is logged as key-value text in gwei under the TOML key names, showing only the fields the configured transaction type uses (was a raw struct dump in wei), the fees actually signed on the broadcast line, and at debug the observed base fee plus any priority-fee clamp — a tip pinned below market previously timed out every attempt with nothing explaining why. The relay loop, which bumped gas per attempt and logged none of it, now logs it too.
- The relay's silent outcomes (reconciled-accepted, nonce consumed by another transaction, and every pre-broadcast failure) now log like the submitter's; a reconciled revert prints the reverted transaction hash instead of the unrelated nonce-too-low error, a submitter's own mined-but-reverted transaction is no longer reported as another transaction consuming the nonce, and a not-found receipt during reconciliation — previously the only unlogged branch, and the one that drives a resend — is logged at debug.
- Relay nonce-fetch and terminal send failures log at error, matching the submitter. A send that exhausts its budget with a transaction still outstanding reports "outcome unknown" with the hashes on both paths, now including after post-broadcast timeouts, which were reported as flat failures.
- Relay nonce-too-low reconciliation now bounds each receipt/revert lookup to 5s (was the full 60s tx timeout), so a hung RPC endpoint cannot stall a retry cycle for minutes.
- A send that exhausts retries with an own broadcast still unresolved logs "outcome unknown, a broadcast tx may be on chain" at warning level instead of an unqualified error, distinguishing a possibly-successful send from a confirmed failure.
- Example config: `submit1` start offset moved from 75s to 65s and a note added, keeping same-key submit1/submit2 fires well apart around the round so overlapping sends cannot collide on the nonce.
- Expected payload rejections in the finalizer (bad signature, unregistered signer, duplicate signature, round below the stored window) now log at debug instead of error; only unexpected failures remain at error, so an error from submission processing again signals a real problem.
- Relay and submit transactions now sign with the `chain_id` from config instead of fetching the network id from the node on every send, so a transient `net_version` failure can no longer abort a send; `chain_id` is validated as set at startup.
- The gas-limit estimate and fee reads of a transaction build run concurrently: a send's pre-broadcast phase costs at most one RPC round-trip timeout instead of three.

### Fixed

- Nonce-too-low reconciliation now decodes the revert reason from the JSON-RPC error geth/coreth return for a reverting `eth_call`; previously the reason was never recovered, so a relay tx that mined reverting with the non-fatal "Already relayed" was treated as undetermined and retried instead of recognized as already done.
- A relay transaction that mines but reverts for a non-fatal reason is no longer retried (whether observed directly or via reconciliation): the revert is deterministic on the signed payload, so a resend only re-mines and wastes gas.
- Nonce-too-low reconciliation treats a not-found receipt as undetermined rather than conclusive, and an unresolved undetermined now refreshes the nonce and resends instead of retrying the same nonce until the budget is exhausted: a submission whose nonce was consumed by another transaction is resent instead of dropped, at the cost of a rare benign duplicate (submits are idempotent; a duplicate relay reverts non-fatally with "Already relayed").
- A stuck relay send no longer blocks the finalization queue processor for its full retry budget (~11 minutes), delaying other rounds' grace-period finalizations past their window: each queued item's send is bounded to 50 seconds, with per-attempt timeouts sized (inter-attempt backoff included) so gas-bumped replacements still fire within the bound, the nonce prefetch capped so a flaky fetch cannot starve the send attempts, and the item then falls back to the delayed queue where already-relayed rounds are skipped; a fallback target already in the past — previously every post-grace "send now" item silently lost its retry this way — is rescheduled a few seconds ahead, and the delayed queue logs any dropped past-time entry instead of discarding it silently.
- Receipt waits poll every 400ms (relaxing to 1s after 10s of waiting) instead of relying on `bind.WaitMined`, whose fixed 1s interval added up to a second of latency per attempt against 1–2s Flare blocks — enough to push a bounded send attempt past its slice and turn a mined transaction into a retry.
- Restored the relay nonce fetch's retry budget so a transient RPC failure no longer drops a finalization after only a few hundred milliseconds.
- Transaction send-retry helpers now abort promptly on context cancellation instead of sleeping through the remaining retries, unblocking graceful shutdown (previously up to ~50s for the finalizer, longer for epoch paths).
- A transient indexer-DB error during the delayed queue's already-relayed check no longer drops the whole batch of pending finalizations (the items are consumed from the queue before processing and were never retried): the check is skipped instead, and the dry-run send catches already-relayed rounds pre-broadcast.

### Removed

- Support for the previous generation of the `voterRegistry` and `voterPreRegistry` contracts. Flare, Songbird and Coston are all past their breaking reward epochs (417, 417 and 5451), and the contract registry on each resolves both names to the new deployments, so the per-chain breaking-epoch switch and the old contract addresses no longer had a live case. Registration and pre-registration now always target the configured address.
- The old registration message form. The signed message is now always `keccak256(abi.encode(chainID, nextRewardEpochID, address))`; the two-argument form without the chain ID is gone. This also changes behaviour for any chain outside Flare, Songbird, Coston and Coston2, which previously fell back to the old form regardless of epoch — `chain_id` is now load-bearing for signature validity everywhere.
- The hardcoded per-chain registry address table, along with the check that the configured address matched the configured `chain_id`. That check only covered three addresses and went silently inert whenever the contracts were redeployed, which is what this release does. The code-presence check above covers every address on every chain instead, but only warns, so a wrong-network address no longer stops startup.

## [v1.1.1](https://github.com/flare-foundation/flare-system-client/tree/v1.1.1) - 2026-7-15

### Added

- Per-submitter `enabled` option (default true) reintroduced in `[submit1]`, `[submit2]`, and `[submit_signatures]` sections, allowing individual submitter opt-out via `enabled = false`.
- Startup warnings for penalised opt-out combinations: submit1 enabled without submit2 (FTSO penalises a commit with no reveal), and submit2 enabled without submit_signatures (FDC penalises a reveal with no signatures).
- Startup validation rejecting all three submitters disabled while `enabled_protocol_voting = true`, and validation that `enabled_finalizer = true` requires `submit_signatures` to be enabled.

### Changed

- **Config:** leftover `enabled` keys in `[submit1]`/`[submit2]`/`[submit_signatures]` — documented as ignored in v1.1.0 — take effect again: a stale `enabled = false` now opts that submitter out of every round. Check configs for stale keys when upgrading.

## [v1.1.0](https://github.com/flare-foundation/flare-system-client/tree/v1.1.0) - 2026-7-14

### Added

- CI: test coverage reporting and full pipelines on merge requests.
- Tests covering payload extraction, signature transforms, finalization storage cleanup and concurrent access, gas config validation, reward data bounds, protocol client shutdown, and protocol client HTTP response parsing.
- Startup validation of the submitter config sections (`submit1`, `submit2`, `submit_signatures`): rejects negative start offsets, non-positive submit/data-fetch timeouts, retry counts below one, a `submit_signatures` deadline at or before its start offset, negative `max_cycles`/`cycle_duration`, and a `submit_signatures` start offset scheduled before the `submit2` reveal.
- Startup validation that the finalizer requires protocol voting: `enabled_finalizer` needs `enabled_protocol_voting`, otherwise the finalizer would never receive submitted messages.
- Startup validation of type-2 gas multipliers: rejects non-finite (`inf`/`nan`) or non-positive `max_priority_fee_multiplier` / `base_fee_multiplier`, requires `base_fee_multiplier` to be at least 1 (unless `base_fee_per_gas_cap` is set) so the fee cap covers the base fee, and rejects a non-finite or below-1 `gas_price_multiplier` — surfacing bad values at startup instead of panicking at transaction time.
- Smooth transition for Flare and Songbird (previously Coston-only) for the voterRegistry and voterPreRegistry smart contracts: switches to the new contract addresses at a per-chain breaking epoch, and the voter registration message hash now also includes the chain ID.

### Changed

- **Config (breaking):** removed the per-submitter `enabled` option from `[submit1]`, `[submit2]`, and `[submit_signatures]`. All three submitters now always run when protocol voting is enabled; use `enabled_protocol_voting` to turn submission on or off. Any leftover `enabled` key in these sections is ignored.

- Type-2 gas config: `max_priority_fee_multiplier` and `base_fee_multiplier` now accept fractional (float) values (e.g. `1.5`) instead of only whole numbers; quoted-string values (`"2"`) from the previous `big.Int` format are still parsed for backward compatibility.
- Finalization storage returns the live signatures collection guarded by a per-collection mutex instead of a deep copy on every read.
- Migrated error handling from `github.com/pkg/errors` to the standard library; cleaned up error and log messages and fixed logger format misuse.
- Bumped go-ethereum to v1.17.3, go-flare-common to v1.2.1, and Go to 1.26.4.

### Fixed

- Type-2 transactions built via `SetGas` (voter registration, signing-policy, systems-manager) now clamp the priority fee to `[minimal_max_priority_fee, maximal_max_priority_fee]`, matching the submit/finalize path. Both paths now default the gas config identically, so an unset priority-fee cap can no longer be dereferenced.
- Panic in `FromSignedPayload` when a submitSignatures transaction contained a zero-length payload; the empty slice is now rejected with an error and skipped by the caller.
- Integer overflow in `ExtractPayloads` length handling that could bypass the bounds check on crafted submitSignatures calldata.
- Panics on malformed input: signature transforms and payload extraction now validate slice lengths instead of assuming 65-byte signatures and a 4-byte function selector.
- Delayed finalization queue compared already-relayed rounds by seed pointer instead of value, so finalizations could be re-sent for rounds that were already relayed.
- Unauthenticated submitSignatures payloads are now capped to one buffered payload and one signature-collection allocation per sender per round and protocol, bounding memory growth and the ECDSA-recovery burst from crafted transactions (DOS-01).
- Panic in `unpackError` on short revert data.
- Data race on shared `TransactOpts` in the epoch client.
- `WaitGroup` misuse in protocol client submitter scheduling.
- Panics on nil `big.Int` values during reward data verification.
- Panic in gas config validation when `gas_price_fixed` was unset for type 0 transactions.
- Round leak in finalization storage: `RemoveRoundsBefore` left one stale round stored forever and rejected new payloads for it.
- Protocol client registration checks (`isRegistered`/`waitUntilRegistered`) now query the same old-or-new voterRegistry as the send path, based on the reward epoch relative to the breaking epoch, instead of always querying the new registry.

## [v.1.0.12](https://github.com/flare-foundation/flare-system-client/tree/v1.0.12) - 2026-4-17

### Added

- Smooth transition for Coston for voterRegistry and voterPreRegistry smart contracts

### Fixed

- Issues pointed out by github [issue #4](https://github.com/flare-foundation/flare-system-client/issues/4).
  Improved context handling, more decoupled submitter and finalizer client, and immediate client shutdown if an unexpected error in finalizer happens.

## [v.1.0.11](https://github.com/flare-foundation/flare-system-client/tree/v1.0.11) - 2026-3-25

### Added

- Changed registry and preregistry smart contracts with updated message for signing.

## [v.1.0.10](https://github.com/flare-foundation/flare-system-client/tree/v1.0.10) - 2026-3-2

### Added

- Automated releases on GitHub.

### Changed

- Config examples and template fully moved from README to cong.toml.

### Fixed

- Minor issues found by AI review.

### Removed

- Changes needed for Relay contract address update from v1.0.8.

## [v1.0.9](https://github.com/flare-foundation/flare-system-client/tree/v1.0.9) - 2026-2-20

### Fixed

- Copying of big.Int in Gas configs.

## [v1.0.8](https://github.com/flare-foundation/flare-system-client/tree/v1.0.8) - 2026-2-19

### Changed

- Addressed change of Relay contract address on all chains.

## [v1.0.7](https://github.com/flare-foundation/flare-system-client/tree/v1.0.7) - 2026-2-17

### Changed

- For type 2 gas configs, both MaximalMaxPriorityFe and MinimalMaxPriorityFee are increased by 11% on each retry to ensure valid replacement transactions.

### Fixed

- Nil pointer for gas config for signing policy signing.

## [v1.0.6](https://github.com/flare-foundation/flare-system-client/tree/v1.0.6) - 2026-2-13

### Added

- SECURITY.md
- CHANGELOG.md
- CONTRIBUTING.md
- CODEOWNERS

### Changed

- go version update
- Config for gas for raw transactions
  - Type 2 is default
  - Removed:
    - MaxPriorityFeePerGas (max_priority_fee_per_gas)
  - Added:
    - MaxPriorityFeeMultiplier (max_priority_fee_multiplier) withe default 2
    - MaximalMaxPriorityFee (maximal_max_priority_fee) with default 5000 Gwei (5_000_000_000_000)
    - MinimalMaxPriorityFee (minimal_max_priority_fee) with default 100 Gwei (100_000_000_000)
  - For type 2, the MaxPriorityFee is set a product of MaxPriorityFeeMultiplier and estimation of the baseFee caped with MaximalMaxPriorityFee and MinimalMaxPriorityFee, respectively.
- More aggressive gas settings for signingPolicy signing. Raised maximal number of transaction retries.
- Dependency updates
- README.md updates
- Small refactors
