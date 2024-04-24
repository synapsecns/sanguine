import { gql } from 'graphql-request'

const SENT_INFO_FRAGMENT = gql`
  fragment InterchainTransactionSentFields on InterchainTransactionSent {
    id
    address
    srcSender
    srcChainId
    dstChainId
    dstReceiver
    transactionHash
    options
    timestamp
    dbNonce
    count
  }
`

const RECEIVED_INFO_FRAGMENT = gql`
  fragment InterchainTransactionReceivedFields on InterchainTransactionReceived {
    id
    address
    srcSender
    srcChainId
    dstChainId
    dstReceiver
    transactionHash
    timestamp
    dbNonce
    count
  }
`

const BATCH_INFO_FRAGMENT = gql`
  fragment InterchainBatchFields on InterchainBatch {
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
`

export const GET_INTERCHAIN_TRANSACTIONS = gql`
  ${BATCH_INFO_FRAGMENT}
  ${SENT_INFO_FRAGMENT}
  ${RECEIVED_INFO_FRAGMENT}

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
          ...InterchainBatchFields
        }
        interchainTransactionSent {
          ...InterchainTransactionSentFields
        }
        interchainTransactionReceived {
          ...InterchainTransactionReceivedFields
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
  ${BATCH_INFO_FRAGMENT}
  ${SENT_INFO_FRAGMENT}
  ${RECEIVED_INFO_FRAGMENT}

  query GetInterchainTransaction($id: String!) {
    interchainTransaction(id: $id) {
      id
      status
      interchainBatch {
        ...InterchainBatchFields
      }
      interchainTransactionSent {
        ...InterchainTransactionSentFields
      }
      interchainTransactionReceived {
        ...InterchainTransactionReceivedFields
      }
    }
  }
`
