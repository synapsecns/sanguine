import { Address } from 'viem'
import { BridgeQuote, Token } from '@/utils/types'
import { formatBigIntToString } from '../bigint/format'
import { commify } from '@ethersproject/units'
import { calculateExchangeRate } from '../calculateExchangeRate'

export interface BridgeQuoteResponse extends BridgeQuote {
  destinationToken: Token
  destinationChainId: number
}

export interface BridgeQuoteRequest {
  originChainId: number
  originToken: Token
  destinationChainId: number
  destinationTokenAddress: Address
  destinationToken: Token
  amount: bigint
}

export async function fetchBridgeQuote(
  request: BridgeQuoteRequest,
  synapseSDK: any
): Promise<BridgeQuoteResponse> {
  if (request && synapseSDK) {
    try {
      const {
        originChainId,
        originToken,
        destinationChainId,
        destinationTokenAddress,
        destinationToken,
        amount,
      }: BridgeQuoteRequest = request

      const {
        feeAmount,
        routerAddress,
        maxAmountOut,
        originQuery,
        destQuery,
        estimatedTime,
        bridgeModuleName,
        gasDropAmount,
      } = await synapseSDK.bridgeQuote(
        originChainId,
        destinationChainId,
        originToken.addresses[originChainId],
        destinationTokenAddress,
        amount
      )

      const toValueBigInt: bigint = BigInt(maxAmountOut.toString()) ?? 0n
      // Bridge Lifecycle: originToken -> bridgeToken -> destToken
      // amount is in originToken decimals
      // originQuery.minAmountOut and feeAmount is in bridgeToken decimals
      // Adjust feeAmount to be in originToken decimals
      const adjustedFeeAmount =
        (BigInt(feeAmount) * BigInt(amount)) / BigInt(originQuery.minAmountOut)

      const {
        originQuery: originQueryWithSlippage,
        destQuery: destQueryWithSlippage,
      } = synapseSDK.applyBridgeSlippage(
        bridgeModuleName,
        originQuery,
        destQuery
      )

      return {
        outputAmount: toValueBigInt,
        outputAmountString: commify(
          formatBigIntToString(
            toValueBigInt,
            destinationToken.decimals[destinationChainId],
            8
          )
        ),
        routerAddress,
        allowance: null, // update for allowances
        exchangeRate: calculateExchangeRate(
          amount - adjustedFeeAmount,
          originToken.decimals[originChainId],
          toValueBigInt,
          destinationToken.decimals[destinationChainId]
        ),
        feeAmount,
        delta: BigInt(maxAmountOut.toString()),
        quotes: {
          originQuery: originQueryWithSlippage,
          destQuery: destQueryWithSlippage,
        },
        destinationToken: request.destinationToken,
        destinationChainId: destinationChainId,
        estimatedTime: estimatedTime,
        bridgeModuleName: bridgeModuleName,
        gasDropAmount: BigInt(gasDropAmount.toString()),
      }
    } catch (error) {
      console.error('Error fetching bridge quote:', error)
      throw error
    }
  } else {
    console.error('Missing request or synapse SDK')
  }
}

export async function fetchBridgeQuotes(
  requests: BridgeQuoteRequest[],
  synapseSDK: any,
  maxConcurrentRequests: number = 2, // Set the maximum number of concurrent requests
  requestDelay: number = 1000 // Set the delay between requests in milliseconds (adjust as needed)
): Promise<BridgeQuoteResponse[]> {
  try {
    const bridgeQuotes: BridgeQuoteResponse[] = []

    for (let i = 0; i < requests.length; i += maxConcurrentRequests) {
      const batchRequests = requests.slice(i, i + maxConcurrentRequests)
      const bridgeQuotesPromises: Promise<BridgeQuoteResponse>[] =
        batchRequests.map(async (request: BridgeQuoteRequest) => {
          try {
            const results: BridgeQuoteResponse = await fetchBridgeQuote(
              request,
              synapseSDK
            )
            return results
          } catch (error) {
            console.error('Error in individual bridge quote request: ', error)
            return null
          }
        })

      const batchBridgeQuotes = (
        await Promise.all(bridgeQuotesPromises)
      ).filter((quote) => quote !== null)
      bridgeQuotes.push(...batchBridgeQuotes)

      // Add a delay between batches of requests to avoid overloading the server
      if (i + maxConcurrentRequests < requests.length) {
        await new Promise((resolve) => setTimeout(resolve, requestDelay))
      }
    }

    return bridgeQuotes
  } catch (error) {
    console.error('Error fetching bridge quotes: ', error)
    return []
  }
}

export function locateBestExchangeRateToken(
  quotes: BridgeQuoteResponse[]
): Token | null {
  if (quotes?.length === 0) {
    return null
  }

  let bestQuote: BridgeQuoteResponse | null = null

  quotes.forEach((quote) => {
    if (!bestQuote || quote.exchangeRate > bestQuote.exchangeRate) {
      bestQuote = quote
    }
  })

  return bestQuote ? bestQuote.destinationToken : null
}
