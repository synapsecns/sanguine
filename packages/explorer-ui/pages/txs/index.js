import {HorizontalDivider} from "@components/misc/HorizontalDivider";
import {UniversalSearch} from "@components/pages/Home/UniversalSearch";
import {BridgeTransactionTable} from "@components/BridgeTransaction/BridgeTransactionTable";
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import {
  GET_BRIDGE_TRANSACTIONS_QUERY
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
  let { bridgeTransactions: bridgeTransactionsTable } = queryResult

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


export async function getServerSideProps(context) {
  let result

  if (context.query.account) {
    let { data } = await client.query({
      query: GET_BRIDGE_TRANSACTIONS_QUERY,
      variables: {
        address: context.query.account,
        includePending: false,
        page: context.query.p ?? 1,
      },
    })
    result = data
  } else if (context.query.chainId) {
    let { data } = await client.query({
      query: GET_BRIDGE_TRANSACTIONS_QUERY,
      variables: {
        chainId: context.query.chainId,
        includePending: false,
        page: context.query.p ?? 1,
      },
    })
    result = data
  } else {
    let { data } = await client.query({
      query: GET_BRIDGE_TRANSACTIONS_QUERY,
      variables: {
        includePending: false,
        page: context.query.p ?? 1,
      },
    })
    result = data
  }

  return {
    props: {
      queryResult: result,
    },
  }
}
