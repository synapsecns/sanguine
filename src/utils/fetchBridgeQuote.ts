export interface BridgeQuoteRequest {
  originChainId: number
  destinationChainId: number
  originTokenAddress: string
  destinationTokenAddress: string
  amount: bigint
}

export async function fetchBridgeQuote(
  bridgeQuoteRequest: BridgeQuoteRequest,
  synapseSDK: any
) {
  const {
    originChainId,
    destinationChainId,
    originTokenAddress,
    destinationTokenAddress,
    amount,
  } = bridgeQuoteRequest

  const {
    feeAmount,
    routerAddress,
    maxAmountOut,
    originQuery,
    destQuery,
    estimatedTime,
    bridgeModuleName,
  } = await synapseSDK.bridgeQuote(
    originChainId,
    destinationChainId,
    originTokenAddress,
    destinationTokenAddress,
    amount
  )

  const x = BigInt(maxAmountOut.toString()) ?? 0n

  return {
    feeAmount,
    routerAddress,
    maxAmountOut,
    originQuery,
    destQuery,
    estimatedTime,
    bridgeModuleName,
    x,
  }
}
