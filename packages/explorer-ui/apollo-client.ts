import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { API_URL } from '@graphql'

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link,
  cache: new InMemoryCache(),
})

export default client
