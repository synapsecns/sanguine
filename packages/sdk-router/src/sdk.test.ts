import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, PopulatedTransaction } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'
import { parseFixed } from '@ethersproject/bignumber'

import { SynapseSDK } from './sdk'
import {
  ARB_GMX,
  ARB_NETH,
  ARB_NUSD,
  ARB_POOL_ETH_WRAPPER,
  ARB_POOL_NETH,
  ARB_POOL_NUSD,
  ARB_USDC,
  ARB_USDC_E,
  ARB_USDT,
  ARB_WETH,
  AVAX_GMX,
  AVAX_GOHM,
  AVAX_USDC_E,
  BSC_GOHM,
  BSC_USDC,
  ETH_DAI,
  ETH_POOL_NUSD,
  ETH_USDC,
  ETH_USDT,
  NATIVE_ADDRESS,
} from './constants/testValues'
import { getTestProvider } from './constants/testProviders'
import {
  CCTP_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_BRIDGE,
  MEDIAN_TIME_CCTP,
  ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from './constants'
import {
  BridgeQuote,
  FeeConfig,
  Query,
  RouterQuery,
  SwapQuote,
  SynapseModuleSet,
} from './module'
import * as operations from './operations'

// Override fetch to exclude RFQ from tests
global.fetch = jest.fn(() =>
  Promise.resolve({
    json: () => Promise.resolve({}),
  })
) as any

const EXPECTED_GAS_DROP: { [chainId: number]: BigNumber } = {
  [SupportedChainId.ETH]: BigNumber.from(0),
  [SupportedChainId.ARBITRUM]: parseFixed('0.0003', 18),
  [SupportedChainId.BSC]: parseFixed('0.002', 18),
  [SupportedChainId.AVALANCHE]: parseFixed('0.025', 18),
}

const expectCorrectFeeConfig = (feeConfig: FeeConfig) => {
  expect(feeConfig).toBeDefined()
  expect(feeConfig.bridgeFee).toBeGreaterThan(0)
  expect(feeConfig.minFee.gt(0)).toBe(true)
  expect(feeConfig.maxFee.gt(0)).toBe(true)
}

const expectCorrectBridgeQuote = (bridgeQuote: BridgeQuote) => {
  expect(bridgeQuote).toBeDefined()
  expect(bridgeQuote.feeAmount.gt(0)).toBe(true)
  expectCorrectFeeConfig(bridgeQuote.feeConfig)
  expect(bridgeQuote.routerAddress?.length).toBeGreaterThan(0)
  expect(bridgeQuote.maxAmountOut.gt(0)).toBe(true)
  expect(bridgeQuote.originQuery).toBeDefined()
  expect(bridgeQuote.destQuery).toBeDefined()
}

const expectCorrectPopulatedTransaction = (
  populatedTransaction: PopulatedTransaction,
  expectedValue: BigNumber = Zero
) => {
  expect(populatedTransaction).toBeDefined()
  expect(populatedTransaction.data?.length).toBeGreaterThan(0)
  expect(populatedTransaction.to?.length).toBeGreaterThan(0)
  expect(populatedTransaction.value).toEqual(expectedValue)
}

const createBridgeQuoteTests = (
  synapse: SynapseSDK,
  originChainId: number,
  destChainId: number,
  token: string,
  amount: BigNumber,
  resultPromise: Promise<BridgeQuote>
) => {
  let result: BridgeQuote
  beforeAll(async () => {
    result = await resultPromise
  })

  it('Fetches a bridge quote', async () => {
    expectCorrectBridgeQuote(result)
  })

  it('Could be used for bridging', async () => {
    const expectedValue = token === NATIVE_ADDRESS ? amount : Zero
    const data = await synapse.bridge(
      '0x0000000000000000000000000000000000001337',
      result.routerAddress,
      originChainId,
      destChainId,
      token,
      amount,
      result.originQuery,
      result.destQuery
    )
    expectCorrectPopulatedTransaction(data, expectedValue)
  })
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

  const opProvider: Provider = getTestProvider(SupportedChainId.OPTIMISM)

  const avaxProvider: Provider = getTestProvider(SupportedChainId.AVALANCHE)

  const bscProvider: Provider = getTestProvider(SupportedChainId.BSC)

  // Chain where CCTP is unlikely to be deployed
  const moonbeamProvider: Provider = getTestProvider(SupportedChainId.MOONBEAM)

  describe('#constructor', () => {
    const synapse = new SynapseSDK(
      [
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        SupportedChainId.MOONBEAM,
      ],
      [ethProvider, arbProvider, moonbeamProvider]
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
        synapse.synapseRouterSet.routers[SupportedChainId.MOONBEAM]
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
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.MOONBEAM]
      ).toBeUndefined()
    })

    it('Does not instantiate SynapseCCTPRouters for chains without providers', () => {
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.AVALANCHE]
      ).toBeUndefined()
    })

    it('Saves providers', () => {
      expect(synapse.providers[SupportedChainId.ETH]).toBe(ethProvider)
      expect(synapse.providers[SupportedChainId.ARBITRUM]).toBe(arbProvider)
      expect(synapse.providers[SupportedChainId.MOONBEAM]).toBe(
        moonbeamProvider
      )
    })
  })

  describe('Bridging: ETH -> ARB', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
      [ethProvider, arbProvider]
    )

    describe('ETH USDC -> ARB USDC', () => {
      const amount = BigNumber.from(10).pow(9)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        ARB_USDC,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        amount,
        resultPromise
      )
    })

    describe('ETH Native -> ARB Native', () => {
      const amount = BigNumber.from(10).pow(18)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        NATIVE_ADDRESS,
        NATIVE_ADDRESS,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        NATIVE_ADDRESS,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.ETH]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.ARBITRUM]
        )
      })
    })
  })

  describe('Bridging: AVAX -> BSC', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.AVALANCHE, SupportedChainId.BSC],
      [avaxProvider, bscProvider]
    )

    describe('AVAX USDC.e -> BSC USDC', () => {
      const amount = BigNumber.from(10).pow(9)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.AVALANCHE,
        SupportedChainId.BSC,
        AVAX_USDC_E,
        BSC_USDC,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.AVALANCHE,
        SupportedChainId.BSC,
        AVAX_USDC_E,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.AVALANCHE]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.AVALANCHE]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.BSC]
        )
      })
    })

    describe('AVAX gOHM -> BSC gOHM', () => {
      const amount = BigNumber.from(10).pow(21)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.AVALANCHE,
        SupportedChainId.BSC,
        AVAX_GOHM,
        BSC_GOHM,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.AVALANCHE,
        SupportedChainId.BSC,
        AVAX_GOHM,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.AVALANCHE]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.AVALANCHE]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.BSC]
        )
      })
    })
  })

  describe.skip('Bridging: ARB -> ETH', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
      [ethProvider, arbProvider]
    )

    describe('ARB USDC -> ETH USDC (using Bridge)', () => {
      const amount = BigNumber.from(10).pow(12).add(1)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        ARB_USDC,
        ETH_USDC,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        ETH_USDC,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.ARBITRUM]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.ETH]
        )
      })
    })

    describe('ARB USDC -> ETH USDC (using CCTP)', () => {
      const amount = BigNumber.from(10).pow(12)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        ARB_USDC,
        ETH_USDC,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        ETH_USDC,
        amount,
        resultPromise
      )

      it('Fetches a CCTP bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
        )
        // SynapseCCTPRouterQuery has routerAdapter property
        expect(result.originQuery.routerAdapter).toBeDefined()
        // Estimated time must match the SynapseCCTP median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_CCTP[SupportedChainId.ARBITRUM]
        )
        expect(result.bridgeModuleName).toEqual('SynapseCCTP')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.ETH]
        )
      })
    })

    describe('ARB USDT -> ETH USDC (excludeCCTP flag tests)', () => {
      // Use $1000 USDT as amount. SynapseCCTP requires less gas on Ethereum to be completed,
      // when USDC is used as a tokenOut (compared to SynapseBridge route).
      // Therefore we can expect that the min fees would be lower. Meaning for amount this low,
      // we should get a CCTP quote unless we explicitly exclude CCTP.
      const amount = BigNumber.from(10).pow(9)

      describe('excludeCCTP flag omitted', () => {
        const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
          SupportedChainId.ARBITRUM,
          SupportedChainId.ETH,
          ARB_USDT,
          ETH_USDC,
          amount
        )

        createBridgeQuoteTests(
          synapse,
          SupportedChainId.ARBITRUM,
          SupportedChainId.ETH,
          ETH_USDC,
          amount,
          resultPromise
        )

        it('Fetches a CCTP bridge quote', async () => {
          const result = await resultPromise
          expect(result.routerAddress).toEqual(
            CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
          )
          // SynapseCCTPRouterQuery has routerAdapter property
          expect(result.originQuery.routerAdapter).toBeDefined()
          // Estimated time must match the SynapseCCTP median time
          expect(result.estimatedTime).toEqual(
            MEDIAN_TIME_CCTP[SupportedChainId.ARBITRUM]
          )
          expect(result.bridgeModuleName).toEqual('SynapseCCTP')
          expect(result.gasDropAmount).toEqual(
            EXPECTED_GAS_DROP[SupportedChainId.ETH]
          )
        })
      })

      describe('excludeCCTP flag off', () => {
        const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
          SupportedChainId.ARBITRUM,
          SupportedChainId.ETH,
          ARB_USDT,
          ETH_USDC,
          amount,
          undefined,
          false
        )

        createBridgeQuoteTests(
          synapse,
          SupportedChainId.ARBITRUM,
          SupportedChainId.ETH,
          ETH_USDC,
          amount,
          resultPromise
        )

        it('Fetches a CCTP bridge quote', async () => {
          const result = await resultPromise
          expect(result.routerAddress).toEqual(
            CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
          )
          // SynapseCCTPRouterQuery has routerAdapter property
          expect(result.originQuery.routerAdapter).toBeDefined()
          // Estimated time must match the SynapseCCTP median time
          expect(result.estimatedTime).toEqual(
            MEDIAN_TIME_CCTP[SupportedChainId.ARBITRUM]
          )
          expect(result.bridgeModuleName).toEqual('SynapseCCTP')
          expect(result.gasDropAmount).toEqual(
            EXPECTED_GAS_DROP[SupportedChainId.ETH]
          )
        })
      })

      describe('excludeCCTP flag on', () => {
        const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
          SupportedChainId.ARBITRUM,
          SupportedChainId.ETH,
          ARB_USDT,
          ETH_USDC,
          amount,
          undefined,
          true
        )

        createBridgeQuoteTests(
          synapse,
          SupportedChainId.ARBITRUM,
          SupportedChainId.ETH,
          ETH_USDC,
          amount,
          resultPromise
        )

        it('Fetches a Synapse bridge quote', async () => {
          const result = await resultPromise
          expect(result.routerAddress).toEqual(
            ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
          )
          // SynapseRouterQuery has swapAdapter property
          expect(result.originQuery.swapAdapter).toBeDefined()
          // Estimated time must match the SynapseBridge median time
          expect(result.estimatedTime).toEqual(
            MEDIAN_TIME_BRIDGE[SupportedChainId.ARBITRUM]
          )
          expect(result.bridgeModuleName).toEqual('SynapseBridge')
          expect(result.gasDropAmount).toEqual(
            EXPECTED_GAS_DROP[SupportedChainId.ETH]
          )
        })
      })
    })

    describe('ARB Native -> ETH Native', () => {
      const amount = BigNumber.from(10).pow(18)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        NATIVE_ADDRESS,
        NATIVE_ADDRESS,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        NATIVE_ADDRESS,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.ARBITRUM]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.ETH]
        )
      })
    })
  })

  describe('Bridging: BSC -> AVAX', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.AVALANCHE, SupportedChainId.BSC],
      [avaxProvider, bscProvider]
    )

    describe('BSC USDC -> AVAX USDC.e', () => {
      // USDC has 18 decimals on BSC. Don't ask me why.
      const amount = BigNumber.from(10).pow(21)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.BSC,
        SupportedChainId.AVALANCHE,
        BSC_USDC,
        AVAX_USDC_E,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.BSC,
        SupportedChainId.AVALANCHE,
        AVAX_USDC_E,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.BSC]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.BSC]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.AVALANCHE]
        )
      })
    })

    describe('BSC gOHM -> AVAX gOHM', () => {
      const amount = BigNumber.from(10).pow(21)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.BSC,
        SupportedChainId.AVALANCHE,
        BSC_GOHM,
        AVAX_GOHM,
        amount
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.BSC,
        SupportedChainId.AVALANCHE,
        AVAX_GOHM,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        const result = await resultPromise
        expect(result.routerAddress).toEqual(
          ROUTER_ADDRESS_MAP[SupportedChainId.BSC]
        )
        // SynapseRouterQuery has swapAdapter property
        expect(result.originQuery.swapAdapter).toBeDefined()
        // Estimated time must match the SynapseBridge median time
        expect(result.estimatedTime).toEqual(
          MEDIAN_TIME_BRIDGE[SupportedChainId.BSC]
        )
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.AVALANCHE]
        )
      })
    })
  })

  describe('Gas drop edge cases', () => {
    const synapse = new SynapseSDK(
      [
        SupportedChainId.ARBITRUM,
        SupportedChainId.AVALANCHE,
        SupportedChainId.MOONBEAM,
      ],
      [arbProvider, avaxProvider, moonbeamProvider]
    )

    describe('GMX', () => {
      it('ARB -> AVAX: non-zero gas drop', async () => {
        const result = await synapse.bridgeQuote(
          SupportedChainId.ARBITRUM,
          SupportedChainId.AVALANCHE,
          ARB_GMX,
          AVAX_GMX,
          parseFixed('100', 18)
        )
        expect(result).toBeDefined()
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(
          EXPECTED_GAS_DROP[SupportedChainId.AVALANCHE]
        )
      })

      it('AVAX -> ARB: zero gas drop', async () => {
        const result = await synapse.bridgeQuote(
          SupportedChainId.AVALANCHE,
          SupportedChainId.ARBITRUM,
          AVAX_GMX,
          ARB_GMX,
          parseFixed('100', 18)
        )
        expect(result).toBeDefined()
        expect(result.bridgeModuleName).toEqual('SynapseBridge')
        expect(result.gasDropAmount).toEqual(Zero)
      })
    })
  })

  describe('allBridgeQuotes', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.ARBITRUM],
      [ethProvider, arbProvider]
    )

    it('Fetches SynapseBridge and SynapseCCTP quotes for USDC', async () => {
      const allQuotes = await synapse.allBridgeQuotes(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        ARB_USDT,
        BigNumber.from(10).pow(9)
      )
      expect(allQuotes.length).toEqual(2)
      expectCorrectBridgeQuote(allQuotes[0])
      expectCorrectBridgeQuote(allQuotes[1])
      // First quote should have better quote
      expect(allQuotes[0].maxAmountOut.gte(allQuotes[1].maxAmountOut)).toBe(
        true
      )
      // One should be SynapseBridge and the other SynapseCCTP
      expect(allQuotes[0].bridgeModuleName).not.toEqual(
        allQuotes[1].bridgeModuleName
      )
      expect(
        allQuotes[0].bridgeModuleName === 'SynapseBridge' ||
          allQuotes[1].bridgeModuleName === 'SynapseBridge'
      ).toBe(true)
      expect(
        allQuotes[0].bridgeModuleName === 'SynapseCCTP' ||
          allQuotes[1].bridgeModuleName === 'SynapseCCTP'
      ).toBe(true)
      expect(allQuotes[0].gasDropAmount).toEqual(
        EXPECTED_GAS_DROP[SupportedChainId.ARBITRUM]
      )
      expect(allQuotes[1].gasDropAmount).toEqual(
        EXPECTED_GAS_DROP[SupportedChainId.ARBITRUM]
      )
    })

    it('Fetches only SynapseBridge quotes for ETH', async () => {
      const allQuotes = await synapse.allBridgeQuotes(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        NATIVE_ADDRESS,
        NATIVE_ADDRESS,
        BigNumber.from(10).pow(18)
      )
      expect(allQuotes.length).toEqual(1)
      expectCorrectBridgeQuote(allQuotes[0])
      expect(allQuotes[0].bridgeModuleName).toEqual('SynapseBridge')
      expect(allQuotes[0].gasDropAmount).toEqual(
        EXPECTED_GAS_DROP[SupportedChainId.ARBITRUM]
      )
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

    const createApplySlippageTests = (moduleSet: SynapseModuleSet) => {
      describe(`${moduleSet.bridgeModuleName} module`, () => {
        beforeEach(() => {
          jest.spyOn(moduleSet, 'applySlippage').mockImplementation(jest.fn())
        })

        it('Applies slippage', () => {
          synapse.applyBridgeSlippage(
            moduleSet.bridgeModuleName,
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
            moduleSet.bridgeModuleName,
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
            moduleSet.bridgeModuleName,
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
    }

    createApplySlippageTests(synapse.synapseRouterSet)

    createApplySlippageTests(synapse.synapseCCTPRouterSet)

    createApplySlippageTests(synapse.fastBridgeRouterSet)

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
      [SupportedChainId.ETH, SupportedChainId.BSC],
      [ethProvider, bscProvider]
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
          SupportedChainId.BSC,
          ETH_USDC,
          BSC_USDC,
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

  describe('Bridge Tx Status', () => {
    const synapse = new SynapseSDK(
      [
        SupportedChainId.ARBITRUM,
        SupportedChainId.ETH,
        SupportedChainId.OPTIMISM,
      ],
      [arbProvider, ethProvider, opProvider]
    )

    // https://etherscan.io/tx/0xe3f0f0c1d139c48730492c900f9978449d70c0939c654d5abbfd6b191f9c7b3d
    // https://arbiscan.io/tx/0xb13d5c9156e2d88662fa2f252bd2e1d77d768f0de9d27ca60a79e40b493f6ef2
    const bridgeEthToArbTx = {
      txHash:
        '0xe3f0f0c1d139c48730492c900f9978449d70c0939c654d5abbfd6b191f9c7b3d',
      synapseTxId:
        '0x2f223fb1509f04f777b5c9dd2287931b6e63d994a6a697db7a08cfbe784b5e90',
    }

    // https://arbiscan.io/tx/0xe226c7e38e4b83072aa9d947e533be32c8bb38120bbdd8f490c5c6a5894e62c9
    // https://etherscan.io/tx/0xb88feb2a92690448b840851dff41dbc7cdc975c1fb740f0523b5c2e407ac9f38
    const bridgeArbToEthTx = {
      txHash:
        '0xe226c7e38e4b83072aa9d947e533be32c8bb38120bbdd8f490c5c6a5894e62c9',
      synapseTxId:
        '0xf7b8085d96b1ea3f6bf7a07ad93d1861b8fcd551ef56665d6a22c9fb7633a097',
    }

    // https://etherscan.io/tx/0x1a25b0dfde1e2cc43f1dc659ba60f2b8e7ff8177555773fea0c4fba2d6e9c393
    // https://arbiscan.io/tx/0x0166c1e99b0ec8942ed10527cd7ac9003111ee697e0c0519312228e669a61378
    const cctpEthToArbTx = {
      txHash:
        '0x1a25b0dfde1e2cc43f1dc659ba60f2b8e7ff8177555773fea0c4fba2d6e9c393',
      synapseTxId:
        '0x492b923b5a0ace2715a8d0a80fb93c094bf6d35b142a010bdc3761b8613439fc',
    }

    // https://arbiscan.io/tx/0x2a6d04ba5a48331454f00d136b3666869d03f004395fea25d97d42715c119096
    // https://etherscan.io/tx/0xefb946d2acf8343ac5526de66de498e0d5f70ae73c81b833181616ee058a22d7
    const cctpArbToEthTx = {
      txHash:
        '0x2a6d04ba5a48331454f00d136b3666869d03f004395fea25d97d42715c119096',
      synapseTxId:
        '0xed98b02f712c940d3b37a1aa9005a5986ecefa5cdbb4505118a22ae65d4903af',
    }

    // https://optimistic.etherscan.io/tx/0xf8c736aa8f0455853e68fc4c26c251b6264d77e613efe1cde2a8400ec7a9355f
    // https://arbiscan.io/tx/0x93287ed477a034a4b843088bd60affc78be7fa0a199ae1ae399e82ccebea8a43
    const rfqOpToArbTx = {
      txHash:
        '0xf8c736aa8f0455853e68fc4c26c251b6264d77e613efe1cde2a8400ec7a9355f',
      synapseTxId:
        '0xfe7914246c17a5069024168fda5ceb8f31ed1b1c929da7f586b2a415f75fdc5e',
    }

    // https://arbiscan.io/tx/0xc93e4abbbad0e5c6f724928bf42ed9b8ea9c4ac70483c1a00e374a7f002cdb72
    // https://optimistic.etherscan.io/tx/0xbd45074e933e68795e02c5b3b7378f1911972415aac6e00d33e29933f86b5462
    const rfqArbToOpTx = {
      txHash:
        '0xc93e4abbbad0e5c6f724928bf42ed9b8ea9c4ac70483c1a00e374a7f002cdb72',
      synapseTxId:
        '0xd8e5d4b4658beccfa7ccd69d85c84181cd24b1f8f35d88993c033c0f732b1dd3',
    }

    describe('getSynapseTxId', () => {
      describe('SynapseBridge', () => {
        const ethSynBridge = '0x2796317b0fF8538F253012862c06787Adfb8cEb6'
        const events =
          'TokenDeposit, TokenDepositAndSwap, TokenRedeem, TokenRedeemAndRemove, TokenRedeemAndSwap, TokenRedeemV2'

        it('ETH -> ARB', async () => {
          const synapseTxId = await synapse.getSynapseTxId(
            SupportedChainId.ETH,
            'SynapseBridge',
            bridgeEthToArbTx.txHash
          )
          expect(synapseTxId).toEqual(bridgeEthToArbTx.synapseTxId)
        })

        it('ARB -> ETH', async () => {
          const synapseTxId = await synapse.getSynapseTxId(
            SupportedChainId.ARBITRUM,
            'SynapseBridge',
            bridgeArbToEthTx.txHash
          )
          expect(synapseTxId).toEqual(bridgeArbToEthTx.synapseTxId)
        })

        it('Throws when given a txHash that does not exist', async () => {
          // Use txHash for another chain
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseBridge',
              bridgeArbToEthTx.txHash
            )
          ).rejects.toThrow('Failed to get transaction receipt')
        })

        it('Throws when origin tx does not refer to SynapseBridge', async () => {
          const errorMsg =
            `Contract ${ethSynBridge} in transaction ${cctpEthToArbTx.txHash}` +
            ` did not emit any of the expected events: ${events}`
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseBridge',
              cctpEthToArbTx.txHash
            )
          ).rejects.toThrow(errorMsg)
        })

        it('Throws when given a destination tx', async () => {
          // Destination tx hash for ARB -> ETH
          const txHash =
            '0xefb946d2acf8343ac5526de66de498e0d5f70ae73c81b833181616ee058a22d7'
          const errorMsg =
            `Contract ${ethSynBridge} in transaction ${txHash}` +
            ` did not emit any of the expected events: ${events}`
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseBridge',
              txHash
            )
          ).rejects.toThrow(errorMsg)
        })
      })

      describe('SynapseCCTP', () => {
        const ethSynCCTP = '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E'
        const events = 'CircleRequestSent'

        it('ETH -> ARB', async () => {
          const synapseTxId = await synapse.getSynapseTxId(
            SupportedChainId.ETH,
            'SynapseCCTP',
            cctpEthToArbTx.txHash
          )
          expect(synapseTxId).toEqual(cctpEthToArbTx.synapseTxId)
        })

        it('ARB -> ETH', async () => {
          const synapseTxId = await synapse.getSynapseTxId(
            SupportedChainId.ARBITRUM,
            'SynapseCCTP',
            cctpArbToEthTx.txHash
          )
          expect(synapseTxId).toEqual(cctpArbToEthTx.synapseTxId)
        })

        it('Throws when given a txHash that does not exist', async () => {
          // Use txHash for another chain
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseCCTP',
              cctpArbToEthTx.txHash
            )
          ).rejects.toThrow('Failed to get transaction receipt')
        })

        it('Throws when origin tx does not refer to SynapseCCTP', async () => {
          const errorMsg =
            `Contract ${ethSynCCTP} in transaction ${bridgeEthToArbTx.txHash}` +
            ` did not emit any of the expected events: ${events}`
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseCCTP',
              bridgeEthToArbTx.txHash
            )
          ).rejects.toThrow(errorMsg)
        })

        it('Throws when given a destination tx', async () => {
          // Destination tx hash for ARB -> ETH
          const txHash =
            '0xefb946d2acf8343ac5526de66de498e0d5f70ae73c81b833181616ee058a22d7'
          const errorMsg =
            `Contract ${ethSynCCTP} in transaction ${txHash}` +
            ` did not emit any of the expected events: ${events}`
          await expect(
            synapse.getSynapseTxId(SupportedChainId.ETH, 'SynapseCCTP', txHash)
          ).rejects.toThrow(errorMsg)
        })
      })

      describe('SynapseRFQ', () => {
        const arbSynRFQ = '0x6C0771aD91442D670159a8171C35F4828E19aFd2'
        const events = 'BridgeRequested'

        it('OP -> ARB', async () => {
          const synapseTxId = await synapse.getSynapseTxId(
            SupportedChainId.OPTIMISM,
            'SynapseRFQ',
            rfqOpToArbTx.txHash
          )
          expect(synapseTxId).toEqual(rfqOpToArbTx.synapseTxId)
        })

        it('ARB -> OP', async () => {
          const synapseTxId = await synapse.getSynapseTxId(
            SupportedChainId.ARBITRUM,
            'SynapseRFQ',
            rfqArbToOpTx.txHash
          )
          expect(synapseTxId).toEqual(rfqArbToOpTx.synapseTxId)
        })

        it('Throws when given a txHash that does not exist', async () => {
          // Use txHash for another chain
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.OPTIMISM,
              'SynapseRFQ',
              rfqArbToOpTx.txHash
            )
          ).rejects.toThrow('Failed to get transaction receipt')
        })

        it('Throws when origin tx does not refer to SynapseRFQ', async () => {
          const errorMsg =
            `Contract ${arbSynRFQ} in transaction ${bridgeArbToEthTx.txHash}` +
            ` did not emit any of the expected events: ${events}`
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ARBITRUM,
              'SynapseRFQ',
              bridgeArbToEthTx.txHash
            )
          ).rejects.toThrow(errorMsg)
        })

        it('Throws when given a destination tx', async () => {
          // Destination tx hash for OP -> ARB
          const txHash =
            '0x53a8e543bc0e3f0c1cae509e50d9435c3b62073eecf1aee7ece63c3be285db30'
          const errorMsg =
            `Contract ${arbSynRFQ} in transaction ${txHash}` +
            ` did not emit any of the expected events: ${events}`
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ARBITRUM,
              'SynapseRFQ',
              txHash
            )
          ).rejects.toThrow(errorMsg)
        })
      })

      it('Throws when bridge module name is invalid', async () => {
        await expect(
          synapse.getSynapseTxId(
            SupportedChainId.ETH,
            'SynapseSynapse',
            bridgeEthToArbTx.txHash
          )
        ).rejects.toThrow('Unknown bridge module')
      })
    })

    describe('getBridgeTxStatus', () => {
      describe('SynapseBridge', () => {
        it('ETH -> ARB', async () => {
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ARBITRUM,
            'SynapseBridge',
            bridgeEthToArbTx.synapseTxId
          )
          expect(txStatus).toBe(true)
        })

        it('ARB -> ETH', async () => {
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseBridge',
            bridgeArbToEthTx.synapseTxId
          )
          expect(txStatus).toBe(true)
        })

        it('Returns false when unknown synapseTxId', async () => {
          // Using txHash instead of synapseTxId
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseBridge',
            bridgeArbToEthTx.txHash
          )
          expect(txStatus).toBe(false)
        })

        it('Returns false when origin chain is used instead of destination', async () => {
          // First argument should be destination chainId
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseBridge',
            bridgeEthToArbTx.synapseTxId
          )
          expect(txStatus).toBe(false)
        })
      })

      describe('SynapseCCTP', () => {
        it('ETH -> ARB', async () => {
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ARBITRUM,
            'SynapseCCTP',
            cctpEthToArbTx.synapseTxId
          )
          expect(txStatus).toBe(true)
        })

        it('ARB -> ETH', async () => {
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseCCTP',
            cctpArbToEthTx.synapseTxId
          )
          expect(txStatus).toBe(true)
        })

        it('Returns false when unknown synapseTxId', async () => {
          // Using txHash instead of synapseTxId
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseCCTP',
            cctpArbToEthTx.txHash
          )
          expect(txStatus).toBe(false)
        })

        it('Returns false when origin chain is used instead of destination', async () => {
          // First argument should be destination chainId
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseCCTP',
            cctpEthToArbTx.synapseTxId
          )
          expect(txStatus).toBe(false)
        })
      })

      describe('SynapseRFQ', () => {
        it('OP -> ARB', async () => {
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.ARBITRUM,
            'SynapseRFQ',
            rfqOpToArbTx.synapseTxId
          )
          expect(txStatus).toBe(true)
        })

        it('ARB -> OP', async () => {
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.OPTIMISM,
            'SynapseRFQ',
            rfqArbToOpTx.synapseTxId
          )
          expect(txStatus).toBe(true)
        })

        it('Returns false when unknown synapseTxId', async () => {
          // Using txHash instead of synapseTxId
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.OPTIMISM,
            'SynapseRFQ',
            rfqArbToOpTx.txHash
          )
          expect(txStatus).toBe(false)
        })

        it('Returns false when origin chain is used instead of destination', async () => {
          // First argument should be destination chainId
          const txStatus = await synapse.getBridgeTxStatus(
            SupportedChainId.OPTIMISM,
            'SynapseRFQ',
            rfqOpToArbTx.synapseTxId
          )
          expect(txStatus).toBe(false)
        })
      })

      it('Throws when bridge module name is invalid', async () => {
        await expect(
          synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseSynapse',
            bridgeEthToArbTx.txHash
          )
        ).rejects.toThrow('Unknown bridge module')
      })
    })
  })

  describe('getBridgeModuleName', () => {
    const synapse = new SynapseSDK([], [])

    // https://github.com/synapsecns/synapse-contracts/blob/3f592a879baa4487a62ca8d2cfd44d329bc22e62/contracts/bridge/SynapseBridge.sol#L63-L121
    describe('SynapseBridge events', () => {
      const contractEvents = [
        'TokenDeposit',
        'TokenRedeem',
        'TokenWithdraw',
        'TokenMint',
        'TokenDepositAndSwap',
        'TokenMintAndSwap',
        'TokenRedeemAndSwap',
        'TokenRedeemAndRemove',
        'TokenWithdrawAndRemove',
        'TokenRedeemV2',
      ]

      contractEvents.forEach((contractEvent) => {
        it(contractEvent, () => {
          // Event naming in contract and explorer is a bit different
          // schema: TokenDeposit => DepositEvent
          const explorerEvent = `${contractEvent.slice(5)}Event`
          expect(synapse.getBridgeModuleName(explorerEvent)).toEqual(
            'SynapseBridge'
          )
        })
      })
    })

    // https://github.com/synapsecns/synapse-contracts/blob/3f592a879baa4487a62ca8d2cfd44d329bc22e62/contracts/cctp/events/SynapseCCTPEvents.sol#L5-L45
    describe('SynapseCCTP events', () => {
      const contractEvents = ['CircleRequestSent', 'CircleRequestFulfilled']

      contractEvents.forEach((contractEvent) => {
        it(contractEvent, () => {
          // Event naming in contract and explorer is a bit different
          // schema: CircleRequestSent => CircleRequestSentEvent
          const explorerEvent = `${contractEvent}Event`
          expect(synapse.getBridgeModuleName(explorerEvent)).toEqual(
            'SynapseCCTP'
          )
        })
      })
    })

    it('Throws when event name is unknown', () => {
      expect(() => synapse.getBridgeModuleName('SomeUnknownEvent')).toThrow(
        'Unknown event'
      )
    })
  })

  describe('getEstimatedTime', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.MOONBEAM],
      [ethProvider, moonbeamProvider]
    )

    describe('Chain with a provider', () => {
      it('Returns estimated time for SynapseBridge', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseBridge')
        ).toEqual(MEDIAN_TIME_BRIDGE[SupportedChainId.ETH])

        expect(
          synapse.getEstimatedTime(SupportedChainId.MOONBEAM, 'SynapseBridge')
        ).toEqual(MEDIAN_TIME_BRIDGE[SupportedChainId.MOONBEAM])
      })

      it('Returns estimated time for SynapseCCTP', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseCCTP')
        ).toEqual(MEDIAN_TIME_CCTP[SupportedChainId.ETH])
      })

      it('Throws when bridge module does not exist on a chain', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.MOONBEAM, 'SynapseCCTP')
        ).toThrow('No estimated time for chain 1284')
      })

      it('Throws when bridge module name is invalid', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseSynapse')
        ).toThrow('Unknown bridge module')
      })
    })

    describe('Chain without a provider', () => {
      it('Returns estimated time for SynapseBridge', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.BSC, 'SynapseBridge')
        ).toEqual(MEDIAN_TIME_BRIDGE[SupportedChainId.BSC])
      })

      it('Returns estimated time for SynapseCCTP', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.ARBITRUM, 'SynapseCCTP')
        ).toEqual(MEDIAN_TIME_CCTP[SupportedChainId.ARBITRUM])
      })

      it('Throws when bridge module does not exist on a chain', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.DOGECHAIN, 'SynapseCCTP')
        ).toThrow('No estimated time for chain 2000')
      })

      it('Throws when bridge module name is invalid', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.BSC, 'SynapseSynapse')
        ).toThrow('Unknown bridge module')
      })
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

      it('Throws when bridge module name is invalid', () => {
        expect(() =>
          operations.getModuleSet.call(synapse, 'SynapseSynapse')
        ).toThrow('Unknown bridge module')
      })
    })
  })
})
