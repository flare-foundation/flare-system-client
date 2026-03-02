# Changelog

## [v.1.0.10](https://github.com/flare-foundation/flare-system-client/tree/v1.0.10) - 2026-3-?

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
