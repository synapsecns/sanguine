import { gql } from 'graphql-request'

export const GET_INTERCHAIN_TRANSACTIONS = gql`
  query GetInterchainTransactions(
    $limit: Int
    $after: String
    $before: String
  ) {
    interchainTransactions(
      limit: $limit
      after: $after
      before: $before
      orderBy: "sentAt"
      orderDirection: "desc"
    ) {
      pageInfo {
        startCursor
        endCursor
        hasPreviousPage
        hasNextPage
      }
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
`
