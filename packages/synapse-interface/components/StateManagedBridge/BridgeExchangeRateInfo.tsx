import numeral from 'numeral'
import Image from 'next/image'
import { useMemo } from 'react'
import { useAppSelector } from '@/store/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import {
  ELIGIBILITY_DEFAULT_TEXT,
  useStipEligibility,
} from '@/utils/hooks/useStipEligibility'
import { formatBigIntToString } from '@/utils/bigint/format'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'

const MAX_ARB_REBATE_PER_ADDRESS = 2000

const BridgeExchangeRateInfo = () => {
  /* TODO:
   * Upgrade to collapsable element
   * Use dark:border-zinc-800 in <section> className
   */

  return (
    <details open className="my-1 text-sm">
      {/* <RouteEligibility /> */}
      <summary className="block text-right px-1 mb-2 cursor-default pointer-events-none">
        <TimeEstimate />
      </summary>
      <section className="p-2 block rounded leading-relaxed border border-zinc-300 dark:border-separator">
        {' '}
        <GasDropLabel />
        <Router />
        {/* <Rebate /> */}
        <Slippage />
      </section>
    </details>
  )
}

const Slippage = () => {
  const {
    fromValue,
    bridgeQuote: { exchangeRate },
  } = useBridgeState()

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
  } = useBridgeState()
  return (
    <div className="flex justify-between">
      <span className="text-zinc-500 dark:text-zinc-400">Router</span>
      {bridgeModuleName}
    </div>
  )
}

const RouteEligibility = () => {
  const { isRouteEligible, isActiveRouteEligible, rebate } =
    useStipEligibility()

  const { parsedCumulativeRewards } = useAppSelector(
    (state) => state.feeAndRebate
  )

  if (
    !isRouteEligible ||
    !rebate ||
    Number(parsedCumulativeRewards) > MAX_ARB_REBATE_PER_ADDRESS
  ) {
    return (
      <div className="flex justify-between">
        <div className="flex-grow" />
        <TimeEstimate />
      </div>
    )
  }

  return (
    <div className="flex items-center justify-between">
      <div className="flex items-center">
        <Image
          src={CHAINS_BY_ID[CHAINS.ARBITRUM.id].chainImg}
          alt="To chain"
          className="w-4 h-4 mr-2 rounded-full"
        />

        <span className="text-green-300">
          {isActiveRouteEligible && rebate ? (
            <RebateText />
          ) : (
            ELIGIBILITY_DEFAULT_TEXT
          )}
        </span>
      </div>
      <TimeEstimate />
    </div>
  )
}

const RebateText = () => {
  const { rebate } = useStipEligibility()
  const { arbPrice } = useAppSelector((state) => state.priceData)
  const arbInDollars = rebate * arbPrice

  return (
    <div className="overflow-hidden whitespace-nowrap overflow-ellipsis">
      <span className="text-green-300">
        +{numeral(rebate).format('0,0.000')} ARB
      </span>
      <span className="text-secondary"> / </span>
      <span className="text-green-300">
        {numeral(arbInDollars).format('$0,0.00')}
      </span>
    </div>
  )
}

const Rebate = () => {
  const { isRouteEligible, rebate } = useStipEligibility()

  const { parsedCumulativeRewards } = useAppSelector(
    (state) => state.feeAndRebate
  )

  if (
    !isRouteEligible ||
    !rebate ||
    Number(parsedCumulativeRewards) > MAX_ARB_REBATE_PER_ADDRESS
  ) {
    return null
  }

  return (
    <div className="flex items-center justify-between">
      <div className="text-green-300">Rebate</div>
      <RebateText />
    </div>
  )
}

const TimeEstimate = () => {
  const { fromToken, bridgeQuote } = useBridgeState()

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
  const {
    bridgeQuote: { gasDropAmount },
    toChainId,
  } = useBridgeState()
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

export default BridgeExchangeRateInfo
