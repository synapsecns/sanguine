import { useState, useEffect } from 'react'

import _ from 'lodash'
import { useQuery, useLazyQuery } from '@apollo/client'

import {
  GET_BRIDGE_TRANSACTIONS_QUERY,
  BRIDGE_AMOUNT_STATISTIC,
} from '@graphql/queries'

import { Error } from '@components/Error'
import { Pagination } from '@components/Pagination'
import {
  AllTransactions,
  TransactionsLoader,
} from '@components/TransactionCard'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { SummaryStats } from './SummaryStats'
import { mode } from '@utils/math/statistics'
import { CopyButtonIcon } from '@components/CopyButtonIcon'

import { GetCsvButton } from '@components/GetCsvButton'

export function Address() {
  const [search, setSearch] = useSearchParams()
  const p = Number(search.get('page')) || 1

  const [page, setPage] = useState(p)
  const [transactions, setTransactions] = useState([])
  const [totalCount, setTotalCount] = useState(0)

  let { address } = useParams()

  const {
    error: totalError,
    loading: totalLoading,
    data: totalData,
  } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type: 'COUNT_TRANSACTIONS',
      duration: 'ALL_TIME',
      address: address,
    },
  })

  const [getBridgeTransactions, { loading, error, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY
  )

  useEffect(() => {
    if (data) {
      setTransactions(data.bridgeTransactions)
    }

    const num = Number(search.get('page'))

    if (num === 0) {
      setPage(1)
      getBridgeTransactions({
        variables: {
          address,
          page: 1,
        },
      })
    } else {
      setPage(num)

      getBridgeTransactions({
        variables: {
          address,
          page: num,
        },
      })
    }
  }, [data, search])

  useEffect(() => {
    if (totalData) {
      setTotalCount(totalData.bridgeAmountStatistic.value)
    }
  }, [totalData])

  function nextPage() {
    let newPage = page + 1
    setPage(newPage)
    setSearch({ page: newPage })

    getBridgeTransactions({
      variables: { address, page: newPage },
    })
  }

  function prevPage() {
    if (page > 1) {
      let newPage = page - 1
      setPage(newPage)
      setSearch({ page: newPage })
      getBridgeTransactions({
        variables: { address, page: newPage },
      })
    }
  }

  function resetPage() {
    setPage(1)
    setSearch({ page: 1 })
    getBridgeTransactions({
      variables: { address, page: 1 },
    })
  }

  let content

  if (totalLoading || loading) {
    content = <TransactionsLoader number={5} />
  } else if (totalError || error) {
    content = (
      <Error
        text="Sorry, there was a problem associated with that address."
        param={address}
        subtitle="Unknown"
      />
    )
  } else if (transactions.length === 0) {
    content = (
      <Error
        text="Sorry, there are no transactions associated with that address."
        param={address}
        subtitle="No transactions"
      />
    )
  } else {
    let bridgeTransactions = transactions
    bridgeTransactions = _.orderBy(bridgeTransactions, 'fromInfo.time', [
      'desc',
    ])

    let originChainIds = _.map(bridgeTransactions, 'fromInfo.chainId').filter(
      (n) => n
    )
    let destinationChainIds = _.map(
      bridgeTransactions,
      'toInfo.chainId'
    ).filter((n) => n)

    let originChainIdMode = mode(originChainIds)
    let destinationChainIdMode = mode(destinationChainIds)

    content = (
      <>
        <SummaryStats
          address={address}
          topOriginChainId={originChainIdMode}
          topDestinationChainId={destinationChainIdMode}
        />
        <AllTransactions className="mb-5" txns={bridgeTransactions} />
        <Pagination
          totalCount={totalCount}
          page={page}
          resetPage={resetPage}
          prevPage={prevPage}
          nextPage={nextPage}
        />
      </>
    )
  }
  return (
    <StandardPageContainer
      title="Address"
      subtitle={
        <div className="flex ">
          <span className="font-mono sm:text-lg text-slate-400 ">
            {address}
          </span>
          <span className="ml-2 sm:mt-1">
            <CopyButtonIcon
              text={address}
              className="text-slate-600 hover:text-slate-300"
              tooltipText="address"
            />
          </span>
          <span className="ml-2 sm:mt-1">
            <GetCsvButton address={address} />
          </span>
        </div>
      }
    >
      {content}
    </StandardPageContainer>
  )
}
