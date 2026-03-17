import { Provider } from '@ethersproject/abstract-provider'
import { parseFixed } from '@ethersproject/bignumber'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber, PopulatedTransaction } from 'ethers'
import { mock } from 'jest-mock-extended'

import {
  MEDIAN_TIME_BRIDGE,
  MEDIAN_TIME_CCTP,
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
          (moduleSet) => moduleSet.moduleName === 'SynapseBridgeAdapter'
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

      describe.skip('SynapseRFQ', () => {
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

      describe.skip('SynapseRFQ', () => {
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

    describe('SynapseBridgeAdapter events', () => {
      ;['TokenSent', 'TokenReceived'].forEach((contractEvent) => {
        it(contractEvent, () => {
          expect(synapse.getBridgeModuleName(contractEvent)).toEqual(
            'SynapseBridgeAdapter'
          )
          expect(synapse.getBridgeModuleName(`${contractEvent}Event`)).toEqual(
            'SynapseBridgeAdapter'
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
    const mockEthProvider = mock<Provider>()
    const mockOpProvider = mock<Provider>()
    const sbaOriginToken = '0x00000000000000000000000000000000000000a1'
    const sbaRemoteToken = '0x00000000000000000000000000000000000000b1'
    const sbaFinalToken = '0x00000000000000000000000000000000000000c1'
    const sender = '0x0000000000000000000000000000000000000f01'
    const recipient = '0x0000000000000000000000000000000000000f02'

    const noOpRoute = {
      engineID: EngineID.NoOp,
      engineName: EngineID[EngineID.NoOp],
      chainId: SupportedChainId.ETH,
      fromToken: sbaOriginToken,
      fromAmount: BigNumber.from(1000),
      toToken: sbaOriginToken,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(1000),
      steps: [],
    }

    const destinationSwapRoute = {
      engineID: EngineID.DefaultPools,
      engineName: EngineID[EngineID.DefaultPools],
      chainId: SupportedChainId.OPTIMISM,
      fromToken: sbaRemoteToken,
      fromAmount: BigNumber.from(1000),
      toToken: sbaFinalToken,
      expectedToAmount: BigNumber.from(900),
      minToAmount: BigNumber.from(880),
      steps: [
        {
          token: sbaRemoteToken,
          amount: BigNumber.from(1000),
          msgValue: Zero,
          zapData: '0x1234',
        },
      ],
    }

    const setupSynapse = () => {
      const synapse = new SynapseSDK(
        [SupportedChainId.ETH, SupportedChainId.OPTIMISM],
        [mockEthProvider, mockOpProvider]
      )
      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]
      jest
        .spyOn(synapse.synapseBridgeAdapterModuleSet, 'getGasDropAmount')
        .mockResolvedValue(Zero)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getRemoteAddress'
        )
        .mockResolvedValue(sbaRemoteToken)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getNativeFee'
        )
        .mockResolvedValue(BigNumber.from(77))
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getEstimatedTime'
        )
        .mockResolvedValue(42)
      jest
        .spyOn(synapse.swapEngineSet, 'getBestQuote')
        .mockImplementation(async (input) => {
          if (
            input.chainId === SupportedChainId.ETH &&
            input.fromToken.toLowerCase() === sbaOriginToken.toLowerCase() &&
            input.toToken.toLowerCase() === sbaOriginToken.toLowerCase()
          ) {
            return noOpRoute as any
          }
          if (
            input.chainId === SupportedChainId.OPTIMISM &&
            input.fromToken.toLowerCase() === sbaRemoteToken.toLowerCase() &&
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
          SupportedChainId.ETH,
          'SynapseBridgeAdapter',
          txHash
        )
      ).resolves.toEqual(txHash)
    })

    it('delegates getBridgeTxStatus to the SBA module set', async () => {
      const synapse = setupSynapse()
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
          'SynapseBridgeAdapter',
          '0x1234'
        )
      ).resolves.toBe(true)
    })

    it('returns direct SBA bridgeV2 quotes without origin swap module names', async () => {
      const synapse = setupSynapse()

      const quotes: BridgeQuoteV2[] = await synapse.bridgeV2({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: sbaOriginToken,
        toToken: sbaRemoteToken,
        fromAmount: '1000',
        fromSender: sender,
        toRecipient: recipient,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0]).toMatchObject({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: sbaOriginToken,
        toToken: sbaRemoteToken,
        expectedToAmount: '1000',
        minToAmount: '1000',
        routerAddress: SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.ETH],
        moduleNames: ['SynapseBridgeAdapter'],
        nativeFee: '77',
        gasDropAmount: '0',
      })
      expect(quotes[0].tx).toBeDefined()
    })

    it('supports native-wrap origins in bridgeV2', async () => {
      const synapse = setupSynapse()
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getWrappedNativeToken'
        )
        .mockResolvedValue(sbaOriginToken)
      jest
        .spyOn(synapse.swapEngineSet, 'getBestQuote')
        .mockImplementation(async (input) => {
          if (
            input.chainId === SupportedChainId.ETH &&
            input.fromToken === ETH_NATIVE_TOKEN_ADDRESS &&
            input.toToken.toLowerCase() === sbaOriginToken.toLowerCase()
          ) {
            return {
              ...noOpRoute,
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
          return undefined as any
        })
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as any)

      const quotes: BridgeQuoteV2[] = await synapse.bridgeV2({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: ETH_NATIVE_TOKEN_ADDRESS,
        toToken: sbaRemoteToken,
        fromAmount: '1000',
        fromSender: sender,
        toRecipient: recipient,
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0]).toMatchObject({
        fromToken: ETH_NATIVE_TOKEN_ADDRESS,
        toToken: sbaRemoteToken,
        expectedToAmount: '1000',
        minToAmount: '950',
        moduleNames: [EngineID[EngineID.NoOp], 'SynapseBridgeAdapter'],
        routerAddress: SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.ETH],
        nativeFee: '77',
      })
      expect(quotes[0].tx).toBeDefined()
    })

    it('uses SBA as the bridge step inside multi-tx intents', async () => {
      const synapse = setupSynapse()

      const quotes: IntentQuote[] = await synapse.intent({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: sbaOriginToken,
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
        toToken: sbaRemoteToken,
        moduleNames: ['SynapseBridgeAdapter'],
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
      const harmonyToken = '0x0000000000000000000000000000000000000d01'
      const baseToken = '0x0000000000000000000000000000000000000b01'
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
          'getRemoteAddress'
        )
        .mockResolvedValue(baseToken)
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
      jest.spyOn(synapse.swapEngineSet, 'getBestQuote').mockResolvedValue({
        engineID: EngineID.NoOp,
        engineName: EngineID[EngineID.NoOp],
        chainId: SupportedChainId.HARMONY,
        fromToken: harmonyToken,
        fromAmount: BigNumber.from(1000),
        toToken: harmonyToken,
        expectedToAmount: BigNumber.from(1000),
        minToAmount: BigNumber.from(1000),
        steps: [],
      } as any)
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
      expect(quotes[0].moduleNames).toEqual(['SynapseBridgeAdapter'])
      expect(quotes[0].routerAddress).toEqual(
        SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.HARMONY]
      )
    })

    it('supports Harmony as an SBA destination chain', async () => {
      const ethProviderMock = mock<Provider>()
      const harmonyProvider = mock<Provider>()
      const ethToken = '0x0000000000000000000000000000000000000e01'
      const harmonyToken = '0x0000000000000000000000000000000000000d01'
      const synapse = new SynapseSDK(
        [SupportedChainId.ETH, SupportedChainId.HARMONY],
        [ethProviderMock, harmonyProvider]
      )

      synapse.allModuleSets = [synapse.synapseBridgeAdapterModuleSet]
      jest
        .spyOn(synapse.synapseBridgeAdapterModuleSet, 'getGasDropAmount')
        .mockResolvedValue(Zero)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getRemoteAddress'
        )
        .mockResolvedValue(harmonyToken)
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getNativeFee'
        )
        .mockResolvedValue(BigNumber.from(22))
      jest
        .spyOn(
          synapse.synapseBridgeAdapterModuleSet.modules[SupportedChainId.ETH],
          'getEstimatedTime'
        )
        .mockResolvedValue(44)
      jest.spyOn(synapse.swapEngineSet, 'getBestQuote').mockResolvedValue({
        engineID: EngineID.NoOp,
        engineName: EngineID[EngineID.NoOp],
        chainId: SupportedChainId.ETH,
        fromToken: ethToken,
        fromAmount: BigNumber.from(1000),
        toToken: ethToken,
        expectedToAmount: BigNumber.from(1000),
        minToAmount: BigNumber.from(1000),
        steps: [],
      } as any)
      jest
        .spyOn(synapse.swapEngineSet, 'generateRoute')
        .mockImplementation(async (_input, quote) => quote as any)

      const quotes = await synapse.bridgeV2({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.HARMONY,
        fromToken: ethToken,
        toToken: harmonyToken,
        fromAmount: '1000',
      })

      expect(quotes).toHaveLength(1)
      expect(quotes[0].moduleNames).toEqual(['SynapseBridgeAdapter'])
      expect(quotes[0].toToken).toEqual(harmonyToken)
    })
  })

  describe('getEstimatedTime', () => {
    const synapse = new SynapseSDK(
      [SupportedChainId.ETH, SupportedChainId.BSC],
      [ethProvider, bscProvider]
    )

    describe('Chain with a provider', () => {
      it('Returns estimated time for SynapseBridge', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseBridge')
        ).toEqual(MEDIAN_TIME_BRIDGE[SupportedChainId.ETH])

        expect(
          synapse.getEstimatedTime(SupportedChainId.BSC, 'SynapseBridge')
        ).toEqual(MEDIAN_TIME_BRIDGE[SupportedChainId.BSC])
      })

      it('Returns estimated time for SynapseCCTP', () => {
        expect(
          synapse.getEstimatedTime(SupportedChainId.ETH, 'SynapseCCTP')
        ).toEqual(MEDIAN_TIME_CCTP[SupportedChainId.ETH])
      })

      it('Throws when bridge module does not exist on a chain', () => {
        expect(() =>
          synapse.getEstimatedTime(SupportedChainId.BSC, 'SynapseCCTP')
        ).toThrow('No estimated time for chain 56')
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
