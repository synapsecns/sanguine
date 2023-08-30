# Guard

The Guard is an agent that detects and reports fraud, and relays origin states to the Summit.

## Usage

Navigate to `sanguine/agents/main` and run the following command to start the agents CLI:

```bash
$ go run main.go
```
Then the command line will be exposed. The CLI takes in three arguments, in addition to the agent specifier:
1. Specify the agent to run: `guard-run`
2. `--config </Path/to/config.yaml>`: This argument is required. It is the path to the config file
3. `--metrics-port <port>`: Port to expose metrics on
4. `--debug`: Enable debug tracing on the omniRPC client

## Configuration

The Guard requires a config file to run. The config file is a yaml file that contains the following fields:

```yaml
# The `db_config` field specifies the database type and the source (either a path or a connection string).
db_config:
  # Must be mysql or sqlite.
  type: mysql
  # Source is either a path (for sqlite) or a connection string (for mysql).
  source: "root:password@tcp(agents-mysql:3306)/guard?parseTime=true"

# The base omnirpc url which each chain's collection of RPC's will be proxied through.
base_omnirpc_url: http://agents-omnirpc

# For each chain (domain), specify the necessary contracts.
# Remotes need: origin_address
# Summit needs: origin_address, summit_address, inbox_address
# TODO: Update this for fraud.
domains:
  domain_client1:
    domain_id: 11155111
    type: EVM
    required_confirmations: 0
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
  domain_client2:
    domain_id: 901
    type: EVM
    required_confirmations: 0
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    summit_address: 0xFf9F696EE2A213FA0519673f6503F2933A574093
    inbox_address: 0x3E56136B612427AcCb501752e08c0872518Ae81E
  domain_client3:
    domain_id: 43113
    type: EVM
    required_confirmations: 0
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861

# Specify the summit domain id
summit_domain_id: 901
# Guards have a `domain_id` of 0.
domain_id: 0

# The `unbonded_signer` field specifies the path to the file containing the private key of the signer
unbonded_signer:
  type: "File"
  file: "/config/guard_signer.txt"

# The `bonded_signer` is the account that will post a bond to the Summit contract. Specify its path to
# the file containing the private key of the signer
bonded_signer:
  type: "File"
  file: "/config/guard_signer.txt"

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
