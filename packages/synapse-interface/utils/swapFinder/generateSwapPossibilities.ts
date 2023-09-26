import _ from 'lodash'

import { flattenPausedTokens } from '../flattenPausedTokens'
import { Token } from '../types'
import { getSwapFromChainIds } from './getSwapFromChainIds'
import { getSwapFromTokens } from './getSwapFromTokens'
import { getSwapToTokens } from './getSwapToTokens'
import { PAUSED_TO_CHAIN_IDS } from '@/constants/chains'
import { findTokenByRouteSymbol } from '../findTokenByRouteSymbol'
import { getSymbol } from '@/utils/getSymbol'

export interface RouteQueryFields {
  fromChainId?: number
  fromTokenRouteSymbol?: string
  toChainId?: number
  toTokenRouteSymbol?: string
}

export const getSwapPossibilities = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}) => {
  const fromTokenRouteSymbol = fromToken && fromToken.routeSymbol
  const toTokenRouteSymbol = toToken && toToken.routeSymbol

  const fromChainIds: number[] = getSwapFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })

  const fromTokens: Token[] = _(
    getSwapFromTokens({
      fromChainId,
      fromTokenRouteSymbol,
      toChainId,
      toTokenRouteSymbol,
    })
  )
    .difference(flattenPausedTokens())
    .map(getSymbol)
    .uniq()
    .map((symbol) => findTokenByRouteSymbol(symbol))
    .compact()
    .value()

  const toTokens: Token[] = _(
    getSwapToTokens({
      fromChainId,
      fromTokenRouteSymbol,
      toChainId,
      toTokenRouteSymbol,
    })
  )
    .difference(flattenPausedTokens())
    .filter((token) => {
      return !PAUSED_TO_CHAIN_IDS.some((value) => token.endsWith(`-${value}`))
    })
    .map(getSymbol)
    .uniq()
    .map((symbol) => findTokenByRouteSymbol(symbol))
    .compact()
    .value()

  return {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toTokens,
  }
}
