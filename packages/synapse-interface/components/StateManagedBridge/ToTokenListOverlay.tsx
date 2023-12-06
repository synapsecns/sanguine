import _ from 'lodash'
import { useEffect, useRef, useState, useMemo } from 'react'
import { useDispatch } from 'react-redux'
import { Address } from 'viem'
import Fuse from 'fuse.js'

import { useKeyPress } from '@hooks/useKeyPress'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { Token } from '@/utils/types'
import { BridgeState, setToToken } from '@/slices/bridge/reducer'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import SelectSpecificTokenButton from './components/SelectSpecificTokenButton'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

import { sortByPriorityRank } from './helpers/sortByPriorityRank'
import { CHAINS_BY_ID } from '@/constants/chains'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from './components/CloseButton'
import { SearchResults } from './components/SearchResults'
import { formatBigIntToString } from '@/utils/bigint/format'
import { FetchState } from '@/slices/portfolio/actions'
import { calculateEstimatedTransactionTime } from '@/utils/calculateEstimatedTransactionTime'

interface TokenWithRates extends Token {
  exchangeRate: bigint
  estimatedTime: number
}

export const ToTokenListOverlay = () => {
  const {
    fromChainId,
    fromToken,
    toTokens,
    toChainId,
    toToken,
    toTokensBridgeQuotes,
    toTokensBridgeQuotesStatus,
  }: BridgeState = useBridgeState()

  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const dispatch = useDispatch()
  const overlayRef = useRef(null)

  let possibleTokens: Token[] = sortByPriorityRank(toTokens)

  const { toTokens: allToChainTokens } = getRoutePossibilities({
    fromChainId,
    fromToken: null,
    toChainId,
    toToken: null,
  })

  let remainingChainTokens = toChainId
    ? sortByPriorityRank(_.difference(allToChainTokens, toTokens))
    : []

  const { toTokens: allTokens } = getRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherToTokens = toChainId
    ? sortByPriorityRank(_.difference(allTokens, allToChainTokens))
    : sortByPriorityRank(allTokens)

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

  const isLoadingExchangeRate = useMemo(() => {
    const hasRequiredUserInput: boolean = Boolean(
      fromChainId && toChainId && fromToken && toToken
    )
    const isFetchLoading: boolean =
      toTokensBridgeQuotesStatus === FetchState.IDLE ||
      toTokensBridgeQuotesStatus === FetchState.LOADING

    return hasRequiredUserInput && isFetchLoading
  }, [fromChainId, toChainId, fromToken, toToken, toTokensBridgeQuotesStatus])

  const bridgeQuotesMatchDestination: boolean = useMemo(() => {
    return (
      Array.isArray(toTokensBridgeQuotes) &&
      toTokensBridgeQuotes[0]?.destinationChainId === toChainId
    )
  }, [toTokensBridgeQuotes, toChainId])

  const orderedPossibleTokens: TokenWithRates[] | Token[] = useMemo(() => {
    if (
      toTokensBridgeQuotesStatus === FetchState.VALID &&
      bridgeQuotesMatchDestination &&
      possibleTokens &&
      possibleTokens.length > 0
    ) {
      const bridgeQuotesMap = new Map(
        toTokensBridgeQuotes.map((quote) => [quote.destinationToken, quote])
      )

      const tokensWithRates: TokenWithRates[] = possibleTokens.map((token) => {
        const bridgeQuote = bridgeQuotesMap.get(token)
        if (bridgeQuote) {
          return {
            ...token,
            exchangeRate: bridgeQuote?.exchangeRate,
            estimatedTime: bridgeQuote?.estimatedTime,
          }
        } else {
          return token as TokenWithRates
        }
      })

      const sortedTokens = tokensWithRates.sort(
        (a, b) => Number(b.exchangeRate) - Number(a.exchangeRate)
      )

      return sortedTokens
    }
    return possibleTokens
  }, [
    possibleTokens,
    toTokensBridgeQuotes,
    toTokensBridgeQuotesStatus,
    bridgeQuotesMatchDestination,
  ])

  const totalPossibleTokens: number = useMemo(() => {
    return orderedPossibleTokens.length
  }, [orderedPossibleTokens])

  return (
    <div
      ref={overlayRef}
      data-test-id="to-token-list-overlay"
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
      {orderedPossibleTokens && orderedPossibleTokens.length > 0 && (
        <>
          <div className="px-2 pt-2 pb-2 text-sm text-primaryTextColor ">
            Receive…
          </div>
          <div className="px-2 pb-2 md:px-2">
            {orderedPossibleTokens.map((token: TokenWithRates, idx: number) => {
              return (
                <SelectSpecificTokenButton
                  isOrigin={false}
                  key={idx}
                  token={token}
                  selectedToken={toToken}
                  active={idx === currentIdx}
                  showAllChains={false}
                  isLoadingExchangeRate={isLoadingExchangeRate}
                  isBestExchangeRate={totalPossibleTokens > 1 && idx === 0}
                  exchangeRate={formatBigIntToString(
                    token?.exchangeRate,
                    18,
                    4
                  )}
                  estimatedDurationInSeconds={
                    toTokensBridgeQuotesStatus === FetchState.VALID &&
                    bridgeQuotesMatchDestination &&
                    token.estimatedTime
                  }
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
          <div className="px-2 pt-2 pb-2 text-sm text-primaryTextColor ">
            {toChainId
              ? `More on ${CHAINS_BY_ID[toChainId]?.name}`
              : 'All receivable tokens'}
          </div>
          <div className="px-2 pb-2 md:px-2">
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
          <div className="px-2 pt-2 pb-2 text-sm text-primaryTextColor ">
            All receivable tokens
          </div>
          <div className="px-2 pb-2 md:px-2">
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
