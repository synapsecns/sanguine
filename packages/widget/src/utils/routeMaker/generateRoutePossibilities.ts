import _ from 'lodash'
import { BridgeableToken } from 'types'

import { flattenPausedTokens } from '../flattenPausedTokens'
import { getToChainIds } from './getToChainIds'
import { getFromChainIds } from './getFromChainIds'
import { getFromTokens } from './getFromTokens'
import { getToTokens } from './getToTokens'
import { findTokenByRouteSymbol } from '../findTokenByRouteSymbol'

export const PAUSED_TO_CHAIN_IDS = [2000]

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
  fromToken?: BridgeableToken
  toChainId?: number
  toToken?: BridgeableToken
}) => {
  const fromTokenRouteSymbol = fromToken && fromToken.routeSymbol
  const toTokenRouteSymbol = toToken && toToken.routeSymbol

  const fromChainIds: number[] = getFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })

  const fromTokens: BridgeableToken[] = _(
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

  const toTokens: BridgeableToken[] = _(
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
