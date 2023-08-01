import { expect } from '@jest/globals'

import { getToChainIds } from '@/utils/routeMaker/getToChainIds'

jest.mock('../constants/existing-bridge-routes', () => ({
  __esModule: true,
  EXISTING_BRIDGE_ROUTES: {
    'GOHM-1': ['GOHM-10', 'GOHM-25', 'GOHM-56'],
    'GOHM-10': ['GOHM-1', 'GOHM-25', 'GOHM-56'],
    'GOHM-25': ['GOHM-1', 'GOHM-10', 'GOHM-56'],
    'GOHM-56': ['GOHM-1', 'GOHM-10', 'GOHM-25'],
    'HIGHSTREET-1': ['HIGHSTREET-56'],
    'HIGHSTREET-56': ['HIGHSTREET-1'],
    'USDC-1': ['USDC-10', 'USDC-25', 'USDC-56', 'NUSD-10'],
    'NUSD-10': ['USDC-1', 'BUSD-56'],
    'USDC-10': ['USDC-1', 'USDC-25', 'USDC-56'],
    'USDC-25': ['USDC-1', 'USDC-10', 'USDC-56'],
    'USDC-56': ['USDC-1', 'USDC-10', 'USDC-25'],
    'SYN-1': ['SYN-10', 'SYN-25', 'SYN-56'],
    'SYN-10': ['SYN-1', 'SYN-25', 'SYN-56'],
    'SYN-25': ['SYN-1', 'SYN-10', 'SYN-56'],
    'SYN-56': ['SYN-1', 'SYN-10', 'SYN-25'],
    'XYZ-50': ['XYZ-1', 'ABC-1'],
    'XYZ-1': ['XYZ-50', 'ABC-56'],
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

describe('getToChainIds', () => {
  it('all entries null', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([10, 25, 56, 1, 50])
  })

  it('has fromChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([10, 25, 56, 50])
  })

  it('has fromChainId, toChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual([10, 25, 56, 50].sort())
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([10, 25, 56])
  })

  it('has fromChainId, fromToken, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 10,
      fromTokenRouteSymbol: 'NUSD',
      toChainId: null,
      toTokenRouteSymbol: 'BUSD',
    })

    expect(chainIds).toEqual([1, 56])
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds).toEqual([10, 25, 56])
  })

  it('has fromChainId, fromToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([10, 25, 56])
  })

  it('has fromChainId, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(chainIds).toEqual([50])
  })

  it('has toChainId, signle', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual([10, 25, 56, 1, 50].sort())
  })

  it('has toChainId, multiple', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 25,
      toTokenRouteSymbol: null,
    })

    expect(chainIds.sort()).toEqual([10, 25, 56, 1, 50].sort())
  })

  it('has toChainId, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(chainIds.sort()).toEqual([1, 50].sort())
  })

  it('has fromToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(chainIds).toEqual([10, 25, 56, 1])
  })

  it('has fromToken, toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(chainIds).toEqual([10, 25, 56, 1])
  })

  it('has toToken', () => {
    const chainIds = getToChainIds({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(chainIds.sort()).toEqual([1, 50].sort())
  })
})
