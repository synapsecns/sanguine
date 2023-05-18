import {BeakerIcon, FireIcon} from '@heroicons/react/outline'
import Grid from '@components/tailwind/Grid'

import {ContainerCard} from '@components/ContainerCard'
import {StandardPageContainer} from '@components/layouts/StandardPageContainer'
import {TopChains} from './TopChains'
import {TopTokens} from './TopTokens'
import {AddressRanking} from './AddressRanking'

export function Leaderboard() {
  return (
    // @ts-expect-error TS(2749): 'StandardPageContainer' refers to a value, but is ... Remove this comment to see the full error message
    <StandardPageContainer title={'Leaderboard'}>
      // @ts-expect-error TS(2749): 'Grid' refers to a value, but is being used as a t... Remove this comment to see the full error message
      <Grid gap={4} cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} className="mb-5">
        // @ts-expect-error TS(2749): 'ContainerCard' refers to a value, but is being us... Remove this comment to see the full error message
        <ContainerCard
          // @ts-expect-error TS(2304): Cannot find name 'title'.
          title="Top Chains by Transaction Count"
          // @ts-expect-error TS(2304): Cannot find name 'subtitle'.
          subtitle="30 days"
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="mt-10"
          // @ts-expect-error TS(2304): Cannot find name 'icon'.
          icon={<BeakerIcon className="w-5 h-5 text-purple-500" />}
        // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
        >
          // @ts-expect-error TS(2749): 'TopChains' refers to a value, but is being used a... Remove this comment to see the full error message
          <TopChains />
        </ContainerCard>
        <ContainerCard
          // @ts-expect-error TS(2304): Cannot find name 'title'.
          title="Top Tokens by Transaction Count"
          // @ts-expect-error TS(2304): Cannot find name 'subtitle'.
          subtitle="30 days"
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="mt-10"
          // @ts-expect-error TS(2304): Cannot find name 'icon'.
          icon={<FireIcon className="w-5 h-5 text-red-500" />}
        // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
        >
          // @ts-expect-error TS(2749): 'TopTokens' refers to a value, but is being used a... Remove this comment to see the full error message
          <TopTokens />
        </ContainerCard>
      </Grid>
      // @ts-expect-error TS(2362): The left-hand side of an arithmetic operation must... Remove this comment to see the full error message
      <AddressRanking />
    </StandardPageContainer>
  )
}
