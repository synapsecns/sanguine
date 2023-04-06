import { POOLS_BY_CHAIN } from '@constants/tokens'

const POOL_ROUTER_INDEX = {}
const POOL_INVERTED_ROUTER_INDEX = {}

for (const [chainId, arr] of Object.entries(POOLS_BY_CHAIN)) {
  POOL_INVERTED_ROUTER_INDEX[chainId] = {}
  for (const token of arr) {
    POOL_ROUTER_INDEX[token.routerIndex] = token.poolName
    POOL_INVERTED_ROUTER_INDEX[chainId][token.poolName] = token.routerIndex
  }
}

export { POOL_ROUTER_INDEX, POOL_INVERTED_ROUTER_INDEX }
