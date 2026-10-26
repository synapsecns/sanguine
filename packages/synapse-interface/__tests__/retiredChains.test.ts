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
const {
  DISPLAY_POOLS_BY_CHAIN,
  POOL_BY_ROUTER_INDEX,
  tokenAddressToToken,
} = require('../constants/tokens')
const {
  EXISTING_SWAP_ROUTES,
  SWAP_CHAIN_IDS,
} = require('../constants/existingSwapRoutes')
const { USDB, WJEWEL } = require('../constants/tokens/bridgeable')
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

describe.each([
  {
    chainId: 53935,
    name: 'DFK Chain',
    tokenAddress: '0x3AD9DFE640E1A9Cc1D9B0948620820D975c3803a',
    symbol: 'USDC',
    exclusiveToken: WJEWEL,
  },
  {
    chainId: 81457,
    name: 'Blast',
    tokenAddress: '0x3194B0A295D87fDAA54DF852c248F7a6BAF6c6e0',
    symbol: 'nUSD',
    exclusiveToken: USDB,
  },
])(
  'retired chain: $name',
  ({ chainId, name, tokenAddress, symbol, exclusiveToken }) => {
    it('excludes retired chain from active frontend registries', () => {
      expect(isActiveChainId(chainId)).toBe(false)
      expect(CHAINS_ARR.some(({ id }) => id === chainId)).toBe(false)
      expect(ACTIVE_CHAINS_BY_ID[chainId]).toBeUndefined()
      expect(Object.values(CHAIN_IDS)).not.toContain(chainId)
      expect(ORDERED_CHAINS_BY_ID).not.toContain(String(chainId))
      expect(AcceptedChainId[chainId]).toBeUndefined()
      expect(supportedChains.some(({ id }) => id === chainId)).toBe(false)
    })

    it('excludes retired chain from bridge maps and real routes', () => {
      expect(Object.keys(BRIDGE_MAP)).not.toContain(String(chainId))
      expect(
        Object.keys(EXISTING_BRIDGE_ROUTES).some((route) =>
          route.endsWith(`-${chainId}`)
        )
      ).toBe(false)
      expect(
        Object.values(EXISTING_BRIDGE_ROUTES)
          .flat()
          .some((route) => route.endsWith(`-${chainId}`))
      ).toBe(false)
    })

    it('excludes retired chains from swap routes and URL chain choices', () => {
      expect(SWAP_CHAIN_IDS).not.toContain(chainId)
      expect(
        Object.keys(EXISTING_SWAP_ROUTES).some((route) =>
          route.endsWith(`-${chainId}`)
        )
      ).toBe(false)
      expect(
        Object.values(EXISTING_SWAP_ROUTES)
          .flat()
          .some((route: string) => route.endsWith(`-${chainId}`))
      ).toBe(false)
    })

    it('retains retired chain metadata for historical transactions', () => {
      expect(CHAINS_BY_ID[chainId]).toMatchObject({
        id: chainId,
        name,
      })
      expect(tokenAddressToToken(chainId, tokenAddress)).toMatchObject({
        symbol,
      })
    })

    it('rejects retired chain as a new bridge or swap selection', () => {
      const initialBridgeState = bridgeReducer(undefined, { type: '@@INIT' })
      const initialSwapState = swapReducer(undefined, { type: '@@INIT' })

      expect(
        bridgeReducer(initialBridgeState, setFromChainId(chainId))
      ).toEqual(initialBridgeState)
      expect(bridgeReducer(initialBridgeState, setToChainId(chainId))).toEqual(
        initialBridgeState
      )
      expect(
        bridgeReducer(initialBridgeState, setFromToken(exclusiveToken))
      ).toEqual(initialBridgeState)
      expect(
        bridgeReducer(initialBridgeState, setToToken(exclusiveToken))
      ).toEqual(initialBridgeState)
      expect(swapReducer(initialSwapState, setSwapChainId(chainId))).toEqual(
        initialSwapState
      )
      expect(
        swapReducer(initialSwapState, setSwapFromToken(exclusiveToken))
      ).toEqual(initialSwapState)
      expect(
        swapReducer(initialSwapState, setSwapToToken(exclusiveToken))
      ).toEqual(initialSwapState)
    })
  }
)

describe('retired Blast pools', () => {
  it('hides pool promotion while preserving pool and LP recovery metadata', () => {
    expect(DISPLAY_POOLS_BY_CHAIN[81457]).toBeUndefined()
    expect(POOL_BY_ROUTER_INDEX.blast2pool).toMatchObject({
      chainId: 81457,
      swapAddresses: { 81457: '0xa4bd1AAD7cF04567c10f38FC4355E91bba32aC9c' },
      addresses: { 81457: '0xb758FF41E0A811b9eB0Ff973D373c354588524b6' },
    })
    expect(POOL_BY_ROUTER_INDEX.blastethpool).toMatchObject({ chainId: 81457 })
  })
})
