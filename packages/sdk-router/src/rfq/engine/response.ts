import { BigNumber } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'

import {
  getEmptyRoute,
  EngineID,
  RouteInput,
  SwapEngineRoute,
  getForwardTo,
} from './swapEngine'
import { isSameAddress } from '../../utils/addressUtils'
import { AMOUNT_NOT_PRESENT, encodeZapData } from '../zapData'

export type SwapAPIResponse = {
  amountOut: BigNumber
  transaction: {
    chainId: number
    from: string
    to: string
    value: string
    data: string
  }
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
    minFinalAmount: response.amountOut,
  })

  return {
    engineID,
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
