import { expect } from '@jest/globals'

import { getFromChainIds } from '@/utils/routeMaker/getFromChainIds'

jest.mock('../constants/existingBridgeRoutes', () => ({
  __esModule: true,
  EXISTING_BRIDGE_ROUTES: {
    'GOHM-1': ['GOHM-10', 'GOHM-25', 'GOHM-56'],
    'GOHM-10': ['GOHM-1', 'GOHM-25', 'GOHM-56'],
    'GOHM-25': ['GOHM-1', 'GOHM-10', 'GOHM-56'],
    'GOHM-56': ['GOHM-1', 'GOHM-10', 'GOHM-25'],
    'HIGH-1': ['HIGH-56'],
    'HIGH-56': ['HIGH-1'],
    'USDC-1': ['USDC-10', 'USDC-25', 'USDC-56', 'NUSD-10'],
    'NUSD-10': ['USDC-1'],
    'USDC-10': ['USDC-1', 'USDC-25', 'USDC-56'],
    'USDC-25': ['USDC-1', 'USDC-10', 'USDC-56'],
    'USDC-56': ['USDC-1', 'USDC-10', 'USDC-25'],
    'SYN-1': ['SYN-10', 'SYN-25', 'SYN-56'],
    'SYN-10': ['SYN-1', 'SYN-25', 'SYN-56'],
    'SYN-25': ['SYN-1', 'SYN-10', 'SYN-56'],
    'SYN-56': ['SYN-1', 'SYN-10', 'SYN-25'],
    'XYZ-50': ['XYZ-1'],
    'XYZ-1': ['XYZ-50'],
  },
}))

jest.mock('../utils/flattenPausedTokens', () => ({
  __esModule: true,
  flattenPausedTokens: jest.fn(() => {
    return []
  }),
}))

jest.mock('../constants/tokens/master', () => ({
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

    expect(chainIds).toEqual([1, 10, 25, 56, 50])
  })

  it('has fromChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56, 50])
  })

  it('has fromChainId, toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1, 25, 56])
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1, 25, 56])
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds).toEqual([1, 25, 56])
  })

  it('has fromChainId, fromToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56])
  })

  it('has fromChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(chainIds).toEqual([1, 50])
  })

  it('has toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1])
  })

  it('has toChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(chainIds).toEqual([1])
  })

  it('has fromToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56])
  })

  it('has fromToken, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds.sort()).toEqual([1, 10, 25, 56].sort())
  })

  it('has toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(chainIds.sort()).toEqual([50, 1].sort())
  })
})
