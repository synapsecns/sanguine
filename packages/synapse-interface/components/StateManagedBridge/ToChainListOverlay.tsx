import { useCallback, useEffect, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import * as CHAINS from '@constants/chains/master'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { useNetwork } from 'wagmi'
import { sortChains } from '@constants/chains'
import { useDispatch } from 'react-redux'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setToChainId } from '@/slices/bridge/reducer'
import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'

export const ToChainListOverlay = () => {
  const { toChainIds, toChainId } = useBridgeState()
  const { chain } = useNetwork()
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const [networks, setNetworks] = useState([])
  const dispatch = useDispatch()
  const fuse = new Fuse(networks, {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'chainShortName',
      'chainId',
      'nativeCurrency',
    ],
  })

  const dataId = 'bridge-destination-chain-list'

  useEffect(() => {
    let tempNetworks = []
    Object.values(CHAINS).map((chain) => {
      if (toChainIds?.includes(chain.id)) {
        tempNetworks.push(chain)
      }
    })
    tempNetworks = sortChains(tempNetworks)
    if (searchStr?.length > 0) {
      tempNetworks = fuse.search(searchStr).map((i) => i.item)
    }
    setNetworks(tempNetworks)
  }, [chain, searchStr])

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowToChainListOverlay(false))
  }, [setShowToChainListOverlay])

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < networks.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  const enterPressedFunc = () => {
    if (enterPressed && currentIdx > -1) {
      const currentChain = networks[currentIdx]
      dispatch(setToChainId(currentChain.chainId))
      onClose()
    }
  }
  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(escFunc, [escPressed])
  useEffect(arrowUpFunc, [arrowUp])
  useEffect(enterPressedFunc, [enterPressed])

  return (
    <div
      data-test-id="chain-slide-over"
      className="max-h-full pb-4 mt-2 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="flex items-center mb-2 font-medium justfiy-between sm:float-none">
          <SlideSearchBox
            placeholder="Filter"
            searchStr={searchStr}
            onSearch={onSearch}
          />
        </div>
      </div>
      <div
        data-test-id={dataId}
        className="px-2 pt-2 pb-8 bg-[#343036] md:px-2"
      >
        <div className="mb-2 text-sm font-normal text-white">Receive on...</div>
        {networks.map(({ id: mapChainId }, idx) => {
          let onClickSpecificNetwork
          if (toChainId === mapChainId) {
            onClickSpecificNetwork = () => {
              onClose()
            }
          } else {
            onClickSpecificNetwork = () => {
              const eventTitle = `[Bridge User Action] Sets new toChainId`
              const eventData = {
                previousToChainId: toChainId,
                newToChainId: mapChainId,
              }

              segmentAnalyticsEvent(eventTitle, eventData)
              dispatch(setToChainId(mapChainId))
              onClose()
            }
          }
          return (
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={toChainId === mapChainId}
              active={idx === currentIdx}
              onClick={onClickSpecificNetwork}
              dataId={dataId}
            />
          )
        })}
        {searchStr && (
          <div className="px-12 py-4 text-center text-white text-md">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-2 text-white text-opacity-50 align-bottom text-md">
              Want to see a chain supported on Synapse? Let us know!
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
