import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import fastBridgeAbi from '../abi/FastBridge.json'
import fastBridgeRouterAbi from '../abi/FastBridgeRouter.json'
import { FastBridgeRouter as FastBridgeRouterContract } from '../typechain/FastBridgeRouter'
import {
  FastBridge as FastBridgeContract,
  IFastBridge,
} from '../typechain/FastBridge'
import {
  SynapseModule,
  Query,
  narrowToCCTPRouterQuery,
  reduceToQuery,
} from '../module'
import { BigintIsh } from '../constants'
import { getMatchingTxLog } from '../utils/logs'
import { adjustValueIfNative } from '../utils/handleNativeToken'
import { MS_TIMES, SimpleCache } from '../utils/SimpleCache'

// Define type alias
export type BridgeParams = IFastBridge.BridgeParamsStruct

export class FastBridgeRouter implements SynapseModule {
  static fastBridgeInterface = new Interface(fastBridgeAbi)
  static fastBridgeRouterInterface = new Interface(fastBridgeRouterAbi)

  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private readonly routerContract: FastBridgeRouterContract
  private fastBridgeContractCache: FastBridgeContract | undefined

  // All possible events emitted by the FastBridge contract in the origin transaction (in alphabetical order)
  private readonly originEvents = ['BridgeRequested']

  constructor(chainId: number, provider: Provider, address: string) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(FastBridgeRouter.fastBridgeRouterInterface, 'INTERFACE_UNDEFINED')
    invariant(FastBridgeRouter.fastBridgeInterface, 'INTERFACE_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = address
    this.routerContract = new Contract(
      address,
      FastBridgeRouter.fastBridgeRouterInterface,
      provider
    ) as FastBridgeRouterContract
    // this.hydrateCache().then(() => {
    //   console.log('rfq router cache hydrated')
    // })
  }

  private async hydrateCache() {
    try {
      await Promise.all([
        this.getFastBridgeContract(),
        this.getProtocolFeeRate(),
      ])
    } catch (e) {
      console.error('Failed to hydrate rfqrouter cache', e)
    }
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
    const populatedTransaction =
      await this.routerContract.populateTransaction.bridge(
        to,
        destChainId,
        token,
        amount,
        narrowToCCTPRouterQuery(originQuery),
        narrowToCCTPRouterQuery(destQuery)
      )
    // Adjust the tx.value if the initial token is native
    return adjustValueIfNative(
      populatedTransaction,
      token,
      BigNumber.from(amount)
    )
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

  @SimpleCache(MS_TIMES.TEN_MINUTES)
  public async chainGasAmount(): Promise<BigNumber> {
    const fastBridgeContract = await this.getFastBridgeContract()
    return fastBridgeContract.chainGasAmount()
  }

  public async getFastBridgeContract(): Promise<FastBridgeContract> {
    // Populate the cache if necessary
    if (!this.fastBridgeContractCache) {
      const address = await this.routerContract.fastBridge()
      this.fastBridgeContractCache = new Contract(
        address,
        FastBridgeRouter.fastBridgeInterface,
        this.provider
      ) as FastBridgeContract
    }
    return this.fastBridgeContractCache
  }

  // ══════════════════════════════════════════ FAST BRIDGE ROUTER ONLY ══════════════════════════════════════════════

  public async getOriginAmountOut(
    tokenIn: string,
    rfqTokens: string[],
    amountIn: BigintIsh
  ): Promise<Query[]> {
    const queries = await this.routerContract.getOriginAmountOut(
      tokenIn,
      rfqTokens,
      amountIn
    )
    return queries.map(reduceToQuery)
  }

  /**
   * @returns The protocol fee rate, multiplied by 1_000_000 (e.g. 1 basis point = 100).
   */
  @SimpleCache(MS_TIMES.TEN_MINUTES)
  public async getProtocolFeeRate(): Promise<BigNumber> {
    const fastBridgeContract = await this.getFastBridgeContract()
    return fastBridgeContract.protocolFeeRate()
  }
}