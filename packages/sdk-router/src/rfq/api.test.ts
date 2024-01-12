import fetchMock from 'jest-fetch-mock'

import { getAllQuotes } from './api'
import { FastBridgeQuoteAPI, unmarshallFastBridgeQuote } from './quote'

const OK_RESPONSE_TIME = 1900
const SLOW_RESPONSE_TIME = 2100

const delayedAPIPromise = (
  quotes: FastBridgeQuoteAPI[],
  delay: number
): Promise<{ body: string }> => {
  return new Promise((resolve) =>
    setTimeout(() => resolve({ body: JSON.stringify(quotes) }), delay)
  )
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
        delayedAPIPromise(quotesAPI, OK_RESPONSE_TIME)
      )
      const result = await getAllQuotes()
      expect(result).toEqual([
        unmarshallFastBridgeQuote(quotesAPI[0]),
        unmarshallFastBridgeQuote(quotesAPI[1]),
      ])
    })
  })

  describe('Returns an empty array', () => {
    it('when the response is not ok', async () => {
      fetchMock.mockResponseOnce(JSON.stringify(quotesAPI), { status: 500 })
      const result = await getAllQuotes()
      expect(result).toEqual([])
    })

    it('when fetch throws an error', async () => {
      fetchMock.mockRejectOnce(new Error('Error fetching quotes'))
      const result = await getAllQuotes()
      expect(result).toEqual([])
    })

    it('when the response takes too long to return', async () => {
      fetchMock.mockResponseOnce(() =>
        delayedAPIPromise(quotesAPI, SLOW_RESPONSE_TIME)
      )
      const result = await getAllQuotes()
      expect(result).toEqual([])
    })
  })
})
