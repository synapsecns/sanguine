# Swap/Bridge REST API Quoter

To run locally:
\`npm start\`

To make requests, use https://synapse-rest-api-v2.herokuapp.com/

The Synapse Rest API supports four main functions

## /swap

which returns the following

- \`routerAddress\` (string) - The address of the router contract
- \`maxAmountOut\` (object) - The maximum amount of tokens that can be swapped out. Contains:
  - \`type\` (string) - The data type
  - \`hex\` (string) - The amount encoded in hexidecimal
- \`query\` (object) - Parameters for the swap query:
  - \`0\` (string) - Router contract address
  - \`1\` (string) - Address of tokenIn
  - \`2\` (object) - Amount of tokenIn to swap (same structure as maxAmountOut)
  - \`3\` (object) - Minimum amount of tokenOut requested (same structure as maxAmountOut)
  - \`4\` (string) - Encoded params for swap routing
  - \`swapAdapter\` (string) - Address of the swap adapter contract
  - \`tokenOut\` (string) - Address of tokenOut
  - \`minAmountOut\` (object) - Minimum amount of tokenOut required (same structure as maxAmountOut)
  - \`deadline\` (object) - Deadline parameter for the swap (same structure as maxAmountOut)
  - \`rawParams\` (string) - Encoded hex string containing swap parameters
- \`maxAmountOutStr\` (string) - The maxAmountOut value formatted as a decimal string

All \`/swap\` requests should be formatted like such:

\`/swap?chain=1&fromToken=USDC&toToken=DAI&amount=100\`

## /bridge

which returns all transaction information

- \`feeAmount\` (object) - The fee amount for the swap. Contains:
  - \`type\` (string) - Data type
  - \`hex\` (string) - Fee amount encoded in hex
- \`feeConfig\` (array) - Fee configuration parameters, contains:
  - \`0\` (number) - Gas price
  - \`1\` (object) - Fee percentage denominator (hex encoded BigNumber)
  - \`2\` (object) - Protocol fee percentage numerator (hex encoded BigNumber)
- \`routerAddress\` (string) - Address of the router contract
- \`maxAmountOut\` (object) - Maximum amount receivable from swap, structure same as above
- \`originQuery\` (object) - Original swap query parameters, contains:
  - \`swapAdapter\` (string) - Swap adapter address
  - \`tokenOut\` (string) - Address of output token
  - \`minAmountOut\` (object) - Minimum output token amount
  - \`deadline\` (object) - Expiry time
  - \`rawParams\` (string) - Encoded hex params
- \`destQuery\` (object) - Destination swap query parameters, structure similar to originQuery above.
- \`maxAmountOutStr\` (string) - maxAmountOut as a decimal string.

All \`/bridge\` requests should be formatted like such:

\`/bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000\`

## /swapTxInfo

which returns the following

- \`'data'\`: The binary data that forms the input to the transaction.
- \`'to'\`: The address of the Synapse Router (the synapse bridge contract)

All \`/swapTxInfo\` requests should be formatted like such:

\`/swap?chain=1&fromToken=USDC&toToken=DAI&amount=100\`

## /bridgeTxInfo

which returns the following

- \`'data'\`: The binary data that forms the input to the transaction.
- \`'to'\`: The address of the Synapse Router (the synapse bridge contract)

All \`/bridgeTxInfo\` requests should be formatted like such:

\`/bridgeTxInfo?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842\`
