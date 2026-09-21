import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, utils } from 'ethers'
import { mock } from 'jest-mock-extended'

import {
  SupportedChainId,
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
} from '../constants'
import { SynapseSDK } from '../sdk'

const ETH_NUSD = '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F'
const AMOUNT = utils.parseEther('100').toString()

describe('SBA origin routes omitted from the original token snapshot', () => {
  beforeEach(() => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => ({ chains: [] }),
    } as Response)
  })

  afterEach(() => jest.restoreAllMocks())

  it.each<[SupportedChainId, string, string]>([
    [
      SupportedChainId.AURORA,
      '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
      ETH_NUSD,
    ],
    [
      SupportedChainId.CRONOS,
      '0x396c9c192dd323995346632581BEF92a31AC623b',
      ETH_NUSD,
    ],
  ])(
    'quotes chain %s to Ethereum with and without a wallet, preserving destination restrictions',
    async (chainId, fromToken, toToken) => {
      const sdk = new SynapseSDK(
        [chainId, SupportedChainId.ETH],
        [mock<Provider>(), mock<Provider>()]
      )
      sdk.allModuleSets = [sdk.synapseBridgeAdapterModuleSet]
      const module = sdk.synapseBridgeAdapterModuleSet.modules[chainId]
      jest.spyOn(module, 'getNativeFee').mockResolvedValue(BigNumber.from(77))
      jest.spyOn(module, 'getEstimatedTime').mockResolvedValue(120)

      const quotes = await sdk.bridgeV2({
        fromChainId: chainId,
        toChainId: SupportedChainId.ETH,
        fromToken,
        toToken,
        fromAmount: AMOUNT,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0]).toMatchObject({
        expectedToAmount: AMOUNT,
        minToAmount: AMOUNT,
        toToken,
        nativeFee: '77',
        estimatedTime: 120,
        moduleNames: ['SynapseBridge'],
        routerAddress: SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[chainId],
      })
      expect(module.getNativeFee).toHaveBeenCalledWith(30101)

      const sender = '0x00000000000000000000000000000000000000a1'
      const connectedQuotes = await sdk.bridgeV2({
        fromChainId: chainId,
        toChainId: SupportedChainId.ETH,
        fromToken,
        toToken,
        fromAmount: AMOUNT,
        fromSender: sender,
        toRecipient: sender,
      })
      expect(connectedQuotes).toHaveLength(1)
      expect(connectedQuotes[0].tx).toMatchObject({
        to: SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[chainId],
        value: '77',
      })
      expect(connectedQuotes[0].tx?.data).toMatch(/^0x[0-9a-f]+$/i)

      await expect(
        sdk.bridgeV2({
          fromChainId: SupportedChainId.ETH,
          toChainId: chainId,
          fromToken: toToken,
          toToken: fromToken,
          fromAmount: AMOUNT,
        })
      ).resolves.toEqual([])
    }
  )
})
