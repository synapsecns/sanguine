import _ from 'lodash'

import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { SlideSearchBox } from '@components/bridgeSwap/SlideSearchBox'
import type { Token } from '@/utils/types'
import { setFromToken } from '@/slices/bridge/reducer'
import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { SelectSpecificTokenButton } from './components/SelectSpecificTokenButton'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

import { hasBalance } from '@/utils/helpers/hasBalance'
import { sortByPriorityRank } from '@/utils/helpers/sortByPriorityRank'
import { CHAINS_BY_ID } from '@/constants/chains'
import { CloseButton } from '@/components/buttons/CloseButton'
import { SearchResults } from '@/components/bridgeSwap/SearchResults'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { getTokenFuseOptions } from '@/constants/fuseOptions'
import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'

export const FromTokenListOverlay = () => {
  const dispatch = useDispatch()

  const { fromTokens, fromChainId, fromToken } = useBridgeState()
  const portfolioBalances = usePortfolioBalances()

  let possibleTokens = sortByPriorityRank(fromTokens)

  possibleTokens = [
    ...possibleTokens.filter((t) =>
      hasBalance(t, fromChainId, portfolioBalances)
    ),
    ...possibleTokens.filter(
      (t) => !hasBalance(t, fromChainId, portfolioBalances)
    ),
  ]

  const { fromTokens: allFromChainTokens } = getRoutePossibilities({
    fromChainId,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let remainingTokens = sortByPriorityRank(
    _.difference(allFromChainTokens, fromTokens)
  )

  remainingTokens = [
    ...remainingTokens.filter((t) =>
      hasBalance(t, fromChainId, portfolioBalances)
    ),
    ...remainingTokens.filter(
      (t) => !hasBalance(t, fromChainId, portfolioBalances)
    ),
  ]

  const { fromTokens: allTokens } = getRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherFromTokens = sortByPriorityRank(
    _.difference(allTokens, allFromChainTokens)
  )

  const possibleTokensWithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))
  const remainingTokensWithSource = remainingTokens.map((token) => ({
    ...token,
    source: 'remainingTokens',
  }))
  const allOtherFromTokensWithSource = allOtherFromTokens.map((token) => ({
    ...token,
    source: 'allOtherFromTokens',
  }))

  const masterList = [
    ...possibleTokensWithSource,
    ...remainingTokensWithSource,
    ...allOtherFromTokensWithSource,
  ]

  function onCloseOverlay() {
    dispatch(setShowFromTokenListOverlay(false))
  }

  const {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose,
  } = useOverlaySearch(masterList.length, onCloseOverlay)

  const fuse = new Fuse(masterList, getTokenFuseOptions(fromChainId))

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)
    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingTokens = results.filter(
      (item) => item.source === 'remainingTokens'
    )
    allOtherFromTokens = results.filter(
      (item) => item.source === 'allOtherFromTokens'
    )
  }


  const handleSetFromToken = (oldToken: Token, newToken: Token) => {
    const eventTitle = '[Bridge User Action] Sets new fromToken'
    const eventData = {
      previousFromToken: oldToken?.symbol,
      newFromToken: newToken?.symbol,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setFromToken(newToken))
    onClose()
  }

  return (
    <div
      ref={overlayRef}
      data-test-id="token-slide-over"
      className="max-h-full pb-4 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="relative flex items-center mt-2 mb-2 font-medium">
          <SlideSearchBox
            placeholder="Filter by symbol, contract, or name..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <CloseButton onClick={onClose} />
        </div>
      </div>
      {possibleTokens?.length > 0 && (
        <SearchResultsContainer label="Send...">
            {possibleTokens.map((token, idx) =>
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
                  selectedToken={fromToken}
                  active={idx === currentIdx}
                  showAllChains={false}
                  onClick={() => {
                    if (token === fromToken) {
                      onClose()
                    } else {
                      handleSetFromToken(fromToken, token)
                    }
                  }}
                />
            )}
        </SearchResultsContainer>
      )}
      {remainingTokens?.length > 0 && (
        <SearchResultsContainer
          label={
            fromChainId
              ? `More on ${CHAINS_BY_ID[fromChainId]?.name}`
              : 'All sendable tokens'
          }
        >
            {remainingTokens.map((token, idx) =>
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
                  selectedToken={fromToken}
                  active={idx + possibleTokens.length === currentIdx}
                  showAllChains={false}
                  onClick={() => handleSetFromToken(fromToken, token)}
                />
            )}
        </SearchResultsContainer>
      )}
      {allOtherFromTokens?.length > 0 && (
        <SearchResultsContainer label="All sendable tokens">
            {allOtherFromTokens.map((token, idx) =>
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
                  selectedToken={fromToken}
                  active={
                    idx + possibleTokens.length + remainingTokens.length ===
                    currentIdx
                  }
                  showAllChains={true}
                  onClick={() => handleSetFromToken(fromToken, token)}
                  alternateBackground={true}
                />
            )}
        </SearchResultsContainer>
      )}
      <SearchResults searchStr={searchStr} type="token" />
    </div>
  )
}


