import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, utils } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  LZ_EID_MAP,
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
 * Checks finalized OFT receipt. For HyperCore, also verifies that compose submitted
 * the matching SpotSend action. HyperCore processes that action after the EVM block.
 */
export const getSynBridgeStatus = async (
  destChainId: number,
  txHash: string,
  hyperEvmProvider?: Provider
): Promise<boolean> => {
  const isCore = destChainId === HYPERCORE_CHAIN_ID
  const dstEid = LZ_EID_MAP[isCore ? SupportedChainId.HYPEREVM : destChainId]
  if (!dstEid) {
    return false
  }

  const response = await getWithTimeout(
    'LZ API',
    `${LZ_API_URL}/messages/tx/${txHash}`,
    LZ_API_TIMEOUT
  )
  if (!response) {
    return false
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
        return true
      }
      if (
        hyperEvmProvider &&
        message.destination.lzCompose?.status === 'SUCCEEDED' &&
        (await hasMatchingSpotSend(message, hyperEvmProvider))
      ) {
        return true
      }
    }
  } catch {
    // A malformed API response or unavailable destination RPC is not delivery.
  }
  return false
}

/** A successful lzCompose can refund to HyperEVM; require its exact CoreWriter action. */
const hasMatchingSpotSend = async (
  message: LzMessage,
  provider: Provider
): Promise<boolean> => {
  const payload = message.source?.tx?.payload
  if (
    !payload ||
    !utils.isHexString(payload) ||
    utils.hexDataLength(payload) !== 136
  ) {
    return false
  }

  const composerWord = utils.hexDataSlice(payload, 0, 32)
  if (
    composerWord.toLowerCase() !==
    utils.hexZeroPad(SYN_COMPOSER_ADDRESS, 32).toLowerCase()
  ) {
    return false
  }
  // OFT payload: bytes32 composer, uint64 amountSD, bytes32 sender,
  // abi.encode(uint256 minMsgValue, address Core recipient).
  if (!BigNumber.from(utils.hexDataSlice(payload, 72, 104)).isZero()) {
    return false
  }
  const amountSD = BigNumber.from(utils.hexDataSlice(payload, 32, 40))
  if (amountSD.isZero()) {
    return false
  }
  const recipient = utils.getAddress(utils.hexDataSlice(payload, 116, 136))
  // SYN has 6 shared decimals and 8 HyperCore weiDecimals.
  const expectedCoreAmount = amountSD.mul(100)

  for (const { txHash } of message.destination?.lzCompose?.txs ?? []) {
    if (!txHash) {
      continue
    }
    const receipt = await provider.getTransactionReceipt(txHash)
    if (receipt?.status !== 1) {
      continue
    }
    for (const log of receipt.logs) {
      if (log.address.toLowerCase() !== CORE_WRITER) {
        continue
      }
      try {
        const parsed = coreWriterInterface.parseLog(log)
        if (
          parsed.name !== 'RawAction' ||
          parsed.args.user.toLowerCase() !== SYN_COMPOSER_ADDRESS.toLowerCase()
        ) {
          continue
        }
        const action: string = parsed.args.data
        if (utils.hexDataSlice(action, 0, 4) !== SPOT_SEND_HEADER) {
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
          return true
        }
      } catch {
        // Ignore unrelated or malformed logs in the compose receipt.
      }
    }
  }
  return false
}
