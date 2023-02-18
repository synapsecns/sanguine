import { HorizontalDivider } from "@components/misc/HorizontalDivider";
import { UniversalSearch } from "@components/pages/Home/UniversalSearch";
import { BridgeTransactionTable } from "@components/BridgeTransaction/BridgeTransactionTable";
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import {
  GET_BRIDGE_TRANSACTIONS_QUERY
} from "@graphql/queries";
import { CHAIN_ID_NAMES_REVERSE } from '@constants/networks'
import { useSearchParams } from 'next/navigation'

import { API_URL } from '@graphql'
import { useState, useEffect } from "react";
import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import _ from "lodash";
import { Pagination } from "@components/Pagination";
import { useLazyQuery } from '@apollo/client'
import { SynapseLogoSvg } from "@components/layouts/MainLayout/SynapseLogoSvg";

export default function Txs() {
  const search = useSearchParams()
  const p = Number(search.get('p'))
  // const kappaSearch = Number(search.get('kappa'))
  // const txhashSearch = Number(search.get('txhash'))


  const [transactionsArr, setTransactionsArr] = useState([])
  const [pending, setPending] = useState(false)
  const [wallet, setWallet] = useState("")
  const [walletLocale, setWalletLocale] = useState(true)
  const [minSize, setMinSize] = useState({ type: "USD", value: "" })
  const [maxSize, setMaxSize] = useState({ type: "USD", value: "" })
  const [chains, setChains] = useState([])
  const [chainsLocale, setChainsLocale] = useState(false)
  const [tokens, setTokens] = useState([])
  const [startDate, setStartDate] = useState(null)
  const [endDate, setEndDate] = useState(null)
  const [toTx, setToTx] = useState(true)
  const [fromTx, setFromTx] = useState(true)
  const [kappa, setKappa] = useState("")
  const [page, setPage] = useState(1)

  useEffect(() => {
    setPage(p)
    // setKappa(kappaSearch)
    // setToTx(txhashSearch)
    // setFromTx(txhashSearch)
    executeSearch(p)
  }, [p])

  // }, [p, kappaSearch, txhashSearch])

  const [getBridgeTransactions, { loading, error, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY, {
    onCompleted: (data) => {
      setTransactionsArr(data.bridgeTransactions);
    }
  })


  // Get initial data
  useEffect(() => {
    getBridgeTransactions({
      variables: {
        pending: pending,
        page: 1,
        useMv: true,
      },
    })
  }, [])

  const handlePending = (arg) => {
    setPending(arg)
    getBridgeTransactions({
      variables: {
        pending: arg,
        page: 1,
        useMv: true,
      },
    })
  }
  const createQueryField = (field, value, query) => {

    if (value && value !== "") {

      if (field === "endTime" || field === "startTime") {
        let timestamp = parseInt((new Date(value.$d).getTime() / 1000).toFixed(0))
        query[field] = timestamp
      } else if (field === "chainIDTo" || field === "chainIDFrom") {
        let chainIDs = []
        for (let i = 0; i < value.length; i++) {
          chainIDs.push(parseInt(CHAIN_ID_NAMES_REVERSE[value[i]]))

        }
        query[field] = chainIDs

      } else {
        query[field] = value
      }
    }
    return query
  }
  const executeSearch = (p) => {
    let queryPage = p ? p : page
    let variables = { page: queryPage===0 ? 1: queryPage, pending: pending, useMv: true }
    if (chains.length > 0) {
      if (chainsLocale) {
        variables = createQueryField("chainIDFrom", chains, variables)
      } else {
        variables = createQueryField("chainIDTo", chains, variables)
      }
    }
    if (walletLocale) {
      variables = createQueryField("addressFrom", wallet?.toLowerCase(), variables)
    } else {
      variables = createQueryField("addressTo", wallet?.toLowerCase(), variables)
    }

    if (minSize.value !== "") {
      if (minSize.type === "USD") {
        variables = createQueryField("minAmountUsd", parseInt(minSize.value), variables)
      } else {
        variables = createQueryField("minAmount", parseInt(minSize.value), variables)
      }
    }
    if (maxSize.value !== "") {
      if (maxSize.type === "USD") {
        variables = createQueryField("maxAmountUsd", parseInt(maxSize.value), variables)
      } else {
        variables = createQueryField("maxAmount", parseInt(maxSize.value), variables)
      }
    }
    variables = createQueryField("startTime", startDate, variables)
    variables = createQueryField("endTime",endDate, variables)
    if (kappa.length === 64) {
      variables = createQueryField("kappa", kappa?.toLowerCase(), variables)
    } else {
      variables = createQueryField("txnHash", kappa?.toLowerCase(), variables)

    }
    variables = createQueryField("pending", pending, variables)
    getBridgeTransactions({
      variables: variables,
    })
  }


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
          chainsLocale={chainsLocale}
          setChainsLocale={setChainsLocale}
          walletLocale={walletLocale}
          setWalletLocale={setWalletLocale}
        />
        {loading ? <div className="flex justify-center align-center w-full my-[100px] animate-spin"><SynapseLogoSvg /></div> : <BridgeTransactionTable queryResult={transactionsArr} />}


        <HorizontalDivider />
        <Pagination />
      </StandardPageContainer>
    </>
  )
}

