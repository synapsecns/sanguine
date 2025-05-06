import { BigNumber, utils } from 'ethers'
import { Zero } from '@ethersproject/constants'

import {
  getWithTimeout,
  isNativeToken,
  isSameAddress,
  logger,
  Prettify,
  TokenMetadataFetcher,
} from '../../utils'
import { EngineID } from '../core'
import {
  getEmptyQuote,
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from '../models'
import { SupportedChainId } from '../../constants'

const LIQUID_SWAP_API_URL = 'https://api.liqd.ag'
const WHYPE = '0x5555555555555555555555555555555555555555'

type LiquidSwapRouteRequest = {
  tokenA: string
  tokenB: string
  // Human-readable amounts (e.g., 1.5)
  amountIn?: number
  amountOut?: number
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

export class LiquidSwapEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.LiquidSwap

  private tokenMetadataFetcher: TokenMetadataFetcher

  constructor(tokenMetadataFetcher: TokenMetadataFetcher) {
    this.tokenMetadataFetcher = tokenMetadataFetcher
  }

  public async getQuote(
    input: RouteInput,
    timeout: number
  ): Promise<LiquidSwapQuote> {
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
      amountIn: Number(utils.formatUnits(fromAmount, fromTokenDecimals)),
      multiHop: true,
    }
    const response = await this.getRouteResponse(request, timeout)
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
    // TODO: implement route generation logic
    return getEmptyRoute(this.id)
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
