import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { hexDataLength, hexDataSlice } from '@ethersproject/bytes'
import { AddressZero } from '@ethersproject/constants'
import { BigNumber, BigNumberish, Contract, PopulatedTransaction } from 'ethers'

import { MEDIAN_TIME_BLOCK } from '../constants'
import { SynapseModule } from '../module'
import { getWithTimeout } from '../utils'
import {
  getSbaChainMetadata,
  SBA_EXECUTION_BUFFER_SECONDS,
  SBA_MIN_GAS_LIMIT,
} from './metadata'

const SBA_ABI = [
  'function bridgeERC20(uint32 dstEid, address to, address token, uint256 amount, uint64 gasLimit) payable',
  'function getNativeFee(uint32 dstEid, uint64 gasLimit) view returns (uint256 nativeFee)',
]

const LZ_API_URL = 'https://scan.layerzero-api.com/v1'
const LZ_API_TIMEOUT = 5000
const LZ_COMPLETED_STATUSES = ['CONFIRMING', 'DELIVERED']

export type SynapseBridgeAdapterBridgeParams = {
  dstEid: number
  toRecipient: string
  token: string
  amount: BigNumberish
}

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

export class SynapseBridgeAdapterModule implements SynapseModule {
  static sbaInterface = new Interface(SBA_ABI)

  readonly address: string
  readonly chainId: number
  readonly adapterContract: Contract

  private amountPositionCache: number | undefined

  constructor(chainId: number, provider: Provider, address: string) {
    this.chainId = chainId
    this.address = address
    this.adapterContract = new Contract(
      address,
      SynapseBridgeAdapterModule.sbaInterface,
      provider
    )
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
    const fromMetadata = getSbaChainMetadata(this.chainId)
    const toMetadata = getSbaChainMetadata(toChainId)
    const fromBlockTime =
      MEDIAN_TIME_BLOCK[this.chainId as keyof typeof MEDIAN_TIME_BLOCK]
    const toBlockTime =
      MEDIAN_TIME_BLOCK[toChainId as keyof typeof MEDIAN_TIME_BLOCK]
    if (
      !fromMetadata ||
      !toMetadata ||
      !fromBlockTime ||
      !toBlockTime ||
      this.chainId === toChainId
    ) {
      return undefined
    }
    const pathwayId = [
      fromMetadata.lzEid,
      toMetadata.lzEid,
      this.address.toLowerCase(),
      toMetadata.adapterAddress.toLowerCase(),
    ].join('-')
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
    return confirmations * fromBlockTime + SBA_EXECUTION_BUFFER_SECONDS
  }

  public async getNativeFee(dstEid: number): Promise<BigNumber> {
    return this.adapterContract.getNativeFee(dstEid, SBA_MIN_GAS_LIMIT)
  }

  public populateBridgeERC20(
    params: SynapseBridgeAdapterBridgeParams,
    nativeFee: BigNumberish
  ): PopulatedTransaction {
    return {
      to: this.address,
      value: BigNumber.from(nativeFee),
      data: SynapseBridgeAdapterModule.sbaInterface.encodeFunctionData(
        'bridgeERC20',
        [
          params.dstEid,
          params.toRecipient,
          params.token,
          params.amount,
          SBA_MIN_GAS_LIMIT,
        ]
      ),
    }
  }

  public getAmountPosition(): number {
    if (this.amountPositionCache) {
      return this.amountPositionCache
    }
    const amountAA = '0x' + 'aa'.repeat(32)
    const amountBB = '0x' + 'bb'.repeat(32)
    const mockParams = {
      dstEid: 0,
      toRecipient: AddressZero,
      token: AddressZero,
      amount: 0,
    }
    const dataAA = hexDataSlice(
      this.populateBridgeERC20({ ...mockParams, amount: amountAA }, 0).data!,
      4
    )
    const dataBB = hexDataSlice(
      this.populateBridgeERC20({ ...mockParams, amount: amountBB }, 0).data!,
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
}
