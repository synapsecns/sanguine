import { useEffect } from 'react'
import type { Address } from 'viem'
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
import type { Token } from '@/utils/types'

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
    const isValueInvalid: boolean =
      hasOnlyZeroes(debouncedToTokensFromValue) ||
      isEmptyString(debouncedToTokensFromValue)

    const isSelectionsInvalid: boolean = [
      fromChainId,
      toChainId,
      fromToken,
      toToken,
    ].some(_.isNull)

    /** Conditions required for fetching bridge quotes */
    if (!isValueInvalid && !isSelectionsInvalid && synapseSDK) {
      const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
        (token: Token) => ({
          originChainId: fromChainId,
          originToken: fromToken,
          destinationChainId: toChainId,
          destinationTokenAddress: token.addresses[toChainId] as Address,
          destinationToken: token,
          amount: stringToBigInt(
            debouncedToTokensFromValue,
            fromToken?.decimals[fromChainId]
          ),
        })
      )
      dispatch(
        fetchAndStoreBridgeQuotes({
          requests: bridgeQuoteRequests,
          synapseSDK,
        })
      )
    }

    /** Reset bridge quotes state when value is invalid */
    if (isValueInvalid) {
      dispatch(resetFetchedBridgeQuotes())
    }
  }, [debouncedToTokensFromValue])
}
