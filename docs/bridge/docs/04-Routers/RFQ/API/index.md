---
sidebar_position: 0
sidebar_label: API
---

# Quoter API

:::note

This guide is intended for builders who are integrating a quoter or frontend with the Synapse RFQ system.

If you are interested in running a relayer, please also see [Relayer](../Relayer).

:::

The Quoter API is an off-chain RESTful service that allows market makers / solvers to post open quotes which communicate an intent to fulfill any transaction that occurs upon specifically quoted routes and meets specified limits, pricing, and fee criteria. This method of quoting is referred to as "Passive" quoting and is similar to an order book.

Starting with [Fast Bridge V2](https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html), a new "Active" quoting method has been introduced where a solver can listen and respond to live quote requests individually. This creates a hybrid system, where Active and Passive quoting can be utilized together by solvers in any desired combination to maximize their efficiency.

Active Quoting is more complicated to implement and maintain, but allow for more granular & customized quotes that can improve efficiency among other benefits. Quoters who prefer a simpler approach are free to use nothing but Passive Quotes if they choose.

Integrators and users can then utilize the data from these quotes to construct and submit a corresponding transaction on-chain through the [Fast Bridge Contract](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html).

Quoters are responsible for keeping their quotes fresh and accurate. Likewise, they are responsible for completing their part of fulfillment for any transactions which act upon their quotes. To these effects, quoters should push updates as rapidly as possible in reaction to consequential changes in prices, balances, etc. By default, the canonical [relayer](../Relayer) continuously updates quotes by checking on-chain balances, in-flight requests, and gas prices - custom implementations should take a similar approach.

The implementation of the Quoter API can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq/api).

Please note that end-users and solvers will not need to run their own version of the API.


**Integrating the API**

## Passive Quotes

  ## Endpoints for Quoters:

  Authorized quoters can push passive quotes via these endpoints:

  - [`PUT /quotes`](./upsert-quote.api.mdx) - Upsert a passive quote
  - [`PUT /bulk_quotes`](./upsert-quotes.api.mdx) - Upsert an array of passive quotes in bulk

  ## Endpoints for Integrators / Users

  To view all current passive quotes, this permissionless endpoint can be used:

  - [`GET /quotes`](./get-quotes.api.mdx) - Get all quotes, can be filtered by different parameters.



## Active Quotes

  Active Quoting is more complicated than passive and requires listening for & responding to individual Requests for Quotes (RFQs).

  ## Endpoints for Quoters

  - [`GET /rfq_stream`](./rfq-stream.api.mdx) - Connect via WebSocket to listen for streamed RFQs
  - [`GET /rfq`](./get-rfq-request.api.mdx) - Retrieve currently open RFQs.

  ## Endpoints for Integrators / Users

  - [`PUT /rfq`](./put-rfq-request.api.mdx) - Initiate an RFQ and receive the best available quote.



**API Version Changes**

An http response header "X-Api-Version" will be returned on each call response.

Any systems that integrate with the API should use this header to detect version changes and perform appropriate follow-up actions & alerts.

Upon a version change, [versions.go](https://github.com/synapsecns/sanguine/blob/master/services/rfq/api/rest/versions.go) can be referred to for further detail on the version including deprecation alerts, etc.

Please note, while Synapse may choose to take additional steps to alert & advise on API changes through other communication channels, it will remain the responsibility of the API users & integrators to set up their own detection & notifications of version changes as they use these endpoints. Likewise, it will be their responsibility review the versions.go file, to research & understand how any changes may affect their integration, and to implement any necessary adjustments resulting from the API changes.

**Authentication & Authorization**

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

Once the message has been authenticated, the authorization of the sender/signer will be checked against the assigned roles of the respective FastBridge contract. If `RELAYER_ROLE` is not assigned, the request will be rejected. If you wish to be added as an authorized quoter, contact us.

:::

### API Urls

 - Mainnet: `api.synapseprotocol.com/quotes`
 - Testnet: `rfq-api-testnet.omnirpc.io`
 -

## Running the API:

Users and relayers **are not** expected to run their own version of the Quoter API. The API is a service that should be run by Quoters and interfaces that allow Solvers to post quotes. The Quoter API takes in a yaml config that allows the user to specify which contracts, chains and interfaces it should run on. The config is structured like this:

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

Yaml settings:

- `database` - The database settings for the API backend. A database is required to store quotes and other information. Using SQLite with a dsn set to a `/tmp/` directory is recommended for development.
  - `type` - the database driver to use, can be `mysql` or `sqlite`.
  - `dsn` - the dsn of your database. If using sqlite, this can be a path, if using mysql please see [here](https://dev.mysql.com/doc/connector-odbc/en/connector-odbc-configuration.html) for more information.
- `omnirpc_url` - The omnirpc url to use for querying chain data (no trailing slash). For more information on omnirpc, see [here](/docs/Services/Omnirpc).
- `bridges` - A key value map of chain id to FastBridge contract address. The API will only allow quotes to be posted on these chains.
- `port` - The port to run the http server on.

**Building From Source:**

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
