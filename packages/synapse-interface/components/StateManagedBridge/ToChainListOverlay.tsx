import _ from 'lodash'

import { useBridgeState } from '@/slices/bridge/hooks'
import { setToChainId } from '@/slices/bridge/reducer'
import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'

import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { ChainListOverlay } from '../bridgeSwap/ChainListOverlay'

export const ToChainListOverlay = () => {
  const { toChainIds, toChainId } = useBridgeState()
  const dataId = 'bridge-destination-chain-list'


  return (
    <ChainListOverlay
      isOrigin={false}
      primaryLabel={"To"}
      chainId={toChainId}
      chainIds={toChainIds}
      setChainId={setToChainId}
      setShowOverlay={setShowToChainListOverlay}
      SelectNetworkButtonComponent={SelectSpecificNetworkButton}
    />
  )
}
