import { BigNumber } from 'ethers'
import fetchMock from 'jest-fetch-mock'

import { getGasZipQuote } from './api'

describe('getGasZipQuote', () => {
  const emptyQuote = {
    amountOut: BigNumber.from(0),
    speed: 0,
    usd: 0,
  }

  beforeEach(() => {
    fetchMock.enableMocks()
    fetchMock.resetMocks()
    jest.spyOn(console, 'warn').mockImplementation(() => undefined)
  })

  afterEach(() => {
    jest.restoreAllMocks()
  })

  afterAll(() => {
    fetchMock.disableMocks()
  })

  it.each(['', ' ', '\t', '\r\n  '])(
    'preserves large integer quote amounts with JSON whitespace %j',
    async (whitespace) => {
      fetchMock.mockResponseOnce(
        `{"quotes":[{"chain":1,"expected":${whitespace}987654321098765432,"speed":30,"usd":1}]}`
      )

      await expect(
        getGasZipQuote(42161, 1, '10000000000000000')
      ).resolves.toEqual({
        amountOut: BigNumber.from('987654321098765432'),
        speed: 30,
        usd: 1,
      })
    }
  )

  it('keeps quotes with missing speed eligible for the default estimated time', async () => {
    fetchMock.mockResponseOnce('{"quotes":[{"expected":"100","usd":1}]}')

    await expect(getGasZipQuote(42161, 1, '100')).resolves.toEqual({
      amountOut: BigNumber.from(100),
      speed: 0,
      usd: 1,
    })
  })

  it.each(['"fast"', '"30"', '-1', 'null', '1e309'])(
    'returns no quote for invalid speed %s',
    async (speed) => {
      fetchMock.mockResponseOnce(
        `{"quotes":[{"expected":"100","speed":${speed},"usd":1}]}`
      )

      await expect(getGasZipQuote(42161, 1, '100')).resolves.toEqual(emptyQuote)
    }
  )

  it.each([
    ['the HTTP 200 disabled response', '{"error":"Temporarily Disabled"}'],
    ['an empty quote list', '{"quotes":[]}'],
    ['a null response', 'null'],
    ['a non-array quote list', '{"quotes":{}}'],
    ['a null quote', '{"quotes":[null]}'],
    ['a quote without an amount', '{"quotes":[{"speed":30}]}'],
    ['invalid JSON', 'temporarily unavailable'],
    ['an invalid amount', '{"quotes":[{"expected":"unavailable"}]}'],
  ])('returns no quote for %s', async (_description, body) => {
    fetchMock.mockResponseOnce(body)

    await expect(
      getGasZipQuote(42161, 1, '100000000000000000')
    ).resolves.toEqual(emptyQuote)
  })
})
