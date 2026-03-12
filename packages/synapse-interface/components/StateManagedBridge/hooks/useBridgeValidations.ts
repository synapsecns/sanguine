import { useMemo } from 'react'
import { useAccount } from 'wagmi'

import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { BridgeQuoteState } from '@/slices/bridgeQuote/reducer'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { useBridgeSelections } from './useBridgeSelections'
import { ARBITRUM, HYPERLIQUID } from '@/constants/chains/master'

export const useBridgeValidations = () => {
  const { chainId } = useAccount()
  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    debouncedFromValue,
  }: BridgeState = useBridgeState()
  const { bridgeQuote }: BridgeQuoteState = useBridgeQuoteState()
  const { fromTokenBalance, debouncedFromValueBigInt } = useBridgeSelections()

  const hasValidInput: boolean = useMemo(() => {
    if (debouncedFromValue === '') return false
    if (hasOnlyZeroes(debouncedFromValue)) return false
    return debouncedFromValueBigInt > 0n
  }, [debouncedFromValue, debouncedFromValueBigInt])

  const hasValidFromSelections = useMemo(() => {
    return Boolean(fromChainId && fromToken)
  }, [fromChainId, fromToken])

  const hasValidSelections = useMemo(() => {
    return Boolean(fromChainId && fromToken && toChainId && toToken)
  }, [fromChainId, fromToken, toChainId, toToken])

  const hasValidQuote: boolean = useMemo(() => {
    return bridgeQuote !== EMPTY_BRIDGE_QUOTE
  }, [bridgeQuote])

  const hasSufficientBalance: boolean = useMemo(() => {
    return hasValidSelections
      ? debouncedFromValueBigInt <= fromTokenBalance
      : false
  }, [hasValidSelections, debouncedFromValueBigInt, fromTokenBalance])

  const stringifiedBridgeQuote = useMemo(() => {
    return constructStringifiedBridgeSelections(
      bridgeQuote.inputAmountForQuote,
      bridgeQuote.originChainId,
      bridgeQuote.originTokenForQuote,
      bridgeQuote.destChainId,
      bridgeQuote.destTokenForQuote
    )
  }, [
    bridgeQuote.inputAmountForQuote,
    bridgeQuote.originChainId,
    bridgeQuote.originTokenForQuote,
    bridgeQuote.destChainId,
    bridgeQuote.destTokenForQuote,
  ])

  const stringifiedBridgeState = useMemo(() => {
    return constructStringifiedBridgeSelections(
      debouncedFromValue,
      fromChainId,
      fromToken,
      toChainId === HYPERLIQUID.id ? ARBITRUM.id : toChainId,
      toToken
    )
  }, [debouncedFromValue, fromChainId, fromToken, toChainId, toToken])

  const doesBridgeStateMatchQuote = useMemo(() => {
    return stringifiedBridgeQuote === stringifiedBridgeState
  }, [stringifiedBridgeQuote, stringifiedBridgeState])

  const onSelectedChain: boolean = useMemo(() => {
    return chainId === fromChainId
  }, [fromChainId, chainId])

  return {
    hasValidInput,
    hasValidFromSelections,
    hasValidSelections,
    hasValidQuote,
    hasSufficientBalance,
    doesBridgeStateMatchQuote,
    onSelectedChain,
  }
}

export const constructStringifiedBridgeSelections = (
  originAmount,
  originChainId,
  originToken,
  destChainId,
  destToken
) => {
  const state = {
    originAmount,
    originChainId,
    originToken,
    destChainId,
    destToken,
  }
  return JSON.stringify(state)
}
