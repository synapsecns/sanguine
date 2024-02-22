import { JsonRpcProvider } from '@ethersproject/providers'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits, parseUnits } from '@ethersproject/units'
import * as express from 'express'

import * as tokensList from './config/bridgeable'
import * as chainsData from './config/chains.json'
import { BRIDGE_MAP } from './config/bridgeMap'

// eslint-disable-next-line @typescript-eslint/no-unused-vars
interface Token {
  addresses: {
    [key: string]: string
  }
  decimals: number | { [key: string]: number }
  symbol: string
  name: string
  swapableType: string
  color: string
  visibilityRank: number
  priorityRank: number
  routeSymbol: string
  swapableOn?: number[]
  wrapperAddresses?: {
    [key: string]: string
  }
  isNative?: boolean
  docUrl?: string
}
// interface Tokens {
//   [key: string]: Token
// }

// const tokens: Tokens = tokensList as any

interface Chains {
  id: number
  name: string
  rpc: string
}

const chains: Chains[] = chainsData as any
// eslint-disable-next-line prefer-arrow/prefer-arrow-functions
function findTokenInfo(chain: string, tokenSymbol: string) {
  const chainData = BRIDGE_MAP[chain]
  if (!chainData) {
    return null
  }

  for (const tokenAddress in chainData) {
    if (chainData[tokenAddress].symbol === tokenSymbol) {
      return {
        address: tokenAddress,
        decimals: chainData[tokenAddress].decimals,
      }
    }
  }

  return null
}

const tokenHtml = Object.values(tokensList)
  .map((token: any) => {
    return (
      '<b>Token: ' +
      token.symbol +
      '</b><br/>' +
      Object.keys(token.addresses)
        .map((chainId) => {
          const tokenAddress = token.addresses[chainId]
          return (
            '<li>Chain Id: ' +
            String(chainId) +
            ', Address: ' +
            String(tokenAddress) +
            '</li>'
          )
        })
        .join('') +
      '</br>'
    )
  })
  .join('')

// Set up Synapse SDK
const providers = []
const chainIds = []

for (const chain of chains) {
  providers.push(new JsonRpcProvider(chain.rpc))
  chainIds.push(chain.id)
}
// Define the sdk
const Synapse = new SynapseSDK(chainIds, providers)

// Set up express server
const app = express()
const port = process.env.PORT || 3000

//Intro Message for UI
app.get('/', (_req, res) => {
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

app.get('/tokenList', (_req, res) => {
  res.send(tokensList)
})

//Swap Quote get request
app.get('/swap', async (req, res) => {
  try {
    // Access query params
    const query = req.query

    // Chain
    const chainId = query.chain

    // Symbols
    const fromTokenSymbol = String(query.fromToken)
    const toTokenSymbol = String(query.toToken)

    const fromTokenInfo = findTokenInfo(chainId, fromTokenSymbol)
    const toTokenInfo = findTokenInfo(chainId, toTokenSymbol)

    if (!fromTokenInfo || !toTokenInfo) {
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
    // Get Token Addresses and decimals (new)
    const fromTokenAddress = fromTokenInfo.address
    const toTokenAddress = toTokenInfo.address
    const fromTokenDecimals = fromTokenInfo.decimals
    const toTokenDecimals = toTokenInfo.decimals

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
    const amount = parseUnits(
      query.amount.toString(),
      fromTokenDecimals
    ).toString()
    // Send request w/Synapse SDK
    Synapse.swapQuote(
      Number(chainId),
      fromTokenAddress,
      toTokenAddress,
      BigNumber.from(amount)
    )
      .then((resp) => {
        const payload: any = resp
        payload.maxAmountOutStr = formatBNToString(
          resp.maxAmountOut,
          toTokenDecimals
        )
        res.json(payload)
      })
      .catch((err) => {
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
  } catch (err) {
    res.status(400).send({
      message: 'Invalid request',
      error: err.message,
    })
  }
})

//BridgeQuote get request
app.get('/bridge', async (req, res) => {
  // Access query params
  try {
    const query = req.query

    // Chains
    const fromChain = query.fromChain
    const toChain = query.toChain

    // Symbols
    const fromTokenSymbol = String(query.fromToken)
    const toTokenSymbol = String(query.toToken)

    // Get Token Info
    const fromTokenInfo = findTokenInfo(fromChain, fromTokenSymbol)
    const toTokenInfo = findTokenInfo(toChain, toTokenSymbol)

    if (!fromTokenInfo || !toTokenInfo) {
      res.send(
        `
        <h1>Invalid Params</h1>
        <hr/>
        <b>Ensure that your request matches the following format: /bridge?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000
        </b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
      )
      return
    }
    // Get Token Addresses and decimals (new)
    const fromTokenAddress = fromTokenInfo.address
    const toTokenAddress = toTokenInfo.address
    const fromTokenDecimals = fromTokenInfo.decimals
    const toTokenDecimals = toTokenInfo.decimals

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
    const amount = parseUnits(
      query.amount.toString(),
      fromTokenDecimals
    ).toString()
    // Send request w/Synapse SDK
    Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenAddress,
      toTokenAddress,
      BigNumber.from(amount)
    )
      .then((resp) => {
        const firstQuote = resp[0]
        const payload: any = firstQuote
        payload.maxAmountOutStr = formatBNToString(
          firstQuote.maxAmountOut,
          toTokenDecimals
        )
        payload.bridgeFeeFormatted = formatBNToString(
          firstQuote.feeAmount,
          toTokenDecimals
        )
        res.json(payload)
      })
      .catch((err) => {
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
  } catch (err) {
    res.status(400).send({
      message: 'Invalid request',
      error: err.message,
    })
  }
})

// Beginning of txInfo functions --> These return the txInfo to actually bridge
app.get('/swapTxInfo', async (req, res) => {
  try {
    // Access query params
    const query = req.query

    // Chain
    const chainId = query.chain

    // Symbols
    const fromTokenSymbol = String(query.fromToken)
    const toTokenSymbol = String(query.toToken)

    // Get Token Info
    const fromTokenInfo = findTokenInfo(chainId, fromTokenSymbol)
    const toTokenInfo = findTokenInfo(chainId, toTokenSymbol)

    if (!fromTokenInfo || !toTokenInfo) {
      res.send(
        `
        <h1>Invalid Params</h1>
        <hr/>
        <b>Ensure that your request matches the following format: /swapTxInfo?chain=1&fromToken=USDC&toToken=DAI&amount=100
        </b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
      )
      return
    }
    // Get Token Addresses and decimals (new)
    const fromTokenAddress = fromTokenInfo.address
    const toTokenAddress = toTokenInfo.address
    const fromTokenDecimals = fromTokenInfo.decimals
    const toTokenDecimals = toTokenInfo.decimals

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
    const amount = parseUnits(
      query.amount.toString(),
      fromTokenDecimals
    ).toString()

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
  } catch (err) {
    res.status(400).send({
      message: 'Invalid request',
      error: err.message,
    })
  }
})

//BridgeTxInfo
app.get('/bridgeTxInfo', async (req, res) => {
  try {
    // Access query params
    const query = req.query

    // Chains
    const fromChain = query.fromChain
    const toChain = query.toChain

    // Symbols
    const fromTokenSymbol = String(query.fromToken)
    const toTokenSymbol = String(query.toToken)

    // Get Token Info
    const fromTokenInfo = findTokenInfo(fromChain, fromTokenSymbol)
    const toTokenInfo = findTokenInfo(toChain, toTokenSymbol)

    if (!fromTokenInfo || !toTokenInfo) {
      res.send(
        `
        <h1>Invalid Params</h1>
        <hr/>
        <b>Ensure that your request matches the following format: /bridgeTxInfo?fromChain=1&toChain=42161&fromToken=USDC&toToken=USDC&amount=1000000&destAddress=0xcc78d2f004c9de9694ff6a9bbdee4793d30f3842
        </b>
        <h2>Available Tokens (symbols to use)</h2>
        ${tokenHtml}`
      )
      return
    }
    // Get Token Addresses and decimals (new)
    const fromTokenAddress = fromTokenInfo.address
    const toTokenAddress = toTokenInfo.address
    const fromTokenDecimals = fromTokenInfo.decimals
    const toTokenDecimals = toTokenInfo.decimals

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
    const amount = parseUnits(
      query.amount.toString(),
      fromTokenDecimals
    ).toString()

    // Send request with Synapse SDK
    Synapse.allBridgeQuotes(
      Number(fromChain),
      Number(toChain),
      fromTokenAddress,
      toTokenAddress,
      BigNumber.from(amount)
    )
      .then((resp) => {
        const firstQuote = resp[0]
        Synapse.bridge(
          destAddress,
          routerAddress,
          Number(fromChain),
          Number(toChain),
          fromTokenAddress,
          BigNumber.from(amount),
          firstQuote.originQuery,
          firstQuote.destQuery
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
  } catch (err) {
    res.status(400).send({
      message: 'Invalid request',
      error: err.message,
    })
  }
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
