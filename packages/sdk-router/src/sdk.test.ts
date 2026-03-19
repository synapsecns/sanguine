import { Provider } from '@ethersproject/abstract-provider'
import { parseFixed } from '@ethersproject/bignumber'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber, PopulatedTransaction } from 'ethers'
import { mock } from 'jest-mock-extended'

import {
  MEDIAN_TIME_BLOCK,
  ROUTER_ADDRESS_MAP,
  SupportedChainId,
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
} from './constants'
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
import { SynapseSDK } from './sdk'
import { EngineID } from './swap'
import { BridgeQuoteV2, IntentQuote, SwapQuote } from './types'
import { ETH_NATIVE_TOKEN_ADDRESS } from './utils'

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

  const opProvider: Provider = getTestProvider(SupportedChainId.OPTIMISM)

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

    it('Saves providers', () => {
      expect(synapse.providers[SupportedChainId.ETH]).toBe(ethProvider)
      expect(synapse.providers[SupportedChainId.ARBITRUM]).toBe(arbProvider)
      expect(synapse.providers[SupportedChainId.BSC]).toBe(bscProvider)
    })

    it('Registers the SynapseBridgeAdapter module set', () => {
      expect(synapse.synapseBridgeAdapterModuleSet).toBeDefined()
      expect(
        synapse.allModuleSets.some(
          (moduleSet) => moduleSet.moduleName === 'SynapseBridge'
        )
      ).toBe(true)
    })

    it('builds the shared intent path for DFK, Harmony, and Klaytn', () => {
      const dfkProvider = mock<Provider>()
      const harmonyProvider = mock<Provider>()
      const klaytnProvider = mock<Provider>()

      const intentsSynapse = new SynapseSDK(
        [
          SupportedChainId.DFK,
          SupportedChainId.HARMONY,
          SupportedChainId.KLAYTN,
        ],
        [dfkProvider, harmonyProvider, klaytnProvider]
      )

      expect(
        intentsSynapse.synapseBridgeAdapterModuleSet.modules[
          SupportedChainId.DFK
        ]
      ).toBeDefined()
      expect(
        intentsSynapse.synapseBridgeAdapterModuleSet.modules[
          SupportedChainId.HARMONY
        ]
      ).toBeDefined()
      expect(
        intentsSynapse.synapseBridgeAdapterModuleSet.modules[
          SupportedChainId.KLAYTN
        ]
      ).toBeDefined()
      expect(() =>
        intentsSynapse.swapEngineSet.getTokenZap(SupportedChainId.DFK)
      ).not.toThrow()
      expect(
        intentsSynapse.sirSet.getSirAddress(SupportedChainId.HARMONY)
      ).toEqual(SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.HARMONY])
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

  describe('Bridge Tx Status', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.OPTIMISM],
      [ethProvider, opProvider]
    )
    const sbaTxHash = '0x1234'

    afterEach(() => {
      jest.restoreAllMocks()
    })

    describe('getSynapseTxId', () => {
      describe('SynapseBridge', () => {
        it('passes through the SBA origin tx hash', async () => {
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseBridge',
              sbaTxHash
            )
          ).resolves.toEqual(sbaTxHash)
        })
      })

      describe('SynapseCCTP', () => {
        it('throws when the module set is not registered', async () => {
          await expect(
            synapse.getSynapseTxId(
              SupportedChainId.ETH,
              'SynapseCCTP',
              sbaTxHash
            )
          ).rejects.toThrow('Unknown bridge module')
        })
      })

      it('Throws when bridge module name is invalid', async () => {
        await expect(
          synapse.getSynapseTxId(
            SupportedChainId.ETH,
            'SynapseSynapse',
            sbaTxHash
          )
        ).rejects.toThrow('Unknown bridge module')
      })
    })

    describe('getBridgeTxStatus', () => {
      describe('SynapseBridge', () => {
        it('delegates status checks to the SBA module set', async () => {
          jest
            .spyOn(
              synapse.synapseBridgeAdapterModuleSet.modules[
                SupportedChainId.OPTIMISM
              ],
              'getBridgeTxStatus'
            )
            .mockResolvedValue(true)

          await expect(
            synapse.getBridgeTxStatus(
              SupportedChainId.OPTIMISM,
              'SynapseBridge',
              sbaTxHash
            )
          ).resolves.toBe(true)
        })

        it('Returns false when unknown synapseTxId', async () => {
          await expect(
            synapse.getBridgeTxStatus(
              SupportedChainId.ETH,
              'SynapseBridge',
              sbaTxHash
            )
          ).resolves.toBe(false)
        })
      })

      describe('SynapseCCTP', () => {
        it('throws when the module set is not registered', async () => {
          await expect(
            synapse.getBridgeTxStatus(
              SupportedChainId.OPTIMISM,
              'SynapseCCTP',
              sbaTxHash
            )
          ).rejects.toThrow('Unknown bridge module')
        })
      })

      it('Throws when bridge module name is invalid', async () => {
        await expect(
          synapse.getBridgeTxStatus(
            SupportedChainId.ETH,
            'SynapseSynapse',
            sbaTxHash
          )
        ).rejects.toThrow('Unknown bridge module')
      })
    })
  })

  describe('getBridgeModuleName', () => {
    const synapse = new SynapseSDK([], [])

    describe('SynapseBridge SBA events', () => {
      ;['TokenSent', 'TokenReceived'].forEach((contractEvent) => {
        it(contractEvent, () => {
          expect(synapse.getBridgeModuleName(contractEvent)).toEqual(
            'SynapseBridge'
          )
          expect(synapse.getBridgeModuleName(`${contractEvent}Event`)).toEqual(
            'SynapseBridge'
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

  describe('SynapseBridgeAdapter V2 integration', () => {
    const mockHarmonyProvider = mock<Provider>()
    const mockBaseProvider = mock<Provider>()
    const mockDfkProvider = mock<Provider>()
    const mockKlaytnProvider = mock<Provider>()
    const sbaDirectOriginToken = '0xE55e19Fb4F2D85af758950957714292DAC1e25B2'
    const sbaDirectRemoteToken = '0x432036208d2717394d2614d6697c46DF3Ed69540'
    const sbaNativeBridgeToken = sbaDirectOriginToken
    const sbaNativeRemoteToken = sbaDirectRemoteToken
    const sbaWrappedNativeOriginToken =
      '0x97855Ba65aa7ed2F65Ed832a776537268158B78a'
    const sbaWrappedNativeRemoteToken =
      '0x5819b6af194A78511c79C85Ea68D2377a7e9335f'
    const sbaFinalToken = '0x00000000000000000000000000000000000000c1'
    const sender = '0x0000000000000000000000000000000000000f01'
    const recipient = '0x0000000000000000000000000000000000000f02'

    const createNoOpRoute = (chainId: number, token: string) => ({
      engineID: EngineID.NoOp,
      engineName: EngineID[EngineID.NoOp],
      chainId,
      fromToken: token,
      fromAmount: BigNumber.from(1000),
      toToken: token,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(1000),
      steps: [],
    })

    const destinationSwapRoute = {
      engineID: EngineID.DefaultPools,
      engineName: EngineID[EngineID.DefaultPools],
      chainId: SupportedChainId.BASE,
      fromToken: sbaDirectRemoteToken,
      fromAmount: BigNumber.from(1000),
      toToken: sbaFinalToken,
      expectedToAmount: BigNumber.from(900),
      minToAmount: BigNumber.from(880),
      steps: [
        {
          token: sbaDirectRemoteToken,
          amount: BigNumber.from(1000),
          msgValue: Zero,
          zapData: '0x1234',
        },
      ],
    }

    const setupSynapse = () => {
      const synapse = new SynapseSDK(
        [SupportedChainId.HARMONY, SupportedChainId.BASE],
        [mockHarmonyProvider, mockBaseProvider]
      )
      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]
      jest
        .spyOn(synapse.synapseBridgeAdapterModuleSet, 'getGasDropAmount')
        .mockResolvedValue(Zero)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[
            SupportedChainId.HARMONY
          ],
          'getNativeFee'
        )
        .mockResolvedValue(BigNumber.from(77))
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[
            SupportedChainId.HARMONY
          ],
          'getEstimatedTime'
        )
        .mockResolvedValue(42)
      jest
        .spyOn(synapse.swapEngineSet, 'getBestQuote')
        .mockImplementation(async (input) => {
          if (
            input.chainId === SupportedChainId.HARMONY &&
            input.fromToken.toLowerCase() ===
              sbaDirectOriginToken.toLowerCase() &&
            input.toToken.toLowerCase() === sbaDirectOriginToken.toLowerCase()
          ) {
            return createNoOpRoute(
              SupportedChainId.HARMONY,
              sbaDirectOriginToken
            ) as any
          }
          if (
            input.chainId === SupportedChainId.HARMONY &&
            input.fromToken.toLowerCase() ===
              sbaNativeBridgeToken.toLowerCase() &&
            input.toToken.toLowerCase() === sbaNativeBridgeToken.toLowerCase()
          ) {
            return createNoOpRoute(
              SupportedChainId.HARMONY,
              sbaNativeBridgeToken
            ) as any
          }
          if (
            input.chainId === SupportedChainId.HARMONY &&
            input.fromToken === ETH_NATIVE_TOKEN_ADDRESS &&
            input.toToken.toLowerCase() === sbaNativeBridgeToken.toLowerCase()
          ) {
            return {
              ...createNoOpRoute(
                SupportedChainId.HARMONY,
                sbaNativeBridgeToken
              ),
              engineID: EngineID.DefaultPools,
              engineName: EngineID[EngineID.DefaultPools],
              fromToken: ETH_NATIVE_TOKEN_ADDRESS,
              expectedToAmount: BigNumber.from(1000),
              minToAmount: BigNumber.from(950),
              steps: [
                {
                  token: ETH_NATIVE_TOKEN_ADDRESS,
                  amount: BigNumber.from(1000),
                  msgValue: BigNumber.from(1000),
                  zapData: '0x1234',
                },
              ],
            } as any
          }
          if (
            input.chainId === SupportedChainId.BASE &&
            input.fromToken.toLowerCase() ===
              sbaDirectRemoteToken.toLowerCase() &&
            input.toToken.toLowerCase() === sbaFinalToken.toLowerCase()
          ) {
            return destinationSwapRoute as any
          }
          return undefined as any
        })
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as any)
      jest
        .spyOn(synapse.tokenMetadataFetcher, 'getTokenDecimals')
        .mockResolvedValue(18)
      return synapse
    }

    afterEach(() => {
      jest.restoreAllMocks()
    })

    it('delegates getSynapseTxId to the SBA module set', async () => {
      const synapse = setupSynapse()
      const txHash = '0x1234'

      await expect(
        synapse.getSynapseTxId(
          SupportedChainId.HARMONY,
          'SynapseBridge',
          txHash
        )
      ).resolves.toEqual(txHash)
    })

    it('delegates getBridgeTxStatus to the SBA module set', async () => {
      const synapse = setupSynapse()
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.BASE],
          'getBridgeTxStatus'
        )
        .mockResolvedValue(true)

      await expect(
        synapse.getBridgeTxStatus(
          SupportedChainId.BASE,
          'SynapseBridge',
          '0x1234'
        )
      ).resolves.toBe(true)
    })

    it('returns direct SBA bridgeV2 quotes without origin swap module names', async () => {
      const synapse = setupSynapse()

      const quotes: BridgeQuoteV2[] = await synapse.bridgeV2({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.BASE,
        fromToken: sbaDirectOriginToken,
        toToken: sbaDirectRemoteToken,
        fromAmount: '1000',
        fromSender: sender,
        toRecipient: recipient,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0]).toMatchObject({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.BASE,
        fromToken: sbaDirectOriginToken,
        toToken: sbaDirectRemoteToken,
        expectedToAmount: '1000',
        minToAmount: '1000',
        routerAddress:
          SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.HARMONY],
        moduleNames: ['SynapseBridge'],
        nativeFee: '77',
        gasDropAmount: '0',
      })
      expect(quotes[0].tx).toBeDefined()
    })

    it('returns no SBA quotes when either bridge chain is outside the temporary allowlist', async () => {
      const synapse = new SynapseSDK(
        [SupportedChainId.ETH, SupportedChainId.BASE],
        [mock<Provider>(), mock<Provider>()]
      )

      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]

      await expect(
        synapse.bridgeV2({
          fromChainId: SupportedChainId.ETH,
          toChainId: SupportedChainId.BASE,
          fromToken: '0x0f2D719407FdBeFF09D87557AbB7232601FD9F29',
          toToken: sbaDirectRemoteToken,
          fromAmount: '1000',
        })
      ).resolves.toEqual([])
    })

    it('supports native-origin bridgeV2 quotes through the generic swap path', async () => {
      const synapse = setupSynapse()

      const quotes: BridgeQuoteV2[] = await synapse.bridgeV2({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.BASE,
        fromToken: ETH_NATIVE_TOKEN_ADDRESS,
        toToken: sbaNativeRemoteToken,
        fromAmount: '1000',
        fromSender: sender,
        toRecipient: recipient,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0]).toMatchObject({
        fromToken: ETH_NATIVE_TOKEN_ADDRESS,
        toToken: sbaNativeRemoteToken,
        expectedToAmount: '1000',
        minToAmount: '950',
        moduleNames: [EngineID[EngineID.DefaultPools], 'SynapseBridge'],
        routerAddress:
          SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.HARMONY],
        nativeFee: '77',
      })
      expect(quotes[0].tx).toBeDefined()
    })

    it('surfaces native output when SBA unwraps a wrapped-native destination token', async () => {
      const synapse = new SynapseSDK(
        [SupportedChainId.DFK, SupportedChainId.KLAYTN],
        [mockDfkProvider, mockKlaytnProvider]
      )

      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]
      jest
        .spyOn(synapse.synapseBridgeAdapterModuleSet, 'getGasDropAmount')
        .mockResolvedValue(Zero)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.DFK],
          'getNativeFee'
        )
        .mockResolvedValue(BigNumber.from(11))
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.DFK],
          'getEstimatedTime'
        )
        .mockResolvedValue(33)
      jest
        .spyOn(synapse.swapEngineSet, 'getBestQuote')
        .mockImplementation(async (input) => {
          if (
            input.chainId === SupportedChainId.DFK &&
            input.fromToken.toLowerCase() ===
              sbaWrappedNativeOriginToken.toLowerCase() &&
            input.toToken.toLowerCase() ===
              sbaWrappedNativeOriginToken.toLowerCase()
          ) {
            return createNoOpRoute(
              SupportedChainId.DFK,
              sbaWrappedNativeOriginToken
            ) as any
          }
          return undefined as any
        })
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as any)

      const nativeQuotes = await synapse.bridgeV2({
        fromChainId: SupportedChainId.DFK,
        toChainId: SupportedChainId.KLAYTN,
        fromToken: sbaWrappedNativeOriginToken,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
        fromAmount: '1000',
        fromSender: sender,
        toRecipient: recipient,
      })

      expect(nativeQuotes).toHaveLength(1)
      expect(nativeQuotes[0]).toMatchObject({
        fromToken: sbaWrappedNativeOriginToken,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
        expectedToAmount: '1000',
        minToAmount: '1000',
        moduleNames: ['SynapseBridge'],
      })

      await expect(
        synapse.bridgeV2({
          fromChainId: SupportedChainId.DFK,
          toChainId: SupportedChainId.KLAYTN,
          fromToken: sbaWrappedNativeOriginToken,
          toToken: sbaWrappedNativeRemoteToken,
          fromAmount: '1000',
        })
      ).resolves.toEqual([])
    })

    it('uses SBA as the bridge step inside multi-tx intents', async () => {
      const synapse = setupSynapse()

      const quotes: IntentQuote[] = await synapse.intent({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.BASE,
        fromToken: sbaDirectOriginToken,
        toToken: sbaFinalToken,
        fromAmount: '1000',
        fromSender: sender,
        toRecipient: recipient,
        allowMultipleTxs: true,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0].toToken).toEqual(sbaFinalToken)
      expect(quotes[0].steps).toHaveLength(2)
      expect(quotes[0].steps[0]).toMatchObject({
        toToken: sbaDirectRemoteToken,
        moduleNames: ['SynapseBridge'],
        nativeFee: '77',
      })
      expect(quotes[0].steps[1]).toMatchObject({
        toToken: sbaFinalToken,
        moduleNames: [EngineID[EngineID.DefaultPools]],
      })
      expect(quotes[0].steps[0].tx).toBeDefined()
      expect(quotes[0].steps[1].tx).toBeDefined()
    })

    it('supports Harmony as an SBA origin chain', async () => {
      const harmonyProvider = mock<Provider>()
      const baseProvider = mock<Provider>()
      const harmonyToken = '0xE55e19Fb4F2D85af758950957714292DAC1e25B2'
      const baseToken = '0x432036208d2717394d2614d6697c46DF3Ed69540'
      const harmonySender = '0x0000000000000000000000000000000000000f11'
      const harmonyRecipient = '0x0000000000000000000000000000000000000f12'
      const synapse = new SynapseSDK(
        [SupportedChainId.HARMONY, SupportedChainId.BASE],
        [harmonyProvider, baseProvider]
      )

      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]
      jest
        .spyOn(synapse.synapseBridgeAdapterModuleSet, 'getGasDropAmount')
        .mockResolvedValue(Zero)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[
            SupportedChainId.HARMONY
          ],
          'getNativeFee'
        )
        .mockResolvedValue(BigNumber.from(11))
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[
            SupportedChainId.HARMONY
          ],
          'getEstimatedTime'
        )
        .mockResolvedValue(33)
      jest
        .spyOn(synapse.swapEngineSet, 'getBestQuote')
        .mockImplementation(async (input) => {
          if (
            input.chainId === SupportedChainId.HARMONY &&
            input.fromToken.toLowerCase() === harmonyToken.toLowerCase() &&
            input.toToken.toLowerCase() === harmonyToken.toLowerCase()
          ) {
            return createNoOpRoute(
              SupportedChainId.HARMONY,
              harmonyToken
            ) as any
          }
          return undefined as any
        })
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as any)

      const quotes = await synapse.bridgeV2({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.BASE,
        fromToken: harmonyToken,
        toToken: baseToken,
        fromAmount: '1000',
        fromSender: harmonySender,
        toRecipient: harmonyRecipient,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0].moduleNames).toEqual(['SynapseBridge'])
      expect(quotes[0].routerAddress).toEqual(
        SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.HARMONY]
      )
    })

    it('supports Harmony as an SBA destination chain', async () => {
      const ethProviderMock = mock<Provider>()
      const harmonyProvider = mock<Provider>()
      const baseToken = '0x432036208d2717394d2614d6697c46DF3Ed69540'
      const harmonyToken = '0xE55e19Fb4F2D85af758950957714292DAC1e25B2'
      const synapse = new SynapseSDK(
        [SupportedChainId.BASE, SupportedChainId.HARMONY],
        [ethProviderMock, harmonyProvider]
      )

      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]
      jest
        .spyOn(synapse.synapseBridgeAdapterModuleSet, 'getGasDropAmount')
        .mockResolvedValue(Zero)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.BASE],
          'getNativeFee'
        )
        .mockResolvedValue(BigNumber.from(22))
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.BASE],
          'getEstimatedTime'
        )
        .mockResolvedValue(44)
      jest
        .spyOn(synapse.swapEngineSet, 'getBestQuote')
        .mockImplementation(async (input) => {
          if (
            input.chainId === SupportedChainId.BASE &&
            input.fromToken.toLowerCase() === baseToken.toLowerCase() &&
            input.toToken.toLowerCase() === baseToken.toLowerCase()
          ) {
            return createNoOpRoute(SupportedChainId.BASE, baseToken) as any
          }
          return undefined as any
        })
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as any)

      const quotes = await synapse.bridgeV2({
        fromChainId: SupportedChainId.BASE,
        toChainId: SupportedChainId.HARMONY,
        fromToken: baseToken,
        toToken: harmonyToken,
        fromAmount: '1000',
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0].moduleNames).toEqual(['SynapseBridge'])
      expect(quotes[0].toToken).toEqual(harmonyToken)
    })
  })

  describe('getEstimatedTime', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.BSC],
      [ethProvider, bscProvider]
    )
    const expectedEthSbaEta = Math.ceil(
      67 * MEDIAN_TIME_BLOCK[SupportedChainId.ETH]
    )
    const expectedBscSbaEta = Math.ceil(
      103 * MEDIAN_TIME_BLOCK[SupportedChainId.BSC]
    )

    describe('Chain with a provider', () => {
      it('Returns estimated time for SynapseBridge', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseBridge')
        ).toEqual(expectedEthSbaEta)

        expect(
          synapse.getEstimatedTime(SupportedChainId.BSC, 'SynapseBridge')
        ).toEqual(expectedBscSbaEta)
      })

      it('Throws when the bridge module is not registered', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseCCTP')
        ).toThrow('Unknown bridge module')
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
        ).toEqual(expectedBscSbaEta)
      })

      it('Throws when the bridge module is not registered', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.ARBITRUM, 'SynapseCCTP')
        ).toThrow('Unknown bridge module')
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
