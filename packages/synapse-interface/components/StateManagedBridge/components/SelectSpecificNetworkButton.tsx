import { isChainEligible } from '@/utils/hooks/useStipEligibility'
import { useBridgeState } from '@/slices/bridge/hooks'
import { SelectNetworkButton } from '@/components/bridgeSwap/SelectNetworkButton'

export const SelectSpecificNetworkButton = ({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
  dataId,
  isOrigin,
  alternateBackground = false,
}: {
  itemChainId: number
  isCurrentChain: boolean
  active: boolean
  onClick: () => void
  dataId: string
  isOrigin: boolean
  alternateBackground?: boolean
}) => {
  const { fromChainId, fromToken } = useBridgeState()
  const isEligible = isChainEligible(fromChainId, itemChainId, fromToken)

  return (
    <SelectNetworkButton
      itemChainId={itemChainId}
      isCurrentChain={isCurrentChain}
      active={active}
      onClick={onClick}
      dataId={dataId}
      isOrigin={isOrigin}
      isEligible={isEligible}
      alternateBackground={alternateBackground}
    />
  )
}

