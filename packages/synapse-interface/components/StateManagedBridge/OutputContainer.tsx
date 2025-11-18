import { useAccount } from 'wagmi'
import { useMemo } from 'react'

import { ChainSelector } from '@/components/ui/ChainSelector'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { AmountInput } from '@/components/ui/AmountInput'
import { useToChainListArray } from '@/components/StateManagedBridge/hooks/useToChainListArray'
import { useToTokenListArray } from '@/components/StateManagedBridge/hooks/useToTokenListArray'
import { DestinationAddressInput } from '@/components/StateManagedBridge/DestinationAddressInput'
import { CHAINS_BY_ID } from '@/constants/chains'
import { setToChainId, setToToken } from '@/slices/bridge/reducer'
import { useBridgeDisplayState, useBridgeState } from '@/slices/bridge/hooks'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { useTranslations } from 'next-intl'
import { ARBITRUM, HYPERLIQUID } from '@/constants/chains/master'
import { useUsdDisplay } from '@hooks/useUsdDisplay'
import { formatInlineUsdDifference } from '@utils/calculateUsdValue'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { formatAmount } from '@/utils/formatAmount'
import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { stringToBigInt } from '@/utils/bigint/format'

interface OutputContainerProps {
  isQuoteStale: boolean
}

export const OutputContainer = ({ isQuoteStale }: OutputContainerProps) => {
  const t = useTranslations('Bridge')
  const { address, isConnected } = useAccount()
  const { bridgeQuote, isLoading } = useBridgeQuoteState()
  const { showDestinationAddress } = useBridgeDisplayState()
  const { hasValidInput, hasValidQuote } = useBridgeValidations()
  const { debouncedFromValue, fromChainId, toChainId, toToken } =
    useBridgeState()

  const showValue = useMemo(() => {
    if (!hasValidInput) {
      return ''
    } else if (hasValidQuote) {
      return bridgeQuote?.outputAmountString
    } else {
      return ''
    }
  }, [bridgeQuote, hasValidInput, hasValidQuote])

  const inputClassName = isQuoteStale ? 'opacity-50' : undefined

  // Fetch token price and calculate USD value
  const outputValue =
    fromChainId === ARBITRUM.id && toChainId === HYPERLIQUID.id
      ? debouncedFromValue
      : showValue
  const usdValue = useUsdDisplay(toToken, outputValue)

  // Convert input amount to bigint for slippage calculation
  const inputAmount =
    bridgeQuote.inputAmountForQuote &&
    bridgeQuote.inputAmountForQuote !== '0' &&
    bridgeQuote.originTokenForQuote &&
    fromChainId
      ? stringToBigInt(
          bridgeQuote.inputAmountForQuote,
          typeof bridgeQuote.originTokenForQuote.decimals === 'number'
            ? bridgeQuote.originTokenForQuote.decimals
            : bridgeQuote.originTokenForQuote.decimals[fromChainId]
        )
      : null

  // Calculate USD-based slippage to get USD difference
  const { usdDifference } = useUsdSlippage({
    originToken: bridgeQuote.originTokenForQuote,
    destToken: bridgeQuote.destTokenForQuote,
    originChainId: fromChainId,
    destChainId: toChainId,
    inputAmount,
    outputAmount: bridgeQuote.outputAmount,
  })

  // Get destination token balance
  const balances = usePortfolioBalances()
  const toTokenAddress = toToken?.addresses[toChainId]
  const toChainBalances = balances[toChainId]
  const toTokenBalance = toChainBalances?.find(
    (t) => t.tokenAddress === toTokenAddress
  )?.balance
  const toTokenDecimals =
    typeof toToken?.decimals === 'number'
      ? toToken.decimals
      : toToken?.decimals?.[toChainId]
  const parsedBalance =
    toTokenBalance !== undefined && toTokenDecimals !== undefined
      ? getParsedBalance(toTokenBalance, toTokenDecimals)
      : '0.0'
  const formattedBalance = formatAmount(parsedBalance)
  const formattedUsdValue = `${usdValue}${formatInlineUsdDifference(usdDifference)}`

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <ToChainSelector />
        {showDestinationAddress && toChainId !== HYPERLIQUID.id ? (
          <DestinationAddressInput connectedAddress={address} />
        ) : null}
      </div>

      <BridgeAmountContainer>
        <ToTokenSelector />
        <div className="flex flex-col w-full">
          <AmountInput
            disabled={true}
            showValue={
              fromChainId === ARBITRUM.id && toChainId === HYPERLIQUID.id
                ? debouncedFromValue
                : showValue
            }
            isLoading={isLoading}
            className={inputClassName}
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

const ToChainSelector = () => {
  const { toChainId } = useBridgeState()
  const { isWalletPending } = useWalletState()

  const t = useTranslations('Bridge')

  return (
    <ChainSelector
      dataTestId="bridge-destination-chain"
      isOrigin={false}
      selectedItem={CHAINS_BY_ID[toChainId]}
      label={t('To')}
      itemListFunction={useToChainListArray}
      setFunction={setToChainId}
      action="Bridge"
      disabled={isWalletPending}
    />
  )
}

const ToTokenSelector = () => {
  const { toToken } = useBridgeState()
  const { isWalletPending } = useWalletState()
  const t = useTranslations('Bridge')

  return (
    <TokenSelector
      dataTestId="bridge-destination-token"
      isOrigin={false}
      selectedItem={toToken}
      placeholder={t('In')}
      itemListFunction={useToTokenListArray}
      setFunction={setToToken}
      action="Bridge"
      disabled={isWalletPending}
    />
  )
}
