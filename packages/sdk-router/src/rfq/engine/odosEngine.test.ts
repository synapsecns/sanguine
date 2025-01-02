import { ETH_USDC, ETH_USDT } from '../../constants/testValues'
import { USER_SIMULATED_ADDRESS } from './swapEngine'
import { OdosEngine, OdosQuoteResponse } from './odosEngine'

global.fetch = require('node-fetch')

describe('Integration test: ParaSwapEngine', () => {
  it('Ethereum USDC -> USDT', async () => {
    const odosEngine = new OdosEngine({})
    let response = await odosEngine.getQuoteResponse(
      {
        chainId: 1,
        inputTokens: [{ amount: '1000000000', tokenAddress: ETH_USDC }],
        outputTokens: [{ proportion: 1, tokenAddress: ETH_USDT }],
        userAddr: USER_SIMULATED_ADDRESS,
        slippageLimitPercent: 1,
        simple: true,
      },
      2000
    )
    expect(response).not.toBeNull()
    const quoteResponse: OdosQuoteResponse = await response?.json()
    console.log(JSON.stringify(quoteResponse, null, 2))

    response = await odosEngine.getAssembleResponse(
      {
        userAddr: USER_SIMULATED_ADDRESS,
        pathId: quoteResponse.pathId,
      },
      2000
    )
    expect(response).not.toBeNull()
    const assembleResponse = await response?.json()
    console.log(JSON.stringify(assembleResponse, null, 2))
  })
})
