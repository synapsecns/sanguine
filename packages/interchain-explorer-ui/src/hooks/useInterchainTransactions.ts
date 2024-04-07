import { type InterchainTransaction } from '@/types'
import { useQuery } from '@tanstack/react-query'
import { GraphQLClient, gql } from 'graphql-request'

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
            interchainTransactions {
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
        console.error('Error fetching deposits:', error)
        throw error
      }
    },
    staleTime: 60_000, // 1 minute
    refetchInterval: 5_000, // 5 seconds
  })
}
