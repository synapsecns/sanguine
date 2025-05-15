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

import { SynapseModule } from '../module'

const OFT_ABI = [
  'function quoteOFT(tuple(uint32,bytes32,uint256,uint256,bytes,bytes,bytes)) view returns (tuple(uint256,uint256), tuple(int256,string)[], tuple(uint256,uint256))',
  'function quoteSend(tuple(uint32,bytes32,uint256,uint256,bytes,bytes,bytes), bool) view returns (tuple(uint256,uint256))',
  'function send(tuple(uint32,bytes32,uint256,uint256,bytes,bytes,bytes), tuple(uint256,uint256), address) payable returns (tuple(bytes32,uint64,tuple(uint256,uint256)), tuple(uint256,uint256))',
]

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
  readonly oftContract: Contract

  private amountPositionCache: number | undefined

  constructor(provider: Provider, address: string) {
    this.address = address
    this.oftContract = new Contract(address, UsdtModule.oftInterface, provider)
  }

  public async bridge(): Promise<PopulatedTransaction> {
    throw new Error('bridge V1 not supported')
  }

  public async getSynapseTxId(txHash: string): Promise<string> {
    return txHash
  }

  public async getBridgeTxStatus(): Promise<boolean> {
    // TODO: implement
    return false
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
