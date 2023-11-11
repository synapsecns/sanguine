import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, PopulatedTransaction, providers } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'
import { beforeAll, describe, expect, it } from 'vitest'

import { SynapseSDK } from './sdk'
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
  AVAX_GOHM,
  AVAX_USDC_E,
  BSC_GOHM,
  BSC_USDC,
  CCTP_ROUTER_ADDRESS_MAP,
  ETH_DAI,
  ETH_POOL_NUSD,
  ETH_USDC,
  ETH_USDT,
  MEDIAN_TIME_BRIDGE,
  MEDIAN_TIME_CCTP,
  PUBLIC_PROVIDER_URLS,
  ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from './constants'
import { BridgeQuote, FeeConfig, RouterQuery, SwapQuote } from './router'

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
  populatedTransaction: PopulatedTransaction
) => {
  expect(populatedTransaction).toBeDefined()
  expect(populatedTransaction.data?.length).toBeGreaterThan(0)
  expect(populatedTransaction.to?.length).toBeGreaterThan(0)
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
    synapse
      .bridge(
        '0x0000000000000000000000000000000000001337',
        result.routerAddress,
        originChainId,
        destChainId,
        token,
        amount,
        result.originQuery,
        result.destQuery
      )
      .then(expectCorrectPopulatedTransaction)
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
    synapse
      .swap(
        chainId,
        '0x0000000000000000000000000000000000001337',
        token,
        amount,
        result.query
      )
      .then(expectCorrectPopulatedTransaction)
  })
}

describe('SynapseSDK', () => {
  const ethProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ETH]
  )

  const arbProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ARBITRUM]
  )

  const avaxProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.AVALANCHE]
  )

  const bscProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.BSC]
  )

  // Chain where CCTP is unlikely to be deployed
  const moonbeamProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.MOONBEAM]
  )

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

    describe('ETH USDC -> ARB USDC.e (excludeCCTP flag omitted)', () => {
      // Try to find ETH USDC -> ARB USDC.e quote for 1M USDC,
      // which by default is routed through USDC
      const amount = BigNumber.from(10).pow(12)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        ARB_USDC_E,
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

      it('Fetches a CCTP bridge quote', async () => {
        resultPromise.then((result) => {
          expect(result.routerAddress).toEqual(
            CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
          )
          // SynapseCCTPRouterQuery has routerAdapter property
          expect(result.originQuery.routerAdapter).toBeDefined()
          // Estimated time must match the SynapseCCTP median time
          expect(result.estimatedTime).toEqual(
            MEDIAN_TIME_CCTP[SupportedChainId.ETH]
          )
          expect(result.bridgeModuleName).toEqual('SynapseCCTP')
        })
      })
    })

    describe('ETH USDC -> ARB USDC.e (excludeCCTP flag off)', () => {
      // Try to find ETH USDC -> ARB USDC.e quote for 1M USDC,
      // which by default is routed through USDC
      const amount = BigNumber.from(10).pow(12)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        ARB_USDC_E,
        amount,
        undefined,
        false
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        amount,
        resultPromise
      )

      it('Fetches a CCTP bridge quote', async () => {
        resultPromise.then((result) => {
          expect(result.routerAddress).toEqual(
            CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
          )
          // SynapseCCTPRouterQuery has routerAdapter property
          expect(result.originQuery.routerAdapter).toBeDefined()
          // Estimated time must match the SynapseCCTP median time
          expect(result.estimatedTime).toEqual(
            MEDIAN_TIME_CCTP[SupportedChainId.ETH]
          )
          expect(result.bridgeModuleName).toEqual('SynapseCCTP')
        })
      })
    })

    describe('ETH USDC -> ARB USDC.e (excludeCCTP flag on)', () => {
      // Try to find ETH USDC -> ARB USDC.e quote for 1M USDC,
      // which by default is routed through USDC
      const amount = BigNumber.from(10).pow(12)
      const resultPromise: Promise<BridgeQuote> = synapse.bridgeQuote(
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        ARB_USDC_E,
        amount,
        undefined,
        true
      )

      createBridgeQuoteTests(
        synapse,
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        resultPromise.then((result) => {
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
        })
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
        resultPromise.then((result) => {
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
        })
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
        resultPromise.then((result) => {
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
        })
      })
    })
  })

  describe('Bridging: ARB -> ETH', () => {
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
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        resultPromise.then((result) => {
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
        })
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
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        ETH_USDC,
        amount,
        resultPromise
      )

      it('Fetches a CCTP bridge quote', async () => {
        resultPromise.then((result) => {
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
        })
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
        SupportedChainId.AVALANCHE,
        SupportedChainId.BSC,
        AVAX_USDC_E,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        resultPromise.then((result) => {
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
        })
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
        SupportedChainId.AVALANCHE,
        SupportedChainId.BSC,
        AVAX_GOHM,
        amount,
        resultPromise
      )

      it('Fetches a Synapse bridge quote', async () => {
        resultPromise.then((result) => {
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
        })
      })
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
})
