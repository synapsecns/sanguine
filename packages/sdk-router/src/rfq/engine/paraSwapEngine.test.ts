import { ETH_USDC, ETH_USDT } from '../../constants/testValues'
import { USER_SIMULATED_ADDRESS } from './swapEngine'
import { ParaSwapEngine } from './paraSwapEngine'

global.fetch = require('node-fetch')

describe('Integration test: ParaSwapEngine', () => {
  it('returns a response', async () => {
    const paraSwapEngine = new ParaSwapEngine([], {})
    const response = await paraSwapEngine.getResponse({
      srcToken: ETH_USDC,
      srcDecimals: 6,
      destToken: ETH_USDT,
      destDecimals: 6,
      amount: '1000000000',
      side: 'SELL',
      userAddress: USER_SIMULATED_ADDRESS,
      network: '1',
      slippage: 0,
      version: '6.2',
    })
    expect(response).toBeDefined()
    console.log(JSON.stringify(response, null, 2))
  })
})
