import { isSameAddress } from '../../utils/addressUtils'
import { EmptyRoute, EngineID, SwapEngine, SwapEngineRoute } from './swapEngine'

export class ParaSwapEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.ParaSwap

  public async findRoute(
    _chainId: number,
    tokenIn: string,
    tokenOut: string
  ): Promise<SwapEngineRoute> {
    if (isSameAddress(tokenIn, tokenOut)) {
      return EmptyRoute
    }
    // TODO: implement
    return EmptyRoute
  }

  public applySlippage(
    _chainId: number,
    route: SwapEngineRoute
  ): SwapEngineRoute {
    // TODO: implement
    return route
  }
}
