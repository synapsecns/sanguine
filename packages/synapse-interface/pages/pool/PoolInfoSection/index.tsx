import { Zero } from '@ethersproject/constants'

import {
  commifyBnToString,
  commifyBnWithDefault,
  bnPercentFormat,
} from '@bignumber/format'

import AugmentWithUnits from '../components/AugmentWithUnits'
import { Token } from '@types'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import LoadingSpinner from '@tw/LoadingSpinner'
const PoolInfoSection = ({
  pool,
  poolData,
  chainId,
}: {
  pool: Token
  poolData: any
  chainId: number
}) => {
  // const swapFee = bnPercentFormat('0.02')
  // this needs to be fixed, need admin fee

  return (
    <div className="space-y-4">
      {/* <UserPoolInfoCard data={userData} /> */}
      <CurrencyReservesCard
        title="Currency Reserves"
        chainId={chainId}
        poolData={poolData}
      />
      <InfoSectionCard title="Pool Info">
        <InfoListItem
          labelText="Trading Fee"
          content={true ? '0.02%' : <LoadingSpinner />}
        />
        <InfoListItem
          labelText="Virtual Price"
          content={
            poolData?.virtualPriceStr ? (
              <AugmentWithUnits
                content={poolData.virtualPriceStr}
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
