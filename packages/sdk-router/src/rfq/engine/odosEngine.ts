import { BigNumber } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'

import { postWithTimeout } from '../api'
import { generateAPIRoute } from './response'
import {
  SwapEngine,
  EngineID,
  SwapEngineRoute,
  RouteInput,
  getEmptyRoute,
  toPercentFloat,
  SlippageMax,
  SwapEngineQuote,
} from './swapEngine'
import { AddressMap } from '../../constants'
import { isSameAddress } from '../../utils/addressUtils'
import { isNativeToken } from '../../utils/handleNativeToken'

const ODOS_API_URL = 'https://api.odos.xyz/sor'
const ODOS_API_TIMEOUT = 2000

type OdosQuoteRequest = {
  chainId: number
  inputTokens: {
    amount: string
    tokenAddress: string
  }[]
  outputTokens: {
    proportion: number
    tokenAddress: string
  }[]
  userAddr: string
  slippageLimitPercent: number
  referralCode?: number
  simple?: boolean
}

export type OdosQuoteResponse = {
  pathId: string
  outAmounts: string[]
}

type OdosAssembleRequest = {
  userAddr: string
  pathId: string
  receiver?: string
}

type OdosAssembleResponse = {
  transaction: {
    chainId: number
    gas: number
    gasPrice: number
    value: string
    to: string
    from: string
    data: string
  }
}

type OdosQuote = SwapEngineQuote & {
  assembleRequest: OdosAssembleRequest
}

const EmptyOdosQuote: OdosQuote = {
  engineID: EngineID.Odos,
  expectedAmountOut: Zero,
  assembleRequest: {
    userAddr: AddressZero,
    pathId: '',
  },
}

export class OdosEngine implements SwapEngine {
  readonly id: EngineID = EngineID.Odos

  private readonly tokenZapAddressMap: AddressMap

  constructor(tokenZapAddressMap: AddressMap) {
    this.tokenZapAddressMap = tokenZapAddressMap
  }

  public async getQuote(input: RouteInput): Promise<OdosQuote> {
    const { chainId, tokenIn, tokenOut, amountIn } = input
    const tokenZap = this.tokenZapAddressMap[chainId]
    if (
      !tokenZap ||
      isSameAddress(tokenIn, tokenOut) ||
      BigNumber.from(amountIn).eq(Zero)
    ) {
      return EmptyOdosQuote
    }
    const request: OdosQuoteRequest = {
      chainId,
      inputTokens: [
        {
          amount: amountIn.toString(),
          tokenAddress: this.handleNativeToken(tokenIn),
        },
      ],
      outputTokens: [
        {
          proportion: 1,
          tokenAddress: this.handleNativeToken(tokenOut),
        },
      ],
      userAddr: tokenZap,
      // slippage settings are applied when generating the zap data as minFwdAmount
      slippageLimitPercent: toPercentFloat(SlippageMax),
      simple: true,
    }
    const response = await this.getQuoteResponse(request)
    if (!response) {
      return EmptyOdosQuote
    }
    const odosQuoteResponse: OdosQuoteResponse = await response.json()
    if (
      !odosQuoteResponse.outAmounts ||
      !odosQuoteResponse.pathId ||
      odosQuoteResponse.outAmounts.length !== 1
    ) {
      console.error({ request, response }, 'Odos: invalid quote response')
      return EmptyOdosQuote
    }
    const amountOut = odosQuoteResponse.outAmounts[0]
    if (amountOut === '0') {
      console.info({ request, response }, 'Odos: zero amount out')
      return EmptyOdosQuote
    }
    return {
      engineID: this.id,
      expectedAmountOut: BigNumber.from(amountOut),
      assembleRequest: {
        userAddr: tokenZap,
        pathId: odosQuoteResponse.pathId,
      },
    }
  }

  public async generateRoute(
    input: RouteInput,
    quote: OdosQuote
  ): Promise<SwapEngineRoute> {
    if (quote.engineID !== this.id || !quote.assembleRequest) {
      console.error({ quote }, 'Odos: unexpected quote')
      return getEmptyRoute(this.id)
    }
    const response = await this.getAssembleResponse(quote.assembleRequest)
    if (!response) {
      return getEmptyRoute(this.id)
    }
    const odosAssembleResponse: OdosAssembleResponse = await response.json()
    if (!odosAssembleResponse.transaction) {
      console.error(
        { request: quote.assembleRequest, response },
        'Odos: invalid assemble response'
      )
      return getEmptyRoute(this.id)
    }
    return generateAPIRoute(input, this.id, {
      amountOut: quote.expectedAmountOut,
      transaction: odosAssembleResponse.transaction,
    })
  }

  public async getAssembleResponse(
    params: OdosAssembleRequest
  ): Promise<Response | null> {
    return postWithTimeout(
      'Odos',
      `${ODOS_API_URL}/assemble`,
      ODOS_API_TIMEOUT,
      params
    )
  }

  public async getQuoteResponse(
    params: OdosQuoteRequest
  ): Promise<Response | null> {
    return postWithTimeout(
      'Odos',
      `${ODOS_API_URL}/quote/v2`,
      ODOS_API_TIMEOUT,
      params
    )
  }

  private handleNativeToken(token: string): string {
    return isNativeToken(token) ? AddressZero : token
  }
}
