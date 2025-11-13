import { useMemo } from 'react'
import Image from 'next/image'
import { useTranslations } from 'next-intl'

import { CHAINS_BY_ID } from '@constants/chains'
import { Token } from '@/utils/types'
import { useUsdSlippage } from '@hooks/useUsdSlippage'
import { formatBigIntToString } from '@/utils/bigint/format'

const SwapExchangeRateInfo = ({
  fromAmount,
  fromToken,
  toToken,
  exchangeRate,
  toChainId,
  outputAmount,
}: {
  fromAmount: bigint
  fromToken: Token
  toToken: Token
  exchangeRate: bigint
  toChainId: number
  outputAmount: bigint
}) => {
  const safeExchangeRate = useMemo(() => exchangeRate ?? 0n, [exchangeRate])
  const safeFromAmount = useMemo(() => fromAmount ?? 0n, [fromAmount])
  const formattedExchangeRate = formatBigIntToString(safeExchangeRate, 18, 5)

  // Calculate USD-based slippage
  const { slippage, isLoading, error, textColor } = useUsdSlippage({
    originToken: fromToken,
    destToken: toToken,
    originChainId: toChainId, // Swap happens on same chain
    destChainId: toChainId,
    inputAmount: safeFromAmount > 0n ? safeFromAmount : null,
    outputAmount: outputAmount > 0n ? outputAmount : null,
  })

  const expectedToChain = useMemo(() => {
    return toChainId && <ChainInfoLabel chainId={toChainId} />
  }, [toChainId])

  return (
    <div className="mt-1 mb-2 text-sm">
      <div className="block p-2 leading-relaxed border rounded border-zinc-300 dark:border-separator">
        <ExpectedPrice
          expectedToChain={expectedToChain}
          safeFromAmount={safeFromAmount}
          formattedExchangeRate={formattedExchangeRate}
          toToken={toToken}
        />
        <Slippage
          safeFromAmount={safeFromAmount}
          slippage={slippage}
          isLoading={isLoading}
          error={error}
          textColor={textColor}
        />
      </div>
    </div>
  )
}

const ExpectedPrice = ({
  expectedToChain,
  safeFromAmount,
  formattedExchangeRate,
  toToken,
}) => {
  const t = useTranslations('Swap')

  return (
    <div className="flex justify-between">
      <div className="flex space-x-2 text-[#88818C]">
        <p>{t('Expected price on')}</p> {expectedToChain}
      </div>
      <span className="text-[#88818C]">
        {safeFromAmount != 0n ? (
          <>
            {formattedExchangeRate}{' '}
            <span className="text-white">{toToken?.symbol}</span>
          </>
        ) : (
          '—'
        )}
      </span>
    </div>
  )
}

interface SlippageProps {
  safeFromAmount: bigint
  slippage: number | null
  isLoading: boolean
  error: string | null
  textColor: string
}

const Slippage = ({
  safeFromAmount,
  slippage,
  isLoading,
  error,
  textColor,
}: SlippageProps) => {
  const t = useTranslations('Swap')

  const shouldShow = safeFromAmount > 0n

  return (
    <div className="flex justify-between">
      <p className="text-[#88818C] ">{t('Slippage')}</p>
      {shouldShow ? (
        <>
          {isLoading && <span className="text-[#88818C]">Calculating...</span>}
          {!isLoading && error && <span className="text-[#88818C]">{error}</span>}
          {!isLoading && !error && slippage !== null && (
            <span className={textColor}>
              {slippage >= 0 ? '+' : ''}
              {slippage.toFixed(2)}%
            </span>
          )}
          {!isLoading && !error && slippage === null && (
            <span className="text-[#88818C]">—</span>
          )}
        </>
      ) : (
        <span className="text-[#88818C]">—</span>
      )}
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
