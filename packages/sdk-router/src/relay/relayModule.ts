import { PopulatedTransaction } from 'ethers'

import { SynapseModule } from '../module'
import { getRequests, Status } from './api'

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

  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    const requests = await getRequests({ hash: synapseTxId })
    if (!requests) {
      return false
    }
    const request = requests.requests.find(
      (r) => !!r.data.inTxs.find((tx) => tx.hash === synapseTxId)
    )
    return request?.status === Status.Success
  }
}
