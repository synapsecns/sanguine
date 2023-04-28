import { JsonRpcProvider } from '@ethersproject/providers'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'
import express from 'express'

import chains from './config/chains.json' assert { type: 'json' };
import tokens from './config/tokens.json' assert { type: 'json' };

// Constants
const TEN = BigNumber.from(10)
const tokenHtml = Object.keys(tokens).map(
  (symbol) =>
    "<b>" + String(symbol) + '</b>: <br/>' + String(Object.keys(tokens[symbol].addresses).map((addrChainId) => "<li>" + String(addrChainId) + ': ' + tokens[symbol].addresses[addrChainId] + '</li>').join(""))
    + "</br>").join("")


// Set up Synapse SDK
const providers = []
const chainIds = []
for (let i = 0; i < chains.length; i++) {
  providers.push(new JsonRpcProvider(chains[i].rpc))
  chainIds.push(chains[i].id)
}
const Synapse = new SynapseSDK(chainIds, providers)

// Format helper function
const formatBNToString = (
  bn,
  nativePrecison,
  decimalPlaces,
) => {
  const fullPrecision = formatUnits(bn, nativePrecison)
  const decimalIdx = fullPrecision.indexOf('.')
  if (decimalPlaces === undefined || decimalIdx === -1) {
    return fullPrecision
  } else {
    const rawNumber = Number(fullPrecision)
    return rawNumber === 0 ? rawNumber.toFixed(1) : rawNumber.toFixed(decimalPlaces)
  }
}


// Set up express server
const app = express()
const port = process.env.PORT || 3000

//Intro Message for UI
app.get('/', (req, res) => {
  res.send(
    `
    <h1>Welcome to the Synapse Rest API for swap and bridge quotes</h1>
    <hr/>
    <h2>Available Chains</h2>
    <ul>
     ${chains.map(
      (chain) => "<li>" + String(chain.name) + ' (' + String(chain.id) + ')' + '</li>'
    ).join("")}
    </ul>
    <h2>Available Tokens (symbols to use)</h2>
    ${tokenHtml}`
  )
})

//Swap get request
app.get('/swap', async (req, res) => {
  // Access query params
  const query = req.query

  // Chain
  const chainId = query.chain

  // Symbols
  const fromTokenSymbol = query.fromToken
  const toTokenSymbol = query.toToken

  // Get Token Addresses
  const fromTokenAddress = tokens[fromTokenSymbol]?.addresses?.[chainId]
  const toTokenAddress = tokens[toTokenSymbol]?.addresses?.[chainId]

  // Get Token Decimals
  const fromTokenDecimals = tokens[fromTokenSymbol]?.decimals?.[chainId]
  const toTokenDecimals = tokens[fromTokenSymbol]?.decimals?.[chainId]

  // Handle invalid params (either token symbols or chainIDs)
  // TODO: add error handling for missing params
  if (!fromTokenAddress || !toTokenAddress || !fromTokenDecimals || !toTokenDecimals) {
    res.send(
      `
      <h1>Invalid Params</h1>
      <hr/>
      <b>Ensure that your request matches the following format: /swap?chain=1fromToken=USDC&toToken=DAI&amount=100</b>
      <h2>Available Tokens (symbols to use)</h2>
      ${tokenHtml}`)
    return
  }

  // Handle amount
  const amount = BigNumber.from(query.amount).mul(TEN.pow(fromTokenDecimals))

  // Send request w/Synapse SDK
  Synapse.swapQuote(
    Number(chainId),
    fromTokenAddress,
    toTokenAddress,
    BigNumber.from(amount)
  ).then((resp) => {
    // Check for stable swap (going in its 6 decimals but coming out its 18 decimals so we need to adjust)
    // Using arbitrary 6 decimals as a threshold for now
    // TODO: Router contract v2 should return the amount out with decimals for the out-out token not the out-in token (eg.nusd).
    const adjustedDecimals = resp.maxAmountOut.gt(amount.mul(TEN.pow(6))) ? 18 : toTokenDecimals

    // Add response field with adjusted maxAmountOutStr (to account for decimals)
    resp.maxAmountOutStr = formatBNToString(resp.maxAmountOut, adjustedDecimals, 10)
    res.json(resp)
  }).catch((err) => {
    // TODO: do a better return here
    res.send(
      `
      <h1>Invalid Request</h1>
      <code>${err}</code>
      <hr/>
      <b>Ensure that your request matches the following format: /swap?chain=1fromToken=USDC&toToken=DAI&amount=100</b>
      <h2>Available Tokens (symbols to use)</h2>
      ${tokenHtml}`)
  })
})

//Bridge get request
app.get(
  '/bridge',
  async (req, res) => {
    // Access query params
    const query = req.query

    // Chains
    const fromChain = query.fromChain
    const toChain = query.toChain

    // Symbols
    const fromTokenSymbol = query.fromToken
    const toTokenSymbol = query.toToken

    // Get Token Addresses
    const fromTokenAddress = tokens[fromTokenSymbol]?.addresses?.[fromChain]
    const toTokenAddress = tokens[toTokenSymbol]?.addresses?.[toChain]

    // Get Token Decimals
    const fromTokenDecimals = tokens[fromTokenSymbol]?.decimals?.[fromChain]
    const toTokenDecimals = tokens[fromTokenSymbol]?.decimals?.[toChain]

    // Handle invalid params (either token symbols or chainIDs)
    // TODO: add error handling for missing params
    if (!fromTokenAddress || !toTokenAddress || !fromTokenDecimals || !toTokenDecimals) {
      res.send(
        `
        <h1>Invalid Request</h1>
        <hr/>
        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000</b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`)
      return
    }


    // Handle amount
    const amount = BigNumber.from(query.amount).mul(TEN.pow(fromTokenDecimals))

    // Send request w/Synapse SDK
    Synapse.bridgeQuote(
      Number(fromChain),
      Number(toChain),
      fromTokenAddress,
      toTokenAddress,
      BigNumber.from(amount)
    ).then((resp) => {
      // Check for stable swap (going in its 6 decimals but coming out its 18 decimals so we need to adjust)
      // Using arbitrary 6 decimals as a threshold for now
      // TODO: Router contract v2 should return the amount out with decimals for the out-out token not the out-in token (eg.nusd).
      const adjustedDecimals = resp.maxAmountOut.gt(amount.mul(TEN.pow(6))) ? 18 : toTokenDecimals

      // Add response field with adjusted maxAmountOutStr (to account for decimals)
      resp.maxAmountOutStr = formatBNToString(resp.maxAmountOut, adjustedDecimals, 10)

      res.json(resp)
    }).catch((err) => {
      // TODO: do a better return here
      res.send(
        `
        <h1>Invalid Request</h1>
        <code>${err}</code>
        <hr/>
        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000</b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`)
    })
  }
)

app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})
