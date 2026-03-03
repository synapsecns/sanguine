import { Provider } from '@ethersproject/abstract-provider'
import { parseFixed } from '@ethersproject/bignumber'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber, PopulatedTransaction } from 'ethers'

import { ROUTER_ADDRESS_MAP, SupportedChainId } from './constants'
import { getTestProvider } from './constants/testProviders'
import {
  ARB_NETH,
  ARB_NUSD,
  ARB_POOL_ETH_WRAPPER,
  ARB_POOL_NETH,
  ARB_POOL_NUSD,
  ARB_USDC,
  ARB_USDC_E,
  ARB_USDT,
  ARB_WETH,
  AVAX_USDC_E,
  BOBA_USDC,
  BSC_USDC,
  ETH_DAI,
  ETH_POOL_NUSD,
  ETH_USDC,
  ETH_USDT,
  NATIVE_ADDRESS,
} from './constants/testValues'
import { Query, RouterQuery } from './module'
import * as operations from './operations'
import { SynapseSDK } from './sdk'
import { SwapQuote } from './types'

// Override fetch to exclude RFQ from tests
global.fetch = jest.fn(() =>
  Promise.resolve({
    ok: true,
    json: () => Promise.resolve([]),
  })
) as any

// Retry the flaky tests up to 3 times
jest.retryTimes(3)

const expectCorrectPopulatedTransaction = (
  populatedTransaction: PopulatedTransaction,
  expectedValue: BigNumber = Zero
) => {
  expect(populatedTransaction).toBeDefined()
  expect(populatedTransaction.data?.length).toBeGreaterThan(0)
  expect(populatedTransaction.to?.length).toBeGreaterThan(0)
  expect(populatedTransaction.value).toEqual(expectedValue)
}

const createSwapQuoteTests = (
  synapse: SynapseSDK,
  chainId: number,
  token: string,
  amount: BigNumber,
  resultPromise: Promise<SwapQuote>
) => {
  let result: SwapQuote
  beforeAll(async () => {
    result = await resultPromise
  })

  it('Fetches a swap quote', async () => {
    expect(result).toBeDefined()
    expect(result.routerAddress?.length).toBeGreaterThan(0)
    expect(result.maxAmountOut.gt(0)).toBe(true)
    expect(result.query).toBeDefined()
  })

  it('Could be used for swapping', async () => {
    const expectedValue = token === NATIVE_ADDRESS ? amount : Zero
    const data = await synapse.swap(
      chainId,
      '0x0000000000000000000000000000000000001337',
      token,
      amount,
      result.query
    )
    expectCorrectPopulatedTransaction(data, expectedValue)
  })
}

describe('SynapseSDK', () => {
  const ethProvider: Provider = getTestProvider(SupportedChainId.ETH)

  const arbProvider: Provider = getTestProvider(SupportedChainId.ARBITRUM)

  const bscProvider: Provider = getTestProvider(SupportedChainId.BSC)

  describe('#constructor', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM, SupportedChainId.BSC],
      [ethProvider, arbProvider, bscProvider]
    )

    it('fails with unequal amount of chains to providers', () => {
      const chainIds = [SupportedChainId.ETH, SupportedChainId.ARBITRUM]
      const testProviders = [ethProvider]
      expect(() => new SynapseSDK(chainIds, testProviders)).toThrow(
        'Amount of chains and providers does not equal'
      )
    })

    it('Instantiates SynapseRouters for each chain', () => {
      expect(synapse.synapseRouterSet).toBeDefined()
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.ETH]
      ).toBeDefined()
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.ARBITRUM]
      ).toBeDefined()
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.BSC]
      ).toBeDefined()
    })

    it('Does not instantiate SynapseRouters for chains without providers', () => {
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.AVALANCHE]
      ).toBeUndefined()
    })

    it('Instantiates SynapseCCTPRouters for each chain with CCTP', () => {
      expect(synapse.synapseCCTPRouterSet).toBeDefined()
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.ETH]
      ).toBeDefined()
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.ARBITRUM]
      ).toBeDefined()
    })

    it('Does not instantiate SynapseCCTPRouters for chains without CCTP', () => {
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.BSC]
      ).toBeUndefined()
    })

    it('Does not instantiate SynapseCCTPRouters for chains without providers', () => {
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.AVALANCHE]
      ).toBeUndefined()
    })

    it('Instantiates CircleCCTPV2 modules for supported chains with providers', () => {
      expect(synapse.circleCCTPV2ModuleSet).toBeDefined()
      expect(
        synapse.circleCCTPV2ModuleSet.modules[SupportedChainId.ETH]
      ).toBeDefined()
      expect(
        synapse.circleCCTPV2ModuleSet.modules[SupportedChainId.ARBITRUM]
      ).toBeDefined()
    })

    it('Does not instantiate CircleCCTPV2 modules for unsupported chains or chains without providers', () => {
      expect(
        synapse.circleCCTPV2ModuleSet.modules[SupportedChainId.BSC]
      ).toBeUndefined()
      expect(
        synapse.circleCCTPV2ModuleSet.modules[SupportedChainId.AVALANCHE]
      ).toBeUndefined()
    })

    it('Registers CircleCCTPV2 module set by default', () => {
      expect(
        synapse.allModuleSets.some(
          (moduleSet) => moduleSet.moduleName === 'CircleCCTPV2'
        )
      ).toBe(true)
    })

    it('Saves providers', () => {
      expect(synapse.providers[SupportedChainId.ETH]).toBe(ethProvider)
      expect(synapse.providers[SupportedChainId.ARBITRUM]).toBe(arbProvider)
      expect(synapse.providers[SupportedChainId.BSC]).toBe(bscProvider)
    })
  })

  describe('applyBridgeSlippage', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
      [ethProvider, arbProvider]
    )

    const originQuery: Query = {
      routerAdapter: '1',
      tokenOut: '2',
      minAmountOut: BigNumber.from(3),
      deadline: BigNumber.from(4),
      rawParams: '5',
    }
    const destQuery: Query = {
      routerAdapter: '6',
      tokenOut: '7',
      minAmountOut: BigNumber.from(8),
      deadline: BigNumber.from(9),
      rawParams: '10',
    }
    const moduleSet = synapse.fastBridgeRouterSet

    describe(`${moduleSet.moduleName} module`, () => {
      beforeEach(() => {
        jest.spyOn(moduleSet, 'applySlippage').mockImplementation(jest.fn())
      })

      it('Applies slippage', () => {
        synapse.applyBridgeSlippage(
          moduleSet.moduleName,
          originQuery,
          destQuery,
          10,
          100
        )
        expect(moduleSet.applySlippage).toHaveBeenCalledWith(
          originQuery,
          destQuery,
          10,
          100
        )
      })

      it('Uses default denominator of 10000', () => {
        synapse.applyBridgeSlippage(
          moduleSet.moduleName,
          originQuery,
          destQuery,
          10
        )
        expect(moduleSet.applySlippage).toHaveBeenCalledWith(
          originQuery,
          destQuery,
          10,
          10000
        )
      })

      it('Uses default slippage of 10 bips', () => {
        synapse.applyBridgeSlippage(
          moduleSet.moduleName,
          originQuery,
          destQuery
        )
        expect(moduleSet.applySlippage).toHaveBeenCalledWith(
          originQuery,
          destQuery,
          10,
          10000
        )
      })
    })

    it('Throws on unknown bridge module', () => {
      expect(() =>
        synapse.applyBridgeSlippage(
          'UnknownBridgeModule',
          originQuery,
          destQuery,
          10,
          10000
        )
      ).toThrow('Unknown bridge module')
    })
  })

  describe('Errors', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
      [ethProvider, arbProvider]
    )

    const amount = BigNumber.from(10).pow(9)
    const emptyQuery: RouterQuery = {
      swapAdapter: AddressZero,
      tokenOut: ETH_USDC,
      minAmountOut: amount,
      deadline: BigNumber.from(0),
      rawParams: '0x',
    }
    const mockAddress = '0x0000000000000000000000000000000000001337'

    describe('origin == destination', () => {
      it('bridgeQuote throws', async () => {
        await expect(
          synapse.bridgeQuote(
            SupportedChainId.ETH,
            SupportedChainId.ETH,
            ETH_USDC,
            BSC_USDC,
            amount
          )
        ).rejects.toThrow(
          'Origin chainId cannot be equal to destination chainId'
        )
      })

      it('bridge throws', async () => {
        await expect(
          synapse.bridge(
            mockAddress,
            ROUTER_ADDRESS_MAP[SupportedChainId.ETH],
            SupportedChainId.ETH,
            SupportedChainId.ETH,
            ETH_USDC,
            amount,
            emptyQuery,
            emptyQuery
          )
        ).rejects.toThrow(
          'Origin chainId cannot be equal to destination chainId'
        )
      })
    })

    it('bridgeQuote throws on unknown chainId', async () => {
      await expect(
        synapse.bridgeQuote(
          SupportedChainId.ETH,
          SupportedChainId.AVALANCHE,
          ETH_USDC,
          AVAX_USDC_E,
          amount
        )
      ).rejects.toThrow('No route found')
    })

    it('bridgeQuote throws when amount too low', async () => {
      await expect(
        synapse.bridgeQuote(
          SupportedChainId.ETH,
          SupportedChainId.ARBITRUM,
          ETH_USDC,
          ARB_USDC,
          BigNumber.from(10).pow(3)
        )
      ).rejects.toThrow('No route found')
    })

    it('bridge throws when incorrect router address', async () => {
      // Use MockAddress as router address
      await expect(
        synapse.bridge(
          mockAddress,
          mockAddress,
          SupportedChainId.ETH,
          SupportedChainId.BSC,
          ETH_USDC,
          amount,
          emptyQuery,
          emptyQuery
        )
      ).rejects.toThrow('Invalid router address')
    })
  })

  describe('Swap', () => {
    const synapse = new SynapseSDK([SupportedChainId.ARBITRUM], [arbProvider])
    const amount = BigNumber.from(10).pow(9)
    const resultPromise: Promise<SwapQuote> = synapse.swapQuote(
      SupportedChainId.ARBITRUM,
      ARB_USDC,
      ARB_USDC_E,
      amount
    )

    createSwapQuoteTests(
      synapse,
      SupportedChainId.ARBITRUM,
      ARB_USDC,
      amount,
      resultPromise
    )
  })

  describe('bridgeV2', () => {
    it('Includes CircleCCTPV2 in module names when Circle route is selected', async () => {
      const synapse = new SynapseSDK(
        [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
        [ethProvider, arbProvider]
      )
      const amount = BigNumber.from(10).pow(9)
      const bridgeToken = {
        originChainId: SupportedChainId.ETH,
        destChainId: SupportedChainId.ARBITRUM,
        originToken: ETH_USDC,
        destToken: ARB_USDC,
      }
      const originSwapRoute = {
        engineID: 0,
        engineName: 'DefaultPools',
        chainId: SupportedChainId.ETH,
        fromToken: ETH_USDC,
        fromAmount: amount,
        toToken: ETH_USDC,
        expectedToAmount: amount,
        steps: [],
        minToAmount: amount,
      }
      const bridgeRoute = {
        bridgeToken,
        toToken: ARB_USDC,
        expectedToAmount: amount,
        minToAmount: amount,
        nativeFee: Zero,
        zapData: '0x',
      }
      const baseQuote = {
        id: 'test-circle-cctp-v2-quote',
        fromChainId: SupportedChainId.ETH,
        fromToken: ETH_USDC,
        fromAmount: amount.toString(),
        toChainId: SupportedChainId.ARBITRUM,
        toToken: ARB_USDC,
        expectedToAmount: amount.toString(),
        minToAmount: amount.toString(),
        routerAddress: '0x0000000000000000000000000000000000001111',
        estimatedTime: 0,
        moduleNames: [],
        gasDropAmount: '0',
        nativeFee: '0',
      }

      // Keep this test fully offline by forcing bridgeV2 to evaluate CircleCCTPV2 only.
      synapse.allModuleSets = [synapse.circleCCTPV2ModuleSet]

      jest
        .spyOn(synapse.circleCCTPV2ModuleSet, 'getBridgeTokenCandidates')
        .mockResolvedValue([bridgeToken])
      jest
        .spyOn(synapse.circleCCTPV2ModuleSet, 'getBridgeRouteV2')
        .mockResolvedValue(bridgeRoute as any)
      jest
        .spyOn(synapse.swapEngineSet, 'getTokenZap')
        .mockReturnValue('0x0000000000000000000000000000000000001234')
      jest.spyOn(synapse.swapEngineSet, 'getBestQuote').mockResolvedValue({
        engineID: 0,
        engineName: 'DefaultPools',
        chainId: SupportedChainId.ETH,
        fromToken: ETH_USDC,
        fromAmount: amount,
        toToken: ETH_USDC,
        expectedToAmount: amount,
      } as any)
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockResolvedValue(originSwapRoute as any)
      jest
        .spyOn(synapse.sirSet, 'finalizeBridgeRouteV2')
        .mockResolvedValue(baseQuote)

      const quotes = await synapse.bridgeV2({
        fromChainId: SupportedChainId.ETH,
        fromToken: ETH_USDC,
        fromAmount: amount.toString(),
        toChainId: SupportedChainId.ARBITRUM,
        toToken: ARB_USDC,
        slippagePercentage: 1,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0].moduleNames).toContain('CircleCCTPV2')
    })
  })
  describe('Get bridge gas', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
      [ethProvider, arbProvider]
    )

    it('Returns zero for ETH', async () => {
      const gas = await synapse.getBridgeGas(SupportedChainId.ETH)
      expect(gas).toEqual(Zero)
    })

    it('Returns non-zero for ARBITRUM', async () => {
      const gas = await synapse.getBridgeGas(SupportedChainId.ARBITRUM)
      expect(gas.gt(0)).toBe(true)
    })
  })

  // TODO: improve tests
  describe('Pool inspection', () => {
    const synapse = new SynapseSDK([SupportedChainId.ARBITRUM], [arbProvider])

    it('Get all pools', async () => {
      const pools = await synapse.getAllPools(SupportedChainId.ARBITRUM)
      expect(pools).toBeDefined()
      expect(pools.length).toBeGreaterThan(0)
      expect(pools[0]?.tokens?.[0]?.token?.length).toBeGreaterThan(0)
    })

    it('Get pool info', async () => {
      const poolInfo = await synapse.getPoolInfo(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NUSD
      )
      expect(poolInfo).toBeDefined()
      expect(poolInfo.lpToken).toEqual(
        '0xcFd72be67Ee69A0dd7cF0f846Fc0D98C33d60F16'
      )
      expect(poolInfo.tokens).toEqual(BigNumber.from(3))
    })

    it('Get pool tokens', async () => {
      const poolTokens = await synapse.getPoolTokens(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NUSD
      )
      expect(poolTokens).toBeDefined()
      expect(poolTokens.length).toEqual(3)
      expect(poolTokens[0].token).toEqual(ARB_NUSD)
      expect(poolTokens[1].token).toEqual(ARB_USDC_E)
      expect(poolTokens[2].token).toEqual(ARB_USDT)
    })
  })

  describe('calculate add liquidity', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ARBITRUM, SupportedChainId.ETH],
      [arbProvider, ethProvider]
    )

    it('Arbitrum nETH pool', async () => {
      const amounts: Record<string, BigNumber> = {}
      amounts[ARB_NETH] = BigNumber.from(10).pow(18)
      amounts[ARB_WETH] = BigNumber.from(10).pow(18).mul(2)
      const result = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amounts
      )
      expect(result).toBeDefined()
      expect(result.amount.gt(0)).toBe(true)
      expect(result.routerAddress).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Handles lowercase token addresses', async () => {
      const amounts: Record<string, BigNumber> = {}
      amounts[ARB_NETH.toLowerCase()] = BigNumber.from(10).pow(18)
      amounts[ARB_WETH.toLowerCase()] = BigNumber.from(10).pow(18).mul(2)
      const result = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amounts
      )
      expect(result).toBeDefined()
      expect(result.amount.gt(0)).toBe(true)
      expect(result.routerAddress).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Handles uppercase token addresses', async () => {
      const amounts: Record<string, BigNumber> = {}
      amounts[ARB_NETH.toUpperCase()] = BigNumber.from(10).pow(18)
      amounts[ARB_WETH.toLowerCase()] = BigNumber.from(10).pow(18).mul(2)
      const result = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amounts
      )
      expect(result).toBeDefined()
      expect(result.amount.gt(0)).toBe(true)
      expect(result.routerAddress).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Handles single token: first one', async () => {
      const amounts: Record<string, BigNumber> = {}
      amounts[ARB_NETH] = BigNumber.from(10).pow(18)
      const amountsFull: Record<string, BigNumber> = {}
      amountsFull[ARB_NETH] = BigNumber.from(10).pow(18)
      amountsFull[ARB_WETH] = Zero
      const result = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amounts
      )
      const expectedResult = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amountsFull
      )
      expect(result).toBeDefined()
      expect(expectedResult).toBeDefined()
      expect(result).toEqual(expectedResult)
    })

    it('Handles single token: second one', async () => {
      const amounts: Record<string, BigNumber> = {}
      amounts[ARB_WETH] = BigNumber.from(10).pow(18)
      const amountsFull: Record<string, BigNumber> = {}
      amountsFull[ARB_NETH] = Zero
      amountsFull[ARB_WETH] = BigNumber.from(10).pow(18)
      const result = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amounts
      )
      const expectedResult = await synapse.calculateAddLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_NETH,
        amountsFull
      )
      expect(result).toBeDefined()
      expect(expectedResult).toBeDefined()
      expect(result).toEqual(expectedResult)
    })

    it('Ethereum nUSD pool', async () => {
      const amounts: Record<string, BigNumber> = {}
      amounts[ETH_USDC] = BigNumber.from(10).pow(6)
      amounts[ETH_DAI] = BigNumber.from(10).pow(18).mul(2)
      amounts[ETH_USDT] = BigNumber.from(10).pow(6).mul(3)
      const result = await synapse.calculateAddLiquidity(
        SupportedChainId.ETH,
        ETH_POOL_NUSD,
        amounts
      )
      expect(result).toBeDefined()
      expect(result.amount.gt(0)).toBe(true)
      expect(result.routerAddress).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
      )
    })
  })

  describe('calculate remove liquidity', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ARBITRUM, SupportedChainId.ETH],
      [arbProvider, ethProvider]
    )

    it('Arbitrum EthWrapper', async () => {
      const result = await synapse.calculateRemoveLiquidity(
        SupportedChainId.ARBITRUM,
        ARB_POOL_ETH_WRAPPER,
        BigNumber.from(10).pow(18)
      )
      expect(result).toBeDefined()
      expect(result.amounts.length).toEqual(2)
      expect(result.amounts[0].value.gt(0)).toBe(true)
      expect(result.amounts[0].index).toEqual(0)
      expect(result.amounts[1].value.gt(0)).toBe(true)
      expect(result.amounts[1].index).toEqual(1)
      expect(result.routerAddress).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })
  })

  describe('calculate remove liquidity one', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ARBITRUM, SupportedChainId.ETH],
      [arbProvider, ethProvider]
    )

    it('Arbitrum EthWrapper', async () => {
      const poolIndex = 0
      const result = await synapse.calculateRemoveLiquidityOne(
        SupportedChainId.ARBITRUM,
        ARB_POOL_ETH_WRAPPER,
        BigNumber.from(10).pow(18),
        poolIndex
      )
      expect(result).toBeDefined()
      expect(result.amount.value.gt(0)).toBe(true)
      expect(result.amount.index).toEqual(poolIndex)
      expect(result.routerAddress).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })
  })
  describe('Internal functions', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ARBITRUM, SupportedChainId.ETH],
      [arbProvider, ethProvider]
    )
    describe('getModuleSet', () => {
      it('Returns correct set for SynapseBridge', () => {
        const routerSet = operations.getModuleSet.call(synapse, 'SynapseBridge')
        expect(routerSet).toEqual(synapse.synapseRouterSet)
      })

      it('Returns correct set for SynapseCCTP', () => {
        const routerSet = operations.getModuleSet.call(synapse, 'SynapseCCTP')
        expect(routerSet).toEqual(synapse.synapseCCTPRouterSet)
      })

      it('Returns correct set for SynapseRFQ', () => {
        const routerSet = operations.getModuleSet.call(synapse, 'SynapseRFQ')
        expect(routerSet).toEqual(synapse.fastBridgeRouterSet)
      })

      it('Returns correct set for CircleCCTPV2', () => {
        const routerSet = operations.getModuleSet.call(synapse, 'CircleCCTPV2')
        expect(routerSet).toEqual(synapse.circleCCTPV2ModuleSet)
      })

      it('Throws when bridge module name is invalid', () => {
        expect(() =>
          operations.getModuleSet.call(synapse, 'SynapseSynapse')
        ).toThrow('Unknown bridge module')
      })
    })
  })
})

describe('Paused Chain Tests', () => {
  let synapseSDK: SynapseSDK

  beforeEach(() => {
    // Setup SDK with test providers
    const chainIds = [SupportedChainId.ETH, SupportedChainId.BOBA] // Include a paused chain
    const providers = chainIds.map((chainId) => getTestProvider(chainId))
    synapseSDK = new SynapseSDK(chainIds, providers)
  })

  describe('Bridge Quote Generation', () => {
    it('should not find quotes when origin chain is paused', async () => {
      // Try to get quote from paused chain (BOBA)
      await expect(
        synapseSDK.bridgeQuote(
          SupportedChainId.BOBA, // Paused chain as origin
          SupportedChainId.ETH,
          BOBA_USDC, // Example token addresses
          ETH_USDC,
          parseFixed('100', 6)
        )
      ).rejects.toThrow('No route found')
    })

    it('should not find quotes when destination chain is paused', async () => {
      await expect(
        synapseSDK.bridgeQuote(
          SupportedChainId.ETH,
          SupportedChainId.BOBA, // Paused chain as destination
          ETH_USDC,
          BOBA_USDC,
          parseFixed('100', 6)
        )
      ).rejects.toThrow('No route found')
    })

    it('should not find quotes when allBridgeQuotes is called with paused chains', async () => {
      const quotes = await synapseSDK.allBridgeQuotes(
        SupportedChainId.BOBA, // Paused chain
        SupportedChainId.ETH,
        BOBA_USDC,
        ETH_USDC,
        parseFixed('100', 6)
      )

      expect(quotes).toHaveLength(0)
    })
  })
})
