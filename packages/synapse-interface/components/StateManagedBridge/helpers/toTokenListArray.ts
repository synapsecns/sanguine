import _ from 'lodash'
import { useMemo } from 'react'
import Fuse from 'fuse.js'

import { Token } from '@/utils/types'
import { BridgeState } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'

import { sortByPriorityRank } from './sortByPriorityRank'
import { FetchState } from '@/slices/portfolio/actions'

interface TokenWithRates extends Token {
  exchangeRate: bigint
  estimatedTime: number
}

export const toTokenListArray = (searchStr: string = '') => {
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

  return {
    'Receive…': possibleTokens,
    'All receivable tokens': remainingChainTokens,
    'All other tokens': allOtherToTokens,
  }
}
