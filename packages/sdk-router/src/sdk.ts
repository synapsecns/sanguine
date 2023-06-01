import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { BigNumber } from '@ethersproject/bignumber'
import { PopulatedTransaction } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'
import { Interface } from '@ethersproject/abi'
import { Contract } from '@ethersproject/contracts'

import {
  handleNativeToken,
  ETH_NATIVE_TOKEN_ADDRESS,
} from './utils/handleNativeToken'
import { BigintIsh } from './constants'
import { SynapseRouter } from './synapseRouter'
import bridgeAbi from './abi/SynapseBridge.json'
import {
  Query,
  FeeConfig,
  convertQuery,
  RawQuery,
  PoolToken,
} from './utils/types'
const ONE_WEEK_DEADLINE = BigNumber.from(Math.floor(Date.now() / 1000) + 604800) // one week in the future
const TEN_MIN_DEADLINE = BigNumber.from(Math.floor(Date.now() / 1000) + 600) // ten minutes in the future

type SynapseRouters = {
  [key: number]: SynapseRouter
}
class SynapseSDK {
  public synapseRouters: SynapseRouters
  public providers: { [x: number]: Provider }
  public bridgeAbi: Interface = new Interface(bridgeAbi)
  public bridgeTokenCache: {
    [x: string]: { symbol: string; token: string }[]
  } = {}
  constructor(chainIds: number[], providers: Provider[]) {
    invariant(
      chainIds.length === providers.length,
      `Amount of chains and providers does not equal`
    )
    this.synapseRouters = {}
    this.providers = {}
    for (let i = 0; i < chainIds.length; i++) {
      this.synapseRouters[chainIds[i]] = new SynapseRouter(
        chainIds[i],
        providers[i]
      )
      this.providers[chainIds[i]] = providers[i]
    }
  }

  public async bridgeQuote(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    deadline?: BigNumber
  ): Promise<{
    feeAmount: BigNumber | undefined
    feeConfig: FeeConfig | undefined
    routerAddress: string | undefined
    maxAmountOut: BigNumber | undefined
    originQuery: Query | undefined
    destQuery: Query | undefined
  }> {
    tokenOut = handleNativeToken(tokenOut)
    tokenIn = handleNativeToken(tokenIn)
    const originRouter: SynapseRouter = this.synapseRouters[originChainId]
    const destRouter: SynapseRouter = this.synapseRouters[destChainId]
    const routerAddress = originRouter.routerContract.address

    let bridgeTokens = this.bridgeTokenCache[destChainId + '_' + tokenOut]
    if (!bridgeTokens) {
      const routerBridgeTokens =
        await destRouter.routerContract.getConnectedBridgeTokens(tokenOut)

      // Filter tokens with a valid symbol and address
      bridgeTokens = routerBridgeTokens.filter(
        (bridgeToken) =>
          bridgeToken.symbol.length && bridgeToken.token !== AddressZero
      )

      // Throw error if no valid bridge tokens found
      if (!bridgeTokens?.length) {
        throw new Error('No bridge tokens found for this route')
      }

      // Store only the symbol and token fields in the cache
      bridgeTokens = bridgeTokens.map(({ symbol, token }) => ({
        symbol,
        token,
      }))

      // Cache the bridge tokens
      this.bridgeTokenCache[destChainId + '_' + tokenOut] = bridgeTokens
    }

    // Get quotes from origin SynapseRouter
    const originQueries: RawQuery[] =
      await originRouter.routerContract.getOriginAmountOut(
        tokenIn,
        bridgeTokens.map((bridgeToken) => bridgeToken.symbol),
        amountIn
      )

    // create requests for destination router
    const requests: { symbol: string; amountIn: BigintIsh }[] = []
    for (let i = 0; i < bridgeTokens.length; i++) {
      requests.push({
        symbol: bridgeTokens[i].symbol,
        amountIn: originQueries[i].minAmountOut,
      })
    }
    if (originQueries.length === 0) {
      throw Error('No origin queries found for this route')
    }

    // Get quotes from destination SynapseRouter
    const destQueries: RawQuery[] =
      await destRouter.routerContract.getDestinationAmountOut(
        requests,
        tokenOut
      )
    if (destQueries.length === 0) {
      throw Error('No destination queries found for this route')
    }

    // Find the best query (in practice, we could return them all)
    let destInToken
    let rawOriginQuery
    let rawDestQuery
    let maxAmountOut: BigNumber = BigNumber.from(0)
    for (let i = 0; i < destQueries.length; i++) {
      if (destQueries[i].minAmountOut.gt(maxAmountOut)) {
        maxAmountOut = destQueries[i].minAmountOut
        rawOriginQuery = originQueries[i]
        rawDestQuery = destQueries[i]
        destInToken = bridgeTokens[i].token
      }
    }

    if (!rawOriginQuery || !rawDestQuery) {
      throw Error('No route found')
    }

    // Set default deadline
    const originQuery = convertQuery(rawOriginQuery)
    originQuery.deadline = deadline ?? TEN_MIN_DEADLINE
    const destQuery = convertQuery(rawDestQuery)
    destQuery.deadline = ONE_WEEK_DEADLINE

    // Get fee data
    let feeAmount
    let feeConfig

    if (originQuery && destInToken) {
      feeAmount = destRouter.routerContract.calculateBridgeFee(
        destInToken,
        originQuery.minAmountOut
      )
      feeConfig = destRouter.routerContract.fee(destInToken)
    }

    return {
      feeAmount: await feeAmount,
      feeConfig: await feeConfig,
      routerAddress,
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
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    token = handleNativeToken(token)
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

  public async swapQuote(
    chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    deadline?: BigNumber
  ): Promise<{
    routerAddress: string | undefined
    maxAmountOut: BigNumber | undefined
    query: Query | undefined
  }> {
    tokenOut = handleNativeToken(tokenOut)
    tokenIn = handleNativeToken(tokenIn)

    const router: SynapseRouter = this.synapseRouters[chainId]
    const routerAddress = router.routerContract.address

    const rawQuery = await router.routerContract.getAmountOut(
      tokenIn,
      tokenOut,
      amountIn
    )

    // Check if call was unsuccessful
    if (rawQuery?.length !== 5) {
      throw Error('No queries found for this route')
    }

    const query = convertQuery(rawQuery)
    query.deadline = deadline ?? DEFAULT_DEADLINE
    const maxAmountOut = query.minAmountOut

    return {
      routerAddress,
      maxAmountOut,
      query,
    }
  }

  public async swap(
    chainId: number,
    to: string,
    token: string,
    amount: BigintIsh,
    query: Query
  ): Promise<PopulatedTransaction> {
    token = handleNativeToken(token)
    const originRouter: SynapseRouter = this.synapseRouters[chainId]
    return originRouter.routerContract.populateTransaction.swap(
      to,
      token,
      amount,
      query
    )
  }
  public async getBridgeGas(chainId: number): Promise<BigintIsh> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const bridgeAddress = await router.routerContract.synapseBridge()
    const bridgeContract = new Contract(
      bridgeAddress,
      this.bridgeAbi,
      this.providers[chainId]
    )
    return bridgeContract.chainGasAmount()
  }

  public async getPoolTokens(
    chainId: number,
    poolAddress: string
  ): Promise<PoolToken[]> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const poolTokens = await router.routerContract.poolTokens(poolAddress)
    return poolTokens.map((token) => {
      return { token: token.token, isWeth: token?.isWeth }
    })
  }

  public async getPoolInfo(
    chainId: number,
    poolAddress: string
  ): Promise<{ tokens: BigNumber | undefined; lpToken: string | undefined }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const poolInfo = await router.routerContract.poolInfo(poolAddress)
    return { tokens: poolInfo?.[0], lpToken: poolInfo?.[1] }
  }

  public async getAllPools(chainId: number): Promise<
    {
      poolAddress: string | undefined
      tokens: PoolToken[] | undefined
      lpToken: string | undefined
    }[]
  > {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const pools = await router.routerContract.allPools()
    const res = pools.map((pool) => {
      return {
        poolAddress: pool?.pool,
        tokens: pool?.tokens.map((token) => {
          return { token: token.token, isWeth: token?.isWeth }
        }),
        lpToken: pool?.lpToken,
      }
    })
    return res
  }

  public async calculateAddLiquidity(
    chainId: number,
    poolAddress: string,
    amounts: Record<string, BigNumber>
  ): Promise<{ amount: BigNumber; routerAddress: string }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const poolTokens = await router.routerContract.poolTokens(poolAddress)
    const amountArr: BigNumber[] = []
    poolTokens.map((token) => {
      amountArr.push(amounts[token.token] ?? Zero)
    })
    if (amountArr.filter((amount) => !amount.isZero()).length === 0) {
      return { amount: Zero, routerAddress: router.routerContract.address }
    }
    return {
      amount: await router.routerContract.calculateAddLiquidity(
        poolAddress,
        amountArr
      ),
      routerAddress: router.routerContract.address,
    }
  }

  public async calculateRemoveLiquidity(
    chainId: number,
    poolAddress: string,
    amount: BigNumber
  ): Promise<{
    amounts: Record<string, { value: BigNumber; index: number }>
    routerAddress: string
  }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const amounts = await router.routerContract.calculateRemoveLiquidity(
      poolAddress,
      amount
    )
    const poolTokens = await router.routerContract.poolTokens(poolAddress)
    const amountsOut: Record<string, { value: BigNumber; index: number }> = {}
    poolTokens.map((token, index) => {
      amountsOut[token.token] = { value: amounts[index], index }
    })
    return {
      amounts: amountsOut,
      routerAddress: router.routerContract.address,
    }
  }

  public async calculateRemoveLiquidityOne(
    chainId: number,
    poolAddress: string,
    amount: BigNumber,
    token: string
  ): Promise<{
    amount: { value: BigNumber; index: number }
    routerAddress: string
  }> {
    const router: SynapseRouter = this.synapseRouters[chainId]

    let poolIndex = 0
    const poolTokens = await router.routerContract.poolTokens(poolAddress)
    poolTokens.map((poolToken, index) => {
      if (poolToken.token === token) {
        poolIndex = index
      }
    })

    const outAmount = await router.routerContract.calculateWithdrawOneToken(
      poolAddress,
      amount,
      poolIndex
    )

    return {
      amount: { value: outAmount, index: poolIndex },
      routerAddress: router.routerContract.address,
    }
  }
}

export { SynapseSDK, ETH_NATIVE_TOKEN_ADDRESS, Query, PoolToken }
