import { useMemo } from 'react'
import { useNetwork } from 'wagmi'

import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState, initialState } from '@/slices/bridge/reducer'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { stringToBigInt } from '@/utils/bigint/format'

export const useBridgeStatus = (): {
  hasValidRoute: boolean
  hasEnoughBalance: boolean
  hasInputAmount: boolean
  onSelectedChain: boolean
} => {
  const { chain } = useNetwork()
  const {
    debouncedFromValue,
    fromChainId,
    fromToken,
    isLoading,
    bridgeQuote,
  }: BridgeState = useBridgeState()
  const balances = usePortfolioBalances()

  const hasValidRoute: boolean = useMemo(() => {
    return !isLoading && bridgeQuote.outputAmount > 0n
  }, [isLoading, bridgeQuote])

  const hasEnoughBalance: boolean = useMemo(() => {
    const currentBalance: bigint = balances[fromChainId]?.find(
      (token) => token.token === fromToken
    )?.balance
    const precisedebouncedFromValue: bigint = stringToBigInt(
      debouncedFromValue,
      fromToken?.decimals[fromChainId]
    )

    return currentBalance >= precisedebouncedFromValue
  }, [balances, debouncedFromValue, fromToken, fromChainId])

  const hasInputAmount: boolean = useMemo(() => {
    const isEmpty: boolean =
      debouncedFromValue === initialState.debouncedFromValue
    const isZero: boolean = parseFloat(debouncedFromValue) === 0

    return Boolean(!isEmpty && !isZero)
  }, [debouncedFromValue])

  const onSelectedChain: boolean = useMemo(() => {
    return chain?.id === fromChainId
  }, [fromChainId, chain])

  return {
    hasValidRoute,
    hasEnoughBalance,
    hasInputAmount,
    onSelectedChain,
  }
}
