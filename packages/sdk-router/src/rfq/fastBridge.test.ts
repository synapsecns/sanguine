import { mock } from 'jest-mock-extended'
import { BigNumber, providers } from 'ethers'
import { Log, TransactionReceipt } from '@ethersproject/abstract-provider'

import { BridgeParams, FastBridge } from './fastBridge'
import { FAST_BRIDGE_ADDRESS_MAP, SupportedChainId } from '../constants'
import { NATIVE_ADDRESS } from '../constants/testValues'
import { Query } from '../module'

jest.mock('@ethersproject/contracts', () => {
  return {
    Contract: jest.fn().mockImplementation((...args: any[]) => {
      const actualContract = jest.requireActual('@ethersproject/contracts')
      const actualInstance = new actualContract.Contract(...args)
      return {
        address: args[0],
        interface: args[1],
        bridgeRelays: jest.fn(),
        populateTransaction: {
          bridge: actualInstance.populateTransaction.bridge,
        },
      }
    }),
  }
})

const expectCorrectPopulatedTx = (
  populatedTx: any,
  expectedAddress: string,
  expectedBridgeParams: BridgeParams
) => {
  expect(populatedTx).toBeDefined()
  expect(populatedTx.to).toEqual(expectedAddress)
  expect(populatedTx.data).toEqual(
    FastBridge.fastBridgeInterface.encodeFunctionData('bridge', [
      expectedBridgeParams,
    ])
  )
  if (
    expectedBridgeParams.originToken.toLowerCase() ===
    NATIVE_ADDRESS.toLowerCase()
  ) {
    expect(populatedTx.value).toEqual(
      BigNumber.from(expectedBridgeParams.originAmount)
    )
  } else {
    expect(populatedTx.value).toEqual(BigNumber.from(0))
  }
}

const createBridgeTests = (
  fastBridge: FastBridge,
  expectedBridgeParams: BridgeParams,
  originQuery: Query,
  destQuery: Query
) => {
  it('bridge without sendChainGas', async () => {
    const populatedTx = await fastBridge.bridge(
      expectedBridgeParams.to,
      SupportedChainId.OPTIMISM,
      expectedBridgeParams.originToken,
      BigNumber.from(expectedBridgeParams.originAmount),
      originQuery,
      destQuery
    )
    expectCorrectPopulatedTx(
      populatedTx,
      FAST_BRIDGE_ADDRESS_MAP[SupportedChainId.ARBITRUM],
      expectedBridgeParams
    )
  })

  it.skip('bridge with sendChainGas', async () => {
    const bridgeParamsWithGas = {
      ...expectedBridgeParams,
      sendChainGas: true,
    }
    // TODO: adjust this test once sendChainGas is supported
    const destQueryWithGas = {
      ...destQuery,
      rawParams: '0x0',
    }
    const populatedTx = await fastBridge.bridge(
      expectedBridgeParams.to,
      SupportedChainId.OPTIMISM,
      expectedBridgeParams.originToken,
      BigNumber.from(expectedBridgeParams.originAmount),
      originQuery,
      destQueryWithGas
    )
    expectCorrectPopulatedTx(
      populatedTx,
      FAST_BRIDGE_ADDRESS_MAP[SupportedChainId.ARBITRUM],
      bridgeParamsWithGas
    )
  })
}

describe('FastBridge', () => {
  const mockProvider = mock<providers.Provider>()

  const fastBridge = new FastBridge(
    SupportedChainId.ARBITRUM,
    mockProvider,
    FAST_BRIDGE_ADDRESS_MAP[SupportedChainId.ARBITRUM]
  )

  const mockedTxHash = '0x1234'
  const mockedSynapseTxId = '0x4321'

  describe('getSynapseTxId', () => {
    const bridgeRequestedTopic =
      '0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a'
    const mockedOriginLog = {
      address: fastBridge.address,
      // keccak256('BridgeRequested(bytes32,address,bytes)')
      topics: [bridgeRequestedTopic],
    } as Log
    const mockedUnrelatedLog = {
      address: fastBridge.address,
      topics: ['0x0'],
    } as Log
    const mockedReceipt = {
      logs: [mockedUnrelatedLog, mockedOriginLog],
    } as TransactionReceipt

    it('should return the Synapse transaction ID', async () => {
      // Return the mocked receipt only for the mocked transaction hash
      mockProvider.getTransactionReceipt.mockImplementation((txHash) => {
        if (txHash === mockedTxHash) {
          return Promise.resolve(mockedReceipt)
        } else {
          return Promise.resolve(undefined as any)
        }
      })
      // Return the mocked Synapse transaction ID for the mocked origin log
      fastBridge['fastBridgeContract'].interface.parseLog = jest.fn(
        (log: { topics: Array<string>; data: string }) => ({
          args: {
            transactionId:
              log.topics[0] === bridgeRequestedTopic ? mockedSynapseTxId : '',
          },
        })
      ) as any
      const result = await fastBridge.getSynapseTxId(mockedTxHash)
      expect(result).toEqual(mockedSynapseTxId)
    })
  })

  describe('getBridgeTxStatus', () => {
    it('returns false when bridgeRelays returns false', async () => {
      // Returns false only when mockedSynapseTxId is passed
      jest
        .spyOn(fastBridge['fastBridgeContract'], 'bridgeRelays')
        .mockImplementation((synapseTxId) =>
          Promise.resolve(synapseTxId !== mockedSynapseTxId)
        )
      const result = await fastBridge.getBridgeTxStatus(mockedSynapseTxId)
      expect(result).toEqual(false)
    })

    it('returns true when bridgeRelays returns true', async () => {
      // Returns true only when mockedSynapseTxId is passed
      jest
        .spyOn(fastBridge['fastBridgeContract'], 'bridgeRelays')
        .mockImplementation((synapseTxId) =>
          Promise.resolve(synapseTxId === mockedSynapseTxId)
        )
      const result = await fastBridge.getBridgeTxStatus(mockedSynapseTxId)
      expect(result).toEqual(true)
    })
  })

  describe('bridge', () => {
    const expectedBridgeParamsFragment = {
      dstChainId: SupportedChainId.OPTIMISM,
      sender: '0x0000000000000000000000000000000000000001',
      to: '0x0000000000000000000000000000000000000001',
      originAmount: 1234,
      destAmount: 5678,
      deadline: 9999,
    }

    const originQueryFragment = {
      routerAdapter: '0x0000000000000000000000000000000000000000',
      minAmountOut: BigNumber.from(expectedBridgeParamsFragment.originAmount),
      deadline: BigNumber.from(8888),
      rawParams: '0x',
    }

    const destQueryFragment = {
      routerAdapter: '0x0000000000000000000000000000000000000000',
      minAmountOut: BigNumber.from(expectedBridgeParamsFragment.destAmount),
      deadline: BigNumber.from(expectedBridgeParamsFragment.deadline),
      rawParams: '0x',
    }

    describe('bridge ERC20 token', () => {
      const expectedBridgeParams: BridgeParams = {
        ...expectedBridgeParamsFragment,
        originToken: '0x000000000000000000000000000000000000000A',
        destToken: '0x000000000000000000000000000000000000000b',
        sendChainGas: false,
      }

      const originQuery: Query = {
        ...originQueryFragment,
        tokenOut: expectedBridgeParams.originToken,
      }

      const destQuery: Query = {
        ...destQueryFragment,
        tokenOut: expectedBridgeParams.destToken,
      }

      createBridgeTests(
        fastBridge,
        expectedBridgeParams,
        originQuery,
        destQuery
      )
    })

    describe('bridge native token', () => {
      const expectedBridgeParams: BridgeParams = {
        ...expectedBridgeParamsFragment,
        originToken: NATIVE_ADDRESS,
        destToken: NATIVE_ADDRESS,
        sendChainGas: false,
      }

      const originQuery: Query = {
        ...originQueryFragment,
        tokenOut: expectedBridgeParams.originToken,
      }

      const destQuery: Query = {
        ...destQueryFragment,
        tokenOut: expectedBridgeParams.destToken,
      }

      createBridgeTests(
        fastBridge,
        expectedBridgeParams,
        originQuery,
        destQuery
      )
    })
  })
})
