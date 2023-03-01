import { ChainId } from '@constants/networks'

import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/lp'

import { POOLS_BY_CHAIN } from '@constants/tokens/poolsByChain'

export const STAKABLE_TOKENS = {
  ...POOLS_BY_CHAIN,
  [ChainId.ETH]: [...POOLS_BY_CHAIN[ChainId.ETH], SYN_ETH_SUSHI_TOKEN],
}

// The numbers in staking maps are significant contract wise, important to leave as is
let STAKING_MAP_TOKENS = {}
for (const [chainId, arr] of Object.entries(STAKABLE_TOKENS)) {
  STAKING_MAP_TOKENS[chainId] = {}
  for (const token of arr) {
    STAKING_MAP_TOKENS[chainId][token.poolName] = token
  }
}

export { STAKING_MAP_TOKENS }
