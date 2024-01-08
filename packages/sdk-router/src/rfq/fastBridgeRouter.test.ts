import { mock } from 'jest-mock-extended'
import { BigNumber, providers } from 'ethers'
import { Log, TransactionReceipt } from '@ethersproject/abstract-provider'

import { FastBridgeRouter } from './fastBridgeRouter'
import { FAST_BRIDGE_ROUTER_ADDRESS_MAP, SupportedChainId } from '../constants'
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
        fastBridge: jest.fn(),
        populateTransaction: {
          bridge: actualInstance.populateTransaction.bridge,
        },
      }
    }),
  }
})

const expectCorrectPopulatedTx = (
  populatedTx: any,
  originToken: string,
  originAmount: number,
  expectedAddress: string,
  expectedData: string
) => {
  expect(populatedTx).toBeDefined()
  expect(populatedTx.to).toEqual(expectedAddress)
  expect(populatedTx.data).toEqual(expectedData)
  if (originToken.toLowerCase() === NATIVE_ADDRESS.toLowerCase()) {
    expect(populatedTx.value).toEqual(BigNumber.from(originAmount))
  } else {
    expect(populatedTx.value).toEqual(BigNumber.from(0))
  }
}

type BridgeTestsParams = {
  dstChainId: number
  sender: string
  to: string
  originAmount: number
  destAmount: number
  deadline: number
}

const createBridgeTest = (
  fastBridgeRouter: FastBridgeRouter,
  bridgeParams: BridgeTestsParams,
  originQuery: Query,
  destQuery: Query
) => {
  it('bridge', async () => {
    const populatedTx = await fastBridgeRouter.bridge(
      bridgeParams.to,
      bridgeParams.dstChainId,
      originQuery.tokenOut,
      bridgeParams.originAmount,
      originQuery,
      destQuery
    )
    const expectedData =
      FastBridgeRouter.fastBridgeRouterInterface.encodeFunctionData('bridge', [
        bridgeParams.to,
        bridgeParams.dstChainId,
        originQuery.tokenOut,
        bridgeParams.originAmount,
        originQuery,
        destQuery,
      ])
    expectCorrectPopulatedTx(
      populatedTx,
      originQuery.tokenOut,
      bridgeParams.originAmount,
      fastBridgeRouter.address,
      expectedData
    )
  })
}

describe('FastBridgeRouter', () => {
  const mockProvider = mock<providers.Provider>()

  const fastBridgeRouter = new FastBridgeRouter(
    SupportedChainId.ARBITRUM,
    mockProvider,
    FAST_BRIDGE_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
  )

  const mockedTxHash = '0x1234'
  const mockedSynapseTxId = '0x4321'

  const mockedFastBridgeAddress = '0x000000000000000000000000000000000000dEaD'

  beforeAll(async () => {
    // Override .fastBridge()
    jest
      .spyOn(fastBridgeRouter['routerContract'], 'fastBridge')
      .mockImplementation(() => Promise.resolve(mockedFastBridgeAddress))
    // Populate the cache
    await fastBridgeRouter.getFastBridgeContract()
  })

  describe('getSynapseTxId', () => {
    // keccak256('BridgeRequested(bytes32,address,bytes)')
    const bridgeRequestedTopic =
      '0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf'
    const mockedOriginLog = {
      address: mockedFastBridgeAddress,
      topics: [bridgeRequestedTopic],
    } as Log
    const mockedUnrelatedLog = {
      address: mockedFastBridgeAddress,
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
      fastBridgeRouter['fastBridgeContractCache']!.interface.parseLog = jest.fn(
        (log: { topics: Array<string>; data: string }) => ({
          args: {
            transactionId:
              log.topics[0] === bridgeRequestedTopic ? mockedSynapseTxId : '',
          },
        })
      ) as any
      const result = await fastBridgeRouter.getSynapseTxId(mockedTxHash)
      expect(result).toEqual(mockedSynapseTxId)
    })
  })

  describe('getBridgeTxStatus', () => {
    it('returns false when bridgeRelays returns false', async () => {
      // Returns false only when mockedSynapseTxId is passed
      jest
        .spyOn(fastBridgeRouter['fastBridgeContractCache']!, 'bridgeRelays')
        .mockImplementation((synapseTxId) =>
          Promise.resolve(synapseTxId !== mockedSynapseTxId)
        )
      const result = await fastBridgeRouter.getBridgeTxStatus(mockedSynapseTxId)
      expect(result).toEqual(false)
    })

    it('returns true when bridgeRelays returns true', async () => {
      // Returns true only when mockedSynapseTxId is passed
      jest
        .spyOn(fastBridgeRouter['fastBridgeContractCache']!, 'bridgeRelays')
        .mockImplementation((synapseTxId) =>
          Promise.resolve(synapseTxId === mockedSynapseTxId)
        )
      const result = await fastBridgeRouter.getBridgeTxStatus(mockedSynapseTxId)
      expect(result).toEqual(true)
    })
  })

  describe('bridge', () => {
    const bridgeParams: BridgeTestsParams = {
      dstChainId: SupportedChainId.OPTIMISM,
      sender: '0x0000000000000000000000000000000000000001',
      to: '0x0000000000000000000000000000000000000002',
      originAmount: 1234,
      destAmount: 5678,
      deadline: 9999,
    }

    const originQueryFragment = {
      routerAdapter: '0x0000000000000000000000000000000000000000',
      minAmountOut: BigNumber.from(bridgeParams.originAmount),
      deadline: BigNumber.from(8888),
      rawParams: '0x',
    }

    const destQueryFragment = {
      routerAdapter: '0x0000000000000000000000000000000000000000',
      minAmountOut: BigNumber.from(bridgeParams.destAmount),
      deadline: BigNumber.from(bridgeParams.deadline),
      rawParams: '0x2a',
    }

    describe('ERC20 token', () => {
      const originQuery: Query = {
        ...originQueryFragment,
        tokenOut: '0x000000000000000000000000000000000000000A',
      }

      const destQuery: Query = {
        ...destQueryFragment,
        tokenOut: '0x000000000000000000000000000000000000000b',
      }

      createBridgeTest(fastBridgeRouter, bridgeParams, originQuery, destQuery)
    })

    describe('Native token', () => {
      const originQuery: Query = {
        ...originQueryFragment,
        tokenOut: NATIVE_ADDRESS,
      }

      const destQuery: Query = {
        ...destQueryFragment,
        tokenOut: NATIVE_ADDRESS,
      }

      createBridgeTest(fastBridgeRouter, bridgeParams, originQuery, destQuery)
    })
  })
})
