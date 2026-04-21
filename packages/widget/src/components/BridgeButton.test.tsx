import { fireEvent, render, screen } from '@testing-library/react'

import { BridgeButton } from '@/components/BridgeButton'
import { useNativeFeeValidation } from '@/hooks/useNativeFeeValidation'
import { useValidations } from '@/hooks/useValidations'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { Web3Context } from '@/providers/Web3Provider'

jest.mock('@/hooks/useValidations')
jest.mock('@/hooks/useNativeFeeValidation')
jest.mock('@/state/slices/bridge/hooks')
jest.mock('@/utils/actions/switchNetwork', () => ({
  switchNetwork: jest.fn(),
}))
jest.mock('@/constants/chains', () => ({
  CHAINS_BY_ID: {
    1: {
      nativeCurrency: {
        symbol: 'ETH',
      },
    },
  },
}))

const originChain = {
  id: 1,
  name: 'Ethereum',
} as any

const renderBridgeButton = () => {
  return render(
    <Web3Context.Provider
      value={{
        web3Provider: {
          connectedAddress: '0xabc',
          networkId: 1,
          provider: {},
          signer: null,
        },
        setWeb3Provider: jest.fn() as any,
      }}
    >
      <BridgeButton
        originChain={originChain}
        isValidQuote={true}
        handleApprove={jest.fn()}
        handleBridge={jest.fn()}
        isApprovalPending={false}
        isBridgePending={false}
        isBridgePaused={false}
      />
    </Web3Context.Provider>
  )
}

const hoverTooltip = (container: HTMLElement) => {
  const tooltip = container.querySelector('[data-test-id="tool-tip"]')

  expect(tooltip).not.toBeNull()

  fireEvent.mouseEnter(tooltip as HTMLElement)
}

describe('BridgeButton', () => {
  beforeEach(() => {
    ;(useBridgeState as jest.Mock).mockReturnValue({
      originChainId: 1,
    })
    ;(useValidations as jest.Mock).mockReturnValue({
      hasEnoughBalance: true,
      isInputValid: true,
      isApproved: false,
      onSelectedChain: true,
    })
    ;(useNativeFeeValidation as jest.Mock).mockReturnValue({
      hasEnoughNativeBalanceForQuoteFee: true,
      nativeFeeValidationMessage: null,
    })
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('disables Approve & Sign when an ERC20 route cannot cover the native bridge fee', () => {
    ;(useNativeFeeValidation as jest.Mock).mockReturnValue({
      hasEnoughNativeBalanceForQuoteFee: false,
      nativeFeeValidationMessage:
        'Insufficient ETH balance to complete bridge transaction',
    })

    const { container } = renderBridgeButton()
    const button = screen.getByRole('button', { name: 'Approve & Sign' })

    expect(button).toBeDisabled()

    hoverTooltip(container)

    expect(
      screen.getByText(
        'Insufficient ETH balance to complete bridge transaction'
      )
    ).toBeInTheDocument()
  })

  it('disables Send when a native route cannot cover amount plus native bridge fee', () => {
    ;(useValidations as jest.Mock).mockReturnValue({
      hasEnoughBalance: true,
      isInputValid: true,
      isApproved: true,
      onSelectedChain: true,
    })
    ;(useNativeFeeValidation as jest.Mock).mockReturnValue({
      hasEnoughNativeBalanceForQuoteFee: false,
      nativeFeeValidationMessage:
        'Insufficient ETH balance to complete bridge transaction',
    })

    const { container } = renderBridgeButton()
    const button = screen.getByRole('button', { name: 'Send' })

    expect(button).toBeDisabled()

    hoverTooltip(container)

    expect(
      screen.getByText(
        'Insufficient ETH balance to complete bridge transaction'
      )
    ).toBeInTheDocument()
  })

  it('keeps the existing amount-balance tooltip ahead of the native-fee validation', () => {
    ;(useValidations as jest.Mock).mockReturnValue({
      hasEnoughBalance: false,
      isInputValid: true,
      isApproved: true,
      onSelectedChain: true,
    })
    ;(useNativeFeeValidation as jest.Mock).mockReturnValue({
      hasEnoughNativeBalanceForQuoteFee: false,
      nativeFeeValidationMessage:
        'Insufficient ETH balance to complete bridge transaction',
    })

    const { container } = renderBridgeButton()
    const button = screen.getByRole('button', { name: 'Send' })

    expect(button).toBeDisabled()

    hoverTooltip(container)

    expect(
      screen.getByText('Amount may not exceed available balance')
    ).toBeInTheDocument()
    expect(
      screen.queryByText(
        'Insufficient ETH balance to complete bridge transaction'
      )
    ).not.toBeInTheDocument()
  })
})
