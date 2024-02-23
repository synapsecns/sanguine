import { useEffect, useRef } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'
import {
  getNetworkHover,
  getNetworkButtonBorder,
  getNetworkButtonBorderHover,
  getNetworkButtonBgClassName,
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getMenuItemStyleForChain,
} from '@/styles/chains'
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
  const ref = useRef<any>(null)
  const chain = CHAINS_BY_ID[itemChainId]

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  let bgClassName

  if (isCurrentChain) {
    bgClassName = `
      ${getNetworkButtonBgClassName(chain.color)}
      ${getNetworkButtonBorder(chain.color)}
      bg-opacity-30
    `
  }

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

