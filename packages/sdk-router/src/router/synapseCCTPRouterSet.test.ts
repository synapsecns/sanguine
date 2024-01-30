import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, parseFixed } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

import { getTestProvider } from '../constants/testProviders'
import {
  ROUTER_ADDRESS_MAP,
  CCTP_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_CCTP,
  CCTP_SUPPORTED_CHAIN_IDS,
  SupportedChainId,
} from '../constants'
import { ChainProvider } from './routerSet'
import { SynapseCCTPRouterSet } from './synapseCCTPRouterSet'
import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { CCTPRouterQuery, Query, SynapseModuleSet } from '../module'

export const createSlippageTests = (
  moduleSet: SynapseModuleSet,
  originQuery: Query,
  destQuery: Query,
  expectedOriginMinAmountOut: BigNumber,
  expectedDestMinAmountOut: BigNumber,
  slipNumerator: number,
  slipDenominator: number
) => {
  // Create a copy of the queries to check that the original query is not modified
  const originQueryCopy = { ...originQuery }
  const destQueryCopy = { ...destQuery }

  it('Applies slippage to origin query', () => {
    const { originQuery: originQueryNew } = moduleSet.applySlippage(
      originQuery,
      destQuery,
      slipNumerator,
      slipDenominator
    )
    expect(originQueryNew).toEqual({
      ...originQueryCopy,
      minAmountOut: expectedOriginMinAmountOut,
    })
  })

  it('Applies slippage to dest query', () => {
    const { destQuery: destQueryNew } = moduleSet.applySlippage(
      originQuery,
      destQuery,
      slipNumerator,
      slipDenominator
    )
    expect(destQueryNew).toEqual({
      ...destQueryCopy,
      minAmountOut: expectedDestMinAmountOut,
    })
  })

  it('Does not modify the original queries', () => {
    moduleSet.applySlippage(
      originQuery,
      destQuery,
      slipNumerator,
      slipDenominator
    )
    expect(originQuery).toEqual(originQueryCopy)
    expect(destQuery).toEqual(destQueryCopy)
  })
}

describe('SynapseCCTPRouterSet', () => {
  const ethProvider: Provider = getTestProvider(SupportedChainId.ETH)

  const arbProvider: Provider = getTestProvider(SupportedChainId.ARBITRUM)

  // Chain where CCTP is unlikely to be deployed
  const moonbeamProvider: Provider = getTestProvider(SupportedChainId.MOONBEAM)

  const testProviders: ChainProvider[] = [
    {
      chainId: SupportedChainId.ETH,
      provider: ethProvider,
    },
    {
      chainId: SupportedChainId.ARBITRUM,
      provider: arbProvider,
    },
    {
      chainId: SupportedChainId.MOONBEAM,
      provider: moonbeamProvider,
    },
  ]

  const routerSet = new SynapseCCTPRouterSet(testProviders)

  describe('#constructor', () => {
    it('Creates SynapseCCTPRouter instances for chains with CCTP', () => {
      expect(routerSet.routers[SupportedChainId.ETH]).toBeDefined()
      expect(routerSet.routers[SupportedChainId.ARBITRUM]).toBeDefined()
    })

    it('Uses correct addresses for each chain', () => {
      expect(routerSet.routers[SupportedChainId.ETH].address).toEqual(
        CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
      )
      expect(routerSet.routers[SupportedChainId.ARBITRUM].address).toEqual(
        CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Does not create SynapseCCTPRouter instances for chains without CCTP', () => {
      expect(routerSet.routers[SupportedChainId.MOONBEAM]).toBeUndefined()
    })

    it('Does not create SynapseCCTPRouter instances for chains without providers', () => {
      expect(routerSet.routers[SupportedChainId.AVALANCHE]).toBeUndefined()
    })

    it('Correct bridge module name', () => {
      expect(routerSet.bridgeModuleName).toEqual('SynapseCCTP')
    })
  })

  describe('getEstimatedTime', () => {
    it('Returns the correct estimated time for all supported chains', () => {
      CCTP_SUPPORTED_CHAIN_IDS.forEach((chainId) => {
        expect(routerSet.getEstimatedTime(Number(chainId))).toEqual(
          MEDIAN_TIME_CCTP[chainId as keyof typeof MEDIAN_TIME_CCTP]
        )
      })
    })

    it('Throws error for unsupported chain with a provider', () => {
      expect(() =>
        routerSet.getEstimatedTime(SupportedChainId.MOONBEAM)
      ).toThrow('No estimated time for chain 1284')
    })

    it('Throws error for unsupported chain without a provider', () => {
      // 5 is the chain ID for Goerli testnet
      expect(() => routerSet.getEstimatedTime(5)).toThrow(
        'No estimated time for chain 5'
      )
    })
  })

  describe('getModuleWithAddress', () => {
    it('Returns the correct router given correct address', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.ETH,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toEqual(routerSet.routers[SupportedChainId.ETH])
    })

    it('Returns undefined given incorrect address', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.ETH,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })

    it('Returns undefined given unknown chainId', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.AVALANCHE,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })

    it('Returns undefined given chainId without CCTP', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.MOONBEAM,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })
  })

  describe('getSynapseCCTPRouter', () => {
    it('Returns the correct router for supported chain', () => {
      const router = routerSet.getSynapseCCTPRouter(SupportedChainId.ETH)
      expect(router).toEqual(routerSet.routers[SupportedChainId.ETH])
      expect(router).toBeInstanceOf(SynapseCCTPRouter)
    })

    it('Throws error for unsupported chain', () => {
      expect(() =>
        routerSet.getSynapseCCTPRouter(SupportedChainId.AVALANCHE)
      ).toThrow('No module found for chain 43114')
    })
  })

  describe('applySlippage', () => {
    const originQuery: CCTPRouterQuery = {
      routerAdapter: '1',
      tokenOut: '2',
      minAmountOut: parseFixed('1000', 18),
      deadline: BigNumber.from(3),
      rawParams: '4',
    }

    const destQuery: CCTPRouterQuery = {
      routerAdapter: '5',
      tokenOut: '6',
      minAmountOut: parseFixed('2000', 6),
      deadline: BigNumber.from(8),
      rawParams: '9',
    }

    describe('0% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        parseFixed('1000', 18),
        parseFixed('2000', 6),
        0,
        10000
      )
    })

    describe('1% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        parseFixed('990', 18),
        parseFixed('1980', 6),
        100,
        10000
      )
    })

    describe('10% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        parseFixed('900', 18),
        parseFixed('1800', 6),
        1000,
        10000
      )
    })

    describe('100% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        Zero,
        Zero,
        10000,
        10000
      )
    })

    describe('Rounds down', () => {
      const originQueryPlusOne: CCTPRouterQuery = {
        ...originQuery,
        minAmountOut: originQuery.minAmountOut.add(1),
      }
      const destQueryPlusOne: CCTPRouterQuery = {
        ...destQuery,
        minAmountOut: destQuery.minAmountOut.add(1),
      }

      createSlippageTests(
        routerSet,
        originQueryPlusOne,
        destQueryPlusOne,
        parseFixed('990', 18).add(1),
        parseFixed('1980', 6).add(1),
        100,
        10000
      )
    })
  })

  // TODO (Chi): more tests
})
