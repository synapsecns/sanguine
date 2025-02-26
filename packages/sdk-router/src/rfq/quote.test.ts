import { BigNumber, parseFixed } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  marshallFastBridgeQuote,
  unmarshallFastBridgeQuote,
  applyQuote,
} from './quote'

const expectEqualBigNumbers = (a: BigNumber, b: BigNumber) => {
  expect(a.toString()).toEqual(b.toString())
}

const expectNotEqualBigNumbers = (a: BigNumber, b: BigNumber) => {
  expect(a.toString()).not.toEqual(b.toString())
}

const createZeroAmountTests = (quote: FastBridgeQuote) => {
  describe('Returns zero', () => {
    it('If origin amount is zero', () => {
      expectEqualBigNumbers(applyQuote(quote, Zero), Zero)
    })

    it('If origin amount is greater than max origin amount', () => {
      const amount = quote.maxOriginAmount.add(quote.fixedFee).add(1)
      expectEqualBigNumbers(applyQuote(quote, amount), Zero)
    })

    it('If dest amount is lower than fixed fee', () => {
      const amount = quote.fixedFee
        .mul(quote.maxOriginAmount)
        .div(quote.destAmount)
      expectEqualBigNumbers(applyQuote(quote, amount), Zero)
    })
  })

  describe('Returns non-zero', () => {
    it('If origin amount is equal to max origin amount', () => {
      expectNotEqualBigNumbers(applyQuote(quote, quote.maxOriginAmount), Zero)
    })
  })
}

const createCorrectAmountTest = (
  quote: FastBridgeQuote,
  amount: BigNumber,
  expected: BigNumber
) => {
  it(`${amount.toString()} -> ${expected.toString()}`, () => {
    expectEqualBigNumbers(applyQuote(quote, amount), expected)
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
        fixedFee: parseFixed('1', destDecimals),
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
        fixedFee: parseFixed('1', destDecimals),
      }

      // 10 origin -> 10.0010 dest -> 9.0010 dest after fee
      createCorrectAmountTest(
        quote,
        parseFixed('10', originDecimals),
        parseFixed('9.0010', destDecimals)
      )
      createZeroAmountTests(quote)
    })

    describe(`origin:destination price 1:0.9999`, () => {
      const quote: FastBridgeQuote = {
        ...quoteTemplate,
        maxOriginAmount: parseFixed('100000', originDecimals),
        destAmount: parseFixed('99990', destDecimals),
        fixedFee: parseFixed('1', destDecimals),
      }

      // 10 origin -> 9.9990 dest -> 8.9990 dest after fee
      createCorrectAmountTest(
        quote,
        parseFixed('10', originDecimals),
        parseFixed('8.9990', destDecimals)
      )
      createZeroAmountTests(quote)
    })
  })
}

const createRoundDownTest = (
  quoteTemplate: FastBridgeQuote,
  maxOriginAmount: BigNumber,
  destAmount: BigNumber,
  fixedFee: BigNumber,
  amountIn: BigNumber,
  expected: BigNumber
) => {
  describe(`Rounds down with price ${maxOriginAmount.toString()} -> ${destAmount.toString()} and fixed fee ${fixedFee.toString()}`, () => {
    const quote: FastBridgeQuote = {
      ...quoteTemplate,
      maxOriginAmount,
      destAmount,
      fixedFee,
    }

    createCorrectAmountTest(quote, amountIn, expected)
  })
}

describe('quote', () => {
  const quoteAPI: FastBridgeQuoteAPI = {
    origin_chain_id: 1,
    origin_token_addr: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
    dest_chain_id: 2,
    dest_token_addr: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    dest_amount: '4000000000000000000000',
    max_origin_amount: '3000000000000000000000',
    fixed_fee: '1000000000000000000',
    origin_fast_bridge_address: '0x1',
    dest_fast_bridge_address: '0x2',
    relayer_addr: '0xB300efF6B57AA09e5fCcf7221FCB9E676A74d931',
    updated_at: '2023-01-02T03:04:05.678Z',
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
    originFastBridge: '0x1',
    destFastBridge: '0x2',
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
    createRoundDownTest(
      quote,
      parseFixed('1234', 18),
      parseFixed('2345', 18),
      parseFixed('1', 18),
      parseFixed('2', 18),
      // 2 * 2345 / 1234 - 1 = 2.800648298217179902
      BigNumber.from('2800648298217179902')
    )

    // // Bigger decimals
    createQuoteTests(quote, 6, 18)
    createRoundDownTest(
      quote,
      parseFixed('1234', 6),
      parseFixed('2345', 18),
      parseFixed('1', 18),
      parseFixed('2', 6),
      // 2 * 2345 / 1234 - 1 = 2.800648298217179902
      BigNumber.from('2800648298217179902')
    )

    // Smaller decimals
    createQuoteTests(quote, 18, 6)
    createRoundDownTest(
      quote,
      parseFixed('1234', 18),
      parseFixed('2345', 6),
      parseFixed('1', 6),
      parseFixed('2', 18),
      // 2 * 2345 / 1234 - 1 = 2.800648298217179902
      BigNumber.from('2800648')
    )

    it('Returns zero when max origin amount is zero', () => {
      const zeroQuote: FastBridgeQuote = {
        ...quote,
        maxOriginAmount: Zero,
      }
      const amount = zeroQuote.fixedFee.mul(2)
      expectEqualBigNumbers(applyQuote(zeroQuote, amount), Zero)
    })

    it('Returns zero when dest amount is zero', () => {
      const zeroQuote: FastBridgeQuote = {
        ...quote,
        destAmount: Zero,
      }
      const amount = zeroQuote.fixedFee.mul(2)
      expectEqualBigNumbers(applyQuote(zeroQuote, amount), Zero)
    })

    it('Returns zero when max origin amount and dest amount are zero', () => {
      const zeroQuote: FastBridgeQuote = {
        ...quote,
        maxOriginAmount: Zero,
        destAmount: Zero,
      }
      const amount = zeroQuote.fixedFee.mul(2)
      expectEqualBigNumbers(applyQuote(zeroQuote, amount), Zero)
    })
  })
})
