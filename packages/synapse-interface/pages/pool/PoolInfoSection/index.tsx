import { Zero } from '@ethersproject/constants'

import {
  commifyBnToString,
  commifyBnWithDefault,
  bnPercentFormat,
} from '@bignumber/format'

import { PRICE_UNITS_INDEX } from '@constants/priceUnits'

import InfoListItem from '@components/InfoListItem'
import AugmentWithUnits from '@components/AugmentWithUnits'

import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

export default function PoolInfoSection({ data, userData }) {
  const { chainId } = useActiveWeb3React()
  const swapFee = bnPercentFormat(data?.swapFee)
  // const defaultDepositFee = bnPercentFormat(data?.defaultDepositFee)
  let adminFee = bnPercentFormat(data?.adminFee)

  if (swapFee && adminFee) {
    adminFee = `${adminFee} of ${swapFee}`
  }

  const standardUnits = PRICE_UNITS_INDEX[data?.name] ?? ''

  const tokens = data?.tokens

  let displayDecimals
  if (standardUnits === 'ETH') {
    displayDecimals = 3
  } else {
    displayDecimals = 0
  }
  const totalLocked = commifyBnWithDefault(data?.totalLocked, displayDecimals)
  const totalLockedUSD = commifyBnWithDefault(data?.totalLockedUSD ?? Zero, 0)

  const virtualPrice = data?.virtualPrice
    ? commifyBnToString(data.virtualPrice, 5)
    : null

  return (
    <div className="space-y-4">
      {/* <UserPoolInfoCard data={userData} /> */}
      {tokens && (
        <CurrencyReservesCard
          title="Currency Reserves"
          chainId={chainId}
          tokens={tokens}
          poolToken={data.poolToken}
        />
      )}
      <InfoSectionCard title="Pool Info">
        <InfoListItem labelText="Trading Fee" content={swapFee} />
        <InfoListItem
          labelText="Virtual Price"
          content={
            <AugmentWithUnits content={virtualPrice} label={standardUnits} />
          }
        />
        <InfoListItem
          labelText="Total Liquidity"
          content={
            <AugmentWithUnits content={totalLocked} label={standardUnits} />
          }
        />
        <InfoListItem
          labelText="Total Liquidity USD"
          content={`$${totalLockedUSD}`}
        />
      </InfoSectionCard>
    </div>
  )
}
