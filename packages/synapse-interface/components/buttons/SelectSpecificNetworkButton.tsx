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
} from '@/styles/chains'

// TODO could probably use a re-write
export const SelectSpecificNetworkButton = ({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
  dataId,
}: {
  itemChainId: number
  isCurrentChain: boolean
  active: boolean
  onClick: () => void
  dataId: string
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
      bg-opacity-50
    `
  } else {
    bgClassName = 'bg-[#58535B] hover:bg-[#58535B] active:bg-[#58535B]'
  }

  return (
    <button
      ref={ref}
      tabIndex={active ? 1 : 0}
      className={`
        flex items-center
        transition-all duration-75
        w-full rounded-md
        px-2 py-3
        cursor-pointer
        border border-transparent
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
        className="w-10 h-10 ml-2 mr-4 rounded-full"
      />
      <div className="flex-col text-left">
        <div className="text-lg font-medium text-white">{chain.name}</div>
        <div className="text-sm text-white opacity-50">Layer {chain.layer}</div>
      </div>
    </>
  ) : null
}
