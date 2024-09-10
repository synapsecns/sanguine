import { setFromToken } from '@/slices/bridge/reducer'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useFromTokenListArray } from './hooks/useFromTokenListArray'
import { useWalletState } from '@/slices/wallet/hooks'

export const FromTokenSelector = () => {
  const { fromToken } = useBridgeState()
  const { isWalletPending } = useWalletState()

  return (
    <TokenSelector
      dataTestId="bridge-origin-token"
      selectedItem={fromToken}
      isOrigin={true}
      placeholder="Out"
      itemListFunction={useFromTokenListArray}
      setFunction={setFromToken}
      action="Bridge"
      disabled={isWalletPending}
    />
  )
}
