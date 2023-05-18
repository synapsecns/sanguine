import Grid from '@components/tailwind/Grid'

import {getAddressesUrl} from '@urls'
import {ContainerCard} from '@components/ContainerCard'
import {CurrencyDollarIcon, FireIcon, LightningBoltIcon,} from '@heroicons/react/outline'
import {infoBlockIconClassName} from '@constants'

// @ts-expect-error TS(2307): Cannot find module '@pages/Address/GetMostCommonTo... Remove this comment to see the full error message
import {GetMostCommonTokens} from '@pages/Address/GetMostCommonTokens'

// @ts-expect-error TS(2307): Cannot find module '@pages/Address/StatisticBlock'... Remove this comment to see the full error message
import {StatisticBlock} from '@pages/Address/StatisticBlock'
import {InfoBlock} from '@components/misc/InfoBlock'
import {ellipsizeString} from '@utils/ellipsizeString'

export function LeaderCard({ rank, address, count }) {
  const hoverColors = {
    1: ' !shadow-yellow-600 hover:shadow-yellow-700 hover:!shadow-yellow-500 transition-all',
    2: ' !shadow-blue-600 hover:shadow-blue-700 hover:!shadow-blue-500 transition-all',
    3: ' !shadow-purple-600 hover:shadow-purple-700 hover:!shadow-purple-500 transition-all',
  }
  let hoverColor
  let borderType
  let shadowSize
  let textSize

  switch (rank) {
    case 1:
      hoverColor = hoverColors[rank]
      borderType = ' border-2 border-double'
      shadowSize = ' !shadow-md'
      textSize = ' sm:text-md lg:text-xl md:text-xl'
      break
    case 2:
      hoverColor = hoverColors[rank]
      borderType = ' border-2 border-double'
      shadowSize = ' !shadow-md'
      textSize = ' sm:text-md lg:text-xl md:text-xl'
      break
    case 3:
      hoverColor = hoverColors[rank]
      borderType = ' border-2 border-double'
      shadowSize = ' !shadow-md'
      textSize = ' sm:text-md lg:text-xl md:text-xl'
      break
    default:
      hoverColor = ' !shadow-red-600 hover:shadow-red-500'
      borderType = ' border'
      shadowSize = ' !shadow-sm'
      textSize = ' text-md'
  }

  const iconClassName = 'w-5 h-5 text-purple-500 inline -mt-1'

  const formattedAddress = ellipsizeString({ string: address, limiter: 10 })

  return (
    // @ts-expect-error TS(2304): Cannot find name 'a'.
    <a href={getAddressesUrl({ address })} className={hoverColor}>
      // @ts-expect-error TS(2749): 'ContainerCard' refers to a value, but is being us... Remove this comment to see the full error message
      <ContainerCard
        // @ts-expect-error TS(2304): Cannot find name 'title'.
        title={formattedAddress}
        // @ts-expect-error TS(2304): Cannot find name 'className'.
        className="mt-2 mb-5 font-mono"
        // @ts-expect-error TS(2304): Cannot find name 'subtitle'.
        subtitle={rank}
        // @ts-expect-error TS(2304): Cannot find name 'subtitleClassName'.
        subtitleClassName="text-xl text-yellow-500"
      >
        // @ts-expect-error TS(2749): 'Grid' refers to a value, but is being used as a t... Remove this comment to see the full error message
        <Grid gap={4} cols={{ sm: 1, md: 3, lg: 3 }} className="pt-4 font-sans">
          <StatisticBlock
            // @ts-expect-error TS(2304): Cannot find name 'title'.
            title="Total Bridge Volume"
            // @ts-expect-error TS(2304): Cannot find name 'logo'.
            logo={<CurrencyDollarIcon className={iconClassName} />}
            // @ts-expect-error TS(2304): Cannot find name 'address'.
            address={address}
            // @ts-expect-error TS(2304): Cannot find name 'type'.
            type="TOTAL_VOLUME_USD"
            // @ts-expect-error TS(2304): Cannot find name 'duration'.
            duration="PAST_DAY"
            // @ts-expect-error TS(2304): Cannot find name 'prefix'.
            prefix="$"
          />
          <StatisticBlock
            // @ts-expect-error TS(2304): Cannot find name 'title'.
            title="Total Bridge Transactions"
            // @ts-expect-error TS(2304): Cannot find name 'logo'.
            logo={<FireIcon className={iconClassName} />}
            // @ts-expect-error TS(2304): Cannot find name 'address'.
            address={address}
            // @ts-expect-error TS(2304): Cannot find name 'type'.
            type="COUNT_TRANSACTIONS"
            // @ts-expect-error TS(2304): Cannot find name 'duration'.
            duration="PAST_DAY"
          />
          // @ts-expect-error TS(2749): 'InfoBlock' refers to a value, but is being used a... Remove this comment to see the full error message
          <InfoBlock
            // @ts-expect-error TS(2304): Cannot find name 'title'.
            title="Most Bridged"
            // @ts-expect-error TS(2304): Cannot find name 'logo'.
            logo={<LightningBoltIcon className={infoBlockIconClassName} />}
            // @ts-expect-error TS(2304): Cannot find name 'content'.
            content={<GetMostCommonTokens address={address} hours={720} />}
          // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
          />
        </Grid>
      </ContainerCard>
    </a>
  )
}
