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
import { SynapseLogoSvg } from "@components/layouts/MainLayout/SynapseLogoSvg";

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
  const [wallet, setWallet] = useState("")
  const [minSize, setMinSize] = useState("")
  const [maxSize, setMaxSize] = useState("")
  const [chains, setChains] = useState([])
  const [tokens, setTokens] = useState([])
  const [startDate, setStartDate] = useState(null)
  const [endDate, setEndDate] = useState(null)
  const [toTx, setToTx] = useState(true)
  const [fromTx, setFromTx] = useState(true)
  const [kappa, setKappa] = useState("")




  const [getBridgeTransactions, { loading, error, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY
  )

  useEffect(() => {
    if (data) {
      setTransactionsArr(data.bridgeTransactions
      )
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
  const createQueryField = (field, value, query) => {
    if (value !== "") {
      // if (field === "endTime" || field === "startTime") {
      //   query[field] = parseInt((new Date(field).getTime() / 1000).toFixed(0))
      // } else {
      //   query[field] = value
      // }
      query[field] = value
    }
    return query
  }
  const executeSearch = () => {
    let variables = { page: 1 }
    variables = createQueryField("address", wallet, variables)
    variables = createQueryField("minAmount", minSize, variables)
    variables = createQueryField("maxAmount", maxSize, variables)
    variables = createQueryField("startTime", startDate, variables)
    variables = createQueryField("endTime", endDate, variables)
    variables = createQueryField("kappa", kappa, variables)
    variables = createQueryField("pending", pending, variables)
    console.log(variables)
    getBridgeTransactions({
      variables: variables,
    })
  }
  let bridgeTransactionsTable = _.orderBy(
    transactionsArr,
    'fromInfo.time',
    ['desc']
  ).slice(0, 25)

  return (
    <>
      <StandardPageContainer title="Synapse Analytics">
        <div className="flex items-center mt-10 mb-2">
          <h3 className="text-white text-2xl font-semibold">Bridge Transactions</h3>
        </div>

        <HorizontalDivider />
        <UniversalSearch placeholder={'Search bridge transactions by bridge tx'}
          setPending={handlePending}
          pending={pending}
          loading={loading}
          setWallet={setWallet}
          wallet={wallet}
          setMinSize={setMinSize}
          minSize={minSize}
          setMaxSize={setMaxSize}
          maxSize={maxSize}
          setStartDate={setStartDate}
          startDate={startDate}
          setEndDate={setEndDate}
          endDate={endDate}
          setToTx={setToTx}
          toTx={toTx}
          setFromTx={setFromTx}
          FromTx={fromTx}
          setKappa={setKappa}
          kappa={kappa}
          executeSearch={executeSearch}
          chains={chains}
          setChains={setChains}
          tokens={tokens}
          setTokens={setTokens}
        />
        {loading ? <div className="flex justify-center align-center w-full my-[100px] animate-spin"><SynapseLogoSvg /></div> : <BridgeTransactionTable queryResult={bridgeTransactionsTable} />}


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
