import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { XOR } from 'ts-xor'

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
