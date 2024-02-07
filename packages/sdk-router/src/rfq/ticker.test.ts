import {
  ChainToken,
  Ticker,
  marshallChainToken,
  marshallTicker,
  unmarshallChainToken,
  unmarshallTicker,
} from './ticker'

describe('ticker operations', () => {
  const arbUSDC: ChainToken = {
    chainId: 42161,
    token: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
  }
  const arbUSDCStr = '42161:0xaf88d065e77c8cC2239327C5EDb3A432268e5831'

  const ethUSDC: ChainToken = {
    chainId: 1,
    token: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
  }
  const ethUSDCStr = '1:0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48'

  const noSeparator = arbUSDCStr.replace(':', '')
  const twoSeparators =
    arbUSDCStr + ':0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48'
  const invalidChainId = 'abc:0xaf88d065e77c8cC2239327C5EDb3A432268e5831'
  const invalidAddress = '42161:abc'

  describe('ChainToken', () => {
    it('marshalls a ChainToken', () => {
      const marshalled = marshallChainToken(arbUSDC)
      expect(marshalled).toEqual(arbUSDCStr)
    })

    it('unmarshalls a checksummed ChainToken', () => {
      const unmarshalled = unmarshallChainToken(arbUSDCStr)
      expect(unmarshalled).toEqual(arbUSDC)
    })

    it('unmarshalls a non-checksummed ChainToken', () => {
      const unmarshalled = unmarshallChainToken(arbUSDCStr.toLowerCase())
      expect(unmarshalled).toEqual(arbUSDC)
    })

    describe('Throws during unmarshalling', () => {
      it('No token separator is found', () => {
        expect(() => unmarshallChainToken(noSeparator)).toThrow(
          `Can not unmarshall "${noSeparator}": invalid format`
        )
      })

      it('More than one token separator is found', () => {
        expect(() => unmarshallChainToken(twoSeparators)).toThrow(
          `Can not unmarshall "${twoSeparators}": invalid format`
        )
      })

      it('Chain ID is not a number', () => {
        expect(() => unmarshallChainToken(invalidChainId)).toThrow(
          `Can not unmarshall "${invalidChainId}": abc is not a chain ID`
        )
      })

      it('Token is not a valid address', () => {
        expect(() => unmarshallChainToken(invalidAddress)).toThrow(
          'invalid address'
        )
      })
    })
  })

  describe('Ticker', () => {
    const ticker: Ticker = {
      originToken: arbUSDC,
      destToken: ethUSDC,
    }
    const tickerStr = `${arbUSDCStr}-${ethUSDCStr}`

    it('marshalls a Ticker', () => {
      const marshalled = marshallTicker(ticker)
      expect(marshalled).toEqual(tickerStr)
    })

    it('unmarshalls a Ticker', () => {
      const unmarshalled = unmarshallTicker(tickerStr)
      expect(unmarshalled).toEqual(ticker)
    })

    it('unmarshalls a Ticker with non-checksummed addresses', () => {
      const unmarshalled = unmarshallTicker(tickerStr.toLowerCase())
      expect(unmarshalled).toEqual(ticker)
    })

    describe('Throws during unmarshalling', () => {
      describe('Invalid ticker format', () => {
        const noTickerSeparator = tickerStr.replace('-', '')
        const twoTickerSeparators = tickerStr + `-10:${ethUSDCStr}`

        it('No ticker separator is found', () => {
          expect(() => unmarshallTicker(noTickerSeparator)).toThrow(
            `Can not unmarshall "${noTickerSeparator}": invalid format`
          )
        })

        it('More than one ticker separator is found', () => {
          expect(() => unmarshallTicker(twoTickerSeparators)).toThrow(
            `Can not unmarshall "${twoTickerSeparators}": invalid format`
          )
        })
      })

      describe('Invalid origin token', () => {
        it('No origin token separator is found', () => {
          expect(() =>
            unmarshallTicker(`${noSeparator}-${ethUSDCStr}`)
          ).toThrow(`Can not unmarshall "${noSeparator}": invalid format`)
        })

        it('More than one origin token separator is found', () => {
          expect(() =>
            unmarshallTicker(`${twoSeparators}-${ethUSDCStr}`)
          ).toThrow(`Can not unmarshall "${twoSeparators}": invalid format`)
        })

        it('Origin chainId is not a number', () => {
          expect(() =>
            unmarshallTicker(`${invalidChainId}-${ethUSDCStr}`)
          ).toThrow(
            `Can not unmarshall "${invalidChainId}": abc is not a chain ID`
          )
        })

        it('Origin token is not a valid address', () => {
          expect(() =>
            unmarshallTicker(`${invalidAddress}-${ethUSDCStr}`)
          ).toThrow('invalid address')
        })
      })

      describe('Invalid destination token', () => {
        it('No destination token separator is found', () => {
          expect(() =>
            unmarshallTicker(`${arbUSDCStr}-${noSeparator}`)
          ).toThrow(`Can not unmarshall "${noSeparator}": invalid format`)
        })

        it('More than one destination token separator is found', () => {
          expect(() =>
            unmarshallTicker(`${arbUSDCStr}-${twoSeparators}`)
          ).toThrow(`Can not unmarshall "${twoSeparators}": invalid format`)
        })

        it('Destination chainId is not a number', () => {
          expect(() =>
            unmarshallTicker(`${arbUSDCStr}-${invalidChainId}`)
          ).toThrow(
            `Can not unmarshall "${invalidChainId}": abc is not a chain ID`
          )
        })

        it('Destination token is not a valid address', () => {
          expect(() =>
            unmarshallTicker(`${arbUSDCStr}-${invalidAddress}`)
          ).toThrow('invalid address')
        })
      })
    })
  })
})
