import { expect } from '@jest/globals'

import { getToChainIds } from '@/utils/routeMaker/getToChainIds'
import { mockRoutes } from './__fixtures__/mockRoutes'

jest.mock('../constants/existingBridgeRoutes', () => ({
  get EXISTING_BRIDGE_ROUTES() {
    return mockRoutes
  },
}))

jest.mock('../utils/flattenPausedTokens', () => ({
  __esModule: true,
  flattenPausedTokens: jest.fn(() => {
    return []
  }),
}))

jest.mock('../constants/tokens/bridgeable', () => ({
  __esModule: true,
  '*': jest.fn(() => {
    return []
  }),
}))

describe('getToChainIds', () => {
  it('all entries null', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([
      8217, 43114, 53935, 1666600000, 1284, 1, 10, 25, 137, 250, 288, 1088,
      7700, 42161, 1313161554, 56, 2000, 8453, 1285,
    ])
  })

  it('has fromChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: 8217,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1284, 43114, 53935, 1666600000, 1, 2000])
  })

  it('has fromChainId, toChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual(
      [
        10, 1088, 1284, 1285, 1313161554, 137, 1666600000, 2000, 25, 250, 288,
        42161, 43114, 53935, 56, 7700, 8217, 8453,
      ].sort()
    )
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([
      10, 25, 56, 137, 250, 288, 1088, 2000, 7700, 8217, 8453, 42161, 43114,
      53935, 1313161554, 1666600000,
    ])
  })

  it('has fromChainId, fromToken, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 10,
      fromTokenRouteSymbol: 'nUSD',
      toChainId: null,
      toTokenRouteSymbol: 'BUSD',
    })

    expect(chainIds).toEqual([56])
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds).toEqual([
      56, 1, 10, 137, 288, 8453, 42161, 43114, 7700, 25, 2000, 8217, 53935,
      1313161554, 1088, 250, 1666600000,
    ])
  })

  it('has fromChainId, fromToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([
      10, 25, 56, 137, 250, 288, 1088, 2000, 7700, 8217, 8453, 42161, 43114,
      53935, 1313161554, 1666600000,
    ])
  })

  it('has fromChainId, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds).toEqual([
      10, 25, 56, 137, 288, 2000, 7700, 8217, 8453, 42161, 43114, 53935,
    ])
  })

  it('has toChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 25,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual(
      [
        1, 10, 1088, 1284, 1285, 1313161554, 137, 1666600000, 2000, 25, 250,
        288, 42161, 43114, 53935, 56, 7700, 8217, 8453,
      ].sort()
    )
  })

  it('has toChainId, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 1088,
      toTokenRouteSymbol: 'm.USDC',
    })

    expect(chainIds).toEqual([1088])
  })

  it('has fromToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([
      10, 25, 56, 137, 250, 288, 1088, 2000, 7700, 8217, 8453, 42161, 43114,
      53935, 1313161554, 1666600000, 1,
    ])
  })

  it('has fromToken, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: 'BUSD',
    })

    expect(chainIds).toEqual([56])
  })

  it('has toToken, m.USDC', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'm.USDC',
    })

    expect(chainIds).toEqual([1088])
  })
})
