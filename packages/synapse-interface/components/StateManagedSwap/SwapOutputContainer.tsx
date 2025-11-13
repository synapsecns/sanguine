import { useAccount } from 'wagmi'

import { useSwapState } from '@/slices/swap/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { setSwapToToken } from '@/slices/swap/reducer'
import { useSwapToTokenListArray } from './hooks/useSwapToTokenListArray'
import { AmountInput } from '@/components/ui/AmountInput'
import { useWalletState } from '@/slices/wallet/hooks'
import { useDefiLlamaPrice } from '@hooks/useDefiLlamaPrice'
import { calculateUsdValue } from '@utils/calculateUsdValue'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { formatAmount } from '@/utils/formatAmount'

export const SwapOutputContainer = () => {
  const { isConnected } = useAccount()
  const { swapQuote, isLoading, swapChainId, swapToToken } = useSwapState()

  const showValue =
    swapQuote.outputAmountString === '0' ? '' : swapQuote.outputAmountString

  // Fetch token price and calculate USD value
  const swapToTokenPrice = useDefiLlamaPrice(swapToToken, swapChainId)
  const usdValue = calculateUsdValue(showValue, swapToTokenPrice)

  // Get output token balance
  const { balances } = usePortfolioState()
  const tokenData = balances[swapChainId]?.find(
    (token) => token.tokenAddress === swapToToken?.addresses[swapChainId]
  )
  const balance = tokenData?.balance
  const decimals = tokenData?.token?.decimals[swapChainId]
  const parsedBalance =
    balance !== undefined && decimals !== undefined
      ? getParsedBalance(balance, decimals)
      : '0.0'
  const formattedBalance = formatAmount(parsedBalance)

  return (
    <BridgeSectionContainer>
      <BridgeAmountContainer>
        <SwapToTokenSelector />
        <div className="flex flex-col w-full gap-1">
          <AmountInput
            disabled={true}
            showValue={showValue}
            isLoading={isLoading}
          />
          <div className="flex justify-between items-center">
            <div className="text-xs text-zinc-500 dark:text-zinc-400">
              {usdValue}
            </div>
            {isConnected && (
              <div className="text-xs text-zinc-500 dark:text-zinc-400">
                Balance: {formattedBalance ?? '0.0'}
              </div>
            )}
          </div>
        </div>
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const SwapToTokenSelector = () => {
  const { swapToToken } = useSwapState()
  const { isWalletPending } = useWalletState()

  return (
    <TokenSelector
      dataTestId="swap-destination-token"
      isOrigin={false}
      selectedItem={swapToToken}
      placeholder="Out"
      itemListFunction={useSwapToTokenListArray}
      setFunction={setSwapToToken}
      action="Swap"
      disabled={isWalletPending}
    />
  )
}
