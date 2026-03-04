import { Interface } from '@ethersproject/abi'
import { hexDataLength, hexDataSlice } from '@ethersproject/bytes'
import { AddressZero } from '@ethersproject/constants'
import { BigNumberish, PopulatedTransaction, utils } from 'ethers'

import {
  CCTP_V2_DOMAIN_MAP,
  CCTP_V2_FORWARD_HOOK_DATA,
  CCTP_V2_SUPPORTED_CHAIN_IDS,
} from '../constants'
import { SynapseModule } from '../module'
import { CctpV2Message, getMessages } from './api'

const TOKEN_MESSENGER_V2_ABI = [
  'function depositForBurnWithHook(uint256 amount,uint32 destinationDomain,bytes32 mintRecipient,address burnToken,bytes32 destinationCaller,uint256 maxFee,uint32 minFinalityThreshold,bytes hookData)',
]

const MESSAGE_SUCCESS_STATUS = 'complete'
const FORWARD_SUCCESS_STATES = new Set([
  'COMPLETE',
  'COMPLETED',
  'SUCCESS',
  'SUCCEEDED',
])

export type CctpV2BurnParams = {
  amount: BigNumberish
  destinationDomain: number
  mintRecipient: string
  burnToken: string
  maxFee: BigNumberish
  minFinalityThreshold: number
}

export class CCTPv2Module implements SynapseModule {
  static tokenMessengerV2Interface = new Interface(TOKEN_MESSENGER_V2_ABI)

  readonly address: string
  readonly chainId: number

  private amountPositionCache: number | undefined

  constructor(chainId: number, address: string) {
    this.chainId = chainId
    this.address = address
  }

  public async bridge(): Promise<PopulatedTransaction> {
    throw new Error('bridge V1 not supported')
  }

  public async getSynapseTxId(txHash: string): Promise<string> {
    return txHash
  }

  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    try {
      const message = await this.getMessageByTxHash(synapseTxId)
      if (!message) {
        return false
      }
      if (message.status?.toLowerCase() !== MESSAGE_SUCCESS_STATUS) {
        return false
      }
      const forwardState = message.forwardState?.toUpperCase()
      return !!forwardState && FORWARD_SUCCESS_STATES.has(forwardState)
    } catch {
      // Status lookup should never fail open due to transient API errors.
      return false
    }
  }

  public populateDepositForBurnWithHook(
    params: CctpV2BurnParams
  ): PopulatedTransaction {
    const data =
      CCTPv2Module.tokenMessengerV2Interface.encodeFunctionData(
        'depositForBurnWithHook',
        [
          params.amount,
          params.destinationDomain,
          utils.hexZeroPad(params.mintRecipient, 32),
          params.burnToken,
          utils.hexZeroPad(AddressZero, 32),
          params.maxFee,
          params.minFinalityThreshold,
          CCTP_V2_FORWARD_HOOK_DATA,
        ]
      )
    return {
      to: this.address,
      data,
    }
  }

  /**
   * Finds the amount position within encoded depositForBurnWithHook function data.
   */
  public getAmountPosition(params: Omit<CctpV2BurnParams, 'amount'>): number {
    if (this.amountPositionCache !== undefined) {
      return this.amountPositionCache
    }
    const amountAA = '0x' + 'aa'.repeat(32)
    const amountBB = '0x' + 'bb'.repeat(32)
    const dataAA = hexDataSlice(
      this.populateDepositForBurnWithHook({
        ...params,
        amount: amountAA,
      }).data!,
      4
    )
    const dataBB = hexDataSlice(
      this.populateDepositForBurnWithHook({
        ...params,
        amount: amountBB,
      }).data!,
      4
    )
    const length = hexDataLength(dataAA)
    if (length !== hexDataLength(dataBB)) {
      throw new Error(
        `Unable to find amount position: ${dataAA} and ${dataBB} are of different lengths`
      )
    }
    for (let i = 0; i < length / 32; i++) {
      const offset = i * 32
      if (
        hexDataSlice(dataAA, offset, offset + 32) === amountAA &&
        hexDataSlice(dataBB, offset, offset + 32) === amountBB
      ) {
        this.amountPositionCache = 4 + offset
        return this.amountPositionCache
      }
    }
    throw new Error(
      `Unable to find amount position within ${dataAA} and ${dataBB}`
    )
  }

  private async getMessageByTxHash(
    txHash: string
  ): Promise<CctpV2Message | null> {
    const sourceDomains = [
      ...new Set(
        CCTP_V2_SUPPORTED_CHAIN_IDS.map(
          (chainId) => CCTP_V2_DOMAIN_MAP[chainId]
        )
      ),
    ]
    const messagesByDomain = await Promise.all(
      sourceDomains.map((sourceDomainId) => getMessages(sourceDomainId, txHash))
    )
    for (const messages of messagesByDomain) {
      if (messages?.length) {
        return messages[0]
      }
    }
    return null
  }
}
