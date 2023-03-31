import { ChevronDownIcon } from '@heroicons/react/outline'
import { CHAIN_INFO_MAP, CHAIN_ID_DISPLAY_ORDER } from '@constants/networks'
import { getNetworkButtonBorder } from '@styles/networks'
import { getOrderedChains } from '@utils/getOrderedChains'
import Image from 'next/image'
import Tooltip from '@tw/Tooltip'
import { useNetwork } from 'wagmi'
import { use, useEffect, useState } from 'react'

export const ChainLabel = ({
  isOrigin,
  chains,
  chainId,
  titleText,
  connectedChainId,
  labelClassNameOverride,
  onChangeChain,
  setDisplayType,
}: {
  isOrigin: boolean
  chains: string[] | undefined
  chainId: number
  titleText?: string
  connectedChainId: number
  labelClassNameOverride?: string
  onChangeChain: (v: number) => void
  setDisplayType: (v: string) => void
}) => {
  const displayType = isOrigin ? 'fromChain' : 'toChain'
  const title = titleText ?? (isOrigin ? 'Origin' : 'Dest.')
  const labelClassName = 'text-sm'
  const [orderedChains, setOrderedChains] = useState<number[]>([])
  useEffect(() => {
    setOrderedChains(
      chainOrderBySwapSide(connectedChainId, isOrigin, chainId, chains)
    )
  }, [chainId, connectedChainId, chains])

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

const PossibleChain = ({
  chainId,
  onChangeChain,
}: {
  chainId: number
  onChangeChain: (v: number) => void
}) => {
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

const SelectedChain = ({ chainId }: { chainId: number }) => {
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

const chainOrderBySwapSide = (
  connectedChain: number,
  isOrigin: boolean,
  chainId: number,
  chains: string[] | undefined
) => {
  let orderedChains
  if (isOrigin) {
    orderedChains = CHAIN_ID_DISPLAY_ORDER.filter((e) => e !== chainId)
    orderedChains = orderedChains.slice(0, 5)
    orderedChains.unshift(chainId)
    return orderedChains
  } else {
    return getOrderedChains(connectedChain, chainId, chains)
  }
}
