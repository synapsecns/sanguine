import fetchMock from 'jest-fetch-mock'
import { parseFixed } from '@ethersproject/bignumber'

import {
  getAllQuotes,
  getBestRelayerQuote,
  PutRFQResponseAPI,
  RelayerQuote,
} from './api'
import { Ticker } from './ticker'
import { FastBridgeQuoteAPI, unmarshallFastBridgeQuote } from './quote'

const OK_RESPONSE_TIME = 1900
const SLOW_RESPONSE_TIME = 2100

const delayedAPIPromise = (
  body: string,
  delay: number
): Promise<{ body: string }> => {
  return new Promise((resolve) => setTimeout(() => resolve({ body }), delay))
}

describe('getAllQuotes', () => {
  const quotesAPI: FastBridgeQuoteAPI[] = [
    {
      origin_chain_id: 1,
      origin_token_addr: '0x0000000000000000000000000000000000000001',
      dest_chain_id: 2,
      dest_token_addr: '0x0000000000000000000000000000000000000002',
      dest_amount: '3',
      max_origin_amount: '4',
      fixed_fee: '5',
      origin_fast_bridge_address: '10',
      dest_fast_bridge_address: '11',
      relayer_addr: '0x0000000000000000000000000000000000000003',
      updated_at: '2023-01-01T00:00:00.420Z',
    },
    {
      origin_chain_id: 3,
      origin_token_addr: '0x0000000000000000000000000000000000000004',
      dest_chain_id: 4,
      dest_token_addr: '0x0000000000000000000000000000000000000005',
      dest_amount: '6',
      max_origin_amount: '7',
      fixed_fee: '8',
      origin_fast_bridge_address: '20',
      dest_fast_bridge_address: '21',
      relayer_addr: '0x0000000000000000000000000000000000000006',
      updated_at: '2023-01-02T00:00:00.420Z',
    },
  ]

  beforeEach(() => {
    fetchMock.enableMocks()
  })

  afterEach(() => {
    fetchMock.resetMocks()
  })

  describe('Returns a list of quotes', () => {
    it('when the response is ok', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(quotesAPI))
      const result = await getAllQuotes()
      expect(result).toEqual([
        unmarshallFastBridgeQuote(quotesAPI[0]),
        unmarshallFastBridgeQuote(quotesAPI[1]),
      ])
    })

    it('when the response takes a long, but reasonable time to return', async () => {
      fetchMock.mockResponseOnce(() =>
        delayedAPIPromise(JSON.stringify(quotesAPI), OK_RESPONSE_TIME)
      )
      const result = await getAllQuotes()
      expect(result).toEqual([
        unmarshallFastBridgeQuote(quotesAPI[0]),
        unmarshallFastBridgeQuote(quotesAPI[1]),
      ])
    })
  })

  describe('Returns an empty array', () => {
    beforeEach(() => {
      jest.spyOn(console, 'error').mockImplementation(() => {
        // Do nothing
      })
    })

    afterEach(() => {
      jest.restoreAllMocks()
    })

    it('when the response is not ok', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(quotesAPI), { status: 500 })
      const result = await getAllQuotes()
      expect(result).toEqual([])
      expect(console.error).toHaveBeenCalled()
    })

    it('when fetch throws an error', async () => {
      fetchMock.mockRejectOnce(new Error('Error fetching quotes'))
      const result = await getAllQuotes()
      expect(result).toEqual([])
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response takes too long to return', async () => {
      fetchMock.mockResponseOnce(() =>
        delayedAPIPromise(JSON.stringify(quotesAPI), SLOW_RESPONSE_TIME)
      )
      const result = await getAllQuotes()
      expect(result).toEqual([])
      expect(console.error).toHaveBeenCalled()
    })
  })
})

describe('getBestRelayerQuote', () => {
  const bigAmount = parseFixed('1', 24)
  const bigAmountStr = '1000000000000000000000000'
  const relayerAddress = '0x0000000000000000000000000000000000001337'
  const quoteID = 'acbdef-123456'
  const userAddress = '0x0000000000000000000000000000000000007331'

  const ticker: Ticker = {
    originToken: {
      chainId: 1,
      token: '0x0000000000000000000000000000000000000001',
    },
    destToken: {
      chainId: 2,
      token: '0x0000000000000000000000000000000000000002',
    },
  }

  const noQuotesFound: PutRFQResponseAPI = {
    success: false,
    reason: 'No quotes found',
  }

  const quoteFound: PutRFQResponseAPI = {
    success: true,
    quote_id: quoteID,
    dest_amount: bigAmountStr,
    relayer_address: relayerAddress,
  }

  const quote: RelayerQuote = {
    destAmount: bigAmount,
    relayerAddress,
    quoteID,
  }

  const quoteZero: RelayerQuote = {
    destAmount: parseFixed('0'),
  }

  beforeEach(() => {
    fetchMock.enableMocks()
  })

  afterEach(() => {
    fetchMock.resetMocks()
  })

  describe('Returns a non-zero quote', () => {
    it('when the response is ok', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(quoteFound))
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quote)
    })

    it('when the response takes a long, but reasonable time to return', async () => {
      fetchMock.mockResponseOnce(() =>
        delayedAPIPromise(JSON.stringify(quoteFound), OK_RESPONSE_TIME)
      )
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quote)
    })

    it('when the response does not contain quote ID', async () => {
      const responseWithoutID = { ...quoteFound, quote_id: undefined }
      const quoteWithoutID = { ...quote, quoteID: undefined }
      fetchMock.mockResponseOnce(JSON.stringify(responseWithoutID))
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteWithoutID)
    })
  })

  describe('Returns a zero quote', () => {
    beforeEach(() => {
      jest.spyOn(console, 'error').mockImplementation(() => {
        // Do nothing
      })
    })

    afterEach(() => {
      jest.restoreAllMocks()
    })

    it('when the user address is not provided', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(quoteFound))
      const result = await getBestRelayerQuote(ticker, bigAmount)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response is not ok', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(quoteFound), { status: 500 })
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response success is false', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(noQuotesFound))
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response takes too long to return', async () => {
      fetchMock.mockResponseOnce(() =>
        delayedAPIPromise(JSON.stringify(quoteFound), SLOW_RESPONSE_TIME)
      )
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response does not contain dest amount', async () => {
      const responseWithoutDestAmount = {
        ...quoteFound,
        dest_amount: undefined,
      }
      fetchMock.mockResponseOnce(JSON.stringify(responseWithoutDestAmount))
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response does not contain relayer address', async () => {
      const responseWithoutRelayerAddress = {
        ...quoteFound,
        relayer_address: undefined,
      }
      fetchMock.mockResponseOnce(JSON.stringify(responseWithoutRelayerAddress))
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })

    it('when the response dest amount is zero', async () => {
      const responseWithZeroDestAmount = { ...quoteFound, dest_amount: '0' }
      fetchMock.mockResponseOnce(JSON.stringify(responseWithZeroDestAmount))
      const result = await getBestRelayerQuote(ticker, bigAmount, userAddress)
      expect(result).toEqual(quoteZero)
      expect(console.error).toHaveBeenCalled()
    })
  })
})
