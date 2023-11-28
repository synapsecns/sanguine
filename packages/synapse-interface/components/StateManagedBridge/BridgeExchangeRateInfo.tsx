import { useState, useMemo, useEffect } from 'react'
import Image from 'next/image'
import { useCoingeckoPrice } from '@hooks/useCoingeckoPrice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { getEstimatedBridgeTimeInMinutes } from '@/utils/getEstimatedBridgeTime'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
import { useGasDropAmount } from '@/utils/hooks/useGasDropAmount'
import { formatBigIntToString } from '@/utils/bigint/format'
import { Chain } from '@/utils/types'
import { CHAINS_BY_ID } from '@constants/chains'
import * as CHAINS from '@constants/chains/master'
import { PieChart } from '../PieChart'

const BridgeExchangeRateInfo = ({ showGasDrop }: { showGasDrop: boolean }) => {
  const [gasDropChainId, setGasDropChainId] = useState<number>(null)

  const {
    fromValue: fromAmount,
    fromChainId,
    toChainId,
    toToken,
    bridgeQuote,
  }: BridgeState = useBridgeState()

  const bridgeOriginChain: Chain = CHAINS_BY_ID[fromChainId]
  const exchangeRate: bigint = bridgeQuote?.exchangeRate
  const bridgeModuleName: string = bridgeQuote?.bridgeModuleName

  const { gasDrop: gasDropAmount, loading } = useGasDropAmount(toChainId)

  const safeExchangeRate = typeof exchangeRate === 'bigint' ? exchangeRate : 0n
  const safeFromAmount = fromAmount ?? '0'

  const formattedExchangeRate = formatBigIntToString(safeExchangeRate, 18, 4)
  const numExchangeRate = Number(formattedExchangeRate)
  const slippage = safeExchangeRate - 1000000000000000000n
  const formattedPercentSlippage = formatBigIntToPercentString(slippage, 18)
  const underFee = safeExchangeRate === 0n && safeFromAmount != '0'

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
    <div className="py-3.5 px-1 space-y-2 text-sm md:text-base md:px-6">
      <div className="flex justify-between">
        <p className="text-[#88818C] ">Est. time</p>
        {fromChainId && toChainId ? (
          <div className="inline-flex space-x-1">
            <span className="text-[#88818C]">
              {getEstimatedBridgeTimeInMinutes({
                bridgeOriginChain,
                bridgeModuleName,
              })}{' '}
              min
            </span>

            <PieChart
              activeAmount={getEstimatedBridgeTimeInMinutes({
                bridgeOriginChain,
                bridgeModuleName,
              })}
              totalAmount={60}
            />
          </div>
        ) : (
          <span className="text-[#88818C]">—</span>
        )}
      </div>

      <DetailRow
        leftContent={
          <div className="flex space-x-1">
            <p>Expected Price on</p>
            {expectedToChain}
          </div>
        }
        rightContent={
          <span className="text-primary">
            {safeFromAmount != '0' ? (
              <>
                {formattedExchangeRate}{' '}
                <span className="text-secondary">{toToken?.symbol}</span>
              </>
            ) : (
              '—'
            )}
          </span>
        }
      />

      {/* <div className="flex justify-between">
        <div className="flex space-x-2 text-[#88818C]">
          <p>Expected Price on</p>
          {expectedToChain}
        </div>
        <span className="text-[#88818C]">
          {safeFromAmount != '0' ? (
            <>
              {formattedExchangeRate}{' '}
              <span className="text-white">{toToken?.symbol}</span>
            </>
          ) : (
            '—'
          )}
        </span>
      </div> */}

      {isGasDropped && (
        <div className="flex justify-between">
          <div className="flex space-x-2 text-[#88818C]">
            <p>Gas bonus</p>
          </div>
          {memoizedGasDropLabel}
        </div>
      )}

      <div className="flex justify-between">
        <p className="text-[#88818C]">Slippage</p>
        {safeFromAmount != '0' && !underFee ? (
          <span className={` ${textColor}`}>{formattedPercentSlippage}</span>
        ) : (
          <span className="text-[#88818C]">—</span>
        )}
      </div>
    </div>
  )
}

export const DetailRow = ({
  leftContent,
  rightContent,
}: {
  leftContent: any
  rightContent: any
}) => {
  return (
    <div className="flex justify-between text-secondary">
      <div>{leftContent}</div>
      <div>{rightContent}</div>
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
  const chain: Chain = CHAINS_BY_ID[toChainId]
  const symbol: string = chain?.nativeCurrency.symbol
  const icon = chain?.chainImg

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

  // @TO-DO: Remove commented out lines when new gas label approved
  return (
    // <div className="flex justify-between text-[#88818C]">
    //   <span className="text-[#88818C]">
    //     Will also receive {formattedGasDropAmount}{' '}
    //   </span>
    //   <span className="ml-1 font-medium text-white">
    //     {' '}
    //     {symbol}{' '}
    //     <span className="text-[#88818C] font-normal">
    //       {airdropInDollars && `($${airdropInDollars})`}
    //     </span>
    //   </span>
    // </div>
    <div
      data-test-id="gas-drop-label"
      className="flex justify-between space-x-1"
    >
      <span className="text-primary">+{formattedGasDropAmount}</span>
      <span className="text-secondary">{symbol}</span>
      <Image
        loading="lazy"
        alt={`${symbol} img`}
        className="w-4 h-4 m-auto rounded-md"
        src={icon}
      />
    </div>
  )
}

const ChainInfoLabel = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return chain ? (
    <span className="flex items-center space-x-1">
      <Image
        alt="chain image"
        src={chain?.chainImg}
        className="w-4 h-4 rounded-full"
      />
      <span className="text-white">
        {chain?.name?.length > 10 ? chain?.chainSymbol : chain?.name}
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

export default BridgeExchangeRateInfo
