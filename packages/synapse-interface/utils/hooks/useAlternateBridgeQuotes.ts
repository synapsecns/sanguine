import { useEffect } from 'react'
import { Address } from 'viem'

import { useAppDispatch } from '@/store/hooks'
import {
  useBridgeState,
  fetchAndStoreBridgeQuotes,
} from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { resetFetchedBridgeQuotes } from '@/slices/bridge/actions'
import { useSynapseContext } from '../providers/SynapseProvider'
import { hasOnlyZeroes } from '../hasOnlyZeroes'
import { isEmptyString } from '../isEmptyString'
import { BridgeQuoteRequest } from '../actions/fetchBridgeQuotes'
import { Token } from '../types'
import { stringToBigInt } from '../bigint/format'

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
