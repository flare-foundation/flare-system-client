<div align="center">
  <a href="https://flare.network/" target="blank">
    <img src="https://content.flare.network/Flare-2.svg" width="300" alt="Flare Logo" />
  </a>
  <br />
  <a href="CONTRIBUTING.md">Contributing</a>
  ·
  <a href="SECURITY.md">Security</a>
  ·
  <a href="CHANGELOG.md">Changelog</a>
</div>

# Flare Systems Protocol Client

...

[![API Reference](https://pkg.go.dev/badge/github.com/flare-foundation/flare-system-client)](https://pkg.go.dev/github.com/flare-foundation/flare-system-client?tab=doc)

## Configuration

The configuration is read from `toml` file. Some configuration
parameters can also be configured using environment variables. See the list below.

Config file can be specified using the command line parameter `--config`, e.g., `./fsc-client --config config.local.toml`. The default config file name is `config.toml`.

Below is the list of configuration parameters for all clients. Clients that are not enabled can be omitted from the config file.

```toml
[db]
host = "localhost"  # MySql db address, or env variable DB_HOST
port = 3306         # MySql db port, env DB_PORT
database = "flare_fsc"        # database name, env DB_DATABASE
username = "flarefscuser"     # db username, env DB_USERNAME
password = "P.a.s.s.W.O.R.D"  # db password, env DB_PASSWORD
log_queries = false  # Log db queries (for debugging)

[logger]
level = "INFO"      # valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (as in zap logger)
file = "./logs/fsp.log"  # logger file
max_file_size = 10  # max file size before rotating, in MB
console = true      # also log to console

[metrics]
prometheus_address = "localhost:2112"  # expose client metrics to this address (empty value does not expose this endpoint)

[chain]
eth_rpc_url = "http://localhost:9650/ext/C/rpc"  # Ethereum RPC URL
chain_id = 162  # chain id

[contract_addresses]
submission = "0xfae0fd738dabc8a0426f47437322b6d026a9fd95"
systems_manager = "0x22474d350ec2da53d717e30b96e9a2b7628ede5b"
voter_registry = "0xa4bcdf64cdd5451b6ac3743b414124a6299b65ff"
relay = "0x18b9306737eaf6e8fc8e737f488a1ae077b18053"

[identity]
address = "0xd7de703d9bbc4602242d0f3149e5ffcd30eb3adf" # identity account not private key

# reading the private keys from the file is discouraged, it is only allowed when
# INSECURE_PRIVATE_KEYS environment variable is set to true
#
# Otherwise the private keys should be set as environment variables
#  - SYSTEM_CLIENT_SENDER_PRIVATE_KEY
#  - SIGNING_POLICY_PRIVATE_KEY
#  - PROTOCOL_MANAGER_SUBMIT_PRIVATE_KEY
#  - PROTOCOL_MANAGER_SUBMIT_SIGNATURES_PRIVATE_KEY
[credentials]
system_client_sender_private_key_file = "../credentials/sender-private-key.txt" # any account
signing_policy_private_key_file = "../credentials/policy-private-key.txt" # for signing and submitting votes
protocol_manager_submit_private_key_file = "../credentials/submit-private-key.txt"
protocol_manager_submit_signatures_private_key_file = "../credentials/signatures-private-key.txt"

[clients]
enabled_registration = true     # enable/disable voter registration AND new signing policy signing
enabled_uptime_voting = true    # enable/disable uptime vote signing
enabled_reward_signing = false  # enable/disable reward signing
enabled_protocol_voting = true  # enable/disable protocol data submission
enabled_finalizer = true        # enable/disable finalizer client

[protocol.ftso1]
id = 1
api_url = "http://localhost:3000/ftso1"
type = 0 # payload type: currently available 0 and 1
# To specify an API key for this endpoint set it via PROTOCOL_X_API_KEY_1 env var

[protocol.ftso2]
id = 2
api_url = "http://localhost:3000/ftso2"
type = 0
# To specify an API key for this endpoint set it via PROTOCOL_X_API_KEY_2 env var

[submit1]
enabled = true             # (optional) set to false to disable a specific submitter, default: true
start_offset = "85"        # start fetching data and submitting txs after this offset from the start of the epoch
tx_submit_retries = 1      # (optional) number of retries for submitting txs, default: 1
tx_submit_timeout = "10s"  # (optional) timeout for waiting tx to be mined, default: 10s
data_fetch_retries = 1     # (optional) number of retries for fetching data from the API, default: 1
data_fetch_timeout = "5s"  # (optional) timeout for fetching data from the API, default: 5s

[submit2]
enabled = true
start_offset = "30s"       # start fetching data and submitting txs after this offset from the start of the NEXT epoch
tx_submit_retries = 1
tx_submit_timeout = "10s"
data_fetch_retries = 1
data_fetch_timeout = "5s"
```

Signature submission is set differently than submit and submit2. See
`func (s *SignatureSubmitter) RunEpoch(currentEpoch int64)` in `submitter.go`.

```toml
[submit_signatures]
enabled = true
start_offset = "45s"      # start fetching data and submitting txs after this offset from the start of the NEXT epoch
deadline = "60s"          # submit transaction until deadline
tx_submit_retries = 1
data_fetch_retries = 1     # number of retries for fetching data from the API, timeout is 1 second
data_fetch_timeout = "2s"
cycle_duration = "2s"
max_cycles = 3             # max number of rounds to fetch data and submit signatures

[finalizer]
starting_reward_epoch = 0
starting_voting_round = 1005
start_offset = "500s"            # how far in the past we start fetching reward epochs from the indexer at the start of the finalizer client default is 7 days
grace_period_end_offset = "65s"  # Offset from the start of the voting round
```

Type 0 and type 2 are supported for transactions to Submission and Relay contracts

```toml
[gas_submit] # applies to all submit1, submit2 and submitSignatures transactions.
tx_type = 2                            # 0 for legacy and 2 for eip-1559 transaction
gas_limit = 0                          # (optional) gas limit for transaction. Defaults to 0, which will use gas limit estimates.
# type 0 settings // Note: only one of gas_price_multiplier and gas_price_fixed can be set.
gas_price_multiplier = 0               # (optional for type 0 tx) sets the gas price to be a multiplier of the estimated gas price. Defaults to 0, which will simply use the estimate, OR a fixed gas price if gas_price_fixed is set (!= 0).
gas_price_fixed = 0                    # (optional for type 0 tx) sets a fixed gas price for the transaction. Defaults to 0, which will use an estimate OR a multiplier of the estimate if gas_price_multiplier is set (!= 0).
# type 2 settings
max_priority_fee_multiplier = 2 # (optional for type 2 tx) sets the max priority fee per gas to be a multiple of the estimated base fee. Defaults to 2.
maximal_max_priority_fee_per_gas = 5000_000_000_000 # (optional for type 2 tx) maximal max priority fee per gas. Defaults to 5000 Gwei.
minimal_max_priority_fee =100_000_000_000 # (optional for type 2 tx) minimal max priority fee per gas. Defaults to 100 Gwei.
base_fee_multiplier = 4 # (optional for type 2 tx) sets the base fee to be a multiple of the estimated base fee. Defaults to 4.



[gas_relay] # applies to finalization transaction
tx_type = 2
max_priority_fee_multiplier = 2
maximal_max_priority_fee = 5000_000_000_000 # 5000 Gwei
minimal_max_priority_fee = 100_000_000_000 # 100 Gwei
base_fee_multiplier = 4


[gas_register] # applies to all voter registration transaction
tx_type = 2
max_priority_fee_multiplier = 2
maximal_max_priority_fee = 5000_000_000_000 # 5000 Gwei
minimal_max_priority_fee = 100_000_000_000 # 100 Gwei
base_fee_multiplier = 4


[gas_systems_manager] # applies to transactions to FlareSystemsManager contract
tx_type = 2
max_priority_fee_multiplier = 2
maximal_max_priority_fee = 5000_000_000_000 # 5000 Gwei
minimal_max_priority_fee = 100_000_000_000 # 100 Gwei
base_fee_multiplier = 4
```

```toml
[rewards] # reward signing configuration - clients.enabled_reward_signing must be set to true
# URL prefix for retrieving reward distribution data.

# A full URL will be constructed by appending the epoch id and expected file name: <prefix>/<epochId>/reward-distribution-data.json

#

# For example, if reward data for an epoch can be retrieved at https://raw.githubusercontent.com/flare-foundation/fsp-rewards/refs/heads/main/songbird/240/reward-distribution-data.json,

# then the url_prefix should be set to "https://raw.githubusercontent.com/flare-foundation/fsp-rewards/refs/heads/main/songbird"

url_prefix = ""
min_reward = 0 # minimum acceptable claim amount in wei for the identity address of this provider, default 0.
max_reward = 0 # (optional) maximum acceptable claim amount in wei for the identity address of this provider. If 0 or not set, no maximum is enforced.
retries = 8 # (optional) number of retries for fetching and signing reward data, default: 8.
retry_interval = "6h" # (optional) interval between retries, default: 6 hours.

```
