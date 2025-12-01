import { useAccount } from 'wagmi'
import { useMemo } from 'react'

import { ChainSelector } from '@/components/ui/ChainSelector'
import { GasInfoBadge } from '@/components/StateManagedBridge/GasInfoBadge'
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
import { useDefiLlamaPrice } from '@hooks/useDefiLlamaPrice'
import {
  formatInlineUsdDifference,
  formatUsdValue,
  formatUsdBreakdownTooltip,
} from '@utils/calculateUsdValue'
import { HoverTooltip } from '@/components/HoverTooltip'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { formatAmount, formatAmountByPrice, getTooltipValue } from '@/utils/formatAmount'
import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { getTokenDecimals, parseTokenAmount } from '@/utils/decimals'

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

  const bridgeQuoteValue = useMemo(() => {
    if (!hasValidInput) {
      return ''
    } else if (hasValidQuote) {
      return bridgeQuote?.outputAmountString
    } else {
      return ''
    }
  }, [bridgeQuote, hasValidInput, hasValidQuote])

  const inputClassName = isQuoteStale ? 'opacity-50' : undefined

  const outputValue =
    fromChainId === ARBITRUM.id && toChainId === HYPERLIQUID.id
      ? debouncedFromValue
      : bridgeQuoteValue
  // Format output amount based on price (each decimal digit = $0.01)
  const toTokenPrice = useDefiLlamaPrice(toToken)
  const showValue = formatAmountByPrice(
    outputValue ?? '',
    toTokenPrice
  )
  const tooltipValue = getTooltipValue(showValue, outputValue ?? '', toToken?.symbol)

  // Convert input amount to bigint for slippage calculation
  const inputAmount = parseTokenAmount(
    bridgeQuote.inputAmountForQuote,
    bridgeQuote.originTokenForQuote,
    fromChainId
  )

  // Calculate USD-based slippage to get USD difference
  const { valueOut, gasDropUsd, usdDifference } = useUsdSlippage({
    originToken: bridgeQuote.originTokenForQuote,
    destToken: bridgeQuote.destTokenForQuote,
    originChainId: fromChainId,
    destChainId: toChainId,
    inputAmount,
    outputAmount: bridgeQuote.outputAmount,
    formattedGasDrop: bridgeQuote.formattedGasDrop,
    formattedNativeFee: bridgeQuote.formattedNativeFee,
  })

  // Get destination token balance
  const balances = usePortfolioBalances()
  const toTokenAddress = toToken?.addresses[toChainId]
  const toChainBalances = balances[toChainId]
  const toTokenBalance = toChainBalances?.find(
    (t) => t.tokenAddress === toTokenAddress
  )?.balance
  const toTokenDecimals = getTokenDecimals(toToken, toChainId)
  const parsedBalance =
    toTokenBalance !== undefined && toTokenDecimals !== undefined
      ? getParsedBalance(toTokenBalance, toTokenDecimals)
      : '0.0'
  const formattedBalance = formatAmount(parsedBalance)
  const formattedUsdValue = `${formatUsdValue(valueOut)}${formatInlineUsdDifference(usdDifference)}`
  const usdBreakdown = formatUsdBreakdownTooltip(gasDropUsd, valueOut, t('gas airdrop'))
  const toChainNativeSymbol = CHAINS_BY_ID[toChainId]?.nativeCurrency.symbol

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
          <div className="flex justify-between items-center">
            <AmountInput
              disabled={true}
              showValue={showValue}
              isLoading={isLoading}
              className={inputClassName}
              tooltipValue={tooltipValue}
            />
            {bridgeQuote.formattedGasDrop && !isLoading && (
              <GasInfoBadge
                amount={bridgeQuote.formattedGasDrop}
                symbol={toChainNativeSymbol}
                tooltipText={`${bridgeQuote.bridgeModuleName} ${t('gas airdrop')}`}
              />
            )}
          </div>
          <div className="flex justify-between items-center">
            <HoverTooltip
              hoverContent={usdBreakdown}
              position="bottom"
            >
              <div className="text-xs text-zinc-500 dark:text-zinc-400">
                {!isLoading && formattedUsdValue}
              </div>
            </HoverTooltip>
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
