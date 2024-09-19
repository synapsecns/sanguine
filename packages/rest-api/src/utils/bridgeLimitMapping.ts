import { BRIDGE_MAP } from '../constants/bridgeMap'
import * as ALL_TOKENS from '../constants/bridgeable'

const constructJSON = (swappableMap, exclusionList) => {
  const result = {}

  // Iterate through the origin chains
  for (const originChainId in swappableMap) {
    for (const originTokenAddress in swappableMap[originChainId]) {
      const originToken = swappableMap[originChainId][originTokenAddress]
      const originKey = `${originToken.symbol}-${originChainId}`

      // Use transformPair to get token object
      const transformedOriginToken = transformPair(originKey)

      if (!transformedOriginToken || exclusionList.includes(originKey)) {
        continue
      }

      // Initialize origin chain and origin token with symbol and swappableType if not existing
      if (!result[originChainId]) {
        result[originChainId] = {}
      }

      if (!result[originChainId][transformedOriginToken.address]) {
        result[originChainId][transformedOriginToken.address] = {
          symbol: transformedOriginToken.symbol,
          swappableType: transformedOriginToken.swapableType, // Fetch swappableType
          routes: {},
        }
      }

      // Iterate through destination chains
      for (const destinationChainId in swappableMap) {
        if (originChainId === destinationChainId) {
          continue
        }

        for (const destinationTokenAddress in swappableMap[
          destinationChainId
        ]) {
          const destinationToken =
            swappableMap[destinationChainId][destinationTokenAddress]
          const destinationKey = `${destinationToken.symbol}-${destinationChainId}`

          // Use transformPair for destination token as well
          const transformedDestinationToken = transformPair(destinationKey)

          if (
            !transformedDestinationToken ||
            exclusionList.includes(destinationKey)
          ) {
            continue
          }

          // Check for bridge compatibility by comparing origin and destination symbols
          for (const bridgeSymbol of originToken.origin) {
            if (
              originToken.origin.includes(bridgeSymbol) &&
              destinationToken.destination.includes(bridgeSymbol)
            ) {
              // Initialize destination token with symbol, minValue, maxValue if not existing
              if (
                !result[originChainId][transformedOriginToken.address].routes[
                  destinationChainId
                ]
              ) {
                result[originChainId][transformedOriginToken.address].routes[
                  destinationChainId
                ] = {}
              }

              result[originChainId][transformedOriginToken.address].routes[
                destinationChainId
              ][transformedDestinationToken.address] = {
                symbol: transformedDestinationToken.symbol,
                minOriginValue: null,
                maxOriginValue: null,
              }
            }
          }
        }
      }
    }
  }

  return result
}

export const transformPair = (string: string): any => {
  const [symbol, chainId] = string.split('-')
  const token = Object.values(ALL_TOKENS).find((t) => t.routeSymbol === symbol)
  const address = token?.addresses[chainId]
  if (token && address) {
    return {
      symbol,
      chainId,
      address,
      swapableType: token.swapableType,
    }
  }
}

export const BRIDGE_LIMIT_MAPPING = constructJSON(BRIDGE_MAP, [])
