import { useMemo } from 'react'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { formatBigIntToString } from '@/utils/bigint/format'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'

const BridgeExchangeRateInfo = () => {
  return (
    <div className="py-3.5 px-1 space-y-3 text-sm md:px-6 tracking-wide">
      <TimeEstimate />
      <section className="p-2 space-y-1 text-sm border rounded-sm border-[#504952] text-secondary font-light">
        <GasDropLabel />
        <Router />
        <Slippage />
      </section>
    </div>
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
      <div>Slippage</div>
      {safeFromAmount !== '0' && !underFee ? (
        <span className={textColor}>{formattedPercentSlippage}</span>
      ) : (
        <span className="">—</span>
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
      <div>Router</div>
      <div className="text-primaryTextColor">{bridgeModuleName}</div>
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
    showText = null
  }

  if (!fromToken) {
    showText = `Select origin token`
  }

  return <div className="text-right text-secondary">{showText}</div>
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
    <div className="flex items-center text-secondary">
      <span className="">Will also receive {formattedGasDropAmount} </span>
      <span className="ml-1 font-medium text-white">
        {symbol}{' '}
        <span className="font-normal ">
          {airdropInDollars && `($${airdropInDollars})`}
        </span>
      </span>
    </div>
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
      return 'text-green-300'
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
