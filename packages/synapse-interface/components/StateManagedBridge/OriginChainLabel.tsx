import { useRef } from 'react'
import { useSelector } from 'react-redux'
import { ChevronDownIcon } from '@heroicons/react/outline'
import { CHAINS_BY_ID, ORDERED_CHAINS_BY_ID } from '@constants/chains'
import { getNetworkButtonBorder } from '@/styles/chains'
import Image from 'next/image'
import Tooltip from '@tw/Tooltip'
import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { setFromChainId } from '@/slices/bridgeSlice'
import { setShowFromChainSlideOver } from '@/slices/bridgeDisplaySlice'
import { RootState } from '@/store/store'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import useElementWidth from '@/utils/hooks/useElementWidth'

export const OriginChainLabel = ({
  chains,
  chainId,
  connectedChainId,
}: {
  chains: number[] | undefined
  chainId: number
  connectedChainId: number
}) => {
  const chainContainerRef = useRef<HTMLDivElement>(null)
  const leftRef = useRef<HTMLDivElement>(null)
  const rightRef = useRef<HTMLDivElement>(null)
  const scrollableRef = useRef<HTMLDivElement>(null)

  // const containerLength = useElementWidth(chainContainerRef)
  // const leftLength = useElementWidth(leftRef)
  // const rightLength = useElementWidth(rightRef)

  const [orderedChains, setOrderedChains] = useState<number[]>([])

  useEffect(() => {
    setOrderedChains(chainOrderBySwapSide(chainId))
  }, [chainId, connectedChainId, chains])

  const dispatch = useDispatch()

  const resetScrollPosition = () => {
    if (scrollableRef.current) {
      scrollableRef.current.scrollLeft = 0
    }
  }

  return (
    <div
      ref={chainContainerRef}
      data-test-id="origin-chain-label"
      className="flex items-center"
    >
      <div
        ref={leftRef}
        className={`text-gray-400 block text-sm mr-2 min-w-[40px]`}
      >
        Origin
      </div>
      <div className="relative flex w-full">
        <div
          ref={scrollableRef}
          className={`
            flex items-center relative
            overflow-x-auto overflow-y-hidden
            scrollbar-hide`}
          // style={{
          //   width: `${containerLength - leftLength - rightLength - 20}px`,
          // }}
        >
          {/* <div className="block sticky min-w-[5px] h-full left-[-2px] max-[475px]:bg-gradient-to-l from-transparent to-bgLight ">
            &nbsp;
          </div> */}
          <div
            className={`
            flex items-center
            [&>*:nth-child(2)]:hidden [&>*:nth-child(2)]:min-[360px]:flex
            [&>*:nth-child(3)]:hidden [&>*:nth-child(3)]:min-[375px]:flex
            [&>*:nth-child(4)]:hidden [&>*:nth-child(4)]:min-[414px]:flex
            [&>*:nth-child(5)]:hidden [&>*:nth-child(5)]:min-[450px]:flex
            `}
          >
            {orderedChains.map((id: number, key: number) => {
              return Number(id) === chainId ? (
                <SelectedChain chainId={Number(id)} key={id} />
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

        <div ref={rightRef} className="max-[475px]:pl-1 ml-auto">
          <button
            onClick={() => {
              dispatch(setShowFromChainSlideOver(true))
            }}
            tabIndex={0}
            data-test-id="bridge-origin-chain-list-button"
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
  const { fromChainId } = useSelector((state: RootState) => state.bridge)

  const dispatch = useDispatch()

  const onChangeChain = () => {
    const eventTitle = `[Bridge User Action] Change Origin Chain`
    const eventData = {
      previousFromChainId: fromChainId,
      newFromChainId: chainId,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setFromChainId(chainId))
    resetScrollPosition()
  }

  return chain ? (
    <button
      data-test-id="origin-possible-chain"
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
      data-test-id="origin-selected-chain"
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

const chainOrderBySwapSide = (chainId: number, count?: number) => {
  let orderedChains
  orderedChains = ORDERED_CHAINS_BY_ID.filter((e) => e !== String(chainId))
  orderedChains = orderedChains.slice(0, count ?? 5)
  orderedChains.unshift(chainId)

  return orderedChains
}
