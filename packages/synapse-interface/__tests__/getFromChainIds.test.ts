import { expect } from '@jest/globals'

import { getFromChainIds } from '@/utils/generateRoutePossibilities'
import { Token } from '@/utils/types'

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
      fromToken: null,
      toChainId: null,
      toToken: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56, 50])
  })

  it('has fromChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromToken: null,
      toChainId: null,
      toToken: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56, 50])
  })

  it('has fromChainId, toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromToken: null,
      toChainId: 10,
      toToken: null,
    })

    expect(chainIds).toEqual([1, 25, 56])
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromToken: 'USDC',
      toChainId: 10,
      toToken: null,
    })

    expect(chainIds).toEqual([1, 25, 56])
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromToken: 'USDC',
      toChainId: 10,
      toToken: 'USDC',
    })

    expect(chainIds).toEqual([1, 25, 56])
  })

  it('has fromChainId, fromToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromToken: 'USDC',
      toChainId: null,
      toToken: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56])
  })

  it('has fromChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: 1,
      fromToken: null,
      toChainId: null,
      toToken: 'XYZ',
    })

    expect(chainIds).toEqual([1, 50])
  })

  it('has toChainId', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromToken: null,
      toChainId: 50,
      toToken: null,
    })

    expect(chainIds).toEqual([1])
  })

  it('has toChainId, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromToken: null,
      toChainId: 50,
      toToken: 'XYZ',
    })

    expect(chainIds).toEqual([1])
  })

  it('has fromToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromToken: 'USDC',
      toChainId: null,
      toToken: null,
    })

    expect(chainIds).toEqual([1, 10, 25, 56])
  })

  it('has fromToken, toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromToken: 'USDC',
      toChainId: null,
      toToken: 'USDC',
    })

    expect(chainIds).toEqual([1, 10, 25, 56])
  })

  it('has toToken', () => {
    const chainIds = getFromChainIds({
      fromChainId: null,
      fromToken: null,
      toChainId: null,
      toToken: 'XYZ',
    })

    expect(chainIds).toEqual([50, 1])
  })
})
