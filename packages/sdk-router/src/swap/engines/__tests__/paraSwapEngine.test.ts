import { SupportedChainId } from '../../../constants'
import { ETH_USDC } from '../../../constants/testValues'
import { ETH_NATIVE_TOKEN_ADDRESS } from '../../../utils'
import { USER_SIMULATED_ADDRESS } from '../../core'
import { ParaSwapEngine, ParaSwapPricesResponse } from '../paraSwapEngine'

global.fetch = require('node-fetch')

const TEST_TIMEOUT = 5000

// Unskip to check if integration is working
describe.skip('Integration test: ParaSwapEngine', () => {
  it('Ethereum USDC -> ETH', async () => {
    const paraSwapEngine = new ParaSwapEngine([])
    const amount = '1234567890'
    let response = await paraSwapEngine.getPricesResponse(
      {
        srcToken: ETH_USDC,
        srcDecimals: 6,
        destToken: ETH_NATIVE_TOKEN_ADDRESS,
        destDecimals: 18,
        amount,
        side: 'SELL',
        network: SupportedChainId.ETH,
        excludeRFQ: true,
        userAddress: USER_SIMULATED_ADDRESS,
        version: '6.2',
      },
      TEST_TIMEOUT
    )
    expect(response).not.toBeNull()
    if (!response) {
      return
    }
    const pricesResponse: ParaSwapPricesResponse = await response.json()
    console.log(JSON.stringify(pricesResponse, null, 2))

    response = await paraSwapEngine.getTransactionsResponse(
      SupportedChainId.ETH,
      {
        srcToken: ETH_USDC,
        srcDecimals: 6,
        destToken: ETH_NATIVE_TOKEN_ADDRESS,
        destDecimals: 18,
        srcAmount: amount,
        userAddress: USER_SIMULATED_ADDRESS,
        slippage: 100,
        priceRoute: pricesResponse.priceRoute,
      },
      TEST_TIMEOUT
    )
    expect(response).not.toBeNull()
    if (!response) {
      return
    }
    const transactionsResponse = await response.json()
    console.log(JSON.stringify(transactionsResponse, null, 2))
  })
})
