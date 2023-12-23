import { BigNumber } from '@ethersproject/bignumber'

import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  marshallFastBridgeQuote,
  unmarshallFastBridgeQuote,
  applyQuote,
} from './quote'

describe('quote', () => {
  const quoteAPI: FastBridgeQuoteAPI = {
    OriginChainID: 1,
    OriginTokenAddr: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
    DestChainID: 2,
    DestTokenAddr: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    DestAmount: '4000000000000000000000',
    MaxOriginAmount: '3000000000000000000000',
    FixedFee: '1000000000000000000',
    RelayerAddr: '0xB300efF6B57AA09e5fCcf7221FCB9E676A74d931',
    UpdatedAt: '2023-01-02T03:04:05.678Z',
  }

  const quote: FastBridgeQuote = {
    ticker: {
      originToken: {
        chainId: 1,
        token: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
      },
      destToken: {
        chainId: 2,
        token: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      },
    },
    destAmount: BigNumber.from(10).pow(18).mul(4000),
    maxOriginAmount: BigNumber.from(10).pow(18).mul(3000),
    fixedFee: BigNumber.from(10).pow(18),
    relayerAddr: '0xB300efF6B57AA09e5fCcf7221FCB9E676A74d931',
    updatedAt: 1672628645678,
  }

  it('should unmarshall a quote', () => {
    expect(unmarshallFastBridgeQuote(quoteAPI)).toEqual(quote)
  })

  it('should marshall a quote', () => {
    expect(marshallFastBridgeQuote(quote)).toEqual(quoteAPI)
  })

  describe('applyQuote', () => {
    it('Returns zero if origin amount is lower than fixed fee', () => {
      expect(applyQuote(quote, quote.fixedFee.sub(1))).toEqual(
        BigNumber.from(0)
      )
    })

    it('Returns zero if origin amount is equal to fixed fee', () => {
      expect(applyQuote(quote, quote.fixedFee)).toEqual(BigNumber.from(0))
    })

    it('Returns zero if origin amount is greater than max origin amount', () => {
      expect(applyQuote(quote, quote.maxOriginAmount.add(1))).toEqual(
        BigNumber.from(0)
      )
    })

    it('Returns zero if resulted quote is higher than dest amount', () => {
      // Modify dest amount to be lower than the max origin amount
      quote.destAmount = BigNumber.from(10).pow(18).mul(100)
      expect(
        applyQuote(quote, quote.destAmount.add(quote.fixedFee).add(1))
      ).toEqual(BigNumber.from(0))
    })

    it('Returns a correct quote', () => {
      expect(applyQuote(quote, BigNumber.from(10).pow(18).mul(100))).toEqual(
        BigNumber.from(10).pow(18).mul(99)
      )
    })
  })
})
