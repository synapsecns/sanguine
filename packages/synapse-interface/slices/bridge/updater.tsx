import { useEffect, useMemo } from 'react'
import { Address } from 'viem'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from './reducer'
import {
  BridgeQuoteRequest,
  fetchBridgeQuotes,
} from '@/utils/actions/fetchBridgeQuotes'
import { BridgeQuote, Token } from '@/utils/types'
import { stringToBigInt } from '@/utils/bigint/format'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

export default function Updater(): null {
  const { synapseSDK } = useSynapseContext()
  const {
    fromChainId,
    toChainId,
    fromToken,
    toTokens,
    fromValue,
  }: BridgeState = useBridgeState()

  useEffect(() => {
    if (fromChainId && toChainId && fromToken && fromValue && synapseSDK) {
      const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
        (token: Token) => {
          return {
            originChainId: fromChainId,
            destinationChainId: toChainId,
            originTokenAddress: fromToken?.addresses[fromChainId] as Address,
            destinationTokenAddress: token?.addresses[toChainId] as Address,
            amount: stringToBigInt(fromValue, fromToken.decimals[fromChainId]),
            token: token,
          }
        }
      )

      ;(async () => {
        const bridgeQuotes: [BridgeQuote][] = await fetchBridgeQuotes(
          bridgeQuoteRequests,
          synapseSDK
        )

        console.log('bridgeQuotes:', bridgeQuotes)
      })()
    }
  }, [fromChainId, toChainId, fromToken, fromValue, toTokens, synapseSDK])

  return null
}
