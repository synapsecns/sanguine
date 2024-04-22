import { useQuery } from '@tanstack/react-query'

import { type InterchainTransaction } from '@/types'
import { GET_INTERCHAIN_TRANSACTION } from '@/graphql/queries'
import { client } from '@/graphql/client'

type InterchainTransactionResponse = {
  interchainTransaction: InterchainTransaction
}

export const useInterchainTransaction = (transactionId: string) => {
  return useQuery<InterchainTransaction>({
    queryKey: ['interchain-transaction', transactionId],
    queryFn: async () => {
      try {
        const variables = { id: transactionId }
        const response = (await client.request(
          GET_INTERCHAIN_TRANSACTION,
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
    staleTime: 60_000,
    refetchInterval: 5_000,
  })
}
