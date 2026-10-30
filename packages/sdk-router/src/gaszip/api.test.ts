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

  it('preserves integer quote amounts above Number.MAX_SAFE_INTEGER', async () => {
    fetchMock.mockResponseOnce(
      '{"quotes":[{"chain":1,"expected":987654321098765432,"speed":30,"usd":1}]}'
    )

    await expect(
      getGasZipQuote(42161, 1, '10000000000000000')
    ).resolves.toEqual({
      amountOut: BigNumber.from('987654321098765432'),
      speed: 30,
      usd: 1,
    })
  })

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
