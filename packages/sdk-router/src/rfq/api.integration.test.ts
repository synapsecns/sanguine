import { parseFixed } from '@ethersproject/bignumber'

import { getAllQuotes, getBestRelayerQuote, RelayerQuote } from './api'
import { Ticker } from './ticker'
import { ETH_NATIVE_TOKEN_ADDRESS } from '../utils/handleNativeToken'

global.fetch = require('node-fetch')

// Retry the flaky tests up to 3 times
jest.retryTimes(3)

describe('Integration test: getAllQuotes', () => {
  it('returns a non-empty array', async () => {
    const result = await getAllQuotes()
    // console.log('Current quotes: ' + JSON.stringify(result, null, 2))
    expect(result.length).toBeGreaterThan(0)
  })
})

describe('Integration test: getBestRelayerQuote', () => {
  const ticker: Ticker = {
    originToken: {
      chainId: 42161,
      token: ETH_NATIVE_TOKEN_ADDRESS,
    },
    destToken: {
      chainId: 10,
      token: ETH_NATIVE_TOKEN_ADDRESS,
    },
  }
  const userAddress = '0x0000000000000000000000000000000000007331'

  describe('Cases where a non-zero quote is returned', () => {
    it('ARB ETH -> OP ETH; 0.01 ETH', async () => {
      const result = await getBestRelayerQuote(
        ticker,
        parseFixed('0.01', 18),
        userAddress
      )
      expect(result?.destAmount.gt(0)).toBe(true)
      expect(result?.relayerAddress).toBeDefined()
    })
  })

  describe('Cases where a zero quote is returned', () => {
    const quoteZero: RelayerQuote = {
      destAmount: parseFixed('0'),
    }

    beforeEach(() => {
      jest.spyOn(console, 'error').mockImplementation(() => {
        // Do nothing
      })
    })

    afterEach(() => {
      jest.restoreAllMocks()
    })

    it('ARB ETH -> OP ETH; 1337 wei', async () => {
      const result = await getBestRelayerQuote(
        ticker,
        parseFixed('1337'),
        userAddress
      )
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('ARB ETH -> OP ETH; 10**36 wei', async () => {
      const result = await getBestRelayerQuote(
        ticker,
        parseFixed('1', 36),
        userAddress
      )
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })
  })
})
