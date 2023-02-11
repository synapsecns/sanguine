import {ApolloClient, InMemoryCache} from '@apollo/client'
import {API_URL} from '@graphql'

const client = new ApolloClient({
  uri: API_URL,
  cache: new InMemoryCache(),
  fetchPolicy: 'network-only',
  fetchOptions: {
    mode: 'no-cors',
  },
})

export default client
