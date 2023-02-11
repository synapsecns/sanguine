import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'
import {API_URL} from '@graphql'
import numeral from 'numeral'
import {BridgeTransactionTable} from "@components/BridgeTransaction/BridgeTransactionTable";

import _ from 'lodash'
import {useEffect, useState} from 'react'

import {useLazyQuery, useQuery} from '@apollo/client'

import {BRIDGE_AMOUNT_STATISTIC, GET_BRIDGE_TRANSACTIONS_QUERY,} from '@graphql/queries'
import {StandardPageContainer} from '@components/layouts/StandardPageContainer'
import {Pagination} from '@components/Pagination'

import Grid from '@components/tailwind/Grid'

import {TokenOnChain} from '@components/misc/TokenOnChain'

import {StatCard} from '@components/pages/Home/Stats'

import {useRouter} from 'next/router'
import {useSearchParams} from 'next/navigation'


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

export default function tokenAddressRoute({
  allTimeBridgeVolume,
  allTimeTransactionCount,
  allTimeAddresses,
  bridgeTransactions,
}) {
  const router = useRouter()
  const { tokenAddress } = router.query
  const search = useSearchParams()
  const p = Number(search.get('page')) || 1
  const chainId = Number(search.get('chainId')) || 1
  // const chainId = Number(search.get('chainId'))

  const [page, setPage] = useState(p)
  const [transactions, setTransactions] = useState([])

  const [getBridgeTransactions, { error: pageError, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY
  )

  const nextPage = () => {
    let newPage = page + 1
    setPage(newPage)
    // setSearch({ page: newPage, chainId })

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
      // setSearch({ page: newPage, chainId })
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
    // setSearch({ page: 1, chainId })
    getBridgeTransactions({
      variables: {
        tokenAddress,
        chainId: Number(chainId),
        page: 1,
      },
    })
  }

  let content

    bridgeTransactions = _.orderBy(bridgeTransactions, 'fromInfo.time', [
      'desc',
    ]).slice(0, 10)

    content = <BridgeTransactionTable queryResult={bridgeTransactions} />

  let title = <TokenOnChain tokenAddress={tokenAddress} chainId={chainId} />

  return (
    <StandardPageContainer title={title}>
      <Grid cols={{ sm: 1, md: 3, lg: 3 }} gap={4} className="my-5">
        <StatCard title="Volume" active={true} duration="All-Time">
          <div className="text-4xl font-bold text-white">
            ${allTimeBridgeVolume}
          </div>
        </StatCard>
        <StatCard title="Transaction Count" active={true} duration="All-Time">
          <div className="text-4xl font-bold text-white">
            {allTimeTransactionCount}
          </div>
        </StatCard>
        <StatCard title="Addresses" active={true} duration="All-Time">
          <div className="text-4xl font-bold text-white">
            {allTimeAddresses}
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

export async function getServerSideProps(context) {
  const { tokenAddress, chainId } = context.query
  const { data: allTimeBridgeVolume } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      chainId: chainId,
      duration: 'ALL_TIME',
      tokenAddress: tokenAddress,
      type: 'TOTAL_VOLUME_USD',
    },
  })

  const { data: allTimeTransactionCount } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      chainId: chainId,
      duration: 'ALL_TIME',
      tokenAddress: tokenAddress,
      type: 'COUNT_TRANSACTIONS',
    },
  })

  const { data: allTimeAddresses } = await client.query({
    query: BRIDGE_AMOUNT_STATISTIC,
    variables: {
      chainId: chainId,
      duration: 'ALL_TIME',
      tokenAddress: tokenAddress,
      type: 'COUNT_ADDRESSES',
    },
  })

  const { data: bridgeTransactions } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: chainId,
      tokenAddress: tokenAddress,
      page: 1,
    },
  })

  return {
    props: {
      allTimeBridgeVolume: normalizeValue(
        allTimeBridgeVolume?.bridgeAmountStatistic?.value
      ),
      allTimeTransactionCount: normalizeValue(
        allTimeTransactionCount?.bridgeAmountStatistic?.value
      ),
      allTimeAddresses: normalizeValue(
        allTimeAddresses?.bridgeAmountStatistic?.value
      ),
      bridgeTransactions: bridgeTransactions,
    },
  }
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
