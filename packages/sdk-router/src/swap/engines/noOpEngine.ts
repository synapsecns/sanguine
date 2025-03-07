import { BigNumber } from 'ethers'

import { isSameAddress } from '../../utils/addressUtils'
import { logger } from '../../utils/logger'
import { EngineID } from '../core'
import {
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineRoute,
} from '../models'

export class NoOpEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.NoOp

  public async getQuote(input: RouteInput): Promise<SwapEngineRoute> {
    const { chainId, tokenIn, tokenOut, amountIn } = input
    if (!isSameAddress(tokenIn, tokenOut)) {
      return getEmptyRoute(this.id)
    }
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      expectedAmountOut: BigNumber.from(amountIn),
      steps: [],
    }
  }

  public async generateRoute(
    _input: RouteInput,
    quote: SwapEngineRoute
  ): Promise<SwapEngineRoute> {
    if (quote.engineID !== this.id || !quote.steps) {
      logger.error({ quote }, 'NoOpEngine: unexpected quote')
      return getEmptyRoute(this.id)
    }
    return quote
  }
}
