import { useQuery } from '@apollo/client'
import { FireIcon, LocationMarkerIcon } from '@heroicons/react/outline'

import { COUNT_BY_CHAIN_ID } from '@graphql/queries'
import Grid from '@components/tailwind/Grid'

import { StyledChainAndLink } from '@components/misc/StyledChainAndLink'
import { ContainerCard } from '@components/ContainerCard'
import { InfoDisplay } from './InfoDisplay'
import { InfoLoader } from './InfoLoader'
import { infoBlockIconClassName } from '@constants'

export function ChainActivityCard({ chainId }) {
  return (
    <ContainerCard
      title="Chain Activity"
      subtitle="30 days"
      icon={<FireIcon className="w-5 h-5 text-orange-500" />}
    >
      <ChainActivity chainId={chainId} />
    </ContainerCard>
  )
}

export function getChainActivity({ chainId, direction }) {
  const { data } = useQuery(COUNT_BY_CHAIN_ID, {
    variables: {
      chainId: Number(chainId),
      direction,
      hours: 720,
    },
  })

  const { countByChainId } = data ?? {}

  return countByChainId
}

function ChainActivity({ chainId }) {
  let showBoth = !chainId

  let origins
  let destinations

  if (showBoth) {
    const fromCountByChainId = getChainActivity({ direction: 'OUT' })

    origins = {
      title: 'Origins',
      content: fromCountByChainId ? (
        <Activity data={fromCountByChainId} />
      ) : (
        <InfoLoader />
      ),
      logo: <LocationMarkerIcon className={infoBlockIconClassName} />,
    }
  }

  const toCountByChainId = getChainActivity({ chainId, direction: 'IN' })

  destinations = {
    title: 'Destinations',
    content: toCountByChainId ? (
      <Activity data={toCountByChainId} />
    ) : (
      <InfoLoader />
    ),
    logo: <LocationMarkerIcon className={infoBlockIconClassName} />,
  }

  return showBoth ? (
    <Grid gap={1} cols={{ sm: 1, lg: 2 }} className="lg:mt-6">
      <InfoDisplay arr={[origins]} />
      <InfoDisplay arr={[destinations]} />
    </Grid>
  ) : (
    <Grid gap={1} cols={{ sm: 1 }} className="lg:mt-6">
      <InfoDisplay arr={[destinations]} />
    </Grid>
  )
}

function Activity({ data }) {
  return (
    <div className="mb-1 text-xl">
      {data.slice(0, 3).map(({ chainId }, i) => (
        <div key={i}>
          <StyledChainAndLink chainId={chainId} />
        </div>
      ))}
    </div>
  )
}
