import { mock } from 'jest-mock-extended'
import { providers } from 'ethers'
import { Log, TransactionReceipt } from '@ethersproject/abstract-provider'

import { FastBridge } from './fastBridge'
import { FAST_BRIDGE_ADDRESS_MAP, SupportedChainId } from '../constants'

jest.mock('@ethersproject/contracts', () => {
  return {
    Contract: jest.fn().mockImplementation((...args: any[]) => {
      return {
        address: args[0],
        interface: args[1],
        bridgeRelays: jest.fn(),
        populateTransaction: {
          bridge: jest.fn(),
        },
      }
    }),
  }
})

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
      '0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf'
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

  // TODO: test bridge()
})
