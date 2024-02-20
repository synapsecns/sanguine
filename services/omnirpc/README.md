# OmniRPC

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/services/omnirpc.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/services/omnirpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/services/omnirpc)](https://goreportcard.com/report/github.com/synapsecns/sanguine/services/omnirpc)



A highly experimental rpc load balancer. This is designed to be embeddable in sanguine agents, but can also be used generally. It reduces trust in any one rpc provider by verifying the response against multiple.

For instance, the following config can be used to make sure a response is matched across 5 Ethereum rpc clients (all from chainlist):

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
# port to run on
port: 5000
# expressed in seconds
refreshInterval: 60
```

There are some nuances here. First of all, we can only check against "verifiable" queries not dependent on synchronization status. For instance:

```shell
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

Will check against 5 rpcs using the config above, but the following will only check against 1:

```shell
curl --location --request POST 'http://localhost:5000/rpc/1' \
--header 'Content-Type: application/json' \
--data-raw '{
	"jsonrpc":"2.0",
	"method":"eth_getTransactionCount",
	"params":[
		"0x230a1ac45690b9ae1176389434610b9526d2f21b",
		"latest"
	],
	"id":1
}'
```

This is because chains might have different states for "latest", resulting in false positives. The same goes for the following queries:

You can also query using a customizable number of confirmations using (for 2 confirmations on eth): `http://localhost:5000/confirmations/2/rpc/1`.

A postman collection is also available at `/collection.json`

Confirmable only when latest or pending are not passed:

- `eth_getBlockByNumber`
- `eth_getBlockTransactionCountByNumber`
- `eth_getBalance`
- `eth_getCode`
- `eth_getTransactionCount`
- `eth_call`
- `eth_getStorageAt`
- `eth_getLogs`

Never confirmable:

- `eth_blockNumber`
- `eth_syncing`
- `eth_gasPrice`
- `eth_maxPriorityFeePerGas`
- `eth_estimateGas`
- `eth_sendRawTransaction`: This is because the tx might not be pending

This is also verifiable in the headers. A successful response will have the following headers:

<!-- generated with https://www.tablesgenerator.com/markdown_tables -->

| Header Name              | Description                                                                                                                                                                            | Example Value (from request above)                                                                                                                                                       |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Content-Type             | Content type of the response                                                                                                                                                           | application/json                                                                                                                                                                         |
| X-Checked-URLS           | Comma-separated list of the response was checked against                                                                                                                               | https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7,https://rpc.ankr.com/eth,https://cloudflare-eth.com/,https://api.mycryptoapi.com/eth,https://nodes.mewapi.io/rpc/eth |
| X-Confirmable            | Whether the request is a confirmable type (see above)                                                                                                                                  | true                                                                                                                                                                                     |
| X-Forwarded-From         | The actual url the response was forwarded from. While json responses are asserted to be identical in substance (minified), there can be formatting difference/key-ordering differences | https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7                                                                                                                      |
| X-Json-Hash              | Sha256 hash of the minified json response.                                                                                                                                             | 17b9f3ec9687bb9ea7771d919cb19889b617868102a217b6761f86f4209f8d1f                                                                                                                         |
| X-Request-Id             | Request id used for tracing. This is a random-uuid if not passed by the user in the request                                                                                            | a75026e6-c8d6-46ac-a168-16163220765f                                                                                                                                                     |
| X-Required-Confirmations | Number of confirmations the request was checked against, always 1 if confirmable is false                                                                                              | 5                                                                                                                                                                                        |

# Chainlist

You can also quickly start a server running against all public chainlist rpcs with a confirmation threshold of 1. Just run `./omnirpc chainlist-server`

# To Do:

- Optionally use latest block number instead of latest to make latest queries confirmable
- Customizable Confirmation Count
- Real error handling

## Additional functionality

Please see [modules/README.md](modules/README.md) for additional functionality.
