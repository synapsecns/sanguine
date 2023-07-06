import AugmentWithUnits from '../components/AugmentWithUnits'
import { Token } from '@types'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import LoadingSpinner from '@tw/LoadingSpinner'
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
            poolData && poolData.swapFee ? (
              // what decimals should this be?
              formatBigIntToPercentString(poolData.swapFee, 8, 2)
            ) : (
              <LoadingSpinner />
            )
          }
        />
        <InfoListItem
          labelText="Virtual Price"
          content={
            poolData && poolData?.virtualPrice ? (
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
            poolData && poolData?.totalLockedUSDStr ? (
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
            poolData && poolData?.totalLockedUSDStr ? (
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

export default PoolInfoSection
