import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import { Abi } from '../utils/types'
import { SynapseRouter as SynapseRouterContract } from '../typechain/SynapseRouter'
import { Router } from './router'
import { Query, narrowToRouterQuery, reduceToQuery } from './query'
import { BigintIsh } from '../constants'
import {
  BridgeToken,
  DestRequest,
  FeeConfig,
  reduceToBridgeToken,
  reduceToFeeConfig,
} from './types'

export class SynapseRouter extends Router {
  public readonly routerContract: SynapseRouterContract

  constructor(chainId: number, provider: Provider, address: string, abi: Abi) {
    super(chainId, provider)
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(abi?.length, 'ABI_UNDEFINED')
    this.routerContract = new Contract(
      address,
      new Interface(abi),
      provider
    ) as SynapseRouterContract
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
    amount: BigNumber
  ): Promise<{ feeAmount: BigNumber; feeConfig: FeeConfig }> {
    const feeAmount = await this.routerContract.calculateBridgeFee(
      token,
      amount
    )
    const feeConfig = await this.routerContract
      .fee(token)
      .then(reduceToFeeConfig)
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
      narrowToRouterQuery(originQuery),
      narrowToRouterQuery(destQuery)
    )
  }
}
