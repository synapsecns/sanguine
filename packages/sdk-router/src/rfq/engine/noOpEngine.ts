import { BigNumber } from 'ethers'

import { isSameAddress } from '../../utils/addressUtils'
import {
  SwapEngine,
  SwapEngineRoute,
  EmptyRoute,
  EngineID,
  RouteInput,
} from './swapEngine'

export class NoOpEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.NoOp

  public async findRoute(input: RouteInput): Promise<SwapEngineRoute> {
    return this._findRoute(input)
  }

  public async findRoutes(inputs: RouteInput[]): Promise<SwapEngineRoute[]> {
    return inputs.map((input) => this._findRoute(input))
  }

  private _findRoute(input: RouteInput): SwapEngineRoute {
    const { tokenIn, tokenOut, amountIn } = input
    if (!isSameAddress(tokenIn, tokenOut)) {
      return EmptyRoute
    }
    return {
      engineID: this.id,
      expectedAmountOut: BigNumber.from(amountIn),
      steps: [],
    }
  }
}
