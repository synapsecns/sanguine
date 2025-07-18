import { PopulatedTransaction } from 'ethers'

import { SynapseModule } from '../module'

export class RelayModule implements SynapseModule {
  readonly address: string
  readonly chainId: number

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

  public async getBridgeTxStatus(): Promise<boolean> {
    // TODO: implement
    return false
  }
}
