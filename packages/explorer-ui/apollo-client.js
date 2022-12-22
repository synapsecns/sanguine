import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client'
import { API_URL } from '@graphql'

const link = new HttpLink({
  uri: 'http://24.199.82.219:80/graphql',
  useGETForQueries: true,
})

const client = new ApolloClient({
  uri: link,
  cache: new InMemoryCache(),
  fetchPolicy: 'network-only',
  fetchOptions: {
    mode: 'no-cors',
  },
})

export default client
