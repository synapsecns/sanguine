/** @jest-environment jsdom */

import { fireEvent, render, screen } from '@testing-library/react'

import { BridgeTransactionButton } from '@/components/StateManagedBridge/BridgeTransactionButton'
import { SynHyperCoreRecipientWarning } from '@/components/StateManagedBridge/BridgeWarnings'

const syn = {
  symbol: 'SYN',
  routeSymbol: 'SYN',
  addresses: { 1: '0x1111111111111111111111111111111111111111' },
}

jest.mock('viem', () => ({ isAddress: jest.fn(() => true) }))
jest.mock('wagmi', () => ({
  useAccount: () => ({ isConnected: true }),
  useAccountEffect: jest.fn(),
  useSwitchChain: () => ({
    chains: [{ id: 1, name: 'Ethereum' }],
    switchChain: jest.fn(),
  }),
}))
jest.mock('@rainbow-me/rainbowkit', () => ({
  useConnectModal: () => ({ openConnectModal: jest.fn() }),
}))
jest.mock('next-intl', () => ({
  useTranslations: () => (key: string, values?: Record<string, string>) =>
    values?.symbol ? key.replace('{symbol}', values.symbol) : key,
}))
jest.mock('../store/hooks', () => ({ useAppDispatch: () => jest.fn() }))
jest.mock('../slices/wallet/hooks', () => ({
  useWalletState: () => ({ isWalletPending: false }),
}))
jest.mock('../slices/bridgeQuote/hooks', () => ({
  useBridgeQuoteState: () => ({
    bridgeQuote: { bridgeModuleName: 'SYN' },
    isLoading: false,
  }),
}))
jest.mock('../slices/bridge/hooks', () => ({
  useBridgeState: () => ({
    destinationAddress: null,
    fromToken: syn,
    fromChainId: 1,
    toToken: syn,
    toChainId: 1337,
    debouncedFromValue: '1',
  }),
  useBridgeDisplayState: () => ({
    showDestinationWarning: false,
    isDestinationWarningAccepted: true,
  }),
}))
jest.mock(
  '../components/StateManagedBridge/hooks/useBridgeValidations',
  () => ({
    useBridgeValidations: () => ({
      hasValidInput: true,
      hasValidQuote: true,
      hasSufficientBalance: true,
      doesBridgeStateMatchQuote: true,
      onSelectedChain: true,
    }),
  })
)
jest.mock(
  '../components/StateManagedBridge/hooks/useConfirmNewBridgePrice',
  () => ({
    useConfirmNewBridgePrice: () => ({
      isPendingConfirmChange: false,
      onUserAcceptChange: jest.fn(),
    }),
  })
)
jest.mock('../utils/hyperliquid', () => ({
  isHyperliquidUsdcDeposit: () => false,
}))
jest.mock('../contexts/SegmentAnalyticsProvider', () => ({
  segmentAnalyticsEvent: jest.fn(),
}))
jest.mock('../components/buttons/TransactionButton', () => ({
  TransactionButton: ({ label, disabled, onClick }) => (
    <button disabled={disabled} onClick={onClick}>
      {label}
    </button>
  ),
}))

describe('inactive HyperCore recipient acknowledgement', () => {
  it('keeps the bridge button disabled until the warning is accepted', () => {
    const props = {
      approveTxn: jest.fn(),
      executeBridge: jest.fn(),
      isApproved: true,
      isBridgePaused: false,
      isTyping: false,
      isQuoteStale: false,
    }
    const { rerender } = render(
      <BridgeTransactionButton {...props} isRecipientWarningAccepted={false} />
    )

    expect(screen.getByRole('button', { name: 'Bridge SYN' })).toHaveProperty(
      'disabled',
      true
    )
    fireEvent.click(screen.getByRole('button', { name: 'Bridge SYN' }))
    expect(props.executeBridge).not.toHaveBeenCalled()

    rerender(
      <BridgeTransactionButton {...props} isRecipientWarningAccepted={true} />
    )
    expect(screen.getByRole('button', { name: 'Bridge SYN' })).toHaveProperty(
      'disabled',
      false
    )
    fireEvent.click(screen.getByRole('button', { name: 'Bridge SYN' }))
    expect(props.executeBridge).toHaveBeenCalledTimes(1)
  })

  it('exposes a labelled checkbox and reports acceptance', () => {
    const onAccept = jest.fn()
    render(
      <SynHyperCoreRecipientWarning accepted={false} onAccept={onAccept} />
    )

    const checkbox = screen.getByRole('checkbox', {
      name: 'SynHyperCoreRecipientWarningAccept',
    })
    expect(checkbox).toHaveProperty('checked', false)

    fireEvent.click(checkbox)
    expect(onAccept).toHaveBeenCalledWith(true)
  })
})
