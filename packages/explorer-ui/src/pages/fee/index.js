/* eslint-disable react/jsx-key */
import { addOrSetObject, OverviewChart } from 'components/Charts'
import { Header } from 'components/Header'
import { Card, GradientCard, PrimaryContainer } from 'components/Sections'
import { StatisticsSection } from 'components/Statistics'
import { LabelSwitch } from 'components/Switch'
import { OverviewTable } from 'components/Table'
import { CHAIN_ID_NAMES_REVERSE } from 'constants/chains'
import _ from 'lodash'
import { fillInCumulativeData } from 'pages/index'
import { useState } from 'react'
import useSWR from 'swr'
import { fetcherMany } from 'utils/fetcher'

const feesColumns = ['Name', 'Total Fees']

const Page = () => {
  const [showAggregatedFees, setShowAggregatedFees] = useState(false)
  const [showCumulativeFees, setShowCumulativeFees] = useState(false)
  const [showMonthlyFees, setShowMonthlyFees] = useState(false)

  const chainNames = _.keys(CHAIN_ID_NAMES_REVERSE)

  const { data: poolData } = useSWR(
    _.flatten(
      _.map(chainNames, (chain) =>
        _.map(['nusd', 'neth'], (pool) => `https://analytics-api.bridgesyn.com/api/v1/analytics/pools/volume/${chain}/${pool}`),
      ),
    ),
    fetcherMany,
  )

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

  // Manipulate `tokenData` into a similar schema to `volume` (src/pages/index.js) so i can copy pasta
  // `volumeChartData` (src/pages/index.js) code, because i honestly have no clue how this shit works.
  let totalFees = {}
  const bridgeFeesTotals = {}
  _.map(chainNames, (chain) => (bridgeFeesTotals[chain] = 0))

  _.forEach(tokenData, (obj) => {
    const chain = _.keys(obj.stats.volume)[0]

    for (const [date, data] of Object.entries(obj.data)) {
      if (!(date in totalFees)) {
        totalFees[date] = { [chain]: data.price_usd }
      } else {
        addOrSetObject(totalFees[date], chain, data.price_usd)
      }

      bridgeFeesTotals[chain] += data.price_usd
    }
  })

  const poolFeesTotals = {}
  _.map(chainNames, (chain) => (poolFeesTotals[chain] = 0))

  let i = 0
  _.map(poolData, (obj) => {
    // poolData is given in order, [eth (nUSD), eth(nETH), optimism (nUSD), ...]
    const chain = chainNames[Math.floor(i++ / 2)]

    for (const [date, data] of Object.entries(obj)) {
      for (const _data of Object.values(data)) {
        if (!(date in totalFees)) {
          totalFees[date] = { [chain]: _data.admin_fees_usd }
        } else {
          addOrSetObject(totalFees[date], chain, _data.admin_fees_usd)
        }

        poolFeesTotals[chain] += _data.admin_fees_usd
      }
    }
  })

  // Make sure dates are ordered (oldest first, newest last).
  totalFees = _.sortBy(_.keys(totalFees)).reduce((x, key) => ((x[key] = totalFees[key]), x), {})

  if (showAggregatedFees) {
    for (const [key, value] of Object.entries(totalFees)) {
      totalFees[key].total = _.sum(_.values(value))
    }
  }

  const previousFees = {}
  const feeChartData = _.orderBy(
    _.map(totalFees, (value, key) => {
      const { ...rest } = value
      _.map(rest, (value, key) => {
        if (previousFees[key]) {
          previousFees[key] += value
        } else {
          previousFees[key] = value
        }
      })

      const cumulativeOutput = {}
      if (showCumulativeFees) {
        cumulativeOutput['date'] = key
        _.map(rest, (value, key) => {
          cumulativeOutput[key] = previousFees[key]
        })

        return cumulativeOutput
      }

      return { date: key, ...rest }
    }),
    'date',
  )

  if (showCumulativeFees) {
    fillInCumulativeData({ feeChartData })
  }

  const bridgeFeeTableData = _.orderBy(
    _.map(bridgeFeesTotals, (value, key) => {
      return { name: key, value: value }
    }),
    'value',
    'desc',
  )

  const poolFeeTableData = _.orderBy(
    _.map(poolFeesTotals, (value, key) => {
      return { name: key, value: value }
    }),
    'value',
    'desc',
  )

  return (
    <div>
      <Header />
      <PrimaryContainer>
        <h1 className="mt-8 mb-8 text-4xl font-bold text-gray-200">Fee Statistics</h1>
        <StatisticsSection />
        <div className="mb-8 text-gray-200">
          <div>
            <GradientCard>
              <div className="flex justify-between">
                <h2 className="mb-2 text-2xl font-bold">Protocol Revenue</h2>
                <div className="space-y-2">
                  <LabelSwitch
                    enabled={showAggregatedFees}
                    setEnabled={setShowAggregatedFees}
                    text="Show aggregated fees"
                  />
                  <LabelSwitch
                    enabled={showCumulativeFees}
                    setEnabled={setShowCumulativeFees}
                    text="Show cumulative fees"
                  />
                  <LabelSwitch enabled={showMonthlyFees} setEnabled={setShowMonthlyFees} text="Show monthly fees" />
                </div>
              </div>
              <OverviewChart
                data={feeChartData}
                isCumulativeData={showCumulativeFees}
                showAggregated={showAggregatedFees}
                monthlyData={showMonthlyFees}
                currency
              />
              <br />
              <p className="mr-4 text-gray-400">
                Note: Protocol revenues currently accrue to the Synapse treasury (after gas costs). It does not include
                revenues earned by liquidity providers.
              </p>  
            </GradientCard>
          </div>
        </div>
        <br />  
        <div className="grid gap-8 mb-8 text-gray-200 xl:grid-cols-12">
          {bridgeFeeTableData && (
            <div className="xl:col-span-6">
              <h2 className="mb-4 text-2xl font-bold">Bridge Fees Accumulated by Chain</h2>
              {<OverviewTable columns={feesColumns} data={bridgeFeeTableData} currency></OverviewTable>}
            </div>
          )}
          {poolFeeTableData && (
            <div className="xl:col-span-6">
              <h2 className="mb-4 text-2xl font-bold">Admin Pool Fees Accumulated by Chain</h2>
              <OverviewTable columns={feesColumns} data={poolFeeTableData} currency></OverviewTable>
            </div>
          )}
        </div>
      </PrimaryContainer>
    </div>
  )
}

export default Page
