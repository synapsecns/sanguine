import { useQuery } from '@tanstack/react-query'
import { GraphQLClient, gql } from 'graphql-request'

import { type InterchainTransaction } from '@/types'

const client = new GraphQLClient('https://sanguine-production.up.railway.app')

type InterchainTransactionsResponse = {
  interchainTransactions: {
    items: InterchainTransaction[]
  }
}

export const useInterchainTransactions = () => {
  return useQuery<InterchainTransaction[]>({
    queryKey: ['interchain-transactions'],
    queryFn: async () => {
      try {
        const r = (await client.request(gql`
          {
            interchainTransactions(
              orderBy: "sentAt"
              orderDirection: "desc"
              limit: 1000
            ) {
              items {
                id
                interchainTransactionSent {
                  id
                  chainId
                  address
                  srcSender
                  dstChainId
                  dstReceiver
                  transactionHash
                  options
                  timestamp
                  count
                }
                interchainTransactionReceived {
                  id
                  chainId
                  address
                  srcSender
                  srcChainId
                  transactionHash
                  dstReceiver
                  timestamp
                  count
                }
              }
            }
          }
        `)) as InterchainTransactionsResponse

        return r.interchainTransactions.items.map((d) => ({
          id: d.id,
          interchainTransactionSent: d.interchainTransactionSent,
          interchainTransactionReceived: d.interchainTransactionReceived,
        }))
      } catch (error) {
        console.error('Error fetching interchain transactions:', error)
        throw error
      }
    },
    staleTime: 60_000, // 1 minute
    refetchInterval: 1_000, // 1 second
  })
}
