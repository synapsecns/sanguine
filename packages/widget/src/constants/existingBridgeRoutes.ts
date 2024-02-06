import { BRIDGE_MAP } from '@/constants/bridgeMap'
import { flattenPausedTokens } from '@/utils/flattenPausedTokens'

export type BridgeRoutes = Record<string, string[]>

const constructJSON = (swappableMap, exclusionList) => {
  const result = {}

  for (const chainA in swappableMap) {
    for (const tokenA in swappableMap[chainA]) {
      const symbolA = swappableMap[chainA][tokenA].symbol
      const key = `${symbolA}-${chainA}`

      if (exclusionList.includes(key)) continue

      for (const chainB in swappableMap) {
        if (chainA !== chainB) {
          for (const tokenB in swappableMap[chainB]) {
            const symbolB = swappableMap[chainB][tokenB].symbol
            const value = `${symbolB}-${chainB}`

            if (exclusionList.includes(value)) continue

            for (const bridgeSymbol of swappableMap[chainA][tokenA].origin) {
              if (
                swappableMap[chainA][tokenA].origin.includes(bridgeSymbol) &&
                swappableMap[chainB][tokenB].destination.includes(bridgeSymbol)
              ) {
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

export const EXISTING_BRIDGE_ROUTES: BridgeRoutes = constructJSON(
  BRIDGE_MAP,
  PAUSED_TOKENS
)
