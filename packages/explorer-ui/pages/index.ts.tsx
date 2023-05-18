import { Home } from '@components/pages/Home'
import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import {
  AMOUNT_STATISTIC,
  DAILY_STATISTICS_BY_CHAIN,
  RANKED_CHAINIDS_BY_VOLUME,
} from '@graphql/queries'
import { API_URL } from '@graphql'

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  ssrMode: true,
  cache: new InMemoryCache(),
  // @ts-expect-error TS(2345): Argument of type '{ link: HttpLink; ssrMode: true;... Remove this comment to see the full error message
  fetchPolicy: 'cache-and-network',
  fetchOptions: {
    mode: 'no-cors',
  },
})

function Index() {
  return (
    // @ts-expect-error TS(2749): 'Home' refers to a value, but is being used as a t... Remove this comment to see the full error message
    <Home/>
  )
}

export default Index
