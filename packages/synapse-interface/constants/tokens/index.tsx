import * as CHAINS from '@constants/chains/master'
import * as all from './master'
import * as allPool from './poolMaster'
import * as allSwap from './swapMaster'
import { GMX, ETH } from './master'
import { WETH } from './swapMaster'
import { SYN_ETH_SUSHI_TOKEN } from './sushiMaster'
import { Token } from '@/utils/types'
import _ from 'lodash'

// TODO change this to token by key
interface TokensByChain {
  [cID: string]: Token[]
}

interface TokenByKey {
  [cID: string]: Token
}
interface BridgeChainsByType {
  [swapableType: string]: string[]
}

interface BridgeTypeByChain {
  [cID: string]: string[]
}

interface SwapableTokensByType {
  [cID: string]: {
    [swapableType: string]: Token[]
  }
}
export const sortTokens = (tokens: Token[]) =>
  Object.values(tokens).sort((a, b) => b.visibilityRank - a.visibilityRank)

const sortedTokens = Object.values(all).sort(
  (a, b) => b.visibilityRank - a.visibilityRank
)

// This should be an object where keys are chain IDs and values are arrays of token keys that you want to pause on each chain
const PAUSED_TOKENS_BY_CHAIN = {
  [CHAINS.FANTOM.id]: ['USDC', 'USDT', 'FTMETH'],
  [CHAINS.AVALANCHE.id]: ['AVWETH'],
}

const getBridgeableTokens = (): TokensByChain => {
  const bridgeableTokens: TokensByChain = {}
  Object.entries(all).map(([key, token]) => {
    for (const cID of Object.keys(token.addresses)) {
      // Skip if the token is paused on the current chain
      if (PAUSED_TOKENS_BY_CHAIN[cID]?.includes(key)) continue

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
  const bridgeChainsByType: BridgeChainsByType = {}
  Object.entries(all).map(([key, token]) => {
    const swapableType = String(token?.swapableType)
    const keys = Object.keys(token.addresses).filter((cID) => {
      // Skip if the token is paused on the current chain
      if (PAUSED_TOKENS_BY_CHAIN[cID]?.includes(key)) return false

      return !bridgeChainsByType[swapableType]?.includes(cID)
    })

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
  const bridgeChainByType = getBridgeChainsByType()
  const bridgeTypeByChain: BridgeTypeByChain = {}
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

const getBridgeableTokensByType = (): SwapableTokensByType => {
  const bridgeTypeByChain = getBridgeTypeByChain()
  const bridgeSwapableTokensByType = Object.fromEntries(
    Object.entries(bridgeTypeByChain).map(([k, v]) => [
      k,
      convertArrayToObject(v),
    ])
  )

  Object.entries(all).map(([key, token]) => {
    const swapableType = String(token?.swapableType)

    for (const cID of Object.keys(token.addresses)) {
      // Skip if the token is paused on the current chain
      if (PAUSED_TOKENS_BY_CHAIN[cID]?.includes(key)) continue

      if (bridgeSwapableTokensByType[cID][swapableType].length === 0) {
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

const getTokenHashMap = () => {
  let tokenHashMap = {}

  for (const [chainId, tokensOnChain] of _.toPairs(BRIDGABLE_TOKENS)) {
    tokenHashMap[chainId] = {}
    for (const token of tokensOnChain) {
      tokenHashMap[chainId][token.addresses[chainId]] = token
    }
  }

  tokenHashMap[CHAINS.AVALANCHE.id][GMX.wrapperAddresses[CHAINS.AVALANCHE.id]] =
    GMX
  Object.keys(WETH.addresses).map((chain) => {
    tokenHashMap[chain][WETH.addresses[chain]] = ETH
  })
  return tokenHashMap
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
export const BRIDGE_SWAPABLE_TOKENS_BY_TYPE = getBridgeableTokensByType()
export const tokenSymbolToToken = (chainId: number, symbol: string) => {
  if (chainId) {
    const token = BRIDGABLE_TOKENS[chainId].find((token) => {
      return token.symbol === symbol
    })
    return token
  }
}
export const TOKEN_HASH_MAP = getTokenHashMap()

// SWAPS
const allTokensWithSwap = [...Object.values(all), ...Object.values(allSwap)]
const getSwapableTokens = (): TokensByChain => {
  const swapTokens: TokensByChain = {}
  allTokensWithSwap.map((token) => {
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

const getSwapableTokensByType = (): SwapableTokensByType => {
  const swapTokens: SwapableTokensByType = {}
  allTokensWithSwap.map((token) => {
    if (!(token?.swapableOn?.length > 0)) return
    for (const cID of token.swapableOn) {
      if (!swapTokens[cID]) {
        swapTokens[cID] = { [token.swapableType]: [token] }
      } else if (!swapTokens[cID][token.swapableType]) {
        swapTokens[cID][token.swapableType] = [token]
      } else if (!swapTokens[cID][token.swapableType].includes(token)) {
        swapTokens[cID][token.swapableType] = [
          ...swapTokens[cID][token.swapableType],
          token,
        ]
      }
    }
  })
  return swapTokens
}
const getSwapPriorityRanking = () => {
  const swapPriorityRanking = {}
  allTokensWithSwap.map((token) => {
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
export const SWAPABLE_TOKENS_BY_TYPE = getSwapableTokensByType()
export const POOL_PRIORITY_RANKING = getSwapPriorityRanking()

// POOLS
const getPoolsByChain = (displayOnly: boolean): TokensByChain => {
  const poolTokens: TokensByChain = {}
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
  const CHAINS_BY_POOL_NAME = {}
  const poolsByChain = getPoolsByChain(false)
  Object.keys(poolsByChain).map((chainId) => {
    for (const swapToken of poolsByChain[chainId]) {
      CHAINS_BY_POOL_NAME[swapToken.poolName] = chainId
    }
  })
  return CHAINS_BY_POOL_NAME
}

const getTokensByPoolTypeByChain = (type: string) => {
  const poolTokens: TokensByChain = {}
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
  const poolTokens: TokensByChain = {}
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

const getPoolByRouterIndex = () => {
  const poolTokens: TokenByKey = {}
  Object.values(allPool).map((token) => {
    poolTokens[token.routerIndex] = token
  })
  return poolTokens
}

export const POOL_CHAINS_BY_NAME = getChainsByPoolName()
export const POOL_BY_ROUTER_INDEX = getPoolByRouterIndex()

export const POOLS_BY_CHAIN = getPoolsByChain(false)
export const DISPLAY_POOLS_BY_CHAIN = getPoolsByChain(true)
export const USD_POOLS_BY_CHAIN = getTokensByPoolTypeByChain('USD')
export const ETH_POOLS_BY_CHAIN = getTokensByPoolTypeByChain('ETH')
export const LEGACY_POOLS_BY_CHAIN = getLegacyTokensByChain()

export const STAKABLE_TOKENS = {
  ...getPoolsByChain(false),
  [CHAINS.ETH.id]: [SYN_ETH_SUSHI_TOKEN],
}

const getStakingMap = () => {
  const STAKING_MAP_TOKENS = {}
  Object.keys(POOLS_BY_CHAIN).map((chainId) => {
    if (!STAKING_MAP_TOKENS[chainId]) {
      STAKING_MAP_TOKENS[chainId] = {}
    }

    STAKING_MAP_TOKENS[chainId] = {}
    for (const token of STAKABLE_TOKENS[chainId]) {
      STAKING_MAP_TOKENS[chainId][token.poolName] = token
    }
  })
  return STAKING_MAP_TOKENS
}
export const STAKING_MAP_TOKENS = getStakingMap()
