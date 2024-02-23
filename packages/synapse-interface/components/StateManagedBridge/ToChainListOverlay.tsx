import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import * as ALL_CHAINS from '@constants/chains/master'
import { CHAINS_BY_ID, PAUSED_TO_CHAIN_IDS, sortChains } from '@constants/chains'
import { CHAIN_FUSE_OPTIONS } from '@/constants/fuseOptions'

import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setToChainId } from '@/slices/bridge/reducer'
import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'

import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'

export const ToChainListOverlay = () => {
  const { toChainIds, toChainId } = useBridgeState()

  const dispatch = useDispatch()


  const dataId = 'bridge-destination-chain-list'

  let possibleChains = _(ALL_CHAINS)
    .pickBy((value) => _.includes(toChainIds, value.id))
    .values()
    .value()
    .filter((chain) => !PAUSED_TO_CHAIN_IDS.includes(chain.id))

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
      toChainIds.map((id) => CHAINS_BY_ID[id])
    )
  ).filter((chain) => !PAUSED_TO_CHAIN_IDS.includes(chain.id))

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
    dispatch(setShowToChainListOverlay(false))
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
    remainingChains = results.filter((item) => item.source === 'remainingChains')
  }


  const handleSetToChainId = (chainId) => {
    const eventTitle = `[Bridge User Action] Sets new toChainId`
    const eventData = {
      previousToChainId: toChainId,
      newToChainId: chainId,
    }

    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setToChainId(chainId))
    onClose()
  }

  return (
    <SearchOverlayContent
      overlayRef={overlayRef}
      searchStr={searchStr}
      onSearch={onSearch}
      onClose={onClose}
      type="chain"
    >
      {possibleChains?.length > 0 && (
        <SearchResultsContainer label="To…">
          {possibleChains.map(({ id: mapChainId }, idx) =>
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={toChainId === mapChainId}
              isOrigin={false}
              active={idx === currentIdx}
              onClick={() => {
                if (toChainId === mapChainId) {
                  onClose()
                } else {
                  handleSetToChainId(mapChainId)
                }
              }}
              dataId={dataId}
            />
          )}
        </SearchResultsContainer>
      )}
      {remainingChains?.length > 0 && (
        <SearchResultsContainer label="All chains">
          {remainingChains.map(({ id: mapChainId }, idx) =>
            <SelectSpecificNetworkButton
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={toChainId === mapChainId}
              isOrigin={false}
              active={idx + possibleChains.length === currentIdx}
              onClick={() => handleSetToChainId(mapChainId)}
              dataId={dataId}
              alternateBackground={true}
            />
          )}
        </SearchResultsContainer>
      )}
    </SearchOverlayContent>
  )
}
