import { PopulatedTransaction } from 'ethers'
import { BigNumber } from '@ethersproject/bignumber'
import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { AddressZero } from '@ethersproject/constants'

import { BigintIsh } from '../constants'
import { Query } from '../module/query'
import { DestRequest } from './types'
import { BridgeToken, FeeConfig, SynapseModule } from '../module'

/**
 * Abstract class for a router contract deployed on a chain.
 * Handles contract interaction: the Router users don't need to know about the contract,
 * or the data structures used to interact with it.
 *
 * Instead, they use the Router class and generic types such as Query and BridgeToken.
 *
 * @property address The address of the router contract.
 * @property chainId The chain ID of chain the router is deployed on.
 * @property provider The provider used to interact with the chain router is deployed on.
 */
export abstract class Router implements SynapseModule {
  abstract readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private bridgeTokensCache: { [tokenOut: string]: BridgeToken[] } = {}

  constructor(chainId: number, provider: Provider) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
  }

  abstract getOriginAmountOut(
    tokenIn: string,
    bridgeTokens: string[],
    amountIn: BigintIsh
  ): Promise<Query[]>

  abstract getDestinationAmountOut(
    requests: DestRequest[],
    tokenOut: string
  ): Promise<Query[]>

  abstract getConnectedBridgeTokens(tokenOut: string): Promise<BridgeToken[]>

  abstract getBridgeFees(
    token: string,
    amount: BigNumber,
    isSwap: boolean
  ): Promise<{ feeAmount: BigNumber; feeConfig: FeeConfig }>

  /**
   * @inheritdoc SynapseModule.bridge
   */
  abstract bridge(
    to: string,
    chainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction>

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  abstract getSynapseTxId(txHash: string): Promise<string>

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  abstract getBridgeTxStatus(synapseTxId: string): Promise<boolean>

  /**
   * Fetches bridge tokens for a destination chain and output token.
   *
   * Checks the cache first, and fetches from the router if not cached. Filters invalid tokens and caches the result.
   *
   * @param destChainId - The destination chain ID.
   * @param tokenOut - The output token.
   * @param destRouter - The SynapseRouter or SynapseCCTPRouter to use.
   * @returns An array of BridgeToken objects for valid bridge tokens.
   */
  public async getBridgeTokens(tokenOut: string): Promise<BridgeToken[]> {
    // Populate the cache if necessary
    if (!this.bridgeTokensCache[tokenOut]) {
      // Fetch tokens from the router
      const routerBridgeTokens = await this.getConnectedBridgeTokens(tokenOut)
      // Filter out invalid tokens and cache the result
      this.bridgeTokensCache[tokenOut] = routerBridgeTokens.filter(
        (token) => token.symbol && token.token !== AddressZero
      )
    }
    // Return cached result
    return this.bridgeTokensCache[tokenOut]
  }

  /**
   * Fetches origin queries from either a SynapseRouter or SynapseCCTPRouter.
   *
   * @param tokenIn - The input token
   * @param tokenSymbols - The token symbols
   * @param amountIn - The input amount
   * @returns A promise that resolves to an array of Query objects with the same length as tokenSymbols.
   * @throws Will throw an error if unable to fetch origin queries
   */
  public async getOriginQueries(
    tokenIn: string,
    tokenSymbols: string[],
    amountIn: BigintIsh
  ): Promise<Query[]> {
    try {
      // Don't filter anything, as the amount of returned queries should match the amount of symbols
      return this.getOriginAmountOut(tokenIn, tokenSymbols, amountIn)
    } catch (error) {
      console.error('Failed to fetch origin queries', error)
      throw error
    }
  }

  /**
   * Fetches destination queries from either a SynapseRouter or SynapseCCTPRouter.
   *
   * @param requests - The requests with symbol and amount in.
   * @param tokenOut - The output token.
   * @returns A promise that resolves to an array of Query objects with the same length as requests.
   * @throws Will throw an error if unable to fetch destination queries.
   */
  public async getDestinationQueries(
    requests: DestRequest[],
    tokenOut: string
  ): Promise<Query[]> {
    try {
      // Don't filter anything, as the amount of returned queries should match the amount of requests
      return this.getDestinationAmountOut(requests, tokenOut)
    } catch (error) {
      console.error('Failed to fetch destination queries', error)
      throw error
    }
  }
}
