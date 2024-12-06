# Synapse REST API
To make requests, use https://api.synapseprotocol.com/

To run locally:
```bash
yarn dev
```

# Documentation
[Swagger Documentation](https://api.synapseprotocol.com/api-docs/)

[GitBook Documentation](https://docs.synapseprotocol.com/developers/rest-api)

## REST API
The Synapse REST API provides a set of endpoints for quoting and executing cross-chain token swaps and bridges. It supports various functions including swap quotes, bridge quotes, transaction information retrieval, and token list management.

## Formatting Requests
Here's an example of how to format a bridge request:

```bash
/bridge?fromChain=1&toChain=42161&fromToken=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48&toToken=0xaf88d065e77c8cC2239327C5EDb3A432268e5831&amount=100
```

To use this in a basic application:

```javascript
async function getBridgeQuote() {
  const response = await fetch('https://api.synapseprotocol.com/bridge?fromChain=1&toChain=42161&fromToken=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48&toToken=0xaf88d065e77c8cC2239327C5EDb3A432268e5831&amount=100');
  const data = await response.json();
  console.log(data);
}
getBridgeQuote();
```


## Bridging

1. Get a Bridge Quote to confirm the expected amount out and the bridge route

```javascript
const quoteResponse = await fetch('https://api.synapseprotocol.com/bridge?fromChain=1&toChain=42161&fromToken=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48&toToken=0xaf88d065e77c8cC2239327C5EDb3A432268e5831&amount=100');
const quoteData = await quoteResponse.json();
```

2. Get the structured transaction information

```javascript
const txInfoResponse = await fetch('https://api.synapseprotocol.com/bridgeTxInfo?fromChain=1&toChain=42161&fromToken=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48&toToken=0xaf88d065e77c8cC2239327C5EDb3A432268e5831&amount=100&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842');
const txInfoData = await txInfoResponse.json();
```

3. Execute the transaction

```javascript
const provider = new ethers.providers.Web3Provider(window.ethereum);
const signer = provider.getSigner();
const transaction = await signer.sendTransaction(txInfoData);
const receipt = await transaction.wait();
```

## Other Functions:
1. `/destinationTokens`: This endpoint provides information about possible destination tokens for a given source chain and token. It's useful for showing users their options when initiating a cross-chain transfer.

```javascript
const response = await fetch('https://api.synapseprotocol.com/destinationTokens?fromChain=1&fromToken=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48');
const destinationTokens = await response.json();
```

2. `/bridgeLimits`: This endpoint returns the minimum and maximum amounts that can be bridged for a specific token pair. It's helpful for validating user input and displaying available limits.

```javascript
const limitsResponse = await fetch('https://api.synapseprotocol.com/bridgeLimits?fromChain=1&toChain=42161&fromToken=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48&toToken=0xaf88d065e77c8cC2239327C5EDb3A432268e5831');
const limitsData = await limitsResponse.json();
```

There are other additional functions that are included and documented in the Swagger documentation. Suggested changes to the API can be made by creating a new branch, making the changes, and opening a pull request. Any changes should include the appropriate test coverage and update the relevant documentation (this README, Swagger, and GitBook).
