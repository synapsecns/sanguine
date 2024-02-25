import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'
import { CHAINS_ARR, PAUSED_FROM_CHAIN_IDS, sortChains } from '@constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { CHAIN_FUSE_OPTIONS } from '@/constants/fuseOptions'
import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'
import { SelectNetworkButton } from './SelectNetworkButton'


export const ChainListOverlay = ({
  chainId,
  chainIds,
  isOrigin,
  primaryLabel,
  setChainId,
  setShowOverlay,
  filterPausedChains=true,
  SelectNetworkButtonComponent=SelectNetworkButton
}: {
  chainId: any
  chainIds: any[]
  isOrigin: boolean
  primaryLabel: string
  setChainId: (chainId: any) => any
  setShowOverlay: (showOverlay: boolean) => any
  filterPausedChains?: boolean
  SelectNetworkButtonComponent?: React.FC<any>
}) => {
  const dispatch = useDispatch()
  const dataId = 'chain-list'

  let possibleChains = sortAndFilterPausedWithSource({
    chains: CHAINS_ARR.filter((chain) => chainIds.includes(chain.id)),
    filterPausedChains,
    source: 'possibleChains'
  })

  let remainingChains = sortAndFilterPausedWithSource({
    chains: CHAINS_ARR.filter((chain) => !chainIds.includes(chain.id)),
    filterPausedChains,
    source: 'remainingChains'
  })

  const masterList = [...possibleChains, ...remainingChains]

  function onCloseOverlay() {
    dispatch(setShowOverlay(false))
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


  const handleSetChainId = (newChainId) => {
    const eventTitle = `[Bridge User Action] Sets new chainId`
    const eventData = {
      previousChainId: chainId,
      newFromChainId: newChainId,
    }

    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setChainId(newChainId))
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
        <SearchResultsContainer label={`${primaryLabel}…`}>
          {possibleChains.map(({ id: mapChainId }, idx) =>
            <SelectNetworkButtonComponent
              key={idx}
              itemChainId={mapChainId}
              isCurrentChain={chainId === mapChainId}
              isOrigin={isOrigin}
              active={idx === currentIdx}
              onClick={() => {
                if (chainId === mapChainId) {
                  onClose()
                } else {
                  handleSetChainId(mapChainId)
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
            <SelectNetworkButtonComponent
              key={mapChainId}
              itemChainId={mapChainId}
              isCurrentChain={chainId === mapChainId}
              isOrigin={isOrigin}
              active={idx + possibleChains.length === currentIdx}
              onClick={() => handleSetChainId(mapChainId)}
              dataId={dataId}
              alternateBackground={true}
            />
          )}
        </SearchResultsContainer>
      )}
    </SearchOverlayContent>
  )
}


function sortAndFilterPausedWithSource({
  chains,
  filterPausedChains,
  source
}) {
  const filterFunc = filterPausedChains
    ? (chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id)
    : () => true

  return sortChains(chains).filter(filterFunc).map((chain) => ({...chain, source}))

}

