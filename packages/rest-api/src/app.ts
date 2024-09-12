import { JsonRpcProvider } from '@ethersproject/providers'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits, parseUnits } from '@ethersproject/units'
import express from 'express'

import * as tokensList from './constants/bridgeable'
import { CHAINS_ARRAY } from './constants/chains'
import { Chain } from './types'
import { BRIDGE_MAP } from './constants/bridgeMap'

const chains: Chain[] = CHAINS_ARRAY
const findTokenInfo = (chain: string, tokenSymbol: string) => {
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
    const chainList = Object.entries(token.addresses)
      .map(
        ([chainId, tokenAddress]) =>
          `<li>Chain Id: ${chainId}, Address: ${tokenAddress}</li>`
      )
      .join('')

    return `
      <b>Token: ${token.symbol}</b><br/>
      ${chainList}
      <br/>
    `
  })
  .join('')

// Set up Synapse SDK
const providers = []
const chainIds = []

for (const chain of chains) {
  providers.push(new JsonRpcProvider(chain.rpcUrls.primary))
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
     ${chains.map((chain) => `<li>${chain.name} (${chain.id})</li>`).join('')}
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
    const chainId = String(query.chain)

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
    const fromChain = String(query.fromChain)
    const toChain = String(query.toChain)

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
        const payload = resp.map((quote) => ({
          ...quote,
          maxAmountOutStr: formatBNToString(
            quote.maxAmountOut,
            toTokenDecimals
          ),
          bridgeFeeFormatted: formatBNToString(
            quote.feeAmount,
            toTokenDecimals
          ),
        }))
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
    const chainId = String(query.chain)

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
    const fromChain = String(query.fromChain)
    const toChain = String(query.toChain)

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

    //Get to Address on destination chain
    const destAddress = String(query.destAddress)

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
      .then(async (resp) => {
        const txInfoArray = await Promise.all(
          resp.map(async (quote) => {
            const txInfo = await Synapse.bridge(
              destAddress,
              quote.routerAddress,
              Number(fromChain),
              Number(toChain),
              fromTokenAddress,
              BigNumber.from(amount),
              quote.originQuery,
              quote.destQuery
            )
            return txInfo
          })
        )
        res.json(txInfoArray)
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

// Get Synapse Transaction ID
app.get('/getSynapseTxId', (req, res) => {
  try {
    const query = req.query
    const originChainId = Number(query.originChainId)
    const bridgeModule = String(query.bridgeModule)
    const txHash = String(query.txHash)

    if (!originChainId || !bridgeModule || !txHash) {
      res.status(400).send({
        message: 'Invalid request: Missing required parameters',
      })
      return
    }

    Synapse.getSynapseTxId(originChainId, bridgeModule, txHash)
      .then((synapseTxId) => {
        res.json({ synapseTxId })
      })
      .catch((err) => {
        res.status(400).send({
          message:
            'Ensure that your request matches the following format: /getSynapseTxId?originChainId=8453&bridgeModule=SynapseRFQ&txHash=0x4acd82091b54cf584d50adcad9f57c61055beaca130016ecc3798d3d61f5487a',
          error: err.message,
        })
      })
  } catch (err) {
    res.status(400).send({
      message:
        'Ensure that your request matches the following format: /getSynapseTxId?originChainId=8453&bridgeModule=SynapseRFQ&txHash=0x4acd82091b54cf584d50adcad9f57c61055beaca130016ecc3798d3d61f5487a',
      error: err.message,
    })
  }
})

// Get Bridge Transaction Status
app.get('/getBridgeTxStatus', async (req, res) => {
  try {
    const query = req.query
    const destChainId = Number(query.destChainId)
    const bridgeModule = String(query.bridgeModule)
    const synapseTxId = String(query.synapseTxId)

    if (!destChainId || !bridgeModule || !synapseTxId) {
      res.status(400).send({
        message: 'Invalid request: Missing required parameters',
      })
      return
    }

    try {
      const status = await Synapse.getBridgeTxStatus(
        destChainId,
        bridgeModule,
        synapseTxId
      )

      if (status) {
        const txIdWithout0x = synapseTxId.startsWith('0x')
          ? synapseTxId.slice(2)
          : synapseTxId
        const graphqlEndpoint = 'https://explorer.omnirpc.io/graphql'
        const graphqlQuery = `
          {
            bridgeTransactions(
              useMv: true
              kappa: "${txIdWithout0x}"
            ) {
              toInfo {
                chainID
                address
                txnHash
                USDValue
                tokenSymbol
                blockNumber
                formattedTime
              }
            }
          }
        `

        const graphqlResponse = await fetch(graphqlEndpoint, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ query: graphqlQuery }),
        })

        const graphqlData = await graphqlResponse.json()
        const toInfo = graphqlData.data.bridgeTransactions[0]?.toInfo || null

        res.json({ status, toInfo })
      } else {
        res.json({ status })
      }
    } catch (err) {
      res.status(400).send({
        message: 'Error fetching bridge transaction status',
        error: err.message,
      })
    }
  } catch (err) {
    res.status(400).send({
      message: 'Invalid request',
      error: err.message,
    })
  }
})

// Get Destination Transaction Hash
app.get('/getDestinationTx', async (req, res) => {
  try {
    const query = req.query
    const originChainId = Number(query.originChainId)
    const txHash = String(query.txHash)

    if (!originChainId || !txHash) {
      res.status(400).send({
        message: 'Invalid request: Missing required parameters',
      })
      return
    }

    try {
      const graphqlEndpoint = 'https://explorer.omnirpc.io/graphql'
      const graphqlQuery = `
          {
            bridgeTransactions(
              useMv: true
              chainIDFrom: ${originChainId}
              txnHash: "${txHash}"
            ) {
              toInfo {
                chainID
                address
                txnHash
                USDValue
                tokenSymbol
                blockNumber
                formattedTime
              }
            }
          }
        `

      const graphqlResponse = await fetch(graphqlEndpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query: graphqlQuery }),
      })

      const graphqlData = await graphqlResponse.json()
      const toInfo = graphqlData.data.bridgeTransactions[0]?.toInfo || null

      if (toInfo === null) {
        res.json({ status: 'pending' })
      } else {
        res.json({ status: 'completed', toInfo })
      }
    } catch (err) {
      res.status(400).send({
        message: 'Error fetching bridge transaction status',
        error: err.message,
      })
    }
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
