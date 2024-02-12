# Flare Top level protocol client

...

## Configuration

The configuration is read from `toml` file. Some configuration
parameters can also be configured using environment variables. See the list below.

Config file can be specified using the command line parameter `--config`, e.g., `./tlc-client --config config.local.toml`. The default config file name is `config.toml`.

Below is the list of configuration parameters for all clients. Clients that are not enabled can be omitted from the config file.

```toml
[db]
host = "localhost"  # MySql db address, or env variable DB_HOST
port = 3306         # MySql db port, env DB_PORT
database = "flare_tlc"        # database name, env DB_DATABASE
username = "flaretlcuser"     # db username, env DB_USERNAME
password = "P.a.s.s.W.O.R.D"  # db password, env DB_PASSWORD
log_queries = false  # Log db queries (for debugging)

[logger]
level = "INFO"      # valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (as in zap logger)
file = "./logs/flare-tlc.log"  # logger file
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

[credentials]
system_client_sender_private_key_file = "../credentials/sender-private-key.txt" # any account
signing_policy_private_key_file = "../credentials/policy-private-key.txt" # for signing and submitting votes
protocol_manager_submit_private_key_file = "../credentials/submit-private-key.txt"
protocol_manager_submit_signatures_private_key_file = "../credentials/signatures-private-key.txt"

[clients]
enabled_registration = true       # enable/disable registration - send RegisterVoter and SignNewSigningPolicy txs
enabled_protocol_voting = false
enabled_finalizer = false

[protocol.ftso1]
id = 1
api_endpoint = "http://localhost:3000/ftso1"

[protocol.ftso2]
id = 2
api_endpoint = "http://localhost:3000/ftso2"

[submit1]
start_offset = "5s"    # start fetching data and submitting txs after this offset from the start of the epoch
tx_submit_retries = 1  # number of retries for submitting txs

[submit2]
start_offset = "15s"
tx_submit_retries = 1

[submit_signatures]
start_offset = "10s"    # offset from the start of the epoch
tx_submit_retries = 3
data_fetch_retries = 5  # number of retries for fetching data from the API, timeout is 1 second
max_rounds = 3          # max number of rounds to fetch data and submit signatures

[finalizer]
starting_reward_epoch = 0
starting_voting_round = 1005
start_offset = "500s" # how far in the past we start fetching reward epochs from the indexer at the start of the finalizer client default is 7 days
grace_period_end_offset = "40s"  # Offset from the start of the voting round
```
