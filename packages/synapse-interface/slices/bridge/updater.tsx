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
    toToken,
    toTokens,
    fromValue,
    debouncedFromValue,
    debouncedToTokensFromValue,
  }: BridgeState = useBridgeState()

  /**
   * Debounce user input to fetch primary bridge quote (in ms)
   * Delay loading animation when user input updates
   */
  useEffect(() => {
    const debounceDelay = 300
    const animationDelay = 200

    const animationTimer = setTimeout(() => {
      if (debouncedFromValue !== initialState.debouncedFromValue) {
        dispatch(setIsLoading(true))
      }
    }, animationDelay)

    const debounceTimer = setTimeout(() => {
      dispatch(updateDebouncedFromValue(fromValue))
    }, debounceDelay)

    return () => {
      clearTimeout(debounceTimer)
      clearTimeout(animationTimer)
    }
  }, [fromValue])

  // Debounce alternative destination token bridge quotes
  useEffect(() => {
    const alternativeOptionsDebounceDelay = 500

    const alternativeOptionsDebounceTimer = setTimeout(() => {
      dispatch(updateDebouncedToTokensFromValue(debouncedFromValue))
    }, alternativeOptionsDebounceDelay)

    return () => {
      clearTimeout(alternativeOptionsDebounceTimer)
    }
  }, [debouncedFromValue])

  // Conditions for fetching alternative bridge quotes
  useEffect(() => {
    if (fromChainId && toChainId && fromToken && toToken && synapseSDK) {
      const userInputExists: boolean = debouncedToTokensFromValue !== ''
      const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
        (token: Token) => {
          return {
            originChainId: fromChainId,
            originToken: fromToken as Token,
            destinationChainId: toChainId,
            destinationTokenAddress: token?.addresses[toChainId] as Address,
            destinationToken: token as Token,
            amount: stringToBigInt(
              userInputExists
                ? debouncedToTokensFromValue
                : getDefaultBridgeAmount(fromToken),
              fromToken.decimals[fromChainId]
            ),
          }
        }
      )

      if (userInputExists) {
        console.log('fetching')
        dispatch(
          fetchAndStoreBridgeQuotes({
            requests: bridgeQuoteRequests,
            synapseSDK,
          })
        )
      }
    }

    if (!fromToken) {
      dispatch(resetFetchedBridgeQuotes())
    }
  }, [debouncedToTokensFromValue, toTokens])

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
