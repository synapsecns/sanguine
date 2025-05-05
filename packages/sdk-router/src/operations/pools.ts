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
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
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
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
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
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
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
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
 */
export async function calculateAddLiquidity(
  this: SynapseSDK,
  chainId: number,
  poolAddress: string,
  amounts: Record<string, BigNumber>
): Promise<{ amount: BigNumber; routerAddress: string }> {
  // TODO (Chi): use amounts array as the input instead of map in the first place
  const router = this.synapseRouterSet.getSynapseRouter(chainId)
  const routerAddress = router.address
  const poolTokens = await router.getPoolTokens(poolAddress)
  // Create a map that uses lowercase token addresses as keys
  const lowerCaseAmounts: Record<string, BigNumber> = {}
  Object.keys(amounts).forEach((key) => {
    lowerCaseAmounts[key.toLowerCase()] = amounts[key]
  })
  // Populate amounts array by combining amounts map and pool tokens, preserving tokens order
  // and adding 0 for tokens not in the map
  const amountsArray: BigNumber[] = poolTokens.map(
    (poolToken) => lowerCaseAmounts[poolToken.token.toLowerCase()] ?? Zero
  )
  // Don't do a contract call if all amounts are 0
  const amount = amountsArray.every((value) => value.isZero())
    ? Zero
    : await router.calculateAddLiquidity(poolAddress, amountsArray)
  return { amount, routerAddress }
}

/**
 * Calculates the amounts received when removing liquidity.
 *
 * @param chainId The chain ID
 * @param poolAddress The pool address
 * @param amount The amount of LP tokens to remove
 * @returns The amounts of each token received and router address
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
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
  const routerAddress = router.address
  const amountsOut = await router.calculateRemoveLiquidity(poolAddress, amount)
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
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
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
  const routerAddress = router.address
  const amountOut = {
    value: await router.calculateWithdrawOneToken(
      poolAddress,
      amount,
      poolIndex
    ),
    index: poolIndex,
  }
  return { amount: amountOut, routerAddress }
}
