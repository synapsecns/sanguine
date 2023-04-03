import * as all from './master'
import * as allPool from './poolMaster'
import { SYN_ETH_SUSHI_TOKEN } from './sushiMaster'

import * as CHAINS from '@constants/chains/master'
import { Token } from '@/utils/types'

interface TokensByChain {
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

const sortedTokens = Object.values(all).sort(
  (a, b) => b.visibilityRank - a.visibilityRank
)

const getBridgeableTokens = (): TokensByChain => {
  let bridgeableTokens: TokensByChain = {}
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
  Object.values(all).map((token) => {
    let swapableType = String(token?.swapableType)

    for (const cID of Object.keys(token.addresses)) {
      if (bridgeSwapableTokensByType[cID][swapableType].length == 0) {
        bridgeSwapableTokensByType[cID][swapableType] = [token]
      } else if (
        !bridgeSwapableTokensByType[cID][swapableType]?.includes(token)
      ) {
        bridgeSwapableTokensByType[cID][swapableType] = [
          ...bridgeSwapableTokensByType[cID][swapableType],
          token,
        ]
      }
    }
  })

  return bridgeSwapableTokensByType
}

export const TOKENS_SORTED_BY_SWAPABLETYPE = Array.from(
  new Set(sortedTokens.map((token) => token.swapableType))
)
export const TOKENS_SORTED_BY_SYMBOL = Array.from(
  new Set(sortedTokens.map((token) => token.symbol))
)
export const BRIDGABLE_TOKENS = getBridgeableTokens()
export const BRIDGE_CHAINS_BY_TYPE = getBridgeChainsByType()
export const BRIDGE_TYPES_BY_CHAIN = getBridgeTypeByChain()
export const BRIDGE_SWAPABLE_TOKENS_BY_TYPE = getBridgeSwapableTokensByType()
export const tokenSymbolToToken = (chainId: number, symbol: string) => {
  const token = BRIDGABLE_TOKENS[chainId].find((token) => {
    return token.symbol === symbol
  })
  return token
}

// SWAPS
const getSwapableTokens = (): TokensByChain => {
  let swapTokens: TokensByChain = {}
  Object.values(all).map((token) => {
    if (!(token?.swapableOn?.length > 0)) return
    for (const cID of token.swapableOn) {
      if (!swapTokens[cID]) {
        swapTokens[cID] = [token]
      } else if (!swapTokens[cID]?.includes(token)) {
        swapTokens[cID] = [...swapTokens[cID], token]
      }
    }
  })
  return swapTokens
}
const getSwapPriorityRanking = () => {
  let swapPriorityRanking = {}
  Object.values(allPool).map((token) => {
    if (!token.priorityPool) return
    for (const cID of Object.keys(token.addresses)) {
      if (!swapPriorityRanking[cID]) {
        swapPriorityRanking[cID] = {}
      }
      for (const poolToken of token.poolTokens) {
        swapPriorityRanking[cID][poolToken.symbol] = token
      }
    }
  })
  return swapPriorityRanking
}
export const SWAPABLE_TOKENS = getSwapableTokens()
export const POOL_PRIORITY_RANKING = getSwapPriorityRanking()

// POOLS
const getPoolsByChain = (displayOnly: boolean): TokensByChain => {
  let poolTokens: TokensByChain = {}
  Object.values(allPool).map((token) => {
    if (displayOnly && !token.display) return
    for (const cID of Object.keys(token.addresses)) {
      if (!poolTokens[cID]) {
        poolTokens[cID] = [token]
      } else {
        if (!poolTokens[cID]?.includes(token)) {
          poolTokens[cID] = [...poolTokens[cID], token]
        }
      }
    }
  })
  return poolTokens
}

const getChainsByPoolName = () => {
  let CHAINS_BY_POOL_NAME = {}
  const poolsByChain = getPoolsByChain(false)
  Object.keys(poolsByChain).map((chainId) => {
    for (const swapToken of poolsByChain[chainId]) {
      CHAINS_BY_POOL_NAME[swapToken.poolName] = chainId
    }
  })
  return CHAINS_BY_POOL_NAME
}

const getTokensByPoolTypeByChain = (type: string) => {
  let poolTokens: TokensByChain = {}
  Object.values(allPool).map((token) => {
    if (!token.display || !token?.poolType?.includes(type)) return
    for (const cID of Object.keys(token.addresses)) {
      if (!poolTokens[cID]) {
        poolTokens[cID] = [token]
      } else {
        if (!poolTokens[cID]?.includes(token)) {
          poolTokens[cID] = [...poolTokens[cID], token]
        }
      }
    }
  })
  return poolTokens
}

const getLegacyTokensByChain = () => {
  let poolTokens: TokensByChain = {}
  Object.values(allPool).map((token) => {
    if (!token.legacy) return
    for (const cID of Object.keys(token.addresses)) {
      if (!poolTokens[cID]) {
        poolTokens[cID] = [token]
      } else {
        if (!poolTokens[cID]?.includes(token)) {
          poolTokens[cID] = [...poolTokens[cID], token]
        }
      }
    }
  })
  return poolTokens
}

export const POOL_CHAINS_BY_NAME = getChainsByPoolName()
export const POOLS_BY_CHAIN = getPoolsByChain(false)
export const DISPLAY_POOLS_BY_CHAIN = getPoolsByChain(true)
export const USD_POOLS_BY_CHAIN = getTokensByPoolTypeByChain('USD')
export const ETH_POOLS_BY_CHAIN = getTokensByPoolTypeByChain('ETH')
export const LEGACY_POOLS_BY_CHAIN = getLegacyTokensByChain()

export const STAKABLE_TOKENS = {
  ...getChainsByPoolName(),
  [CHAINS.ETH.id]: [...POOLS_BY_CHAIN[CHAINS.ETH.id], SYN_ETH_SUSHI_TOKEN],
}

const getStakingMap = () => {
  let STAKING_MAP_TOKENS = {}
  Object.keys(STAKABLE_TOKENS).map((chainId) => {
    STAKING_MAP_TOKENS[chainId] = {}
    for (const token of STAKABLE_TOKENS[chainId]) {
      STAKING_MAP_TOKENS[chainId][token.poolName] = token
    }
  })
  return STAKING_MAP_TOKENS
}
export const STAKING_MAP_TOKENS = getStakingMap()
