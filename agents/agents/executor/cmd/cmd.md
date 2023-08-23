# Executor

The Executor is an agent that verifies messages and executes them on the destination chain.

## Usage

Navigate to `sanguine/agents/main` and run the following command to start the agents CLI:

```bash
$ go run main.go
```
Then the command line will be exposed. The CLI takes in three arguments, in addition to the agent specifier:
1. Specify the agent to run: `executor-run`
2. `--config </Path/to/config.yaml>`: This argument is required. It is the path to the config file
3. `--metrics-port <port>`: Port to expose metrics on
4. `--debug`: Enable debug tracing on the omniRPC client

## Configuration

The Executor requires a config file to run. The config file is a yaml file that contains the following fields:

```yaml
# The `db_config` field specifies the database type and the source (either a path or a connection string).
db_config:
  # Must be mysql or sqlite.
  type: mysql
  # Source is either a path (for sqlite) or a connection string (for mysql).
  source: "root:password@tcp(agents-mysql:3306)/executor?parseTime=true"

# The base omnirpc url which each chain's collection of RPC's will be proxied through.
base_omnirpc_url: http://agents-omnirpc
execute_interval: 5
summit_chain_id: 901
summit_address: 0xFf9F696EE2A213FA0519673f6503F2933A574093
inbox_address: 0x3E56136B612427AcCb501752e08c0872518Ae81E

# For each chain (including the Summit chain), specify the origin and destination addresses.
# For remote chains, also specify the light inbox address.
chains:
  - chain_id: 11155111
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    destination_address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
    light_inbox_address: 0x22EaC155C605Ffaf1002A2733206BCC0A52EB8C8
  - chain_id: 901
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    destination_address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
  - chain_id: 43113
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    destination_address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
    light_inbox_address: 0x22EaC155C605Ffaf1002A2733206BCC0A52EB8C8

# The `unbonded_signer` field specifies the path to the file containing the private key of the signer
unbonded_signer:
  type: "File"
  file: "/path/executor_signer.txt"

# The `submitter_config` field specifies how the submitter should submit messages to the chains.
submitter_config:
  chains:
    11155111:
      supports_eip_1559: true
      gas_estimate: 7500000
    901:
      gas_bump_percentage: 40
      is_l2: true
      gas_estimate: 7500000
    43113:
      gas_estimate: 7500000
```

The last component that this config needs is a `Scribe` config. There are two types of Scribes that will have different configs:
1. Embedded Scribe: This is a `Scribe` that is embedded in the Executor. When running an Executor locally, this is the simplest option. The config for this is as follows:
```yaml
type: embedded

# The `embedded_db_config` is the database specification for the embedded Scribe.
embedded_db_config:
  type: sqlite
  source: "/path/to/scribe.db"

# The `embedded_scibe_config` is a whole configuration for a Scribe.
embedded_scribe_config:
  # Every chain and contract that the Executor is listening for logs from must be specified here.
  chains:
    - chain_id: 11155111
      required_confirmations: 0
      contract_sub_chunk_size: 512
      contract_chunk_size: 512
      store_concurrency_threshold: 100000
      contracts:
        - address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
          start_block: 3924722
        - address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
          start_block: 3924722
        - address: 0x22EaC155C605Ffaf1002A2733206BCC0A52EB8C8
          start_block: 3924722
    - chain_id: 901
      required_confirmations: 0
      contract_sub_chunk_size: 512
      contract_chunk_size: 512
      store_concurrency: 1
      store_concurrency_threshold: 100000
      contracts:
        - address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
          start_block: 3143
        - address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
          start_block: 3143
        - address: 0x3E56136B612427AcCb501752e08c0872518Ae81E
          start_block: 3143
        - address: 0xFf9F696EE2A213FA0519673f6503F2933A574093
          start_block: 3143
    - chain_id: 43113
      required_confirmations: 0
      contract_sub_chunk_size: 512
      contract_chunk_size: 512
      store_concurrency_threshold: 100000
      contracts:
        - address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
          start_block: 24307130
        - address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
          start_block: 24307127
        - address: 0x22EaC155C605Ffaf1002A2733206BCC0A52EB8C8
          start_block: 24307126
  # The `rpc_url` is an omnirpc endpoint, similar to the `base_omnirpc_url` in the Executor config.
  # However, notice the inclusion of `/confirmations`. This could be the same base url as above with an added `/confirmations` path.
  rpc_url: http://scribe-omnirpc/confirmations
  refresh_rate: 0
```

2. Remote Scribe: This is a `Scribe` that is running separately, but accessible to the Executor via a gRPC connection. This is the recommended option for running an Executor in production. The config for this is as follows:
```yaml
scribe_config:
  type: remote
  port: 80
  url: executor-scribe.url
```
Note that the remote `Scribe` instance also needs to be listening for all chains and contracts that the Executor will need, but the configuration for the remote `Scribe` is used when deploying the remote `Scribe`.

Append whichever `Scribe` config you choose to the end of the Executor config file, and the Executor should be ready to run.
