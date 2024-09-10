import numeral from 'numeral'
import { useTranslations } from 'next-intl'

import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'
import { usePoolDataState } from '@/slices/pools/hooks'

const PoolInfoSection = () => {
  const { pool, poolData, isLoading } = usePoolDataState()

  const t = useTranslations('Pools.Other')

  const usdFormat = poolData.totalLockedUSD > 1000000 ? '$0,0.0' : '$0,0'

  return (
    <div className="space-y-4">
      <CurrencyReservesCard />
      <InfoSectionCard title={t('Pool Info')}>
        <InfoListItem
          labelText={t('Trading Fee')}
          content={
            poolData && poolData.swapFee && !isLoading ? (
              formatBigIntToPercentString(poolData.swapFee, 8, 2, false)
            ) : (
              <LoadingDots className="mr-4" />
            )
          }
        />
        <InfoListItem
          labelText={t('Virtual Price')}
          content={
            poolData && poolData?.virtualPrice && !isLoading ? (
              <AugmentWithUnits
                content={formatBigIntToString(poolData.virtualPrice, 18, 5)}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingDots className="mr-4" />
            )
          }
        />
        <InfoListItem
          labelText={t('Total Liquidity')}
          content={
            poolData && poolData?.totalLocked && !isLoading ? (
              <AugmentWithUnits
                content={numeral(poolData.totalLocked).format('0,0')}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingDots className="mr-4" />
            )
          }
        />
        <InfoListItem
          labelText={t('Total Liquidity USD')}
          content={
            poolData && poolData?.totalLockedUSD && !isLoading ? (
              <AugmentWithUnits
                content={numeral(poolData.totalLockedUSD).format(usdFormat)}
                label="USD"
              />
            ) : (
              <LoadingDots className="mr-4" />
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
      <div className="text-white">{labelText} </div>
      <div className="self-center ml-auto text-white">{content}</div>
    </li>
  )
}

export default PoolInfoSection
