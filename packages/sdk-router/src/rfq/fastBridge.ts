import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'

import { BigintIsh } from '../constants'
import { SynapseModule, Query } from '../module'
import { getMatchingTxLog } from '../utils/logs'

export class FastBridge implements SynapseModule {
  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  // All possible events emitted by the FastBridge contract in the origin transaction (in alphabetical order)
  private readonly originEvents = ['BridgeRequested']

  constructor(chainId: number, provider: Provider, address: string) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    invariant(address, 'ADDRESS_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = address
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
    // TODO: initiate the bridge transaction using FastBridge binding
    console.log(to, destChainId, token, amount, originQuery, destQuery)
    return null as any
  }

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    const fastBridgeContract = await this.getFastBridgeContract()
    const fastBridgeLog = await getMatchingTxLog(
      this.provider,
      txHash,
      fastBridgeContract,
      this.originEvents
    )
    // transactionId always exists in the log as we are using the correct ABI
    const parsedLog = fastBridgeContract.interface.parseLog(fastBridgeLog)
    return parsedLog.args.transactionId
  }

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    const fastBridgeContract = await this.getFastBridgeContract()
    return fastBridgeContract.bridgeRelays(synapseTxId)
  }

  private async getFastBridgeContract(): Promise<Contract> {
    // TODO: implement
    return null as any
  }
}
