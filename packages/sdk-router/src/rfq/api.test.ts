import { getAllQuotes } from './api'
import { FastBridgeQuoteAPI, unmarshallFastBridgeQuote } from './quote'

describe('getAllQuotes', () => {
  const quotesAPI: FastBridgeQuoteAPI[] = [
    {
      OriginChainID: 1,
      OriginTokenAddr: '0x0000000000000000000000000000000000000001',
      DestChainID: 2,
      DestTokenAddr: '0x0000000000000000000000000000000000000002',
      DestAmount: '3',
      MaxOriginAmount: '4',
      FixedFee: '5',
      OriginFastBridgeAddress: '10',
      DestFastBridgeAddress: '11',
      RelayerAddr: '0x0000000000000000000000000000000000000003',
      UpdatedAt: '2023-01-01T00:00:00.420Z',
    },
    {
      OriginChainID: 3,
      OriginTokenAddr: '0x0000000000000000000000000000000000000004',
      DestChainID: 4,
      DestTokenAddr: '0x0000000000000000000000000000000000000005',
      DestAmount: '6',
      MaxOriginAmount: '7',
      FixedFee: '8',
      OriginFastBridgeAddress: '20',
      DestFastBridgeAddress: '21',
      RelayerAddr: '0x0000000000000000000000000000000000000006',
      UpdatedAt: '2023-01-02T00:00:00.420Z',
    },
  ]

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
