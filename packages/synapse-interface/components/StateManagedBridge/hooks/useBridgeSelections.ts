import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { stringToBigInt } from '@/utils/bigint/format'
import { CHAINS_BY_ID } from '@/constants/chains'

export const useBridgeSelections = () => {
  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    debouncedFromValue,
  }: BridgeState = useBridgeState()
  const balances = usePortfolioBalances()

  const fromChain = CHAINS_BY_ID[fromChainId]
  const fromChainName = fromChain?.name

  const fromTokenSymbol = fromToken?.symbol
  const fromTokenDecimals = fromToken?.decimals[fromChainId]
  const fromTokenAddress = fromToken?.addresses[fromChainId]

  const fromChainBalances = balances[fromChainId]
  const fromTokenBalance = fromChainBalances?.find(
    (t) => t.tokenAddress === fromTokenAddress
  )?.balance

  const toChain = CHAINS_BY_ID[toChainId]
  const toChainName = toChain?.name

  const toTokenSymbol = toToken?.symbol
  const toTokenDecimals = toToken?.decimals[toChainId]
  const toTokenAddress = toToken?.addresses[toChainId]

  const debouncedFromValueBigInt = stringToBigInt(
    debouncedFromValue,
    fromTokenDecimals
  )

  return {
    fromTokenBalance,
    fromChain,
    fromChainName,
    fromTokenSymbol,
    fromTokenDecimals,
    fromTokenAddress,
    toChain,
    toChainName,
    toTokenSymbol,
    toTokenDecimals,
    toTokenAddress,
    debouncedFromValueBigInt,
  }
}
