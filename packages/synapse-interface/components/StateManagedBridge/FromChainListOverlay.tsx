import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'
import * as ALL_CHAINS from '@constants/chains/master'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { CloseButton } from '@/components/buttons/CloseButton'
import { SearchResults } from '@/components/SearchResults'
import { PAUSED_FROM_CHAIN_IDS } from '@constants/chains'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { CHAIN_FUSE_OPTIONS } from '@/constants/fuseOptions'
import { SearchResultsContainer } from '@/components/SearchResultsContainer'
import { useMemo } from 'react'

export const FromChainListOverlay = () => {
  const { fromChainIds, fromChainId } = useBridgeState()
  const dispatch = useDispatch()
  const dataId = 'bridge-origin-chain-list'

  let possibleChains = _(ALL_CHAINS)
    .pickBy((value) => _.includes(fromChainIds, value.id))
    .values()
    .value()
    .filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
      fromChainIds.map((id) => CHAINS_BY_ID[id])
    )
  ).filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

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
    dispatch(setShowFromChainListOverlay(false))
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

    possibleChains = results.filter(item => item.source === 'possibleChains')
    remainingChains = results.filter(item => item.source === 'remainingChains')
  }


  const handleSetFromChainId = (chainId) => {
    const eventTitle = `[Bridge User Action] Sets new fromChainId`
    const eventData = {
      previousFromChainId: fromChainId,
      newFromChainId: chainId,
    }

    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setFromChainId(chainId))
    onClose()
  }

  return (
    <div
      ref={overlayRef}
      data-test-id="fromChain-list-overlay"
      className="max-h-full pb-4 mt-2 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="relative flex items-center mb-2 font-medium">
          <SlideSearchBox
            placeholder="Filter by chain name, id, or native currency"
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <CloseButton onClick={onClose} />
        </div>
      </div>
      <div data-test-id={dataId} className="px-2 pt-2 pb-8 md:px-2">
        {possibleChains?.length > 0 && (
          <SearchResultsContainer label="From…">
            {possibleChains.map(({ id: mapChainId }, idx) =>
                <SelectSpecificNetworkButton
                  key={idx}
                  itemChainId={mapChainId}
                  isCurrentChain={fromChainId === mapChainId}
                  isOrigin={true}
                  active={idx === currentIdx}
                  onClick={() => {
                    if (fromChainId === mapChainId) {
                      onClose()
                    } else {
                      handleSetFromChainId(mapChainId)
                    }
                  }}
                  dataId={dataId}
                />
            )}
          </SearchResultsContainer>
        )}
        {remainingChains?.length > 0 && (
          <SearchResultsContainer label="All Chains">
            {remainingChains.map(({ id: mapChainId }, idx) =>
                <SelectSpecificNetworkButton
                  key={mapChainId}
                  itemChainId={mapChainId}
                  isCurrentChain={fromChainId === mapChainId}
                  isOrigin={true}
                  active={idx + possibleChains.length === currentIdx}
                  onClick={() => handleSetFromChainId(mapChainId)}
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
