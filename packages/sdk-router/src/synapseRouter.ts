import { Interface } from '@ethersproject/abi'
import { Contract } from '@ethersproject/contracts'
import { Provider } from '@ethersproject/abstract-provider'
import { BytesLike } from '@ethersproject/bytes'
import invariant from 'tiny-invariant'
import { PopulatedTransaction } from 'ethers'

import abi from './abi/SynapseRouter.json'
import { BigintIsh, ROUTER_ADDRESS } from './constants'
import { SynapseRouter as SynapseRouterContract } from './typechain/SynapseRouter'
export class SynapseRouter {
  public static INTERFACE: Interface = new Interface(abi)
  public readonly chainId: number
  public readonly provider: Provider
  public readonly routerContract: SynapseRouterContract

  constructor(chainId: number, provider: Provider) {
    invariant(chainId !== undefined, 'CHAIN_ID_UNDEFINED')
    invariant(provider !== undefined, 'PROVIDER_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.routerContract = new Contract(
      ROUTER_ADDRESS[chainId as keyof object],
      SynapseRouter.INTERFACE,
      provider
    ) as SynapseRouterContract
  }

  public async getAmountOut(
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<any> {
    return this.routerContract.getAmountOut(tokenIn, tokenOut, amountIn)
  }

  public async getOriginAmountOut(
    tokenIn: string,
    bridgeTokens: string[],
    amountIn: BigintIsh
  ): Promise<any> {
    return this.routerContract.getOriginAmountOut(
      tokenIn,
      bridgeTokens,
      amountIn
    )
  }

  public async getDestinationAmountOut(
    requests: { symbol: string; amountIn: BigintIsh }[],
    tokenOut: string
  ): Promise<any> {
    return this.routerContract.getDestinationAmountOut(requests, tokenOut)
  }

  public async getConnectedBridgeTokens(tokenOut: string): Promise<any> {
    return this.routerContract.getConnectedBridgeTokens(tokenOut)
  }

  public async bridge(
    to: string,
    chainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: {
      swapAdapter: string
      tokenOut: string
      minAmountOut: BigintIsh
      deadline: BigintIsh
      rawParams: BytesLike
    },
    destQuery: {
      swapAdapter: string
      tokenOut: string
      minAmountOut: BigintIsh
      deadline: BigintIsh
      rawParams: BytesLike
    }
  ): Promise<PopulatedTransaction> {
    return this.routerContract.populateTransaction.bridge(
      to,
      chainId,
      token,
      amount,
      originQuery,
      destQuery,
      { value: 0 }
    )
  }
}
