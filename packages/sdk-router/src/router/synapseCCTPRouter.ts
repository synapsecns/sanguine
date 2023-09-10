import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import { Abi } from '../utils/types'
import { SynapseCCTPRouter as SynapseCCTPRouterContract } from '../typechain/SynapseCCTPRouter'
import { Router } from './router'
import { Query, narrowToCCTPRouterQuery, reduceToQuery } from './query'
import { BigintIsh } from '../constants'
import {
  BridgeToken,
  DestRequest,
  FeeConfig,
  reduceToBridgeToken,
} from './types'

export class SynapseCCTPRouter extends Router {
  public readonly routerContract: SynapseCCTPRouterContract
  public readonly address: string

  constructor(chainId: number, provider: Provider, address: string, abi: Abi) {
    super(chainId, provider)
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(abi?.length, 'ABI_UNDEFINED')
    this.routerContract = new Contract(
      address,
      new Interface(abi),
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
}
