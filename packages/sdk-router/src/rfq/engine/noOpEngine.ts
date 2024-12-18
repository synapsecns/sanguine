import { BigNumber } from 'ethers'

import { isSameAddress } from '../../utils/addressUtils'
import { SwapEngine, SwapEngineRoute, EmptyRoute, EngineID } from './swapEngine'
import { BigintIsh } from '../../constants'

export class NoOpEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.NoOp

  public async findRoute(
    _chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<SwapEngineRoute> {
    if (!isSameAddress(tokenIn, tokenOut)) {
      return EmptyRoute
    }
    return {
      engineID: this.id,
      expectedAmountOut: BigNumber.from(amountIn),
      minAmountOut: BigNumber.from(amountIn),
      steps: [],
    }
  }

  public applySlippage(
    _chainId: number,
    route: SwapEngineRoute
  ): SwapEngineRoute {
    // Slippage settings are ignored for NoOpEngine
    return route
  }
}
