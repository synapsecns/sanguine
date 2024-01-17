import _ from 'lodash'

import { useEffect, useRef, useState } from 'react'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { Token } from '@/utils/types'
import { setFromToken } from '@/slices/bridge/reducer'
import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

import { hasBalance } from './helpers/sortByBalance'
import { sortByPriorityRank } from './helpers/sortByPriorityRank'
import { CHAINS_BY_ID } from '@/constants/chains'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from './components/CloseButton'
import { SearchResults } from './components/SearchResults'

export const FromTokenListOverlay = () => {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  const overlayRef = useRef(null)

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
      `addresses.${fromChainId}`,
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
    allOtherFromTokens = results.filter(
      (item) => item.source === 'allOtherFromTokens'
    )
  }

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  function onClose() {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowFromTokenListOverlay(false))
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
      className="max-h-full overflow-auto"
    >
      <div className="flex items-center m-2 rounded bg-zinc-200 dark:bg-zinc-900 custom-shadow">
        <SlideSearchBox
          placeholder="Filter by symbol, contract, or name..."
          searchStr={searchStr}
          onSearch={onSearch}
        />
        <CloseButton onClick={onClose} />
      </div>
      {possibleTokens && possibleTokens.length > 0 && (
        <>
          <div className="px-2 pt-2 pb-4 text-sm">
            Send…
          </div>
          <div className="px-2 pb-2 md:px-2 ">
            {possibleTokens.map((token, idx) => {
              return (
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
              )
            })}
          </div>
        </>
      )}
      {remainingTokens && remainingTokens.length > 0 && (
        <>
          <div className="px-2 pb-4 text-sm">
            {fromChainId
              ? `More on ${CHAINS_BY_ID[fromChainId]?.name}`
              : 'All sendable tokens'}
          </div>
          <div className="px-2 pb-2 md:px-2">
            {remainingTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
                  selectedToken={fromToken}
                  active={idx + possibleTokens.length === currentIdx}
                  showAllChains={false}
                  onClick={() => handleSetFromToken(fromToken, token)}
                />
              )
            })}
          </div>
        </>
      )}
      {allOtherFromTokens && allOtherFromTokens.length > 0 && (
        <>
          <div className="px-2 pb-4 text-sm">
            All sendable tokens
          </div>
          <div className="px-2 pb-2 md:px-2">
            {allOtherFromTokens.map((token, idx) => {
              return (
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
              )
            })}
          </div>
        </>
      )}
      <SearchResults searchStr={searchStr} type="token" />
    </div>
  )
}
