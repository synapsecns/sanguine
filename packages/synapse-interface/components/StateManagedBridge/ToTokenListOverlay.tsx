import _ from 'lodash'
import { useEffect, useMemo, useState } from 'react'
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
import { getToTokens } from '@/utils/routeMaker/getToTokens'
import { getSymbol } from '@/utils/routeMaker/generateRoutePossibilities'

import * as ALL_TOKENS from '@constants/tokens/master'
import { toTokenText } from './helpers/toTokensText'
import { sortByBalances } from './helpers/sortByBalance'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'

export const ToTokenListOverlay = () => {
  const { fromChainId, fromToken, toTokens, toChainId, toToken } =
    useBridgeState()
  const portfolioBalances = usePortfolioBalances()

  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()

  let tokenList = sortTokens(toTokens).sort((t) =>
    sortByBalances(t, toChainId, portfolioBalances)
  )

  const fuse = new Fuse(tokenList, {
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
  })

  if (searchStr?.length > 0) {
    tokenList = fuse.search(searchStr).map((i) => i.item)
  }

  const allToTokens = _.uniq(
    getToTokens({
      fromChainId,
      fromTokenRouteSymbol: null,
      toChainId,
      toTokenRouteSymbol: null,
    })
      .map(getSymbol)
      .map((symbol) => ALL_TOKENS[symbol])
  )

  const remainingTokens = _.difference(allToTokens, toTokens).sort((t) =>
    sortByBalances(t, toChainId, portfolioBalances)
  )

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')
  const enterPressed = useKeyPress('Enter')

  function onClose() {
    setCurrentIdx(-1)
    setSearchStr('')
    dispatch(setShowToTokenListOverlay(false))
  }

  function onMenuItemClick(token: Token) {
    dispatch(setToToken(token))
    dispatch(setShowToTokenListOverlay(false))
    onClose()
  }

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  function arrowDownFunc() {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < tokenList.length) {
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
      onMenuItemClick(tokenList[currentIdx])
    }
  }

  useEffect(enterPressedFunc, [enterPressed])

  function onSearch(str: string) {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  const toTokensText = useMemo(() => {
    return toTokenText({ fromChainId, fromToken, toChainId, toToken })
  }, [fromChainId, fromToken, toChainId, toToken])

  const remainingTokensText = useMemo(() => {
    return 'Other tokens'
  }, [])

  return (
    <div
      data-test-id="token-slide-over"
      className="max-h-full pb-4 mt-2 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="flex items-center mb-2 font-medium justfiy-between sm:float-none">
          <SlideSearchBox
            placeholder="Filter by symbol, contract, or name..."
            searchStr={searchStr}
            onSearch={onSearch}
          />
        </div>
      </div>
      <div className="px-2 pb-2 pt-2 text-secondaryTextColor text-sm bg-[#343036]">
        {toTokensText}
      </div>
      <div className="px-2 pb-2 bg-[#343036] md:px-2">
        {tokenList.map((token, idx) => {
          return (
            <SelectSpecificTokenButton
              isOrigin={false}
              key={idx}
              token={token}
              selectedToken={toToken}
              active={idx === currentIdx}
              onClick={() => {
                const eventTitle = `[Bridge User Action] Sets new toToken`
                const eventData = {
                  previousToToken: toToken?.symbol,
                  newToToken: token?.symbol,
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
          <div className="px-2 pb-2 pt-2 text-secondaryTextColor text-sm bg-[#343036]">
            {remainingTokensText}
          </div>
          <div className="px-2 pb-2 bg-[#343036] md:px-2">
            {remainingTokens.map((token, idx) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={false}
                  key={idx}
                  token={token}
                  selectedToken={toToken}
                  active={idx === currentIdx}
                  onClick={() => {
                    const eventTitle = `[Bridge User Action] Sets new toToken`
                    const eventData = {
                      previousToToken: toToken?.symbol,
                      newToToken: token?.symbol,
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

      <div>
        {searchStr && (
          <div className="px-12 py-4 text-xl text-center text-white">
            No other results found for{' '}
            <i className="text-white text-opacity-60">{searchStr}</i>.
            <div className="pt-4 text-lg text-white text-opacity-50 align-bottom text-medium">
              Want to see a token supported on Synapse? Submit a request{' '}
              <span className="text-white text-opacity-70 hover:underline hover:cursor-pointer">
                here
              </span>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
