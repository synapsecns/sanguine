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
import { BigintIsh, CCTP_ROUTER_ADDRESS } from './constants'
import { SynapseRouter } from './synapseRouter'
import bridgeAbi from './abi/SynapseBridge.json'
import {
  Query,
  FeeConfig,
  convertQuery,
  RawQuery,
  PoolToken,
} from './utils/types'
import { SynapseCCTPRouter } from './SynapseCCTPRouter'
const ONE_WEEK_DEADLINE = BigNumber.from(Math.floor(Date.now() / 1000) + 604800) // one week in the future
const TEN_MIN_DEADLINE = BigNumber.from(Math.floor(Date.now() / 1000) + 600) // ten minutes in the future

type SynapseRouters = {
  [key: number]: SynapseRouter
}
type SynapseCCTPRouters = {
  [key: number]: SynapseCCTPRouter
}
type BridgeQuote = {
  feeAmount: BigNumber | undefined;
  feeConfig: FeeConfig | undefined;
  routerAddress: string | undefined;
  maxAmountOut: BigNumber | undefined;
  originQuery: Query | undefined;
  destQuery: Query | undefined;
}

class SynapseSDK {
  public synapseRouters: SynapseRouters
  public synapseCCTPRouters: SynapseCCTPRouters
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
    this.synapseCCTPRouters = {}
    this.providers = {}
    for (let i = 0; i < chainIds.length; i++) {
      this.synapseRouters[chainIds[i]] = new SynapseRouter(
        chainIds[i],
        providers[i]
      )
      this.providers[chainIds[i]] = providers[i]
      // check if the chain id is in the CCTP_ROUTER_ADDRESS object
      if(CCTP_ROUTER_ADDRESS.hasOwnProperty(chainIds[i])) {
        this.synapseCCTPRouters[chainIds[i]] = new SynapseCCTPRouter(
          chainIds[i],
          providers[i]
        )
      }
    }
  }


  public async getBridgeTokens(
    destChainId: number,
    tokenOut: string,
    destRouters: (SynapseRouter | SynapseCCTPRouter)[]
  ): Promise<{ symbol: string; token: string }[][]> {

    const bridgeTokensPromises = destRouters.map(async (destRouter) => {
      // Check the cache first
      let bridgeTokens = this.bridgeTokenCache[destChainId + '_' + tokenOut];

      // If not in cache, fetch from destination router
      if (!bridgeTokens) {
        const routerBridgeTokens =
          await destRouter.routerContract.getConnectedBridgeTokens(tokenOut);

        // Filter tokens with a valid symbol and address
        bridgeTokens = routerBridgeTokens.filter(
          (bridgeToken) =>
            bridgeToken.symbol.length && bridgeToken.token !== AddressZero
        );

        // Throw error if no valid bridge tokens found
        if (!bridgeTokens?.length) {
          // throw new Error('No bridge tokens found for this route');
          return bridgeTokens
        }

        // Store only the symbol and token fields in the cache
        bridgeTokens = bridgeTokens.map(({ symbol, token }) => ({
          symbol,
          token,
        }));

        // Cache the bridge tokens
        this.bridgeTokenCache[destChainId + '_' + tokenOut] = bridgeTokens;
      }
      return bridgeTokens;
    });

    // Fetch bridge tokens for all routers in parallel
    const allBridgeTokens = await Promise.all(bridgeTokensPromises);

    return allBridgeTokens;
  }


    // Function to fetch origin queries from a router
  public async getOriginQueries(
      router: SynapseRouter | SynapseCCTPRouter,
      tokenIn: string,
      tokenSymbols: string[],
      amountIn: BigintIsh
    ): Promise<RawQuery[]> {
      const routerQueries = await router.routerContract.getOriginAmountOut(
        tokenIn,
        tokenSymbols,
        amountIn
      );

      const filteredOriginQueries = (routerQueries as any).filter((query: { minAmountOut: { eq: (arg0: number) => any } }) => !query.minAmountOut.eq(0));

      return filteredOriginQueries.map((routerQuery: any) => {
        // Normalize key differences between router types
        return {
          ...routerQuery,
          swapAdapter: routerQuery.swapAdapter ?? routerQuery.routerAdapter,
        };
      });
    }


    public async getDestinationQueries(
      router:  SynapseRouter | SynapseCCTPRouter,
      requests: { symbol: string; amountIn: BigintIsh }[],
      tokenOut: string
    ): Promise<RawQuery[]> {
      const routerQueries = await router.routerContract.getDestinationAmountOut(
        requests,
        tokenOut
      );

      const filteredDestQueries = (routerQueries as any).filter((query: { minAmountOut: { eq: (arg0: number) => any } }) => !query.minAmountOut.eq(0));

      return filteredDestQueries.map((routerQuery: any) => {
        // Normalize key differences between router types
        return {
          ...routerQuery,
          swapAdapter: routerQuery.swapAdapter ?? routerQuery.routerAdapter,
        };
      });
    }

    public findBestQuery(
      destQueries: RawQuery[],
      originQueries: RawQuery[],
      bridgeTokens: { symbol: string; token: string }[]
    ): [RawQuery, RawQuery, { symbol: string; token: string }] {
      let maxAmountOut: BigNumber = BigNumber.from(0);
      let bestDestQuery: RawQuery | null = null;
      let bestOriginQuery: RawQuery | null = null;
      let bestBridgeToken: { symbol: string; token: string } | null = null;

      for (let i = 0; i < destQueries.length; i++) {
        if (destQueries[i].minAmountOut.gt(maxAmountOut)) {
          maxAmountOut = destQueries[i].minAmountOut;
          bestDestQuery = destQueries[i];
          bestOriginQuery = originQueries[i];
          bestBridgeToken = bridgeTokens[i];
        }
      }

      if (!bestDestQuery || !bestOriginQuery || !bestBridgeToken) {
        throw new Error('No best queries found');
      }

      return [bestOriginQuery, bestDestQuery, bestBridgeToken];
    }


    public async finalizeQuote(
      bestQuery: [RawQuery, RawQuery, { symbol: string; token: string }],
      router: SynapseRouter | SynapseCCTPRouter,
      deadline: BigNumber | undefined,
      isCCTP: boolean = false
    ): Promise<any> {
      const [bestOriginQuery, bestDestQuery, bestBridgeToken] = bestQuery;

      // Set default deadlines
      const originQuery = convertQuery(bestOriginQuery);
      originQuery.deadline = deadline ?? TEN_MIN_DEADLINE;
      const destQuery = convertQuery(bestDestQuery);
      destQuery.deadline = ONE_WEEK_DEADLINE;

      let feeAmount;
      let feeConfig;

      // Get fee data
      if (isCCTP) {
        // Cast router to SynapseCCTPRouter
        const cctpRouter = <SynapseCCTPRouter>router;

        feeAmount = await cctpRouter.routerContract.calculateFeeAmount(
          bestBridgeToken.token,
          originQuery.minAmountOut,
          //TODO: how to determine isSwap?
          false
        );

        const [relayerFee, minBaseFee, , maxFee] = await cctpRouter.routerContract.feeStructures(
          bestBridgeToken.token
        );

        // return only the compatible fields
        feeConfig = {
          bridgeFee: relayerFee, // Assuming relayerFee is compatible with bridgeFee
          minFee: minBaseFee, // assuming minBaseFee is compatible with minFee
          maxFee // maxFee is the same in both cases
        };
      } else {
        // Cast router to SynapseRouter
        const synapseRouter = <SynapseRouter>router;

        feeAmount = await synapseRouter.routerContract.calculateBridgeFee(
          bestBridgeToken.token,
          originQuery.minAmountOut
        );

        feeConfig = await synapseRouter.routerContract.fee(bestBridgeToken.token);
      }

      return {
        feeAmount,
        feeConfig,
        routerAddress: router.routerContract.address,
        maxAmountOut: destQuery.minAmountOut,
        originQuery,
        destQuery,
      };
    }

    public async bridgeQuote(
      originChainId: number,
      destChainId: number,
      tokenIn: string,
      tokenOut: string,
      amountIn: BigintIsh,
      deadline?: BigNumber
    ): Promise<BridgeQuote | undefined> {
      tokenOut = handleNativeToken(tokenOut);
      tokenIn = handleNativeToken(tokenIn);

      const originSynapseRouter = this.synapseRouters[originChainId];
      const destSynapseRouter = this.synapseRouters[destChainId];

      const synapseQuotePromise = this.calculateBestQuote(originSynapseRouter, destSynapseRouter, destChainId, tokenIn, tokenOut, amountIn, deadline);

      let cctpQuotePromise: Promise<BridgeQuote | undefined>;
      if(CCTP_ROUTER_ADDRESS[originChainId] && CCTP_ROUTER_ADDRESS[destChainId]) {
        const originCCTPRouter = this.synapseCCTPRouters[originChainId];
        const destCCTPRouter = this.synapseCCTPRouters[destChainId];
        cctpQuotePromise = this.calculateBestQuote(originCCTPRouter, destCCTPRouter, destChainId, tokenIn, tokenOut, amountIn, deadline);
      } else {
        cctpQuotePromise = Promise.resolve(undefined);
      }

      const [synapseQuote, cctpQuote] = await Promise.all([synapseQuotePromise, cctpQuotePromise]);

      const bestQuote = [synapseQuote, cctpQuote].reduce((prev, current) =>
        (!prev || (current && current.maxAmountOut && prev.maxAmountOut && current.maxAmountOut.gt(prev.maxAmountOut)) ? current : prev), undefined);

      if (!bestQuote) {
        throw new Error('No route found');
      }

      return bestQuote;
    }


private async calculateBestQuote(
  originRouter: SynapseRouter | SynapseCCTPRouter | undefined,
  destRouter: SynapseRouter | SynapseCCTPRouter | undefined,
  destChainId: number,
  tokenIn: string,
  tokenOut: string,
  amountIn: BigintIsh,
  deadline?: BigNumber
): Promise<BridgeQuote | undefined> {
  if (!originRouter || !destRouter) {
    return;
  }

  let bestQuote: BridgeQuote | undefined;

  // Getting bridge tokens from cache or fetch from destination router
  const bridgeTokensArray = await this.getBridgeTokens(destChainId, tokenOut, [destRouter]);

  // Iterate through each array of bridge tokens
  for (let bridgeTokens of bridgeTokensArray[0]) {
    // Fetching queries from origin router
    const originQueries = await this.getOriginQueries(originRouter, tokenIn, [bridgeTokens.symbol], amountIn);
    if (!originQueries.length) continue;  // Skip if no origin queries for these bridge tokens

    // Building request for destination queries
    const requests = originQueries.map((query) => ({ symbol: bridgeTokens.symbol, amountIn: query.minAmountOut }));
    // Fetching queries from destination router
    const destQueries = await this.getDestinationQueries(destRouter, requests, tokenOut);
    if (!destQueries.length) continue;  // Skip if no destination queries for these requests

    // Finding the best query
    const bestQuery = this.findBestQuery(destQueries, originQueries, [bridgeTokens]);

    // Finalizing quote
    const quote = await this.finalizeQuote(bestQuery, destRouter, deadline, destRouter instanceof SynapseCCTPRouter);
    // Check if this quote is better than previous best
    if (!bestQuote || quote.maxAmountOut.gt(bestQuote.maxAmountOut)) {
      bestQuote = quote;
    }
  }

  return bestQuote;
}




  // public async oldBridgeQuote(
  //   originChainId: number,
  //   destChainId: number,
  //   tokenIn: string,
  //   tokenOut: string,
  //   amountIn: BigintIsh,
  //   deadline?: BigNumber
  // ): Promise<{
  //   feeAmount: BigNumber | undefined
  //   feeConfig: FeeConfig | undefined
  //   routerAddress: string | undefined
  //   maxAmountOut: BigNumber | undefined
  //   originQuery: Query | undefined
  //   destQuery: Query | undefined
  // }> {
  //   tokenOut = handleNativeToken(tokenOut)
  //   tokenIn = handleNativeToken(tokenIn)
  //   const originRouter: SynapseRouter = this.synapseRouters[originChainId]
  //   const destRouter: SynapseRouter = this.synapseRouters[destChainId]
  //   const routerAddress = originRouter.routerContract.address

  //   let bridgeTokens = this.bridgeTokenCache[destChainId + '_' + tokenOut]
  //   if (!bridgeTokens) {
  //     const routerBridgeTokens =
  //       await destRouter.routerContract.getConnectedBridgeTokens(tokenOut)

  //     // Filter tokens with a valid symbol and address
  //     bridgeTokens = routerBridgeTokens.filter(
  //       (bridgeToken) =>
  //         bridgeToken.symbol.length && bridgeToken.token !== AddressZero
  //     )

  //     // Throw error if no valid bridge tokens found
  //     if (!bridgeTokens?.length) {
  //       throw new Error('No bridge tokens found for this route')
  //     }

  //     // Store only the symbol and token fields in the cache
  //     bridgeTokens = bridgeTokens.map(({ symbol, token }) => ({
  //       symbol,
  //       token,
  //     }))

  //     // Cache the bridge tokens
  //     this.bridgeTokenCache[destChainId + '_' + tokenOut] = bridgeTokens
  //   }

  //   // Get quotes from origin SynapseRouter
  //   const originQueries: RawQuery[] =
  //     await originRouter.routerContract.getOriginAmountOut(
  //       tokenIn,
  //       bridgeTokens.map((bridgeToken) => bridgeToken.symbol),
  //       amountIn
  //     )

  //   // create requests for destination router
  //   const requests: { symbol: string; amountIn: BigintIsh }[] = []
  //   for (let i = 0; i < bridgeTokens.length; i++) {
  //     requests.push({
  //       symbol: bridgeTokens[i].symbol,
  //       amountIn: originQueries[i].minAmountOut,
  //     })
  //   }
  //   if (originQueries.length === 0) {
  //     throw Error('No origin queries found for this route')
  //   }

  //   // Get quotes from destination SynapseRouter
  //   const destQueries: RawQuery[] =
  //     await destRouter.routerContract.getDestinationAmountOut(
  //       requests,
  //       tokenOut
  //     )
  //   if (destQueries.length === 0) {
  //     throw Error('No destination queries found for this route')
  //   }

  //   // Find the best query (in practice, we could return them all)
  //   let destInToken
  //   let rawOriginQuery
  //   let rawDestQuery
  //   let maxAmountOut: BigNumber = BigNumber.from(0)
  //   for (let i = 0; i < destQueries.length; i++) {
  //     if (destQueries[i].minAmountOut.gt(maxAmountOut)) {
  //       maxAmountOut = destQueries[i].minAmountOut
  //       rawOriginQuery = originQueries[i]
  //       rawDestQuery = destQueries[i]
  //       destInToken = bridgeTokens[i].token
  //     }
  //   }

  //   if (!rawOriginQuery || !rawDestQuery) {
  //     throw Error('No route found')
  //   }

  //   // Set default deadline
  //   const originQuery = convertQuery(rawOriginQuery)
  //   originQuery.deadline = deadline ?? TEN_MIN_DEADLINE
  //   const destQuery = convertQuery(rawDestQuery)
  //   destQuery.deadline = ONE_WEEK_DEADLINE

  //   // Get fee data
  //   let feeAmount
  //   let feeConfig

  //   if (originQuery && destInToken) {
  //     feeAmount = destRouter.routerContract.calculateBridgeFee(
  //       destInToken,
  //       originQuery.minAmountOut
  //     )
  //     feeConfig = destRouter.routerContract.fee(destInToken)
  //   }

  //   return {
  //     feeAmount: await feeAmount,
  //     feeConfig: await feeConfig,
  //     routerAddress,
  //     maxAmountOut,
  //     originQuery,
  //     destQuery,
  //   }
  // }

  public async bridge(
    to: string,
    originRouterAddress: string,
    originChainId: number,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    token = handleNativeToken(token)
    let bridgeOriginQuery: {
      swapAdapter?: string
      routerAdapter?: string
      tokenOut: string
      minAmountOut: BigNumber
      deadline: BigNumber
      rawParams: string } = { ...originQuery}

      let bridgeDestQuery: {
        swapAdapter?: string
        routerAdapter: string
        tokenOut: string
        minAmountOut: BigNumber
        deadline: BigNumber
        rawParams: string
      } = { ...destQuery, routerAdapter: destQuery.swapAdapter}

    let isCCTP = (originRouterAddress.toLowerCase() === this.synapseCCTPRouters[originChainId].routerContract.address.toLowerCase())
    // Determine which router to use based on the address
    const originRouter: SynapseRouter | SynapseCCTPRouter = isCCTP ? this.synapseCCTPRouters[originChainId] : this.synapseRouters[originChainId];

    if (isCCTP) {
      bridgeOriginQuery = {
        ...originQuery,
        routerAdapter: originQuery.swapAdapter,
      };
      delete bridgeOriginQuery.swapAdapter;

      bridgeDestQuery = {
        ...destQuery,
        routerAdapter: destQuery.swapAdapter,
      };
      delete bridgeDestQuery.swapAdapter;
    }

    // Call the bridge method on the selected router
    return (originRouter as any).routerContract.populateTransaction.bridge(
      to,
      destChainId,
      token,
      amount,
      bridgeOriginQuery,
      bridgeDestQuery
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

    // Check if call was unsuccessful.
    if (rawQuery?.length !== 5 || rawQuery.minAmountOut.isZero()) {
      throw Error('No queries found for this route')
    }

    const query = convertQuery(rawQuery)
    query.deadline = deadline ?? TEN_MIN_DEADLINE
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
    amounts: Array<{ value: BigNumber; index: number }>
    routerAddress: string
  }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const amounts = await router.routerContract.calculateRemoveLiquidity(
      poolAddress,
      amount
    )
    const amountsOut: Array<{ value: BigNumber; index: number }> = amounts.map(
      (respAmount, index) => ({
        value: respAmount,
        index,
      })
    )

    return {
      amounts: amountsOut,
      routerAddress: router.routerContract.address,
    }
  }

  public async calculateRemoveLiquidityOne(
    chainId: number,
    poolAddress: string,
    amount: BigNumber,
    poolIndex: number
  ): Promise<{
    amount: { value: BigNumber; index: number }
    routerAddress: string
  }> {
    const router: SynapseRouter = this.synapseRouters[chainId]

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
