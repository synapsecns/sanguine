import Grid from '@components/tailwind/Grid'
import { getAddressesUrl } from '@urls'
import { ContainerCard } from '@components/ContainerCard'
import {
  CurrencyDollarIcon,
  FireIcon,
  LightningBoltIcon,
} from '@heroicons/react/outline'
import { infoBlockIconClassName } from '@constants'
import { GetMostCommonTokens } from '@pages/Address/GetMostCommonTokens'
import { StatisticBlock } from '@pages/Address/StatisticBlock'
import { InfoBlock } from '@components/misc/InfoBlock'
import { ellipsizeString } from '@utils/ellipsizeString'

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
    <a href={getAddressesUrl({ address })} className={hoverColor}>
      <ContainerCard
        title={formattedAddress}
        className="mt-2 mb-5 font-mono"
        subtitle={rank}
        subtitleClassName="text-xl text-yellow-500"
      >
        <Grid gap={4} cols={{ sm: 1, md: 3, lg: 3 }} className="pt-4 font-sans">
          <StatisticBlock
            title="Total Bridge Volume"
            logo={<CurrencyDollarIcon className={iconClassName} />}
            address={address}
            type="TOTAL_VOLUME_USD"
            duration="PAST_DAY"
            prefix="$"
          />
          <StatisticBlock
            title="Total Bridge Transactions"
            logo={<FireIcon className={iconClassName} />}
            address={address}
            type="COUNT_TRANSACTIONS"
            duration="PAST_DAY"
          />
          <InfoBlock
            title="Most Bridged"
            logo={<LightningBoltIcon className={infoBlockIconClassName} />}
            content={<GetMostCommonTokens address={address} hours={720} />}
          />
        </Grid>
      </ContainerCard>
    </a>
  )
}
