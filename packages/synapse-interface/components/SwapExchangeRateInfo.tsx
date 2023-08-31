import { useMemo } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'

import { Token } from '@/utils/types'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'

const SwapExchangeRateInfo = ({
  fromAmount,
  toToken,
  exchangeRate,
  toChainId,
}: {
  fromAmount: bigint
  toToken: Token
  exchangeRate: bigint
  toChainId: number
}) => {
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

  const expectedToChain = useMemo(() => {
    return toChainId && <ChainInfoLabel chainId={toChainId} />
  }, [toChainId])

  return (
    <div className="py-3.5 px-1 space-y-2 text-xs md:text-base lg:text-base md:px-6">
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

export default SwapExchangeRateInfo
