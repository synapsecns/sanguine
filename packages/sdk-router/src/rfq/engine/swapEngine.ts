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
  steps: StepParams[]
}

export const EmptyRoute: SwapEngineRoute = {
  engineID: EngineID.Null,
  expectedAmountOut: Zero,
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

export const applySlippage = (
  amount: BigNumber,
  slippage: Slippage
): BigNumber => {
  return amount.sub(amount.mul(slippage.numerator).div(slippage.denominator))
}
