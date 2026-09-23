import _ from 'lodash'
import { RELAY_SUPPORTED_CHAIN_IDS } from '@synapsecns/sdk-router'

import { BRIDGE_MAP } from '@/constants/bridgeMap'
import { flattenPausedTokens } from '@/utils/flattenPausedTokens'
import { ETH, HYPEREVM, HYPERLIQUID } from './chains/master'

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

const addUSDCHyperLiquid = (routes: BridgeRoutes): BridgeRoutes => {
  const usdcHyperliquid = `USDC-${HYPERLIQUID.id}`

  return _.mapValues(routes, (innerList, key) => {
    const originChainId = Number(key.slice(key.lastIndexOf('-') + 1))
    if (RELAY_SUPPORTED_CHAIN_IDS.includes(originChainId)) {
      return [...new Set([...innerList, usdcHyperliquid])]
    }
    return innerList
  })
}
const PAUSED_TOKENS = flattenPausedTokens()

const bridgeRoutes: BridgeRoutes = constructJSON(BRIDGE_MAP, PAUSED_TOKENS)

// Separate from the legacy SYN graph, which would advertise every SYN chain.
const synRoutes: BridgeRoutes = {
  [`SYN-${ETH.id}`]: [`SYN-${HYPEREVM.id}`, `SYN-${HYPERLIQUID.id}`],
  [`SYN-${HYPEREVM.id}`]: [`SYN-${ETH.id}`],
}
Object.entries(synRoutes).forEach(([origin, destinations]) => {
  if (PAUSED_TOKENS.includes(origin)) return
  bridgeRoutes[origin] = [
    ...new Set([
      ...(bridgeRoutes[origin] ?? []),
      ...destinations.filter(
        (destination) => !PAUSED_TOKENS.includes(destination)
      ),
    ]),
  ]
})

export const EXISTING_BRIDGE_ROUTES: BridgeRoutes = addUSDCHyperLiquid(
  bridgeRoutes
)
