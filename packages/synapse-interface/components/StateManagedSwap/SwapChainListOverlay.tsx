import _ from 'lodash'
import { useCallback, useEffect, useRef, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import * as ALL_CHAINS from '@constants/chains/master'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { useDispatch } from 'react-redux'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
// import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { SelectSpecificNetworkButton } from '../StateManagedBridge/components/SelectSpecificNetworkButton'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from './components/CloseButton'
import { SearchResults } from './components/SearchResults'
import { setShowSwapChainListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapChainId } from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'

export const SwapChainListOverlay = () => {
  const { swapChainId, swapFromChainIds } = useSwapState()
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  const dataId = 'swap-origin-chain-list'
  const overlayRef = useRef(null)

  let possibleChains = sortChains(
    _(ALL_CHAINS)
      .pickBy((value) => _.includes(swapFromChainIds, value.id))
      .values()
      .value()
  )

  let remainingChains = swapFromChainIds
    ? sortChains(
        _.difference(
          Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
          swapFromChainIds.map((id) => CHAINS_BY_ID[id])
        )
      )
    : []

  const possibleChainsWithSource = possibleChains.map((chain) => ({
    ...chain,
    source: 'possibleChains',
  }))

  const remainingChainsWithSource = remainingChains.map((chain) => ({
    ...chain,
    source: 'remainingChains',
  }))

  const masterList = [...possibleChainsWithSource, ...remainingChainsWithSource]

  const fuseOptions = {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'id',
      'nativeCurrency.symbol',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter((item) => item.source === 'possibleChains')
    remainingChains = results.filter(
      (item) => item.source === 'remainingChains'
    )
  }

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setSearchStr('')
    setOpen(false)
    dispatch(setShowSwapChainListOverlay(false))
  }, [dispatch])

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(escFunc, [escPressed])
  useEffect(arrowUpFunc, [arrowUp])
  useCloseOnOutsideClick(overlayRef, onClose)

  const handleSetSwapChainId = (chainId) => {
    const eventTitle = `[Swap User Action] Sets new fromChainId`
    const eventData = {
      previousFromChainId: swapChainId,
      newFromChainId: chainId,
    }

    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setSwapChainId(chainId))
    onClose()
  }

  const [open, setOpen] = useState(true)

  useEffect(() => {
    const ref = overlayRef.current
    const { y, height } = ref.getBoundingClientRect()
    const screen = window.innerHeight
    console.log(ref.style)
    if (y + height > screen) {
      ref.style.position = 'fixed'
      ref.style.bottom = '4px'
    }
    if (y < 0) {
      ref.style.position = 'fixed'
      ref.style.top = '4px'
    }
  }, [])

  return (
    <div
      ref={overlayRef}
      data-test-id="swapChain-list-overlay"
      className={`${
        open ? 'block' : 'hidden'
      } z-20 absolute bg-bgLight border border-separator rounded overflow-y-auto max-h-96 animate-slide-down origin-top shadow-md`}
    >
      <div className="p-1 flex items-center font-medium">
        <SlideSearchBox
          placeholder="Find"
          searchStr={searchStr}
          onSearch={onSearch}
        />
        <CloseButton onClick={onClose} />
      </div>
      <div data-test-id={dataId}>
        {
          possibleChains && possibleChains.length > 0 && (
            <>
              <div className="p-2 text-sm text-secondary">From…</div>
              {possibleChains.map(({ id: mapChainId }, idx) => {
                return (
                  <SelectSpecificNetworkButton
                    key={idx}
                    itemChainId={mapChainId}
                    isCurrentChain={swapChainId === mapChainId}
                    active={idx === currentIdx}
                    onClick={() =>
                      swapChainId === mapChainId
                        ? onClose()
                        : handleSetSwapChainId(mapChainId)
                    }
                    dataId={dataId}
                    isOrigin={true}
                  />
                )
              })}
            </>
          )
        }
        {
          remainingChains && remainingChains.length > 0 && (
            <div className="bg-bgBase rounded">
              <div className="px-2 py-2 text-sm text-secondary">All chains</div>
              {remainingChains.map(({ id: mapChainId }, idx) => {
                return (
                  <SelectSpecificNetworkButton
                    key={mapChainId}
                    itemChainId={mapChainId}
                    isCurrentChain={swapChainId === mapChainId}
                    active={idx + possibleChains.length === currentIdx}
                    onClick={() =>
                      swapChainId === mapChainId
                        ? onClose()
                        : handleSetSwapChainId(mapChainId)
                    }
                    dataId={dataId}
                    alternateBackground={true}
                    isOrigin={true}
                  />
                )
              })}
            </div>
          )
        }
        <SearchResults searchStr={searchStr} type="chain" />
      </div>
    </div>
  )
}
