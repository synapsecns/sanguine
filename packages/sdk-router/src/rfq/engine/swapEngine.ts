import { BigNumber } from 'ethers'
import { Zero, WeiPerEther } from '@ethersproject/constants'

import { StepParams } from '../steps'
import { BigintIsh } from '../../constants'

export enum EngineID {
  Null,
  NoOp,
  Default,
  ParaSwap,
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

export const SlippageDefault: Slippage = {
  numerator: 10,
  denominator: 10000,
}

export const SlippageFull: Slippage = {
  numerator: 1,
  denominator: 1,
}

export const USER_SIMULATED_ADDRESS =
  '0xFAcefaCEFACefACeFaCefacEFaCeFACEFAceFAcE'

export interface SwapEngine {
  readonly id: EngineID

  findRoute(
    chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    finalRecipient: Recipient,
    slippage: Slippage
  ): Promise<SwapEngineRoute>

  applySlippage(
    chainId: number,
    route: SwapEngineRoute,
    slippage: Slippage
  ): SwapEngineRoute
}

export const validateEngineID = (engineID: number): engineID is EngineID => {
  return Object.values(EngineID).includes(engineID)
}

export const toBasisPoints = (slippage: Slippage): number => {
  return (slippage.numerator * 10000) / slippage.denominator
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
