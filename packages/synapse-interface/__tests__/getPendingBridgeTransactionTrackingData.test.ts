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

  it('keeps the Hyperliquid path working without an explicit tracked timestamp', () => {
    const trackedTransaction = getPendingBridgeTransactionTrackingData(
      {
        id: 333,
        originChain: mockChain,
        originToken: mockToken,
        originValue: '5',
        destinationChain: HYPERLIQUID,
        transactionHash: '0xhyper',
        isSubmitted: true,
      },
      zeroAddress
    )

    expect(trackedTransaction).toMatchObject({
      originTxHash: '0xhyper',
      estimatedTime: 0,
      timestamp: 333,
      bridgeModuleName: '',
      routerAddress: zeroAddress,
      destinationToken: mockToken,
    })
  })
})
