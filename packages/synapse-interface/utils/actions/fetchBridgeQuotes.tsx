import { useState, useEffect } from 'react'
import { Address } from 'viem'
import { useSynapseContext } from '../providers/SynapseProvider'

interface BridgeQuoteRequest {
  originChainId: number
  destinationChainId: number
  originTokenAddress: Address
  destinationTokenAddress: Address
  amount: bigint
}

export function useBridgeQuote(request: BridgeQuoteRequest) {
  const [bridgeQuote, setBridgeQuote] = useState(null)
  const { synapseSDK } = useSynapseContext()

  console.log('request:', request)

  useEffect(() => {
    ;(async () => {
      if (request) {
        const {
          feeAmount,
          routerAddress,
          maxAmountOut,
          originQuery,
          destQuery,
        } = await synapseSDK.bridgeQuote(
          request.originChainId,
          request.destinationChainId,
          request.originTokenAddress,
          request.destinationTokenAddress,
          request.amount
        )

        setBridgeQuote({
          feeAmount: feeAmount,
          routerAddress: routerAddress,
          maxAmountOut: maxAmountOut,
          originQuery: originQuery,
          destQuery: destQuery,
        })
      }
    })()
  }, [request])

  return bridgeQuote
}
