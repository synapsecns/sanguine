import _ from 'lodash'
import Fuse from 'fuse.js'
import * as ALL_CHAINS from '@constants/chains/master'
import { SlideSearchBox } from '@components/bridgeSwap/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { useDispatch } from 'react-redux'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { CloseButton } from '@/components/buttons/CloseButton'
import { SearchResults } from '@/components/bridgeSwap/SearchResults'
import { setShowSwapChainListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapChainId } from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { CHAIN_FUSE_OPTIONS } from '@/constants/fuseOptions'
import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'

export const SwapChainListOverlay = () => {
  const { swapChainId, swapFromChainIds } = useSwapState()

  const dispatch = useDispatch()
  const dataId = 'swap-origin-chain-list'

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

  function onCloseOverlay() {
    dispatch(setShowSwapChainListOverlay(false))
  }

  const {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose,
  } = useOverlaySearch(masterList.length, onCloseOverlay)

  const fuse = new Fuse(masterList, CHAIN_FUSE_OPTIONS)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter((item) => item.source === 'possibleChains')
    remainingChains = results.filter(
      (item) => item.source === 'remainingChains'
    )
  }

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

  return (
    <div
      ref={overlayRef}
      data-test-id="fromChain-list-overlay"
      className="max-h-full pb-4 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="relative flex items-center my-2 font-medium">
          <SlideSearchBox
            placeholder="Filter by chain name, id, or native currency"
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <CloseButton onClick={onClose} />
        </div>
      </div>
      <div data-test-id={dataId} className="px-2 pt-2 pb-8 md:px-2">
        {possibleChains && possibleChains.length > 0 && (
          <SearchResultsContainer label="From…">
            {possibleChains.map(({ id: mapChainId }, idx) =>
                <SelectSpecificNetworkButton
                  key={idx}
                  itemChainId={mapChainId}
                  isCurrentChain={swapChainId === mapChainId}
                  active={idx === currentIdx}
                  onClick={() => {
                    if (swapChainId === mapChainId) {
                      onClose()
                    } else {
                      handleSetSwapChainId(mapChainId)
                    }
                  }}
                  dataId={dataId}
                />
            )}
          </SearchResultsContainer>
        )}
        {remainingChains && remainingChains.length > 0 && (
          <SearchResultsContainer label="All Chains">
            {remainingChains.map(({ id: mapChainId }, idx) =>
                <SelectSpecificNetworkButton
                  key={mapChainId}
                  itemChainId={mapChainId}
                  isCurrentChain={swapChainId === mapChainId}
                  active={idx + possibleChains.length === currentIdx}
                  onClick={() => handleSetSwapChainId(mapChainId)}
                  dataId={dataId}
                  alternateBackground={true}
                />
            )}
          </SearchResultsContainer>
        )}
        <SearchResults searchStr={searchStr} type="chain" />
      </div>
    </div>
  )
}
