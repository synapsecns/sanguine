import _ from 'lodash'

import { useEffect, useMemo, useState } from 'react'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'

import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { sortTokens } from '@constants/tokens'
import { Token } from '@/utils/types'
import {
  resetState,
  setFromChainId,
  setFromToken,
  setToToken,
} from '@/slices/bridge/reducer'
import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'
import { fromTokenText } from './helpers/fromTokenText'
import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
import {
  getRoutePossibilities,
  getSymbol,
} from '@/utils/routeMaker/generateRoutePossibilities'

import * as ALL_TOKENS from '@constants/tokens/master'
import { sortByBalances } from './helpers/sortByBalance'

export const FromTokenListOverlay = () => {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()

  const { fromTokens, fromChainId, fromToken, toChainId, toToken } =
    useBridgeState()
  const portfolioBalances = usePortfolioBalances()

  let possibleTokens = sortTokens(fromTokens).sort((t) =>
    sortByBalances(t, fromChainId, portfolioBalances)
  )

  const allFromChainTokens = _.uniq(
    getFromTokens({
      fromChainId,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })
      .map(getSymbol)
      .map((symbol) => ALL_TOKENS[symbol])
  )

  let remainingTokens = _.difference(allFromChainTokens, fromTokens).sort((t) =>
    sortByBalances(t, fromChainId, portfolioBalances)
  )

  const { fromTokens: allTokens } = getRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherFromTokens = _.difference(allTokens, allFromChainTokens)

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
  const enterPressed = useKeyPress('Enter')

  function onClose() {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowFromTokenListOverlay(false))
  }

  function onMenuItemClick(token: Token) {
    if (allFromChainTokens.includes(token)) {
      dispatch(setFromToken(token))
      onClose()
    } else {
      dispatch(resetState())
      dispatch(setFromToken(token))
      const fromChainId = Object.keys(token.addresses)[0]
      dispatch(setFromChainId(Number(fromChainId)))
      dispatch(setToToken(token))
      onClose()
    }
  }

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  function arrowDownFunc() {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  useEffect(arrowDownFunc, [arrowDown])

  function arrowUpFunc() {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  useEffect(arrowUpFunc, [arrowUp])

  function enterPressedFunc() {
    if (enterPressed && currentIdx > -1) {
      onMenuItemClick(masterList[currentIdx])
    }
  }

  useEffect(enterPressedFunc, [enterPressed])

  function onSearch(str: string) {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  const fromTokensText = useMemo(() => {
    return fromTokenText({ fromChainId, fromToken, toChainId, toToken })
  }, [fromChainId, fromToken, toChainId, toToken])

  return (
    <div
      data-test-id="token-slide-over"
      className="max-h-full pb-4 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="flex items-center mt-2 mb-2 font-medium justfiy-between sm:float-none">
          <SlideSearchBox
            placeholder="Filter by symbol, contract, or name..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
        </div>
      </div>
      <div className="px-2 pb-2 pt-2 text-secondaryTextColor text-sm bg-[#343036]">
        {fromTokensText}
      </div>
      <div className="px-2 pb-4 bg-[#343036] md:px-2 ">
        {possibleTokens.map((token, idx) => {
          return (
            <SelectSpecificTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              selectedToken={fromToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = '[Bridge User Action] Sets new fromToken'
                const eventData = {
                  previousFromToken: fromToken?.symbol,
                  newFromToken: token?.symbol,
                }
                segmentAnalyticsEvent(eventTitle, eventData)
                onMenuItemClick(token)
              }}
            />
          )
        })}
      </div>
      {remainingTokens && (
        <>
          <div className="px-2 pb-2 text-secondaryTextColor text-sm bg-[#343036]">
            Other tokens
          </div>
          <div className="px-2 pb-2 bg-[#343036] md:px-2">
            {remainingTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={true}
                  key={idx}
                  token={token}
                  selectedToken={fromToken}
                  active={idx === currentIdx}
                  onClick={() => {
                    const eventTitle = '[Bridge User Action] Sets new fromToken'
                    const eventData = {
                      previousFromToken: fromToken?.symbol,
                      newFromToken: token?.symbol,
                    }
                    segmentAnalyticsEvent(eventTitle, eventData)
                    onMenuItemClick(token)
                  }}
                />
              )
            })}
          </div>
        </>
      )}
      <div className="px-2 pb-2 text-secondaryTextColor text-sm bg-[#343036]">
        All other sendable tokens
      </div>
      <div className="px-2 pb-2 bg-[#343036] md:px-2">
        {allOtherFromTokens.map((token, idx) => {
          return (
            <SelectSpecificTokenButton
              isOrigin={true}
              key={idx}
              token={token}
              selectedToken={fromToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = '[Bridge User Action] Sets new fromToken'
                const eventData = {
                  previousFromToken: fromToken?.symbol,
                  newFromToken: token?.symbol,
                }
                segmentAnalyticsEvent(eventTitle, eventData)
                onMenuItemClick(token)
              }}
            />
          )
        })}
      </div>
      <div>
        {searchStr && (
          <div className="px-12 py-4 text-center text-white text-md">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-2 text-white text-opacity-50 align-bottom text-md">
              Want to see a token supported on Synapse? Let us know!
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
