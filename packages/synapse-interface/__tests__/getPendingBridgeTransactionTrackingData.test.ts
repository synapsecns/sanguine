import { zeroAddress } from 'viem'

import { HYPERLIQUID } from '@/constants/chains/master'
import { getPendingBridgeTransactionTrackingData } from '@/utils/getPendingBridgeTransactionTrackingData'

const mockChain = {
  id: 1,
  name: 'Ethereum',
} as any

const mockToken = {
  routeSymbol: 'USDC',
  symbol: 'USDC',
} as any

describe('getPendingBridgeTransactionTrackingData', () => {
  it('tracks SYN composer deposits as pending bridge transactions', () => {
    const transaction = {
      id: 111,
      originChain: mockChain,
      originToken: { routeSymbol: 'SYN', symbol: 'SYN' } as any,
      originValue: '1',
      destinationChain: HYPERLIQUID,
      transactionHash: '0xsyn',
      isSubmitted: true,
      estimatedTime: 840,
      bridgeModuleName: 'SYN',
      routerAddress: zeroAddress,
    }
    expect(
      getPendingBridgeTransactionTrackingData(transaction, zeroAddress)
    ).toBeNull()
    expect(
      getPendingBridgeTransactionTrackingData(
        { ...transaction, timestamp: 222 },
        zeroAddress
      )
    ).toMatchObject({
      destinationChain: HYPERLIQUID,
      bridgeModuleName: 'SYN',
      timestamp: 222,
      status: 'pending',
    })
  })

  it('does not promote regular bridge transactions until a tracked timestamp exists', () => {
    const trackedTransaction = getPendingBridgeTransactionTrackingData(
      {
        id: 111,
        originChain: mockChain,
        originToken: mockToken,
        originValue: '1',
        destinationChain: { id: 42161, name: 'Arbitrum' } as any,
        destinationToken: mockToken,
        transactionHash: '0xhash',
        isSubmitted: true,
        estimatedTime: 90,
        bridgeModuleName: 'SynapseRFQ',
        routerAddress: zeroAddress,
      },
      zeroAddress
    )

    expect(trackedTransaction).toBeNull()
  })

  it('uses the stored post-confirmation timestamp for regular bridge transactions', () => {
    const trackedTransaction = getPendingBridgeTransactionTrackingData(
      {
        id: 111,
        originChain: mockChain,
        originToken: mockToken,
        originValue: '1',
        destinationChain: { id: 42161, name: 'Arbitrum' } as any,
        destinationToken: mockToken,
        transactionHash: '0xhash',
        timestamp: 222,
        isSubmitted: true,
        estimatedTime: 90,
        bridgeModuleName: 'SynapseRFQ',
        routerAddress: zeroAddress,
        destinationAddress: zeroAddress,
      },
      zeroAddress
    )

    expect(trackedTransaction).toMatchObject({
      originTxHash: '0xhash',
      estimatedTime: 90,
      timestamp: 222,
    })
  })

  it('does not promote legacy Hyperliquid deposits without a tracked timestamp', () => {
    expect(
      getPendingBridgeTransactionTrackingData(
        {
          id: 333,
          originChain: mockChain,
          originToken: mockToken,
          originValue: '5',
          destinationChain: { ...HYPERLIQUID, id: 998 },
          transactionHash: '0xhyper',
          isSubmitted: true,
        },
        zeroAddress
      )
    ).toBeNull()
  })
})
