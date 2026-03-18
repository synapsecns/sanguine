import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'

import { SupportedChainId } from '../constants'
import { getTestProvider } from '../constants/testProviders'
import { BridgeTokenCandidate, Query } from '../module'
import { ChainProvider } from '../router'
import { EngineID, SwapEngineRoute } from '../swap'
import { ETH_NATIVE_TOKEN_ADDRESS } from '../utils'
import * as gasZipApi from './api'
import { GasZipModuleSet } from './gasZipModuleSet'

describe('GasZipModuleSet', () => {
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

  const moduleSet = new GasZipModuleSet(testProviders)

  afterEach(() => {
    jest.restoreAllMocks()
  })

  describe('getBridgeRoutes', () => {
    it('returns an empty array for legacy bridge quotes', async () => {
      await expect(moduleSet.getBridgeRoutes()).resolves.toEqual([])
    })
  })

  describe('getBridgeTokenCandidates', () => {
    it('returns the native token candidate for supported chains', async () => {
      jest
        .spyOn(moduleSet as never, 'getAllChainIds' as never)
        .mockResolvedValue([
          SupportedChainId.ETH,
          SupportedChainId.ARBITRUM,
        ] as never)

      await expect(
        moduleSet.getBridgeTokenCandidates({
          fromChainId: SupportedChainId.ETH,
          toChainId: SupportedChainId.ARBITRUM,
          fromToken: ETH_NATIVE_TOKEN_ADDRESS,
          toToken: ETH_NATIVE_TOKEN_ADDRESS,
        })
      ).resolves.toEqual([
        {
          originChainId: SupportedChainId.ETH,
          destChainId: SupportedChainId.ARBITRUM,
          originToken: ETH_NATIVE_TOKEN_ADDRESS,
          destToken: ETH_NATIVE_TOKEN_ADDRESS,
        },
      ])
    })
  })

  describe('getBridgeRouteV2', () => {
    const bridgeToken: BridgeTokenCandidate = {
      originChainId: SupportedChainId.ETH,
      destChainId: SupportedChainId.ARBITRUM,
      originToken: ETH_NATIVE_TOKEN_ADDRESS,
      destToken: ETH_NATIVE_TOKEN_ADDRESS,
    }
    const originSwapRoute: SwapEngineRoute = {
      engineID: EngineID.DefaultPools,
      engineName: EngineID[EngineID.DefaultPools],
      chainId: SupportedChainId.ETH,
      fromToken: ETH_NATIVE_TOKEN_ADDRESS,
      fromAmount: BigNumber.from(1_000),
      toToken: ETH_NATIVE_TOKEN_ADDRESS,
      expectedToAmount: BigNumber.from(1_000),
      minToAmount: BigNumber.from(1_000),
      steps: [],
    }

    it('returns a V2 route for a supported native-token path', async () => {
      jest.spyOn(gasZipApi, 'getGasZipQuote').mockResolvedValue({
        amountOut: BigNumber.from(123),
        speed: 45,
        usd: 1,
      })
      jest
        .spyOn(moduleSet as never, 'checkBlockHeights' as never)
        .mockResolvedValue(true as never)
      jest
        .spyOn(moduleSet as never, 'getGasZipZapData' as never)
        .mockResolvedValue('0xdeadbeef' as never)

      await expect(
        moduleSet.getBridgeRouteV2({
          originSwapRoute,
          bridgeToken,
          toToken: ETH_NATIVE_TOKEN_ADDRESS,
          toRecipient: '0x0000000000000000000000000000000000000137',
        })
      ).resolves.toEqual({
        bridgeToken,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
        expectedToAmount: BigNumber.from(123),
        minToAmount: BigNumber.from(123),
        nativeFee: BigNumber.from(0),
        estimatedTime: 45,
        zapData: '0xdeadbeef',
      })
    })
  })

  describe('applySlippage', () => {
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

    it('returns the original queries unchanged', () => {
      expect(
        moduleSet.applySlippage(originQuery, destQuery)
      ).toEqual({
        originQuery,
        destQuery,
      })
    })
  })
})
