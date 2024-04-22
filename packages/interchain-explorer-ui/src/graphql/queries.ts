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
        status
        interchainBatch {
          id
          status
          verifiedAt
          appConfig {
            id
            requiredResponses
            optimisticPeriod
            modules
          }
        }
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
          dbNonce
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
          dbNonce
          count
        }
      }
    }
  }
`

export const GET_STATS = gql`
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

export const GET_INTERCHAIN_TRANSACTION = gql`
  query GetInterchainTransaction($id: String!) {
    interchainTransaction(id: $id) {
      id
      status
      interchainBatch {
        id
        status
      }
      interchainTransactionSent {
        id
        chainId
        address
        srcSender
        dstChainId
        dstReceiver
        transactionHash
        options
        dbNonce
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
        dbNonce
        timestamp
      }
    }
  }
`
