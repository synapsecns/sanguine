import { Zero } from '@ethersproject/constants'

import { getWithTimeout, postWithTimeout } from '../api'
import {
  SwapEngine,
  EngineID,
  SwapEngineRoute,
  getEmptyRoute,
  SwapEngineQuote,
} from './swapEngine'
import { logExecutionTime } from '../../utils/logger'

const PARASWAP_API_URL = 'https://api.paraswap.io'

export type ParaSwapPricesRequest = {
  srcToken: string
  srcDecimals: number
  destToken: string
  destDecimals: number
  amount: string
  side: string
  network: number
  excludeRFQ: boolean
  userAddress: string
  partner?: string
  version: string
}

export type ParaSwapPriceRoute = {
  destAmount: string
}

export type ParaSwapPricesResponse = {
  priceRoute: ParaSwapPriceRoute
}

export type ParaSwapTransactionsRequest = {
  srcToken: string
  srcDecimals: number
  destToken: string
  destDecimals: number
  srcAmount: string
  priceRoute: ParaSwapPriceRoute
  slippage: number
  userAddress: string
}

export type ParaSwapTransactionsResponse = {
  from: string
  to: string
  value: string
  data: string
  gasPrice: string
  chainId: number
}

type ParaSwapQuote = SwapEngineQuote & {
  priceRoute: ParaSwapPriceRoute
}

const EmptyParaSwapQuote: ParaSwapQuote = {
  engineID: EngineID.ParaSwap,
  chainId: 0,
  expectedAmountOut: Zero,
  priceRoute: {
    destAmount: '0',
  },
}

export class ParaSwapEngine implements SwapEngine {
  readonly id: EngineID = EngineID.ParaSwap

  public async getQuote(): Promise<ParaSwapQuote> {
    return EmptyParaSwapQuote
  }

  public async generateRoute(): Promise<SwapEngineRoute> {
    return getEmptyRoute(this.id)
  }

  @logExecutionTime('ParaSwapEngine.getPricesResponse')
  public async getPricesResponse(
    params: ParaSwapPricesRequest,
    timeout: number
  ): Promise<Response | null> {
    return getWithTimeout(
      'ParaSwap',
      `${PARASWAP_API_URL}/prices`,
      timeout,
      params
    )
  }

  @logExecutionTime('ParaSwapEngine.getTransactionsResponse')
  public async getTransactionsResponse(
    chainId: number,
    params: ParaSwapTransactionsRequest,
    timeout: number
  ): Promise<Response | null> {
    return postWithTimeout(
      'ParaSwap',
      `${PARASWAP_API_URL}/transactions/${chainId}?ignoreChecks=true`,
      timeout,
      params
    )
  }
}
