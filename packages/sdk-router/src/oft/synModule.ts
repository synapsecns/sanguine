import { BigNumber, constants, utils } from 'ethers'

import {
  SYN_COMPOSER_ADDRESS,
  SYN_CORE_DECIMALS,
  SYN_EVM_DECIMALS,
  SYN_SHARED_DECIMALS,
} from '../constants'
import { OftSendParams, UsdtModule } from './usdtModule'

export type SynSendParams = OftSendParams & {
  minAmount?: BigNumber
  hyperCoreRecipient?: string
}

const OFT_CONVERSION_RATE = BigNumber.from(10).pow(
  SYN_EVM_DECIMALS - SYN_SHARED_DECIMALS
)
const CORE_CONVERSION_RATE = BigNumber.from(10).pow(
  SYN_EVM_DECIMALS - SYN_CORE_DECIMALS
)
const MAX_UINT64 = BigNumber.from(2).pow(64).sub(1)

export class SynModule extends UsdtModule {
  private pendingNativeFees = new Map<string, Promise<BigNumber>>()

  public async getDestinationQuote(params: SynSendParams): Promise<BigNumber> {
    const amount = BigNumber.from(params.amount)
    const minAmount = BigNumber.from(params.minAmount ?? 0)
    if (amount.lt(0) || amount.gt(constants.MaxUint256)) {
      throw new Error('SYN amount is outside uint256 range')
    }
    if (minAmount.lt(0) || minAmount.gt(constants.MaxUint256)) {
      throw new Error('SYN minimum amount is outside uint256 range')
    }
    // The deployed SYN adapters use the default OFT debit view: no fee, just
    // dust removal from 18 local decimals to 6 shared decimals.
    const sharedAmount = amount.div(OFT_CONVERSION_RATE)
    if (sharedAmount.gt(MAX_UINT64)) {
      throw new Error('SYN amount exceeds OFT shared amount range')
    }
    const received = sharedAmount.mul(OFT_CONVERSION_RATE)
    if (received.lt(minAmount)) {
      throw new Error('SYN amount is below the minimum received amount')
    }
    return received
  }

  public getNativeFee(params: SynSendParams): Promise<BigNumber> {
    const key = UsdtModule.oftInterface.encodeFunctionData('quoteSend', [
      this.getSendParamTuple(params),
      false,
    ])
    const pending = this.pendingNativeFees.get(key)
    if (pending) {
      return pending
    }
    // Share simultaneous fee reads, but never reuse a settled price on the next
    // refresh. The calldata key includes the amount, minimum, route and recipient.
    const request = super.getNativeFee(params)
    this.pendingNativeFees.set(key, request)
    const clear = () => this.pendingNativeFees.delete(key)
    void request.then(clear, clear)
    return request
  }

  public getSendParamTuple(
    params: SynSendParams
  ): ReturnType<UsdtModule['getSendParamTuple']> {
    return [
      params.toEid,
      utils.hexZeroPad(
        params.hyperCoreRecipient ? SYN_COMPOSER_ADDRESS : params.toRecipient,
        32
      ),
      params.amount,
      params.minAmount ?? 0,
      '0x', // The adapter enforces receive/compose execution options.
      params.hyperCoreRecipient
        ? utils.defaultAbiCoder.encode(
            ['uint256', 'address'],
            [0, params.hyperCoreRecipient]
          )
        : '0x',
      '0x',
    ]
  }
}

export const quoteSynHyperCore = (amount: BigNumber): BigNumber => {
  const coreAmount = amount.div(CORE_CONVERSION_RATE)
  if (amount.lt(0) || coreAmount.gt(MAX_UINT64)) {
    throw new Error('SYN amount exceeds HyperCore amount range')
  }
  // This deployment's Core bridge reserve exceeds SYN's maximum supply. The
  // composer still enforces capacity at execution; quotes need no balance read.
  // Destination quotes use native HyperCore units; OFT calldata remains in EVM units.
  return coreAmount
}
