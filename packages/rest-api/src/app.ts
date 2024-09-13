import express from 'express'
import { check, validationResult } from 'express-validator'
import { JsonRpcProvider } from '@ethersproject/providers'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { formatUnits, parseUnits } from '@ethersproject/units'

import { formatBNToString } from './utils/formatBNToString'
import * as tokensList from './constants/bridgeable'
import { CHAINS_ARRAY } from './constants/chains'
import { validateTokens } from './validations/validateTokens'
import { showFirstValidationError } from './validations/showFirstValidationError'
import { Chain } from './types'

const chains: Chain[] = CHAINS_ARRAY

const providers = []
const chainIds = []

for (const chain of chains) {
  providers.push(new JsonRpcProvider(chain.rpcUrls.primary))
  chainIds.push(chain.id)
}

const Synapse = new SynapseSDK(chainIds, providers)

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())

app.get('/', (_req, res) => {
  const tokensWithChains = Object.values(tokensList).map((token: any) => ({
    symbol: token.symbol,
    chains: Object.entries(token.addresses).map(([chainId, tokenAddress]) => ({
      chainId,
      address: tokenAddress,
    })),
  }))

  res.json({
    message: 'Welcome to the Synapse REST API for swap and bridge quotes',
    availableChains: chains.map((chain) => ({
      name: chain.name,
      id: chain.id,
    })),
    availableTokens: tokensWithChains,
  })
})

app.get('/tokenList', (_req, res) => {
  res.json(tokensList)
})

app.get(
  '/swap',
  [
    check('chain')
      .isNumeric()
      .custom((value) => chains.some((c) => c.id === Number(value)))
      .withMessage('Unsupported chain')
      .exists()
      .withMessage('chain is required'),
    validateTokens('chain', 'fromToken', 'fromToken'),
    validateTokens('chain', 'toToken', 'toToken'),
    check('amount').isNumeric(),
  ],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }
    try {
      const { chain, amount } = req.query
      const fromTokenInfo = res.locals.tokenInfo.fromToken
      const toTokenInfo = res.locals.tokenInfo.toToken

      const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)
      const quote = await Synapse.swapQuote(
        Number(chain),
        fromTokenInfo.address,
        toTokenInfo.address,
        amountInWei
      )
      res.json({
        maxAmountOut: formatUnits(quote.maxAmountOut, toTokenInfo.decimals),
        ...quote,
      })
    } catch (err) {
      res.status(500).json({ error: 'Server error' })
    }
  }
)

app.get(
  '/bridge',
  [
    check('fromChain')
      .isNumeric()
      .custom((value) => chains.some((c) => c.id === Number(value)))
      .withMessage('Invalid fromChain'),
    check('toChain')
      .isNumeric()
      .custom((value) => chains.some((c) => c.id === Number(value)))
      .withMessage('Invalid toChain'),
    validateTokens('fromChain', 'fromToken', 'fromToken'),
    validateTokens('toChain', 'toToken', 'toToken'),
    check('amount').isNumeric(),
  ],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }
    try {
      const { fromChain, toChain, amount } = req.query
      const fromTokenInfo = res.locals.tokenInfo.fromToken
      const toTokenInfo = res.locals.tokenInfo.toToken

      const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

      const resp = await Synapse.allBridgeQuotes(
        Number(fromChain),
        Number(toChain),
        fromTokenInfo.address,
        toTokenInfo.address,
        amountInWei
      )
      const payload = resp.map((quote) => ({
        ...quote,
        maxAmountOutStr: formatBNToString(
          quote.maxAmountOut,
          toTokenInfo.decimals
        ),
        bridgeFeeFormatted: formatBNToString(
          quote.feeAmount,
          toTokenInfo.decimals
        ),
      }))
      res.json(payload)
    } catch (err) {
      console.error('Error in bridge route:', err)
      res.status(500).json({ error: 'Server error', details: err.message })
    }
  }
)

app.get(
  '/swapTxInfo',
  [
    check('chain')
      .isNumeric()
      .custom((value) => chains.some((c) => c.id === Number(value)))
      .withMessage('Unsupported chain'),
    validateTokens('chain', 'fromToken', 'fromToken'),
    validateTokens('chain', 'toToken', 'toToken'),
    check('amount').isNumeric(),
  ],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }

    try {
      const { chain, amount } = req.query
      const fromTokenInfo = res.locals.tokenInfo.fromToken
      const toTokenInfo = res.locals.tokenInfo.toToken

      if (!fromTokenInfo || !toTokenInfo) {
        return res.status(400).json({ error: 'Invalid token symbol' })
      }

      const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

      const quote = await Synapse.swapQuote(
        Number(chain),
        fromTokenInfo.address,
        toTokenInfo.address,
        amountInWei
      )

      const txInfo = await Synapse.swap(
        Number(chain),
        fromTokenInfo.address,
        toTokenInfo.address,
        amountInWei,
        quote.query
      )

      res.json(txInfo)
    } catch (err) {
      res.status(500).json({ error: 'Server error' })
    }
  }
)

app.get(
  '/bridgeTxInfo',
  [
    check('fromChain')
      .isNumeric()
      .custom((value) => chains.some((c) => c.id === Number(value)))
      .withMessage('Invalid fromChain'),
    check('toChain')
      .isNumeric()
      .custom((value) => chains.some((c) => c.id === Number(value)))
      .withMessage('Invalid toChain'),
    validateTokens('fromChain', 'fromToken', 'fromToken'),
    validateTokens('toChain', 'toToken', 'toToken'),
    check('amount').isNumeric(),
    check('destAddress').isString(),
  ],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }

    try {
      const { fromChain, toChain, amount, destAddress } = req.query
      const fromTokenInfo = res.locals.tokenInfo.fromToken
      const toTokenInfo = res.locals.tokenInfo.toToken

      const amountInWei = parseUnits(amount.toString(), fromTokenInfo.decimals)

      const quotes = await Synapse.allBridgeQuotes(
        Number(fromChain),
        Number(toChain),
        fromTokenInfo.address,
        toTokenInfo.address,
        amountInWei
      )

      const txInfoArray = await Promise.all(
        quotes.map(async (quote) => {
          const txInfo = await Synapse.bridge(
            destAddress,
            quote.routerAddress,
            Number(fromChain),
            Number(toChain),
            fromTokenInfo.address,
            amountInWei,
            quote.originQuery,
            quote.destQuery
          )
          return txInfo
        })
      )
      res.json(txInfoArray)
    } catch (err) {
      res.status(500).json({ error: 'Server error' })
    }
  }
)

app.get(
  '/getSynapseTxId',
  [
    check('originChainId').isNumeric(),
    check('bridgeModule').isString(),
    check('txHash').isString(),
  ],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }

    try {
      const { originChainId, bridgeModule, txHash } = req.query

      const synapseTxId = await Synapse.getSynapseTxId(
        Number(originChainId),
        bridgeModule,
        txHash
      )
      res.json({ synapseTxId })
    } catch (err) {
      res.status(500).json({ error: 'Server error' })
    }
  }
)

app.get(
  '/getBridgeTxStatus',
  [
    check('destChainId').isNumeric(),
    check('bridgeModule').isString(),
    check('synapseTxId').isString(),
  ],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }

    try {
      const { destChainId, bridgeModule, synapseTxId } = req.query

      const status = await Synapse.getBridgeTxStatus(
        Number(destChainId),
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
      res.status(500).json({ error: 'Server error' })
    }
  }
)

app.get(
  '/getDestinationTx',
  [check('originChainId').isNumeric(), check('txHash').isString()],
  showFirstValidationError,
  async (req, res) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() })
    }

    try {
      const { originChainId, txHash } = req.query

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
      res.status(500).json({ error: 'Server error' })
    }
  }
)

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})
