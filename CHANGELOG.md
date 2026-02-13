# Changelog

## [UNRELEASED]

### Changed

- For type 2 gas configs, both MaximalMaxPriorityFe and MinimalMaxPriorityFee are increased by 11% on each retry to ensure valid replacement transactions

## [v1.0.6] - 2026-2-13

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
