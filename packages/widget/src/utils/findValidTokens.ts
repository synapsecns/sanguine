import _ from 'lodash'
import { BridgeableToken } from 'types'

export const findValidToken = (
  tokens: BridgeableToken[],
  routeSymbol: string,
  swapableType: string
): BridgeableToken | null => {
  const matchingToken = tokens?.find((t) => t.routeSymbol === routeSymbol)
  const swapableToken = _.orderBy(tokens, ['priorityRank'], ['asc'])?.find(
    (t) => t.swapableType === swapableType
  )

  return matchingToken ? matchingToken : swapableToken ? swapableToken : null
}
