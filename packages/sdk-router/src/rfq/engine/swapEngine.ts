import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { StepParams } from '../steps'
import { BigintIsh } from '../../constants'

export type SwapEngineRoute = {
  id: number
  expectedAmountOut: BigNumber
  minAmountOut: BigNumber
  steps: StepParams[]
}

export const EmptyRoute: SwapEngineRoute = {
  id: 0,
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
  readonly id: number

  findRoute(
    chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    finalRecipient: Recipient
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
