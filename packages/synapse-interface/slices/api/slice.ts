import { createApi } from '@reduxjs/toolkit/query/react'
import { graphqlRequestBaseQuery } from '@rtk-query/graphql-request-base-query'

const SCRIBE_URL: string = 'https://explorer.omnirpc.io/graphql'

export const api = createApi({
  baseQuery: graphqlRequestBaseQuery({
    url: SCRIBE_URL,
  }),
  endpoints: () => ({}),
})
