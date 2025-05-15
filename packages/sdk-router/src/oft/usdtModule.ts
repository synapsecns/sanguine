import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
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

type OftSendParams = {
  fromSender: string
  toRecipient: string
  toEid: number
  amount: BigNumberish
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
