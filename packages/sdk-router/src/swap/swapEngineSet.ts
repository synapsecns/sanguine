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
  getEnginePriority,
  Priority,
  RouteInput,
  sanitizeMultiStepQuote,
  sanitizeMultiStepRoute,
  setMinFinalAmount,
  SwapEngine,
  SwapEngineQuote,
  SwapEngineRoute,
} from './models'
import { logExecutionTime, Prettify, TokenMetadataFetcher } from '../utils'

// We don't wait for all quotes to be resolved, so we use a longer common timeout.
const ENGINE_TIMEOUT = 2000

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

  @logExecutionTime('SwapEngineSet.getBestQuote')
  public async getBestQuote(
    input: RouteInput,
    options: QuoteOptions
  ): Promise<SwapEngineQuote | undefined> {
    const enginePromises = Object.values(this.engines).map((engine) => ({
      engine,
      quotePromise: this._getQuote(engine, input, options),
    }))
    // Wait for the first non-Zero quote to resolve from engines with the highest priority (Normal).
    const fastQuotePromise = await this._getFastestQuote(
      enginePromises
        .filter(
          ({ engine }) =>
            getEnginePriority(engine.id, input.chainId) === Priority.Normal
        )
        .map(({ quotePromise }) => quotePromise)
    )
    if (fastQuotePromise) {
      return fastQuotePromise
    }
    // Use the best quote from all the engines as a fallback.
    const allQuotes = await Promise.all(
      enginePromises.map(({ quotePromise }) => quotePromise)
    )
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

  @logExecutionTime('SwapEngineSet.generateRoute')
  public async generateRoute(
    input: RouteInput,
    quote: SwapEngineQuote,
    options: RouteOptions
  ): Promise<SwapEngineRoute | undefined> {
    let route = await this._getEngine(quote.engineID).generateRoute(
      input,
      quote,
      options.timeout ?? ENGINE_TIMEOUT
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
    const quote = await engine.getQuote(
      input,
      options.timeout ?? ENGINE_TIMEOUT
    )
    return options.allowMultiStep ? quote : sanitizeMultiStepQuote(quote)
  }

  /**
   * Applies the engine filter to the engine promises and returns the first non-Zero quote.
   * Returns undefined if no non-Zero quote is found.
   */
  private async _getFastestQuote(
    quotePromises: Promise<SwapEngineQuote>[]
  ): Promise<SwapEngineQuote | undefined> {
    try {
      return Promise.any(
        quotePromises.map(async (quotePromise) => {
          const quote = await quotePromise
          if (quote.expectedToAmount.gt(Zero)) {
            return quote
          }
          throw new Error('Zero Quote')
        })
      )
    } catch (e) {
      // Promise.any throws when all promises reject, so none of the quotes are non-Zero.
      return undefined
    }
  }
}
