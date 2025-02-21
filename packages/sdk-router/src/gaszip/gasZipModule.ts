import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, PopulatedTransaction } from 'ethers'
import invariant from 'tiny-invariant'

import { Query, SynapseModule } from '../module'
import { BigintIsh } from '../constants'
import { isNativeToken } from '../utils/handleNativeToken'
import { getGasZipQuote, getGasZipTxStatus } from './api'
import { isSameAddress } from '../utils/addressUtils'

export class GasZipModule implements SynapseModule {
  readonly address = '0x391E7C679d29bD940d63be94AD22A25d25b5A604'

  public readonly chainId: number
  public readonly provider: Provider

  constructor(chainId: number, provider: Provider) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
  }

  /**
   * @inheritdoc SynapseModule.bridge
   */
  public async bridge(
    to: string,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    if (!isNativeToken(token)) {
      throw new Error('Non-native token not supported by gas.zip')
    }
    if (isSameAddress(to, destQuery.rawParams)) {
      return {
        to: this.address,
        value: BigNumber.from(amount),
        data: originQuery.rawParams,
      }
    }
    const quote = await getGasZipQuote(this.chainId, destChainId, amount, to)
    const amountOut = BigNumber.from(quote.amountOut)
    if (amountOut.lt(destQuery.minAmountOut)) {
      throw new Error('Insufficient amount out')
    }
    return {
      to: this.address,
      value: amountOut,
      data: quote.calldata,
    }
  }

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    return txHash
  }

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    return getGasZipTxStatus(synapseTxId)
  }
}
