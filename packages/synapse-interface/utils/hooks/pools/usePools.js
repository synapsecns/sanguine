
import _ from 'lodash'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useTokenInfo } from '@hooks/tokens/useTokenInfo'

import {
  POOLS_BY_CHAIN
} from '@constants/tokens/poolsByChain'
import { ChainId } from '@constants/networks'
import { HARMONY_JEWEL_SWAP_TOKEN } from '@constants/tokens/jewelswap'


let POOLS_MAP = {}
let POOL_NAME_TOKEN_MAP = {}
for (const [chainId, arr] of Object.entries(POOLS_BY_CHAIN)) {
  POOLS_MAP[chainId] = {}
  POOL_NAME_TOKEN_MAP[chainId] = {}
  for (const token of arr) {
    POOLS_MAP[chainId][token.poolName] = token.poolTokens
    POOL_NAME_TOKEN_MAP[chainId][token.poolName] = token
  }
}
// hackfix for jewel
POOL_NAME_TOKEN_MAP[ChainId.HARMONY][HARMONY_JEWEL_SWAP_TOKEN.poolName] = HARMONY_JEWEL_SWAP_TOKEN

export { POOL_NAME_TOKEN_MAP, POOLS_MAP }








/**
 * @param {string} poolName
 * @return {Token[]}
 */
export function usePool(poolName, otherChainId) {
  const { chainId: activeChainId } = useActiveWeb3React()
  const chainId = otherChainId ?? activeChainId
  return POOLS_MAP[chainId][poolName] ?? []
}





export function usePoolToken(poolName) {
  const { chainId } = useActiveWeb3React()
  return POOL_NAME_TOKEN_MAP[chainId][poolName]
}

export function usePoolTokenInfo(poolName) {
  const { chainId } = useActiveWeb3React()
  const tokenInfo = useTokenInfo(POOL_NAME_TOKEN_MAP[chainId][poolName])
  return tokenInfo
}


