import { ChainToken, marshallChainToken, unmarshallChainToken } from './ticker'

describe('ticker operations', () => {
  const arbUSDC: ChainToken = {
    chainId: 42161,
    token: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
  }
  const arbUSDCStr = '42161:0xaf88d065e77c8cC2239327C5EDb3A432268e5831'

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
})
