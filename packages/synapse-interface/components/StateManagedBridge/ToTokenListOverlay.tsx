import _ from 'lodash'
import { useMemo } from 'react'
import Fuse from 'fuse.js'

import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState, setToToken } from '@/slices/bridge/reducer'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { FetchState } from '@/slices/portfolio/actions'

import { useAppDispatch } from '@/store/hooks'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import type { Token } from '@/utils/types'
import { useAlternateBridgeQuotes } from '@/utils/hooks/useAlternateBridgeQuotes'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'
import { sortByPriorityRank } from '@/utils/helpers/sortByPriorityRank'
import { formatBigIntToString } from '@/utils/bigint/format'

import { CHAINS_BY_ID } from '@/constants/chains'
import { getTokenFuseOptions } from '@/constants/fuseOptions'

import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'
import { SelectSpecificTokenButton } from './components/SelectSpecificTokenButton'


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
    debouncedToTokensFromValue,
    bridgeQuote,
  }: BridgeState = useBridgeState()

  const dispatch = useAppDispatch()


  /** Fetch Alternative Bridge Quotes when component renders */
  /** Temporarily pausing feature */
  // useAlternateBridgeQuotes()

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

  function onCloseOverlay() {
    dispatch(setShowToTokenListOverlay(false))
  }

  const {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose,
  } = useOverlaySearch(masterList.length, onCloseOverlay)


  const fuse = new Fuse(masterList, getTokenFuseOptions(toChainId))

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingChainTokens = results.filter((item) => item.source === 'remainingChainTokens')
    allOtherToTokens = results.filter((item) => item.source === 'allOtherToTokens')
  }


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
    <SearchOverlayContent
      overlayRef={overlayRef}
      searchStr={searchStr}
      onSearch={onSearch}
      onClose={onClose}
      type="token"
    >
      {orderedPossibleTokens?.length > 0 && (
        <SearchResultsContainer label="Receive…">
          {orderedPossibleTokens.map((token: TokenWithRates, idx: number) =>
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
          )}
        </SearchResultsContainer>
      )}
      {remainingChainTokens?.length > 0 && (
        <SearchResultsContainer
          label={
            toChainId
              ? `More on ${CHAINS_BY_ID[toChainId]?.name}`
              : 'All receivable tokens'
          }
        >
          {remainingChainTokens.map((token, idx) =>
            <SelectSpecificTokenButton
              isOrigin={false}
              key={idx}
              token={token}
              selectedToken={toToken}
              active={idx + possibleTokens.length === currentIdx}
              showAllChains={false}
              onClick={() => handleSetToToken(toToken, token)}
            />
          )}
        </SearchResultsContainer>
      )}
      {allOtherToTokens?.length > 0 && (
        <SearchResultsContainer label="All receivable tokens">
          {allOtherToTokens.map((token, idx) =>
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
          )}
        </SearchResultsContainer>
      )}
    </SearchOverlayContent>
  )
}
