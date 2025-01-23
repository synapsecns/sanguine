import { ARB_USDC, ARB_USDT } from '../../constants/testValues'
import { USER_SIMULATED_ADDRESS } from './swapEngine'
import { LiFiEngine } from './liFiEngine'
import { SupportedChainId } from '../../constants'

global.fetch = require('node-fetch')

describe('Integration test: LiFiEngine', () => {
  it('Arbitrum USDC -> USDT /quote', async () => {
    const liFiEngine = new LiFiEngine()
    const response = await liFiEngine.getQuoteResponse(
      {
        fromChain: SupportedChainId.ARBITRUM,
        toChain: SupportedChainId.ARBITRUM,
        fromToken: ARB_USDC,
        toToken: ARB_USDT,
        fromAddress: USER_SIMULATED_ADDRESS,
        fromAmount: '1000000',
        slippage: 0.01,
        skipSimulation: true,
        // swapStepTimingStrategies: 'minWaitTime-300-1-100',
        // routeStepTimingStrategies: 'minWaitTime-100-0-0',
      },
      2000
    )
    expect(response).not.toBeNull()
    if (!response) {
      return
    }
    const quoteResponse = await response.json()
    console.log(JSON.stringify(quoteResponse, null, 2))
  })
})
