import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, PopulatedTransaction, providers } from 'ethers'
import { AddressZero } from '@ethersproject/constants'

import { SynapseSDK } from './sdk'
import {
  ARB_USDC,
  ARB_USDC_E,
  AVAX_GOHM,
  AVAX_USDC_E,
  BSC_GOHM,
  BSC_USDC,
  CCTP_ROUTER_ADDRESS_MAP,
  ETH_USDC,
  PUBLIC_PROVIDER_URLS,
  ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from './constants'
import { BridgeQuote, FeeConfig, RouterQuery } from './router'

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
})
