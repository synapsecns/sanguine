import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, utils } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  LZ_EID_MAP,
  SYN_ADDRESS_MAP,
  SYN_COMPOSER_ADDRESS,
  SYN_CORE_TOKEN_INDEX,
  SupportedChainId,
} from '../constants'
import { getWithTimeout } from '../utils'

const LZ_API_URL = 'https://scan.layerzero-api.com/v1'
const LZ_API_TIMEOUT = 5000
const CORE_WRITER = '0x3333333333333333333333333333333333333333'
const SPOT_SEND_HEADER = '0x01000006'
const coreWriterInterface = new Interface([
  'event RawAction(address indexed user, bytes data)',
])
const tokenInterface = new Interface([
  'event Transfer(address indexed from, address indexed to, uint256 value)',
])

interface LzMessage {
  pathway?: { dstEid?: number }
  status?: { name?: string }
  source?: { tx?: { payload?: string } }
  destination?: {
    status?: string
    lzCompose?: { status?: string; txs?: { txHash?: string }[] }
  }
}

interface LzResponse {
  data?: LzMessage[]
}

/**
 * Checks finalized OFT receipt. For HyperCore, verifies the matching SpotSend
 * submission or automatic SYN fallback to the same recipient on HyperEVM.
 */
export const getSynBridgeDeliveryChainId = async (
  destChainId: number,
  txHash: string,
  hyperEvmProvider?: Provider
): Promise<number | undefined> => {
  const isCore = destChainId === HYPERCORE_CHAIN_ID
  const dstEid = LZ_EID_MAP[isCore ? SupportedChainId.HYPEREVM : destChainId]
  if (!dstEid) {
    return undefined
  }

  const response = await getWithTimeout(
    'LZ API',
    `${LZ_API_URL}/messages/tx/${txHash}`,
    LZ_API_TIMEOUT
  )
  if (!response) {
    return undefined
  }

  try {
    const messages: LzResponse = await response.json()
    for (const message of messages.data ?? []) {
      if (
        message.pathway?.dstEid !== dstEid ||
        message.status?.name !== 'DELIVERED' ||
        message.destination?.status !== 'SUCCEEDED'
      ) {
        continue
      }
      if (!isCore) {
        return destChainId
      }
      if (
        hyperEvmProvider &&
        message.destination.lzCompose?.status === 'SUCCEEDED'
      ) {
        const deliveryChainId = await getComposeDeliveryChainId(
          message,
          hyperEvmProvider
        )
        if (deliveryChainId !== undefined) {
          return deliveryChainId
        }
      }
    }
  } catch {
    // A malformed API response or unavailable destination RPC is not delivery.
  }
  return undefined
}

export const getSynBridgeStatus = async (
  destChainId: number,
  txHash: string,
  hyperEvmProvider?: Provider
): Promise<boolean> =>
  (await getSynBridgeDeliveryChainId(destChainId, txHash, hyperEvmProvider)) !==
  undefined

/** A successful lzCompose must submit the Core action or return the full amount on HyperEVM. */
const getComposeDeliveryChainId = async (
  message: LzMessage,
  provider: Provider
): Promise<number | undefined> => {
  const payload = message.source?.tx?.payload
  if (
    !payload ||
    !utils.isHexString(payload) ||
    utils.hexDataLength(payload) !== 136
  ) {
    return undefined
  }

  const composerWord = utils.hexDataSlice(payload, 0, 32)
  if (
    composerWord.toLowerCase() !==
    utils.hexZeroPad(SYN_COMPOSER_ADDRESS, 32).toLowerCase()
  ) {
    return undefined
  }
  // OFT payload: bytes32 composer, uint64 amountSD, bytes32 sender,
  // abi.encode(uint256 minMsgValue, address Core recipient).
  const [minMsgValue, recipient] = utils.defaultAbiCoder.decode(
    ['uint256', 'address'],
    utils.hexDataSlice(payload, 72)
  )
  if (!BigNumber.from(minMsgValue).isZero()) {
    return undefined
  }
  const amountSD = BigNumber.from(utils.hexDataSlice(payload, 32, 40))
  if (amountSD.isZero()) {
    return undefined
  }
  // SYN has 6 shared decimals and 8 HyperCore weiDecimals.
  const expectedCoreAmount = amountSD.mul(100)
  const expectedEvmAmount = amountSD.mul(BigNumber.from(10).pow(12))

  for (const { txHash } of message.destination?.lzCompose?.txs ?? []) {
    if (!txHash) {
      continue
    }
    const receipt = await provider.getTransactionReceipt(txHash)
    if (receipt?.status !== 1) {
      continue
    }
    for (const log of receipt.logs) {
      try {
        if (
          log.address.toLowerCase() ===
          SYN_ADDRESS_MAP[SupportedChainId.HYPEREVM].toLowerCase()
        ) {
          const transfer = tokenInterface.parseLog(log)
          if (
            transfer.args.from.toLowerCase() ===
              SYN_COMPOSER_ADDRESS.toLowerCase() &&
            transfer.args.to.toLowerCase() === recipient.toLowerCase() &&
            BigNumber.from(transfer.args.value).eq(expectedEvmAmount)
          ) {
            return SupportedChainId.HYPEREVM
          }
          continue
        }
        if (log.address.toLowerCase() !== CORE_WRITER) {
          continue
        }
        const parsed = coreWriterInterface.parseLog(log)
        if (
          parsed.name !== 'RawAction' ||
          parsed.args.user.toLowerCase() !== SYN_COMPOSER_ADDRESS.toLowerCase()
        ) {
          continue
        }
        const action: string = parsed.args.data
        if (
          utils.hexDataLength(action) !== 100 ||
          utils.hexDataSlice(action, 0, 4) !== SPOT_SEND_HEADER
        ) {
          continue
        }
        const [to, index, amount] = utils.defaultAbiCoder.decode(
          ['address', 'uint64', 'uint64'],
          utils.hexDataSlice(action, 4)
        )
        if (
          to.toLowerCase() === recipient.toLowerCase() &&
          BigNumber.from(index).eq(SYN_CORE_TOKEN_INDEX) &&
          BigNumber.from(amount).eq(expectedCoreAmount)
        ) {
          return HYPERCORE_CHAIN_ID
        }
      } catch {
        // Ignore unrelated or malformed logs in the compose receipt.
      }
    }
  }
  return undefined
}
