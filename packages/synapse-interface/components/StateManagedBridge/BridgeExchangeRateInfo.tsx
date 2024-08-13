import { useMemo } from 'react'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useAccount } from 'wagmi'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { formatBigIntToString } from '@/utils/bigint/format'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
import { getValidAddress, isValidAddress } from '@/utils/isValidAddress'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'

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

  const showAddress =
    destinationAddress &&
    getValidAddress(address) !== getValidAddress(destinationAddress)

  const isInputValidAddress: boolean = destinationAddress
    ? isValidAddress(destinationAddress)
    : false

  if (showAddress && isInputValidAddress) {
    return (
      <div className="flex items-center space-x-1">
        <div>To: </div>
        <div className="text-primary">{destinationAddress}</div>
      </div>
    )
  }
}

const Slippage = () => {
  const { fromValue } = useBridgeState()

  const {
    bridgeQuote: { exchangeRate },
  } = useBridgeQuoteState()

  const { formattedPercentSlippage, safeFromAmount, underFee, textColor } =
    useExchangeRateInfo(fromValue, exchangeRate)
  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">Slippage</span>
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
  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">Router</span>
      {bridgeModuleName}
    </div>
  )
}

const TimeEstimate = () => {
  const { fromToken } = useBridgeState()
  const { bridgeQuote } = useBridgeQuoteState()

  let showText
  let showTime
  let timeUnit

  if (fromToken && bridgeQuote?.estimatedTime > 60) {
    showTime = bridgeQuote?.estimatedTime / 60
    timeUnit = 'minutes'
    showText = `${showTime} ${timeUnit} via ${bridgeQuote.bridgeModuleName}`
  }

  if (fromToken && bridgeQuote.estimatedTime <= 60) {
    showTime = bridgeQuote?.estimatedTime
    timeUnit = 'seconds'
    showText = `${showTime} ${timeUnit} via ${bridgeQuote.bridgeModuleName}`
  }

  if (
    !bridgeQuote ||
    bridgeQuote.outputAmount === EMPTY_BRIDGE_QUOTE.outputAmount
  ) {
    showText = (
      <span className="text-zinc-500 dark:text-zinc-400">
        Powered by Synapse
      </span>
    )
  }

  if (!fromToken) {
    showText = `Select origin token`
  }

  return showText
}

const GasDropLabel = () => {
  let decimalsToDisplay
  const { toChainId } = useBridgeState()
  const {
    bridgeQuote: { gasDropAmount },
  } = useBridgeQuoteState()
  const symbol = CHAINS_BY_ID[toChainId]?.nativeCurrency.symbol

  if ([CHAINS.FANTOM.id].includes(toChainId)) {
    decimalsToDisplay = 2
  } else if (
    [CHAINS.BNB.id, CHAINS.AVALANCHE.id, CHAINS.BOBA.id].includes(toChainId)
  ) {
    decimalsToDisplay = 3
  } else {
    decimalsToDisplay = 4
  }

  const formattedGasDropAmount = formatBigIntToString(
    gasDropAmount,
    18,
    decimalsToDisplay
  )

  const airdropInDollars = getAirdropInDollars(symbol, formattedGasDropAmount)

  if (gasDropAmount === EMPTY_BRIDGE_QUOTE.gasDropAmount) {
    return null
  }

  return (
    <>
      <span className="text-zinc-500 dark:text-zinc-400">
        Will also receive {formattedGasDropAmount}
      </span>
      <span>
        {' '}
        {symbol} {airdropInDollars && `($${airdropInDollars})`}
      </span>
    </>
  )
}

const useExchangeRateInfo = (fromValue, exchangeRate) => {
  const safeExchangeRate = typeof exchangeRate === 'bigint' ? exchangeRate : 0n
  const safeFromAmount = fromValue ?? '0'

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
  const price = useCoingeckoPrice(symbol)

  if (price) {
    const airdropInDollars = parseFloat(formattedGasDropAmount) * price

    return airdropInDollars.toFixed(2)
  } else {
    return undefined
  }
}
