import { useMemo } from 'react'
import { useNetwork } from 'wagmi'

import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState, initialState } from '@/slices/bridge/reducer'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { stringToBigInt } from '@/utils/bigint/format'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'

export const useBridgeStatus = (): {
  hasValidSelections: boolean
  hasValidRoute: boolean
  hasEnoughBalance: boolean
  hasInputAmount: boolean
  hasEnoughApproved: boolean
  onSelectedChain: boolean
} => {
  const { chain } = useNetwork()
  const {
    debouncedFromValue,
    fromValue,
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    isLoading,
    bridgeQuote,
  }: BridgeState = useBridgeState()
  const balances = usePortfolioBalances()

  const hasValidSelections: boolean = useMemo(() => {
    return Boolean(fromChainId && toChainId && fromToken && toToken)
  }, [fromChainId, toChainId, fromToken, toToken])

  const hasValidRoute: boolean = useMemo(() => {
    const hasInput: boolean =
      debouncedFromValue !== initialState.debouncedFromValue
    const hasQuote: boolean = bridgeQuote !== initialState.bridgeQuote

    return Boolean(!isLoading && hasInput && hasQuote)
  }, [isLoading, bridgeQuote, debouncedFromValue])

  const hasEnoughBalance: boolean = useMemo(() => {
    const currentBalance: bigint = balances[fromChainId]?.find(
      (token) => token.token === fromToken
    )?.balance
    const preciseFromValue: bigint = stringToBigInt(
      fromValue,
      fromToken?.decimals[fromChainId]
    )

    return currentBalance >= preciseFromValue
  }, [balances, fromValue, fromToken, fromChainId])

  const hasInputAmount: boolean = useMemo(() => {
    const isEmpty: boolean = fromValue === initialState.fromValue

    return Boolean(!hasOnlyZeroes(fromValue) && !isEmpty)
  }, [fromValue])

  const hasEnoughApproved = false

  const onSelectedChain: boolean = useMemo(() => {
    return chain?.id === fromChainId
  }, [fromChainId, chain])

  return {
    hasValidSelections,
    hasValidRoute,
    hasEnoughBalance,
    hasInputAmount,
    hasEnoughApproved,
    onSelectedChain,
  }
}
