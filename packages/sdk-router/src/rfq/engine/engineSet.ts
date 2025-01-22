import { hexlify } from '@ethersproject/bytes'
import { AddressZero, Zero } from '@ethersproject/constants'
import invariant from 'tiny-invariant'

import { TOKEN_ZAP_V1_ADDRESS_MAP } from '../../constants'
import { ChainProvider } from '../../router'
import { DefaultEngine } from './defaultEngine'
import { NoOpEngine } from './noOpEngine'
import {
  SwapEngine,
  SwapEngineRoute,
  RouteInput,
  SwapEngineQuote,
  sanitizeMultiStepQuote,
  sanitizeMultiStepRoute,
} from './swapEngine'
import { compareQuotesWithPriority } from './priority'
import { CCTPRouterQuery } from '../../module'
import { encodeStepParams } from '../steps'
import { KyberSwapEngine } from './kyberSwapEngine'
import { decodeZapData, encodeZapData } from '../zapData'

export enum EngineTimeout {
  Short = 1000,
  Long = 3000,
}

type QuoteOptions = {
  allowMultiStep: boolean
  timeout?: number
}

type RouteOptions = QuoteOptions & {
  useZeroSlippage: boolean
}

export class EngineSet {
  private engines: {
    [engineID: number]: SwapEngine
  }

  private tokenZaps: {
    [chainId: number]: string
  }

  constructor(chains: ChainProvider[]) {
    this.engines = {}
    this._addEngine(new NoOpEngine())
    this._addEngine(new DefaultEngine(chains))
    this._addEngine(new KyberSwapEngine(TOKEN_ZAP_V1_ADDRESS_MAP))

    this.tokenZaps = {}
    chains.forEach(({ chainId }) => {
      const tokenZapAddress = TOKEN_ZAP_V1_ADDRESS_MAP[chainId]
      // Skip chains without a Token Zap address
      if (tokenZapAddress) {
        this.tokenZaps[chainId] = tokenZapAddress
      }
    })
  }

  public async getBestQuote(
    input: RouteInput,
    options: QuoteOptions
  ): Promise<SwapEngineQuote | undefined> {
    // Find the quote for each engine.
    const allQuotes = await Promise.all(
      Object.values(this.engines).map(async (engine) =>
        this._getQuote(engine, input, options)
      )
    )
    // Select the best quote.
    const quote = allQuotes.reduce(compareQuotesWithPriority)
    return quote.expectedAmountOut.gt(Zero) ? quote : undefined
  }

  public async getQuote(
    engineID: number,
    input: RouteInput,
    options: QuoteOptions
  ): Promise<SwapEngineQuote | undefined> {
    const quote = await this._getQuote(
      this._getEngine(engineID),
      input,
      options
    )
    return quote.expectedAmountOut.gt(Zero) ? quote : undefined
  }

  public async generateRoute(
    input: RouteInput,
    quote: SwapEngineQuote,
    options: RouteOptions
  ): Promise<SwapEngineRoute | undefined> {
    // Use longer timeout for route generation by default.
    let route = await this._getEngine(quote.engineID).generateRoute(
      input,
      quote,
      options.timeout ?? EngineTimeout.Long
    )
    route = options.allowMultiStep ? route : sanitizeMultiStepRoute(route)
    if (options.useZeroSlippage && route.steps.length > 0) {
      const lastStepIndex = route.steps.length - 1
      const lastStepZapData = decodeZapData(
        hexlify(route.steps[lastStepIndex].zapData)
      )
      route.steps[lastStepIndex].zapData = encodeZapData({
        ...lastStepZapData,
        minFinalAmount: route.expectedAmountOut,
      })
    }
    return route.expectedAmountOut.gt(Zero) ? route : undefined
  }

  public getOriginQuery(
    chainId: number,
    route: SwapEngineRoute,
    tokenOut: string
  ): CCTPRouterQuery {
    // To preserve consistency with other modules, router adapter is not set for a no-op intent
    return {
      routerAdapter:
        route.steps.length > 0 ? this.getTokenZap(chainId) : AddressZero,
      tokenOut,
      minAmountOut: route.expectedAmountOut,
      // The default deadline will be overridden later in `finalizeBridgeRoute`
      deadline: Zero,
      rawParams: encodeStepParams(route.steps),
    }
  }

  public getTokenZap(chainId: number): string {
    const tokenZap = this.tokenZaps[chainId]
    if (!tokenZap) {
      throw new Error('Token Zap address not found for chain ' + chainId)
    }
    return tokenZap
  }

  private _addEngine(engine: SwapEngine) {
    invariant(!this.engines[engine.id], 'ENGINE_ALREADY_EXISTS')
    this.engines[engine.id] = engine
  }

  private _getEngine(engineID: number): SwapEngine {
    const engine = this.engines[engineID]
    if (!engine) {
      throw new Error('ENGINE_NOT_FOUND')
    }
    return engine
  }

  private async _getQuote(
    engine: SwapEngine,
    input: RouteInput,
    options: QuoteOptions
  ): Promise<SwapEngineQuote> {
    // Use shorter timeout for quote fetching by default.
    const quote = await engine.getQuote(
      input,
      options.timeout ?? EngineTimeout.Short
    )
    return options.allowMultiStep ? quote : sanitizeMultiStepQuote(quote)
  }
}
