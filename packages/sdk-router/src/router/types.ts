import { BigNumber } from '@ethersproject/bignumber'

import { BigintIsh } from '../constants'

/**
 * Matches DestRequest passed to SynapseRouter (V1) and SynapseCCTPRouter.
 */
export type DestRequest = {
  symbol: string
  amountIn: BigintIsh
}

export type PoolToken = { isWeth: boolean; token: string }

/**
 * Reduces the object to contain only the keys that are present in the PoolToken type.
 */
export const reduceToPoolToken = (poolToken: PoolToken): PoolToken => {
  return {
    isWeth: poolToken.isWeth,
    token: poolToken.token,
  }
}

export type PoolInfo = { tokens: BigNumber; lpToken: string }

export type Pool = { poolAddress: string; tokens: PoolToken[]; lpToken: string }
