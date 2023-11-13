import { expect } from '@jest/globals'

import { getFromChainIds } from '@/utils/routeMaker/getFromChainIds'
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

describe('getFromChainIds', () => {
  it('all entries null', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual(
      [
        1, 10, 1284, 8217, 43114, 53935, 1666600000, 56, 137, 288, 2000, 8453,
        42161, 7700, 250, 1285, 1088, 25, 1313161554,
      ].sort()
    )
  })

  it('has fromChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual(
      [
        1, 10, 1284, 8217, 43114, 53935, 1666600000, 56, 137, 288, 2000, 8453,
        42161, 7700, 250, 1285, 1088, 25, 1313161554,
      ].sort()
    )
  })

  it('has fromChainId, toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual(
      [
        56, 1, 137, 288, 8453, 42161, 43114, 7700, 53935, 1284, 1285, 25, 250,
        1088, 2000, 1313161554, 1666600000,
      ].sort()
    )
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 8453,
      fromTokenRouteSymbol: 'crvUSD',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([8453])
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds).toEqual([1, 8453, 42161, 43114])
  })

  it('has fromChainId, fromToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 8453,
      fromTokenRouteSymbol: 'crvUSD',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([8453])
  })

  it('has fromChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'crvUSD',
    })

    expect(chainIds).toEqual([
      10, 25, 56, 137, 250, 288, 1088, 2000, 7700, 8217, 8453, 42161, 43114,
      53935, 1313161554, 1666600000, 1,
    ])
  })

  it('has toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 8217,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1284, 43114, 53935, 1666600000, 1, 2000])
  })

  it('has toChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 1088,
      toTokenRouteSymbol: 'm.USDC',
    })

    expect(chainIds).toEqual([
      56, 1, 10, 137, 288, 42161, 43114, 7700, 25, 53935, 1313161554, 250,
      1666600000,
    ])
  })

  it('has toChainId, toToken, non-existent path', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: 'crvUSD',
    })

    expect(chainIds).toEqual([])
  })

  it('has fromToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'crvUSD',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([8453])
  })

  it('has fromToken, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds.sort()).toEqual(
      [
        1, 10, 137, 2000, 25, 288, 42161, 43114, 53935, 56, 7700, 8217, 8453,
      ].sort()
    )
  })

  it('has toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'USDC.e',
    })

    expect(chainIds.sort()).toEqual(
      [
        1, 10, 1088, 1313161554, 137, 1666600000, 2000, 25, 250, 288, 42161,
        43114, 53935, 56, 7700, 8217, 8453,
      ].sort()
    )
  })
})
