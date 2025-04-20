import { Zero } from '@ethersproject/constants'
import invariant from 'tiny-invariant'

import { TOKEN_ZAP_V1_ADDRESS_MAP } from '../constants'
import { ChainProvider } from '../router'
import { applySlippage, Slippage, SlippageZero } from './core'
import {
  DefaultPoolsEngine,
  KyberSwapEngine,
  LiFiEngine,
  NoOpEngine,
  ParaSwapEngine,
} from './engines'
import {
  compareQuotesWithPriority,
  RouteInput,
  sanitizeMultiStepQuote,
  sanitizeMultiStepRoute,
  setMinFinalAmount,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from './models'
import { Prettify, TokenMetadataFetcher } from '../utils'

export enum EngineTimeout {
  Short = 1000,
  Long = 3000,
}

type QuoteOptions = {
  allowMultiStep: boolean
  timeout?: number
}

type RouteOptions = Prettify<
  QuoteOptions & {
    slippage?: Slippage
  }
>

export class SwapEngineSet {
  private engines: {
    [engineID: number]: SwapEngine
  }

  private tokenZaps: {
    [chainId: number]: string
  }

  constructor(
    chains: ChainProvider[],
    tokenMetadataFetcher?: TokenMetadataFetcher
  ) {
    this.engines = {}
    this._addEngine(new NoOpEngine())
    this._addEngine(new DefaultPoolsEngine(chains))
    this._addEngine(new KyberSwapEngine())
    this._addEngine(new ParaSwapEngine(chains, tokenMetadataFetcher))
    this._addEngine(new LiFiEngine())

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
    return quote.expectedToAmount.gt(Zero) ? quote : undefined
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
    return quote.expectedToAmount.gt(Zero) ? quote : undefined
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
    if (route.steps.length > 0) {
      route.minToAmount = applySlippage(
        quote.expectedToAmount,
        options.slippage ?? SlippageZero
      )
      route.steps = setMinFinalAmount(route.steps, route.minToAmount)
    }
    return route.expectedToAmount.gt(Zero) ? route : undefined
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
