import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { isSameAddress, getWithTimeout, logger, Prettify } from '../../utils'
import { EngineID, SlippageMax, toFloat } from '../core'
import {
  getEmptyQuote,
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from '../models'
import { generateAPIRoute, TransactionData } from './response'

const LIFI_API_URL = 'https://li.quest/v1'

export type LiFiQuoteRequest = {
  fromChain: number
  toChain: number
  fromToken: string
  toToken: string
  fromAddress: string
  fromAmount: string
  slippage: number
  skipSimulation: boolean
  swapStepTimingStrategies?: string
  routeTimingStrategies?: string
  name?: string
}

export type LiFiQuoteResponse = {
  estimate: {
    toAmount: string
  }
  transactionRequest: TransactionData
}

type LiFiQuote = Prettify<
  SwapEngineQuote & {
    tx?: TransactionData
  }
>

const EmptyLiFiQuote: LiFiQuote = getEmptyQuote(EngineID.LiFi)

// TODO: figure out optimal values
export const LI_FI_STRATEGY = `minWaitTime-0-3-300`

export class LiFiEngine implements SwapEngine {
  readonly id: EngineID = EngineID.LiFi

  public async getQuote(
    input: RouteInput,
    timeout: number
  ): Promise<LiFiQuote> {
    const { chainId, fromToken, toToken, swapper, fromAmount } = input
    if (
      isSameAddress(fromToken, toToken) ||
      BigNumber.from(fromAmount).eq(Zero)
    ) {
      return EmptyLiFiQuote
    }
    const response = await this.getQuoteResponse(
      {
        fromChain: chainId,
        toChain: chainId,
        fromToken,
        toToken,
        fromAddress: swapper,
        fromAmount: fromAmount.toString(),
        slippage: toFloat(SlippageMax),
        skipSimulation: true,
        routeTimingStrategies: LI_FI_STRATEGY,
        name: 'cortex_protocol',
      },
      timeout
    )
    if (!response) {
      return EmptyLiFiQuote
    }
    const liFiResponse: LiFiQuoteResponse = await response.json()
    if (!liFiResponse.estimate?.toAmount || !liFiResponse.transactionRequest) {
      return EmptyLiFiQuote
    }
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      fromToken,
      toToken,
      fromAmount: BigNumber.from(fromAmount),
      expectedToAmount: BigNumber.from(liFiResponse.estimate.toAmount),
      tx: liFiResponse.transactionRequest,
    }
  }

  public async generateRoute(
    input: RouteInput,
    quote: LiFiQuote
  ): Promise<SwapEngineRoute> {
    if (quote.engineID !== this.id || !quote.tx) {
      logger.error({ quote }, 'LiFiEngine: unexpected quote')
      return getEmptyRoute(this.id)
    }
    return generateAPIRoute(input, this.id, SlippageMax, {
      expectedToAmount: quote.expectedToAmount,
      transaction: quote.tx,
    })
  }

  public async getQuoteResponse(
    params: LiFiQuoteRequest,
    timeout: number
  ): Promise<Response | null> {
    const headers = process.env.LIFI_API_KEY && {
      'x-lifi-api-key': process.env.LIFI_API_KEY,
    }
    return getWithTimeout(
      'LiFi',
      `${LIFI_API_URL}/quote`,
      timeout,
      params,
      headers
    )
  }
}
