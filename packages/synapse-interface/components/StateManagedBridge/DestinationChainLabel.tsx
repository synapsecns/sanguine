import { createRef, useRef } from 'react'
import { ChevronDownIcon } from '@heroicons/react/outline'
import { CHAINS_BY_ID } from '@constants/chains'
import { getNetworkButtonBorder } from '@/styles/chains'
import Image from 'next/image'
import Tooltip from '@tw/Tooltip'
import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setToChainId } from '@/slices/bridge/reducer'
import { setShowToChainSlideOver } from '@/slices/bridgeDisplaySlice'
import { getOrderedChains } from '@/utils/getOrderedChains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

const scrollableRef = createRef<HTMLDivElement>()

export const DestinationChainLabel = ({
  chains,
  chainId,
  connectedChainId,
}: {
  chains: number[] | undefined
  chainId: number
  connectedChainId: number
}) => {
  const [orderedChains, setOrderedChains] = useState<number[]>([])

  const dispatch = useDispatch()

  useEffect(() => {
    setOrderedChains(
      reorderArrayWithChainIdFirst(
        chainId,
        chainOrderBySwapSide(connectedChainId, chainId, chains)
      )
    )
  }, [chainId, connectedChainId, chains])

  const resetScrollPosition = () => {
    if (scrollableRef.current) {
      scrollableRef.current.scroll({ left: 0, behavior: 'smooth' })
    }
  }

  return (
    <div data-test-id="destination-chain-label" className="flex items-center">
      <div className={`text-gray-400 block text-sm mr-2 min-w-[40px]`}>
        Dest.
      </div>
      <div className="relative flex w-full">
        <div className="relative flex items-center overflow-x-auto overflow-y-hidden scrollbar-hide md:ml-auto">
          <div
            className={`
            flex items-center
            [&>*:nth-child(2)]:hidden [&>*:nth-child(2)]:min-[360px]:flex
            [&>*:nth-child(3)]:hidden [&>*:nth-child(3)]:min-[390px]:flex
            [&>*:nth-child(4)]:hidden [&>*:nth-child(4)]:min-[420px]:flex
            [&>*:nth-child(5)]:hidden [&>*:nth-child(5)]:min-[450px]:flex
            `}
          >
            {orderedChains.map((id: number, key: number) => {
              return Number(id) === chainId ? (
                <SelectedChain chainId={id} key={id} />
              ) : (
                <PossibleChain
                  chainId={Number(id)}
                  key={id}
                  resetScrollPosition={resetScrollPosition}
                />
              )
            })}
          </div>
          <div className="ml-0 sticky min-w-[15px] h-full right-[-3px] max-[475px]:bg-gradient-to-r from-transparent to-bgLight">
            &nbsp;
          </div>
        </div>

        <div className="max-[475px]:pl-1 ml-auto md:ml-0">
          <button
            onClick={() => {
              dispatch(setShowToChainSlideOver(true))
            }}
            tabIndex={0}
            data-test-id="bridge-destination-chain-list-button"
            className="w-8 h-8 px-1.5 py-1.5 bg-[#C4C4C4] bg-opacity-10 rounded-full hover:cursor-pointer group"
          >
            <ChevronDownIcon className="text-gray-300 transition transform-gpu group-hover:opacity-50 group-active:rotate-180" />
          </button>
        </div>
      </div>
    </div>
  )
}

const PossibleChain = ({
  chainId,
  resetScrollPosition,
}: {
  chainId: number
  resetScrollPosition: () => void
}) => {
  const chain = CHAINS_BY_ID[chainId]
  const { toChainId } = useBridgeState()

  const dispatch = useDispatch()

  const onChangeChain = () => {
    const eventTitle = `[Bridge User Action] Change Origin Chain`
    const eventData = {
      previousToChain: toChainId,
      newToChainId: chainId,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setToChainId(chainId))
    resetScrollPosition()
  }
  return chain ? (
    <button
      data-test-id="destination-possible-chain"
      className="
        min-w-[1.75rem] min-h-[1.75rem]
        max-w-[1.75rem] max-h-[1.75rem]
        px-0.5 py-0.5
        border border-gray-500 rounded-full
        mr-3
      "
      tabIndex={0}
      onClick={onChangeChain}
    >
      <Tooltip content={chain.name}>
        <Image
          src={chain.chainImg}
          className="duration-300 rounded-full hover:scale-125"
          alt={chain.name}
        />
      </Tooltip>
    </button>
  ) : null
}

const SelectedChain = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return chain ? (
    <div
      ref={scrollableRef}
      data-test-id="destination-selected-chain"
      className={`
        px-1 mr-3
        flex items-center
        bg-bgLight
        text-white
        border ${getNetworkButtonBorder(chain.color)}
        rounded-full
        min-w-fit
      `}
    >
      <Image
        alt="chain image"
        src={chain.chainImg}
        className="w-5 h-5 my-1 mr-1 rounded-full opacity-80"
      />
      <div className="flex">
        <div className="mr-2 text-sm text-white">
          {chain.name === 'Boba Network' ? 'Boba' : chain.name}
        </div>
      </div>
    </div>
  ) : null
}

const chainOrderBySwapSide = (
  connectedChainId: number,
  chainId: number,
  chains
): number[] => {
  return getOrderedChains(
    connectedChainId,
    chainId,
    chains?.map((id) => `${id}`)
  )
}

const reorderArrayWithChainIdFirst = (
  chainId: number,
  orderedChains: number[]
): number[] => {
  const chainIndex = orderedChains.indexOf(chainId)

  if (chainIndex !== -1) {
    orderedChains.splice(chainIndex, 1)
  }

  orderedChains.unshift(chainId)

  return orderedChains
}
