import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import LoadingSpinner from '@tw/LoadingSpinner'
import {
  commify,
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'
import { stringToBigInt } from '@/utils/bigint/format'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

const PoolInfoSection = ({ chainId }: { chainId: number }) => {
  const { pool, poolData } = useSelector((state: RootState) => state.poolData)

  return (
    <div className="space-y-4">
      <CurrencyReservesCard />
      <InfoSectionCard title="Pool Info">
        <InfoListItem
          labelText="Trading Fee"
          content={
            poolData && poolData.swapFee ? (
              formatBigIntToPercentString(poolData.swapFee, 8, 2, false)
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
                content={formatBigIntToString(poolData.virtualPrice, 18, 6)}
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
            poolData && poolData?.totalLocked ? (
              <AugmentWithUnits
                content={commify(
                  formatBigIntToString(
                    stringToBigInt(
                      `${poolData.totalLocked}`,
                      pool.decimals[chainId]
                    ),
                    18,
                    -1
                  )
                )}
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
            poolData && poolData?.totalLockedUSD ? (
              `$${commify(
                formatBigIntToString(
                  stringToBigInt(
                    `${poolData.totalLockedUSD}`,
                    pool.decimals[chainId]
                  ),
                  18,
                  -1
                )
              )}`
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
