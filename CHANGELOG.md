# Changelog

## [Unreleased]

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
