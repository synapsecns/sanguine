import { Home } from '@components/pages/Home'
import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import {
  AMOUNT_STATISTIC,
  DAILY_STATISTICS_BY_CHAIN,
  RANKED_CHAINIDS_BY_VOLUME,
} from '@graphql/queries'
import { API_URL } from '@graphql'

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  ssrMode: true,
  cache: new InMemoryCache(),
  fetchPolicy: 'cache-and-network',
  fetchOptions: {
    mode: 'no-cors',
  },
})

function Index({
  // dailyStats,
  // rankedChainIDs,
  // totalVolume,
  // totalFee,
  // totalAddresses,
  // totalTransactions,
}) {
  // console.log("YO MAMA", totalVolume,
  // totalFee,
  // totalAddresses,
  // totalTransactions)

  return (
    <Home
      // dailyStats={dailyStats.dailyStatisticsByChain}
      // rankedChainIDs={rankedChainIDs.rankedChainIDsByVolume}
      // totalVolume={totalVolume.amountStatistic.value}
      // totalFee={totalFee.amountStatistic.value}
      // totalAddresses={totalAddresses.amountStatistic.value}
      // totalTransactions={totalTransactions.amountStatistic.value}
    />
  )
}

export default Index

// export async function getServerSideProps() {
//   const { data: dailyStats } = await client.query({
//     query: DAILY_STATISTICS_BY_CHAIN,
//     variables: {
//       type: 'VOLUME',
//       duration: 'PAST_MONTH',
//     },
//   })

//   const { data: rankedChainIDs } = await client.query({
//     query: RANKED_CHAINIDS_BY_VOLUME,
//     variables: {
//       duration: 'PAST_MONTH',
//     },
//   })

//   const { data: totalVolume } = await client.query({
//     query: AMOUNT_STATISTIC,
//     variables: {
//       platform: "ALL",
//       duration: "ALL_TIME",
//       type: "TOTAL_VOLUME_USD",
//     },
//   })
//   const { data: totalFee } = await client.query({
//     query: AMOUNT_STATISTIC,
//     variables: {
//       platform: "ALL",
//       duration: "ALL_TIME",
//       type: "TOTAL_FEE_USD",
//     },
//   })
//   const { data: totalAddresses } = await client.query({
//     query: AMOUNT_STATISTIC,
//     variables: {
//       platform: "ALL",
//       duration: "ALL_TIME",
//       type: "COUNT_ADDRESSES",
//     },
//   })

//   const { data: totalTransactions } = await client.query({
//     query: AMOUNT_STATISTIC,
//     variables: {
//       platform: "ALL",
//       duration: "ALL_TIME",
//       type: "COUNT_TRANSACTIONS",
//     },
//   })

//   return {
//     props: {
//       dailyStats: dailyStats,
//       rankedChainIDs: rankedChainIDs,
//       totalVolume: totalVolume,
//       totalFee: totalFee,
//       totalAddresses: totalAddresses,
//       totalTransactions: totalTransactions,
//     }, // will be passed to the page component as props
//   }
// }
