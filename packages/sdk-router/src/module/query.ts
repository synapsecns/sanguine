import { BigNumber } from '@ethersproject/bignumber'
import { AddressZero } from '@ethersproject/constants'
import invariant from 'tiny-invariant'
import { XOR } from 'ts-xor'

import { ETH_NATIVE_TOKEN_ADDRESS } from '../utils/handleNativeToken'

/**
 * Matches SwapQuery passed to/returned by SynapseRouter (V1).
 */
export type RouterQuery = {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

/**
 * Matches SwapQuery passed to/returned by SynapseCCTPRouter.
 */
export type CCTPRouterQuery = {
  routerAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

/**
 * Matches SwapQuery passed to/returned by SynapseRouter (V1) or SynapseCCTPRouter, but not both.
 */
export type Query = XOR<RouterQuery, CCTPRouterQuery>

/**
 * Reduces the object to contain only the keys that are present in the Query type.
 */
export const reduceToQuery = (query: Query): Query => {
  // Extract the common properties.
  const common = {
    tokenOut: query.tokenOut,
    minAmountOut: query.minAmountOut,
    deadline: query.deadline,
    rawParams: query.rawParams,
  }
  // Extract the properties that are unique to each router.
  if (query.swapAdapter === undefined) {
    return {
      routerAdapter: query.routerAdapter,
      ...common,
    }
  } else {
    return {
      swapAdapter: query.swapAdapter,
      ...common,
    }
  }
}

/**
 * Narrows the query to the SynapseRouter (V1) type.
 *
 * @param query The query to narrow.
 * @returns The narrowed query object compatible with SynapseRouter (V1).
 * @throws If the query is not compatible with SynapseRouter (V1).
 */
export const narrowToRouterQuery = (query: Query): RouterQuery => {
  invariant(query.swapAdapter, 'swapAdapter is undefined')
  return query
}

/**
 * Narrows the query to the SynapseCCTPRouter type.
 *
 * @param query The query to narrow.
 * @returns The narrowed query object compatible with SynapseCCTPRouter.
 * @throws If the query is not compatible with SynapseCCTPRouter.
 */
export const narrowToCCTPRouterQuery = (query: Query): CCTPRouterQuery => {
  invariant(query.routerAdapter, 'routerAdapter is undefined')
  return query
}

/**
 * Checks if the query will lead to a complex bridge action, when used as the destination query.
 * Complex action is defined as an additional external call that the Bridge contract will have to perform
 * in order to complete the bridge action (e.g. mintAndSwap).
 *
 * @param destQuery The query to check.
 * @returns True if the query will lead to a complex bridge action, false otherwise.
 */
export const hasComplexBridgeAction = (destQuery: Query): boolean => {
  // Extract the adapter address: swapAdapter for SynapseRouter (V1), routerAdapter for SynapseCCTPRouter.
  const adapterAddress = destQuery.swapAdapter ?? destQuery.routerAdapter
  // Complex action will happen if the adapter address is not the zero address.
  // The tokenOut also needs to be different from ETH_ADDRESS, because WETH -> ETH is done in the Bridge contract,
  // without any additional external calls. We don't check tokenIn, as it could never be ETH_ADDRESS in destQuery,
  // because the Bridge contract does not natively work with ETH.
  return (
    adapterAddress !== AddressZero &&
    destQuery.tokenOut !== ETH_NATIVE_TOKEN_ADDRESS
  )
}

/**
 * Modifies the deadline of the query and returns the modified query.
 * Note: the original query is preserved unchanged.
 *
 * @param query - The query to modify.
 * @param deadline - The new deadline.
 * @returns The modified query with the new deadline.
 */
export const modifyDeadline = (query: Query, deadline: BigNumber): Query => {
  return {
    ...query,
    deadline,
  }
}

/**
 * Applies the slippage to the query's minAmountOut (rounded down), and returns the modified query
 * with the reduced minAmountOut.
 * Note: the original query is preserved unchanged.
 *
 * @param query - The query to modify.
 * @param slipNumerator - The numerator of the slippage.
 * @param slipDenominator - The denominator of the slippage.
 * @returns The modified query with the reduced minAmountOut.
 * @throws If the slippage fraction is invalid (<0, >1, or NaN)
 */
export const applySlippage = (
  query: Query,
  slipNumerator: number,
  slipDenominator: number
): Query => {
  invariant(slipDenominator > 0, 'Slippage denominator cannot be zero')
  invariant(slipNumerator >= 0, 'Slippage numerator cannot be negative')
  invariant(
    slipNumerator <= slipDenominator,
    'Slippage cannot be greater than 1'
  )
  const slippageAmount = query.minAmountOut
    .mul(slipNumerator)
    .div(slipDenominator)
  return {
    ...query,
    minAmountOut: query.minAmountOut.sub(slippageAmount),
  }
}

/**
 * Applies the slippage (in basis points) to the query's minAmountOut (rounded down), and returns the modified query
 * with the reduced minAmountOut.
 * Note: the original query is preserved unchanged.
 * Note: the slippage is applied as a fraction of 10000, e.g. 100 bips = 1%.
 *
 * @param query - The query to modify.
 * @param slipBasisPoints - The slippage in basis points.
 * @returns The modified query with the reduced minAmountOut.
 * @throws If the basis points are invalid (<0, >10000)
 */
export const applySlippageInBips = (
  query: Query,
  slipBasisPoints: number
): Query => {
  return applySlippage(query, slipBasisPoints, 10000)
}

/**
 * Creates a Query object for a no-swap bridge action.
 *
 * @param token - The token to bridge.
 * @param amount - The amount of token to bridge.
 * @returns The Query object for a no-swap bridge action.
 */
export const createNoSwapQuery = (token: string, amount: BigNumber): Query => {
  return {
    routerAdapter: AddressZero,
    tokenOut: token,
    minAmountOut: amount,
    deadline: BigNumber.from(0),
    rawParams: '0x',
  }
}
