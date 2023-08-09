import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { graphqlRequestBaseQuery } from '@rtk-query/graphql-request-base-query'
import { gql } from 'graphql-request'

const SCRIBE_URL: string = 'https://explorer.omnirpc.io/graphiql'

export const api = createApi({
  baseQuery: graphqlRequestBaseQuery({
    url: SCRIBE_URL,
  }),
  endpoints: () => ({}),
})

export const api = createApi({
  reducerPath: 'dataApi',
  baseQuery: fetchBaseQuery({ baseUrl: SCRIBE_URL }),
  endpoints: (builder) => ({
    userHistoricalActivityQuery: builder.query({
      query: ({ address, startTime }) => gql`
        query GetUserHistoricalActivity($address: String!, $startTime: Int!) {
          bridgeTransactions(
            pending: false
            addressFrom: $address
            startTime: $startTime
            page: 1
          ) {
            fromInfo {
              chainID
              destinationChainID
              address
              txnHash
              value
              formattedValue
              USDValue
              tokenAddress
              tokenSymbol
              blockNumber
              time
              formattedTime
              formattedEventType
              eventType
            }
            toInfo {
              chainID
              destinationChainID
              address
              txnHash
              value
              formattedValue
              USDValue
              tokenAddress
              tokenSymbol
              blockNumber
              time
              formattedTime
              formattedEventType
              eventType
            }
            kappa
          }
        }
      `,
    }),
    userPendingActivityQuery: builder.query({
      query: ({ address, startTime }) => gql`
        query GetUserHistoricalActivity($address: String!, $startTime: Int!) {
          bridgeTransactions(
            pending: true
            addressFrom: $address
            startTime: $startTime
            page: 1
          ) {
            fromInfo {
              chainID
              destinationChainID
              address
              txnHash
              value
              formattedValue
              USDValue
              tokenAddress
              tokenSymbol
              blockNumber
              time
              formattedTime
              formattedEventType
              eventType
            }
            toInfo {
              chainID
              destinationChainID
              address
              txnHash
              value
              formattedValue
              USDValue
              tokenAddress
              tokenSymbol
              blockNumber
              time
              formattedTime
              formattedEventType
              eventType
            }
            kappa
          }
        }
      `,
    }),
  }),
})
