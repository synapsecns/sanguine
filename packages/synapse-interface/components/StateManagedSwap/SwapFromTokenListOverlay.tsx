import _ from 'lodash'

import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { getTokenFuseOptions } from '@/constants/fuseOptions'
import { CHAINS_BY_ID } from '@/constants/chains'

import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'

import type { Token } from '@/utils/types'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { sortByPriorityRankAndBalance } from '@/utils/helpers/sortByPriorityRankAndBalance'

import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapFromToken } from '@/slices/swap/reducer'

import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'
import { SelectTokenButton } from '@/components/bridgeSwap/SelectTokenButton'

export const SwapFromTokenListOverlay = () => {
  const dispatch = useDispatch()

  const { swapFromTokens, swapChainId, swapFromToken } = useSwapState()
  const portfolioBalances = usePortfolioBalances()

  const common = {
    chainId: swapChainId,
    portfolioBalances: portfolioBalances
  }
  let possibleTokens = sortByPriorityRankAndBalance({
    tokens: swapFromTokens,
    source: 'possibleTokens',
    ...common
  })


  const { fromTokens: allSwapChainTokens } = getSwapPossibilities({
    fromChainId: swapChainId,
    fromToken: null,
    toChainId: swapChainId,
    toToken: null,
  })

  let remainingTokens = sortByPriorityRankAndBalance({
    tokens: _.difference(allSwapChainTokens, swapFromTokens),
    source: 'remainingTokens',
    ...common
  })

  const masterList = [...possibleTokens, ...remainingTokens]

  function onCloseOverlay() {
    dispatch(setShowSwapFromTokenListOverlay(false))
  }

  const {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose,
  } = useOverlaySearch(masterList.length, onCloseOverlay)

  const fuse = new Fuse(masterList, getTokenFuseOptions(swapChainId))

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)
    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingTokens = results.filter((item) => item.source === 'remainingTokens')
  }

  const handleSetFromToken = (oldToken: Token, newToken: Token) => {
    const eventTitle = '[Swap User Action] Sets new fromToken'
    const eventData = {
      previousFromToken: oldToken?.symbol,
      newFromToken: newToken?.symbol,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setSwapFromToken(newToken))
    onClose()
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
        <SearchResultsContainer label="Swap…">
          {possibleTokens.map((token, idx) =>
            <SelectTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              chainId={swapChainId}
              selectedToken={swapFromToken}
              active={idx === currentIdx}
              showAllChains={false}
              onClick={() => {
                if (token === swapFromToken) {
                  onClose()
                } else {
                  handleSetFromToken(swapFromToken, token)
                }
              }}
            />
          )}
        </SearchResultsContainer>
      )}
      {remainingTokens?.length > 0 && (
        <SearchResultsContainer
          label={
            swapChainId
              ? `More on ${CHAINS_BY_ID[swapChainId]?.name}`
              : 'All swappable tokens'
          }
        >
          {remainingTokens.map((token, idx) =>
            <SelectTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              chainId={swapChainId}
              selectedToken={swapFromToken}
              active={idx + possibleTokens.length === currentIdx}
              showAllChains={false}
              onClick={() => handleSetFromToken(swapFromToken, token)}
            />
          )}
        </SearchResultsContainer>
      )}
    </SearchOverlayContent>

  )
}
