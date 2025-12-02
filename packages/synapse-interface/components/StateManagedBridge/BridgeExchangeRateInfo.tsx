import { useBridgeState } from '@/slices/bridge/hooks'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { getValidAddress, isValidAddress } from '@/utils/isValidAddress'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { DATA_PLACEHOLDER } from '@/constants/placeholders'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { formatSlippage } from '@/utils/formatSlippage'
import { parseTokenAmount } from '@/utils/decimals'

export const BridgeExchangeRateInfo = () => {
  /* TODO:
   * Upgrade to collapsable element
   * Convert from div to details (conflict on mobile for details/summary)
   * Use dark:border-zinc-800 in <section> className
   */

  return (
    <div className="mt-1 mb-2 text-sm">
      <div className="block p-2 leading-relaxed border rounded border-zinc-300 dark:border-separator">
        <Router />
        <EstimatedTime />
        <Slippage />
        <DestinationAddress />
      </div>
    </div>
  )
}

const DestinationAddress = () => {
  const { address } = useAccount()
  const { destinationAddress } = useBridgeState()

  const t = useTranslations('Bridge')

  const showAddress =
    destinationAddress &&
    getValidAddress(address) !== getValidAddress(destinationAddress)

  const isInputValidAddress: boolean = destinationAddress
    ? isValidAddress(destinationAddress)
    : false

  if (showAddress && isInputValidAddress) {
    return (
      <div className="flex items-center space-x-1">
        <div>{t('To')}: </div>
        <div className="text-primary">{destinationAddress}</div>
      </div>
    )
  }
}

const Slippage = () => {
  const { debouncedFromValue, fromChainId, toChainId } = useBridgeState()
  const t = useTranslations('Slippage')
  const {
    bridgeQuote: {
      inputAmountForQuote,
      outputAmount,
      originTokenForQuote,
      destTokenForQuote,
      formattedGasDrop,
      formattedNativeFee,
    },
    isLoading: isQuoteLoading,
  } = useBridgeQuoteState()

  // Parse input amount - convert decimal string to bigint
  const inputAmount = parseTokenAmount(
    inputAmountForQuote,
    originTokenForQuote,
    fromChainId
  )

  // Calculate USD-based slippage
  const {
    slippage,
    isLoading: isSlippageLoading,
    error,
    textColor,
  } = useUsdSlippage({
    originToken: originTokenForQuote,
    destToken: destTokenForQuote,
    originChainId: fromChainId,
    destChainId: toChainId,
    inputAmount,
    outputAmount,
    formattedGasDrop,
    formattedNativeFee,
  })

  // Show content
  const isLoading = isSlippageLoading || isQuoteLoading
  const shouldShow =
    !isLoading &&
    debouncedFromValue !== '0' &&
    inputAmount &&
    inputAmount > 0n &&
    outputAmount &&
    outputAmount > 0n

  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">{t('Slippage')}</span>
      {shouldShow ? (
        <>
          {error && (
            <span className="text-zinc-400">{t(error)}</span>
          )}
          {!error && slippage !== null && (
            <span className={textColor}>{formatSlippage(slippage)}</span>
          )}
          {!error && slippage === null && (
            DATA_PLACEHOLDER
          )}
        </>
      ) : (
        DATA_PLACEHOLDER
      )}
    </div>
  )
}

const Router = () => {
  const {
    bridgeQuote: { bridgeModuleName },
    isLoading,
  } = useBridgeQuoteState()
  const t = useTranslations('Bridge')
  const shouldShow = !isLoading && bridgeModuleName
  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">{t('Router')}</span>
      {shouldShow ? bridgeModuleName : DATA_PLACEHOLDER}
    </div>
  )
}

const EstimatedTime = () => {
  const { fromToken } = useBridgeState()
  const { bridgeQuote, isLoading } = useBridgeQuoteState()
  const t = useTranslations('Time')

  const shouldShow =
    !isLoading &&
    fromToken &&
    bridgeQuote &&
    bridgeQuote.outputAmount !== EMPTY_BRIDGE_QUOTE.outputAmount &&
    typeof bridgeQuote.estimatedTime === 'number' &&
    Number.isFinite(bridgeQuote.estimatedTime)

  let timeValue: number
  let timeUnit: string

  if (shouldShow) {
    if (bridgeQuote.estimatedTime > 60) {
      timeValue = bridgeQuote.estimatedTime / 60
      timeUnit = t('minutes')
    } else {
      timeValue = bridgeQuote.estimatedTime
      timeUnit = t('seconds')
    }
  }

  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">
        {t('Estimated Time')}
      </span>
      {shouldShow ? (
        <span>
          {timeValue} {timeUnit}
        </span>
      ) : (
        DATA_PLACEHOLDER
      )}
    </div>
  )
}

