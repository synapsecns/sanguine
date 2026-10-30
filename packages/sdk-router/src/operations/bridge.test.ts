import { Provider } from '@ethersproject/abstract-provider'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'
import { mock } from 'jest-mock-extended'

import {
  SupportedChainId,
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
} from '../constants'
import { BridgeRouteV2 } from '../module'
import { SynapseSDK } from '../sdk'
import { EngineID, SwapEngineRoute } from '../swap'
import { ETH_NATIVE_TOKEN_ADDRESS, logger } from '../utils'

const ARB_NETH = '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e'
const AMOUNT = '100000000000000000'

describe('bridgeV2 candidate failures', () => {
  beforeEach(() => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => ({ chains: [] }),
    } as Response)
    jest.spyOn(logger, 'error').mockImplementation(() => undefined)
  })

  afterEach(() => jest.restoreAllMocks())

  it.each(['route', 'transaction finalization', 'quote finalization'])(
    'preserves the SBA native ETH quote when Gas.zip fails during %s',
    async (failureStage) => {
      const sdk = new SynapseSDK(
        [SupportedChainId.ARBITRUM, SupportedChainId.ETH],
        [mock<Provider>(), mock<Provider>()]
      )
      const gasZip = sdk.gasZipModuleSet
      const sba = sdk.synapseBridgeAdapterModuleSet
      sdk.allModuleSets = [gasZip, sba]
      const sbaOriginModule = sba.modules[SupportedChainId.ARBITRUM]
      jest
        .spyOn(sbaOriginModule, 'getNativeFee')
        .mockResolvedValue(BigNumber.from(77))
      jest.spyOn(sbaOriginModule, 'getEstimatedTime').mockResolvedValue(120)

      const gasZipBridgeToken = {
        originChainId: SupportedChainId.ARBITRUM,
        destChainId: SupportedChainId.ETH,
        originToken: ETH_NATIVE_TOKEN_ADDRESS,
        destToken: ETH_NATIVE_TOKEN_ADDRESS,
      }
      jest
        .spyOn(gasZip, 'getBridgeTokenCandidates')
        .mockResolvedValue([gasZipBridgeToken])
      const gasZipRoute: BridgeRouteV2 = {
        bridgeToken: gasZipBridgeToken,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
        expectedToAmount: BigNumber.from(AMOUNT),
        minToAmount: BigNumber.from(AMOUNT),
        nativeFee: Zero,
      }
      jest.spyOn(gasZip, 'getBridgeRouteV2').mockResolvedValue(gasZipRoute)

      // Origin conversion is available for Gas.zip as well as the direct nETH route.
      jest.spyOn(sdk.swapEngineSet, 'getBestQuote').mockImplementation(
        async (input): Promise<SwapEngineRoute> => ({
          engineID: EngineID.NoOp,
          engineName: 'NoOp',
          chainId: input.chainId,
          fromToken: input.fromToken,
          fromAmount: BigNumber.from(input.fromAmount),
          toToken: input.toToken,
          expectedToAmount: BigNumber.from(input.fromAmount),
          minToAmount: BigNumber.from(input.fromAmount),
          steps: [],
        })
      )
      jest
        .spyOn(sdk.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as SwapEngineRoute)

      const failure = new Error('Gas.zip temporarily disabled')
      if (failureStage === 'route') {
        jest.spyOn(gasZip, 'getBridgeRouteV2').mockRejectedValue(failure)
      } else if (failureStage === 'transaction finalization') {
        const finalize = sdk.sirSet.finalizeBridgeRouteV2.bind(sdk.sirSet)
        jest
          .spyOn(sdk.sirSet, 'finalizeBridgeRouteV2')
          .mockImplementation(async (...args) => {
            if (args[3] === gasZipRoute) {
              throw failure
            }
            return finalize(...args)
          })
      } else {
        jest.spyOn(gasZip, 'finalizeBridgeQuoteV2').mockRejectedValue(failure)
      }

      const quotes = await sdk.bridgeV2({
        fromChainId: SupportedChainId.ARBITRUM,
        toChainId: SupportedChainId.ETH,
        fromToken: ARB_NETH,
        toToken: AddressZero,
        fromAmount: AMOUNT,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0]).toMatchObject({
        fromToken: ARB_NETH,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
        expectedToAmount: AMOUNT,
        minToAmount: AMOUNT,
        nativeFee: '77',
        moduleNames: ['SynapseBridge'],
        routerAddress:
          SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM],
      })
      expect(logger.error).toHaveBeenCalledWith(
        expect.stringContaining('Failed to get Gas.zip bridge quote')
      )
    }
  )
})
