import _ from 'lodash'
import { useState, useEffect } from 'react'

import { useLazyQuery, useQuery } from '@apollo/client'
import { getAddress } from '@ethersproject/address'
import numeral from 'numeral'

import {
  TransactionCard,
  TransactionCardLoader,
} from '@components/TransactionCard'
import {
  BRIDGE_AMOUNT_STATISTIC,
  GET_BRIDGE_TRANSACTIONS_QUERY,
} from '@graphql/queries'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { Pagination } from '@components/Pagination'

import Grid from '@components/tailwind/Grid'

import { TokenOnChain } from '@components/misc/TokenOnChain'

import { StatCard } from '@pages/Home/Stats'

export function TokenAddress() {
  const [search, setSearch] = useSearchParams()
  const p = Number(search.get('page')) || 1
  const chainId = Number(search.get('chainId'))

  const [page, setPage] = useState(p)
  const [transactions, setTransactions] = useState([])
  const { tokenAddress } = useParams()

  const [getBridgeTransactions, { error: pageError, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY
  )

  useEffect(() => {
    if (data) {
      setTransactions(data.bridgeTransactions, {
        variables: {
          tokenAddress,
          chainId: Number(chainId),
        },
      })
    }

    const num = Number(search.get('page'))

    if (num === 0) {
      setPage(1)
      getBridgeTransactions({
        variables: {
          tokenAddress,
          chainId: Number(chainId),
          page: 1,
        },
      })
    } else {
      setPage(num)
      getBridgeTransactions({
        variables: {
          tokenAddress,
          chainId: Number(chainId),
          page: num,
        },
      })
    }
  }, [data, search, chainId, tokenAddress])

  const nextPage = () => {
    let newPage = page + 1
    setPage(newPage)
    setSearch({ page: newPage, chainId })

    getBridgeTransactions({
      variables: {
        tokenAddress,
        chainId: Number(chainId),
        page: newPage,
      },
    })
  }

  const prevPage = () => {
    if (page > 1) {
      let newPage = page - 1
      setPage(newPage)
      setSearch({ page: newPage, chainId })
      getBridgeTransactions({
        variables: {
          tokenAddress,
          chainId: Number(chainId),
          page: newPage,
        },
      })
    }
  }

  const resetPage = () => {
    setPage(1)
    setSearch({ page: 1, chainId })
    getBridgeTransactions({
      variables: {
        tokenAddress,
        chainId: Number(chainId),
        page: 1,
      },
    })
  }

  let content

  if (!data) {
    content = [...Array(10).keys()].map((i) => (
      <TransactionCardLoader key={i} />
    ))
  } else if (pageError) {
    content = 'Error'
  } else {
    let bridgeTransactions = transactions

    bridgeTransactions = _.orderBy(bridgeTransactions, 'fromInfo.time', [
      'desc',
    ]).slice(0, 10)

    content = bridgeTransactions.map((txn, i) => (
      <TransactionCard txn={txn} key={i} />
    ))
  }

  let title = <TokenOnChain tokenAddress={tokenAddress} chainId={chainId} />

  return (
    <StandardPageContainer title={title}>
      <Grid cols={{ sm: 1, md: 3, lg: 3 }} gap={4} className="my-5">
        <StatCard title="Volume" active={true} duration="All-Time">
          <div className="text-4xl font-bold text-white">
            $
            {getBridgeVolume({
              chainId,
              tokenAddress,
              duration: 'ALL_TIME',
            })}
          </div>
        </StatCard>
        <StatCard title="Transaction Count" active={true} duration="All-Time">
          <div className="text-4xl font-bold text-white">
            {getTransactionCount({
              chainId,
              tokenAddress,
              duration: 'ALL_TIME',
            })}
          </div>
        </StatCard>
        <StatCard title="Addresses" active={true} duration="All-Time">
          <div className="text-4xl font-bold text-white">
            {getAddressCount({
              chainId,
              tokenAddress,
              duration: 'ALL_TIME',
            })}
          </div>
        </StatCard>
      </Grid>
      {content}
      <Pagination
        page={page}
        resetPage={resetPage}
        prevPage={prevPage}
        nextPage={nextPage}
      />
    </StandardPageContainer>
  )
}

function getTransactionCount({ chainId, tokenAddress, duration }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      chainId: chainId && Number(chainId),
      duration,
      tokenAddress: getAddress(tokenAddress),
      type: 'COUNT_TRANSACTIONS',
    },
  })

  return normalizeValue(data?.bridgeAmountStatistic?.value)
}

function getBridgeVolume({ chainId, tokenAddress, duration }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      chainId: Number(chainId),
      duration,
      tokenAddress: getAddress(tokenAddress),
      type: 'TOTAL_VOLUME_USD',
    },
  })

  return normalizeValue(data?.bridgeAmountStatistic?.value)
}

function getAddressCount({ chainId, tokenAddress, duration }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      chainId: Number(chainId),
      duration,
      tokenAddress: getAddress(tokenAddress),
      type: 'COUNT_ADDRESSES',
    },
  })

  return normalizeValue(data?.bridgeAmountStatistic?.value)
}

function normalizeValue(value) {
  if (value >= 1000000000) {
    return (
      numeral(value / 1000000000)
        .format('0.00')
        .toString() + 'B'
    )
  } else if (value >= 1000000) {
    return (
      numeral(value / 1000000)
        .format('0.00')
        .toString() + 'M'
    )
  }
  return numeral(value).format('0,0').toString()
}
