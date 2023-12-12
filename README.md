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
node_url = "http://localhost:9650/"  # node client address
address_hrp = "localflare"  # HRP (human readable part) of chain -- used to properly encode/decode addresses
chain_id = 162  # chain id
eth_rpc_url = "http://localhost:9650/ext/C/rpc"  # Ethereum RPC URL
api_key = ""    # API key (in case the node is protected by API key), adds ?x-apikey=... to all requests if not empty
private_key_file = "../credentials/pk.txt"  # file containing the private key of an account (for voting and mirroring clients), in hex


[contract_addresses]
voting = "0xf956df3800379fdFA31D0A45FDD5001D02F4109c"       # voting contract address
mirroring = "0xE64Df6a7e4f4c277C5299f0FE12D7BbB8A207175"    # mirror contract address
```

## Deployment configuration

...

## Running tests

...
