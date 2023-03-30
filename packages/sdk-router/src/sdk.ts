import { Provider } from '@ethersproject/abstract-provider'
import { AddressZero } from '@ethersproject/constants'
import invariant from 'tiny-invariant'
import { BigNumber } from '@ethersproject/bignumber'
import { BytesLike } from '@ethersproject/bytes'
import { PopulatedTransaction } from 'ethers'

import { BigintIsh } from './constants'
import { SynapseRouter } from './synapseRouter'

type SynapseRouters = {
  [key: number]: SynapseRouter
}

type Query = [string, string, BigNumber, BigNumber, string] & {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

type FeeConfig = [number, BigNumber, BigNumber] & {
  bridgeFee: number
  minFee: BigNumber
  maxFee: BigNumber
}

class SynapseSDK {
  public synapseRouters: SynapseRouters

  constructor(chainIds: number[], providers: Provider[]) {
    invariant(
      chainIds.length === providers.length,
      `Amount of chains and providers does not equal`
    )
    this.synapseRouters = {}
    for (let i = 0; i < chainIds.length; i++) {
      this.synapseRouters[chainIds[i]] = new SynapseRouter(
        chainIds[i],
        providers[i]
      )
    }
  }

  public async bridgeQuote(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<{
    feeAmount?: BigNumber | undefined
    feeConfig?: FeeConfig | undefined
    maxAmountOut?: BigNumber | undefined
    originQuery?: Query | undefined
    destQuery?: Query | undefined
  }> {
    let originQuery
    let destQuery
    const originRouter: SynapseRouter = this.synapseRouters[originChainId]
    const destRouter: SynapseRouter = this.synapseRouters[destChainId]

    // Step 0: find connected bridge tokens on destination
    const bridgeTokens =
      await destRouter.routerContract.getConnectedBridgeTokens(tokenOut)

    if (bridgeTokens.length === 0) {
      throw Error('No bridge tokens found for this route')
    }

    const filteredTokens = bridgeTokens.filter(
      (bridgeToken) =>
        bridgeToken.symbol.length !== 0 && bridgeToken.token !== AddressZero
    )

    // Step 1: perform a call to origin SynapseRouter
    const originQueries = await originRouter.routerContract.getOriginAmountOut(
      tokenIn,
      filteredTokens.map((bridgeToken) => bridgeToken.symbol),
      amountIn
    )

    // Step 2: form a list of Destination Requests
    // In practice, there is no need to pass the requests with amountIn = 0, but we will do it for code simplicity
    const requests: { symbol: string; amountIn: BigintIsh }[] = []

    for (let i = 0; i < filteredTokens.length; i++) {
      requests.push({
        symbol: filteredTokens[i].symbol,
        amountIn: originQueries[i].minAmountOut,
      })
    }

    // Step 3: perform a call to destination SynapseRouter
    const destQueries = await destRouter.routerContract.getDestinationAmountOut(
      requests,
      tokenOut
    )
    // Step 4: find the best query (in practice, we could return them all)
    let destInToken
    let maxAmountOut: BigNumber = BigNumber.from(0)
    for (let i = 0; i < destQueries.length; i++) {
      if (destQueries[i].minAmountOut.gt(maxAmountOut)) {
        maxAmountOut = destQueries[i].minAmountOut
        originQuery = originQueries[i]
        destQuery = destQueries[i]
        destInToken = filteredTokens[i].token
      }
    }

    // Get fee data
    const feeAmount =
      originQuery && destInToken
        ? await destRouter.routerContract.calculateBridgeFee(
            destInToken,
            originQuery.minAmountOut
          )
        : undefined

    const feeConfig = destInToken
      ? await destRouter.routerContract.fee(destInToken)
      : undefined

    return {
      feeAmount,
      feeConfig,
      maxAmountOut,
      originQuery,
      destQuery,
    }
  }

  public async bridge(
    to: string,
    originChainId: number,
    destChainId: number,
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
    const originRouter: SynapseRouter = this.synapseRouters[originChainId]
    return originRouter.routerContract.populateTransaction.bridge(
      to,
      destChainId,
      token,
      amount,
      originQuery,
      destQuery
    )
  }
}

export { SynapseSDK }
