import { RFQ_MAP } from './rfqMap'
import { transformRFQMap } from './existingRfqRoutes'
import { flattenPausedTokens } from '@/utils/flattenPausedTokens'
import { BRIDGE_MAP } from '@/constants/bridgeMap'

export type BridgeRoutes = Record<string, string[]>

const constructJSON = (swappableMap, exclusionList) => {
  const result = {}

  // Iterate through the chains
  for (const chainA in swappableMap) {
    for (const tokenA in swappableMap[chainA]) {
      const symbolA = swappableMap[chainA][tokenA].symbol
      const key = `${symbolA}-${chainA}`

      if (exclusionList.includes(key)) continue

      // Iterate through other chains to compare
      for (const chainB in swappableMap) {
        if (chainA !== chainB) {
          for (const tokenB in swappableMap[chainB]) {
            const symbolB = swappableMap[chainB][tokenB].symbol
            const value = `${symbolB}-${chainB}`

            if (exclusionList.includes(value)) continue

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

const PAUSED_TOKENS = flattenPausedTokens()

const EXISTING_BRIDGE_CCTP_ROUTES: BridgeRoutes = constructJSON(
  BRIDGE_MAP,
  PAUSED_TOKENS
)

const RFQ_ROUTES = transformRFQMap(RFQ_MAP)

const mergeObjectsUnique = (firstObj, secondObj) => {
  Object.keys(secondObj).forEach((key) => {
    if (!firstObj[key]) {
      firstObj[key] = secondObj[key]
    } else {
      secondObj[key].forEach((value) => {
        if (!firstObj[key].includes(value)) {
          firstObj[key].push(value)
        }
      })
    }
  })

  return firstObj
}

export const EXISTING_BRIDGE_ROUTES: BridgeRoutes = mergeObjectsUnique(
  EXISTING_BRIDGE_CCTP_ROUTES,
  RFQ_ROUTES
)
