# OmniRPC

OmniRPC is an RPC load balancer and verifier that allows users to query chain data from multiple chains. It is a service that should be run by Quoters and interfaces that allow Solvers to post quotes. OmniRPC takes in a yaml config that allows the user to specify which chains it should run on.

## Running OmniRPC

### Building From Source

To build OmniRPC from source, you will need to have Go installed. You can install Go by following the instructions [here](https://golang.org/doc/install). Once you have Go installed, you can build the relayer by running the following commands:

1. `git clone https://github.com/synapsecns/sanguine --recursive`
2. `cd sanguine/services/omnirpc`
3. `go run main.go --config /path/to/config.yaml`

### Running the Docker Image

The relayer can also be run with docker. To do this, you will need to pull the [docker image](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Frfq-relayer) and run it with the config file:

```bash
docker run ghcr.io/synapsecns/sanguine/omnirpc:latest --config /path/to/config
```

There is also a helm chart available for OmniRPC [here](https://artifacthub.io/packages/helm/synapse/omnirpc).

### Configuration

OmniRPC is configured with a yaml file. The following is an example configuration:

```yaml
chains:
  1:
    rpcs:
      - https://api.mycryptoapi.com/eth
      - https://rpc.flashbots.net/
      - https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79
      - https://cloudflare-eth.com/
      - https://mainnet-nethermind.blockscout.com/
      - https://nodes.mewapi.io/rpc/eth
      - https://main-rpc.linkpool.io/
      - https://mainnet.eth.cloud.ava.do/
      - https://ethereumnodelight.app.runonflux.io
      - https://rpc.ankr.com/eth
      - https://eth-rpc.gateway.pokt.network
      - https://main-light.eth.linkpool.io
      - https://eth-mainnet.public.blastapi.io
      - http://18.211.207.34:8545
      - https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7
      - wss://eth-mainnet.nodereal.io/ws/v1/1659dfb40aa24bbb8153a677b98064d7
      - https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet
    confirmations: 5
  10:
    rpcs:
      - https://mainnet.optimism.io
    confirmations: 1
# port to run on
port: 5000
# expressed in seconds
refreshInterval: 60
```

In this example, any request to ethereum (chain id: 1) will need at least 5 rpcs to agree on the data before it is considered valid, but op will only need 1.

Data can be fetched like so, where the last character is the chain id:

```bash
curl --location --request POST 'http://localhost:5000/rpc/1' \
--header 'Content-Type: application/json' \
--data-raw '{
	"jsonrpc":"2.0",
	"method":"eth_getTransactionCount",
	"params":[
		"0x230a1ac45690b9ae1176389434610b9526d2f21b",
		"0xec1d40"
	],
	"id":1
}'
```

A full postman collection can be found [here](https://github.com/synapsecns/sanguine/blob/master/services/omnirpc/swagger/collection.json) or at the `/collection.json endpoint`. Swagger docs are also available at `/swagger`.

Not all requests are confirmable. Please see [here](https://pkg.go.dev/github.com/synapsecns/sanguine/services/omnirpc#section-readme) for details.
