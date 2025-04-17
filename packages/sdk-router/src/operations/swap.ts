import { BigNumber } from '@ethersproject/bignumber'
import { AddressZero } from '@ethersproject/constants'
import { PopulatedTransaction } from '@ethersproject/contracts'
import { BigNumberish } from 'ethers'
import { uuidv7 } from 'uuidv7'

import { areIntentsSupported, MEDIAN_TIME_BLOCK } from '../constants'
import { Query, applySlippageToQuery, applyDeadlineToQuery } from '../module'
import { SynapseSDK } from '../sdk'
import {
  RecipientEntity,
  RouteInput,
  slippageFromPercentage,
  USER_SIMULATED_ADDRESS,
} from '../swap'
import { SwapQuote, SwapQuoteV2, SwapV2Parameters } from '../types'
import {
  handleNativeToken,
  TEN_MINUTES,
  applyOptionalDeadline,
  calculateDeadline,
  isSameAddress,
  stringifyPopulatedTransaction,
  handleParams,
} from '../utils'

const getEmptyQuoteV2 = (params: SwapV2Parameters): SwapQuoteV2 => {
  return {
    id: '',
    chainId: params.chainId,
    fromToken: params.fromToken,
    fromAmount: params.fromAmount,
    toToken: params.toToken,
    expectedToAmount: '0',
    minToAmount: '0',
    moduleNames: [],
    estimatedTime: 0,
    routerAddress: AddressZero,
  }
}

export async function swapV2(
  this: SynapseSDK,
  params: SwapV2Parameters
): Promise<SwapQuoteV2> {
  params = handleParams(params)
  if (!areIntentsSupported(params.chainId)) {
    return getEmptyQuoteV2(params)
  }
  if (isSameAddress(params.fromToken, params.toToken)) {
    return getEmptyQuoteV2(params)
  }
  const slippage = slippageFromPercentage(params.slippagePercentage)
  const input: RouteInput = {
    chainId: params.chainId,
    fromToken: params.fromToken,
    fromAmount: params.fromAmount,
    swapper: this.swapEngineSet.getTokenZap(params.chainId),
    toToken: params.toToken,
    toRecipient: {
      entity: params.toRecipient
        ? RecipientEntity.User
        : RecipientEntity.UserSimulated,
      address: params.toRecipient || USER_SIMULATED_ADDRESS,
    },
    restrictComplexity: params.restrictComplexity ?? false,
  }
  const quote = await this.swapEngineSet.getBestQuote(input, {
    allowMultiStep: true,
  })
  if (!quote) {
    return getEmptyQuoteV2(params)
  }
  const route = await this.swapEngineSet.generateRoute(input, quote, {
    allowMultiStep: true,
    slippage,
  })
  if (!route) {
    return getEmptyQuoteV2(params)
  }
  const tx = params.toRecipient
    ? await this.sirSet.completeIntentWithBalanceChecks(
        params.chainId,
        params.fromToken,
        params.fromAmount,
        params.deadline ?? calculateDeadline(TEN_MINUTES),
        route.steps
      )
    : undefined
  const expectedToAmount = route.expectedToAmount.toString()
  return {
    id: uuidv7(),
    chainId: params.chainId,
    fromToken: params.fromToken,
    fromAmount: params.fromAmount,
    toToken: params.toToken,
    expectedToAmount,
    minToAmount: route.minToAmount?.toString() ?? expectedToAmount,
    routerAddress: this.sirSet.getSirAddress(params.chainId),
    estimatedTime: MEDIAN_TIME_BLOCK[params.chainId],
    moduleNames: [route.engineName],
    tx: stringifyPopulatedTransaction(tx),
  }
}

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
  amount: BigNumberish,
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
  amountIn: BigNumberish,
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
  return {
    routerAddress,
    maxAmountOut,
    query: applySwapDeadline(query, deadline),
  }
}

/**
 * Applies a deadline to the given swap query.
 *
 * @param queryInitial - The swap query
 * @param deadline - The deadline to apply (optional, defaults to 10 minutes from now)
 * @returns The swap query with deadline applied
 */
export const applySwapDeadline = (
  queryInitial: Query,
  deadline?: BigNumber
): Query => {
  return applyDeadlineToQuery(
    queryInitial,
    applyOptionalDeadline(deadline, TEN_MINUTES)
  )
}
/**
 * Applies slippage to the given swap query.
 * Note: default slippage is 10 bips (0.1%).
 *
 * @param queryInitial - The swap query, coming from `swapQuote()`
 * @param slipNumerator - The numerator of the slippage percentage, defaults to 10
 * @param slipDenominator - The denominator of the slippage percentage, defaults to 10000
 * @returns The swap query with slippage applied
 */
export const applySwapSlippage = (
  queryInitial: Query,
  slipNumerator: number = 10,
  slipDenominator: number = 10000
): Query => {
  return applySlippageToQuery(queryInitial, slipNumerator, slipDenominator)
}
