/* eslint-disable @typescript-eslint/no-var-requires */
import { TextDecoder, TextEncoder } from 'util'

import { expect } from '@jest/globals'

Object.assign(global, { TextDecoder, TextEncoder })

jest.mock('@wagmi/core/chains', () => ({
  mainnet: { id: 1 },
  optimism: { id: 10 },
  cronos: { id: 25 },
  bsc: { id: 56 },
  polygon: { id: 137 },
  fantom: { id: 250 },
  boba: { id: 288 },
  metis: { id: 1088 },
  moonbeam: { id: 1284 },
  moonriver: { id: 1285 },
  canto: { id: 7700 },
  klaytn: { id: 8217 },
  base: { id: 8453 },
  arbitrum: { id: 42161 },
  avalanche: { id: 43114 },
  linea: { id: 59144 },
  blast: { id: 81457 },
  scroll: { id: 534352 },
  aurora: { id: 1313161554 },
  harmonyOne: { id: 1666600000 },
}))

const { BRIDGE_MAP } = require('../constants/bridgeMap')
const {
  AcceptedChainId,
  ACTIVE_CHAINS_BY_ID,
  CHAINS_ARR,
  CHAINS_BY_ID,
  CHAIN_IDS,
  isActiveChainId,
  ORDERED_CHAINS_BY_ID,
} = require('../constants/chains')
const { supportedChains } = require('../constants/chains/supportedChains')
const {
  EXISTING_BRIDGE_ROUTES,
}: {
  EXISTING_BRIDGE_ROUTES: Record<string, string[]>
} = require('../constants/existingBridgeRoutes')
const { tokenAddressToToken } = require('../constants/tokens')
const { WJEWEL } = require('../constants/tokens/bridgeable')
const {
  default: bridgeReducer,
  setFromChainId,
  setFromToken,
  setToChainId,
  setToToken,
} = require('../slices/bridge/reducer')
const {
  default: swapReducer,
  setSwapChainId,
  setSwapFromToken,
  setSwapToToken,
} = require('../slices/swap/reducer')

const DFK_CHAIN_ID = 53935
const DFK_ROUTE_SUFFIX = `-${DFK_CHAIN_ID}`
const DFK_USDC_ADDRESS = '0x3AD9DFE640E1A9Cc1D9B0948620820D975c3803a'

describe('retired chains', () => {
  it('excludes DFK Chain from active frontend registries', () => {
    expect(isActiveChainId(DFK_CHAIN_ID)).toBe(false)
    expect(CHAINS_ARR.some(({ id }) => id === DFK_CHAIN_ID)).toBe(false)
    expect(ACTIVE_CHAINS_BY_ID[DFK_CHAIN_ID]).toBeUndefined()
    expect(Object.values(CHAIN_IDS)).not.toContain(DFK_CHAIN_ID)
    expect(ORDERED_CHAINS_BY_ID).not.toContain(String(DFK_CHAIN_ID))
    expect(AcceptedChainId[DFK_CHAIN_ID]).toBeUndefined()
    expect(supportedChains.some(({ id }) => id === DFK_CHAIN_ID)).toBe(false)
  })

  it('excludes DFK Chain from bridge maps and real routes', () => {
    expect(Object.keys(BRIDGE_MAP)).not.toContain(String(DFK_CHAIN_ID))
    expect(
      Object.keys(EXISTING_BRIDGE_ROUTES).some((route) =>
        route.endsWith(DFK_ROUTE_SUFFIX)
      )
    ).toBe(false)
    expect(
      Object.values(EXISTING_BRIDGE_ROUTES)
        .flat()
        .some((route) => route.endsWith(DFK_ROUTE_SUFFIX))
    ).toBe(false)
  })

  it('retains DFK Chain metadata for historical transactions', () => {
    expect(CHAINS_BY_ID[DFK_CHAIN_ID]).toMatchObject({
      id: DFK_CHAIN_ID,
      name: 'DFK Chain',
    })
    expect(tokenAddressToToken(DFK_CHAIN_ID, DFK_USDC_ADDRESS)).toMatchObject({
      symbol: 'USDC',
    })
  })

  it('rejects DFK Chain as a new bridge or swap selection', () => {
    const initialBridgeState = bridgeReducer(undefined, { type: '@@INIT' })
    const initialSwapState = swapReducer(undefined, { type: '@@INIT' })

    expect(
      bridgeReducer(initialBridgeState, setFromChainId(DFK_CHAIN_ID))
    ).toEqual(initialBridgeState)
    expect(
      bridgeReducer(initialBridgeState, setToChainId(DFK_CHAIN_ID))
    ).toEqual(initialBridgeState)
    expect(bridgeReducer(initialBridgeState, setFromToken(WJEWEL))).toEqual(
      initialBridgeState
    )
    expect(bridgeReducer(initialBridgeState, setToToken(WJEWEL))).toEqual(
      initialBridgeState
    )
    expect(swapReducer(initialSwapState, setSwapChainId(DFK_CHAIN_ID))).toEqual(
      initialSwapState
    )
    expect(swapReducer(initialSwapState, setSwapFromToken(WJEWEL))).toEqual(
      initialSwapState
    )
    expect(swapReducer(initialSwapState, setSwapToToken(WJEWEL))).toEqual(
      initialSwapState
    )
  })
})
