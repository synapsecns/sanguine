import { fireEvent, render, screen } from '@testing-library/react'

import { AvailableBalance } from '@/components/AvailableBalance'
import { useCurrentTokenBalance } from '@/hooks/useCurrentTokenBalance'
import { useValidations } from '@/hooks/useValidations'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { useWalletState } from '@/state/slices/wallet/hooks'
import { FetchState } from '@/state/slices/wallet/reducer'

jest.mock('@/hooks/useCurrentTokenBalance')
jest.mock('@/hooks/useValidations')
jest.mock('@/state/slices/bridge/hooks')
jest.mock('@/state/slices/wallet/hooks')

describe('AvailableBalance', () => {
  const setInputAmount = jest.fn()

  beforeEach(() => {
    ;(useBridgeState as jest.Mock).mockReturnValue({
      originChainId: 1,
      originToken: { symbol: 'Token' },
    })
    ;(useCurrentTokenBalance as jest.Mock).mockReturnValue({
      decimals: 18,
      parsedBalance: '1.5000',
      rawBalance: '1500000000000000000',
    })
    ;(useValidations as jest.Mock).mockReturnValue({
      hasEnoughBalance: true,
    })
    ;(useWalletState as jest.Mock).mockReturnValue({
      balancesFetchStatus: FetchState.VALID,
    })
    setInputAmount.mockReset()
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('keeps the existing non-native balance label and raw-balance fill behavior', () => {
    render(
      <AvailableBalance
        connectedAddress="0xabc"
        nativeSafeMax={{
          amountWei: null,
          fillAmount: null,
          isClickable: false,
          isNativeOriginToken: false,
          labelAmount: null,
          status: 'idle',
        }}
        setInputAmount={setInputAmount}
      />
    )

    fireEvent.click(screen.getByText('Available 1.5000'))

    expect(setInputAmount).toHaveBeenCalledWith('1.500000000000000000')
  })

  it('shows the native bridgeable amount and fills the safe max instead of the raw balance', () => {
    render(
      <AvailableBalance
        connectedAddress="0xabc"
        nativeSafeMax={{
          amountWei: 1250000000000000000n,
          fillAmount: '1.250000000000000000',
          isClickable: true,
          isNativeOriginToken: true,
          labelAmount: '1.2500',
          status: 'ready',
        }}
        setInputAmount={setInputAmount}
      />
    )

    fireEvent.click(screen.getByText('Bridgeable 1.2500'))

    expect(setInputAmount).toHaveBeenCalledWith('1.250000000000000000')
  })

  it.each([
    [
      'idle',
      {
        amountWei: null,
        fillAmount: null,
        isClickable: false,
        isNativeOriginToken: true,
        labelAmount: null,
        status: 'idle' as const,
      },
      'Available 1.5000',
    ],
    [
      'loading',
      {
        amountWei: null,
        fillAmount: null,
        isClickable: false,
        isNativeOriginToken: true,
        labelAmount: null,
        status: 'loading' as const,
      },
      'Available 1.5000',
    ],
    [
      'unavailable',
      {
        amountWei: null,
        fillAmount: null,
        isClickable: false,
        isNativeOriginToken: true,
        labelAmount: null,
        status: 'unavailable' as const,
      },
      'Available 1.5000',
    ],
    [
      'zero',
      {
        amountWei: 0n,
        fillAmount: '0.0',
        isClickable: false,
        isNativeOriginToken: true,
        labelAmount: '0.0',
        status: 'ready' as const,
      },
      'Bridgeable 0.0',
    ],
  ])(
    'renders the native %s state as non-clickable',
    (_scenario, nativeSafeMax, label) => {
      render(
        <AvailableBalance
          connectedAddress="0xabc"
          nativeSafeMax={nativeSafeMax}
          setInputAmount={setInputAmount}
        />
      )

      fireEvent.click(screen.getByText(label))

      expect(setInputAmount).not.toHaveBeenCalled()
    }
  )

  it('mirrors the non-native loading label for native assets while the wallet balance is still loading', () => {
    ;(useWalletState as jest.Mock).mockReturnValue({
      balancesFetchStatus: FetchState.LOADING,
    })

    render(
      <AvailableBalance
        connectedAddress="0xabc"
        nativeSafeMax={{
          amountWei: null,
          fillAmount: null,
          isClickable: false,
          isNativeOriginToken: true,
          labelAmount: null,
          status: 'loading',
        }}
        setInputAmount={setInputAmount}
      />
    )

    expect(screen.getByText('loading...')).toBeInTheDocument()
  })
})
