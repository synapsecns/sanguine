import { GraphQLClient } from 'graphql-request'

export const client = new GraphQLClient(
  'https://sanguine-production.up.railway.app'
)
// export const client = new GraphQLClient('http://localhost:42069')
