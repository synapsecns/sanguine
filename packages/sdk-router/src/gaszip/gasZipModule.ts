import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'

import { SynapseModule } from '../module'
import { BigintIsh } from '../constants'

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
    // TODO: implement
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
    // TODO: implement
  }
}
