import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'

import cctpV2WithExecutorAbi from '../abi/ICCTPv2WithExecutor.json'
import { SynapseModule } from '../module'
import { ICCTPv2WithExecutor } from '../typechain/ICCTPv2WithExecutor'

export class CctpModule implements SynapseModule {
  readonly address: string
  readonly chainId: number

  readonly contract: ICCTPv2WithExecutor

  constructor(chainId: number, provider: Provider, address: string) {
    this.chainId = chainId
    this.address = address

    this.contract = new Contract(
      address,
      new Interface(cctpV2WithExecutorAbi),
      provider
    ) as ICCTPv2WithExecutor
  }

  public async bridge(): Promise<PopulatedTransaction> {
    throw new Error('bridge V1 not supported')
  }

  public async getSynapseTxId(txHash: string): Promise<string> {
    return txHash
  }

  public async getBridgeTxStatus(_synapseTxId: string): Promise<boolean> {
    // TODO: Implement
    return false
  }
}
