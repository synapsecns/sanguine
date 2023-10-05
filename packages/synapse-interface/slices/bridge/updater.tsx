import { useEffect, useState } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { Address } from 'viem'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState, setIsLoading, initialState } from './reducer'
import {
  BridgeQuoteRequest,
  fetchBridgeQuotes,
} from '@/utils/actions/fetchBridgeQuotes'
import { fetchAndStoreBridgeQuotes } from '@/slices/bridge/hooks'
import {
  resetFetchedBridgeQuotes,
  updateDebouncedFromValue,
  updateDebouncedToTokensFromValue,
} from './actions'
import { BridgeQuote, Token } from '@/utils/types'
import { stringToBigInt } from '@/utils/bigint/format'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { synapseSDK } = useSynapseContext()
  const {
    fromChainId,
    toChainId,
    fromToken,
    toTokens,
    fromValue,
    debouncedFromValue,
    debouncedToTokensFromValue,
  }: BridgeState = useBridgeState()

  // Debounce user input to prevent unnecessary quote fetching
  useEffect(() => {
    const debounceDelay = 200
    const alternativeDebounceDelay = 500
    const animationDelay = 200
    // dispatch(setIsLoading(true))

    const animationTimer = setTimeout(() => {
      if (debouncedFromValue !== initialState.debouncedFromValue) {
        dispatch(setIsLoading(true))
      }
    }, animationDelay)

    const debounceTimer = setTimeout(() => {
      dispatch(updateDebouncedFromValue(fromValue))
    }, debounceDelay)

    const alternativeDebounceTimer = setTimeout(() => {
      dispatch(updateDebouncedToTokensFromValue(fromValue))
    }, alternativeDebounceDelay)

    return () => {
      clearTimeout(debounceTimer)
      clearTimeout(animationTimer)
      clearTimeout(alternativeDebounceTimer)
      dispatch(setIsLoading(false))
    }
  }, [fromValue])

  // Conditions for fetching alternative bridge quotes
  useEffect(() => {
    if (fromChainId && toChainId && fromToken && synapseSDK) {
      const hasFromValue: boolean = debouncedToTokensFromValue !== ''
      const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
        (token: Token) => {
          return {
            originChainId: fromChainId,
            originToken: fromToken as Token,
            destinationChainId: toChainId,
            destinationTokenAddress: token?.addresses[toChainId] as Address,
            destinationToken: token as Token,
            amount: stringToBigInt(
              hasFromValue
                ? debouncedToTokensFromValue
                : getDefaultBridgeAmount(fromToken),
              fromToken.decimals[fromChainId]
            ),
          }
        }
      )

      dispatch(
        fetchAndStoreBridgeQuotes({
          requests: bridgeQuoteRequests,
          synapseSDK,
        })
      )
    }

    if (!fromToken) {
      dispatch(resetFetchedBridgeQuotes())
    }
  }, [
    fromChainId,
    toChainId,
    fromToken,
    // debouncedFromValue,
    debouncedToTokensFromValue,
    toTokens,
    synapseSDK,
  ])

  return null
}

enum SwappableTypes {
  STABLE = 'USD',
  ETH = 'ETH',
  BTC = 'WBTC',
}

enum DefaultBridgeAmount {
  STABLE = '50',
  ETH = '0.01',
  BTC = '0.001',
}

export const getDefaultBridgeAmount = (
  originToken: Token
): DefaultBridgeAmount => {
  const swappableType: string = originToken.swapableType

  switch (swappableType) {
    case SwappableTypes.STABLE:
      return DefaultBridgeAmount.STABLE
    case SwappableTypes.ETH:
      return DefaultBridgeAmount.ETH
    case SwappableTypes.BTC:
      return DefaultBridgeAmount.BTC
    default:
      return DefaultBridgeAmount.STABLE
  }
}
