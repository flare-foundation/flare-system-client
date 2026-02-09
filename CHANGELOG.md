# Changelog

## [Unreleased]

### Changed

- For type 2 transactions, max priority fee per gas is now computed as configurableMultiplies multiplied with the estimation of the baseFee. The default multiplier is set to 2.
- More aggressive gas settings for signingPolicy signing. Raised maximal number of transaction retries.
