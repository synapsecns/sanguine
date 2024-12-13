import { defaultAbiCoder } from '@ethersproject/abi'
import { BigNumber } from 'ethers'

import { IFastBridgeV2 } from '../typechain/FastBridgeV2'
import { ZapDataV1 } from './zapData'

export type SavedParamsV1 = {
  originSender: string
  destRecipient: string
  destToken: string
  destAmount: BigNumber
}
export type BridgeParamsV2 = IFastBridgeV2.BridgeParamsV2Struct
const savedBridgeParams = [
  'tuple(address,address,address,uint256)',
  'tuple(address,int256,bytes,uint256,bytes)',
  'tuple(address,bytes,uint256,address,address)',
]

export const encodeSavedBridgeParams = (
  paramsV1: SavedParamsV1,
  paramsV2: BridgeParamsV2,
  zapData: ZapDataV1
) => {
  return defaultAbiCoder.encode(savedBridgeParams, [
    [
      paramsV1.originSender,
      paramsV1.destRecipient,
      paramsV1.destToken,
      paramsV1.destAmount,
    ],
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
  paramsV1: SavedParamsV1
  paramsV2: BridgeParamsV2
  zapData: ZapDataV1
} => {
  const [
    [originSender, destRecipient, destToken, destAmount],
    [quoteRelayer, quoteExclusivitySeconds, quoteId, zapNative, zapData],
    [target, payload, amountPosition, finalToken, forwardTo],
  ] = defaultAbiCoder.decode(savedBridgeParams, data)
  return {
    paramsV1: {
      originSender,
      destRecipient,
      destToken,
      destAmount,
    },
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
      amountPosition: amountPosition.toNumber(),
      finalToken,
      forwardTo,
    },
  }
}
