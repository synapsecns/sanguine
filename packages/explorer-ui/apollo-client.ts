import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { API_URL } from '@graphql'

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link,
  cache: new InMemoryCache(),
  // @ts-expect-error TS(2345): Argument of type '{ link: HttpLink; cache: InMemor... Remove this comment to see the full error message
  fetchPolicy: 'network-only',
  fetchOptions: {
    mode: 'no-cors',
  },
})

export default client
