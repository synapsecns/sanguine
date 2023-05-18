import { useQuery } from '@apollo/client'
import { ADDRESS_RANKING } from '@graphql/queries'
import { Fragment } from 'react'

import { LeaderCard } from './LeaderCard'

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
      .map((leader, i) => <LeaderCard key={i} {...leader} />)

    return (
      <>
        <span
          className={`
              text-2xl font-medium text-default
              bg-clip-text text-transparent bg-gradient-to-r
              from-purple-600 to-blue-600
            `}
        >
          Top Transacting Addresses in Last 30 days
        </span>
        {content}
      </>
    )
  } else {
    return (
      <>
        <Fragment>loading</Fragment>
      </>
    )
  }
}
