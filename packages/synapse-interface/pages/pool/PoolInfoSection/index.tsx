import AugmentWithUnits from '../components/AugmentWithUnits'
import { Token } from '@types'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import LoadingSpinner from '@tw/LoadingSpinner'
import { useEffect, useState } from 'react'
import { getPoolFee } from '@utils/actions/getPoolFee'
import { getSwapDepositContractFields } from '@/utils/hooks/useSwapDepositContract'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'

const PoolInfoSection = ({
  pool,
  poolData,
  chainId,
}: {
  pool: Token
  poolData: any
  chainId: number
}) => {
  const [swapFee, setSwapFee] = useState<bigint>(0n)
  const { poolAddress } = getSwapDepositContractFields(pool, chainId)
  useEffect(() => {
    if (pool && chainId) {
      getPoolFee(poolAddress, chainId).then((res) => {
        console.log(`res, get pool fee`, res?.swapFee)
        setSwapFee(res?.swapFee)
      })
    }
  }, [pool, chainId])

  console.log(`sawpFee`, swapFee)
  console.log(`pool`, pool)
  console.log(`poolData`, poolData)
  return (
    <div className="space-y-4">
      <CurrencyReservesCard
        title="Currency Reserves"
        chainId={chainId}
        poolData={poolData}
      />
      <InfoSectionCard title="Pool Info">
        <InfoListItem
          labelText="Trading Fee"
          content={
            swapFee ? (
              // what decimals should this be?
              formatBigIntToPercentString(swapFee, 8, 2)
            ) : (
              <LoadingSpinner />
            )
          }
        />
        <InfoListItem
          labelText="Virtual Price"
          content={
            poolData?.virtualPriceStr ? (
              <AugmentWithUnits
                content={formatBigIntToString(
                  BigInt(poolData.virtualPrice),
                  18,
                  6
                )}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingSpinner />
            )
          }
        />
        <InfoListItem
          labelText="Total Liquidity"
          content={
            poolData?.totalLockedUSDStr ? (
              <AugmentWithUnits
                content={poolData.totalLockedUSDStr}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingSpinner />
            )
          }
        />
        <InfoListItem
          labelText="Total Liquidity USD"
          content={
            poolData?.totalLockedUSDStr ? (
              `$${poolData.totalLockedUSDStr}`
            ) : (
              <LoadingSpinner />
            )
          }
        />
      </InfoSectionCard>
    </div>
  )
}
export default PoolInfoSection

const InfoListItem = ({
  labelText,
  content,
  className = '',
}: {
  labelText: string
  content: any
  className?: string
}) => {
  return (
    <li
      className={`pl-3 pr-4 py-2 text-sm w-full flex border-gray-200 ${className}`}
    >
      <div className="text-white">{labelText} </div>
      <div className="self-center ml-auto text-white">{content}</div>
    </li>
  )
}
