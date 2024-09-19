import { BRIDGE_MAP } from '../constants/bridgeMap'

const constructJSON = (swappableMap, exclusionList) => {
  const result = {}

  // Iterate through the origin chains
  for (const originChainId in swappableMap) {
    for (const originTokenAddress in swappableMap[originChainId]) {
      const originToken = swappableMap[originChainId][originTokenAddress]
      const originKey = `${originToken.symbol}-${originChainId}`

      if (exclusionList.includes(originKey)) {
        continue
      }

      // Initialize origin chain and origin token with symbol if not existing
      if (!result[originChainId]) {
        result[originChainId] = {}
      }

      if (!result[originChainId][originTokenAddress]) {
        result[originChainId][originTokenAddress] = {
          symbol: originToken.symbol,
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

          if (exclusionList.includes(destinationKey)) {
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
                !result[originChainId][originTokenAddress].routes[
                  destinationChainId
                ]
              ) {
                result[originChainId][originTokenAddress].routes[
                  destinationChainId
                ] = {}
              }

              result[originChainId][originTokenAddress].routes[
                destinationChainId
              ][destinationTokenAddress] = {
                symbol: destinationToken.symbol,
                minValue: null,
                maxValue: null,
              }
            }
          }
        }
      }
    }
  }

  return result
}

export const BRIDGE_LIMIT_MAPPING = constructJSON(BRIDGE_MAP, [])
