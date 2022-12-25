import {HorizontalDivider} from "@components/misc/HorizontalDivider";
import {UniversalSearch} from "@components/pages/Home/UniversalSearch";
import {BridgeTransactionTable} from "@components/BridgeTransaction/BridgeTransactionTable";
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import {
  BRIDGE_AMOUNT_STATISTIC, COUNT_BY_CHAIN_ID,
  COUNT_BY_TOKEN_ADDRESS,
  GET_HISTORICAL_STATS,
  GET_LATEST_BRIDGE_TRANSACTIONS_QUERY
} from "@graphql/queries";
import {API_URL} from '@graphql'

import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'
import _ from "lodash";
import {Pagination} from "@components/Pagination";

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  cache: new InMemoryCache(),
  fetchPolicy: 'network-only',
  fetchOptions: {
    mode: 'no-cors',
  },
})


export default function Txs({queryResult}) {
  let { latestBridgeTransactions: bridgeTransactionsTable } = queryResult


  bridgeTransactionsTable = _.orderBy(
    bridgeTransactionsTable,
    'fromInfo.time',
    ['desc']
  ).slice(0, 25)


  return (
    <>
      <StandardPageContainer>
        <div className="flex items-center mb-10">
          <h3 className="text-white text-4xl font-semibold">Bridge Transactions</h3>
        </div>
        <HorizontalDivider />
        <UniversalSearch placeholder={'Search all bridge transactions'} />
        <BridgeTransactionTable queryResult={bridgeTransactionsTable} />

        <HorizontalDivider />
        <Pagination />
        </StandardPageContainer>
    </>
  )
}


export async function getServerSideProps() {
  const { data } = await client.query({
    query: GET_LATEST_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      includePending: false,
      page: 1,
    },
  })

  return {
    props: {
      queryResult: data,
    },
  }
}
