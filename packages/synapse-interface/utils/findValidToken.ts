import { Token } from './types'

export const findValidToken = (
  tokens: Token[],
  routeSymbol: string,
  swapableType: string
) => {
  const matchingToken = tokens?.find((t) => t.routeSymbol === routeSymbol)
  const swapableToken = tokens?.find((t) => t.swapableType === swapableType)

  return matchingToken ? matchingToken : swapableToken ? swapableToken : null
}
