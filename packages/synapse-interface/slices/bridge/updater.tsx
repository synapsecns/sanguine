import { useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { Address } from 'viem'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from './reducer'
import {
  BridgeQuoteRequest,
  fetchBridgeQuotes,
} from '@/utils/actions/fetchBridgeQuotes'
import { fetchAndStoreBridgeQuotes } from '@/slices/bridge/hooks'
import { resetFetchedBridgeQuotes } from './actions'
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
  }: BridgeState = useBridgeState()

  useEffect(() => {
    if (fromChainId && toChainId && fromToken && synapseSDK) {
      const hasFromValue: boolean = fromValue !== ''
      const bridgeQuoteRequests: BridgeQuoteRequest[] = toTokens.map(
        (token: Token) => {
          return {
            originChainId: fromChainId,
            originToken: fromToken as Token,
            destinationChainId: toChainId,
            destinationTokenAddress: token?.addresses[toChainId] as Address,
            destinationToken: token as Token,
            amount: stringToBigInt(
              hasFromValue ? fromValue : '1',
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
  }, [fromChainId, toChainId, fromToken, fromValue, toTokens, synapseSDK])

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

export const getDefaultBridgeAmount = ({
  originToken,
}: {
  originToken: Token
}): DefaultBridgeAmount => {
  const swappableType: string = originToken.swapableType

  switch (swappableType) {
    case SwappableTypes.STABLE:
      return DefaultBridgeAmount.STABLE
    case SwappableTypes.ETH:
      return DefaultBridgeAmount.ETH
    case SwappableTypes.BTC:
      return DefaultBridgeAmount.BTC
  }
}
