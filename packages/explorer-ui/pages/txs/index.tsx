import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { UniversalSearch } from '@components/pages/Home/UniversalSearch'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
import { CHAIN_ID_NAMES_REVERSE } from '@constants/networks'
import { useSearchParams } from 'next/navigation'
import { API_URL } from '@graphql'
import { useState, useEffect } from 'react'
import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import _ from 'lodash'
import { Pagination } from '@components/Pagination'
import { useLazyQuery } from '@apollo/client'
import { SynapseLogoSvg } from '@components/layouts/MainLayout/SynapseLogoSvg'
import { checksumAddress, checkAddressChecksum } from '@utils/checksum'

export default function Txs() {
  const search = useSearchParams()
  const p = Number(search.get('p'))
  const hashSearch = String(search.get('hash'))

  const [transactionsArr, setTransactionsArr] = useState([])
  const [pending, setPending] = useState(false)
  const [wallet, setWallet] = useState('')
  const [walletLocale, setWalletLocale] = useState(true)
  const [minSize, setMinSize] = useState({ type: 'USD', value: '' })
  const [maxSize, setMaxSize] = useState({ type: 'USD', value: '' })
  const [chains, setChains] = useState([])
  const [chainsLocale, setChainsLocale] = useState(false)
  const [tokens, setTokens] = useState([])
  const [startDate, setStartDate] = useState(null)
  const [endDate, setEndDate] = useState(null)
  const [toTx, setToTx] = useState(true)
  const [fromTx, setFromTx] = useState(true)
  const [kappa, setKappa] = useState('')
  const [page, setPage] = useState(1)

  useEffect(() => {
    setPage(p)
    const hash = hashSearch === 'null' ? '' : hashSearch
    setKappa(hash)
    executeSearch(p, hash)
  }, [p, hashSearch])

  useEffect(() => {
    executeSearch()
  }, [pending])

  const [getBridgeTransactions, { loading, error, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY,
    {
      onCompleted: (data) => {
        setTransactionsArr(data.bridgeTransactions)
      },
    }
  )

  // Get initial data
  useEffect(() => {
    getBridgeTransactions({
      variables: {
        pending,
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
    if (value && value !== '') {
      if (field === 'endTime' || field === 'startTime') {
        const timestamp = parseInt(
          (new Date(value.$d).getTime() / 1000).toFixed(0)
        )
        query[field] = timestamp
      } else if (field === 'chainIDTo' || field === 'chainIDFrom') {
        const chainIDs = []
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
  const executeSearch = (p?: any, txOrKappaHash?: any) => {
    const queryPage = p ? p : page
    const queryKappa = txOrKappaHash ? txOrKappaHash : kappa
    if (queryKappa && queryKappa != '' && queryKappa.length < 64) {
      alert('Invalid hash entered')
      return
    }
    if (wallet && wallet != '' && wallet.length !== 42) {
      alert('Invalid wallet address entered')
      return
    }
    let variables = {
      page: queryPage === 0 ? 1 : queryPage,
      pending,
      useMv: true,
    }
    if (chains.length > 0) {
      if (chainsLocale) {
        variables = createQueryField('chainIDFrom', chains, variables)
      } else {
        variables = createQueryField('chainIDTo', chains, variables)
      }
    }
    if (wallet) {
      variables = createQueryField(
        'addressFrom',
        checksumAddress(wallet),
        variables
      )
    } else {
      variables = createQueryField(
        'addressTo',
        checksumAddress(wallet),
        variables
      )
    }

    if (minSize.value !== '') {
      if (minSize.type === 'USD') {
        variables = createQueryField(
          'minAmountUsd',
          parseInt(minSize.value),
          variables
        )
      } else {
        variables = createQueryField(
          'minAmount',
          parseInt(minSize.value),
          variables
        )
      }
    }
    if (maxSize.value !== '') {
      if (maxSize.type === 'USD') {
        variables = createQueryField(
          'maxAmountUsd',
          parseInt(maxSize.value),
          variables
        )
      } else {
        variables = createQueryField(
          'maxAmount',
          parseInt(maxSize.value),
          variables
        )
      }
    }
    variables = createQueryField('startTime', startDate, variables)
    variables = createQueryField('endTime', endDate, variables)
    if (queryKappa.length === 64) {
      variables = createQueryField('kappa', queryKappa, variables)
    } else {
      variables = createQueryField('txnHash', queryKappa, variables)
    }
    variables = createQueryField('pending', pending, variables)
    getBridgeTransactions({
      variables,
    })
  }

  return (
    <>
      <StandardPageContainer title="Synapse Analytics">
        <div className="flex items-center mt-10 mb-2">
          <h3 className="text-white text-2xl font-semibold">
            Bridge Transactions
          </h3>
        </div>

        <HorizontalDivider />
        <UniversalSearch
          placeholder={'Search bridge transactions by bridge tx'}
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
          fromTx={fromTx}
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
        {loading ? (
          <div className="flex justify-center align-center w-full my-[100px]  ">
            <div className="w-[39px] animate-spin">
              <SynapseLogoSvg />
            </div>
          </div>
        ) : (
          <BridgeTransactionTable queryResult={transactionsArr} />
        )}

        <HorizontalDivider />
        <Pagination />
      </StandardPageContainer>
    </>
  )
}
