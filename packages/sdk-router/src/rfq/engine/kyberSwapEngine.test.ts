import { ETH_USDC, ETH_USDT } from '../../constants/testValues'
import { USER_SIMULATED_ADDRESS } from './swapEngine'
import {
  KyberSwapEngine,
  KyberSwapQuoteResponse,
  KyberSwapBuildResponse,
} from './kyberSwapEngine'
import { SupportedChainId } from '../../constants'

global.fetch = require('node-fetch')

describe('Integration test: KyberSwapEngine', () => {
  it('Ethereum USDC -> USDT', async () => {
    const kyberSwapEngine = new KyberSwapEngine()
    let response = await kyberSwapEngine.getQuoteResponse(
      SupportedChainId.ETH,
      {
        tokenIn: ETH_USDC,
        tokenOut: ETH_USDT,
        amountIn: '1000000000',
        gasInclude: true,
      },
      2000
    )
    expect(response).not.toBeNull()
    if (!response) {
      return
    }
    const quoteResponse: KyberSwapQuoteResponse = await response.json()
    console.log(JSON.stringify(quoteResponse, null, 2))

    response = await kyberSwapEngine.getBuildResponse(
      SupportedChainId.ETH,
      {
        routeSummary: quoteResponse.data.routeSummary,
        sender: USER_SIMULATED_ADDRESS,
        recipient: USER_SIMULATED_ADDRESS,
        deadline: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
        slippageTolerance: 100,
        enableGasEstimation: false,
      },
      2000
    )
    expect(response).not.toBeNull()
    if (!response) {
      return
    }
    const buildResponse: KyberSwapBuildResponse = await response.json()
    console.log(JSON.stringify(buildResponse, null, 2))
  })
})
