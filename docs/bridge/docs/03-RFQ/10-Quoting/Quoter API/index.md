---
sidebar_position: 0
sidebar_label: Quoter API
---

<!-- Reference Links -->
[relay]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#relayv2
[prove]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#provev2
[dispute]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#dispute
[claim]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#claimv2
[cancel]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#cancelv2
[proof]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgetxdetails
[BridgeRequested]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerequested
[BridgeTransactionV2]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgetransactionv2
[BridgeRelayed]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerelayed
[BridgeProofProvided]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeproofprovided
[Cancel Delay]: https://rfq-contracts.synapseprotocol.com/contracts/FastBridge.sol/contract.FastBridge.html#refund_delay
[Multicall]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IMulticallTarget.sol/interface.IMulticallTarget.html

[Quoter API]: /docs/RFQ/Quoting/Quoter%20API/
[Dispute Period]: /docs/RFQ/Security/#dispute-period
[Quoting]: /docs/RFQ/Quoting
[Bridging]: /docs/RFQ/Bridging
[Relaying]: /docs/RFQ/Relaying
[Proving]: /docs/RFQ/Proving
[Claiming]: /docs/RFQ/Claiming
[Canceling]: /docs/RFQ/Canceling
[Security]: /docs/RFQ/Security
[Exclusivity]: /docs/RFQ/Exclusivity

[User]: /docs/RFQ/#entities
[Quoter]: /docs/RFQ/#entities
[Prover]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

# Quoter API

:::info

This guide is intended for builders who are integrating a quoter or frontend with the Synapse RFQ system.

If you are interested in running a relayer, please also see [Relaying] and [Canonical Relayer](/docs/RFQ/CanonicalRelayer/) .

:::

The implementation of the Quoter API can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq/api).

Please note that end-users and relayers will not need to run their own version of the API.


## Integrating the API

### Passive Quotes

**Endpoints for Quoters**

Authorized quoters can push passive quotes via these endpoints:

- [`PUT /quotes`](./upsert-quote.api.mdx) - Upsert a passive quote
- [`PUT /bulk_quotes`](./upsert-quotes.api.mdx) - Upsert an array of passive quotes in bulk

**Endpoints for Integrators / Users**

To view all current passive quotes, this permissionless endpoint can be used:

- [`GET /quotes`](./get-quotes.api.mdx) - Get all quotes, can be filtered by different parameters.



### Active Quotes

Active Quoting is more complicated than passive and requires listening for & responding to individual Requests for Quotes (RFQs).

**Endpoints for Quoters**

- [`GET /rfq_stream`](./listen-for-active-rf-qs.api.mdx) - Connect via WebSocket to listen for streamed RFQs
- [`GET /rfq`](./get-open-quote-requests.api.mdx) - Retrieve currently open RFQs.

**Endpoints for Integrators / Users**

- [`PUT /rfq`](./initiate-an-active-rfq.api.mdx) - Initiate an RFQ and receive the best available quote.

  ## Websocket API for Quoters

  The websocket API allows quoters to interact with user quote requests once connected to the `GET /rfq_stream` endpoint.

  The websocket API exposes several operations for quoters:
  - `ping` - sends a heartbeat to the API server to keep the connection alive (must be sent at least once per minute)
  - `subscribe` - subscribes to quote requests for given chain(s)
  - `unsubscribe` - unsubscribes to quote requests for given chain(s)
  - `send_quote` - responds to a quote request

  The API server may respond with the following operations:`
  - `pong` - acknowleges a `ping` message
  - `request_quote` - informs quoter of a new user quote request

  All websocket messages follow this format:
  ```
  {
    op: string,
    content: json,
    success: bool,
  }
  ```

  Quote request content should have the following format:

  ```
  {
    data: {
      origin_chain_id: number,
      dest_chain_id: number,
      origin_token_addr: string,
      dest_token_addr: string,
      origin_amount_exact: string,
      expiration_window: number // number of ms since created_at until request should expire
    },
  }
  ```

  Quote response content should have the following format:

  ```
  {
    request_id: string,
    dest_amount: string,
  }
  ```

  Subscribe / Unsubscribe content should be an array of chain ids.


## API Version Changes

An http response header "X-Api-Version" will be returned on each call response.

Any systems that integrate with the API should use this header to detect version changes and perform appropriate follow-up actions & alerts.

Upon a version change, [versions.go](https://github.com/synapsecns/sanguine/blob/master/services/rfq/api/rest/versions.go) can be referred to for further detail on the version including deprecation alerts, etc.

Please note, while Synapse may choose to take additional steps to alert & advise on API changes through other communication channels, it will remain the responsibility of the API users & integrators to set up their own detection & notifications of version changes as they use these endpoints. Likewise, it will be their responsibility review the versions.go file, to research & understand how any changes may affect their integration, and to implement any necessary adjustments resulting from the API changes.

## Authentication & Authorization

In accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191), authorized Quoter API endpoints require a message signed by the relayer's address to accompany each request. The signature should be sent in the `Authorization` header of the request. We provide a client stub/example implementation in go [here](https://pkg.go.dev/github.com/synapsecns/sanguine/services/rfq@v0.13.3/api/client).

Additional Example in Typescript:

  ```typescript
  import { ecsign, toBuffer, bufferToHex, hashPersonalMessage } from 'ethereumjs-util';

  async function signMessage(privateKey: string) {
      const message = Math.floor(Date.now() / 1000).toString();
      const messageHash = hashPersonalMessage(Buffer.from(message));
      var { v, r, s } = ecsign(messageHash, toBuffer(privateKey));

      v -= 27

      const signature = Buffer.concat([r, s, Buffer.from([v])]);

      return `${message}:${bufferToHex(signature)}`;
  }
  ```

Once the message has been authenticated, the authorization of the sender/signer will be checked against the assigned roles of the respective FastBridge contract. If `QUOTER_ROLE` is not assigned, the request will be rejected. If you wish to be added as an authorized quoter, contact us.

:::

## API Urls

 - Mainnet: `api.synapseprotocol.com`
 - Testnet: `rfq-api-testnet.omnirpc.io`

## Running the API:

Users and relayers **are not** expected to run their own version of the Quoter API. Rather, they are expected to use a Quoter API that is hosted by the the interface they are quoting for. For example, the Quoter API used by the Synapse bridge interface is hosted at the URL above.


### Configuration

The Quoter API takes in a yaml config that allows the user to specify which contracts, chains and interfaces it should run on. The config is structured like this:

```yaml
database:
  type: mysql # can be other mysql or sqlite
  dsn: root:password@hostname:3306)/database?parseTime=true # should be the dsn of your database. If using sqlite, this can be a path
omnirpc_url: https://route-to-my-omnirpc # omnirpc route
bridges:
  1: '0x00......' # FastBridge address on ethereum (chain id: 1)
  10: '0x01....' # FastBridge address on op (chain id: 10)
port: '8081' # port to run your http server on
```

### YAML Descriptions

- `database` - The database settings for the API backend. A database is required to store quotes and other information. Using SQLite with a dsn set to a `/tmp/` directory is recommended for development.
  - `type` - the database driver to use, can be `mysql` or `sqlite`.
  - `dsn` - the dsn of your database. If using sqlite, this can be a path, if using mysql please see [here](https://dev.mysql.com/doc/connector-odbc/en/connector-odbc-configuration.html) for more information.
- `omnirpc_url` - The omnirpc url to use for querying chain data (no trailing slash). For more information on omnirpc, see [here](/docs/Services/Omnirpc).
- `bridges` - A key value map of chain id to FastBridge contract address. The API will only allow quotes to be posted on these chains.
- `port` - The port to run the http server on.

### Building from Source

To build the Quoter API from source, you will need to clone the repository and run the main.go file with the config file. Building from source requires go 1.21 or higher and is generally not recommended for end-users.

1. `git clone https://github.com/synapsecns/sanguine --recursive`
2. `cd sanguine/services/rfq`
3. `go run main.go --config /path/to/config.yaml`

**Running with Docker**

The Quoter API can also be run with docker. To do this, you will need to build the docker image and run it with the config file.

:::tip
Docker versions should always be pinned in production environments. For a full list of tags, see [here](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Frfq-api)
:::

1. `docker run ghcr.io/synapsecns/sanguine/rfq-api:latest --config /path/to/config`
