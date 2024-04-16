import { useQuery } from '@tanstack/react-query'
import { GraphQLClient, gql } from 'graphql-request'

import { type InterchainTransaction } from '@/types'

const client = new GraphQLClient('https://sanguine-production.up.railway.app')

type InterchainTransactionsResponse = {
  interchainTransactions: {
    items: InterchainTransaction[]
  }
}

export const useStats = () => {
  return useQuery({
    queryKey: ['interchain-transaction-stats'],
    queryFn: async () => {
      try {
        const query = gql`
          query GetStats($limit: Int) {
            interchainTransactions(
              orderBy: "sentAt"
              orderDirection: "desc"
              limit: $limit
            ) {
              items {
                interchainTransactionSent {
                  id
                  count
                }
                interchainTransactionReceived {
                  id
                  count
                }
              }
            }
          }
        `

        const variables = { limit: 1 }

        const response = (await client.request(query, variables)) as any
        return response.interchainTransactions.items[0]
      } catch (error) {
        console.error(`Error fetching transaction stats`, error)
        throw error
      }
    },
    // Optionally, set staleness and refetch intervals
    staleTime: 60_000,
    refetchInterval: 5_000,
  })
}
