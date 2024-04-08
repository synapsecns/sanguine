import { type InterchainTransaction } from '@/types'
import { useQuery } from '@tanstack/react-query'
import { GraphQLClient, gql } from 'graphql-request'

const client = new GraphQLClient('https://sanguine-production.up.railway.app')

type InterchainTransactionResponse = {
  interchainTransaction: InterchainTransaction
}

export const useInterchainTransaction = (transactionId: string) => {
  return useQuery<InterchainTransaction>({
    queryKey: ['interchain-transaction', transactionId], // Include transactionId in queryKey for cache uniqueness
    queryFn: async () => {
      try {
        const query = gql`
          query GetInterchainTransaction($id: String!) {
            interchainTransaction(id: $id) {
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
                dstReceiver
                transactionHash
                timestamp
              }
            }
          }
        `
        const variables = { id: transactionId }
        const response = (await client.request(
          query,
          variables
        )) as InterchainTransactionResponse
        return response.interchainTransaction
      } catch (error) {
        console.error(
          `Error fetching transaction with ID ${transactionId}:`,
          error
        )
        throw error
      }
    },
    // Optionally, set staleness and refetch intervals
    staleTime: 60_000,
    refetchInterval: 5_000,
  })
}
