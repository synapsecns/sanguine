import { useBridgeState } from '@/slices/bridge/hooks'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { getValidAddress, isValidAddress } from '@/utils/isValidAddress'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { getSignificantDecimals } from '@/utils/getSignificantDecimals'

export const BridgeExchangeRateInfo = () => {
  /* TODO:
   * Upgrade to collapsable element
   * Convert from div to details (conflict on mobile for details/summary)
   * Use dark:border-zinc-800 in <section> className
   */

  return (
    <div className="mt-1 mb-2 text-sm">
      <div className="block px-1 mb-2 text-right cursor-default pointer-events-none">
        <TimeEstimate />
      </div>
      <div className="block p-2 leading-relaxed border rounded border-zinc-300 dark:border-separator">
        {' '}
        <GasDropLabel />
        <Router />
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
  const t = useTranslations('Bridge')
  const {
    bridgeQuote: {
      inputAmountForQuote,
      outputAmount,
      originTokenForQuote,
      destTokenForQuote,
    },
  } = useBridgeQuoteState()

  // Parse input amount - convert decimal string to bigint
  const inputAmount =
    inputAmountForQuote &&
    inputAmountForQuote !== '0' &&
    originTokenForQuote &&
    fromChainId
      ? stringToBigInt(
          inputAmountForQuote,
          typeof originTokenForQuote.decimals === 'number'
            ? originTokenForQuote.decimals
            : originTokenForQuote.decimals[fromChainId]
        )
      : null

  // Calculate USD-based slippage
  const { slippage, isLoading, error, textColor } = useUsdSlippage({
    originToken: originTokenForQuote,
    destToken: destTokenForQuote,
    originChainId: fromChainId,
    destChainId: toChainId,
    inputAmount,
    outputAmount,
  })

  // Show content
  const shouldShow =
    debouncedFromValue !== '0' && inputAmount && inputAmount > 0n

  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">{t('Slippage')}</span>
      {shouldShow ? (
        <>
          {isLoading && <span className="text-zinc-400">Calculating...</span>}
          {!isLoading && error && (
            <span className="text-zinc-400">{error}</span>
          )}
          {!isLoading && !error && slippage !== null && (
            <span className={textColor}>{slippage.toFixed(2)}%</span>
          )}
        </>
      ) : (
        <span className="">−</span>
      )}
    </div>
  )
}

const Router = () => {
  const {
    bridgeQuote: { bridgeModuleName },
  } = useBridgeQuoteState()
  const t = useTranslations('Bridge')
  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">{t('Router')}</span>
      {bridgeModuleName}
    </div>
  )
}

const TimeEstimate = () => {
  const { fromToken } = useBridgeState()
  const { bridgeQuote } = useBridgeQuoteState()
  const t = useTranslations()

  let showText
  let showTime
  let timeUnit

  if (fromToken && bridgeQuote?.estimatedTime > 60) {
    showTime = bridgeQuote?.estimatedTime / 60
    timeUnit = t('Time.minutes')
    showText = `${showTime} ${timeUnit} via ${bridgeQuote.bridgeModuleName}`
  }

  if (fromToken && bridgeQuote.estimatedTime <= 60) {
    showTime = bridgeQuote?.estimatedTime
    timeUnit = t('Time.seconds')
    showText = `${showTime} ${timeUnit} via ${bridgeQuote.bridgeModuleName}`
  }

  if (
    !bridgeQuote ||
    bridgeQuote.outputAmount === EMPTY_BRIDGE_QUOTE.outputAmount
  ) {
    showText = (
      <span className="text-zinc-500 dark:text-zinc-400">
        {t('Bridge.Powered by Synapse')}
      </span>
    )
  }

  if (!fromToken) {
    showText = t('Bridge.Select origin token')
  }

  return showText
}

const GasDropLabel = () => {
  const { toChainId } = useBridgeState()
  const {
    bridgeQuote: { gasDropAmount },
  } = useBridgeQuoteState()

  const t = useTranslations('Bridge')
  const symbol = CHAINS_BY_ID[toChainId]?.nativeCurrency.symbol

  const stringifiedGasAmount = formatBigIntToString(gasDropAmount, 18)
  const significantDecimals = getSignificantDecimals(stringifiedGasAmount)

  const formattedGasDropAmount = formatBigIntToString(
    gasDropAmount,
    18,
    significantDecimals
  )

  const airdropInDollars = getAirdropInDollars(symbol, formattedGasDropAmount)

  if (gasDropAmount === EMPTY_BRIDGE_QUOTE.gasDropAmount) {
    return null
  }

  return (
    <>
      <span className="text-zinc-500 dark:text-zinc-400">
        {t('Will also receive')} {formattedGasDropAmount}
      </span>
      <span>
        {' '}
        {symbol} {airdropInDollars && `($${airdropInDollars})`}
      </span>
    </>
  )
}

const getAirdropInDollars = (
  symbol: string,
  formattedGasDropAmount: string
) => {
  const decimals = symbol === 'JEWEL' ? 4 : 2
  const price = useCoingeckoPrice(symbol)

  if (price) {
    const airdropInDollars = parseFloat(formattedGasDropAmount) * price

    return airdropInDollars.toFixed(decimals)
  } else {
    return undefined
  }
}
