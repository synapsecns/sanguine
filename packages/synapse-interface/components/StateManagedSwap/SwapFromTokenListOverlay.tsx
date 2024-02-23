import _ from 'lodash'

import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { SlideSearchBox } from '@components/bridgeSwap/SlideSearchBox'
import type { Token } from '@/utils/types'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { SelectSpecificTokenButton } from './components/SelectSpecificTokenButton'

import { hasBalance } from '@/utils/helpers/hasBalance'
import { sortByPriorityRank } from '@/utils/helpers/sortByPriorityRank'
import { CloseButton } from '@/components/buttons/CloseButton'
import { NoSearchResultsFound } from '@/components/bridgeSwap/NoSearchResultsFound'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapFromToken } from '@/slices/swap/reducer'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { getTokenFuseOptions } from '@/constants/fuseOptions'
import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'

export const SwapFromTokenListOverlay = () => {
  const dispatch = useDispatch()

  const { swapFromTokens, swapChainId, swapFromToken } = useSwapState()
  const portfolioBalances = usePortfolioBalances()

  let possibleTokens = sortByPriorityRank(swapFromTokens)

  possibleTokens = [
    ...possibleTokens.filter((t) =>
      hasBalance(t, swapChainId, portfolioBalances)
    ),
    ...possibleTokens.filter(
      (t) => !hasBalance(t, swapChainId, portfolioBalances)
    ),
  ]

  const { fromTokens: allSwapChainTokens } = getSwapPossibilities({
    fromChainId: swapChainId,
    fromToken: null,
    toChainId: swapChainId,
    toToken: null,
  })

  let remainingTokens = sortByPriorityRank(
    _.difference(allSwapChainTokens, swapFromTokens)
  )

  remainingTokens = [
    ...remainingTokens.filter((t) =>
      hasBalance(t, swapChainId, portfolioBalances)
    ),
    ...remainingTokens.filter(
      (t) => !hasBalance(t, swapChainId, portfolioBalances)
    ),
  ]

  const possibleTokensWithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))
  const remainingTokensWithSource = remainingTokens.map((token) => ({
    ...token,
    source: 'remainingTokens',
  }))

  const masterList = [...possibleTokensWithSource, ...remainingTokensWithSource]

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
    remainingTokens = results.filter(
      (item) => item.source === 'remainingTokens'
    )
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
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
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
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
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
