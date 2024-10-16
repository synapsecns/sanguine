import { BRIDGE_MAP } from '../constants/bridgeMap'
import * as ALL_TOKENS from '../constants/bridgeable'

type TokenData = {
  symbol: string
  address: string
  chainId: number
}

type StringifiedBridgeRoutes = Record<string, string[]>
type TransformedBridgeRoutes = Record<string, TokenData[]>

const constructJSON = (
  swappableMap,
  exclusionList
): StringifiedBridgeRoutes => {
  const result = {}

  // Iterate through the chains
  for (const chainA in swappableMap) {
    for (const tokenA in swappableMap[chainA]) {
      const symbolA = swappableMap[chainA][tokenA].symbol
      const key = `${symbolA}-${chainA}`

      if (exclusionList.includes(key)) {
        continue
      }

      // Iterate through other chains to compare
      for (const chainB in swappableMap) {
        if (chainA !== chainB) {
          for (const tokenB in swappableMap[chainB]) {
            const symbolB = swappableMap[chainB][tokenB].symbol
            const value = `${symbolB}-${chainB}`

            if (exclusionList.includes(value)) {
              continue
            }

            // Check if there's a bridge between the origins and destinations
            for (const bridgeSymbol of swappableMap[chainA][tokenA].origin) {
              if (
                swappableMap[chainA][tokenA].origin.includes(bridgeSymbol) &&
                swappableMap[chainB][tokenB].destination.includes(bridgeSymbol)
              ) {
                // Add to the result if the key exists, else create a new array
                if (result[key]) {
                  result[key].push(value)
                } else {
                  result[key] = [value]
                }
              }
            }
          }
        }
      }
    }
  }
  return result
}

const transformPair = (string: string): any => {
  const [symbol, chainId] = string.split('-')
  const token = Object.values(ALL_TOKENS).find((t) => t.routeSymbol === symbol)
  const address = token?.addresses[chainId]
  if (token && address) {
    return {
      symbol,
      chainId,
      address,
    }
  }
}

const transformBridgeRouteValues = (
  routes: StringifiedBridgeRoutes
): TransformedBridgeRoutes => {
  return Object.fromEntries(
    Object.entries(routes).map(([key, values]) => {
      const uniquePairs: TokenData[] = []
      values.forEach((pairStr) => {
        const transformedPair = transformPair(pairStr)
        if (
          transformedPair &&
          !uniquePairs.some(
            (pair) =>
              pair.symbol === transformedPair.symbol &&
              pair.chainId === transformedPair.chainId
          )
        ) {
          uniquePairs.push(transformedPair)
        }
      })
      return [key, uniquePairs]
    })
  )
}

export const BRIDGE_ROUTE_MAPPING_SYMBOLS = constructJSON(BRIDGE_MAP, [])
export const BRIDGE_ROUTE_MAPPING = transformBridgeRouteValues(
  BRIDGE_ROUTE_MAPPING_SYMBOLS
)
