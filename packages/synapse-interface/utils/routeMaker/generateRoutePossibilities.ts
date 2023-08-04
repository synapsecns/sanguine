import _ from 'lodash'
import * as ALL_TOKENS from '@constants/tokens/master'

import { flattenPausedTokens } from '../flattenPausedTokens'
import { Token } from '../types'
import { getToChainIds } from './getToChainIds'
import { getFromChainIds } from './getFromChainIds'
import { getFromTokens } from './getFromTokens'
import { getToTokens } from './getToTokens'

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
  fromChainId: number
  fromToken: Token
  toChainId: number
  toToken: Token
}) => {
  const fromTokenRouteSymbol = fromToken && fromToken.routeSymbol
  const toTokenRouteSymbol = toToken && toToken.routeSymbol

  const fromChainIds = getFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })

  const fromTokens = _.uniq(
    _.difference(
      getFromTokens({
        fromChainId,
        fromTokenRouteSymbol,
        toChainId,
        toTokenRouteSymbol,
      }),
      flattenPausedTokens()
    ).map(getSymbol)
  ).map((symbol) => ALL_TOKENS[symbol])

  const toChainIds = getToChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })?.filter((chainId) => chainId !== fromChainId)

  const toTokens = _.uniq(
    _.difference(
      getToTokens({
        fromChainId,
        fromTokenRouteSymbol,
        toChainId,
        toTokenRouteSymbol,
      }),
      flattenPausedTokens()
    ).map(getSymbol)
  ).map((symbol) => ALL_TOKENS[symbol])

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

/*NOTES

  * Set intelligent defaults when fromTokens/toTokens change
  * check if from/to token is allowed in list


  -- To strip out --
  * paused tokens [x]
  * swap exceptions?

*/
