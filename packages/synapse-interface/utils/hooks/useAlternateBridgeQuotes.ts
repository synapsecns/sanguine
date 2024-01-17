import { useEffect } from 'react'
import { Address } from 'viem'
import _ from 'lodash'

import { useAppDispatch } from '@/store/hooks'
import {
  useBridgeState,
  fetchAndStoreBridgeQuotes,
} from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { resetFetchedBridgeQuotes } from '@/slices/bridge/actions'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { BridgeQuoteRequest } from '@/utils/actions/fetchBridgeQuotes'
import { stringToBigInt } from '@/utils/bigint/format'
import { isEmptyString } from '@/utils/isEmptyString'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { Token } from '@/utils/types'

/**
 * Hook to trigger fetching alternative bridge quotes
 * based on current user input conditions
 */
export const useAlternateBridgeQuotes = () => {
  const dispatch = useAppDispatch()
  const { synapseSDK } = useSynapseContext()
  const {
    fromChainId,
    fromToken,
    toTokens,
    toChainId,
    toToken,
    debouncedToTokensFromValue,
  }: BridgeState = useBridgeState()

  useEffect(() => {
    const isInputInvalid =
      hasOnlyZeroes(debouncedToTokensFromValue) ||
      isEmptyString(debouncedToTokensFromValue)

    if (
      !isInputInvalid &&
      fromChainId &&
      toChainId &&
      fromToken &&
      toToken &&
      synapseSDK
    ) {
      const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
        (token: Token) => {
          return {
            originChainId: fromChainId,
            originToken: fromToken as Token,
            destinationChainId: toChainId,
            destinationTokenAddress: token?.addresses[toChainId] as Address,
            destinationToken: token as Token,
            amount: stringToBigInt(
              debouncedToTokensFromValue,
              fromToken?.decimals[fromChainId]
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

    if (isInputInvalid) {
      dispatch(resetFetchedBridgeQuotes())
    }
  }, [debouncedToTokensFromValue])
}
