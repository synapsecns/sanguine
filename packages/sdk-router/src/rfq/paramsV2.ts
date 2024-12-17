import { defaultAbiCoder } from '@ethersproject/abi'
import { BigNumber } from 'ethers'

import { IFastBridgeV2 } from '../typechain/FastBridgeV2'

export type SavedParamsV1 = {
  originSender: string
  destRecipient: string
  destChainId: number
  destToken: string
  destAmount: BigNumber
}
export type BridgeParamsV2 = IFastBridgeV2.BridgeParamsV2Struct
const savedBridgeParams = [
  'tuple(address,address,uint256,address,uint256)',
  'tuple(address,int256,bytes,uint256,bytes)',
]

export const encodeSavedBridgeParams = (
  paramsV1: SavedParamsV1,
  paramsV2: BridgeParamsV2
) => {
  return defaultAbiCoder.encode(savedBridgeParams, [
    [
      paramsV1.originSender,
      paramsV1.destRecipient,
      paramsV1.destChainId,
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
  ])
}

export const decodeSavedBridgeParams = (
  data: string
): {
  paramsV1: SavedParamsV1
  paramsV2: BridgeParamsV2
} => {
  const [
    [originSender, destRecipient, destChainId, destToken, destAmount],
    [quoteRelayer, quoteExclusivitySeconds, quoteId, zapNative, zapData],
  ] = defaultAbiCoder.decode(savedBridgeParams, data)
  return {
    paramsV1: {
      originSender,
      destRecipient,
      destChainId: destChainId.toNumber(),
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
  }
}
