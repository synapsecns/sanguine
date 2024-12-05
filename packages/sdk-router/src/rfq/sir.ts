import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import fastBridgeV2Abi from '../abi/FastBridgeV2.json'
import synapseIntentRouterAbi from '../abi/SynapseIntentRouter.json'
import { FastBridgeV2 as FastBridgeV2Contract } from '../typechain/FastBridgeV2'
import { SynapseIntentRouter as SIRContract } from '../typechain/SynapseIntentRouter'
import { SynapseModule, Query } from '../module'
import { getMatchingTxLog } from '../utils/logs'
import { CACHE_TIMES, RouterCache } from '../utils/RouterCache'

export class SynapseIntentRouter implements SynapseModule {
  static fastBridgeV2Interface = new Interface(fastBridgeV2Abi)
  static sirInterface = new Interface(synapseIntentRouterAbi)

  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private readonly fastBridgeV2Contract: FastBridgeV2Contract
  private readonly sirContract: SIRContract

  // All possible events emitted by the FastBridgeV2 contract in the origin transaction (in alphabetical order)
  private readonly originEvents = ['BridgeRequested']

  constructor(
    chainId: number,
    provider: Provider,
    address: string,
    fastBridgeV2Address: string
  ) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(SynapseIntentRouter.fastBridgeV2Interface, 'INTERFACE_UNDEFINED')
    invariant(SynapseIntentRouter.sirInterface, 'INTERFACE_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = address
    this.fastBridgeV2Contract = new Contract(
      fastBridgeV2Address,
      fastBridgeV2Abi,
      provider
    ) as FastBridgeV2Contract
    this.sirContract = new Contract(
      address,
      SynapseIntentRouter.sirInterface,
      provider
    ) as SIRContract
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
    // TODO
  }

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    // TODO: this should support older instances of FastBridge to track legacy txs
    const fastBridgeLog = await getMatchingTxLog(
      this.provider,
      txHash,
      this.fastBridgeV2Contract,
      this.originEvents
    )
    // transactionId always exists in the log as we are using the correct ABI
    const parsedLog =
      this.fastBridgeV2Contract.interface.parseLog(fastBridgeLog)
    return parsedLog.args.transactionId
  }

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    // TODO: this should support older instances of FastBridge to track legacy txs
    return this.fastBridgeV2Contract.bridgeRelays(synapseTxId)
  }

  // ══════════════════════════════════════════════ FAST BRIDGE V2 ═══════════════════════════════════════════════════

  /**
   * @returns The protocol fee rate, multiplied by 1_000_000 (e.g. 1 basis point = 100).
   */
  @RouterCache(CACHE_TIMES.TEN_MINUTES)
  public async getProtocolFeeRate(): Promise<BigNumber> {
    return this.fastBridgeV2Contract.protocolFeeRate()
  }
}
