import * as all from './masterTokenList'
import { Token } from '@utils/classes/Token'

interface BridgeableTokens {
  [cID: string]: Token[]
}

interface BridgeChainsByType {
  [swapableType: string]: string[]
}

interface BridgeTypeByChain {
  [cID: string]: string[]
}

interface BridgeSwapableTokensByType {
  [cID: string]: {
    [swapableType: string]: Token[]
  }
}

// generates an object where the keys are "swapType" and the values are the chains that support that swapType

const getBridgeableTokens = (): BridgeableTokens => {
  let bridgeableTokens: BridgeableTokens = {}
  Object.values(all).map((token) => {
    for (const cID of Object.keys(token.addresses)) {
      if (!bridgeableTokens[cID]) {
        bridgeableTokens[cID] = [token]
      } else {
        if (!bridgeableTokens[cID]?.includes(token)) {
          bridgeableTokens[cID] = [...bridgeableTokens[cID], token]
        }
      }
    }
  })
  return bridgeableTokens
}

const getBridgeChainsByType = (): BridgeChainsByType => {
  let bridgeChainsByType: BridgeChainsByType = {}
  Object.values(all).map((token) => {
    let swapableType = String(token?.swapableType)
    let keys = Object.keys(token.addresses).filter(
      (a) => !bridgeChainsByType[swapableType]?.includes(a)
    )
    if (bridgeChainsByType[swapableType]) {
      bridgeChainsByType[swapableType] = [
        ...bridgeChainsByType[swapableType],
        ...keys,
      ]
    } else {
      bridgeChainsByType[swapableType] = keys
    }
  })
  return bridgeChainsByType
}
const getBridgeTypeByChain = (): BridgeTypeByChain => {
  let bridgeChainByType = getBridgeChainsByType()
  let bridgeTypeByChain: BridgeTypeByChain = {}
  Object.keys(bridgeChainByType).forEach((key) => {
    bridgeChainByType[key].forEach((value) => {
      if (bridgeTypeByChain[value]) {
        bridgeTypeByChain[value].push(key)
      } else {
        bridgeTypeByChain[value] = [key]
      }
    })
  })
  return bridgeTypeByChain
}

const convertArrayToObject = (array: any) => {
  return array.reduce((obj: any, value: any) => {
    obj[value] = []
    return obj
  }, {})
}

const getBridgeSwapableTokensByType = (): BridgeSwapableTokensByType => {
  let bridgeTypeByChain = getBridgeTypeByChain()
  let bridgeSwapableTokensByType = Object.fromEntries(
    Object.entries(bridgeTypeByChain).map(([k, v]) => [
      k,
      convertArrayToObject(v),
    ])
  )
  console.log('GG', bridgeSwapableTokensByType)
  Object.values(all).map((token) => {
    let swapableType = String(token?.swapableType)

    for (const cID of Object.keys(token.addresses)) {
      if (bridgeSwapableTokensByType[cID][swapableType].length == 0) {
        bridgeSwapableTokensByType[cID][swapableType] = [token]
      } else {
        if (!bridgeSwapableTokensByType[cID][swapableType]?.includes(token)) {
          bridgeSwapableTokensByType[cID][swapableType] = [
            ...bridgeSwapableTokensByType[cID][swapableType],
            token,
          ]
        }
      }
    }
  })

  return bridgeSwapableTokensByType
}

// const getBridgeSwapableTokensByChain = () => {
//   let bridgeSwapableTokensByChain = {}
//   for (const [coinSymbol, chainIdArr] of _.entries(BRIDGE_CHAINS_BY_TYPE)) {
//     for (const cID of chainIdArr) {
//       if (bridgeSwapableTokensByChain[cID]) {
//         bridgeSwapableTokensByChain[cID].push(coinSymbol)
//       } else {
//         bridgeSwapableTokensByChain[cID] = [coinSymbol]
//       }
//     }
//   }
//   return bridgeSwapableTokensByChain
// }
export const BRIDGABLE_TOKENS = getBridgeableTokens()
export const BRIDGE_CHAINS_BY_TYPE = getBridgeChainsByType()
export const BRIDGE_TYPES_BY_CHAIN = getBridgeTypeByChain()

export const BRIDGE_SWAPABLE_TOKENS_BY_TYPE = getBridgeSwapableTokensByType()
// export const BRIDGE_SWAPABLE_TOKENS_BY_CHAIN = getBridgeSwapableTokensByChain()

// Get bridgeable
