import { OverviewChart } from 'components/Charts'
import { Header } from 'components/Header'
import { GradientCard, PrimaryContainer } from 'components/Sections'
import { StatisticsSection } from 'components/Statistics'
import { LabelSwitch } from 'components/Switch'
import { OverviewTable } from 'components/Table'
import _ from 'lodash'
import { useState } from 'react'
import useSWR from 'swr'
import { fetcher } from 'utils/fetcher'

const volumeColumns = ['Name', 'Total Volume']
const txCountColumns = ['Name', 'Total TX Count', 'Avg TX Size']

export const fillInCumulativeData = (data) => {
  _.map(data, (value, index) => {
    for (const [chain, _value] of _.entries(value)) {
      if (!['date', 'total'].includes(chain)) {
        // We can safely assume we iterate from oldest date to newest,
        // so lets see if the next day's data is given as expected.
        const nextDay = data[index + 1]
        if (nextDay && nextDay[chain] === undefined) {
          nextDay[chain] = _value
        }
      }
    }
  })
}

const Page = () => {
  const [showAggregatedVolume, setShowAggregatedVolume] = useState(false)
  const [showAggregatedTxCount, setShowAggregatedTxCount] = useState(false)
  const [showCumulativeVolume, setShowCumulativeVolume] = useState(false)
  const [showCumulativeTxCount, setShowCumulativeTxCount] = useState(false)
  const [showMonthlyTxCount, setMonthlyTxCount] = useState(false)
  const [showMonthlyVolume, setMonthlyVolume] = useState(false)

  const { data: volume, error: volumeError } = useSWR(
    'https://analytics-api.bridgesyn.com/api/v1/analytics/volume/total/in',
    fetcher,
  )
  const { data: txCount, error: totalTxCountError } = useSWR(
    'https://analytics-api.bridgesyn.com/api/v1/analytics/volume/total/tx_count',
    fetcher,
  )

  const previousVolumes = {}
  const volumeChartData = _.orderBy(
    _.map(volume?.data, (value, key) => {
      const { ...rest } = value
      _.map(rest, (value, key) => {
        if (previousVolumes[key]) {
          previousVolumes[key] += value
        } else {
          previousVolumes[key] = value
        }
      })

      const cumulativeOutput = {}
      if (showCumulativeVolume) {
        cumulativeOutput['date'] = key
        _.map(rest, (value, key) => {
          cumulativeOutput[key] = previousVolumes[key]
        })

        return cumulativeOutput
      }

      return { date: key, ...rest }
    }),
    'date',
  )

  if (showCumulativeVolume) {
    fillInCumulativeData(volumeChartData)
  }

  const volumeTableData = _.orderBy(
    _.map(volume?.totals, (value, key) => {
      return { name: key, value: value }
    }),
    'value',
    'desc',
  )

  const txCountObject = {}
  const txCountTableObject = {}
  const txCountTableObjectData = []
  _.map(txCount?.data, (dates, chain) => {
    _.map(dates, (value, key) => {
      if (key === 'total') {
        txCountTableObject[chain] = value
      } else if (_.isObject(txCountObject[key])) {
        txCountObject[key][chain] = value
      } else {
        txCountObject[key] = { [chain]: value }
      }
    })
  })

  const txCountChartData = _.orderBy(
    _.map(txCountObject, (value, key) => {
      const { ...rest } = value
      const total = _.sum(_.values(rest))
      return { date: key, total, ...rest }
    }),
    'date',
  )

  const previousTxCounts = {}
  const newTxCountChartData = _.map(txCountChartData, (value, index) => {
    const { ...rest } = value

    _.map(rest, (value, key) => {
      if (key === 'date') {
        previousTxCounts[key] = value
      } else if (previousTxCounts[key]) {
        previousTxCounts[key] += value
      } else {
        previousTxCounts[key] = value
      }
    })

    const cumulativeOutput = {}
    if (showCumulativeTxCount) {
      _.map(rest, (value, key) => {
        cumulativeOutput[key] = previousTxCounts[key]
      })

      return cumulativeOutput
    }

    return { ...rest }
  })

  if (showCumulativeTxCount) {
    fillInCumulativeData(newTxCountChartData)
  }

  for (const volumeData of volumeTableData) {
    const value = txCountTableObject[volumeData.name]

    txCountTableObjectData.push({
      name: volumeData.name,
      value,
      volume: volumeData?.value,
      averageTxSize: volumeData?.value / value,
    })
  }

  const txCountTableData = _.orderBy(txCountTableObjectData, 'value', 'desc')

  return (
    <div>
      <Header />
      <PrimaryContainer>
        <h1 className="mt-8 mb-8 text-4xl font-bold text-gray-200">Synapse Analytics</h1>
        <StatisticsSection />
        <div className="mb-8 text-gray-200">
          <div>
            <GradientCard>
              <div className="flex justify-between">
                <h2 className="mb-2 text-2xl font-bold">Bridge Volume</h2>
                <div className="space-y-2">
                  <LabelSwitch
                    enabled={showAggregatedVolume}
                    setEnabled={setShowAggregatedVolume}
                    text="Show aggregated volume"
                  />
                  <LabelSwitch
                    enabled={showCumulativeVolume}
                    setEnabled={setShowCumulativeVolume}
                    text="Show cumulative volume"
                  />
                  <LabelSwitch enabled={showMonthlyVolume} setEnabled={setMonthlyVolume} text="Show monthly data" />
                </div>
              </div>
              <OverviewChart
                data={volumeChartData}
                isCumulativeData={showCumulativeVolume}
                showAggregated={showAggregatedVolume}
                monthlyData={showMonthlyVolume}
                currency
              />
            </GradientCard>
          </div>
          <div>
            <GradientCard>
              <div className="flex justify-between">
                <h2 className="mb-2 text-2xl font-bold">Transaction Count</h2>
                <div className="space-y-2">
                  <LabelSwitch
                    enabled={showAggregatedTxCount}
                    setEnabled={setShowAggregatedTxCount}
                    text="Show aggregated transaction count"
                  />
                  <LabelSwitch
                    enabled={showCumulativeTxCount}
                    setEnabled={setShowCumulativeTxCount}
                    text="Show cumulative transaction count"
                  />
                  <LabelSwitch enabled={showMonthlyTxCount} setEnabled={setMonthlyTxCount} text="Show monthly data" />
                </div>
              </div>
              <OverviewChart
                data={newTxCountChartData}
                isCumulativeData={showCumulativeTxCount}
                showAggregated={showAggregatedTxCount}
                monthlyData={showMonthlyTxCount}
              />
            </GradientCard>
          </div>
        </div>
        <div className="grid gap-8 mb-8 text-gray-200 xl:grid-cols-12">
          {volume && (
            <div className="xl:col-span-5">
              <h2 className="mb-4 text-2xl font-bold">Total Bridge Volume by Chain</h2>
              {<OverviewTable columns={volumeColumns} data={volumeTableData} currency></OverviewTable>}
            </div>
          )}
          {txCount && (
            <div className="xl:col-span-7">
              <h2 className="mb-4 text-2xl font-bold">Total Transaction Count by Chain</h2>
              <OverviewTable columns={txCountColumns} data={txCountTableData}></OverviewTable>
            </div>
          )}
        </div>
      </PrimaryContainer>
    </div>
  )
}

export default Page
