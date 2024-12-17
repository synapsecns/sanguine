import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { StepParams } from '../steps'
import { BigintIsh } from '../../constants'

export enum EngineID {
  Null,
  NoOp,
  Default,
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
    strictOut: boolean
  ): Promise<SwapEngineRoute>

  modifyMinAmountOut(
    chainId: number,
    route: SwapEngineRoute,
    minAmountOut: BigintIsh
  ): SwapEngineRoute

  modifyRecipient(
    chainId: number,
    route: SwapEngineRoute,
    finalRecipient: Recipient
  ): SwapEngineRoute
}

export const validateEngineID = (engineID: number): engineID is EngineID => {
  return Object.values(EngineID).includes(engineID)
}
