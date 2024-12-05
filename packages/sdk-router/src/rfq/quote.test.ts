import { BigNumber } from '@ethersproject/bignumber'

import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  marshallFastBridgeQuote,
  unmarshallFastBridgeQuote,
} from './quote'

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
})
