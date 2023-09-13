import { expect } from '@jest/globals'

import { getToTokens } from '@/utils/routeMaker/getToTokens'

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

jest.mock('../constants/tokens/bridgeable', () => ({
  __esModule: true,
  '*': jest.fn(() => {
    return []
  }),
}))

describe('getToTokens', () => {
  it('all entries null', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(toTokens.sort()).toEqual(
      [
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'GOHM-1',
        'HIGH-56',
        'HIGH-1',
        'NUSD-10',
        'USDC-10',
        'USDC-25',
        'USDC-56',
        'USDC-1',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'SYN-1',
        'XYZ-1',
        'XYZ-50',
      ].sort()
    )
  })

  it('has fromChainId', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(toTokens.sort()).toEqual(
      [
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'HIGH-56',
        'NUSD-10',
        'USDC-10',
        'USDC-25',
        'USDC-56',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'XYZ-50',
      ].sort()
    )
  })

  it('has fromChainId, toChainId', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(toTokens).toEqual(['GOHM-10', 'USDC-10', 'NUSD-10', 'SYN-10'])
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(toTokens).toEqual(['USDC-10', 'NUSD-10'])
  })

  it('has fromTokenRouteSymbol, toChainId', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(toTokens).toEqual(['USDC-10', 'NUSD-10'])
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(toTokens).toEqual(['USDC-10', 'NUSD-10'])
  })

  it('has fromChainId, fromToken', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(toTokens.sort()).toEqual(
      ['NUSD-10', 'USDC-10', 'USDC-25', 'USDC-56'].sort()
    )
  })

  it('has fromChainId, toToken, single', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(toTokens.sort()).toEqual(
      [
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'HIGH-56',
        'USDC-10',
        'USDC-25',
        'USDC-56',
        'NUSD-10',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'XYZ-50',
      ].sort()
    )
  })

  it('has fromChainId, toToken, multiple', () => {
    const toTokens = getToTokens({
      fromChainId: 1,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(toTokens.sort()).toEqual(
      [
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'HIGH-56',
        'USDC-10',
        'USDC-25',
        'USDC-56',
        'NUSD-10',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'XYZ-50',
      ].sort()
    )
  })

  it('has toChainId', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: null,
    })

    expect(toTokens).toEqual(['XYZ-50'])
  })

  it('has toChainId, toToken', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 50,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(toTokens).toEqual(['XYZ-50'])
  })

  it('has fromToken', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(toTokens.sort()).toEqual(
      ['NUSD-10', 'USDC-1', 'USDC-10', 'USDC-25', 'USDC-56'].sort()
    )
  })

  it('has fromToken, toToken', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(toTokens.sort()).toEqual(
      ['NUSD-10', 'USDC-1', 'USDC-10', 'USDC-25', 'USDC-56'].sort()
    )
  })

  it('has toToken', () => {
    const toTokens = getToTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'XYZ',
    })

    expect(toTokens.sort()).toEqual(
      [
        'GOHM-1',
        'GOHM-10',
        'GOHM-25',
        'GOHM-56',
        'HIGH-1',
        'HIGH-56',
        'NUSD-10',
        'SYN-1',
        'SYN-10',
        'SYN-25',
        'SYN-56',
        'USDC-1',
        'USDC-10',
        'USDC-25',
        'USDC-56',
        'XYZ-1',
        'XYZ-50',
      ].sort()
    )
  })
})
