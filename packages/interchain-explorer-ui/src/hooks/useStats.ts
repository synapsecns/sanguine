import { useQuery } from '@tanstack/react-query'

import { GET_STATS } from '@/graphql/queries'
import { InterchainTransaction } from '@/types'
import { client } from '@/graphql/client'

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
        const variables = { limit: 50 }

        const response = (await client.request(
          GET_STATS,
          variables
        )) as StatsResponse
        return response.interchainTransactions.items
      } catch (error) {
        console.error(`Error fetching transaction stats`, error)
        throw error
      }
    },
    staleTime: 60_000,
    refetchInterval: 5_000,
  })
}
