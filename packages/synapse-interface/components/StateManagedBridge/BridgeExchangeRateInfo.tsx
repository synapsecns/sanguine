import { useMemo } from 'react'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { formatBigIntToString } from '@/utils/bigint/format'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
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
  const { debouncedFromValue } = useBridgeState()
  const t = useTranslations('Bridge')

  const {
    bridgeQuote: { exchangeRate },
  } = useBridgeQuoteState()

  const { formattedPercentSlippage, safeFromAmount, underFee, textColor } =
    useExchangeRateInfo(debouncedFromValue, exchangeRate)
  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">{t('Slippage')}</span>
      {safeFromAmount !== '0' && !underFee ? (
        <span className={textColor}>{formattedPercentSlippage}</span>
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

const useExchangeRateInfo = (value, exchangeRate) => {
  const safeExchangeRate = typeof exchangeRate === 'bigint' ? exchangeRate : 0n
  const safeFromAmount = value ?? '0'

  const formattedExchangeRate = formatBigIntToString(safeExchangeRate, 18, 4)
  const numExchangeRate = Number(formattedExchangeRate)
  const slippage = safeExchangeRate - 1000000000000000000n
  const formattedPercentSlippage = formatBigIntToPercentString(slippage, 18)
  const underFee = safeExchangeRate === 0n && safeFromAmount !== '0'

  const textColor: string = useMemo(() => {
    if (numExchangeRate >= 1) {
      return 'text-green-500'
    } else if (numExchangeRate > 0.975) {
      return 'text-amber-500'
    } else {
      return 'text-red-500'
    }
  }, [numExchangeRate])

  return {
    formattedExchangeRate,
    formattedPercentSlippage,
    numExchangeRate,
    safeExchangeRate,
    safeFromAmount,
    slippage,
    underFee,
    textColor,
  }
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
