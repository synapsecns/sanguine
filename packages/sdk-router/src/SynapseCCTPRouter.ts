import { Interface } from '@ethersproject/abi'
import { Contract } from '@ethersproject/contracts'
import { Provider } from '@ethersproject/abstract-provider'
import { BytesLike } from '@ethersproject/bytes'
import invariant from 'tiny-invariant'
import { PopulatedTransaction } from 'ethers'

import abi from './abi/SynapseCCTPRouter.json'
import { BigintIsh, CCTP_ROUTER_ADDRESS } from './constants'
import { SynapseCCTPRouter as SynapseCCTPRouterContract } from './typechain/SynapseCCTPRouter'

export class SynapseCCTPRouter {
  public static INTERFACE: Interface = new Interface(abi)
  public readonly chainId: number
  public readonly provider: Provider
  public readonly routerContract: SynapseCCTPRouterContract

  constructor(chainId: number, provider: Provider) {
    invariant(chainId !== undefined, 'CHAIN_ID_UNDEFINED')
    invariant(provider !== undefined, 'PROVIDER_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.routerContract = new Contract(
      CCTP_ROUTER_ADDRESS[chainId as keyof object],
      SynapseCCTPRouter.INTERFACE,
      provider
    ) as SynapseCCTPRouterContract
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
      routerAdapter: string
      tokenOut: string
      minAmountOut: BigintIsh
      deadline: BigintIsh
      rawParams: BytesLike
    },
    destQuery: {
      routerAdapter: string
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
