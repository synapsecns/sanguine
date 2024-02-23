import { Token } from '@/utils/types'
import { useSwapState } from '@/slices/swap/hooks'

import { SelectTokenButton } from '@/components/bridgeSwap/SelectTokenButton'

export const SelectSpecificTokenButton = ({
  showAllChains,
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
  alternateBackground = false,
}: {
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
  alternateBackground?: boolean
}) => {
  const { swapChainId } = useSwapState()


  return (
    <SelectTokenButton
      showAllChains={showAllChains}
      token={token}
      active={active}
      selectedToken={selectedToken}
      chainId={swapChainId}
      isOrigin={isOrigin}
      onClick={onClick}
      alternateBackground={alternateBackground}
    />
  )
}




