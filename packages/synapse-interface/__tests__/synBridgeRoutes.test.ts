import {
  HYPERCORE_CHAIN_ID,
  RELAY_SUPPORTED_CHAIN_IDS,
} from '@synapsecns/sdk-router'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existingBridgeRoutes'
import { SYN, USDC } from '@/constants/tokens/bridgeable'
import { Token } from '@/utils/types'
import { HYPERLIQUID } from '@/constants/chains/master'
import { getFromChainIds } from '@/utils/routeMaker/getFromChainIds'
import { getToChainIds } from '@/utils/routeMaker/getToChainIds'
import { getToTokens } from '@/utils/routeMaker/getToTokens'
import { getPausedBridgeModuleNamesForRoute } from '@/utils/getPausedBridgeModuleNamesForRoute'
import { isValidBridgeModule } from '@/components/Maintenance/functions/isValidBridgeModule'

describe('SYN routes in the existing bridge selectors', () => {
  it('supports module and ALL pauses on the virtual destination', () => {
    expect(isValidBridgeModule('SYN')).toBe(true)
    for (const bridgeModuleName of ['SYN', 'ALL']) {
      const pausedModules = [{ toChainId: 1337, bridgeModuleName }]
      expect(
        getPausedBridgeModuleNamesForRoute({
          pausedModules,
          fromChainId: 1,
          toChainId: 1337,
        })
      ).toContain('SYN')
      expect(
        getPausedBridgeModuleNamesForRoute({
          pausedModules,
          fromChainId: 1,
          toChainId: 999,
        })
      ).not.toContain('SYN')
    }
  })

  it.each([
    [1, 999],
    [999, 1],
    [1, 1337],
  ])(
    'offers SYN from %s to %s with token metadata',
    (fromChainId, toChainId) => {
      expect(
        getToTokens({
          fromChainId,
          toChainId,
          fromTokenRouteSymbol: 'SYN',
          toTokenRouteSymbol: null,
        })
      ).toContain(`SYN-${toChainId}`)
      expect(SYN.addresses[toChainId]).toBeDefined()
      expect(SYN.decimals[toChainId]).toBe(toChainId === 1337 ? 8 : 18)
    }
  )

  it('uses native token IDs only on HyperCore while preserving EVM address validation', () => {
    const tokenId = '0xf5f05eb8b9aa92365465f06daf5889c9'
    expect(HYPERLIQUID.id).toBe(HYPERCORE_CHAIN_ID)
    expect(SYN.addresses[HYPERLIQUID.id]).toBe(tokenId)
    for (const chainId of [1, 999]) {
      expect(
        () => new Token({ addresses: { [chainId]: tokenId }, priorityRank: 0 })
      ).toThrow()
    }
    expect(
      () =>
        new Token({
          addresses: { [HYPERLIQUID.id]: '0x1234' },
          priorityRank: 0,
        })
    ).toThrow()
  })

  it('keeps HyperCore destination-only and requires Ethereum as its SYN origin', () => {
    expect(
      getFromChainIds({
        fromChainId: null,
        fromTokenRouteSymbol: 'SYN',
        toChainId: 1337,
        toTokenRouteSymbol: 'SYN',
      })
    ).toEqual([1])
    expect(EXISTING_BRIDGE_ROUTES['SYN-1337']).toBeUndefined()
    expect(EXISTING_BRIDGE_ROUTES['SYN-999']).toContain('SYN-1')
  })

  it('keeps HyperCore in the destination selector after all fields are selected', () => {
    const selectedRoute = {
      fromChainId: 1,
      fromTokenRouteSymbol: 'SYN',
      toChainId: 1337,
      toTokenRouteSymbol: 'SYN',
    }
    expect(getToChainIds(selectedRoute)).toEqual(
      getToChainIds({ ...selectedRoute, toChainId: null })
    )
    expect(getToChainIds(selectedRoute)).toContain(1337)
  })

  it('does not advertise the new destinations from other legacy SYN chains', () => {
    Object.entries(EXISTING_BRIDGE_ROUTES).forEach(([origin, destinations]) => {
      if (origin !== 'SYN-1') {
        expect(destinations).not.toContain('SYN-999')
        expect(destinations).not.toContain('SYN-1337')
      }
    })
    expect(EXISTING_BRIDGE_ROUTES['SYN-1']).toContain('SYN-42161')
  })

  it('offers Hyperliquid USDC from every Relay origin token', () => {
    expect(USDC.addresses[1337]).toBe('0x00000000000000000000000000000000')
    expect(USDC.decimals[1337]).toBe(8)
    Object.entries(EXISTING_BRIDGE_ROUTES).forEach(([origin, destinations]) => {
      const originChainId = Number(origin.split('-').at(-1))
      expect(destinations.includes('USDC-1337')).toBe(
        RELAY_SUPPORTED_CHAIN_IDS.includes(originChainId)
      )
    })
    expect(EXISTING_BRIDGE_ROUTES['DAI-8453']).toContain('USDC-1337')
    expect(EXISTING_BRIDGE_ROUTES['HYPE-999']).toContain('USDC-1337')
    expect(
      getToTokens({
        fromChainId: 999,
        fromTokenRouteSymbol: 'HYPE',
        toChainId: 1337,
        toTokenRouteSymbol: null,
      })
    ).toContain('USDC-1337')
    expect(EXISTING_BRIDGE_ROUTES['USDC-25'] ?? []).not.toContain('USDC-1337')
  })
})
