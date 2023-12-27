import { BigNumber, parseFixed } from '@ethersproject/bignumber'

import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  marshallFastBridgeQuote,
  unmarshallFastBridgeQuote,
  applyQuote,
} from './quote'

const createZeroAmountTests = (quote: FastBridgeQuote) => {
  describe('Returns zero', () => {
    it('If origin amount is zero', () => {
      expect(applyQuote(quote, BigNumber.from(0))).toEqual(BigNumber.from(0))
    })

    it('If origin amount is lower than fixed fee', () => {
      expect(applyQuote(quote, quote.fixedFee.sub(1))).toEqual(
        BigNumber.from(0)
      )
    })

    it('If origin amount is equal to fixed fee', () => {
      expect(applyQuote(quote, quote.fixedFee)).toEqual(BigNumber.from(0))
    })

    it('If origin amount is greater than max origin amount + fixed fee', () => {
      const amount = quote.maxOriginAmount.add(quote.fixedFee).add(1)
      expect(applyQuote(quote, amount)).toEqual(BigNumber.from(0))
    })
  })

  describe('Returns non-zero', () => {
    it('If origin amount is equal to max origin amount', () => {
      expect(applyQuote(quote, quote.maxOriginAmount)).not.toEqual(
        BigNumber.from(0)
      )
    })

    it('If origin amount is 1 wei greater than max origin amount', () => {
      const amount = quote.maxOriginAmount.add(1)
      expect(applyQuote(quote, amount)).not.toEqual(BigNumber.from(0))
    })

    it('If origin amount is max origin amount + fixed fee', () => {
      const amount = quote.maxOriginAmount.add(quote.fixedFee)
      expect(applyQuote(quote, amount)).not.toEqual(BigNumber.from(0))
    })
  })
}

const createCorrectAmountTest = (
  quote: FastBridgeQuote,
  amount: BigNumber,
  expected: BigNumber
) => {
  it(`${amount.toString()} -> ${expected.toString()}`, () => {
    expect(applyQuote(quote, amount)).toEqual(expected)
  })
}

const createQuoteTests = (
  quoteTemplate: FastBridgeQuote,
  originDecimals: number,
  destDecimals: number
) => {
  describe(`Origin decimals: ${originDecimals}, dest decimals: ${destDecimals}`, () => {
    describe(`origin:destination price 1:1`, () => {
      const quote: FastBridgeQuote = {
        ...quoteTemplate,
        maxOriginAmount: parseFixed('100000', originDecimals),
        destAmount: parseFixed('100000', destDecimals),
        fixedFee: parseFixed('1', originDecimals),
      }

      // 10 origin -> 9 dest
      createCorrectAmountTest(
        quote,
        parseFixed('10', originDecimals),
        parseFixed('9', destDecimals)
      )
      createZeroAmountTests(quote)
    })

    describe(`origin:destination price 1:1.0001`, () => {
      const quote: FastBridgeQuote = {
        ...quoteTemplate,
        maxOriginAmount: parseFixed('100000', originDecimals),
        destAmount: parseFixed('100010', destDecimals),
        fixedFee: parseFixed('1', originDecimals),
      }

      // 10 origin -> 9.0009 dest
      createCorrectAmountTest(
        quote,
        parseFixed('10', originDecimals),
        parseFixed('9.0009', destDecimals)
      )
      // Should round down 9999 * 1.0001 to 9999
      createCorrectAmountTest(
        quote,
        parseFixed('11', originDecimals).add(9999),
        parseFixed('10.001', destDecimals).add(9999)
      )
      createZeroAmountTests(quote)
    })

    describe(`origin:destination price 1:0.9999`, () => {
      const quote: FastBridgeQuote = {
        ...quoteTemplate,
        maxOriginAmount: parseFixed('100000', originDecimals),
        destAmount: parseFixed('99990', destDecimals),
        fixedFee: parseFixed('1', originDecimals),
      }

      // 10 origin -> 8.9991 dest
      createCorrectAmountTest(
        quote,
        parseFixed('10', originDecimals),
        parseFixed('8.9991', destDecimals)
      )
      // Should round down 1 * 0.9999 to 0
      createCorrectAmountTest(
        quote,
        parseFixed('11', originDecimals).add(1),
        parseFixed('9.999', destDecimals)
      )
      createZeroAmountTests(quote)
    })
  })
}

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
    // Equal decimals
    createQuoteTests(quote, 18, 18)

    // Bigger decimals
    createQuoteTests(quote, 6, 18)

    // Smaller decimals
    createQuoteTests(quote, 18, 6)

    it('Returns zero when max origin amount is zero', () => {
      const zeroQuote: FastBridgeQuote = {
        ...quote,
        maxOriginAmount: BigNumber.from(0),
      }
      const amount = zeroQuote.fixedFee.mul(2)
      expect(applyQuote(zeroQuote, amount)).toEqual(BigNumber.from(0))
    })

    it('Returns zero when dest amount is zero', () => {
      const zeroQuote: FastBridgeQuote = {
        ...quote,
        destAmount: BigNumber.from(0),
      }
      const amount = zeroQuote.fixedFee.mul(2)
      expect(applyQuote(zeroQuote, amount)).toEqual(BigNumber.from(0))
    })

    it('Returns zero when max origin amount and dest amount are zero', () => {
      const zeroQuote: FastBridgeQuote = {
        ...quote,
        maxOriginAmount: BigNumber.from(0),
        destAmount: BigNumber.from(0),
      }
      const amount = zeroQuote.fixedFee.mul(2)
      expect(applyQuote(zeroQuote, amount)).toEqual(BigNumber.from(0))
    })
  })
})
