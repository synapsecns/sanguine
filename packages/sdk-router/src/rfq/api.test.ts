import { getAllQuotes } from './api'
import { FastBridgeQuoteAPI, unmarshallFastBridgeQuote } from './quote'

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
  ];


  it('returns an empty array when the response is not ok', async () => {
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 500,
        ok: false,
      })
    ) as any

    const result = await getAllQuotes()
    expect(result).toEqual([])
  })

  it('returns a list of quotes when the response is ok', async () => {
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 200,
        ok: true,
        json: () => Promise.resolve(quotesAPI),
      })
    ) as any

    const result = await getAllQuotes()
    // You might need to adjust this depending on how your unmarshallFastBridgeQuote function works
    expect(result).toEqual([
      unmarshallFastBridgeQuote(quotesAPI[0]),
      unmarshallFastBridgeQuote(quotesAPI[1]),
    ])
  })

  it('Integration test', async () => {
    global.fetch = require('node-fetch')
    const result = await getAllQuotes()
    console.log('Quotes: ' + JSON.stringify(result, null, 2))
    expect(result.length).toBeGreaterThan(0)
  })
})
