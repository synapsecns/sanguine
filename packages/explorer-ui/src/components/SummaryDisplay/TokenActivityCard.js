import { useQuery } from '@apollo/client'
import {
  PresentationChartLineIcon,
  ArrowSmRightIcon,
  ArrowSmLeftIcon,
} from '@heroicons/react/outline'
import Grid from '@tw/Grid'

import { COUNT_BY_TOKEN_ADDRESS } from '@graphql/queries'

import { MostActive } from '@components/misc/MostActive'
import { ContainerCard } from '@components/ContainerCard'
import { InfoDisplay } from './InfoDisplay'
import { InfoLoader } from './InfoLoader'
import { infoBlockIconClassName } from '@constants'

export function TokenActivityCard({ chainId }) {
  return (
    <ContainerCard
      title="Token Activity"
      subtitle="24 hours"
      icon={<PresentationChartLineIcon className="w-5 h-5 text-purple-500" />}
    >
      <TokenActivity chainId={chainId} />
    </ContainerCard>
  )
}

export function popularTokens({ direction, hours = 24 }) {
  const { data } = useQuery(COUNT_BY_TOKEN_ADDRESS, {
    variables: {
      direction,
      hours,
    },
  })

  const { countByTokenAddress } = data ?? {}

  return countByTokenAddress
}

function TokenActivity({ chainId }) {
  const { data: fromData } = useQuery(COUNT_BY_TOKEN_ADDRESS, {
    variables: {
      chainId: Number(chainId),
      direction: 'OUT',
      hours: 24,
    },
  })
  const { data: toData } = useQuery(COUNT_BY_TOKEN_ADDRESS, {
    variables: {
      chainId: Number(chainId),
      direction: 'IN',
      hours: 24,
    },
  })

  const { countByTokenAddress: fromCountByTokenAddress } = fromData ?? {}
  const { countByTokenAddress: toCountByTokenAddress } = toData ?? {}

  return (
    <Grid gap={4} cols={{ sm: 1 }}>
      <InfoDisplay
        arr={[
          {
            title: 'Most Active Sent',
            content:
              fromCountByTokenAddress && toCountByTokenAddress ? (
                <MostActive data={fromCountByTokenAddress} />
              ) : (
                <InfoLoader />
              ),
            logo: <ArrowSmRightIcon className={infoBlockIconClassName} />,
          },
          {
            title: 'Most Active Received',
            content:
              fromCountByTokenAddress && toCountByTokenAddress ? (
                <MostActive data={toCountByTokenAddress} />
              ) : (
                <InfoLoader />
              ),
            logo: <ArrowSmLeftIcon className={infoBlockIconClassName} />,
          },
        ]}
      />
    </Grid>
  )
}
