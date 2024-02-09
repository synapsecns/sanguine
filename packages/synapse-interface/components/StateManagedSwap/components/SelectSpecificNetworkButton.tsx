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
    <button
      ref={ref}
      tabIndex={active ? 1 : 0}
      className={`
        flex items-center
        transition-all duration-75
        w-full
        px-2 py-4
        cursor-pointer
        rounded-md
        border border-slate-400/10
        mb-1
        ${alternateBackground ? '' : !isCurrentChain && 'bg-slate-400/10'}
        ${bgClassName}
        ${getNetworkButtonBorderHover(chain.color)}
        ${getNetworkHover(chain.color)}
        ${getNetworkButtonBgClassNameActive(chain.color)}
        ${getNetworkButtonBorderActive(chain.color)}
      `}
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
    <>
      <Image
        src={chain.chainImg}
        alt="Switch Network"
        className="ml-2 mr-4 rounded-full w-7 h-7"
      />
      <div className="flex-col text-left">
        <div className="text-lg font-normal text-white">{chain.name}</div>
      </div>
    </>
  ) : null
}
