import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { useSwapState } from '@/slices/swap/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { setSwapToToken } from '@/slices/swap/reducer'
import { useSwapToTokenListArray } from './hooks/useSwapToTokenListArray'
import { AmountInput } from '@/components/ui/AmountInput'
import { useWalletState } from '@/slices/wallet/hooks'
import { useUsdDisplay } from '@hooks/useUsdDisplay'
import { formatInlineUsdDifference } from '@utils/calculateUsdValue'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { formatAmount } from '@/utils/formatAmount'
import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { stringToBigInt } from '@/utils/bigint/format'

export const SwapOutputContainer = () => {
  const t = useTranslations('Swap')
  const { isConnected } = useAccount()
  const {
    swapQuote,
    isLoading,
    swapChainId,
    swapToToken,
    swapFromToken,
    swapFromValue,
  } = useSwapState()

  const showValue =
    swapQuote.outputAmountString === '0' ? '' : swapQuote.outputAmountString

  // Fetch token price and calculate USD value
  const usdValue = useUsdDisplay(swapToToken, showValue, swapChainId)

  // Convert input amount to bigint for slippage calculation
  const inputAmount =
    swapFromToken && swapFromValue && swapFromValue !== '0'
      ? stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId])
      : null

  // Calculate USD-based slippage to get USD difference
  const { usdDifference } = useUsdSlippage({
    originToken: swapFromToken,
    destToken: swapToToken,
    originChainId: swapChainId,
    destChainId: swapChainId,
    inputAmount: inputAmount && inputAmount > 0n ? inputAmount : null,
    outputAmount:
      swapQuote.outputAmount && swapQuote.outputAmount > 0n
        ? swapQuote.outputAmount
        : null,
  })

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
  const formattedUsdValue = `${usdValue}${formatInlineUsdDifference(usdDifference)}`

  return (
    <BridgeSectionContainer>
      <BridgeAmountContainer>
        <SwapToTokenSelector />
        <div className="flex flex-col w-full">
          <AmountInput
            disabled={true}
            showValue={showValue}
            isLoading={isLoading}
          />
          <div className="flex justify-between items-center">
            <div className="text-xs text-zinc-500 dark:text-zinc-400">
              {!isLoading && formattedUsdValue}
            </div>
            {isConnected && (
              <div className="text-xs text-zinc-500 dark:text-zinc-400">
                {t('Balance')}: {formattedBalance ?? '0.0'}
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
