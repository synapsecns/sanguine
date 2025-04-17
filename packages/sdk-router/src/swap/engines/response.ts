import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { isSameAddress } from '../../utils'
import {
  AMOUNT_NOT_PRESENT,
  applySlippage,
  encodeZapData,
  EngineID,
  getForwardTo,
  Slippage,
} from '../core'
import { getEmptyRoute, RouteInput, SwapEngineRoute } from '../models'

export type TransactionData = {
  chainId: number
  from: string
  to: string
  value: string
  data: string
}

export type SwapAPIResponse = {
  expectedToAmount: BigNumber
  transaction: TransactionData
}

export const EMPTY_SWAP_API_RESPONSE: SwapAPIResponse = {
  expectedToAmount: Zero,
  transaction: {
    chainId: 0,
    from: '',
    to: '',
    value: '',
    data: '',
  },
}

export const generateAPIRoute = (
  input: RouteInput,
  engineID: EngineID,
  slippage: Slippage,
  response: SwapAPIResponse
): SwapEngineRoute => {
  if (isSameAddress(input.toRecipient.address, AddressZero)) {
    throw new Error('Missing recipient address')
  }
  if (response.expectedToAmount.eq(Zero)) {
    return getEmptyRoute(engineID)
  }
  const minToAmount = applySlippage(response.expectedToAmount, slippage)
  const zapData = encodeZapData({
    target: response.transaction.to,
    payload: response.transaction.data,
    amountPosition: AMOUNT_NOT_PRESENT,
    finalToken: input.toToken,
    forwardTo: getForwardTo(input.toRecipient),
    minFinalAmount: minToAmount,
  })

  return {
    engineID,
    engineName: EngineID[engineID],
    chainId: input.chainId,
    fromToken: input.fromToken,
    toToken: input.toToken,
    fromAmount: BigNumber.from(input.fromAmount),
    expectedToAmount: response.expectedToAmount,
    minToAmount,
    steps: [
      {
        token: input.fromToken,
        amount: BigNumber.from(input.fromAmount),
        msgValue: BigNumber.from(response.transaction.value),
        zapData,
      },
    ],
  }
}
