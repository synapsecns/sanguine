import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { SupportedChainId } from '../../constants'
import {
  isSameAddress,
  getWithTimeout,
  postWithTimeout,
  ONE_WEEK,
  isNativeToken,
  logger,
  logExecutionTime,
  Prettify,
} from '../../utils'
import { EngineID, SlippageMax, toBasisPoints } from '../core'
import {
  getEmptyQuote,
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from '../models'
import { generateAPIRoute } from './response'

const KYBER_SWAP_API_URL = 'https://aggregator-api.kyberswap.com'

type KyberSwapQuoteRequest = {
  tokenIn: string
  tokenOut: string
  amountIn: string
  gasInclude: boolean
  onlySinglePath?: boolean
}

type KyberSwapRouteSummary = {
  amountOut: string
}

export type KyberSwapQuoteResponse = {
  code: number
  message: string
  data: {
    routeSummary: KyberSwapRouteSummary
    routerAddress: string
  }
}

type KyberSwapBuildRequest = {
  routeSummary: KyberSwapRouteSummary
  sender: string
  recipient: string
  deadline: number
  slippageTolerance: number
  enableGasEstimation: boolean
}

export type KyberSwapBuildResponse = {
  code: number
  message: string
  data: {
    routerAddress: string
    data: string
  }
}

type KyberSwapQuote = Prettify<
  SwapEngineQuote & {
    routeSummary: KyberSwapRouteSummary
  }
>

const EmptyKyberSwapQuote: KyberSwapQuote = {
  ...getEmptyQuote(EngineID.KyberSwap),
  routeSummary: {
    amountOut: '0',
  },
}

/**
 * Chains supported by KyberSwap, ordered lexicographically.
 */
const KyberSwapChainMap: Record<number, string> = {
  [SupportedChainId.ARBITRUM]: 'arbitrum',
  [SupportedChainId.AVALANCHE]: 'avalanche',
  [SupportedChainId.BASE]: 'base',
  [SupportedChainId.BERACHAIN]: 'berachain',
  [SupportedChainId.BLAST]: 'blast',
  [SupportedChainId.BSC]: 'bsc',
  [SupportedChainId.ETH]: 'ethereum',
  [SupportedChainId.FANTOM]: 'fantom',
  [SupportedChainId.LINEA]: 'linea',
  [SupportedChainId.OPTIMISM]: 'optimism',
  [SupportedChainId.POLYGON]: 'polygon',
  [SupportedChainId.SCROLL]: 'scroll',
}

export class KyberSwapEngine implements SwapEngine {
  readonly id: EngineID = EngineID.KyberSwap

  public async getQuote(
    input: RouteInput,
    timeout: number
  ): Promise<KyberSwapQuote> {
    const { chainId, fromToken, toToken, fromAmount } = input
    if (
      isSameAddress(fromToken, toToken) ||
      BigNumber.from(fromAmount).eq(Zero)
    ) {
      return EmptyKyberSwapQuote
    }
    const request: KyberSwapQuoteRequest = {
      tokenIn: fromToken,
      tokenOut: toToken,
      amountIn: fromAmount.toString(),
      gasInclude: true,
      onlySinglePath: input.restrictComplexity,
    }
    const response = await this.getQuoteResponse(
      input.chainId,
      request,
      timeout
    )
    if (!response) {
      return EmptyKyberSwapQuote
    }
    const kyberSwapQuoteResponse: KyberSwapQuoteResponse = await response.json()
    if (
      kyberSwapQuoteResponse.code !== 0 ||
      !kyberSwapQuoteResponse.data?.routeSummary
    ) {
      return EmptyKyberSwapQuote
    }
    const expectedToAmount = BigNumber.from(
      kyberSwapQuoteResponse.data.routeSummary.amountOut ?? '0'
    )
    if (expectedToAmount.eq(Zero)) {
      return EmptyKyberSwapQuote
    }
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      fromToken,
      toToken,
      fromAmount: BigNumber.from(fromAmount),
      expectedToAmount,
      routeSummary: kyberSwapQuoteResponse.data.routeSummary,
    }
  }

  public async generateRoute(
    input: RouteInput,
    quote: KyberSwapQuote,
    timeout: number
  ): Promise<SwapEngineRoute> {
    const { chainId, swapper } = input
    if (quote.engineID !== this.id || !quote.routeSummary) {
      logger.error({ quote }, 'KyberSwap: unexpected quote')
      return getEmptyRoute(this.id)
    }
    const response = await this.getBuildResponse(
      chainId,
      {
        routeSummary: quote.routeSummary,
        sender: swapper,
        recipient: swapper,
        deadline: Math.floor(Date.now() / 1000) + ONE_WEEK,
        slippageTolerance: toBasisPoints(SlippageMax),
        enableGasEstimation: false,
      },
      timeout
    )
    if (!response) {
      return getEmptyRoute(this.id)
    }
    const kyberSwapBuildResponse: KyberSwapBuildResponse = await response.json()
    if (kyberSwapBuildResponse.code !== 0) {
      return getEmptyRoute(this.id)
    }
    return generateAPIRoute(input, this.id, SlippageMax, {
      expectedToAmount: quote.expectedToAmount,
      transaction: {
        chainId,
        from: swapper,
        to: kyberSwapBuildResponse.data.routerAddress,
        value: isNativeToken(input.fromToken)
          ? input.fromAmount.toString()
          : '0',
        data: kyberSwapBuildResponse.data.data,
      },
    })
  }

  @logExecutionTime('KyberSwapEngine.getQuoteResponse')
  public async getQuoteResponse(
    chainId: number,
    params: KyberSwapQuoteRequest,
    timeout: number
  ): Promise<Response | null> {
    const chain = KyberSwapChainMap[chainId]
    if (!chain) {
      return null
    }
    const url = `${this.buildBaseURL(chain)}/routes`
    return getWithTimeout('KyberSwap', url, timeout, params, {
      'x-client-id': 'SynapseIntentNetwork',
    })
  }

  @logExecutionTime('KyberSwapEngine.getBuildResponse')
  public async getBuildResponse(
    chainId: number,
    params: KyberSwapBuildRequest,
    timeout: number
  ): Promise<Response | null> {
    const chain = KyberSwapChainMap[chainId]
    if (!chain) {
      return null
    }
    const url = `${this.buildBaseURL(chain)}/route/build`
    return postWithTimeout('KyberSwap', url, timeout, params, {
      'x-client-id': 'SynapseIntentNetwork',
    })
  }

  private buildBaseURL(chain: string): string {
    return `${KYBER_SWAP_API_URL}/${chain}/api/v1`
  }
}
