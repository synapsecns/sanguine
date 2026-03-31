import { useMemo } from 'react'
import { ZeroAddress } from 'ethers'

import { CHAINS_BY_ID } from '@/constants/chains'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { useBridgeQuoteState } from '@/state/slices/bridgeQuote/hooks'
import { useWalletState } from '@/state/slices/wallet/hooks'
import { stringToBigInt } from '@/utils/stringToBigInt'

export const DEFAULT_NATIVE_FEE_VALIDATION_MESSAGE =
  'Insufficient native token balance to complete bridge transaction'

export const useNativeFeeValidation = (): {
  hasEnoughNativeBalanceForQuoteFee: boolean
  nativeFeeValidationMessage: string | null
} => {
  const { balances } = useWalletState()
  const { bridgeQuote } = useBridgeQuoteState()
  const { debouncedInputAmount, originChainId, originToken } = useBridgeState()

  return useMemo(() => {
    if (!originChainId || !originToken) {
      return {
        hasEnoughNativeBalanceForQuoteFee: true,
        nativeFeeValidationMessage: null,
      }
    }

    const nativeFeeWei = bridgeQuote?.nativeFee
    const hasValidQuoteWithNativeFee =
      typeof nativeFeeWei === 'bigint' &&
      nativeFeeWei > 0n &&
      Boolean(bridgeQuote?.outputAmount)

    if (!hasValidQuoteWithNativeFee) {
      return {
        hasEnoughNativeBalanceForQuoteFee: true,
        nativeFeeValidationMessage: null,
      }
    }

    const nativeSymbol = CHAINS_BY_ID[originChainId]?.nativeCurrency?.symbol
    const isOriginNativeToken =
      originToken.addresses[originChainId] === ZeroAddress
    const nativeBalance = Array.isArray(balances)
      ? balances.find(
          (tokenBalance) =>
            tokenBalance?.token?.addresses?.[originChainId] === ZeroAddress
        )
      : null

    const rawNativeBalanceWei =
      nativeBalance?.balance !== undefined && nativeBalance?.balance !== null
        ? BigInt(nativeBalance.balance)
        : null
    const inputAmountWei = isOriginNativeToken
      ? stringToBigInt(
          debouncedInputAmount ?? '0',
          originToken.decimals[originChainId]
        ) ?? 0n
      : 0n
    const requiredNativeBalanceWei = nativeFeeWei + inputAmountWei
    const hasEnoughNativeBalanceForQuoteFee =
      rawNativeBalanceWei !== null &&
      rawNativeBalanceWei >= requiredNativeBalanceWei

    if (hasEnoughNativeBalanceForQuoteFee) {
      return {
        hasEnoughNativeBalanceForQuoteFee: true,
        nativeFeeValidationMessage: null,
      }
    }

    return {
      hasEnoughNativeBalanceForQuoteFee: false,
      nativeFeeValidationMessage: nativeSymbol
        ? `Insufficient ${nativeSymbol} balance to complete bridge transaction`
        : DEFAULT_NATIVE_FEE_VALIDATION_MESSAGE,
    }
  }, [balances, bridgeQuote, debouncedInputAmount, originChainId, originToken])
}
