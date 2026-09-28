/** @jest-environment node */
import { bridgeQuoteSingleFlight } from '@/utils/bridgeQuoteSingleFlight'

const params = {
  fromChainId: 1,
  toChainId: 1337,
  fromToken: 'SYN',
  toToken: 'SYN_CORE',
  fromAmount: '1000000000000000000',
  fromSender: 'alice',
  toRecipient: 'bob',
  slippagePercentage: 0.1,
}

describe('bridge quote single flight', () => {
  it('coalesces identical in-flight SDK work without caching resolved quotes', async () => {
    let resolve: (value: any) => void
    const sdk = {
      bridgeV2: jest
        .fn()
        .mockImplementationOnce(
          () =>
            new Promise((done) => {
              resolve = done
            })
        )
        .mockResolvedValue(['new quote']),
    }
    const first = bridgeQuoteSingleFlight(sdk, params)
    const second = bridgeQuoteSingleFlight(
      sdk,
      Object.fromEntries(Object.entries(params).reverse())
    )
    expect(second).toBe(first)
    await Promise.resolve()
    expect(sdk.bridgeV2).toHaveBeenCalledTimes(1)
    resolve(['first quote'])
    expect(await first).toEqual(['first quote'])
    expect(await bridgeQuoteSingleFlight(sdk, params)).toEqual(['new quote'])
    expect(sdk.bridgeV2).toHaveBeenCalledTimes(2)
  })

  it('separates every SDK parameter and SDK instance', async () => {
    const sdk = { bridgeV2: jest.fn().mockResolvedValue([]) }
    const calls = [bridgeQuoteSingleFlight(sdk, params)]
    for (const field of Object.keys(params)) {
      calls.push(
        bridgeQuoteSingleFlight(sdk, {
          ...params,
          [field]: String(params[field]) + 'changed',
        })
      )
    }
    const otherSDK = { bridgeV2: jest.fn().mockResolvedValue([]) }
    calls.push(bridgeQuoteSingleFlight(otherSDK, params))
    await Promise.all(calls)
    expect(sdk.bridgeV2).toHaveBeenCalledTimes(Object.keys(params).length + 1)
    expect(otherSDK.bridgeV2).toHaveBeenCalledTimes(1)
  })

  it('releases failed requests so a later refresh can retry', async () => {
    const sdk = {
      bridgeV2: jest
        .fn()
        .mockRejectedValueOnce(new Error('offline'))
        .mockResolvedValue([]),
    }
    const request = bridgeQuoteSingleFlight(sdk, params)
    expect(bridgeQuoteSingleFlight(sdk, params)).toBe(request)
    await expect(request).rejects.toThrow('offline')
    await expect(bridgeQuoteSingleFlight(sdk, params)).resolves.toEqual([])
    expect(sdk.bridgeV2).toHaveBeenCalledTimes(2)
  })

  it('distinguishes undefined, null, and omitted parameters', async () => {
    const sdk = { bridgeV2: jest.fn().mockResolvedValue([]) }
    await Promise.all([
      bridgeQuoteSingleFlight(sdk, { ...params, deadline: undefined }),
      bridgeQuoteSingleFlight(sdk, { ...params, deadline: null }),
      bridgeQuoteSingleFlight(sdk, params),
    ])
    expect(sdk.bridgeV2).toHaveBeenCalledTimes(3)
  })
})
