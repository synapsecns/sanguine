import { fetchJSONData } from './fetchJsonData'

const mockUrl =
  'https://raw.githubusercontent.com/synapsecns/sanguine/test/maintenance/packages/synapse-interface/public/pauses/v1/paused-chains.json'

// Mock the global fetch function
global.fetch = jest.fn()

describe('fetchJSONData', () => {
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

    const data = await fetchJSONData(mockUrl)
    expect(data).toEqual(mockData)
    expect(fetch).toHaveBeenCalledTimes(3)
  })

  test('retries the maximum number of times and fails', async () => {
    const maxRetries = 5
    ;(fetch as jest.Mock).mockResolvedValue({
      ok: false,
      status: 500,
    })

    await expect(fetchJSONData(mockUrl)).rejects.toThrow(
      `Retries failed fetching ${mockUrl}`
    )
    expect(fetch).toHaveBeenCalledTimes(maxRetries)
  }, 30000)
})
