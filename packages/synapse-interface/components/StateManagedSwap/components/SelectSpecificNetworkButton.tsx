import { useEffect, useRef } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'
import { getHoverStyleForButton, getActiveStyleForButton } from '@/styles/hover'
import { joinClassNames } from '@/utils/joinClassNames'

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

  const buttonClass = joinClassNames({
    other: 'whitespace-nowrap',
    grid: 'grid gap-0.5',
    space: 'pl-2 pr-1.5 py-2.5 w-full',
    border: 'border border-transparent',
    transition: 'transition-all duration-75',
    hover: getHoverStyleForButton(chain.color),
    activeStyle: isCurrentChain
      ? getActiveStyleForButton(isCurrentChain && chain.color)
      : '',
  })

  return (
    <button
      ref={ref}
      tabIndex={active ? 1 : 0}
      className={buttonClass}
      onClick={onClick}
      data-test-id={`${dataId}-item`}
    >
      <ButtonContent chainId={itemChainId} />
    </button>
  )
}

function ButtonContent({ chainId }: { chainId: number }) {
  const chain = CHAINS_BY_ID[chainId]

  return chain ? (
    <span className="flex items-center gap-2">
      <Image
        loading="lazy"
        src={chain.chainImg}
        alt="Switch Network"
        width="20"
        height="20"
        className="w-5 h-5 max-w-fit"
      />
      {chain.name}
    </span>
  ) : null
}
