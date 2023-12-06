import { useCallback } from 'react'
import { Address } from 'viem'

interface UseBridgeCallbackArgs {
  destinationAddress: Address
  originRouterAddress: Address
  originChainId: number
  destinationChainId: number
  tokenAddress: Address
  amount: bigint
  originQuery: any
  destinationQuery: any
  synapseSDK?: any
  signer?: any
}

export function useBridgeCallback({
  destinationAddress,
  originRouterAddress,
  originChainId,
  destinationChainId,
  tokenAddress,
  amount,
  originQuery,
  destinationQuery,
  synapseSDK,
  signer,
}: UseBridgeCallbackArgs) {
  const bridgeCallback = useCallback(async () => {
    /**
     * Dependency Guardrail checks
     * Bail on executing Bridge Callback if dependencies missing
     */
    if (!tokenAddress) {
      console.error('useBridgeCallback missing token address')
      return
    } else if (!originRouterAddress) {
      console.error('useBridgeCallback missing origin router address')
      return
    } else if (!destinationAddress) {
      console.error('useBridgeCallback missing destination address')
      return
    } else if (!originChainId) {
      console.error('useBridgeCallback missing origin chain id')
      return
    } else if (!destinationChainId) {
      console.error('useBridgeCallback missing destination chain id')
      return
    } else if (!originQuery) {
      console.error('useBridgeCallback missing origin query')
      return
    } else if (!destinationQuery) {
      console.error('useBridgeCallback missing destination query')
      return
    } else if (!synapseSDK) {
      console.error('useBridgeCallback missing synapseSDK')
      return
    } else if (!signer) {
      console.error('useBridgeCallback missing signer')
      return
    }

    const data = await synapseSDK.bridge(
      destinationAddress,
      originRouterAddress,
      originChainId,
      destinationChainId,
      tokenAddress,
      amount,
      originQuery,
      destinationQuery
    )

    const transactionHash = await signer.sendTransaction(data)

    return transactionHash
  }, [
    destinationAddress,
    originRouterAddress,
    originChainId,
    destinationChainId,
    tokenAddress,
    amount,
    originQuery,
    destinationQuery,
    synapseSDK,
    signer,
  ])

  return bridgeCallback
}
