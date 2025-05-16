import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { hexDataSlice } from '@ethersproject/bytes'
import { AddressZero } from '@ethersproject/constants'
import {
  BigNumber,
  BigNumberish,
  Contract,
  PopulatedTransaction,
  utils,
} from 'ethers'

import {
  isChainIdSupported,
  LZ_EID_MAP,
  MEDIAN_TIME_BLOCK,
  USDT_OFT_ADDRESS_MAP,
} from '../constants'
import { SynapseModule } from '../module'
import { getWithTimeout } from '../utils'

const OFT_ABI = [
  'function quoteOFT(tuple(uint32,bytes32,uint256,uint256,bytes,bytes,bytes)) view returns (tuple(uint256,uint256), tuple(int256,string)[], tuple(uint256,uint256))',
  'function quoteSend(tuple(uint32,bytes32,uint256,uint256,bytes,bytes,bytes), bool) view returns (tuple(uint256,uint256))',
  'function send(tuple(uint32,bytes32,uint256,uint256,bytes,bytes,bytes), tuple(uint256,uint256), address) payable returns (tuple(bytes32,uint64,tuple(uint256,uint256)), tuple(uint256,uint256))',
]

const LZ_API_URL = 'https://scan.layerzero-api.com/v1'
const LZ_API_TIMEOUT = 5000
const LZ_COMPLETED_STATUSES = ['CONFIRMING', 'DELIVERED']

const AVG_DST_BLOCK_TARGET = 3

interface LZResponse {
  data?: {
    config?: {
      outboundConfig?: {
        confirmations?: number
      }
    }
    status?: {
      name?: string
    }
  }[]
}

export type OftSendParams = {
  toEid: number
  toRecipient: string
  amount: BigNumberish
  fromSender: string
}

type SendParamTuple = [
  number,
  string,
  BigNumberish,
  BigNumberish,
  string,
  string,
  string
]

type MsgFeeTuple = [BigNumberish, BigNumberish]

export class UsdtModule implements SynapseModule {
  static oftInterface = new Interface(OFT_ABI)

  readonly address: string
  readonly chainId: number
  readonly oftContract: Contract

  private amountPositionCache: number | undefined

  constructor(chainId: number, provider: Provider, address: string) {
    this.chainId = chainId
    this.address = address
    this.oftContract = new Contract(address, UsdtModule.oftInterface, provider)
  }

  public async bridge(): Promise<PopulatedTransaction> {
    throw new Error('bridge V1 not supported')
  }

  public async getSynapseTxId(txHash: string): Promise<string> {
    return txHash
  }

  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    const response = await getWithTimeout(
      'LZ API',
      `${LZ_API_URL}/messages/tx/${synapseTxId}`,
      LZ_API_TIMEOUT
    )
    if (!response) {
      return false
    }
    const txResponse: LZResponse = await response.json()
    const statusName = txResponse.data?.[0]?.status?.name
    return (
      !!statusName && LZ_COMPLETED_STATUSES.includes(statusName.toUpperCase())
    )
  }

  public async getEstimatedTime(
    toChainId: number
  ): Promise<number | undefined> {
    if (!isChainIdSupported(this.chainId) || !isChainIdSupported(toChainId)) {
      return undefined
    }
    if (toChainId === this.chainId) {
      return undefined
    }
    const [[fromEid, fromOft], [toEid, toOft]] = [this.chainId, toChainId].map(
      (chainId) => [LZ_EID_MAP[chainId], USDT_OFT_ADDRESS_MAP[chainId]]
    )
    if (!fromEid || !toEid || !fromOft || !toOft) {
      return undefined
    }
    // Get confirmations count using the latest message sent on the pathway
    const pathwayId = `${fromEid}-${toEid}-${fromOft.toLowerCase()}-${toOft.toLowerCase()}`
    const response = await getWithTimeout(
      'LZ API',
      `${LZ_API_URL}/messages/pathway/${pathwayId}`,
      LZ_API_TIMEOUT,
      {
        limit: 1,
      }
    )
    if (!response) {
      return undefined
    }
    const pathwayResponse: LZResponse = await response.json()
    const confirmations =
      pathwayResponse.data?.[0]?.config?.outboundConfig?.confirmations
    if (!confirmations) {
      return undefined
    }
    console.log(
      `Confirmations: ${confirmations}, fromChainId: ${this.chainId}, toChainId: ${toChainId}`
    )
    return (
      confirmations * MEDIAN_TIME_BLOCK[this.chainId] +
      AVG_DST_BLOCK_TARGET * MEDIAN_TIME_BLOCK[toChainId]
    )
  }

  public async getDestinationQuote(params: OftSendParams): Promise<BigNumber> {
    const [, , [, amountReceivedLD]] = await this.oftContract.quoteOFT(
      this.getSendParamTuple(params)
    )
    return amountReceivedLD
  }

  public async getNativeFee(params: OftSendParams): Promise<BigNumber> {
    const [nativeFee] = await this.oftContract.quoteSend(
      this.getSendParamTuple(params),
      false
    )
    return nativeFee
  }

  public populateOftSend(
    params: OftSendParams,
    nativeFee: BigNumberish
  ): PopulatedTransaction {
    // Use user origin address as the fee refund address
    const data = UsdtModule.oftInterface.encodeFunctionData('send', [
      this.getSendParamTuple(params),
      this.getMsgFeeTuple(nativeFee),
      params.fromSender,
    ])
    return {
      to: this.address,
      value: BigNumber.from(nativeFee),
      data,
    }
  }

  /**
   * Finds the amount position within encoded oft.send function data
   */
  public getAmountPosition(): number {
    if (this.amountPositionCache) {
      return this.amountPositionCache
    }
    // We use mock amounts (all A's and all B's) to find the amount position
    const amountAA = '0x' + 'aa'.repeat(32)
    const amountBB = '0x' + 'bb'.repeat(32)
    const mockParams = {
      toEid: 0,
      toRecipient: AddressZero,
      amount: 0,
      fromSender: AddressZero,
    }
    // Strip the function selector
    const dataAA = hexDataSlice(
      this.populateOftSend({ ...mockParams, amount: amountAA }, 0).data!,
      4
    )
    const dataBB = hexDataSlice(
      this.populateOftSend({ ...mockParams, amount: amountBB }, 0).data!,
      4
    )
    // Sanity check: both mock amounts should have the same length
    if (dataAA.length !== dataBB.length) {
      throw new Error(
        `Unable to find amount position: ${dataAA} and ${dataBB} are of different lengths`
      )
    }
    // Try offsets from 0 to data.length / 32
    for (let i = 0; i < dataAA.length / 32; i++) {
      const offset = i * 32
      // If data at offset matches both mock amounts, we found the amount position
      if (
        hexDataSlice(dataAA, offset, offset + 32) === amountAA &&
        hexDataSlice(dataBB, offset, offset + 32) === amountBB
      ) {
        this.amountPositionCache = 4 + offset
        return this.amountPositionCache
      }
    }
    // Sanity check: should have found the amount position
    throw new Error(
      `Unable to find amount position within ${dataAA} and ${dataBB}`
    )
  }

  public getSendParamTuple(params: OftSendParams): SendParamTuple {
    return [
      params.toEid,
      utils.hexZeroPad(params.toRecipient, 32),
      params.amount,
      0,
      '0x',
      '0x',
      '0x',
    ]
  }

  public getMsgFeeTuple(nativeFee: BigNumberish): MsgFeeTuple {
    return [nativeFee, 0]
  }
}
