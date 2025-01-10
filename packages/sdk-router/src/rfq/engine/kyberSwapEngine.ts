import { Zero } from '@ethersproject/constants'

import { SupportedChainId } from '../../constants'
import { getWithTimeout, postWithTimeout } from '../api'
import {
  SwapEngine,
  EngineID,
  SwapEngineRoute,
  getEmptyRoute,
  SwapEngineQuote,
} from './swapEngine'
import { logger } from '../../utils/logger'

const KYBER_SWAP_API_URL = 'https://aggregator-api.kyberswap.com'

type KyberSwapQuoteRequest = {
  tokenIn: string
  tokenOut: string
  amountIn: string
  gasInclude: boolean
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

type KyberSwapQuote = SwapEngineQuote

const EmptyKyberSwapQuote: KyberSwapQuote = {
  engineID: EngineID.KyberSwap,
  expectedAmountOut: Zero,
}

const KyberSwapChainMap: Record<number, string> = {
  [SupportedChainId.ETH]: 'ethereum',
  [SupportedChainId.OPTIMISM]: 'optimism',
  [SupportedChainId.BSC]: 'bsc',
  [SupportedChainId.POLYGON]: 'polygon',
  [SupportedChainId.FANTOM]: 'fantom',
  [SupportedChainId.BASE]: 'base',
  [SupportedChainId.ARBITRUM]: 'arbitrum',
  [SupportedChainId.AVALANCHE]: 'avalanche',
  [SupportedChainId.LINEA]: 'linea',
  [SupportedChainId.BLAST]: 'blast',
  [SupportedChainId.SCROLL]: 'scroll',
}

export class KyberSwapEngine implements SwapEngine {
  readonly id: EngineID = EngineID.KyberSwap

  public async getQuote(): Promise<KyberSwapQuote> {
    // TODO: implement
    return EmptyKyberSwapQuote
  }

  public async generateRoute(): Promise<SwapEngineRoute> {
    // TODO: implement
    return getEmptyRoute(this.id)
  }

  public async getQuoteResponse(
    chainId: number,
    params: KyberSwapQuoteRequest,
    timeout: number
  ): Promise<Response | null> {
    const startTime = Date.now()
    const chain = KyberSwapChainMap[chainId]
    if (!chain) {
      return null
    }
    const url = `${this.buildBaseURL(chain)}/routes`
    const response = await getWithTimeout('KyberSwap', url, timeout, params, {
      'x-client-id': 'SynapseIntentNetwork',
    })
    logger.info(`KyberSwap quote response time: ${Date.now() - startTime}ms`)
    return response
  }

  public async getBuildResponse(
    chainId: number,
    params: KyberSwapBuildRequest,
    timeout: number
  ): Promise<Response | null> {
    const startTime = Date.now()
    const chain = KyberSwapChainMap[chainId]
    if (!chain) {
      return null
    }
    const url = `${this.buildBaseURL(chain)}/route/build`
    const response = await postWithTimeout('KyberSwap', url, timeout, params, {
      'x-client-id': 'SynapseIntentNetwork',
    })
    logger.info(`KyberSwap build response time: ${Date.now() - startTime}ms`)
    return response
  }

  private buildBaseURL(chain: string): string {
    return `${KYBER_SWAP_API_URL}/${chain}/api/v1`
  }
}
