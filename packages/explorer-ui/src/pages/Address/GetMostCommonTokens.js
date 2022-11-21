import { useQuery } from '@apollo/client'

import { COUNT_BY_TOKEN_ADDRESS } from '@graphql/queries'

import { MostActive } from '@components/misc/MostActive'

export function GetMostCommonTokens({ address, hours = 50000 }) {
  const { data } = useQuery(COUNT_BY_TOKEN_ADDRESS, {
    variables: {
      hours,
      address: address,
    },
  })

  let count = data?.countByTokenAddress

  if (count) {
    return (
      <div className="mb-2">
        <MostActive data={count} />
      </div>
    )
  } else {
    return (
      <div className="w-10 h-4 bg-slate-400 dark:bg-slate-500 animate-pulse" />
    )
  }
}
