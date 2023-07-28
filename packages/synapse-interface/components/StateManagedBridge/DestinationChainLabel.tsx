import { createRef, useRef } from 'react'
import { ChevronDownIcon } from '@heroicons/react/outline'
import { CHAINS_BY_ID } from '@constants/chains'
import { getNetworkButtonBorder } from '@/styles/chains'
import Image from 'next/image'
import Tooltip from '@tw/Tooltip'
import { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { setToChainId } from '@/slices/bridgeSlice'
import { setShowToChainSlideOver } from '@/slices/bridgeDisplaySlice'
import { getOrderedChains } from '@/utils/getOrderedChains'
import { RootState } from '@/store/store'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import useWindowSize from '@/utils/hooks/useWindowSize'

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
  const leftRef = useRef<HTMLDivElement>(null)
  const rightRef = useRef<HTMLDivElement>(null)
  const [orderedChains, setOrderedChains] = useState<number[]>([])
  const [chainListWidth, setChainListWidth] = useState<number>(0)
  const [width, height] = useWindowSize()

  const dispatch = useDispatch()

  useEffect(() => {
    setOrderedChains(chainOrderBySwapSide(connectedChainId, chainId, chains))
  }, [chainId, connectedChainId, chains])

  const resetScrollPosition = () => {
    if (scrollableRef.current) {
      scrollableRef.current.scroll({ left: 0, behavior: 'smooth' })
    }
  }

  useEffect(() => {
    if (width && leftRef.current && rightRef.current) {
      const distance = getDistanceBetweenElements(
        leftRef.current,
        rightRef.current
      )
      setChainListWidth(distance)
    }
  }, [width])

  return (
    <div
      data-test-id="destination-chain-label"
      className="flex items-center justify-between"
    >
      <div ref={leftRef} className={`text-gray-400 block text-sm mr-2`}>
        Dest.
      </div>
      <div className="relative flex ml-auto">
        <div className="flex items-center relative overflow-x-auto overflow-y-hidden w-[200px] min-[400px]:w-[260px] min-[475px]:w-full scrollbar-hide">
          <div className="hidden sticky min-w-[15px] h-full left-[-3px] bg-gradient-to-l from-transparent to-bgLight max-[475px]:block">
            &nbsp;
          </div>
          <div className="flex items-center [&>*:last-child]:mr-0">
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
          <div className="ml-0 sticky min-w-[15px] h-full right-[-3px] bg-gradient-to-r from-transparent to-bgLight">
            &nbsp;
          </div>
        </div>

        <div ref={rightRef} className="max-[475px]:pl-1">
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
  const { toChainId } = useSelector((state: RootState) => state.bridge)

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
      <div className="inline-block">
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
    chains.map((id) => `${id}`)
  )
}

export function getDistanceBetweenElements(
  element1: HTMLElement,
  element2: HTMLElement
): number {
  const rect1 = element1.getBoundingClientRect()
  const rect2 = element2.getBoundingClientRect()

  const distance = Math.abs(rect2.left - rect1.right)

  return distance
}
