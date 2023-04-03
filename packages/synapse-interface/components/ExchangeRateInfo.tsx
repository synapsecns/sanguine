import { useState } from 'react'
import { BigNumber } from '@ethersproject/bignumber'
import { formatBNToPercentString, formatBNToString } from '@bignumber/format'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import Image from 'next/image'
import { Token } from '@/utils/types'

export default function ExchangeRateInfo({
  fromAmount,
  toToken,
  exchangeRate,
  toChainId,
}: {
  fromAmount: BigNumber
  toToken: Token
  exchangeRate: BigNumber
  toChainId: number
}) {
  const [showExchangeRateInfo, toggleExchangeRateInfo] = useState(false)

  const formattedExchangeRate = formatBNToString(exchangeRate, 18, 4)

  // rewrite the below
  const numExchangeRate = Number(formattedExchangeRate)

  const slippage = exchangeRate.sub(BigNumber.from(10).pow(18))
  const formattedPercentSlippage = formatBNToPercentString(slippage, 18)
  const underFee = exchangeRate.eq(0) && !fromAmount.eq(0)

  let textColor
  if (numExchangeRate >= 1) {
    textColor = 'text-green-500'
  } else if (numExchangeRate > 0.975) {
    textColor = 'text-amber-500'
  } else {
    textColor = 'text-red-500'
  }

  const isGasDropped = exchangeRate.gt(0)

  return (
    <div className="py-3.5 px-1 space-y-2 text-xs md:text-base lg:text-base">
      <div
        className={
          isGasDropped
            ? 'flex items-center justify-between'
            : 'flex justify-end'
        }
      >
        {/*
        TODO need to add gas retrieval to sdk
        {isGasDropped && (
          <GasDropLabel gasDropAmount={gasDropAmount} toChainId={toChainId} />
        )} */}
      </div>
      <div className="flex justify-between">
        <div className="flex space-x-2 text-[#88818C]">
          <p>Expected Price on</p>
          {toChainId && <ChainInfoLabel chainId={toChainId} />}
        </div>
        <span className="text-[#88818C]">
          {!fromAmount.eq(0) ? (
            <>
              {formattedExchangeRate}{' '}
              <span className="text-white">{toToken.symbol}</span>
            </>
          ) : (
            '—'
          )}
        </span>
      </div>
      <div className="flex justify-between">
        <p className="text-[#88818C] ">Slippage</p>
        {!fromAmount.eq(0) && !underFee ? (
          <span className={` ${textColor}`}>{formattedPercentSlippage}</span>
        ) : (
          <span className="text-[#88818C]">—</span>
        )}
      </div>
    </div>
  )
}

function GasDropLabel({
  gasDropAmount,
  toChainId,
}: {
  gasDropAmount: BigNumber
  toChainId: number
}) {
  let decimalsToDisplay
  const symbol = CHAINS_BY_ID[toChainId].nativeCurrency.symbol

  if ([CHAINS.FANTOM.id].includes(toChainId)) {
    decimalsToDisplay = 2
  } else if (
    [CHAINS.BNB.id, CHAINS.AVALANCHE.id, CHAINS.BOBA.id].includes(toChainId)
  ) {
    decimalsToDisplay = 3
  } else {
    decimalsToDisplay = 4
  }

  const formattedGasDropAmount = formatBNToString(
    gasDropAmount,
    18,
    decimalsToDisplay
  )

  const airdropInDollars = getAirdropInDollars(symbol, formattedGasDropAmount)

  return (
    <div className="flex justify-between text-[#88818C]">
      <span className="text-[#88818C]">
        Will also receive {formattedGasDropAmount}{' '}
      </span>{' '}
      <span className="ml-1 font-medium text-white">
        {symbol}
        <span className="text-[#88818C] font-normal">
          {' '}
          {airdropInDollars && `($${airdropInDollars})`}
        </span>
      </span>
    </div>
  )
}

function ChainInfoLabel({ chainId }: { chainId: number }) {
  const chain = CHAINS_BY_ID[chainId]
  return chain ? (
    <span className="flex items-center space-x-1">
      <Image
        alt="chain image"
        src={chain.chainImg}
        className="w-4 h-4 rounded-full"
      />
      <span className="text-white">
        {chain.chainName.length > 10 ? chain.chainSymbol : chain.chainName}
      </span>
    </span>
  ) : null
}

function getAirdropInDollars(symbol: string, formattedGasDropAmount: string) {
  const price = useCoingeckoPrice(symbol)

  if (price) {
    const airdropInDollars = parseFloat(formattedGasDropAmount) * price

    return airdropInDollars.toFixed(2)
  } else {
    return undefined
  }
}
