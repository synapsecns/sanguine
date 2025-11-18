import { useBridgeState } from '@/slices/bridge/hooks'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { getValidAddress, isValidAddress } from '@/utils/isValidAddress'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { CHAINS_BY_ID } from '@constants/chains'
import {DATA_PLACEHOLDER} from '@/constants/placeholders'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { getSignificantDecimals } from '@/utils/getSignificantDecimals'
import { formatSlippage } from '@/utils/formatSlippage'
import { useDefiLlamaPrice } from '@/utils/hooks/useDefiLlamaPrice'
import { zeroAddress } from 'viem'
import { calculateUsdValue } from '@/utils/calculateUsdValue'

export const BridgeExchangeRateInfo = () => {
  /* TODO:
   * Upgrade to collapsable element
   * Convert from div to details (conflict on mobile for details/summary)
   * Use dark:border-zinc-800 in <section> className
   */

  return (
    <div className="mt-1 mb-2 text-sm">
      <div className="block p-2 leading-relaxed border rounded border-zinc-300 dark:border-separator">
        {' '}
        <GasDropLabel />
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
    },
    isLoading: isQuoteLoading,
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

const GasDropLabel = () => {
  const { toChainId } = useBridgeState()
  const {
    bridgeQuote: { gasDropAmount },
    isLoading,
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

  const airdropInDollars = getAirdropInDollars(toChainId, formattedGasDropAmount)

  if (isLoading || gasDropAmount === EMPTY_BRIDGE_QUOTE.gasDropAmount) {
    return null
  }

  return (
    <>
      <span className="text-zinc-500 dark:text-zinc-400">
        {t('Will also receive')} {formattedGasDropAmount}
      </span>
      <span>
        {' '}
        {symbol} {airdropInDollars && `(${airdropInDollars})`}
      </span>
    </>
  )
}

const getAirdropInDollars = (
  chainId: number,
  formattedGasDropAmount: string
) => {
  const price = useDefiLlamaPrice({
    addresses: {
      [chainId]: zeroAddress
    }
  })

  if (price) {
    return calculateUsdValue(formattedGasDropAmount, price)
  } else {
    return undefined
  }
}
