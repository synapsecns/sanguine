import { ETH_USDC, ETH_USDT } from '../../constants/testValues'
import { USER_SIMULATED_ADDRESS } from './swapEngine'
import { OdosEngine } from './odosEngine'

global.fetch = require('node-fetch')

describe('Integration test: ParaSwapEngine', () => {
  it('Ethereum USDC -> USDT', async () => {
    const odosEngine = new OdosEngine({})
    const response = await odosEngine.getResponse({
      chainId: 1,
      inputTokens: [{ amount: '1000000000', tokenAddress: ETH_USDC }],
      outputTokens: [{ proportion: 1, tokenAddress: ETH_USDT }],
      userAddr: USER_SIMULATED_ADDRESS,
      slippageLimitPercent: 1,
      simple: true,
    })
    expect(response).toBeDefined()
    console.log(JSON.stringify(response, null, 2))
  })
})
