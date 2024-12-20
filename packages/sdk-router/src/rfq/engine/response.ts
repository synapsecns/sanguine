import { BigNumber } from 'ethers'
import { AddressZero } from '@ethersproject/constants'

import { BigintIsh } from '../../constants'
import { Recipient } from './swapEngine'
import { StepParams } from '../steps'
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

export const generateAPIStep = (
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  response: SwapAPIResponse,
  finalRecipient: Recipient,
  minAmountOut: BigNumber
): StepParams => {
  if (isSameAddress(finalRecipient.address, AddressZero)) {
    throw new Error('Missing recipient address')
  }
  const zapData = encodeZapData({
    target: response.transaction.to,
    payload: response.transaction.data,
    amountPosition: AMOUNT_NOT_PRESENT,
    finalToken: tokenOut,
    forwardTo: finalRecipient.address,
    minFwdAmount: minAmountOut,
  })
  return {
    token: tokenIn,
    amount: BigNumber.from(amountIn),
    msgValue: BigNumber.from(response.transaction.value),
    zapData,
  }
}
