import { fetchJsonData } from './fetchJsonData'

const mockUrl =
  'https://synapsecns.github.io/sanguine/packages/synapse-interface/public/pauses/v1/paused-chains.json'

// Mock the global fetch function
global.fetch = jest.fn()

describe('fetchJsonData', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  test('retries on failure and eventually succeeds', async () => {
    const mockData = { key: 'value' }
    ;(fetch as jest.Mock)
      .mockResolvedValueOnce({ ok: false, status: 500 })
      .mockResolvedValueOnce({ ok: false, status: 500 })
      .mockResolvedValueOnce({
        ok: true,
        json: jest.fn().mockResolvedValueOnce(mockData),
      })

    const data = await fetchJsonData(mockUrl)
    expect(data).toEqual(mockData)
    expect(fetch).toHaveBeenCalledTimes(3)
  })

  test('retries the maximum number of times and fails', async () => {
    const maxRetries = 3
    ;(fetch as jest.Mock).mockResolvedValue({
      ok: false,
      status: 500,
    })

    await expect(fetchJsonData(mockUrl)).rejects.toThrow(
      `Retries failed fetching ${mockUrl}`
    )
    expect(fetch).toHaveBeenCalledTimes(maxRetries)
  }, 30000)
})
