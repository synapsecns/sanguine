import _ from 'lodash'

import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import type { Token } from '@/utils/types'
import { sortByPriorityRankAndBalance } from '@/utils/helpers/sortByPriorityRankAndBalance'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'

import { setFromToken } from '@/slices/bridge/reducer'
import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { getTokenFuseOptions } from '@/constants/fuseOptions'

import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'
import { SelectSpecificTokenButton } from './components/SelectSpecificTokenButton'


export const FromTokenListOverlay = () => {
  const dispatch = useDispatch()

  const { fromTokens, fromChainId, fromToken } = useBridgeState()
  const portfolioBalances = usePortfolioBalances()

  const common = {
    chainId: fromChainId,
    portfolioBalances: portfolioBalances
  }

  let possibleTokens = sortByPriorityRankAndBalance({
    tokens: fromTokens,
    source: 'possibleTokens',
    ...common
  })

  const { fromTokens: allFromChainTokens } = getRoutePossibilities({
    fromChainId,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let remainingTokens = sortByPriorityRankAndBalance({
    tokens: _.difference(allFromChainTokens, fromTokens),
    source: 'remainingTokens',
    ...common
  })


  const { fromTokens: allTokens } = getRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherFromTokens = sortByPriorityRankAndBalance({
    tokens: _.difference(allTokens, allFromChainTokens),
    source: 'allOtherFromTokens',
    ...common
  })

  const masterList = [
    ...possibleTokens,
    ...remainingTokens,
    ...allOtherFromTokens,
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
    remainingTokens = results.filter((item) => item.source === 'remainingTokens')
    allOtherFromTokens = results.filter((item) => item.source === 'allOtherFromTokens')
  }


  const handleSetFromToken = (oldToken: Token, newToken: Token) => {
    const eventTitle = '[Bridge User Action] Sets new fromToken'
    const eventData = {
      previousFromToken: oldToken?.symbol,
      newFromToken: newToken?.symbol,
    }

    dispatch(setFromToken(newToken))
    onClose()

    segmentAnalyticsEvent(eventTitle, eventData)
  }

  return (
    <SearchOverlayContent
      overlayRef={overlayRef}
      searchStr={searchStr}
      onSearch={onSearch}
      onClose={onClose}
      type="token"
    >
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
    </SearchOverlayContent>
  )
}


