import { Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'

import { Pool, PoolInfo, PoolToken } from '../router'
import { SynapseSDK } from '../sdk'

/**
 * Gets pool tokens for a pool address.
 *
 * @param chainId The chain ID
 * @param poolAddress The pool address
 * @returns The pool tokens
 */
export async function getPoolTokens(
  this: SynapseSDK,
  chainId: number,
  poolAddress: string
): Promise<PoolToken[]> {
  return this.synapseRouterSet
    .getSynapseRouter(chainId)
    .getPoolTokens(poolAddress)
}

/**
 * Gets info for a pool (number of tokens and LP token).
 *
 * @param chainId The chain ID
 * @param poolAddress The pool address
 * @returns The pool info (number of tokens and LP token)
 */
export async function getPoolInfo(
  this: SynapseSDK,
  chainId: number,
  poolAddress: string
): Promise<PoolInfo> {
  return this.synapseRouterSet
    .getSynapseRouter(chainId)
    .getPoolInfo(poolAddress)
}

/**
 * Gets all pools for a chain ID.
 *
 * @param chainId The chain ID
 * @returns An array of all pools (address, tokens, LP token)
 */
export async function getAllPools(
  this: SynapseSDK,
  chainId: number
): Promise<Pool[]> {
  return this.synapseRouterSet.getSynapseRouter(chainId).getAllPools()
}

/**
 * Calculates the amount required to add liquidity for amounts of each token.
 *
 * @param chainId The chain ID
 * @param poolAddress The pool address
 * @param amounts The amounts of each token to add
 * @returns The amount of LP tokens needed and router address
 */
export async function calculateAddLiquidity(
  this: SynapseSDK,
  chainId: number,
  poolAddress: string,
  amounts: Record<string, BigNumber>
): Promise<{ amount: BigNumber; routerAddress: string }> {
  const router = this.synapseRouterSet.getSynapseRouter(chainId)
  const routerAddress = router.routerContract.address
  const poolTokens = await router.getPoolTokens(poolAddress)
  // Populate amounts array by combining amounts map and pool tokens, preserving tokens order
  // and adding 0 for tokens not in the map
  const amountsArr: BigNumber[] = []
  poolTokens.map((token) => {
    amountsArr.push(amounts[token.token] ?? Zero)
  })
  // Don't do a contract call if all amounts are 0
  const amount = amountsArr.every((value) => value.isZero())
    ? Zero
    : await router.routerContract.calculateAddLiquidity(poolAddress, amountsArr)
  return { amount, routerAddress }
}

/**
 * Calculates the amounts received when removing liquidity.
 *
 * @param chainId The chain ID
 * @param poolAddress The pool address
 * @param amount The amount of LP tokens to remove
 * @returns The amounts of each token received and router address
 */
export async function calculateRemoveLiquidity(
  this: SynapseSDK,
  chainId: number,
  poolAddress: string,
  amount: BigNumber
): Promise<{
  amounts: Array<{ value: BigNumber; index: number }>
  routerAddress: string
}> {
  const router = this.synapseRouterSet.getSynapseRouter(chainId)
  const routerAddress = router.routerContract.address
  const amountsOut = await router.routerContract.calculateRemoveLiquidity(
    poolAddress,
    amount
  )
  // Zip amounts with token indexes
  const amounts = amountsOut.map((value, index) => ({ value, index }))
  return { amounts, routerAddress }
}

/**
 * Calculates the amount of one token received when removing liquidity.
 *
 * @param chainId The chain ID
 * @param poolAddress The pool address
 * @param amount The amount of LP tokens to remove
 * @param poolIndex The index of the token to receive
 * @returns The amount received and router address
 */
export async function calculateRemoveLiquidityOne(
  this: SynapseSDK,
  chainId: number,
  poolAddress: string,
  amount: BigNumber,
  poolIndex: number
): Promise<{
  amount: { value: BigNumber; index: number }
  routerAddress: string
}> {
  const router = this.synapseRouterSet.getSynapseRouter(chainId)
  const routerAddress = router.routerContract.address
  const amountOut = {
    value: await router.routerContract.calculateWithdrawOneToken(
      poolAddress,
      amount,
      poolIndex
    ),
    index: poolIndex,
  }
  return { amount: amountOut, routerAddress }
}
