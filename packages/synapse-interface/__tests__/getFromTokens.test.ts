import { expect } from '@jest/globals'

import { getFromTokens } from '@/utils/routeMaker/getFromTokens'

jest.mock('../constants/existingBridgeRoutes', () => ({
  __esModule: true,
  EXISTING_BRIDGE_ROUTES: {
    'GOHM-1': ['GOHM-10', 'GOHM-25', 'GOHM-56'],
    'GOHM-10': ['GOHM-1', 'GOHM-25', 'GOHM-56'],
    'GOHM-25': ['GOHM-1', 'GOHM-10', 'GOHM-56'],
    'GOHM-56': ['GOHM-1', 'GOHM-10', 'GOHM-25'],
    'HIGHSTREET-1': ['HIGHSTREET-56'],
    'HIGHSTREET-56': ['HIGHSTREET-1'],
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

describe('getFromTokens', () => {
  it('all entries null', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(
      [
        'GOHM-1',
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'HIGHSTREET-1',
        'HIGHSTREET-56',
        'USDC-1',
        'USDC-10',
        'NUSD-10',
        'USDC-25',
        'USDC-56',
        'SYN-1',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'XYZ-50',
        'XYZ-1',
      ].sort()
    )
  })

  it('has fromChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual([
      'GOHM-1',
      'HIGHSTREET-1',
      'USDC-1',
      'SYN-1',
      'XYZ-1',
    ])
  })

  it('has fromChainId, toChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual(['GOHM-1', 'USDC-1', 'SYN-1'])
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(['GOHM-1', 'USDC-1', 'SYN-1'].sort())
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(fromTokens).toEqual(['USDC-1'])
  })

  it('has fromChainId, fromToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual([
      'GOHM-1',
      'HIGHSTREET-1',
      'USDC-1',
      'SYN-1',
      'XYZ-1',
    ])
  })

  it('has fromChainId, fromToken, NUSD', () => {
    const fromTokens = getFromTokens({
      fromChainId: 10,
      fromTokenRouteSymbol: 'NUSD',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(
      ['GOHM-10', 'NUSD-10', 'USDC-10', 'SYN-10'].sort()
    )
  })

  it('has fromChainId, toToken, single', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(fromTokens).toEqual(['XYZ-1'])
  })

  it('has fromChainId, toToken, multiple', () => {
    const fromTokens = getFromTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(fromTokens).toEqual(['USDC-1'])
  })

  it('has toChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual(['XYZ-1'])
  })

  it('has toChainId, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(fromTokens).toEqual(['XYZ-1'])
  })

  it('has fromToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(
      [
        'GOHM-1',
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'HIGHSTREET-1',
        'HIGHSTREET-56',
        'USDC-1',
        'NUSD-10',
        'USDC-10',
        'USDC-25',
        'USDC-56',
        'SYN-1',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'XYZ-50',
        'XYZ-1',
      ].sort()
    )
  })

  it('has fromToken, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(fromTokens.sort()).toEqual(
      ['USDC-1', 'NUSD-10', 'USDC-10', 'USDC-25', 'USDC-56'].sort()
    )
  })

  it('has toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(fromTokens.sort()).toEqual(['XYZ-50', 'XYZ-1'].sort())
  })

  it('has FromChainId and fromToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 10,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(
      ['NUSD-10', 'GOHM-10', 'USDC-10', 'SYN-10'].sort()
    )
  })
})
