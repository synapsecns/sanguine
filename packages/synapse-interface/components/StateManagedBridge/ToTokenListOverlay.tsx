import _ from 'lodash'
import { useEffect, useRef, useState } from 'react'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { sortTokens } from '@constants/tokens'
import { Token } from '@/utils/types'
import { setToToken } from '@/slices/bridge/reducer'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

import { sortByBalances } from './helpers/sortByBalance'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from './components/CloseButton'
import { SearchResults } from './components/SearchResults'

export const ToTokenListOverlay = () => {
  const { fromChainId, toTokens, toChainId, toToken } = useBridgeState()
  const portfolioBalances = usePortfolioBalances()

  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  const overlayRef = useRef(null)

  let possibleTokens = sortTokens(toTokens).sort((t) =>
    sortByBalances(t, toChainId, portfolioBalances)
  )

  const { toTokens: allToChainTokens } = getRoutePossibilities({
    fromChainId,
    fromToken: null,
    toChainId,
    toToken: null,
  })

  let remainingChainTokens = _.difference(allToChainTokens, toTokens).sort(
    (t) => sortByBalances(t, toChainId, portfolioBalances)
  )

  const { toTokens: allTokens } = getRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherToTokens = _.difference(allTokens, allToChainTokens)

  const possibleTokenswithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))

  const remainingChainTokensWithSource = remainingChainTokens.map((token) => ({
    ...token,
    source: 'remainingChainTokens',
  }))

  const allOtherToTokensWithSource = allOtherToTokens.map((token) => ({
    ...token,
    source: 'allOtherToTokens',
  }))

  const masterList = [
    ...possibleTokenswithSource,
    ...remainingChainTokensWithSource,
    ...allOtherToTokensWithSource,
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
      `addresses.${toChainId}`,
      'name',
    ],
  }
  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingChainTokens = results.filter(
      (item) => item.source === 'remainingChainTokens'
    )
    allOtherToTokens = results.filter(
      (item) => item.source === 'allOtherToTokens'
    )
  }

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  function onClose() {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowToTokenListOverlay(false))
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

  const handleSetToToken = (oldToken: Token, newToken: Token) => {
    const eventTitle = `[Bridge User Action] Sets new toToken`
    const eventData = {
      previousToToken: oldToken?.symbol,
      newToToken: newToken?.symbol,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setToToken(newToken))
    onClose()
  }

  return (
    <div
      ref={overlayRef}
      data-test-id="token-slide-over"
      className="max-h-full pb-4 mt-2 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="relative flex items-center mb-2 font-medium">
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
          <div className="px-2 pb-2 pt-2 text-primaryTextColor text-sm bg-[#343036]">
            Receive…
          </div>
          <div className="px-2 pb-2 bg-[#343036] md:px-2">
            {possibleTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={false}
                  key={idx}
                  token={token}
                  selectedToken={toToken}
                  active={idx === currentIdx}
                  showAllChains={false}
                  onClick={() => {
                    if (token === toToken) {
                      onClose()
                    } else {
                      handleSetToToken(toToken, token)
                    }
                  }}
                />
              )
            })}
          </div>
        </>
      )}
      {remainingChainTokens && remainingChainTokens.length > 0 && (
        <>
          <div className="px-2 pb-2 pt-2 text-primaryTextColor text-sm bg-[#343036]">
            More on {CHAINS_BY_ID[toChainId]?.name}
          </div>
          <div className="px-2 pb-2 bg-[#343036] md:px-2">
            {remainingChainTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={false}
                  key={idx}
                  token={token}
                  selectedToken={toToken}
                  active={idx + possibleTokens.length === currentIdx}
                  showAllChains={false}
                  onClick={() => handleSetToToken(toToken, token)}
                />
              )
            })}
          </div>
        </>
      )}
      {allOtherToTokens && allOtherToTokens.length > 0 && (
        <>
          <div className="px-2 pb-2 pt-2 text-primaryTextColor text-sm bg-[#343036]">
            All receivable tokens
          </div>
          <div className="px-2 pb-2 bg-[#343036] md:px-2">
            {allOtherToTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={false}
                  key={idx}
                  token={token}
                  selectedToken={toToken}
                  active={
                    idx +
                      possibleTokens.length +
                      remainingChainTokens.length ===
                    currentIdx
                  }
                  showAllChains={true}
                  onClick={() => handleSetToToken(toToken, token)}
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
