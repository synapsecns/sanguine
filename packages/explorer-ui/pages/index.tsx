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
  link,
  ssrMode: true,
  cache: new InMemoryCache(),
})

function Index() {
  return <Home />
}

export default Index
