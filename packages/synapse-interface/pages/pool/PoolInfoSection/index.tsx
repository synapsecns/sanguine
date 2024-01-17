import numeral from 'numeral'
import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

const PoolInfoSection = () => {
  const { pool, poolData } = useSelector((state: RootState) => state.poolData)

  const usdFormat = poolData.totalLockedUSD > 1000000 ? '$0,0.0' : '$0,0'

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
              <LoadingDots />
            )
          }
        />
        <InfoListItem
          labelText="Virtual Price"
          content={
            poolData && poolData?.virtualPrice ? (
              <AugmentWithUnits
                content={formatBigIntToString(poolData.virtualPrice, 18, 5)}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingDots />
            )
          }
        />
        <InfoListItem
          labelText="Total Liquidity"
          content={
            poolData && poolData?.totalLocked ? (
              <AugmentWithUnits
                content={numeral(poolData.totalLocked).format('0,0')}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingDots />
            )
          }
        />
        <InfoListItem
          labelText="Total Liquidity USD"
          content={
            poolData && poolData?.totalLockedUSD ? (
              <AugmentWithUnits
                content={numeral(poolData.totalLockedUSD).format(usdFormat)}
                label="USD"
              />
            ) : (
              <LoadingDots />
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
}: {
  labelText: string
  content: any
}) => {
  return (
    <li className="flex w-full py-2">
      {labelText}
      <span className="ml-auto">{content}</span>
    </li>
  )
}

export default PoolInfoSection
