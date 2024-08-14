import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { stringToBigInt } from '@/utils/bigint/format'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'

export const useBridgeSelections = () => {
  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    debouncedFromValue,
  }: BridgeState = useBridgeState()
  const balances = usePortfolioBalances()

  const fromTokenSymbol = fromToken?.symbol
  const fromTokenDecimals = fromToken?.decimals[fromChainId]
  const fromTokenAddress = fromToken?.addresses[fromChainId]

  const fromChainBalances = balances[fromChainId]
  const fromTokenBalance = fromChainBalances?.find(
    (t) => t.tokenAddress === fromTokenAddress
  )?.balance

  const toTokenSymbol = toToken?.symbol
  const toTokenDecimals = toToken?.decimals[toChainId]
  const toTokenAddress = toToken?.addresses[toChainId]

  const debouncedFromValueBigInt = stringToBigInt(
    debouncedFromValue,
    fromTokenDecimals
  )

  return {
    fromTokenBalance,
    fromTokenSymbol,
    fromTokenDecimals,
    fromTokenAddress,
    toTokenSymbol,
    toTokenDecimals,
    toTokenAddress,
    debouncedFromValueBigInt,
  }
}
