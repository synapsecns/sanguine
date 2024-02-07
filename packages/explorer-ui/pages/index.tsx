import { Home } from '@components/pages/Home'
import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import {
  AMOUNT_STATISTIC,
  DAILY_STATISTICS_BY_CHAIN,
  RANKED_CHAINIDS_BY_VOLUME,
} from '@graphql/queries'
import { API_URL } from '@graphql'
import ReactGA from 'react-ga'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = "G-BBC13LQXBD"
ReactGA.initialize(TRACKING_ID)

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
