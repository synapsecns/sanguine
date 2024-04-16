import { useQuery } from '@tanstack/react-query'
import { GraphQLClient } from 'graphql-request'

import { GET_STATS } from '@/graphql/queries'
import { InterchainTransaction } from '@/types'

const client = new GraphQLClient('https://sanguine-production.up.railway.app')

type StatsResponse = {
  interchainTransactions: {
    items: InterchainTransaction[]
  }
}

export const useStats = () => {
  return useQuery({
    queryKey: ['interchain-transaction-stats'],
    queryFn: async () => {
      try {
        const variables = { limit: 1 }

        const response = (await client.request(
          GET_STATS,
          variables
        )) as StatsResponse
        return response.interchainTransactions.items[0]
      } catch (error) {
        console.error(`Error fetching transaction stats`, error)
        throw error
      }
    },
    staleTime: 60_000,
    refetchInterval: 5_000,
  })
}
