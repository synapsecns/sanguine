import { Interface } from '@ethersproject/abi'
import { Zero } from '@ethersproject/constants'
import { BigNumber, utils } from 'ethers'

import { SupportedChainId } from '../../constants'
import {
  getWithTimeout,
  isNativeToken,
  isSameAddress,
  logExecutionTime,
  logger,
  Prettify,
  TokenMetadataFetcher,
} from '../../utils'
import { applySlippage, EngineID, SlippageMax } from '../core'
import {
  getEmptyQuote,
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from '../models'
import { generateAPIRoute } from './response'

const LIQUID_SWAP_API_URL = 'https://api.liqd.ag'
// TODO: remove custom timeout
const LIQUID_SWAP_API_TIMEOUT = 3000

const WHYPE = '0x5555555555555555555555555555555555555555'
const HYPE = '0x000000000000000000000000000000000000dEaD'
const LIQUID_SWAP_ROUTER = '0x744489Ee3d540777A66f2cf297479745e0852f7A'

// MultiHopRouter ABI (just the executeMultiHopSwap function)
const ROUTER_ABI = [
  'function executeMultiHopSwap(address[] calldata tokens, uint256 amountIn, uint256 minAmountOut, tuple(address tokenIn, address tokenOut, uint8 routerIndex, uint24 fee, uint256 amountIn, bool stable)[][] calldata hopSwaps) external payable returns (uint256 totalAmountOut)',
]

type LiquidSwapRouteRequest = {
  tokenA: string
  tokenB: string
  // Human-readable amounts (e.g., 1.5)
  amountIn?: string
  amountOut?: string
  multiHop?: boolean
  excludeDexes?: string
}

type LiquidSwapTokenInfo = {
  address: string
  symbol: string
  name: string
  decimals: number
}

type LiquidSwapData = {
  status: string
  tokenInfo: {
    [key: string]: LiquidSwapTokenInfo
  } & {
    amountIn: string
  }
  bestPath: {
    amountOut: string
    hop: [
      {
        hopAmountIn: string
        hopAmountOut: string
        tokenIn: string
        tokenOut: string
        allocations: [
          {
            poolAddress: string
            routerIndex: number
            routerName: string
            fee: number
            stable: boolean
            percentage: number
            amountIn: string
            tokenIn: string
            tokenOut: string
          }
        ]
      }
    ]
  }
}

type LiquidSwapRouteResponse = {
  success: boolean
  data?: LiquidSwapData
}

type LiquidSwapQuote = Prettify<
  SwapEngineQuote & {
    data?: LiquidSwapData
  }
>

type HopSwap = {
  tokenIn: string
  tokenOut: string
  routerIndex: number
  fee: number
  amountIn: string
  stable: boolean
}

export class LiquidSwapEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.LiquidSwap

  static routerInterface = new Interface(ROUTER_ABI)

  private tokenMetadataFetcher: TokenMetadataFetcher

  constructor(tokenMetadataFetcher: TokenMetadataFetcher) {
    this.tokenMetadataFetcher = tokenMetadataFetcher
  }

  @logExecutionTime('LiquidSwapEngine.getQuote')
  public async getQuote(input: RouteInput): Promise<LiquidSwapQuote> {
    const { chainId, fromToken, toToken, fromAmount } = input
    if (
      chainId !== SupportedChainId.HYPEREVM ||
      isSameAddress(fromToken, toToken) ||
      BigNumber.from(fromAmount).eq(Zero)
    ) {
      return getEmptyQuote(this.id)
    }
    const [fromTokenDecimals, toTokenDecimals] = await Promise.all([
      this.tokenMetadataFetcher.getTokenDecimals(chainId, fromToken),
      this.tokenMetadataFetcher.getTokenDecimals(chainId, toToken),
    ])
    // Convert native token to WHYPE for getting the quote.
    // We will handle the reverse conversion in the route generation.
    const request: LiquidSwapRouteRequest = {
      tokenA: this.transformNativeToken(fromToken),
      tokenB: this.transformNativeToken(toToken),
      amountIn: utils.formatUnits(fromAmount, fromTokenDecimals),
      multiHop: true,
    }
    const response = await this.getRouteResponse(
      request,
      LIQUID_SWAP_API_TIMEOUT
    )
    if (!response) {
      return getEmptyQuote(this.id)
    }
    const liquidSwapResponse: LiquidSwapRouteResponse = await response.json()
    const data = liquidSwapResponse.data
    if (
      !liquidSwapResponse.success ||
      !data ||
      data.status !== 'success' ||
      !data.tokenInfo ||
      !data.bestPath
    ) {
      return getEmptyQuote(this.id)
    }
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      fromToken,
      toToken,
      fromAmount: BigNumber.from(fromAmount),
      expectedToAmount: utils.parseUnits(
        data.bestPath.amountOut,
        toTokenDecimals
      ),
      data,
    }
  }

  public async generateRoute(
    input: RouteInput,
    quote: LiquidSwapQuote
  ): Promise<SwapEngineRoute> {
    if (quote.engineID !== this.id || !quote.data) {
      logger.error({ quote }, 'LiquidSwapEngine: unexpected quote')
      return getEmptyRoute(this.id)
    }
    const tokens: string[] = []
    const hopSwaps: HopSwap[][] = []
    const tokenInfoList = Object.values(quote.data.tokenInfo).filter(
      (tokenInfo) => typeof tokenInfo !== 'string'
    )
    for (const hopData of quote.data.bestPath.hop) {
      // Fill the tokens array with the tokenOut of each hop
      // Also add the tokenIn of the first hop to the tokens array
      if (tokens.length === 0) {
        tokens.push(hopData.tokenIn)
      }
      tokens.push(hopData.tokenOut)
      const tokenInInfo = tokenInfoList.find((tokenInfo) =>
        isSameAddress(tokenInfo.address, hopData.tokenIn)
      )
      if (!tokenInInfo) {
        logger.error(
          { data: quote.data, token: hopData.tokenIn },
          'LiquidSwapEngine: unexpected token'
        )
        return getEmptyRoute(this.id)
      }
      hopSwaps.push(
        hopData.allocations.map((allocation) => ({
          tokenIn: tokenInInfo.address,
          tokenOut: hopData.tokenOut,
          routerIndex: allocation.routerIndex,
          fee: allocation.fee,
          amountIn: utils
            .parseUnits(allocation.amountIn, tokenInInfo.decimals)
            .toString(),
          stable: allocation.stable,
        }))
      )
    }
    if (tokens.length === 0 || hopSwaps.length === 0) {
      logger.error({ data: quote.data }, 'LiquidSwapEngine: no hops found')
      return getEmptyRoute(this.id)
    }
    // Change fromToken and toToken to HYPE if needed
    if (isNativeToken(input.fromToken)) {
      tokens[0] = HYPE
    }
    if (isNativeToken(input.toToken)) {
      tokens[tokens.length - 1] = HYPE
    }
    return generateAPIRoute(input, this.id, SlippageMax, {
      expectedToAmount: quote.expectedToAmount,
      transaction: {
        chainId: SupportedChainId.HYPEREVM,
        to: LIQUID_SWAP_ROUTER,
        value: isNativeToken(input.fromToken)
          ? input.fromAmount.toString()
          : '0',
        data: LiquidSwapEngine.routerInterface.encodeFunctionData(
          'executeMultiHopSwap',
          [
            tokens,
            input.fromAmount,
            applySlippage(quote.expectedToAmount, SlippageMax),
            hopSwaps,
          ]
        ),
      },
    })
  }

  public async getRouteResponse(
    params: LiquidSwapRouteRequest,
    timeout: number
  ): Promise<Response | null> {
    const url = `${LIQUID_SWAP_API_URL}/route`
    return getWithTimeout('LiquidSwap', url, timeout, params)
  }

  private transformNativeToken(address: string): string {
    return isNativeToken(address) ? WHYPE : address
  }
}
