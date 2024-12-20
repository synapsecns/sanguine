import { BigNumber } from 'ethers'
import { Zero, WeiPerEther } from '@ethersproject/constants'

import { StepParams } from '../steps'
import { BigintIsh } from '../../constants'

export enum EngineID {
  Null,
  NoOp,
  Default,
  ParaSwap,
  Odos,
}

export type SwapEngineRoute = {
  engineID: EngineID
  expectedAmountOut: BigNumber
  minAmountOut: BigNumber
  steps: StepParams[]
}

export const EmptyRoute: SwapEngineRoute = {
  engineID: EngineID.Null,
  expectedAmountOut: Zero,
  minAmountOut: Zero,
  steps: [],
}

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

// Default slippage used by the swap engines, 10 bips (0.1%)
export const SlippageDefault: Slippage = {
  numerator: 10,
  denominator: 10000,
}

// Max slippage that can be used by the swap engines, 100 bips (1%)
export const SlippageMax: Slippage = {
  numerator: 100,
  denominator: 10000,
}

export const USER_SIMULATED_ADDRESS =
  '0xFAcefaCEFACefACeFaCefacEFaCeFACEFAceFAcE'

export type RouteInput = {
  chainId: number
  tokenIn: string
  tokenOut: string
  amountIn: BigintIsh
  finalRecipient: Recipient
  slippage: Slippage
}

export interface SwapEngine {
  readonly id: EngineID

  findRoute(input: RouteInput): Promise<SwapEngineRoute>
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

export const isCorrectSlippage = (slippage: Slippage): boolean => {
  return (
    slippage.numerator >= 0 &&
    slippage.numerator <= slippage.denominator &&
    slippage.denominator > 0
  )
}

export const applySlippage = (
  amount: BigNumber,
  slippage: Slippage
): BigNumber => {
  return amount.sub(amount.mul(slippage.numerator).div(slippage.denominator))
}
