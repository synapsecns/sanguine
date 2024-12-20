import { defaultAbiCoder } from '@ethersproject/abi'
import { BigNumber } from 'ethers'

import { IFastBridgeV2 } from '../typechain/FastBridgeV2'

export type SavedParamsV1 = {
  originSender: string
  destChainId: number
  destEngineID: number
  destRelayRecipient: string
  destRelayToken: string
  destRelayAmount: BigNumber
}
export type BridgeParamsV2 = IFastBridgeV2.BridgeParamsV2Struct
const savedBridgeParams = [
  'tuple(address,uint256,uint256,address,address,uint256)',
  'tuple(address,int256,bytes,uint256,bytes)',
]

export const encodeSavedBridgeParams = (
  paramsV1: SavedParamsV1,
  paramsV2: BridgeParamsV2
) => {
  return defaultAbiCoder.encode(savedBridgeParams, [
    [
      paramsV1.originSender,
      paramsV1.destChainId,
      paramsV1.destEngineID,
      paramsV1.destRelayRecipient,
      paramsV1.destRelayToken,
      paramsV1.destRelayAmount,
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
    [
      originSender,
      destChainId,
      destEngineID,
      destRelayRecipient,
      destRelayToken,
      destRelayAmount,
    ],
    [quoteRelayer, quoteExclusivitySeconds, quoteId, zapNative, zapData],
  ] = defaultAbiCoder.decode(savedBridgeParams, data)
  return {
    paramsV1: {
      originSender,
      destChainId: destChainId.toNumber(),
      destEngineID: destEngineID.toNumber(),
      destRelayRecipient,
      destRelayToken,
      destRelayAmount,
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
