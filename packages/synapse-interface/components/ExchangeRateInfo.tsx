import { useState, useMemo, useEffect } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { useGasDropAmount } from '@/utils/hooks/useGasDropAmount'
import Image from 'next/image'

import { Token } from '@/utils/types'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'

const ExchangeRateInfo = ({
  fromAmount,
  toToken,
  exchangeRate,
  toChainId,
  showGasDrop,
}: {
  fromAmount: bigint
  toToken: Token
  exchangeRate: bigint
  toChainId: number
  showGasDrop: boolean
}) => {
  const [gasDropChainId, setGasDropChainId] = useState<number>(null)
  const { gasDrop: gasDropAmount, loading } = useGasDropAmount(toChainId)

  const safeExchangeRate = useMemo(() => exchangeRate ?? 0n, [exchangeRate]) // todo clean
  const safeFromAmount = useMemo(() => fromAmount ?? 0n, [fromAmount]) // todo clean
  const formattedExchangeRate = formatBigIntToString(safeExchangeRate, 18, 5)
  const numExchangeRate = Number(formattedExchangeRate)
  const slippage = safeExchangeRate - 1000000000000000000n
  const formattedPercentSlippage = formatBigIntToPercentString(slippage, 18)
  const underFee = safeExchangeRate === 0n && safeFromAmount != 0n

  const textColor: string = useMemo(() => {
    if (numExchangeRate >= 1) {
      return 'text-green-500'
    } else if (numExchangeRate > 0.975) {
      return 'text-amber-500'
    } else {
      return 'text-red-500'
    }
  }, [numExchangeRate])

  const isGasDropped = useMemo(() => {
    if (gasDropAmount) {
      return gasDropAmount.gt(0)
    }
  }, [gasDropAmount])

  useEffect(() => {
    setGasDropChainId(toChainId)
  }, [toChainId, isGasDropped])

  const memoizedGasDropLabel = useMemo(() => {
    if (!isGasDropped || !(toChainId == gasDropChainId)) return null
    if (loading) return null
    return <GasDropLabel gasDropAmount={gasDropAmount} toChainId={toChainId} />
  }, [toChainId, gasDropChainId, isGasDropped, loading])

  const expectedToChain = useMemo(() => {
    return toChainId && <ChainInfoLabel chainId={toChainId} />
  }, [toChainId])

  return (
    <div className="py-3.5 px-1 space-y-2 text-xs md:text-base lg:text-base md:px-6">
      {showGasDrop && (
        <div
          className={
            isGasDropped
              ? 'flex items-center justify-between'
              : 'flex justify-end'
          }
        >
          {memoizedGasDropLabel}
        </div>
      )}
      <div className="flex justify-between">
        <div className="flex space-x-2 text-[#88818C]">
          <p>Expected Price on</p>
          {expectedToChain}
        </div>
        <span className="text-[#88818C]">
          {safeFromAmount != 0n ? (
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
        {safeFromAmount != 0n && !underFee ? (
          <span className={` ${textColor}`}>{formattedPercentSlippage}</span>
        ) : (
          <span className="text-[#88818C]">—</span>
        )}
      </div>
    </div>
  )
}

const GasDropLabel = ({
  gasDropAmount,
  toChainId,
}: {
  gasDropAmount: bigint
  toChainId: number
}) => {
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

  const formattedGasDropAmount = formatBigIntToString(
    gasDropAmount,
    18,
    decimalsToDisplay
  )

  const airdropInDollars = getAirdropInDollars(symbol, formattedGasDropAmount)

  return (
    <div className="flex justify-between text-[#88818C]">
      <span className="text-[#88818C]">
        Will also receive {formattedGasDropAmount}{' '}
      </span>
      <span className="ml-1 font-medium text-white">
        {' '}
        {symbol}{' '}
        <span className="text-[#88818C] font-normal">
          {airdropInDollars && `($${airdropInDollars})`}
        </span>
      </span>
    </div>
  )
}

const ChainInfoLabel = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return chain ? (
    <span className="flex items-center space-x-1">
      <Image
        alt="chain image"
        src={chain.chainImg}
        className="w-4 h-4 rounded-full"
      />
      <span className="text-white">
        {chain.name.length > 10 ? chain.chainSymbol : chain.name}
      </span>
    </span>
  ) : null
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
export default ExchangeRateInfo
