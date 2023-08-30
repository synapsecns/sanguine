# Notary

The Notary is an agent that detects and reports fraud, and relays origin states to the Summit.

## Usage

Navigate to `sanguine/agents/main` and run the following command to start the agents CLI:

```bash
$ go run main.go
```
Then the command line will be exposed. The CLI takes in three arguments, in addition to the agent specifier:
1. Specify the agent to run: `notary-run`
2. `--config </Path/to/config.yaml>`: This argument is required. It is the path to the config file
3. `--metrics-port <port>`: Port to expose metrics on
4. `--debug`: Enable debug tracing on the omniRPC client

## Configuration

The Notary requires a config file to run. The config file is a yaml file that contains the following fields:

```yaml
# The `db_config` field specifies the database type and the source (either a path or a connection string).
db_config:
  # Must be mysql or sqlite.
  type: mysql
  # Source is either a path (for sqlite) or a connection string (for mysql).
  source: "root:password@tcp(agents-mysql:3306)/notary?parseTime=true"

# The base omnirpc url which each chain's collection of RPC's will be proxied through.
base_omnirpc_url: http://agents-omnirpc

# For each chain (domain), specify the necessary contracts.
# Remotes need: origin_address, destination_address, light_inbox_address, light_manager_address
# Summit needs: origin_address, destination_address, summit_address, inbox_address, bonding_manager_address
domains:
  domain_client1:
    domain_id: 11155111
    type: EVM
    required_confirmations: 0
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    destination_address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
    light_inbox_address: 0x22EaC155C605Ffaf1002A2733206BCC0A52EB8C8
    light_manager_address: 0xC69119E94b63b03e7Fe80613F3717Df1CaD0055C
  domain_client2:
    domain_id: 901
    type: EVM
    required_confirmations: 0
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    summit_address: 0xFf9F696EE2A213FA0519673f6503F2933A574093
    inbox_address: 0x3E56136B612427AcCb501752e08c0872518Ae81E
    destination_address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
    bonding_manager_address: 0x5a4D4D0283fee81221DBDB545b98424fF5805f4c
  domain_client3:
    domain_id: 43113
    type: EVM
    required_confirmations: 0
    origin_address: 0x1314c7DBdeDEB905CD39D23BB9976090E0853861
    destination_address: 0x7E25eDfA872f1b0F9C4E9F890e645C9F619419E2
    light_inbox_address: 0x22EaC155C605Ffaf1002A2733206BCC0A52EB8C8
    light_manager_address: 0xC69119E94b63b03e7Fe80613F3717Df1CaD0055C

# Specify the summit domain id
summit_domain_id: 901
# A Notary's `domain_id` is the domain id of the chain it has posted a bond for.
domain_id: 11155111

# The `unbonded_signer` field specifies the path to the file containing the private key of the signer
unbonded_signer:
  type: "File
  file: "/config/notary_signer.txt"

# The `bonded_signer` is the account that will post a bond to the Summit contract. Specify its path to
# the file containing the private key of the signer
bonded_signer:
  type: "File"
  file: "/config/notary_signer.txt"

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
