import _ from 'lodash'

import { useEffect, useRef, useState } from 'react'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { Token } from '@/utils/types'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'

import { hasBalance } from './helpers/sortByBalance'
import { sortByPriorityRank } from './helpers/sortByPriorityRank'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from './components/CloseButton'
import { SearchResults } from './components/SearchResults'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapFromToken } from '@/slices/swap/reducer'

export const SwapFromTokenListOverlay = () => {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  const overlayRef = useRef(null)

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

  const possibleTokensWithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))

  const masterList = [...possibleTokensWithSource]

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
  }

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  function onClose() {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowSwapFromTokenListOverlay(false))
  }

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  function arrowDownFunc() {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  function arrowUpFunc() {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  function onSearch(str: string) {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(escFunc, [escPressed])
  useEffect(arrowDownFunc, [arrowDown])
  useEffect(arrowUpFunc, [arrowUp])
  useCloseOnOutsideClick(overlayRef, onClose)

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
            Send…
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
      <SearchResults searchStr={searchStr} type="token" />
    </div>
  )
}
