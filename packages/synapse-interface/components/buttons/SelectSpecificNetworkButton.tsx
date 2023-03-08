import { useRef, useEffect } from 'react'
import { CHAIN_INFO_MAP } from '@constants/networks'
import Image from 'next/image'
import {
  getNetworkButtonBorderHover,
  getNetworkHover,
  getNetworkButtonBgClassName,
  getNetworkButtonBorder,
} from '@styles/networks'

export function SelectSpecificNetworkButton({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
}: {
  itemChainId: number
  isCurrentChain: boolean
  active: boolean
  onClick: () => void
}) {
  const ref = useRef<any>(null)

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  let bgClassName

  if (isCurrentChain) {
    bgClassName = `
      ${getNetworkButtonBgClassName(itemChainId)}
      ${getNetworkButtonBorder(itemChainId)}
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
        w-full rounded-xl
        px-2 py-3
        cursor-pointer
        border border-transparent
        ${getNetworkHover(itemChainId)}
        ${getNetworkButtonBorderHover(itemChainId)}
        ${bgClassName}
      `}
      onClick={onClick}
    >
      <ButtonContent chainId={itemChainId} />
    </button>
  )
}

function ButtonContent({ chainId }: { chainId: number }) {
  const { chainImg, chainName, layer } = CHAIN_INFO_MAP[chainId]

  return (
    <>
      <Image
        src={chainImg}
        alt="Switch Network"
        className="w-10 h-10 ml-2 mr-4 rounded-full"
      />
      <div className="flex-col text-left">
        <div className="text-lg font-medium text-white">{chainName}</div>
        <div className="text-sm text-white opacity-50">Layer {layer}</div>
      </div>
    </>
  )
}
