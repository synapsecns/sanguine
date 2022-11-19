import { Card } from 'components/Sections'
import _ from 'lodash'
import useSWR from 'swr'
import {fetcher, fetcherMany} from 'utils/fetcher'
import {CHAIN_ID_NAMES_REVERSE} from "../constants/chains";

export const totalRevenue = () => {
  const chainNames = _.keys(CHAIN_ID_NAMES_REVERSE)

  let { data: chainTokens } = useSWR(
      _.map(chainNames, (chain) => `https://analytics-api.bridgesyn.com/api/v1/analytics/fees/bridge/${chain}/`),
      fetcherMany,
  )

  chainTokens = _.zip(chainNames, chainTokens)

  const { data: tokenData } = useSWR(
      _.flatten(
          _.map(chainTokens, (data) => {
            if (data.length == 1) return

            const [chain, { valids: tokens }] = data
            return _.map(tokens, (token) => `https://analytics-api.bridgesyn.com/api/v1/analytics/fees/bridge/${chain}/${token}`)
          }),
      ),
      fetcherMany,
  )


    const { data: poolData } = useSWR(
        _.flatten(
            _.map(chainNames, (chain) =>
                _.map(['nusd', 'neth'], (pool) => `https://analytics-api.bridgesyn.com/api/v1/analytics/pools/volume/${chain}/${pool}`),
            ),
        ),
        fetcherMany,
    )

    const bridgeFees = _.reduce(tokenData,  (acc, rev) => {
        if (rev.stats === undefined) {
            acc += 0
            return acc
        }
        acc += rev.stats.usd.adjusted
        return acc
    }, 0)

    const poolFees = _.reduce(poolData, (acc, rev) => {
        acc += Object.values(rev).reduce((a, fees) => {
            a += Object.values(fees).reduce((tot, b) => {
                tot += b.admin_fees_usd

                return tot
            }, 0)

            return a
        }, 0)

        return acc
    }, 0)

    // don't return non-aggregated data
    if (bridgeFees === 0 || poolFees === 0){
        return undefined
    }

    return bridgeFees + poolFees
}

export const StatisticsSection = () => {
  const { data: coingeckoData } = useSWR(
    'https://api.coingecko.com/api/v3/coins/synapse-2?localization=false&tickers=false&community_data=false&developer_data=false',
    fetcher,
  )
  const { data: poolData } = useSWR('https://analytics-api.bridgesyn.com/api/v1/analytics/pools/volume/total', fetcher)
  const { data: bridgeData } = useSWR('https://analytics-api.bridgesyn.com/api/v1/analytics/volume/total/in', fetcher)

  const rev = totalRevenue()
  const synPrice = coingeckoData?.market_data?.current_price?.usd

  let totalBridge, totalPool
  if (bridgeData?.totals) {
    totalBridge = _.sum(_.values(bridgeData?.totals))
  }

  if (poolData?.totals) {
    totalPool = _.sum(_.values(poolData?.totals))
  }

  const transformB = (val) => `$${val ? _.round(val / 1000000000, 2) : '-'}B`
  const transformM = (val) => `$${val ? _.round(val / 1000000, 2) : '-'}M`

  return (
    <div className="grid gap-4 mb-8 text-gray-200 md:grid-cols-2 lg:grid-cols-4">
      <Card>
        <h2 className="mb-2 text-xl font-bold text-gray-400">Total Bridge Volume</h2>
        <p className="text-4xl font-bold">{transformB(totalBridge)}</p>
      </Card>
      <Card>
        <h2 className="mb-2 text-xl font-bold text-gray-400">Total Pool Volume</h2>
        <p className="text-4xl font-bold">{transformB(totalPool)}</p>
      </Card>
      <Card>
        <h2 className="mb-2 text-xl font-bold text-gray-400">Total Revenue</h2>
        <p className="text-4xl font-bold">{`${transformM(rev)}`}</p>
      </Card>
      <Card>
        <h2 className="mb-2 text-xl font-bold text-gray-400">Current Price</h2>
        <p className="text-4xl font-bold">{`$${synPrice ? synPrice.toFixed(2) : '-'}`}</p>
      </Card>
    </div>
  )
}
