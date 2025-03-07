import { BigNumber } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'

import { isSameAddress } from '../../utils/addressUtils'
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
  amountOut: BigNumber
  transaction: TransactionData
}

export const EMPTY_SWAP_API_RESPONSE: SwapAPIResponse = {
  amountOut: Zero,
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
  if (isSameAddress(input.finalRecipient.address, AddressZero)) {
    throw new Error('Missing recipient address')
  }
  if (response.amountOut.eq(Zero)) {
    return getEmptyRoute(engineID)
  }
  const zapData = encodeZapData({
    target: response.transaction.to,
    payload: response.transaction.data,
    amountPosition: AMOUNT_NOT_PRESENT,
    finalToken: input.tokenOut,
    forwardTo: getForwardTo(input.finalRecipient),
    minFinalAmount: applySlippage(response.amountOut, slippage),
  })

  return {
    engineID,
    engineName: EngineID[engineID],
    chainId: input.chainId,
    expectedAmountOut: response.amountOut,
    steps: [
      {
        token: input.tokenIn,
        amount: BigNumber.from(input.amountIn),
        msgValue: BigNumber.from(response.transaction.value),
        zapData,
      },
    ],
  }
}
