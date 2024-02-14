import { useDispatch } from 'react-redux'

import { useBridgeState } from '@/slices/bridge/hooks'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'

import { GenericOutputContainer } from '@/components/bridgeSwap/GenericOutputContainer'
import { ChainSelector } from '@/components/bridgeSwap/ChainSelector'
import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'

export const OutputContainer = () => {
  const dispatch = useDispatch()
  const { bridgeQuote, isLoading, toChainId, toToken } = useBridgeState()

  return (
    <GenericOutputContainer
      chainSelector={
        <ChainSelector
          chainId={toChainId}
          label="To"
          onClick={() => dispatch(setShowToChainListOverlay(true))}
        />
      }
      tokenSelector={
        <TokenSelector
          data-test-id="bridge-destination-token"
          token={toToken}
          label="Out"
          onClick={() => dispatch(setShowToTokenListOverlay(true))}
        />
      }
      isLoading={isLoading}
      quote={bridgeQuote}
    />
  )
}

