import { Contract, PopulatedTransaction } from 'ethers'
import { BigNumber } from '@ethersproject/bignumber'
import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'

import { BigintIsh } from '../constants'
import { Query } from './query'
import { BridgeToken, DestRequest, FeeConfig } from './types'

export abstract class Router {
  abstract readonly routerContract: Contract
  public readonly chainId: number
  public readonly provider: Provider

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

  abstract bridge(
    to: string,
    chainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction>
}
