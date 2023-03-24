import { ChevronDownIcon } from '@heroicons/react/outline'
import { CHAIN_INFO_MAP, CHAIN_ID_DISPLAY_ORDER } from '@constants/networks'
import { getNetworkButtonBorder } from '@styles/networks'
import { getOrderedChains } from '@utils/getOrderedChains'
import Image from 'next/image'
import Tooltip from '@tw/Tooltip'
import { useNetwork } from 'wagmi'
import { use, useEffect, useState } from 'react'

export function ChainLabel({
  isSwapFrom,
  chainId,
  setDisplayType,
  labelClassNameOverride,
  titleText,
  onChangeChain,
  possibleChains,
  connectedChainId,
}: {
  isSwapFrom: boolean
  chainId: number
  setDisplayType: (v: string) => void
  labelClassNameOverride?: string
  titleText?: string
  onChangeChain: (v: number) => void
  possibleChains: string[] | undefined
  connectedChainId: number
}) {
  const [orderedChains, setOrderedChains] = useState<number[]>([])
  useEffect(() => {
    setOrderedChains(
      chainOrderBySwapSide(
        connectedChainId,
        isSwapFrom,
        chainId,
        possibleChains
      )
    )
  }, [chainId, connectedChainId, possibleChains])

  let displayType: string
  let title: string
  let labelClassName

  if (isSwapFrom) {
    title = titleText ?? 'Origin'
    displayType = 'fromChain'
    labelClassName = 'text-sm'
  } else {
    title = titleText ?? 'Dest.'
    displayType = 'toChain'
    labelClassName = 'text-sm'
  }

  return (
    <div className="flex items-center justify-center md:justify-between">
      <div
        className={`text-gray-400 ${labelClassName} ${labelClassNameOverride} hidden md:block lg:block text-sm mr-2`}
      >
        {title}
      </div>
      <div className="flex items-center space-x-4 md:space-x-3">
        {orderedChains.map((id) =>
          id === chainId ? (
            <SelectedChain chainId={id} key={id} />
          ) : (
            <PossibleChain
              chainId={id}
              onChangeChain={onChangeChain}
              key={id}
            />
          )
        )}
        <button
          onClick={() => {
            setDisplayType(displayType)
          }}
          tabIndex={0}
          className="w-8 h-8 px-1.5 py-1.5 bg-[#C4C4C4] bg-opacity-10 rounded-full hover:cursor-pointer group"
        >
          <ChevronDownIcon className="text-gray-300 transition transform-gpu group-hover:opacity-50 group-active:rotate-180" />
        </button>
      </div>
    </div>
  )
}

function PossibleChain({
  chainId,
  onChangeChain,
}: {
  chainId: number
  onChangeChain: (v: number) => void
}) {
  const { chainImg, chainName } = CHAIN_INFO_MAP[chainId]
  const onClick = () => {
    onChangeChain(chainId)
  }
  return (
    <button
      className="
        w-7 h-7
        md:w-7
        px-0.5 py-0.5
        border border-gray-500 rounded-full
      "
      tabIndex={0}
      onClick={onClick}
    >
      <Tooltip content={chainName}>
        <Image
          src={chainImg}
          className="duration-300 rounded-full hover:scale-125"
          alt={chainName}
        />
      </Tooltip>
    </button>
  )
}

function SelectedChain({ chainId }: { chainId: number }) {
  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId]
  return (
    <div
      className={`
        px-1
        flex items-center
        bg-bgLight
        text-white
        border ${getNetworkButtonBorder(chainId)}
        rounded-full
      `}
    >
      <Image
        alt="chain image"
        src={chainImg}
        className="w-5 h-5 my-1 mr-0 rounded-full md:mr-1 opacity-80"
      />
      <div className="hidden md:inline-block lg:inline-block">
        <div className="mr-2 text-sm text-white">
          {chainName === 'Boba Network' ? 'Boba' : chainName}
        </div>
      </div>
    </div>
  )
}

function chainOrderBySwapSide(
  connectedChain: number,
  isSwapFrom: boolean,
  chainId: number,
  possibleChains: string[] | undefined
) {
  let orderedChains
  if (isSwapFrom) {
    orderedChains = CHAIN_ID_DISPLAY_ORDER.filter((e) => e !== chainId)
    orderedChains = orderedChains.slice(0, 5)
    orderedChains.unshift(chainId)
    console.log('YOUOO', chainId, orderedChains)

    return orderedChains
  } else {
    let h = getOrderedChains(connectedChain, chainId, possibleChains)
    console.log('YwwOUOO', chainId, possibleChains, h)

    return h
  }
}
