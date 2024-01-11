import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, parseFixed } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

import { getTestProvider } from '../constants/testProviders'
import {
  ROUTER_ADDRESS_MAP,
  CCTP_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_BRIDGE,
  SUPPORTED_CHAIN_IDS,
  SupportedChainId,
} from '../constants'
import { ChainProvider } from './routerSet'
import { SynapseRouterSet } from './synapseRouterSet'
import { SynapseRouter } from './synapseRouter'
import { RouterQuery } from '../module'
import { createSlippageTests } from './synapseCCTPRouterSet.test'

describe('SynapseRouterSet', () => {
  const ethProvider: Provider = getTestProvider(SupportedChainId.ETH)

  const arbProvider: Provider = getTestProvider(SupportedChainId.ARBITRUM)

  const testProviders: ChainProvider[] = [
    {
      chainId: SupportedChainId.ETH,
      provider: ethProvider,
    },
    {
      chainId: SupportedChainId.ARBITRUM,
      provider: arbProvider,
    },
  ]

  const routerSet = new SynapseRouterSet(testProviders)

  describe('#constructor', () => {
    it('Creates SynapseRouter instances for each chain', () => {
      expect(routerSet.routers[SupportedChainId.ETH]).toBeDefined()
      expect(routerSet.routers[SupportedChainId.ARBITRUM]).toBeDefined()
    })

    it('Uses correct addresses for each chain', () => {
      expect(routerSet.routers[SupportedChainId.ETH].address).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
      )
      expect(routerSet.routers[SupportedChainId.ARBITRUM].address).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Does not create SynapseRouter instances for chains without providers', () => {
      expect(routerSet.routers[SupportedChainId.AVALANCHE]).toBeUndefined()
    })

    it('Correct bridge module name', () => {
      expect(routerSet.bridgeModuleName).toEqual('SynapseBridge')
    })
  })

  describe('getEstimatedTime', () => {
    it('Returns the correct estimated time for all supported chains', () => {
      SUPPORTED_CHAIN_IDS.forEach((chainId) => {
        expect(routerSet.getEstimatedTime(Number(chainId))).toEqual(
          MEDIAN_TIME_BRIDGE[chainId as keyof typeof MEDIAN_TIME_BRIDGE]
        )
      })
    })

    it('Throws error for unsupported chain', () => {
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
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toEqual(routerSet.routers[SupportedChainId.ETH])
    })

    it('Returns undefined given incorrect address', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.ETH,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
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
  })

  describe('getSynapseRouter', () => {
    it('Returns the correct router for supported chain', () => {
      const router = routerSet.getSynapseRouter(SupportedChainId.ETH)
      expect(router).toEqual(routerSet.routers[SupportedChainId.ETH])
      expect(router).toBeInstanceOf(SynapseRouter)
    })

    it('Throws error for unsupported chain', () => {
      expect(() =>
        routerSet.getSynapseRouter(SupportedChainId.AVALANCHE)
      ).toThrow('No module found for chain 43114')
    })
  })

  describe('applySlippage', () => {
    const originQuery: RouterQuery = {
      swapAdapter: '1',
      tokenOut: '2',
      minAmountOut: parseFixed('1000', 18),
      deadline: BigNumber.from(3),
      rawParams: '4',
    }

    const destQuery: RouterQuery = {
      swapAdapter: '5',
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
      const originQueryPlusOne: RouterQuery = {
        ...originQuery,
        minAmountOut: originQuery.minAmountOut.add(1),
      }
      const destQueryPlusOne: RouterQuery = {
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
