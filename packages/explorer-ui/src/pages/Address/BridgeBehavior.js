import {
  FireIcon,
  LightningBoltIcon,
  LocationMarkerIcon,
} from '@heroicons/react/outline'

import Grid from '@tw/Grid'

import { ContainerCard } from '@components/ContainerCard'
import { GetMostCommonTokens } from './GetMostCommonTokens'
import { StyledChainAndLink } from '@components/misc/StyledChainAndLink'
import { InfoBlock } from '@components/misc/InfoBlock'
import { infoBlockIconClassName } from '@constants'

export function BridgeBehavior({
  topDestinationChainId,
  topOriginChainId,
  address,
}) {
  return (
    <ContainerCard
      className="text-gray-500 border border-indigo-500 dark:bg-gray-900 dark:text-gray-200 hover:border-purple-500"
      title="Bridge Behavior"
      icon={<FireIcon className="w-5 h-5 text-orange-500" />}
    >
      <Grid gap={4} cols={{ sm: 1, md: 2 }} className="pt-4">
        <InfoBlock
          title="Active Origin"
          logo={<LocationMarkerIcon className={infoBlockIconClassName} />}
          content={<StyledChainAndLink chainId={topOriginChainId} />}
        />
        <InfoBlock
          title="Active Destination"
          logo={<LocationMarkerIcon className={infoBlockIconClassName} />}
          content={<StyledChainAndLink chainId={topDestinationChainId} />}
        />
        <InfoBlock
          className="col-span-2"
          title="Commonly Bridged"
          logo={<LightningBoltIcon className={infoBlockIconClassName} />}
          content={<GetMostCommonTokens address={address} />}
        />
      </Grid>
    </ContainerCard>
  )
}
