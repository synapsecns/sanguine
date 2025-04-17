import { SupportedChainId } from '../../../constants'
import { ARB_USDC } from '../../../constants/testValues'
import { ETH_NATIVE_TOKEN_ADDRESS } from '../../../utils'
import { USER_SIMULATED_ADDRESS } from '../../core'
import { LiFiEngine } from '../liFiEngine'

global.fetch = require('node-fetch')

const TEST_TIMEOUT = 5000

// Unskip to check if integration is working
describe.skip('Integration test: LiFiEngine', () => {
  it('Arbitrum USDC -> ETH /quote', async () => {
    const liFiEngine = new LiFiEngine()
    const response = await liFiEngine.getQuoteResponse(
      {
        fromChain: SupportedChainId.ARBITRUM,
        toChain: SupportedChainId.ARBITRUM,
        fromToken: ARB_USDC,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
        fromAddress: USER_SIMULATED_ADDRESS,
        fromAmount: '1000000',
        slippage: 0.01,
        skipSimulation: true,
        // swapStepTimingStrategies: 'minWaitTime-300-1-100',
        // routeStepTimingStrategies: 'minWaitTime-100-0-0',
        name: 'cortex_protocol',
      },
      TEST_TIMEOUT
    )
    expect(response).not.toBeNull()
    if (!response) {
      return
    }
    const quoteResponse = await response.json()
    console.log(JSON.stringify(quoteResponse, null, 2))
  })
})
