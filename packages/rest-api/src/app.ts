import { JsonRpcProvider } from '@ethersproject/providers'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'
import express from 'express'

import * as chainsData from './config/chains.json'
import * as tokensData from './config/tokens.json'

// import chains from './config/chains.json'
// import tokens from './config/tokens.json'
// To run locally you may need to add the "node --experimental-json-modules app.js" flag for the following jsons to be read

interface Tokens {
  [key: string]: {
    addresses: {
      [key: string]: string
    }
    decimals: {
      [key: string]: number
    }
    description: string
  }
}

const tokens: Tokens = tokensData as any

interface Chains {
  id: number
  name: string
  rpc: string
}

// @ts-ignore
const chains: Chains[] = chainsData.default as any
// Constants
const TEN = BigNumber.from(10)
const tokenHtml = Object.keys(tokens)
  .filter((symbol) => tokens[symbol].addresses !== undefined)
  .map(
    (symbol) =>
      '<b>' +
      String(symbol) +
      '</b>: <br/>' +
      String(
        Object.keys(tokens[symbol].addresses)
          .map(
            (addrChainId) =>
              '<li>' +
              String(addrChainId) +
              ': ' +
              tokens[symbol].addresses[addrChainId] +
              '</li>'
          )
          .join('')
      ) +
      '</br>'
  )
  .join('')

// Set up Synapse SDK
const providers = []
const chainIds = []

chains.map((chain) => {
  providers.push(new JsonRpcProvider(chain.rpc))
  chainIds.push(chain.id)
})
// Define the sdk
const Synapse = new SynapseSDK(chainIds, providers)

// Set up express server
const app = express()
const port = process.env.PORT || 3000

//Intro Message for UI
app.get('/', ({ res }) => {
  res.send(
    `
    <h1>Welcome to the Synapse Rest API for swap and bridge quotes</h1>
    <hr/>
    <h2>Available Chains</h2>
    <ul>
     ${chains
       .map(
         (chain) =>
           '<li>' + String(chain.name) + ' (' + String(chain.id) + ')' + '</li>'
       )
       .join('')}
    </ul>
    <h2>Available Tokens (symbols to use)</h2>
    ${tokenHtml}`
  )
})

//Swap Quote get request
app.get('/swap', async (req, res) => {
  // Access query params
  const query = req.query

  // Chain
  const chainId = query.chain

  // Symbols
  const fromTokenSymbol = String(query.fromToken)
  const toTokenSymbol = String(query.toToken)

  // Get Token Addresses
  const fromTokenAddress =
    tokens[fromTokenSymbol]?.addresses?.[chainId as string]
  const toTokenAddress = tokens[toTokenSymbol]?.addresses?.[chainId as string]

  // Get Token Decimals
  const fromTokenDecimals =
    tokens[fromTokenSymbol]?.decimals?.[chainId as string]
  const toTokenDecimals = tokens[toTokenSymbol]?.decimals?.[chainId as string]

  // Handle invalid params (either token symbols or chainIDs)
  // TODO: add error handling for missing params
  if (
    !fromTokenAddress ||
    !toTokenAddress ||
    !fromTokenDecimals ||
    !toTokenDecimals
  ) {
    res.send(
      `
      <h1>Invalid Params</h1>
      <hr/>
      <b>Ensure that your request matches the following format: /swap?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>
      <h2>Available Tokens (symbols to use)</h2>
      ${tokenHtml}`
    )
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
  )
    .then((resp) => {
      // Check for stable swap (going in its 6 decimals but coming out its 18 decimals so we need to adjust)
      // Using arbitrary 6 decimals as a threshold for now
      // TODO: Router contract v2 should return the amount out with decimals for the out-out token not the out-in token (eg.nusd).

      // Add response field with adjusted maxAmountOutStr (to account for decimals)
      const payload: any = resp
      payload.maxAmountOutStr = formatBNToString(
        resp.maxAmountOut,
        toTokenDecimals
      )
      res.json(payload)
    })
    .catch((err) => {
      // TODO: do a better return here
      res.send(
        `
      <h1>Invalid Request</h1>
      <code>${err}</code>
      <hr/>
      <b>Ensure that your request matches the following format: /swap?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>
      <h2>Available Tokens (symbols to use)</h2>
      ${tokenHtml}`
      )
    })
})

//BridgeQuote get request
app.get('/bridge', async (req, res) => {
  // Access query params
  const query = req.query

  // Chains
  const fromChain = query.fromChain
  const toChain = query.toChain

  // Symbols
  const fromTokenSymbol = String(query.fromToken)
  const toTokenSymbol = String(query.toToken)

  // Get Token Addresses
  const fromTokenAddress =
    tokens[fromTokenSymbol]?.addresses?.[fromChain as string]
  const toTokenAddress = tokens[toTokenSymbol]?.addresses?.[toChain as string]

  // Get Token Decimals
  const fromTokenDecimals =
    tokens[fromTokenSymbol]?.decimals?.[fromChain as string]
  const toTokenDecimals = tokens[toTokenSymbol]?.decimals?.[toChain as string]

  // Handle invalid params (either token symbols or chainIDs)
  // TODO: add error handling for missing params
  if (
    !fromTokenAddress ||
    !toTokenAddress ||
    !fromTokenDecimals ||
    !toTokenDecimals
  ) {
    res.send(
      `
        <h1>Invalid Request</h1>
        <hr/>
        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000</b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
    )
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
  )
    .then((resp) => {
      // Need to add some sort of execute function here
      // TODO: Router contract v2 should return the amount out with decimals for the out-out token not the out-in token (eg.nusd).

      // Add response field with adjusted maxAmountOutStr (to account for decimals)
      const payload: any = resp
      payload.maxAmountOutStr = formatBNToString(
        resp.maxAmountOut,
        toTokenDecimals
      )
      res.json(payload)
    })
    .catch((err) => {
      // TODO: do a better return here
      res.send(
        `
        <h1>Invalid Request</h1>
        <code>${err}</code>
        <hr/>
        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000</b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
      )
    })
})

// Beginning of txInfo functions --> These return the txInfo to actually bridge
app.get('/swapTxInfo', async (req, res) => {
  // Access query params
  const query = req.query

  // Chain
  const chainId = query.chain

  // Symbols
  const fromTokenSymbol = String(query.fromToken)
  const toTokenSymbol = String(query.toToken)

  // Get Token Addresses
  const fromTokenAddress =
    tokens[fromTokenSymbol]?.addresses?.[chainId as string]
  const toTokenAddress = tokens[toTokenSymbol]?.addresses?.[chainId as string]

  // Get Token Decimals
  const fromTokenDecimals =
    tokens[fromTokenSymbol]?.decimals?.[chainId as string]
  const toTokenDecimals = tokens[toTokenSymbol]?.decimals?.[chainId as string]

  // Handle invalid params (either token symbols or chainIDs)
  // TODO: add error handling for missing params
  if (
    !fromTokenAddress ||
    !toTokenAddress ||
    !fromTokenDecimals ||
    !toTokenDecimals
  ) {
    res.send(
      `
      <h1>Invalid Params</h1>
      <hr/>
      <b>Ensure that your request matches the following format: /swap?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>
      <h2>Available Tokens (symbols to use)</h2>
      ${tokenHtml}`
    )
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
  )
    .then((resp) => {
      Synapse.swap(
        Number(chainId),
        fromTokenAddress,
        toTokenAddress,
        BigNumber.from(amount),
        resp.query
      ).then((txInfo) => {
        res.json(txInfo)
      })
    })
    .catch((err) => {
      // TODO: do a better return here
      res.send(
        `
      <h1>Invalid Request</h1>
      <code>${err}</code>
      <hr/>
      <b>Ensure that your request matches the following format: /swapTxInfo?chain=1&fromToken=USDC&toToken=DAI&amount=100</b>
      <h2>Available Tokens (symbols to use)</h2>
      ${tokenHtml}`
      )
    })
})

//BridgeTxInfo
app.get('/bridgeTxInfo', async (req, res) => {
  // Access query params
  const query = req.query

  // Chains
  const fromChain = query.fromChain
  const toChain = query.toChain

  // Symbols
  const fromTokenSymbol = String(query.fromToken)
  const toTokenSymbol = String(query.toToken)

  // Get Token Addresses
  const fromTokenAddress =
    tokens[fromTokenSymbol]?.addresses?.[fromChain as string]
  const toTokenAddress = tokens[toTokenSymbol]?.addresses?.[toChain as string]

  // Get Token Decimals
  const fromTokenDecimals =
    tokens[fromTokenSymbol]?.decimals?.[fromChain as string]
  const toTokenDecimals = tokens[toTokenSymbol]?.decimals?.[toChain as string]

  //Get to Address on destination chain
  const destAddress = String(query.destAddress)

  //Router Address:
  const routerAddress = '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a'

  // Handle invalid params (either token symbols or chainIDs)
  // TODO: add error handling for missing params
  if (
    !fromTokenAddress ||
    !toTokenAddress ||
    !fromTokenDecimals ||
    !toTokenDecimals
  ) {
    res.send(
      `
        <h1>Invalid Request</h1>
        <hr/>
        <b>Ensure that your request matches the following format: /bridgeTxInfo?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842</b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
    )
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
  )
    .then((resp) => {
      Synapse.bridge(
        destAddress,
        routerAddress,
        Number(fromChain),
        Number(toChain),
        fromTokenAddress,
        BigNumber.from(amount),
        resp.originQuery,
        resp.destQuery
      ).then((txInfo) => {
        res.json(txInfo)
      })
    })
    .catch((err) => {
      // TODO: do a better return here
      res.send(
        `
        <h1>Invalid Request</h1>
        <code>${err}</code>
        <hr/>
        <b>Ensure that your request matches the following format: /bridgeTxInfo?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842</b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
      )
    })
})

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})

const formatBNToString = (
  bn: BigNumber,
  nativePrecison: number,
  decimalPlaces = 18
) => {
  const fullPrecision = formatUnits(bn, nativePrecison)
  const decimalIdx = fullPrecision.indexOf('.')

  if (decimalPlaces === undefined || decimalIdx === -1) {
    return fullPrecision
  } else {
    const rawNumber = Number(fullPrecision)

    if (rawNumber === 0) {
      return rawNumber.toFixed(1)
    }
    return rawNumber.toString()
  }
}
