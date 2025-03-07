import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { Provider } from '@ethersproject/abstract-provider'
import { Contract } from '@ethersproject/contracts'

import erc20ABI from '../../abi/IERC20Metadata.json'
import { IERC20Metadata as ERC20 } from '../../typechain/IERC20Metadata'
import { getWithTimeout, postWithTimeout } from '../../utils/api'
import { isSameAddress } from '../../utils/addressUtils'
import { isNativeToken } from '../../utils/handleNativeToken'
import { logExecutionTime, logger } from '../../utils/logger'
import { ChainProvider } from '../../router'
import { marshallChainToken } from '../../rfq/ticker'
import { EngineID, SlippageMax, toBasisPoints } from '../core'
import {
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from '../models'
import { generateAPIRoute, TransactionData } from './response'

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

type ParaSwapQuote = SwapEngineQuote & {
  priceRoute: ParaSwapPriceRoute
}

const EmptyParaSwapQuote: ParaSwapQuote = {
  engineID: EngineID.ParaSwap,
  engineName: EngineID[EngineID.ParaSwap],
  chainId: 0,
  expectedAmountOut: Zero,
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
  private decimalsCache: {
    [tokenId: string]: number
  }

  constructor(chains: ChainProvider[]) {
    this.providers = {}
    this.decimalsCache = {}
    chains.forEach(({ chainId, provider }) => {
      this.providers[chainId] = provider
    })
  }

  public async getQuote(
    input: RouteInput,
    timeout: number
  ): Promise<ParaSwapQuote> {
    const { chainId, tokenIn, tokenOut, msgSender, amountIn } = input
    if (isSameAddress(tokenIn, tokenOut) || BigNumber.from(amountIn).eq(Zero)) {
      return EmptyParaSwapQuote
    }
    const response = await this.getPricesResponse(
      {
        srcToken: tokenIn,
        srcDecimals: await this.getTokenDecimals(chainId, tokenIn),
        destToken: tokenOut,
        destDecimals: await this.getTokenDecimals(chainId, tokenOut),
        amount: amountIn.toString(),
        side: 'SELL',
        network: chainId,
        excludeRFQ: true,
        userAddress: msgSender,
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
      expectedAmountOut: BigNumber.from(paraSwapResponse.priceRoute.destAmount),
      priceRoute: paraSwapResponse.priceRoute,
    }
  }

  public async generateRoute(
    input: RouteInput,
    quote: ParaSwapQuote,
    timeout: number
  ): Promise<SwapEngineRoute> {
    const { chainId, msgSender } = input
    if (quote.engineID !== this.id || !quote.priceRoute) {
      logger.error({ quote }, 'ParaSwap: unexpected quote')
      return getEmptyRoute(this.id)
    }
    const response = await this.getTransactionsResponse(
      chainId,
      {
        srcToken: input.tokenIn,
        srcDecimals: quote.priceRoute.srcDecimals,
        destToken: input.tokenOut,
        destDecimals: quote.priceRoute.destDecimals,
        srcAmount: input.amountIn.toString(),
        priceRoute: quote.priceRoute,
        slippage: toBasisPoints(SlippageMax),
        userAddress: msgSender,
      },
      timeout
    )
    if (!response) {
      return getEmptyRoute(this.id)
    }
    const paraSwapResponse: ParaSwapTransactionsResponse = await response.json()
    return generateAPIRoute(input, this.id, SlippageMax, {
      amountOut: BigNumber.from(quote.priceRoute.destAmount),
      transaction: paraSwapResponse,
    })
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

  private async getTokenDecimals(
    chainId: number,
    token: string
  ): Promise<number> {
    // TODO: move to utils
    if (isNativeToken(token)) {
      return 18
    }
    const tokenId = marshallChainToken({ chainId, token })
    if (this.decimalsCache[tokenId]) {
      return this.decimalsCache[tokenId]
    }
    const provider = this.providers[chainId]
    if (!provider) {
      logger.error(`No provider found for chainId: ${chainId}`)
      return 0
    }
    const tokenContract = new Contract(token, erc20ABI, provider) as ERC20
    try {
      const decimals = await tokenContract.decimals()
      this.decimalsCache[tokenId] = decimals
      return decimals
    } catch (error) {
      logger.error({ error, chainId, token }, 'Error fetching token decimals')
      return 0
    }
  }
}
