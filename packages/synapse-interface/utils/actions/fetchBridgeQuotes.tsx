import { Address } from 'viem'
import { BridgeQuote, Token } from '@/utils/types'
import {
  stringToBigInt,
  powBigInt,
  formatBigIntToString,
} from '../bigint/format'
import { subtractSlippage } from '../slippage'
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
    const {
      originChainId,
      originToken,
      destinationChainId,
      destinationTokenAddress,
      destinationToken,
      amount,
    }: BridgeQuoteRequest = request
    const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
      await synapseSDK.bridgeQuote(
        originChainId,
        destinationChainId,
        originToken.addresses[originChainId],
        destinationTokenAddress,
        amount
      )

    const toValueBigInt: bigint = BigInt(maxAmountOut.toString()) ?? 0n
    const originTokenDecimals: number = originToken.decimals[originChainId]
    const adjustedFeeAmount: bigint =
      BigInt(feeAmount) < amount
        ? BigInt(feeAmount)
        : BigInt(feeAmount) / powBigInt(10n, BigInt(18 - originTokenDecimals))

    const originMinWithSlippage = subtractSlippage(
      originQuery?.minAmountOut ?? 0n,
      'ONE_TENTH',
      null
    )
    const destMinWithSlippage = subtractSlippage(
      destQuery?.minAmountOut ?? 0n,
      'ONE_TENTH',
      null
    )

    let newOriginQuery = { ...originQuery }
    newOriginQuery.minAmountOut = originMinWithSlippage

    let newDestQuery = { ...destQuery }
    newDestQuery.minAmountOut = destMinWithSlippage

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
        originQuery: newOriginQuery,
        destQuery: newDestQuery,
      },
      destinationToken: request.destinationToken,
      destinationChainId: destinationChainId,
    }
  }
}

export async function fetchBridgeQuotes(
  requests: BridgeQuoteRequest[],
  synapseSDK: any
): Promise<BridgeQuoteResponse[]> {
  try {
    const bridgeQuotesPromises: Promise<BridgeQuoteResponse>[] = requests.map(
      async (request: BridgeQuoteRequest) => {
        const results: BridgeQuoteResponse = await fetchBridgeQuote(
          request,
          synapseSDK
        )

        return results
      }
    )
    const bridgeQuotes = await Promise.all(bridgeQuotesPromises)
    return bridgeQuotes
  } catch (e) {
    console.error('error from fetchBridgeQuotes: ', e)
    return []
  }
}

export function locateBestExchangeRateIndex(
  quotes: BridgeQuoteResponse[]
): number | null {
  if (quotes?.length === 0) {
    return null
  }

  return quotes?.reduce((indexOfHighest, currentQuote, currentIndex) => {
    if (currentQuote.exchangeRate > quotes[indexOfHighest].exchangeRate) {
      return currentIndex
    }
    return indexOfHighest
  }, 0)
}
