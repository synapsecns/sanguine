import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'

import type { Token } from '@/utils/types'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'

export const SwapExchangeRateInfo = ({
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
  const safeExchangeRate = exchangeRate ?? 0n
  const safeFromAmount = fromAmount ?? 0n
  const formattedExchangeRate = formatBigIntToString(safeExchangeRate, 18, 5)
  const numExchangeRate = Number(formattedExchangeRate)
  const slippage = safeExchangeRate - 1000000000000000000n
  const formattedPercentSlippage = formatBigIntToPercentString(slippage, 18)
  const underFee = safeExchangeRate === 0n && safeFromAmount != 0n


  let textColor: string
  if (numExchangeRate >= 1) {
    textColor = 'text-green-500'
  } else if (numExchangeRate > 0.975) {
    textColor = 'text-amber-500'
  } else {
    textColor = 'text-red-500'
  }

  return (
    <div className="py-3.5 px-1 space-y-2 text-xs md:text-base lg:text-base md:px-6">
      <div className="flex justify-between text-white/50">
        <div className="flex space-x-2 ">
          <p>Expected Price on</p>
          {toChainId && <ChainInfoLabel chainId={toChainId} />}
        </div>
        <span className="">
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
        <p className="text-white/50">Slippage</p>
        {safeFromAmount != 0n && !underFee ? (
          <span className={textColor}>{formattedPercentSlippage}</span>
        ) : (
          <span className="text-white/50">—</span>
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


