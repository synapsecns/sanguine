import { Address } from 'viem'
import { BridgeQuote, Token } from '@/utils/types'

// Pass in Origin Token and Destination Token
// To allow fetchBridgeQuote to calculate exchange rate
export interface BridgeQuoteRequest {
  originChainId: number
  destinationChainId: number
  originToken: Token
  destinationTokenAddress: Address
  amount: bigint
  token: Token
}

export async function fetchBridgeQuote(
  request: BridgeQuoteRequest,
  synapseSDK: any
): Promise<BridgeQuote> {
  if (request && synapseSDK) {
    const bridgeQuote = await synapseSDK.bridgeQuote(
      request.originChainId,
      request.destinationChainId,
      request.originToken.addresses[request.originChainId],
      request.destinationTokenAddress,
      request.amount
    )

    return {
      ...bridgeQuote,
      token: request.token,
    }
  }
}

export async function fetchBridgeQuotes(
  requests: BridgeQuoteRequest[],
  synapseSDK: any
): Promise<[BridgeQuote][]> {
  try {
    const bridgeQuotesPromises: Promise<[BridgeQuote]>[] = requests.map(
      async (request: BridgeQuoteRequest) => {
        const results: [BridgeQuote] = await Promise.all([
          fetchBridgeQuote(request, synapseSDK),
        ])
        return results
      }
    )
    const bridgeQuotes = await Promise.all(bridgeQuotesPromises)
    return bridgeQuotes
  } catch (e) {
    console.error('error from fetchBridgeQuotes: ', e)
  }
}
