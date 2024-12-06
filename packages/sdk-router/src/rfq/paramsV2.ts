import { defaultAbiCoder } from '@ethersproject/abi'

import { IFastBridgeV2 } from '../typechain/FastBridgeV2'

export type BridgeParamsV2 = IFastBridgeV2.BridgeParamsV2Struct
const savedBridgeParams = [
  'address',
  'tuple(address,int256,bytes,uint256,bytes)',
]

export const encodeSavedBridgeParams = (
  sender: string,
  paramsV2: BridgeParamsV2
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
  ])
}

export const decodeSavedBridgeParams = (
  data: string
): {
  sender: string
  paramsV2: BridgeParamsV2
} => {
  const [
    sender,
    [quoteRelayer, quoteExclusivitySeconds, quoteId, zapNative, zapData],
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
  }
}
