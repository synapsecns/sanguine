import { Provider } from '@ethersproject/abstract-provider'
import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { generateAPIRoute, TransactionData } from './response'
import { ChainProvider } from '../../router'
import {
  getWithTimeout,
  postWithTimeout,
  isSameAddress,
  TokenMetadataFetcher,
  logger,
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
  srcDecimals: number
  destDecimals: number
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

export type ParaSwapTransactionsResponse = TransactionData

type ParaSwapQuote = Prettify<
  SwapEngineQuote & {
    priceRoute: ParaSwapPriceRoute
  }
>

const EmptyParaSwapQuote: ParaSwapQuote = {
  ...getEmptyQuote(EngineID.ParaSwap),
  priceRoute: {
    srcDecimals: 0,
    destDecimals: 0,
    destAmount: '0',
  },
}

export class ParaSwapEngine implements SwapEngine {
  readonly id: EngineID = EngineID.ParaSwap

  private providers: {
    [chainId: number]: Provider
  }
  private tokenMetadataFetcher: TokenMetadataFetcher

  constructor(
    chains: ChainProvider[],
    tokenMetadataFetcher: TokenMetadataFetcher
  ) {
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      this.providers[chainId] = provider
    })
    this.tokenMetadataFetcher = tokenMetadataFetcher
  }

  public async getQuote(
    input: RouteInput,
    timeout: number
  ): Promise<ParaSwapQuote> {
    const { chainId, fromToken, toToken, swapper, fromAmount } = input
    if (
      isSameAddress(fromToken, toToken) ||
      BigNumber.from(fromAmount).eq(Zero)
    ) {
      return EmptyParaSwapQuote
    }
    const srcDecimals = await this.tokenMetadataFetcher.getTokenDecimals(
      chainId,
      fromToken
    )
    const destDecimals = await this.tokenMetadataFetcher.getTokenDecimals(
      chainId,
      toToken
    )
    const response = await this.getPricesResponse(
      {
        srcToken: fromToken,
        srcDecimals,
        destToken: toToken,
        destDecimals,
        amount: fromAmount.toString(),
        side: 'SELL',
        network: chainId,
        excludeRFQ: true,
        userAddress: swapper,
        version: '6.2',
      },
      timeout
    )
    if (!response) {
      return EmptyParaSwapQuote
    }
    const paraSwapResponse: ParaSwapPricesResponse = await response.json()
    if (!paraSwapResponse.priceRoute?.destAmount) {
      return EmptyParaSwapQuote
    }
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      fromToken,
      toToken,
      fromAmount: BigNumber.from(fromAmount),
      expectedToAmount: BigNumber.from(paraSwapResponse.priceRoute.destAmount),
      priceRoute: paraSwapResponse.priceRoute,
    }
  }

  public async generateRoute(
    input: RouteInput,
    quote: ParaSwapQuote,
    timeout: number
  ): Promise<SwapEngineRoute> {
    const { chainId, swapper } = input
    if (quote.engineID !== this.id || !quote.priceRoute) {
      logger.error({ quote }, 'ParaSwap: unexpected quote')
      return getEmptyRoute(this.id)
    }
    const response = await this.getTransactionsResponse(
      chainId,
      {
        srcToken: input.fromToken,
        srcDecimals: quote.priceRoute.srcDecimals,
        destToken: input.toToken,
        destDecimals: quote.priceRoute.destDecimals,
        srcAmount: input.fromAmount.toString(),
        priceRoute: quote.priceRoute,
        slippage: toBasisPoints(SlippageMax),
        userAddress: swapper,
      },
      timeout
    )
    if (!response) {
      return getEmptyRoute(this.id)
    }
    const paraSwapResponse: ParaSwapTransactionsResponse = await response.json()
    return generateAPIRoute(input, this.id, SlippageMax, {
      expectedToAmount: BigNumber.from(quote.priceRoute.destAmount),
      transaction: paraSwapResponse,
    })
  }

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
