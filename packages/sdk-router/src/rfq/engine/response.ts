import { BigNumber } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'

import {
  applySlippage,
  EmptyRoute,
  EngineID,
  RouteInput,
  SwapEngineRoute,
} from './swapEngine'
import { isSameAddress } from '../../utils/addressUtils'
import { AMOUNT_NOT_PRESENT, encodeZapData } from '../zapData'

export type SwapAPIResponse = {
  amountOut: string
  transaction: {
    chainId: number
    from: string
    to: string
    value: string
    data: string
  }
}

export const EMPTY_SWAP_API_RESPONSE: SwapAPIResponse = {
  amountOut: '0',
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
  const expectedAmountOut = BigNumber.from(response.amountOut)
  if (expectedAmountOut.eq(Zero)) {
    return EmptyRoute
  }
  const minAmountOut = applySlippage(expectedAmountOut, input.slippage)
  const zapData = encodeZapData({
    target: response.transaction.to,
    payload: response.transaction.data,
    amountPosition: AMOUNT_NOT_PRESENT,
    finalToken: input.tokenOut,
    forwardTo: input.finalRecipient.address,
    minFwdAmount: minAmountOut,
  })

  return {
    engineID,
    expectedAmountOut,
    minAmountOut,
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
