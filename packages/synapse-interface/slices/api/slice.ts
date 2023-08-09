import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

const SCRIBE_URL: string = 'https://explorer.omnirpc.io/graphiql'

export const api = createApi({
  reducerPath: 'dataApi',
  baseQuery: fetchBaseQuery({ baseUrl: SCRIBE_URL }),
  endpoints: (builder) => ({}),
})
