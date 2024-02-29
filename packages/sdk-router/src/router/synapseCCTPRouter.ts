import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import cctpRouterAbi from '../abi/SynapseCCTPRouter.json'
import { SynapseCCTP as SynapseCCTPContract } from '../typechain/SynapseCCTP'
import { SynapseCCTPRouter as SynapseCCTPRouterContract } from '../typechain/SynapseCCTPRouter'
import { Router } from './router'
import {
  BridgeToken,
  FeeConfig,
  Query,
  narrowToCCTPRouterQuery,
  reduceToBridgeToken,
  reduceToQuery,
} from '../module'
import cctpAbi from '../abi/SynapseCCTP.json'
import { adjustValueIfNative } from '../utils/handleNativeToken'
import { getMatchingTxLog } from '../utils/logs'
import { BigintIsh, HYDRATION_SUPPORTED_CHAIN_IDS } from '../constants'
import { DestRequest } from './types'
import { CACHE_TIMES, RouterCache } from '../utils/RouterCache'
/**
 * Wrapper class for interacting with a SynapseCCTPRouter contract.
 * Abstracts away the contract interaction: the Router users don't need to know about the contract,
 * or the data structures used to interact with it.
 */
export class SynapseCCTPRouter extends Router {
  static routerInterface = new Interface(cctpRouterAbi)

  public readonly address: string

  private readonly routerContract: SynapseCCTPRouterContract
  private cctpContractCache?: SynapseCCTPContract

  // All possible events emitted by the SynapseCCTP contract in the origin transaction
  private readonly originEvents = ['CircleRequestSent']

  constructor(chainId: number, provider: Provider, address: string) {
    // Parent constructor throws if chainId or provider are undefined
    super(chainId, provider)
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(SynapseCCTPRouter.routerInterface, 'INTERFACE_UNDEFINED')
    this.routerContract = new Contract(
      address,
      SynapseCCTPRouter.routerInterface,
      provider
    ) as SynapseCCTPRouterContract
    this.address = address
    this.hydrateCache()
  }

  /** fully optional but improves perf on first request */
  private async hydrateCache() {
    if (HYDRATION_SUPPORTED_CHAIN_IDS.includes(this.chainId)) {
      try {
        await Promise.all([this.chainGasAmount()])
      } catch (e) {
        console.error('synapseCCTPRouter: Error hydrating cache', e)
      }
    }
  }

  public async getOriginAmountOut(
    tokenIn: string,
    bridgeTokens: string[],
    amountIn: BigintIsh
  ): Promise<Query[]> {
    return this.routerContract
      .getOriginAmountOut(tokenIn, bridgeTokens, amountIn)
      .then((queries) => {
        return queries.map(reduceToQuery)
      })
  }

  public async getDestinationAmountOut(
    requests: DestRequest[],
    tokenOut: string
  ): Promise<Query[]> {
    return this.routerContract
      .getDestinationAmountOut(requests, tokenOut)
      .then((queries) => {
        return queries.map(reduceToQuery)
      })
  }

  @RouterCache(CACHE_TIMES.TEN_MINUTES)
  public async getConnectedBridgeTokens(
    tokenOut: string
  ): Promise<BridgeToken[]> {
    return this.routerContract
      .getConnectedBridgeTokens(tokenOut)
      .then((bridgeTokens) => {
        return bridgeTokens.map(reduceToBridgeToken)
      })
  }

  @RouterCache(CACHE_TIMES.TEN_MINUTES)
  public async getBridgeFees(
    token: string,
    amount: BigNumber,
    isSwap: boolean
  ): Promise<{ feeAmount: BigNumber; feeConfig: FeeConfig }> {
    const [feeAmount, feeConfig] = await Promise.all([
      this.routerContract.calculateFeeAmount(token, amount, isSwap),
      this.routerContract.feeStructures(token).then((feeStructure) => ({
        bridgeFee: feeStructure.relayerFee,
        minFee: isSwap ? feeStructure.minSwapFee : feeStructure.minBaseFee,
        maxFee: feeStructure.maxFee,
      })),
    ])

    return { feeAmount, feeConfig }
  }

  public async bridge(
    to: string,
    chainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    const populatedTransaction =
      await this.routerContract.populateTransaction.bridge(
        to,
        chainId,
        token,
        amount,
        narrowToCCTPRouterQuery(originQuery),
        narrowToCCTPRouterQuery(destQuery)
      )
    // Adjust the tx.value if the token is native
    return adjustValueIfNative(
      populatedTransaction,
      token,
      BigNumber.from(amount)
    )
  }

  /**
   * @inheritdoc Router.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    const cctpContract = await this.getCctpContract()
    const cctpLog = await getMatchingTxLog(
      this.provider,
      txHash,
      cctpContract,
      this.originEvents
    )
    // RequestID always exists in the log as we are using the correct ABI
    const parsedLog = cctpContract.interface.parseLog(cctpLog)
    return parsedLog.args.requestID
  }

  /**
   * @inheritdoc Router.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    const cctpContract = await this.getCctpContract()
    return cctpContract.isRequestFulfilled(synapseTxId)
  }

  @RouterCache(CACHE_TIMES.TEN_MINUTES)
  public async chainGasAmount(): Promise<BigNumber> {
    const cctpContract = await this.getCctpContract()
    return cctpContract.chainGasAmount()
  }

  private async getCctpContract(): Promise<SynapseCCTPContract> {
    // Populate the cache if necessary
    if (!this.cctpContractCache) {
      const cctpAddress = await this.routerContract.synapseCCTP()
      this.cctpContractCache = new Contract(
        cctpAddress,
        new Interface(cctpAbi),
        this.provider
      ) as SynapseCCTPContract
    }
    // Return the cached contract
    return this.cctpContractCache
  }
}
