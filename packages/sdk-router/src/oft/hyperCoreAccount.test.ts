import { getAddress } from '@ethersproject/address'

import { HyperCoreAccountClient } from './hyperCoreAccount'

const recipient = '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd'
const response = (payload: unknown, ok = true, status = 200): Response =>
  ({ ok, status, json: async () => payload } as Response)

let fetchMock: jest.SpyInstance

beforeEach(() => {
  jest.useFakeTimers({ doNotFake: ['performance'] }).setSystemTime(0)
  fetchMock = jest.spyOn(global, 'fetch')
  fetchMock.mockResolvedValue(response({ role: 'user' }))
})

afterEach(() => {
  jest.restoreAllMocks()
  jest.useRealTimers()
})

it.each([
  ['user', true],
  ['vault', true],
  ['subAccount', true],
  ['missing', false],
  ['agent', false],
])('maps the %s role to active=%s', async (role, active) => {
  fetchMock.mockResolvedValue(response({ role }))
  const client = new HyperCoreAccountClient()
  expect(await client.isAccountActive(recipient)).toBe(active)
})

it('normalizes the address and shares cached results across quote amounts', async () => {
  const client = new HyperCoreAccountClient()
  await client.isAccountActive(getAddress(recipient))
  await client.isAccountActive(recipient)
  await client.isAccountActive(recipient)
  expect(fetchMock).toHaveBeenCalledTimes(1)
  expect(fetchMock).toHaveBeenCalledWith(
    'https://api.hyperliquid.xyz/info',
    expect.objectContaining({
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ type: 'userRole', user: recipient }),
      signal: expect.any(AbortSignal),
    })
  )
})

it.each([
  ['user', 60 * 60 * 1000],
  ['missing', 30_000],
])('expires cached %s results after %i ms', async (role, ttl) => {
  fetchMock.mockResolvedValue(response({ role }))
  const client = new HyperCoreAccountClient()
  await client.isAccountActive(recipient)
  jest.setSystemTime(ttl - 1)
  await client.isAccountActive(recipient)
  expect(fetchMock).toHaveBeenCalledTimes(1)
  jest.setSystemTime(ttl)
  await client.isAccountActive(recipient)
  expect(fetchMock).toHaveBeenCalledTimes(2)
})

it('shares a pending request for the same normalized address', async () => {
  let resolveResponse!: (value: Response) => void
  fetchMock.mockReturnValue(
    new Promise<Response>((resolve) => {
      resolveResponse = resolve
    })
  )
  const client = new HyperCoreAccountClient()
  const requests = [
    client.isAccountActive(recipient),
    client.isAccountActive(getAddress(recipient)),
  ]
  expect(fetchMock).toHaveBeenCalledTimes(1)
  resolveResponse(response({ role: 'user' }))
  expect(await Promise.all(requests)).toEqual([true, true])
})

it.each([null, [], 'user', {}, { role: true }, { role: 'futureRole' }])(
  'rejects malformed or unknown role response %j without caching it',
  async (payload) => {
    fetchMock.mockResolvedValueOnce(response(payload))
    const client = new HyperCoreAccountClient()
    await expect(client.isAccountActive(recipient)).rejects.toThrow(
      'Invalid HyperCore account role response.'
    )
    expect(await client.isAccountActive(recipient)).toBe(true)
    expect(fetchMock).toHaveBeenCalledTimes(2)
  }
)

it.each(['HTTP', 'network', 'JSON'])(
  'retries after a %s failure',
  async (failure) => {
    if (failure === 'HTTP') {
      fetchMock.mockResolvedValueOnce(response(null, false, 429))
    } else if (failure === 'network') {
      fetchMock.mockRejectedValueOnce(new Error('Network unavailable'))
    } else {
      fetchMock.mockResolvedValueOnce({
        ok: true,
        json: async () => {
          throw new Error('Invalid JSON')
        },
      } as unknown as Response)
    }
    const client = new HyperCoreAccountClient()
    await expect(client.isAccountActive(recipient)).rejects.toThrow()
    expect(await client.isAccountActive(recipient)).toBe(true)
    expect(fetchMock).toHaveBeenCalledTimes(2)
  }
)

it('times out and permits retry even when the response body stalls', async () => {
  fetchMock.mockResolvedValueOnce({
    ok: true,
    json: () => new Promise(() => undefined),
  } as unknown as Response)
  const client = new HyperCoreAccountClient()
  const pending = client.isAccountActive(recipient)
  const rejection = expect(pending).rejects.toThrow(
    'HyperCore account request timed out.'
  )
  await Promise.resolve()
  jest.advanceTimersByTime(10_000)
  await rejection
  expect(fetchMock.mock.calls[0][1].signal.aborted).toBe(true)
  expect(await client.isAccountActive(recipient)).toBe(true)
  expect(fetchMock).toHaveBeenCalledTimes(2)
})

it.each(['', '0x1234', 'not-an-address', '0x' + 'z'.repeat(40)])(
  'rejects invalid address %s before fetching',
  async (address) => {
    const client = new HyperCoreAccountClient()
    await expect(client.isAccountActive(address)).rejects.toThrow()
    expect(fetchMock).not.toHaveBeenCalled()
  }
)

it('bounds completed cache entries by evicting the oldest result', async () => {
  const client = new HyperCoreAccountClient()
  const address = (index: number) => '0x' + index.toString(16).padStart(40, '0')
  for (let index = 1; index <= 1001; index++) {
    await client.isAccountActive(address(index))
  }
  await client.isAccountActive(address(1001))
  expect(fetchMock).toHaveBeenCalledTimes(1001)
  await client.isAccountActive(address(1))
  expect(fetchMock).toHaveBeenCalledTimes(1002)
})
