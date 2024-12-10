import { AddressZero, MaxUint256 } from '@ethersproject/constants'

import { BigintIsh } from '../constants'
import { TokenZapV1 as TokenZapV1Contract } from '../typechain/TokenZapV1'

export type ZapDataV1 = {
  target: string
  payload: string
  amountPosition: BigintIsh
  finalToken: string
  forwardTo: string
}

export const encodeZapDataBytes = async (
  tokenZapContract: TokenZapV1Contract,
  zapData: Partial<ZapDataV1>
): Promise<string> => {
  if (!zapData.target) {
    return '0x'
  }
  const { target, payload, amountPosition, finalToken, forwardTo } =
    applyDefaultValues(zapData)
  // TODO: rework without RPC calls
  return tokenZapContract.encodeZapData(
    target,
    payload,
    amountPosition,
    finalToken,
    forwardTo
  )
}

export const applyDefaultValues = (zapData: Partial<ZapDataV1>): ZapDataV1 => {
  return {
    target: zapData.target || AddressZero,
    payload: zapData.payload || '0x',
    amountPosition: zapData.amountPosition || MaxUint256,
    finalToken: zapData.finalToken || AddressZero,
    forwardTo: zapData.forwardTo || AddressZero,
  }
}
