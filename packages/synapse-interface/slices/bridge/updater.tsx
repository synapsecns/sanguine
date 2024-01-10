import { useEffect } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { Address } from 'viem'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState, setIsLoading, initialState } from './reducer'
import { BridgeQuoteRequest } from '@/utils/actions/fetchBridgeQuotes'
import { fetchAndStoreBridgeQuotes } from '@/slices/bridge/hooks'
import {
  resetFetchedBridgeQuotes,
  updateDebouncedFromValue,
  updateDebouncedToTokensFromValue,
} from './actions'
import { resetBridgeQuote } from './reducer'
import { Token } from '@/utils/types'
import { stringToBigInt } from '@/utils/bigint/format'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'

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
    bridgeQuote,
    debouncedFromValue,
    debouncedToTokensFromValue,
  }: BridgeState = useBridgeState()

  /**
   * Debounce user input to fetch primary bridge quote (in ms)
   * Delay loading animation when user input updates
   */
  useEffect(() => {
    const DEBOUNCE_DELAY = 300
    const ANIMATION_DELAY = 200

    const animationTimer = setTimeout(() => {
      if (debouncedFromValue !== initialState.debouncedFromValue) {
        dispatch(setIsLoading(true))
      }
    }, ANIMATION_DELAY)

    const debounceTimer = setTimeout(() => {
      dispatch(updateDebouncedFromValue(fromValue))
    }, DEBOUNCE_DELAY)

    return () => {
      clearTimeout(debounceTimer)
      clearTimeout(animationTimer)
      dispatch(setIsLoading(false))
    }
  }, [fromValue])

  // Debounce alternative destination token bridge quotes
  useEffect(() => {
    const ALTERNATE_OPTIONS_DEBOUNCE_DELAY = 1000

    const alternativeOptionsDebounceTimer = setTimeout(() => {
      dispatch(updateDebouncedToTokensFromValue(debouncedFromValue))
    }, ALTERNATE_OPTIONS_DEBOUNCE_DELAY)

    return () => {
      clearTimeout(alternativeOptionsDebounceTimer)
    }
  }, [debouncedFromValue])

  // Conditions for fetching alternative bridge quotes
  // useEffect(() => {
  //   const userInputExists: boolean =
  //     debouncedToTokensFromValue !== initialState.debouncedToTokensFromValue
  //   const userInputIsZero: boolean = hasOnlyZeroes(debouncedFromValue)

  //   if (!userInputIsZero) {
  //     if (fromChainId && toChainId && fromToken && toToken && synapseSDK) {
  //       const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
  //         (token: Token) => {
  //           return {
  //             originChainId: fromChainId,
  //             originToken: fromToken as Token,
  //             destinationChainId: toChainId,
  //             destinationTokenAddress: token?.addresses[toChainId] as Address,
  //             destinationToken: token as Token,
  //             amount: stringToBigInt(
  //               userInputExists
  //                 ? debouncedToTokensFromValue
  //                 : getDefaultBridgeAmount(fromToken),
  //               fromToken?.decimals[fromChainId]
  //             ),
  //           }
  //         }
  //       )
  //       dispatch(
  //         fetchAndStoreBridgeQuotes({
  //           requests: bridgeQuoteRequests,
  //           synapseSDK,
  //         })
  //       )
  //     }
  //   }

  //   if (!fromToken || !userInputExists) {
  //     dispatch(resetFetchedBridgeQuotes())
  //   }
  // }, [debouncedToTokensFromValue, toTokens])

  // Clear bridge quote if input is empty
  useEffect(() => {
    if (debouncedFromValue === initialState.debouncedFromValue) {
      dispatch(resetBridgeQuote())
    }
  }, [debouncedFromValue, bridgeQuote])

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
