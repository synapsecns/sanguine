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
    const isValueInvalid =
      hasOnlyZeroes(debouncedToTokensFromValue) ||
      isEmptyString(debouncedToTokensFromValue)

    const isSelectionsInvalid = [
      fromChainId,
      toChainId,
      fromToken,
      toToken,
    ].some(_.isNull)

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

    if (isValueInvalid) {
      dispatch(resetFetchedBridgeQuotes())
    }
  }, [debouncedToTokensFromValue])
}
