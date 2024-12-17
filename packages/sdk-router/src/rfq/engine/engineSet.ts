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
  EmptyRoute,
  USER_SIMULATED_ADDRESS,
} from './swapEngine'
import { CCTPRouterQuery } from '../../module'
import { encodeStepParams } from '../steps'

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

    this.tokenZaps = {}
    chains.forEach(({ chainId }) => {
      const tokenZapAddress = TOKEN_ZAP_V1_ADDRESS_MAP[chainId]
      // Skip chains without a Token Zap address
      if (tokenZapAddress) {
        this.tokenZaps[chainId] = tokenZapAddress
      }
    })
  }

  public async getOriginRoutes(
    chainId: number,
    tokenIn: TokenInput,
    tokensOut: string[]
  ): Promise<SwapEngineRoute[]> {
    const recipient: Recipient = {
      entity: RecipientEntity.Self,
      address: this.getTokenZap(chainId),
    }
    // Find the route for each token and each engine.
    const allRoutes = await Promise.all(
      tokensOut.map(async (tokenOut) =>
        Promise.all(
          Object.values(this.engines).map(async (engine) =>
            engine.findRoute(
              chainId,
              tokenIn.address,
              tokenOut,
              tokenIn.amount,
              recipient
            )
          )
        )
      )
    )
    // Select the best response for each tokenOut.
    return this._selectBestRoutes(allRoutes)
  }

  public async getDestinationRoutes(
    chainId: number,
    tokensIn: TokenInput[],
    tokenOut: string
  ): Promise<SwapEngineRoute[]> {
    // Check that the chain is supported
    this.getTokenZap(chainId)
    const recipient: Recipient = {
      entity: RecipientEntity.UserSimulated,
      address: USER_SIMULATED_ADDRESS,
    }
    // Find the route for each token and each engine.
    // Remove the routes that have more than one Zap step.
    const allRoutes = await Promise.all(
      tokensIn.map(async (tokenIn) =>
        Promise.all(
          Object.values(this.engines).map(async (engine) => {
            const route = await engine.findRoute(
              chainId,
              tokenIn.address,
              tokenOut,
              tokenIn.amount,
              recipient
            )
            return this.limitSingleZap(route)
          })
        )
      )
    )
    // Select the best response for each tokenIn.
    return this._selectBestRoutes(allRoutes)
  }

  public async findRoute(
    engineID: number,
    chainId: number,
    tokenIn: TokenInput,
    tokenOut: string,
    finalRecipient: Recipient
  ): Promise<SwapEngineRoute> {
    return this._getEngine(engineID).findRoute(
      chainId,
      tokenIn.address,
      tokenOut,
      tokenIn.amount,
      finalRecipient
    )
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
      minAmountOut: route.minAmountOut,
      // The default deadline will be overridden later in `finalizeBridgeRoute`
      deadline: Zero,
      rawParams: encodeStepParams(route.steps),
    }
  }

  public limitSingleZap(route: SwapEngineRoute): SwapEngineRoute {
    return route.steps.length > 1 ? EmptyRoute : route
  }

  public modifyMinAmountOut(
    chainId: number,
    route: SwapEngineRoute,
    minAmountOut: BigintIsh
  ): SwapEngineRoute {
    return this._getEngine(route.engineID).modifyMinAmountOut(
      chainId,
      route,
      minAmountOut
    )
  }

  public modifyRecipient(
    chainId: number,
    route: SwapEngineRoute,
    finalRecipient: Recipient
  ): SwapEngineRoute {
    return this._getEngine(route.engineID).modifyRecipient(
      chainId,
      route,
      finalRecipient
    )
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

  private _selectBestRoutes(routes: SwapEngineRoute[][]): SwapEngineRoute[] {
    return routes.map((route) =>
      route.reduce((best, current) =>
        current.expectedAmountOut.gt(best.expectedAmountOut) ? current : best
      )
    )
  }
}
