import { NativeGasAddress, ZeroAddress } from '../constants'
import { BRIDGE_MAP } from '../constants/bridgeMap'

interface TokenInfo {
  symbol: string
  address: string
  chainId: string
}

interface BridgeRoutes {
  [key: string]: TokenInfo[]
}

const constructJSON = (
  swappableMap: typeof BRIDGE_MAP,
  exclusionList: string[]
): BridgeRoutes => {
  const result: BridgeRoutes = {}

  for (const chainA in swappableMap) {
    for (const addressA in swappableMap[chainA]) {
      const tokenA = swappableMap[chainA][addressA]
      const keyA = `${tokenA.symbol}-${chainA}`

      if (exclusionList.includes(keyA)) {
        continue
      }

      result[keyA] = []

      for (const chainB in swappableMap) {
        if (chainA !== chainB) {
          for (const addressB in swappableMap[chainB]) {
            const tokenB = swappableMap[chainB][addressB]
            const keyB = `${tokenB.symbol}-${chainB}`

            if (exclusionList.includes(keyB)) {
              continue
            }

            const canBridge = tokenA.origin.some(
              (bridgeSymbol) =>
                tokenB.destination.includes(bridgeSymbol) ||
                tokenB.origin.includes(bridgeSymbol)
            )

            if (canBridge) {
              result[keyA].push({
                symbol: tokenB.symbol,
                address: addressB === NativeGasAddress ? ZeroAddress : addressB,
                chainId: chainB,
              })
            }
          }
        }
      }

      if (result[keyA].length === 0) {
        delete result[keyA]
      }
    }
  }

  return result
}

export const BRIDGE_ROUTE_MAPPING: BridgeRoutes = constructJSON(BRIDGE_MAP, [])
