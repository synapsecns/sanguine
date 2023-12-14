import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import cctpRouterAbi from '../abi/SynapseCCTPRouter.json'
import { SynapseCCTPRouter as SynapseCCTPRouterContract } from '../typechain/SynapseCCTPRouter'
import { Router } from './router'
import { Query, narrowToCCTPRouterQuery, reduceToQuery } from './query'
import cctpAbi from '../abi/SynapseCCTP.json'
import { BigintIsh } from '../constants'
import {
  BridgeToken,
  DestRequest,
  FeeConfig,
  reduceToBridgeToken,
} from './types'

/**
 * Wrapper class for interacting with a SynapseCCTPRouter contract.
 * Abstracts away the contract interaction: the Router users don't need to know about the contract,
 * or the data structures used to interact with it.
 */
export class SynapseCCTPRouter extends Router {
  static routerInterface = new Interface(cctpRouterAbi)

  public readonly address: string

  private readonly routerContract: SynapseCCTPRouterContract
  private cctpContractCache: Contract | undefined

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

  public async getConnectedBridgeTokens(
    tokenOut: string
  ): Promise<BridgeToken[]> {
    return this.routerContract
      .getConnectedBridgeTokens(tokenOut)
      .then((bridgeTokens) => {
        return bridgeTokens.map(reduceToBridgeToken)
      })
  }

  public async getBridgeFees(
    token: string,
    amount: BigNumber,
    isSwap: boolean
  ): Promise<{ feeAmount: BigNumber; feeConfig: FeeConfig }> {
    const feeAmount = await this.routerContract.calculateFeeAmount(
      token,
      amount,
      isSwap
    )
    // Get fee structure, then assign minBaseFee/minSwapFee value to minFee based on isSwap flag
    const feeConfig = await this.routerContract
      .feeStructures(token)
      .then((feeStructure) => {
        return {
          bridgeFee: feeStructure.relayerFee,
          minFee: isSwap ? feeStructure.minSwapFee : feeStructure.minBaseFee,
          maxFee: feeStructure.maxFee,
        }
      })
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
    return this.routerContract.populateTransaction.bridge(
      to,
      chainId,
      token,
      amount,
      narrowToCCTPRouterQuery(originQuery),
      narrowToCCTPRouterQuery(destQuery)
    )
  }

  /**
   * @inheritdoc Router.getBridgeID
   */
  public async getBridgeID(txHash: string): Promise<string> {
    return txHash
  }

  /**
   * @inheritdoc Router.getBridgeTxStatus
   */
  public async getBridgeTxStatus(bridgeID: string): Promise<boolean> {
    const cctpContract = await this.getCctpContract()
    return cctpContract.isRequestFulfilled(bridgeID)
  }

  private async getCctpContract(): Promise<Contract> {
    // Populate the cache if necessary
    if (!this.cctpContractCache) {
      const cctpAddress = await this.routerContract.synapseCCTP()
      this.cctpContractCache = new Contract(
        cctpAddress,
        new Interface(cctpAbi),
        this.provider
      )
    }
    // Return the cached contract
    return this.cctpContractCache
  }
}
