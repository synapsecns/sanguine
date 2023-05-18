import { BeakerIcon, FireIcon } from '@heroicons/react/outline'
import Grid from '@components/tailwind/Grid'
import { ContainerCard } from '@components/ContainerCard'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'

import { TopChains } from './TopChains'
import { TopTokens } from './TopTokens'
import { AddressRanking } from './AddressRanking'

export function Leaderboard() {
  return (
    <StandardPageContainer title={'Leaderboard'}>
      <Grid gap={4} cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} className="mb-5">
        <ContainerCard
          title="Top Chains by Transaction Count"
          subtitle="30 days"
          className="mt-10"
          icon={<BeakerIcon className="w-5 h-5 text-purple-500" />}
        >
          <TopChains />
        </ContainerCard>
        <ContainerCard
          title="Top Tokens by Transaction Count"
          subtitle="30 days"
          className="mt-10"
          icon={<FireIcon className="w-5 h-5 text-red-500" />}
        >
          <TopTokens />
        </ContainerCard>
      </Grid>
      <AddressRanking />
    </StandardPageContainer>
  )
}
