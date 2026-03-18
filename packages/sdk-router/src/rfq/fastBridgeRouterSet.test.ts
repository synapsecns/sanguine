import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, parseFixed } from '@ethersproject/bignumber'

import { SupportedChainId } from '../constants'
import { getTestProvider } from '../constants/testProviders'
import { BridgeTokenCandidate, CCTPRouterQuery } from '../module'
import { ChainProvider } from '../router'
import { createSlippageTests } from '../router/synapseCCTPRouterSet.test'
import { decodeZapData, EngineID, SwapEngineRoute } from '../swap'
import { FastBridgeRouter } from './fastBridgeRouter'
import { FastBridgeRouterSet } from './fastBridgeRouterSet'
import { FastBridgeQuote } from './quote'

describe('FastBridgeRouterSet', () => {
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

  const routerSet = new FastBridgeRouterSet(testProviders)

  afterEach(() => {
    jest.restoreAllMocks()
  })

  describe('getBridgeRoutes', () => {
    it('returns an empty array for legacy bridge quotes', async () => {
      await expect(routerSet.getBridgeRoutes()).resolves.toEqual([])
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
      minAmountOut: parseFixed('960', 18),
      deadline: BigNumber.from(8),
      rawParams: '9',
    }

    describe('0% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        originQuery.minAmountOut,
        destQuery.minAmountOut,
        0,
        10000
      )
    })

    describe('0.1% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        originQuery.minAmountOut,
        destQuery.minAmountOut,
        10,
        10000
      )
    })

    describe('1% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        originQuery.minAmountOut,
        destQuery.minAmountOut,
        100,
        10000
      )
    })
  })

  describe('applyProtocolFeeRate', () => {
    const amount = BigNumber.from(1_000_001)

    it('Applies 0 bps fee rate', () => {
      const protocolFeeRate = BigNumber.from(0)
      const result = routerSet.applyProtocolFeeRate(amount, protocolFeeRate)
      expect(result).toEqual(amount)
    })

    it('Applies 10 bps fee rate', () => {
      const protocolFeeRate = BigNumber.from(1_000)
      const result = routerSet.applyProtocolFeeRate(amount, protocolFeeRate)
      expect(result).toEqual(BigNumber.from(999_001))
    })
  })

  describe('getBridgeRouteV2', () => {
    const nowMs = 1_700_000_000_000
    const bridgeToken: BridgeTokenCandidate = {
      originChainId: SupportedChainId.ETH,
      destChainId: SupportedChainId.ARBITRUM,
      originToken: '0x00000000000000000000000000000000000000a1',
      destToken: '0x00000000000000000000000000000000000000b1',
    }
    const originSwapRoute: SwapEngineRoute = {
      engineID: EngineID.DefaultPools,
      engineName: EngineID[EngineID.DefaultPools],
      chainId: SupportedChainId.ETH,
      fromToken: '0x00000000000000000000000000000000000000c1',
      fromAmount: BigNumber.from(1_000),
      toToken: bridgeToken.originToken,
      expectedToAmount: BigNumber.from(1_000),
      minToAmount: BigNumber.from(1_000),
      steps: [],
    }
    const quote: FastBridgeQuote = {
      ticker: {
        originToken: {
          chainId: SupportedChainId.ETH,
          token: bridgeToken.originToken,
        },
        destToken: {
          chainId: SupportedChainId.ARBITRUM,
          token: bridgeToken.destToken,
        },
      },
      destAmount: BigNumber.from(1_500),
      maxOriginAmount: BigNumber.from(1_000),
      fixedFee: BigNumber.from(0),
      originFastBridge: '0x00000000000000000000000000000000000000d1',
      destFastBridge: '0x00000000000000000000000000000000000000d2',
      relayerAddr: '0x00000000000000000000000000000000000000d3',
      updatedAt: nowMs,
    }

    it('preserves the FastBridge zap deadline for direct V2 routes', async () => {
      jest.spyOn(Date, 'now').mockReturnValue(nowMs)
      jest.spyOn(routerSet as never, 'getQuotes' as never).mockResolvedValue([
        quote,
      ] as never)

      const populateBridge = jest.fn(async (params: unknown) => ({
        data: FastBridgeRouter.fastBridgeInterface.encodeFunctionData('bridge', [
          params,
        ]),
      }))
      const router = routerSet.getFastBridgeRouter(SupportedChainId.ETH)
      jest.spyOn(router, 'getProtocolFeeRate').mockResolvedValue(BigNumber.from(0))
      jest.spyOn(router, 'getFastBridgeContract').mockResolvedValue({
        address: '0x00000000000000000000000000000000000000f1',
        populateTransaction: {
          bridge: populateBridge,
        },
      } as never)

      const route = await routerSet.getBridgeRouteV2({
        originSwapRoute,
        bridgeToken,
        toToken: bridgeToken.destToken,
        fromSender: '0x0000000000000000000000000000000000000137',
        toRecipient: '0x0000000000000000000000000000000000000138',
      })

      expect(route).toBeDefined()
      expect(route?.expectedToAmount).toEqual(BigNumber.from(1_500))
      expect(route?.minToAmount).toEqual(BigNumber.from(1_500))
      expect(route?.zapData).toBeDefined()
      const decodedZapData = decodeZapData(route!.zapData!)
      expect(decodedZapData.target).toEqual(
        '0x00000000000000000000000000000000000000f1'
      )
      const decodedBridgeCall = FastBridgeRouter.fastBridgeInterface.decodeFunctionData(
        'bridge',
        decodedZapData.payload!
      )
      const bridgeParams = decodedBridgeCall[0] as { deadline: BigNumber }
      expect(bridgeParams.deadline).toEqual(
        BigNumber.from(Math.floor(nowMs / 1000) + 2 * 60 * 60)
      )
      expect(populateBridge).toHaveBeenCalledTimes(1)
    })
  })
})
