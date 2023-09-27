import { Address } from 'viem'
import { BridgeQuote } from '../types'

export interface BridgeQuoteRequest {
  originChainId: number
  destinationChainId: number
  originTokenAddress: Address
  destinationTokenAddress: Address
  amount: bigint
}

export async function fetchBridgeQuote(
  request: BridgeQuoteRequest,
  synapseSDK: any
): Promise<BridgeQuote> {
  if (request && synapseSDK) {
    return synapseSDK.bridgeQuote(
      request.originChainId,
      request.destinationChainId,
      request.originTokenAddress,
      request.destinationTokenAddress,
      request.amount
    )
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
