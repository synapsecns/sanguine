import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { stringToBigInt } from '@/utils/bigint/format'

export const useBridgeSelections = () => {
  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    debouncedFromValue,
  }: BridgeState = useBridgeState()

  const fromTokenSymbol = fromToken?.symbol
  const fromTokenDecimals = fromToken?.decimals[fromChainId]
  const fromTokenAddress = fromToken?.addresses[fromChainId]

  const toTokenSymbol = toToken?.symbol
  const toTokenDecimals = toToken?.decimals[toChainId]
  const toTokenAddress = toToken?.addresses[toChainId]

  const debouncedFromValueBigInt = stringToBigInt(
    debouncedFromValue,
    fromTokenDecimals
  )

  return {
    fromTokenSymbol,
    fromTokenDecimals,
    fromTokenAddress,
    toTokenSymbol,
    toTokenDecimals,
    toTokenAddress,
    debouncedFromValueBigInt,
  }
}
