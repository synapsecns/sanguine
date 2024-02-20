import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'
import { solidityKeccak256 } from 'ethers/lib/utils'
import { AddressZero } from '@ethersproject/constants'

import routerAbi from '../abi/SynapseRouter.json'
import { SynapseBridge as SynapseBridgeContract } from '../typechain/SynapseBridge'
import {
  SynapseRouter as SynapseRouterContract,
  PoolStructOutput,
} from '../typechain/SynapseRouter'
import { Router } from './router'
import {
  BridgeToken,
  FeeConfig,
  Query,
  narrowToRouterQuery,
  reduceToBridgeToken,
  reduceToFeeConfig,
  reduceToQuery,
} from '../module'
import bridgeAbi from '../abi/SynapseBridge.json'
import { BigintIsh } from '../constants'
import {
  DestRequest,
  Pool,
  PoolInfo,
  PoolToken,
  reduceToPoolToken,
} from './types'
import { adjustValueIfNative } from '../utils/handleNativeToken'
import { getMatchingTxLog } from '../utils/logs'
import { CACHE_TIMES, SimpleCache } from '../utils/SimpleCache'

/**
 * Wraps [tokens, lpToken] returned by the SynapseRouter contract into a PoolInfo object.
 */
const wrapToPoolInfo = (poolInfo: [BigNumber, string]): PoolInfo => {
  return {
    tokens: poolInfo[0],
    lpToken: poolInfo[1],
  }
}

/**
 * Wraps the PoolStructOutput object returned by the SynapseRouter contract into a Pool object.
 */
const wrapToPool = (pool: PoolStructOutput): Pool => {
  return {
    poolAddress: pool.pool,
    tokens: pool.tokens.map(reduceToPoolToken),
    lpToken: pool.lpToken,
  }
}

/**
 * Enum representing the type of a SynapseBridge token.
 * NotSupported: the token is not supported by the SynapseBridge contract.
 * Redeem: the token is supported by the SynapseBridge contract by burning or minting on this chain.
 * Deposit: the token is supported by the SynapseBridge contract by depositing or withdrawing on this chain.
 */
export enum BridgeTokenType {
  NotSupported = -1,
  Redeem,
  Deposit,
}

/**
 * Wrapper class for interacting with a SynapseRouter contract.
 * Abstracts away the contract interaction: the Router users don't need to know about the contract,
 * or the data structures used to interact with it.
 */
export class SynapseRouter extends Router {
  static routerInterface = new Interface(routerAbi)

  public readonly address: string

  private readonly routerContract: SynapseRouterContract
  private bridgeContractCache?: SynapseBridgeContract

  // All possible events emitted by the SynapseBridge contract in the origin transaction (in alphabetical order)
  private readonly originEvents = [
    'TokenDeposit',
    'TokenDepositAndSwap',
    'TokenRedeem',
    'TokenRedeemAndRemove',
    'TokenRedeemAndSwap',
    'TokenRedeemV2',
  ]

  constructor(chainId: number, provider: Provider, address: string) {
    // Parent constructor throws if chainId or provider are undefined
    super(chainId, provider)
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(SynapseRouter.routerInterface, 'INTERFACE_UNDEFINED')
    this.routerContract = new Contract(
      address,
      SynapseRouter.routerInterface,
      provider
    ) as SynapseRouterContract
    this.address = address
    // this.hydrateCache().then(() => {
    //   console.log('router cache hydrated')
    // })
  }

  // private async hydrateCache() {
  //   try {
  //     await Promise.all([
  //       this.getBridgeContract(),
  //       this.chainGasAmount()
  //     ])
  //   } catch (e) {
  //     console.error('Failed to hydrate router cache', e)
  //   }
  // }

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

  @SimpleCache(CACHE_TIMES.TEN_MINUTES)
  public async getConnectedBridgeTokens(
    tokenOut: string
  ): Promise<BridgeToken[]> {
    return this.routerContract
      .getConnectedBridgeTokens(tokenOut)
      .then((bridgeTokens) => {
        return bridgeTokens.map(reduceToBridgeToken)
      })
  }

  @SimpleCache(CACHE_TIMES.TEN_MINUTES)
  public async getBridgeFees(
    token: string,
    amount: BigNumber
  ): Promise<{ feeAmount: BigNumber; feeConfig: FeeConfig }> {
    const [feeAmount, feeConfig] = await Promise.all([
      this.routerContract.calculateBridgeFee(token, amount),
      this.routerContract.fee(token).then(reduceToFeeConfig),
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
        narrowToRouterQuery(originQuery),
        narrowToRouterQuery(destQuery)
      )
    // Adjust the tx.value if the initial token is native
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
    // Check that the transaction hash refers to an origin transaction
    const bridgeContract = await this.getBridgeContract()
    await getMatchingTxLog(
      this.provider,
      txHash,
      bridgeContract,
      this.originEvents
    )
    // Once we know the transaction is an origin transaction, we can calculate the Synapse txId
    return solidityKeccak256(['string'], [txHash])
  }

  /**
   * @inheritdoc Router.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    const bridgeContract = await this.getBridgeContract()
    return bridgeContract.kappaExists(synapseTxId)
  }

  // ═════════════════════════════════════════ SYNAPSE ROUTER (V1) ONLY ══════════════════════════════════════════════

  private async getBridgeContract(): Promise<SynapseBridgeContract> {
    // Populate the cache if necessary
    if (!this.bridgeContractCache) {
      const bridgeAddress = await this.routerContract.synapseBridge()
      this.bridgeContractCache = new Contract(
        bridgeAddress,
        new Interface(bridgeAbi),
        this.provider
      ) as SynapseBridgeContract
    }
    // Return the cached contract
    return this.bridgeContractCache
  }

  @SimpleCache(CACHE_TIMES.TEN_MINUTES)
  public async chainGasAmount(): Promise<BigNumber> {
    const bridgeContract = await this.getBridgeContract()
    return bridgeContract.chainGasAmount()
  }

  @SimpleCache(CACHE_TIMES.TEN_MINUTES)
  public async getBridgeTokenType(token: string): Promise<BridgeTokenType> {
    const tokenConfig = await this.routerContract.config(token)
    // Check if token is supported
    if (tokenConfig.bridgeToken === AddressZero) {
      return BridgeTokenType.NotSupported
    }
    // Otherwise tokenConfig.tokenType is either 0 (Redeem) or 1 (Deposit)
    return tokenConfig.tokenType
  }

  public async getPoolTokens(poolAddress: string): Promise<PoolToken[]> {
    return this.routerContract.poolTokens(poolAddress)
  }

  public async getPoolInfo(poolAddress: string): Promise<PoolInfo> {
    return this.routerContract.poolInfo(poolAddress).then(wrapToPoolInfo)
  }

  public async getAllPools(): Promise<Pool[]> {
    return this.routerContract.allPools().then((pools) => {
      return pools.map(wrapToPool)
    })
  }

  public async calculateAddLiquidity(
    poolAddress: string,
    amounts: BigintIsh[]
  ): Promise<BigNumber> {
    return this.routerContract.calculateAddLiquidity(poolAddress, amounts)
  }

  public async calculateRemoveLiquidity(
    poolAddress: string,
    amount: BigintIsh
  ): Promise<BigNumber[]> {
    return this.routerContract.calculateRemoveLiquidity(poolAddress, amount)
  }

  public async calculateWithdrawOneToken(
    poolAddress: string,
    amount: BigintIsh,
    tokenIndex: number
  ): Promise<BigNumber> {
    return this.routerContract.calculateWithdrawOneToken(
      poolAddress,
      amount,
      tokenIndex
    )
  }

  public async getAmountOut(
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<Query> {
    return this.routerContract
      .getAmountOut(tokenIn, tokenOut, amountIn)
      .then(reduceToQuery)
  }

  public async swap(
    to: string,
    token: string,
    amount: BigintIsh,
    query: Query
  ): Promise<PopulatedTransaction> {
    const populatedTransaction =
      await this.routerContract.populateTransaction.swap(
        to,
        token,
        amount,
        narrowToRouterQuery(query)
      )
    // Adjust the tx.value if the initial token is native
    return adjustValueIfNative(
      populatedTransaction,
      token,
      BigNumber.from(amount)
    )
  }
}