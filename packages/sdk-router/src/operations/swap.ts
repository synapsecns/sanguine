import { PopulatedTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'

import { BigintIsh } from '../constants'
import { Query, SwapQuote } from '../module'
import { handleNativeToken } from '../utils/handleNativeToken'
import { SynapseSDK } from '../sdk'
import { TEN_MINUTES, applyOptionalDeadline } from '../utils/deadlines'

/**
 * Performs a swap through a Synapse Router.
 *
 * @param chainId The chain ID
 * @param to The recipient address
 * @param token The token to swap
 * @param amount The swap amount
 * @param query The swap quote query
 * @returns A populated transaction to perform the swap
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
 */
export async function swap(
  this: SynapseSDK,
  chainId: number,
  to: string,
  token: string,
  amount: BigintIsh,
  query: Query
): Promise<PopulatedTransaction> {
  token = handleNativeToken(token)
  return this.synapseRouterSet
    .getSynapseRouter(chainId)
    .swap(to, token, amount, query)
}

/**
 * Gets a swap quote from a Synapse Router.
 *
 * @param chainId The chain ID
 * @param tokenIn The input token
 * @param tokenOut The output token
 * @param amountIn The input amount
 * @param deadline The deadline to use for the swap. Optional, will default to 10 minutes from now.
 * @returns The swap quote (query, max amount out, and router address)
 * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
 */
export async function swapQuote(
  this: SynapseSDK,
  chainId: number,
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  deadline?: BigNumber
): Promise<SwapQuote> {
  tokenOut = handleNativeToken(tokenOut)
  tokenIn = handleNativeToken(tokenIn)
  // Get SynapseRouter instance for given chain
  const router = this.synapseRouterSet.getSynapseRouter(chainId)
  const routerAddress = router.address
  const query = await router.getAmountOut(tokenIn, tokenOut, amountIn)
  const maxAmountOut = query.minAmountOut
  if (query.minAmountOut.isZero()) {
    throw Error('No queries found for this route')
  }
  query.deadline = applyOptionalDeadline(deadline, TEN_MINUTES)
  return { routerAddress, maxAmountOut, query }
}
