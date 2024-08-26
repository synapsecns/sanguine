import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { stringToBigInt } from '@/utils/bigint/format'

export const useBridgeSelections = () => {
  const { fromChainId, fromToken, debouncedFromValue }: BridgeState =
    useBridgeState()
  const balances = usePortfolioBalances()

  const fromTokenDecimals = fromToken?.decimals[fromChainId]
  const fromTokenAddress = fromToken?.addresses[fromChainId]

  const fromChainBalances = balances[fromChainId]
  const fromTokenBalance = fromChainBalances?.find(
    (t) => t.tokenAddress === fromTokenAddress
  )?.balance

  const debouncedFromValueBigInt = stringToBigInt(
    debouncedFromValue,
    fromTokenDecimals
  )

  return {
    fromTokenBalance,
    fromTokenDecimals,
    fromTokenAddress,
    debouncedFromValueBigInt,
  }
}
