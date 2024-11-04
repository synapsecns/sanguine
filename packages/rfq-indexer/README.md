# RFQ Indexer

## Overview

The RFQ (Request for Quote) Indexer is a system designed to index and track bridge events across multiple blockchain networks. It consists of two main parts: the indexer and the API.

1. What does the rfq-indexer do?
   The rfq-indexer captures and stores bridge events from various blockchain networks, including Ethereum, Optimism, Arbitrum, Base, Blast, Scroll, Linea, and BNB Chain. It indexes events such as bridge requests, relays, proofs, refunds, and claims.

2. Parts of the indexer and their users:
   - Indexer: Used by developers and system administrators to collect and store blockchain data.
   - API: Used by front-end applications, other services, or developers to query the indexed data.

## Directory Structure
<pre>
rfq-indexer
├── <a href="./api">api</a>: API service
│   ├── src/ : API source code
│   ├── package.json : API dependencies and scripts
│   ├── README.md : API documentation
├── <a href="./indexer">indexer</a>: Indexer service
│   ├── src/ : Indexer source code
│   ├── abis/ : Contract ABIs
│   ├── package.json : Indexer dependencies and scripts
│   ├── README.md : Indexer documentation
</pre>


## Example Usage

This example demonstrates how to retrieve details for a specific transaction using the RFQ Indexer API.

## Basic Example
To query information for a specific transaction, use the following `curl` command:

```bash
curl -X 'GET' 'http://api.synapseprotocol.com/transaction-id/{transaction-id}' -H 'accept: application/json'
```
Replace {transaction-id} with the actual transaction ID. This request will return details such as transaction status, timestamps, and associated events in JSON format.

For complete API documentation, including additional endpoints and query parameters, please refer to the [API Documentation](http://api.synapseprotocol.com/api-docs).
