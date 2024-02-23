
import { SelectNetworkButton } from '@/components/bridgeSwap/SelectNetworkButton'

export const SelectSpecificNetworkButton = ({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
  dataId,
  alternateBackground = false,
}: {
  itemChainId: number
  isCurrentChain: boolean
  active: boolean
  onClick: () => void
  dataId: string
  alternateBackground?: boolean
}) => {

  return (
    <SelectNetworkButton
      itemChainId={itemChainId}
      isCurrentChain={isCurrentChain}
      active={active}
      onClick={onClick}
      dataId={dataId}
      isOrigin={true}
      isEligible={false}
      alternateBackground={alternateBackground}
    />
  )
}

