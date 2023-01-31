import { HorizontalDivider } from "@components/misc/HorizontalDivider";
import { UniversalSearch } from "@components/pages/Home/UniversalSearch";
import { BridgeTransactionTable } from "@components/BridgeTransaction/BridgeTransactionTable";
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import {
  GET_BRIDGE_TRANSACTIONS_QUERY
} from "@graphql/queries";
import { API_URL } from '@graphql'
import { useState, useEffect } from "react";
import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import _ from "lodash";
import { Pagination } from "@components/Pagination";
import { useLazyQuery } from '@apollo/client'


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


export default function Txs({ queryResult }) {
  const [transactionsArr, setTransactionsArr] = useState([])
  const [pending, setPending] = useState(false)
  const [getBridgeTransactions, { loading, error, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY
  )

  useEffect(() => {
    if (data) {
      setTransactionsArr(data.bridgeTransactions, {
        variables: {
          pending: pending,
        },
      })
    }

  }, [data, pending])


  // Get initial data
  useEffect(() => {
    getBridgeTransactions({
      variables: {
        pending: pending,
        page: 1,
      },
    })
  }, [])
  const handlePending = (arg) => {
    setPending(arg)
    getBridgeTransactions({
      variables: {
        pending: arg,
        page: 1,
      },
    })
  }

  let bridgeTransactionsTable = _.orderBy(
    transactionsArr,
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
        <UniversalSearch placeholder={'Search all bridge transactions'}
          setPending={handlePending}
          pending={pending}
          loading={loading} />
        {loading ? <div className="text-white">Loading...</div> : <BridgeTransactionTable queryResult={bridgeTransactionsTable} />}


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
        pending: false,
        page: context.query.p ?? 1,
      },
    })
    result = data
  } else if (context.query.chainId) {
    let { data } = await client.query({
      query: GET_BRIDGE_TRANSACTIONS_QUERY,
      variables: {
        chainID: [context.query.chainId],
        pending: false,
        page: context.query.p ?? 1,
      },
    })
    result = data
  } else {
    let { data } = await client.query({
      query: GET_BRIDGE_TRANSACTIONS_QUERY,
      variables: {
        pending: false,
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
