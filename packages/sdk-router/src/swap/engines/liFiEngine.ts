import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { isSameAddress } from '../../utils/addressUtils'
import { getWithTimeout } from '../../utils/api'
import { logExecutionTime, logger } from '../../utils/logger'
import { EngineID, SlippageMax, toFloat } from '../core'
import {
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
  routeStepTimingStrategies?: string
  name?: string
}

export type LiFiQuoteResponse = {
  estimate: {
    toAmount: string
  }
  transactionRequest: TransactionData
}

type LiFiQuote = SwapEngineQuote & {
  tx?: TransactionData
}

const EmptyLiFiQuote: LiFiQuote = {
  engineID: EngineID.LiFi,
  engineName: EngineID[EngineID.LiFi],
  chainId: 0,
  expectedAmountOut: Zero,
}

export class LiFiEngine implements SwapEngine {
  readonly id: EngineID = EngineID.LiFi

  public async getQuote(
    input: RouteInput,
    timeout: number
  ): Promise<LiFiQuote> {
    const { chainId, tokenIn, tokenOut, msgSender, amountIn } = input
    if (isSameAddress(tokenIn, tokenOut) || BigNumber.from(amountIn).eq(Zero)) {
      return EmptyLiFiQuote
    }
    const response = await this.getQuoteResponse(
      {
        fromChain: chainId,
        toChain: chainId,
        fromToken: tokenIn,
        toToken: tokenOut,
        fromAddress: msgSender,
        fromAmount: amountIn.toString(),
        slippage: toFloat(SlippageMax),
        skipSimulation: true,
        // TODO: figure out optimal values
        swapStepTimingStrategies: 'minWaitTime-0-5-100',
        routeStepTimingStrategies: 'minWaitTime-0-5-100',
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
      expectedAmountOut: BigNumber.from(liFiResponse.estimate.toAmount),
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
      amountOut: quote.expectedAmountOut,
      transaction: quote.tx,
    })
  }

  @logExecutionTime('LiFiEngine.getQuoteResponse')
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
