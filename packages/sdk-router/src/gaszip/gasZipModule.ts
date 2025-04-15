import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, BigNumberish, PopulatedTransaction } from 'ethers'
import invariant from 'tiny-invariant'

import gasZipAbi from '../abi/GasZipV2.json'
import { Query, SynapseModule } from '../module'
import { isNativeToken } from '../utils'
import { getGasZipTxStatus } from './api'

export class GasZipModule implements SynapseModule {
  static gasZipInterface = new Interface(gasZipAbi)

  public readonly chainId: number
  public readonly provider: Provider
  public readonly address: string

  constructor(chainId: number, provider: Provider, address: string) {
    invariant(GasZipModule.gasZipInterface, 'INTERFACE_UNDEFINED')
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = address
  }

  /**
   * @inheritdoc SynapseModule.bridge
   */
  public async bridge(
    to: string,
    _destChainId: number,
    token: string,
    amount: BigNumberish,
    _originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    if (!isNativeToken(token)) {
      throw new Error('Non-native token not supported by gas.zip')
    }
    // Decode destQuery.rawParams to get GasZip's "short" chain ID
    const destGasZipChain = parseInt(destQuery.rawParams, 16)
    return this.populateGasZipTransaction(to, destGasZipChain, amount)
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

  private populateGasZipTransaction(
    to: string,
    destGasZipChain: number,
    amount: BigNumberish
  ): PopulatedTransaction {
    const data = GasZipModule.gasZipInterface.encodeFunctionData('deposit', [
      destGasZipChain,
      to,
    ])
    return {
      to: this.address,
      value: BigNumber.from(amount),
      data,
    }
  }
}
