import { BigNumberish } from 'ethers'

import { EngineID, Recipient } from '../core'
import { SwapEngineQuote, SwapEngineRoute } from './route'

/**
 * Input parameters for generating a swap route.
 *
 * @property {number} chainId - The chain ID of the route.
 * @property {string} tokenIn - The input token address.
 * @property {string} tokenOut - The output token address.
 * @property {BigNumberish} amountIn - The amount of input token to swap.
 * @property {string} msgSender - The address that will invoke the swap.
 * @property {Recipient} finalRecipient - The recipient of the output token.
 * @property {boolean} restrictComplexity - Whether to restrict the complexity of the route (no splitting, less steps).
 */
export type RouteInput = {
  chainId: number
  tokenIn: string
  tokenOut: string
  amountIn: BigNumberish
  msgSender: string
  finalRecipient: Recipient
  restrictComplexity: boolean
}

export interface SwapEngine {
  readonly id: EngineID

  /**
   * Gets a swap quote from the engine for the given tokenIn -> tokenOut input.
   * Some of the engines may not be able to generate the route steps at the same time,
   * use the `generateRoute` method to generate the steps.
   */
  getQuote(input: RouteInput, timeout: number): Promise<SwapEngineQuote>

  /**
   * Generates the route steps from the quote obtained from the `getQuote` method.
   */
  generateRoute(
    input: RouteInput,
    quote: SwapEngineQuote,
    timeout: number
  ): Promise<SwapEngineRoute>
}
