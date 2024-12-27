import { AddressZero, Zero } from '@ethersproject/constants'
import invariant from 'tiny-invariant'

import { BigintIsh, TOKEN_ZAP_V1_ADDRESS_MAP } from '../../constants'
import { ChainProvider } from '../../router'
import { DefaultEngine } from './defaultEngine'
import { NoOpEngine } from './noOpEngine'
import {
  SwapEngine,
  SwapEngineRoute,
  Recipient,
  RecipientEntity,
  USER_SIMULATED_ADDRESS,
  RouteInput,
  SwapEngineQuote,
  sanitizeMultiStepQuote,
} from './swapEngine'
import { CCTPRouterQuery } from '../../module'
import { encodeStepParams } from '../steps'
import { OdosEngine } from './odosEngine'

export type TokenInput = {
  address: string
  amount: BigintIsh
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
    this._addEngine(new OdosEngine(TOKEN_ZAP_V1_ADDRESS_MAP))

    this.tokenZaps = {}
    chains.forEach(({ chainId }) => {
      const tokenZapAddress = TOKEN_ZAP_V1_ADDRESS_MAP[chainId]
      // Skip chains without a Token Zap address
      if (tokenZapAddress) {
        this.tokenZaps[chainId] = tokenZapAddress
      }
    })
  }

  public async getOriginQuotes(
    chainId: number,
    tokenIn: TokenInput,
    tokensOut: string[]
  ): Promise<SwapEngineQuote[]> {
    const finalRecipient: Recipient = {
      entity: RecipientEntity.Self,
      address: this.getTokenZap(chainId),
    }
    // Find the quote for each token and each engine.
    const allQuotes = await Promise.all(
      tokensOut.map(async (tokenOut) =>
        Promise.all(
          Object.values(this.engines).map(async (engine) =>
            engine.getQuote({
              chainId,
              tokenIn: tokenIn.address,
              tokenOut,
              amountIn: tokenIn.amount,
              finalRecipient,
            })
          )
        )
      )
    )
    // Select the best quote for each tokenOut.
    return this._selectBestQuotes(allQuotes)
  }

  public async getDestinationQuotes(
    chainId: number,
    tokensIn: TokenInput[],
    tokenOut: string
  ): Promise<SwapEngineQuote[]> {
    // Check that the chain is supported
    this.getTokenZap(chainId)
    const finalRecipient: Recipient = {
      entity: RecipientEntity.UserSimulated,
      address: USER_SIMULATED_ADDRESS,
    }
    // Find the quote for each token and each engine.
    // Remove the quotes that have more than one Zap step (if populated).
    const allQuotes = await Promise.all(
      tokensIn.map(async (tokenIn) =>
        Promise.all(
          Object.values(this.engines).map(async (engine) => {
            const quote = await engine.getQuote({
              chainId,
              tokenIn: tokenIn.address,
              tokenOut,
              amountIn: tokenIn.amount,
              finalRecipient,
            })
            return sanitizeMultiStepQuote(quote)
          })
        )
      )
    )
    // Select the best quote for each tokenIn.
    return this._selectBestQuotes(allQuotes)
  }

  public async getQuote(
    engineID: number,
    chainId: number,
    tokenIn: TokenInput,
    tokenOut: string,
    finalRecipient: Recipient
  ): Promise<SwapEngineQuote> {
    return this._getEngine(engineID).getQuote({
      chainId,
      tokenIn: tokenIn.address,
      tokenOut,
      amountIn: tokenIn.amount,
      finalRecipient,
    })
  }

  public async generateRoute(
    input: RouteInput,
    quote: SwapEngineQuote
  ): Promise<SwapEngineRoute> {
    return this._getEngine(quote.engineID).generateRoute(input, quote)
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

  private _selectBestQuotes(quotes: SwapEngineQuote[][]): SwapEngineQuote[] {
    return quotes.map((quote) =>
      quote.reduce((best, current) =>
        current.expectedAmountOut.gt(best.expectedAmountOut) ? current : best
      )
    )
  }
}
