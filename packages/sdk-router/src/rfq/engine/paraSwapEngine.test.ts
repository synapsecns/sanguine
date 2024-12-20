import { ETH_USDC, ETH_USDT } from '../../constants/testValues'
import { USER_SIMULATED_ADDRESS } from './swapEngine'
import { ParaSwapEngine } from './paraSwapEngine'

global.fetch = require('node-fetch')

describe('Integration test: ParaSwapEngine', () => {
  it('Ethereum USDC -> USDT', async () => {
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

  it('Arbitrum FRAX -> ETH', async () => {
    const paraSwapEngine = new ParaSwapEngine([], {})
    const response = await paraSwapEngine.getResponse({
      srcToken: '0x17FC002b466eEc40DaE837Fc4bE5c67993ddBd6F',
      srcDecimals: 18,
      destToken: '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE',
      destDecimals: 18,
      amount: '250160206404537300',
      side: 'SELL',
      userAddress: '0x289db76b9E19487190D356ecB64324A5c716fFe1',
      network: '42161',
      slippage: 9999,
      version: '6.2',
    })
    console.log(JSON.stringify(response, null, 2))
  })
})
