import { renderHook } from '@testing-library/react'

import { useNativeFeeValidation } from '@/hooks/useNativeFeeValidation'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { useBridgeQuoteState } from '@/state/slices/bridgeQuote/hooks'
import { useWalletState } from '@/state/slices/wallet/hooks'

jest.mock('ethers', () => ({
  ZeroAddress: '0x0000000000000000000000000000000000000000',
}))
jest.mock('@/state/slices/bridge/hooks')
jest.mock('@/state/slices/bridgeQuote/hooks')
jest.mock('@/state/slices/wallet/hooks')
jest.mock('@/constants/chains', () => ({
  CHAINS_BY_ID: {
    1: {
      nativeCurrency: {
        symbol: 'ETH',
      },
    },
  },
}))

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'

const nativeToken = {
  addresses: { 1: ZERO_ADDRESS },
  decimals: { 1: 18 },
  symbol: 'ETH',
  routeSymbol: 'ETH',
}

const erc20Token = {
  addresses: { 1: '0x1' },
  decimals: { 1: 6 },
  symbol: 'USDC',
  routeSymbol: 'USDC',
}

const createNativeBalance = (balance: string) => ({
  token: nativeToken,
  balance,
  parsedBalance: balance,
})

describe('useNativeFeeValidation', () => {
  beforeEach(() => {
    ;(useBridgeState as jest.Mock).mockReturnValue({
      debouncedInputAmount: '1',
      originChainId: 1,
      originToken: erc20Token,
    })
    ;(useBridgeQuoteState as jest.Mock).mockReturnValue({
      bridgeQuote: {
        outputAmount: 1n,
        nativeFee: 100n,
      },
    })
    ;(useWalletState as jest.Mock).mockReturnValue({
      balances: [createNativeBalance('1000')],
    })
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('requires native balance to cover nativeFee for non-native origin tokens', () => {
    ;(useWalletState as jest.Mock).mockReturnValue({
      balances: [createNativeBalance('99')],
    })

    const { result } = renderHook(() => useNativeFeeValidation())

    expect(result.current.hasEnoughNativeBalanceForQuoteFee).toBe(false)
    expect(result.current.nativeFeeValidationMessage).toBe(
      'Insufficient ETH balance to complete bridge transaction'
    )
  })

  it('requires native balance to cover amount plus nativeFee for native origin tokens', () => {
    ;(useBridgeState as jest.Mock).mockReturnValue({
      debouncedInputAmount: '1',
      originChainId: 1,
      originToken: nativeToken,
    })
    ;(useWalletState as jest.Mock).mockReturnValue({
      balances: [createNativeBalance('1000000000000000050')],
    })
    ;(useBridgeQuoteState as jest.Mock).mockReturnValue({
      bridgeQuote: {
        outputAmount: 1n,
        nativeFee: 100n,
      },
    })

    const { result } = renderHook(() => useNativeFeeValidation())

    expect(result.current.hasEnoughNativeBalanceForQuoteFee).toBe(false)
    expect(result.current.nativeFeeValidationMessage).toBe(
      'Insufficient ETH balance to complete bridge transaction'
    )
  })

  it('passes when native balance covers the required native fee', () => {
    const { result } = renderHook(() => useNativeFeeValidation())

    expect(result.current.hasEnoughNativeBalanceForQuoteFee).toBe(true)
    expect(result.current.nativeFeeValidationMessage).toBeNull()
  })

  it('ignores the validation when nativeFee is zero', () => {
    ;(useBridgeQuoteState as jest.Mock).mockReturnValue({
      bridgeQuote: {
        outputAmount: 1n,
        nativeFee: 0n,
      },
    })
    ;(useWalletState as jest.Mock).mockReturnValue({
      balances: [],
    })

    const { result } = renderHook(() => useNativeFeeValidation())

    expect(result.current.hasEnoughNativeBalanceForQuoteFee).toBe(true)
    expect(result.current.nativeFeeValidationMessage).toBeNull()
  })

  it('fails conservatively when the native balance is missing for a quoted fee', () => {
    ;(useWalletState as jest.Mock).mockReturnValue({
      balances: [],
    })

    const { result } = renderHook(() => useNativeFeeValidation())

    expect(result.current.hasEnoughNativeBalanceForQuoteFee).toBe(false)
    expect(result.current.nativeFeeValidationMessage).toBe(
      'Insufficient ETH balance to complete bridge transaction'
    )
  })
})
