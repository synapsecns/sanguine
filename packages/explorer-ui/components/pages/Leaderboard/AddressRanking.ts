import {useQuery} from '@apollo/client'

import {ADDRESS_RANKING} from '@graphql/queries'

import {LeaderCard} from './LeaderCard'

const NUM_TO_SHOW = 10

export function AddressRanking() {
  const { data } = useQuery(ADDRESS_RANKING)

  if (data) {
    let content

    let { addressRanking } = data

    addressRanking = addressRanking.map((entry, i) => {
      return {
        rank: i + 1,
        address: entry.address,
        count: entry.count,
      }
    })

    content = addressRanking
      .slice(0, NUM_TO_SHOW)
      // @ts-expect-error TS(2749): 'LeaderCard' refers to a value, but is being used ... Remove this comment to see the full error message
      .map((leader, i) => <LeaderCard key={i} {...leader} />)

    return (
      <>
        // @ts-expect-error TS(2304): Cannot find name 'span'.
        <span
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className={`
              text-2xl font-medium text-default
              bg-clip-text text-transparent bg-gradient-to-r
              from-purple-600 to-blue-600
            `}
        >
          // @ts-expect-error TS(2552): Cannot find name 'Top'. Did you mean 'top'?
          Top Transacting Addresses in Last 30 days
        </span>
        // @ts-expect-error TS(2304): Cannot find name 'content'.
        {content}
      </>
    )
  } else {
    return 'loading'
  }
}
