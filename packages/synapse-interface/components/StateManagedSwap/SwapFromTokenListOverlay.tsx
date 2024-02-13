import _ from 'lodash'

import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { Token } from '@/utils/types'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'

import { hasBalance } from '@/utils/helpers/hasBalance'
import { sortByPriorityRank } from '@/utils/helpers/sortByPriorityRank'
import { CloseButton } from '@/components/buttons/CloseButton'
import { SearchResults } from '@/components/SearchResults'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapFromToken } from '@/slices/swap/reducer'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'

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

  const fuseOptions = {
    ignoreLocation: true,
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'symbol',
        weight: 2,
      },
      'routeSymbol',
      `addresses.${swapChainId}`,
      'name',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

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
      {possibleTokens && possibleTokens.length > 0 && (
        <>
          <div className="px-2 pt-2 pb-4 text-sm text-primaryTextColor ">
            Swap…
          </div>
          <div className="px-2 pb-2 md:px-2 ">
            {possibleTokens.map((token, idx) => {
              return (
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
              )
            })}
          </div>
        </>
      )}
      {remainingTokens && remainingTokens.length > 0 && (
        <>
          <div className="px-2 pb-4 text-sm text-primaryTextColor">
            {swapChainId
              ? `More on ${CHAINS_BY_ID[swapChainId]?.name}`
              : 'All swappable tokens'}
          </div>
          <div className="px-2 pb-2 md:px-2">
            {remainingTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
                  selectedToken={swapFromToken}
                  active={idx + possibleTokens.length === currentIdx}
                  showAllChains={false}
                  onClick={() => handleSetFromToken(swapFromToken, token)}
                />
              )
            })}
          </div>
        </>
      )}
      <SearchResults searchStr={searchStr} type="token" />
    </div>
  )
}
