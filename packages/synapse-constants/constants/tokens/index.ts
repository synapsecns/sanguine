import _ from 'lodash'
import { Token } from '@types'

import * as CHAINS from '../chains/master'
import * as all from './bridgeable'
import * as allPool from './poolMaster'
import { GMX, ETH, USDC, USDT, WETH } from './bridgeable'
import { SYN_ETH_SUSHI_TOKEN } from './sushiMaster'

const allSwap = [WETH, USDC, USDT]

export * from './bridgeable'

// TODO change this to token by key
interface TokensByChain {
  [cID: string]: Token[]
}

interface TokenByKey {
  [cID: string]: Token
}

export const sortTokens = (tokens: Token[]) =>
  Object.values(tokens).sort((a, b) => b.visibilityRank - a.visibilityRank)

const sortedTokens = Object.values(all).sort(
  (a, b) => b.visibilityRank - a.visibilityRank
)

// Key value pairs here will override bridgeMap to hide particular chain-token pairs
export const PAUSED_TOKENS_BY_CHAIN = {
  [CHAINS.ETH.id]: ['WETH'],
  [CHAINS.OPTIMISM.id]: ['WETH'],
  [CHAINS.BOBA.id]: ['WETH'],
  [CHAINS.MOONBEAM.id]: ['WETH'],
  [CHAINS.BASE.id]: ['WETH'],
  [CHAINS.ARBITRUM.id]: ['WETH'],
  [CHAINS.FANTOM.id]: [],
  [CHAINS.DOGE.id]: ['BUSD', 'WETH'],
  [CHAINS.KLAYTN.id]: ['WETH'],
}

export const findChainIdsWithPausedToken = (routeSymbol: string) => {
  return _.reduce(
    PAUSED_TOKENS_BY_CHAIN,
    (result, tokens, chainId) => {
      if (_.includes(tokens, routeSymbol)) {
        result.push(chainId)
      }
      return result
    },
    []
  )
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

const getTokenHashMap = () => {
  const tokenHashMap = {}

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

export const tokenSymbolToToken = (chainId: number, symbol: string) => {
  if (chainId) {
    const token = BRIDGABLE_TOKENS[chainId].find((token) => {
      return token.symbol === symbol
    })
    return token
  }
}
export const tokenAddressToToken = (
  chainId: number,
  tokenAddress: string
): Token => {
  if (chainId) {
    if (tokenAddress === WETH.addresses[chainId]) {
      return WETH
    } else {
      const token = BRIDGABLE_TOKENS[chainId].find((token: Token) => {
        return token.addresses[chainId] === tokenAddress
      })
      return token
    }
  }
}

export const TOKEN_HASH_MAP = getTokenHashMap()

// SWAPS
const allTokensWithSwap = [...Object.values(all), ...Object.values(allSwap)]

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
