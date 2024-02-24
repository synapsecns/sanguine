import _ from 'lodash'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { ChainListOverlay } from '@components/bridgeSwap/ChainListOverlay'

export const FromChainListOverlay = () => {
  const { fromChainIds, fromChainId } = useBridgeState()
  const dataId = 'bridge-origin-chain-list'

  return (
    <ChainListOverlay
      isOrigin={true}
      primaryLabel={"From"}
      chainId={fromChainId}
      chainIds={fromChainIds}
      setChainId={setFromChainId}
      setShowOverlay={setShowFromChainListOverlay}
      SelectNetworkButtonComponent={SelectSpecificNetworkButton}
    />
  )
}
