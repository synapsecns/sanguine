import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { isSameAddress, logger } from '../../utils'
import { EngineID } from '../core'
import {
  getEmptyQuote,
  getEmptyRoute,
  RouteInput,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from '../models'
import { SupportedChainId } from '../../constants'

export class LiquidSwapEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.LiquidSwap

  public async getQuote(input: RouteInput): Promise<SwapEngineQuote> {
    const { chainId, fromToken, toToken, fromAmount } = input
    if (
      chainId !== SupportedChainId.HYPEREVM ||
      isSameAddress(fromToken, toToken) ||
      BigNumber.from(fromAmount).eq(Zero)
    ) {
      return getEmptyQuote(this.id)
    }
    // TODO: implement quoting logic
    return getEmptyQuote(this.id)
  }

  public async generateRoute(
    input: RouteInput,
    quote: SwapEngineQuote
  ): Promise<SwapEngineRoute> {
    if (quote.engineID !== this.id) {
      logger.error({ quote }, 'LiquidSwapEngine: unexpected quote')
      return getEmptyRoute(this.id)
    }
    // TODO: implement route generation logic
    return getEmptyRoute(this.id)
  }
}
