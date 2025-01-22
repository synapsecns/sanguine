import { BigNumber } from 'ethers'
import { Zero, WeiPerEther, AddressZero } from '@ethersproject/constants'

import { StepParams } from '../steps'
import { BigintIsh } from '../../constants'

export enum EngineID {
  Null,
  NoOp,
  Default,
  KyberSwap,
  ParaSwap,
}

export type SwapEngineQuote = {
  engineID: EngineID
  chainId: number
  expectedAmountOut: BigNumber
  steps?: StepParams[]
}

export type SwapEngineRoute = Required<SwapEngineQuote>

export enum RecipientEntity {
  Self,
  User,
  UserSimulated,
}

export type Recipient = {
  entity: RecipientEntity
  address: string
}

export type Slippage = {
  numerator: number
  denominator: number
}

// Max slippage that can be used by the swap engines, 100 bips (1%)
export const SlippageMax: Slippage = {
  numerator: 100,
  denominator: 10000,
}

export const USER_SIMULATED_ADDRESS =
  '0xFAcefaCEFACefACeFaCefacEFaCeFACEFAceFAcE'

/**
 * Input parameters for generating a swap route.
 *
 * @property {number} chainId - The chain ID of the route.
 * @property {string} tokenIn - The input token address.
 * @property {string} tokenOut - The output token address.
 * @property {BigintIsh} amountIn - The amount of input token to swap.
 * @property {Recipient} finalRecipient - The recipient of the output token.
 * @property {boolean} restrictComplexity - Whether to restrict the complexity of the route (no splitting, less steps).
 */
export type RouteInput = {
  chainId: number
  tokenIn: string
  tokenOut: string
  amountIn: BigintIsh
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

export const validateEngineID = (engineID: number): engineID is EngineID => {
  return Object.values(EngineID).includes(engineID)
}

export const toBasisPoints = (slippage: Slippage): number => {
  return Math.round((slippage.numerator * 10000) / slippage.denominator)
}

export const toPercentFloat = (slippage: Slippage): number => {
  return (slippage.numerator * 100) / slippage.denominator
}

export const toWei = (slippage: Slippage): BigNumber => {
  return BigNumber.from(slippage.numerator)
    .mul(WeiPerEther)
    .div(slippage.denominator)
}

export const applySlippage = (
  amount: BigNumber,
  slippage: Slippage
): BigNumber => {
  return amount.sub(amount.mul(slippage.numerator).div(slippage.denominator))
}

export const getEmptyQuote = (engineID: EngineID): SwapEngineQuote => {
  return {
    engineID,
    chainId: 0,
    expectedAmountOut: Zero,
  }
}

export const getEmptyRoute = (engineID: EngineID): SwapEngineRoute => {
  return {
    engineID,
    chainId: 0,
    expectedAmountOut: Zero,
    steps: [],
  }
}

export const sanitizeMultiStepQuote = (
  quote: SwapEngineQuote
): SwapEngineQuote => {
  if (!quote.steps || quote.steps.length <= 1) {
    return quote
  }
  return getEmptyQuote(quote.engineID)
}

export const sanitizeMultiStepRoute = (
  route: SwapEngineRoute
): SwapEngineRoute => {
  if (route.steps.length <= 1) {
    return route
  }
  return getEmptyRoute(route.engineID)
}

export const getForwardTo = (recipient: Recipient): string => {
  return recipient.entity === RecipientEntity.Self
    ? AddressZero
    : recipient.address
}
