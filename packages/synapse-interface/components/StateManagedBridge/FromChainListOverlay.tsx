import _ from 'lodash'
import { useCallback, useEffect, useRef, useState } from 'react'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import * as ALL_CHAINS from '@constants/chains/master'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from '@/components/buttons/CloseButton'
import { SearchResults } from '@/components/SearchResults'
import { PAUSED_FROM_CHAIN_IDS } from '@constants/chains'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'

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
        {possibleChains && possibleChains.length > 0 && (
          <>
            <div className="mb-4 text-sm text-primaryTextColor">From…</div>
            {possibleChains.map(({ id: mapChainId }, idx) => {
              return (
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
              )
            })}
          </>
        )}
        {remainingChains && remainingChains.length > 0 && (
          <>
            <div className="mt-4 mb-4 text-sm font-normal text-primaryTextColor">
              All chains
            </div>
            {remainingChains.map(({ id: mapChainId }, idx) => {
              return (
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
              )
            })}
          </>
        )}
        <SearchResults searchStr={searchStr} type="chain" />
      </div>
    </div>
  )
}
