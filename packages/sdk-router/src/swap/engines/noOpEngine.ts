import { BigNumber } from 'ethers'

import { isSameAddress, logger } from '../../utils'
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
    const { chainId, fromToken, toToken, fromAmount } = input
    if (!isSameAddress(fromToken, toToken)) {
      return getEmptyRoute(this.id)
    }
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      fromToken,
      toToken,
      fromAmount: BigNumber.from(fromAmount),
      expectedToAmount: BigNumber.from(fromAmount),
      minToAmount: BigNumber.from(fromAmount),
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
