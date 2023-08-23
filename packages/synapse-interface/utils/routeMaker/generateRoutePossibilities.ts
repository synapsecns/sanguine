import _ from 'lodash'

import { flattenPausedTokens } from '../flattenPausedTokens'
import { Token } from '../types'
import { getToChainIds } from './getToChainIds'
import { getFromChainIds } from './getFromChainIds'
import { getFromTokens } from './getFromTokens'
import { getToTokens } from './getToTokens'
import { PAUSED_TO_CHAIN_IDS } from '@/constants/chains'
import { findTokenByRouteSymbol } from '../findTokenByRouteSymbol'

export interface RouteQueryFields {
  fromChainId?: number
  fromTokenRouteSymbol?: string
  toChainId?: number
  toTokenRouteSymbol?: string
}

export const getRoutePossibilities = ({
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

  const fromChainIds: number[] = getFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })

  const fromTokens: Token[] = _(
    getFromTokens({
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

  const toChainIds: number[] = getToChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })
    ?.filter((chainId) => !PAUSED_TO_CHAIN_IDS.includes(chainId))
    .filter((chainId) => chainId !== fromChainId)

  const toTokens: Token[] = _(
    getToTokens({
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
    toChainIds,
    toTokens,
  }
}

export const getSymbol = (tokenAndChainId: string): string => {
  return tokenAndChainId.split('-')[0]
}
