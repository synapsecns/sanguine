import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

import { FiatTokenV2_2 as FiatTokenContract } from '../typechain/FiatTokenV2_2'
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
import cctpRouterAbi from '../abi/SynapseCCTPRouter.json'
import cctpAbi from '../abi/SynapseCCTP.json'
import fiatTokenAbi from '../abi/FiatTokenV2_2.json'
import { adjustValueIfNative } from '../utils/handleNativeToken'
import { getMatchingTxLog } from '../utils/logs'
import { BigintIsh, CCTP_TOKEN_MINTER_MAP } from '../constants'
import { DestRequest } from './types'

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

  /**
   * Returns the remaining minter allowance for a Circle token. Will return zero if the allowance could not be fetched.
   *
   * @param token - The Circle token address.
   * @returns The remaining minter allowance.
   */
  private async getTokenMinterAllowance(token: string): Promise<BigNumber> {
    const tokenMinter = CCTP_TOKEN_MINTER_MAP[this.chainId]
    if (!tokenMinter) {
      return Zero
    }
    const fiatTokenContract = new Contract(
      token,
      new Interface(fiatTokenAbi),
      this.provider
    ) as FiatTokenContract
    try {
      const minterAllowance = await fiatTokenContract.minterAllowance(
        tokenMinter
      )
      return minterAllowance
    } catch (error) {
      return Zero
    }
  }
}
