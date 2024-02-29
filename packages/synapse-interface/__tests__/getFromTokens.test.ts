import { expect } from '@jest/globals'

import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
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
        'AVAX-1284',
        'AVAX-1666600000',
        'AVAX-43114',
        'AVAX-53935',
        'AVAX-8217',
        'BTC.b-43114',
        'BTC.b-53935',
        'BTC.b-8217',
        'BUSD-56',
        'DAI-1',
        'DAI-10',
        'DAI-137',
        'DAI-2000',
        'DAI-288',
        'DAI-42161',
        'DAI-8217',
        'DAI-8453',
        'DAI.e-43114',
        'DOG-1',
        'DOG-137',
        'DOG-56',
        'ETH-1',
        'ETH-10',
        'ETH-288',
        'ETH-42161',
        'ETH-53935',
        'ETH-7700',
        'ETH-8453',
        'FRAX-1',
        'FRAX-2000',
        'FRAX-42161',
        'FTM-250',
        'FTM-53935',
        'GMX-42161',
        'GMX-43114',
        'H2O-1',
        'H2O-10',
        'H2O-1284',
        'H2O-1285',
        'H2O-137',
        'H2O-42161',
        'H2O-43114',
        'H2O-56',
        'HIGH-1',
        'HIGH-56',
        'JEWEL-1666600000',
        'JEWEL-43114',
        'JEWEL-53935',
        'JEWEL-8217',
        'JUMP-1088',
        'JUMP-250',
        'JUMP-56',
        'KLAY-53935',
        'KLAY-8217',
        'L2DAO-10',
        'L2DAO-42161',
        'LINK-1',
        'LINK-8217',
        'MATIC-137',
        'MATIC-53935',
        'MOVR-1284',
        'MOVR-1285',
        'NEWO-1',
        'NEWO-42161',
        'NEWO-43114',
        'NFD-137',
        'NFD-2000',
        'NFD-43114',
        'NFD-56',
        'NOTE-7700',
        'PEPE-1',
        'PEPE-42161',
        'PLS-10',
        'PLS-42161',
        'SDT-1',
        'SDT-1666600000',
        'SDT-250',
        'SDT-42161',
        'SDT-43114',
        'SFI-1',
        'SFI-43114',
        'SYN-1',
        'SYN-10',
        'SYN-1088',
        'SYN-1284',
        'SYN-1285',
        'SYN-1313161554',
        'SYN-137',
        'SYN-1666600000',
        'SYN-2000',
        'SYN-25',
        'SYN-250',
        'SYN-288',
        'SYN-42161',
        'SYN-43114',
        'SYN-56',
        'SYN-7700',
        'SYN-8453',
        'UNIDX-1',
        'UNIDX-10',
        'UNIDX-250',
        'UNIDX-42161',
        'UNIDX-8453',
        'USDC-1',
        'USDC-10',
        'USDC-137',
        'USDC-2000',
        'USDC-25',
        'USDC-288',
        'USDC-42161',
        'USDC-43114',
        'USDC-53935',
        'USDC-56',
        'USDC-7700',
        'USDC-8217',
        'USDC-8453',
        'USDC.e-10',
        'USDC.e-1313161554',
        'USDC.e-42161',
        'USDC.e-43114',
        'USDT-1',
        'USDT-10',
        'USDT-137',
        'USDT-2000',
        'USDT-288',
        'USDT-42161',
        'USDT-43114',
        'USDT-56',
        'USDT-7700',
        'USDT-8217',
        'USDT.e-1313161554',
        'USDT.e-43114',
        'USDbC-8453',
        'VSTA-1',
        'VSTA-42161',
        'WAVAX-43114',
        'WBTC-1',
        'WBTC-2000',
        'WBTC-8217',
        'WETH-1088',
        'WETH.e-43114',
        'WFTM-250',
        'WJEWEL-53935',
        'WKLAY-8217',
        'WMATIC-137',
        'WMOVR-1285',
        'axlUSDC-8453',
        'crvUSD-8453',
        'gOHM-1',
        'gOHM-10',
        'gOHM-1088',
        'gOHM-1284',
        'gOHM-1285',
        'gOHM-137',
        'gOHM-1666600000',
        'gOHM-25',
        'gOHM-250',
        'gOHM-288',
        'gOHM-42161',
        'gOHM-43114',
        'gOHM-56',
        'm.USDC-1088',
        'nETH-10',
        'nETH-1088',
        'nETH-1666600000',
        'nETH-250',
        'nETH-288',
        'nETH-42161',
        'nETH-43114',
        'nETH-7700',
        'nETH-8453',
        'nUSD-1',
        'nUSD-10',
        'nUSD-1088',
        'nUSD-1313161554',
        'nUSD-137',
        'nUSD-1666600000',
        'nUSD-25',
        'nUSD-250',
        'nUSD-288',
        'nUSD-42161',
        'nUSD-43114',
        'nUSD-56',
        'nUSD-7700',
        'sUSD-10',
        'synFRAX-1284',
        'synFRAX-1285',
        'synFRAX-1666600000',
        'synFRAX-250',
        'synJEWEL-1666600000',
        'veSOLAR-1284',
        'veSOLAR-1285',
        'xJEWEL-1666600000',
        'xJEWEL-53935',
      ].sort()
    )
  })

  it('has fromChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: 8453,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual([
      'DAI-8453',
      'ETH-8453',
      'SYN-8453',
      'UNIDX-8453',
      'USDC-8453',
      'USDbC-8453',
      'axlUSDC-8453',
      'crvUSD-8453',
      'nETH-8453',
    ])
  })

  it('has fromChainId, toChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: 8453,
      fromTokenRouteSymbol: null,
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual([
      'DAI-8453',
      'ETH-8453',
      'SYN-8453',
      'UNIDX-8453',
      'USDC-8453',
      'USDbC-8453',
      'axlUSDC-8453',
      'crvUSD-8453',
      'nETH-8453',
    ])
  })

  it('has fromChainId, fromToken, toChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: 8453,
      fromTokenRouteSymbol: 'crvUSD',
      toChainId: 10,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(
      [
        'DAI-8453',
        'ETH-8453',
        'SYN-8453',
        'UNIDX-8453',
        'USDC-8453',
        'USDbC-8453',
        'axlUSDC-8453',
        'crvUSD-8453',
        'nETH-8453',
      ].sort()
    )
  })

  it('has fromChainId, fromToken, toChainId, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 8453,
      fromTokenRouteSymbol: 'crvUSD',
      toChainId: 10,
      toTokenRouteSymbol: 'USDC',
    })

    expect(fromTokens).toEqual([
      'DAI-8453',
      'USDC-8453',
      'USDbC-8453',
      'axlUSDC-8453',
      'crvUSD-8453',
    ])
  })

  it('has fromChainId, fromToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 8453,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual([
      'DAI-8453',
      'ETH-8453',
      'SYN-8453',
      'UNIDX-8453',
      'USDC-8453',
      'USDbC-8453',
      'axlUSDC-8453',
      'crvUSD-8453',
      'nETH-8453',
    ])
  })

  it('has fromChainId, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 8453,
      fromTokenRouteSymbol: null,
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(fromTokens).toEqual([
      'DAI-8453',
      'USDC-8453',
      'USDbC-8453',
      'axlUSDC-8453',
      'crvUSD-8453',
    ])
  })

  it('has toChainId', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 8453,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens).toEqual([
      'DAI-1',
      'DAI-10',
      'DAI-42161',
      'DAI.e-43114',
      'ETH-1',
      'ETH-10',
      'ETH-288',
      'ETH-7700',
      'ETH-42161',
      'ETH-53935',
      'FRAX-42161',
      'SYN-1',
      'SYN-10',
      'SYN-25',
      'SYN-56',
      'SYN-137',
      'SYN-250',
      'SYN-288',
      'SYN-1088',
      'SYN-1284',
      'SYN-1285',
      'SYN-2000',
      'SYN-7700',
      'SYN-42161',
      'SYN-43114',
      'SYN-1313161554',
      'SYN-1666600000',
      'UNIDX-1',
      'UNIDX-10',
      'UNIDX-250',
      'UNIDX-42161',
      'USDC-1',
      'USDC-10',
      'USDC-42161',
      'USDC-43114',
      'USDC.e-10',
      'USDC.e-42161',
      'USDC.e-43114',
      'USDT-1',
      'USDT-10',
      'USDT-42161',
      'USDT-43114',
      'USDT.e-43114',
      'WETH-1088',
      'WETH.e-43114',
      'nETH-10',
      'nETH-250',
      'nETH-288',
      'nETH-1088',
      'nETH-7700',
      'nETH-42161',
      'nETH-43114',
      'nETH-1666600000',
      'nUSD-42161',
      'nUSD-43114',
      'sUSD-10',
    ])
  })

  it('has toChainId, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: null,
      toChainId: 8453,
      toTokenRouteSymbol: 'ETH',
    })

    expect(fromTokens).toEqual([
      'ETH-1',
      'ETH-10',
      'ETH-288',
      'ETH-7700',
      'ETH-42161',
      'ETH-53935',
      'WETH-1088',
      'WETH.e-43114',
      'nETH-10',
      'nETH-250',
      'nETH-288',
      'nETH-1088',
      'nETH-7700',
      'nETH-42161',
      'nETH-43114',
      'nETH-1666600000',
    ])
  })

  it('has fromToken, toToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: null,
      fromTokenRouteSymbol: 'crvUSD',
      toChainId: null,
      toTokenRouteSymbol: 'USDC',
    })

    expect(fromTokens.sort()).toEqual(
      [
        'BUSD-56',
        'DAI-1',
        'DAI-10',
        'DAI-137',
        'DAI-288',
        'DAI-42161',
        'DAI-8453',
        'DAI.e-43114',
        'FRAX-42161',
        'NOTE-7700',
        'USDC-1',
        'USDC-10',
        'USDC-137',
        'USDC-2000',
        'USDC-25',
        'USDC-288',
        'USDC-42161',
        'USDC-43114',
        'USDC-53935',
        'USDC-56',
        'USDC-7700',
        'USDC-8217',
        'USDC-8453',
        'USDC.e-10',
        'USDC.e-1313161554',
        'USDC.e-42161',
        'USDC.e-43114',
        'USDT-1',
        'USDT-10',
        'USDT-137',
        'USDT-288',
        'USDT-42161',
        'USDT-43114',
        'USDT-56',
        'USDT-7700',
        'USDT.e-1313161554',
        'USDT.e-43114',
        'USDbC-8453',
        'axlUSDC-8453',
        'crvUSD-8453',
        'm.USDC-1088',
        'nUSD-1',
        'nUSD-10',
        'nUSD-1088',
        'nUSD-1313161554',
        'nUSD-137',
        'nUSD-1666600000',
        'nUSD-25',
        'nUSD-250',
        'nUSD-288',
        'nUSD-42161',
        'nUSD-43114',
        'nUSD-56',
        'nUSD-7700',
        'sUSD-10',
      ].sort()
    )
  })

  it('has FromChainId and fromToken', () => {
    const fromTokens = getFromTokens({
      fromChainId: 10,
      fromTokenRouteSymbol: 'USDC',
      toChainId: null,
      toTokenRouteSymbol: null,
    })

    expect(fromTokens.sort()).toEqual(
      [
        'DAI-10',
        'ETH-10',
        'H2O-10',
        'L2DAO-10',
        'PLS-10',
        'SYN-10',
        'UNIDX-10',
        'USDC-10',
        'USDC.e-10',
        'USDT-10',
        'gOHM-10',
        'nETH-10',
        'nUSD-10',
        'sUSD-10',
      ].sort()
    )
  })
})
