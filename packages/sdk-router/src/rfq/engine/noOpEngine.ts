import { BigNumber } from 'ethers'

import { isSameAddress } from '../../utils/addressUtils'
import { SwapEngine, SwapEngineRoute, EmptyRoute } from './swapEngine'
import { BigintIsh } from '../../constants'

export class NoOpEngine implements SwapEngine {
  public readonly id = 1

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
      id: this.id,
      expectedAmountOut: BigNumber.from(amountIn),
      minAmountOut: BigNumber.from(amountIn),
      steps: [],
    }
  }

  public modifyMinAmountOut(
    _chainId: number,
    route: SwapEngineRoute
  ): SwapEngineRoute {
    // Slippage settings are ignored for NoOpEngine
    return route
  }

  public modifyRecipient(
    _chainId: number,
    route: SwapEngineRoute
  ): SwapEngineRoute {
    // Recipient settings are ignored for NoOpEngine
    return route
  }
}
