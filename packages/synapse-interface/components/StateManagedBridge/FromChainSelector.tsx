import { setFromChainId } from '@/slices/bridge/reducer'
import { ChainSelector } from '@/components/ui/ChainSelector'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useFromChainListArray } from './hooks/useFromChainListArray'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useWalletState } from '@/slices/wallet/hooks'

export const FromChainSelector = () => {
  const { fromChainId } = useBridgeState()
  const { isWalletPending } = useWalletState()

  return (
    <ChainSelector
      dataTestId="bridge-origin-chain"
      selectedItem={CHAINS_BY_ID[fromChainId]}
      isOrigin={true}
      label="From"
      itemListFunction={useFromChainListArray}
      setFunction={setFromChainId}
      action="Bridge"
      disabled={isWalletPending}
    />
  )
}
