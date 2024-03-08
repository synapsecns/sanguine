import _ from 'lodash'

import { Token } from './types'

export const findValidToken = (
  tokens: Token[],
  routeSymbol: string,
  swapableType: string
): Token | null => {
  const matchingToken = tokens?.find((t) => t.routeSymbol === routeSymbol)

  const swapableToken = _.orderBy(tokens, ['priorityRank'], ['asc'])?.find(
    (t) => t.swapableType === swapableType
  )

  return matchingToken ? matchingToken : swapableToken ? swapableToken : null
}
