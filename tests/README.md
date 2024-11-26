# Simulation

## Simulate Flare environment

### Simulate blockchain with deployed contracts

Using the repository `flare-smart-contracts-v2`
found [here](https://gitlab.com/flarenetwork/flare-smart-contracts-v2)
one can deploy all the Fast Updates contracts
together with the whole Flare system and voter repository. Git clone
and navigate to the
repository. Make sure that the `hardhat` version in `package.json` is at least 2.22.16 and run

```bash
yarn install
yarn compile
yarn sim-node # in first terminal
yarn sim-run # in second terminal
```

This will start a Flare system on a local Hardhat node, register 4
data providers and start a simulation of FTSO v2 feed providers.

### Run indexer and database

Using the repository `flare-system-c-chain-indexer` found [here](https://github.com/flare-foundation/flare-system-c-chain-indexer)
run an indexer of the hardhat chain used in the simulation. The process is Dockerized. First
deploy a database by navigating in `tests/docker` of this repository and run

```bash
docker compose up indexer-db
```

and then run

```bash
docker compose up indexer
```

in the same directory.

## Mock a protocol client

Run a server providing payloads for the system client. Navigate to `tests/mock`
and run

```bash
go run server_mock.go
```

### Simulate 3 providers running system clients

#### Run

Finally run

```bash
INSECURE_PRIVATE_KEYS=true go run client/main/client.go --config tests/configs/test_config1.toml
```

and similarly with configs `test_config2.toml` and `test_config3.toml`.
