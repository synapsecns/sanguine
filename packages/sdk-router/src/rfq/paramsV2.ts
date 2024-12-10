import { defaultAbiCoder } from '@ethersproject/abi'

import { IFastBridgeV2 } from '../typechain/FastBridgeV2'
import { ZapDataV1 } from './zapData'

export type BridgeParamsV2 = IFastBridgeV2.BridgeParamsV2Struct
const savedBridgeParams = [
  'address',
  'tuple(address,int256,bytes,uint256,bytes)',
  'tuple(address,bytes,uint256,address,address)',
]

export const encodeSavedBridgeParams = (
  sender: string,
  paramsV2: BridgeParamsV2,
  zapData: ZapDataV1
) => {
  return defaultAbiCoder.encode(savedBridgeParams, [
    sender,
    [
      paramsV2.quoteRelayer,
      paramsV2.quoteExclusivitySeconds,
      paramsV2.quoteId,
      paramsV2.zapNative,
      paramsV2.zapData,
    ],
    [
      zapData.target,
      zapData.payload,
      zapData.amountPosition,
      zapData.finalToken,
      zapData.forwardTo,
    ],
  ])
}

export const decodeSavedBridgeParams = (
  data: string
): {
  sender: string
  paramsV2: BridgeParamsV2
  zapData: ZapDataV1
} => {
  const [
    sender,
    [quoteRelayer, quoteExclusivitySeconds, quoteId, zapNative, zapData],
    [target, payload, amountPosition, finalToken, forwardTo],
  ] = defaultAbiCoder.decode(savedBridgeParams, data)
  return {
    sender,
    paramsV2: {
      quoteRelayer,
      quoteExclusivitySeconds,
      quoteId,
      zapNative,
      zapData,
    },
    zapData: {
      target,
      payload,
      amountPosition,
      finalToken,
      forwardTo,
    },
  }
}
