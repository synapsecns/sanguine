import _ from 'lodash'
import { Bar, BarChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import { CurrencyTooltip, NumericTooltip } from '@components/misc/ToolTip'
const formatShort = new Intl.NumberFormat('en-US', { notation: 'compact', maximumFractionDigits: 1 })
const formatMonth = new Intl.DateTimeFormat('en-US', {
  year: 'numeric',
  month: 'long',
})
const formatTotalUsdVolumes = (totalUsdVolumes) => {
  if (totalUsdVolumes > 1000000000) {
    return `${_.round(totalUsdVolumes / 1000000000, 3)}B`
  } else if (totalUsdVolumes > 100000) {
    return `${_.round(totalUsdVolumes / 1000000, 2)}M`
  }

  return `${_.round(totalUsdVolumes / 1000, 1)}K`
}

const formatTick = (value, txCount, volume, showUSDVolume) => {
  if (showUSDVolume) {
    return `$${formatShort.format(value)}`
  }

  return formatShort.format(value)
}

export const addOrSetObject = (obj, key, value) => {
  obj[key] ? (obj[key] += value) : (obj[key] = value)
}

const createMonthlyData = (data, isCumulativeData) => {
  const monthlyData = {}

  data.forEach((obj, idx) => {
    const date = new Date(obj.date)
    const month = formatMonth.format(date)

    if (isCumulativeData) {
      // Skip all dates that are not the last date of the month.
      const target = new Date(date.getUTCFullYear(), date.getUTCMonth() + 1, 0)

      // Do not skip the last `obj` i.e the latest date of data.
      if (date.getTime() !== target.getTime() && idx !== data.length - 1) {
        return
      }
    }

    for (const [key, value] of Object.entries(obj)) {
      if (key === 'date') {
        if (!(month in monthlyData)) {
          monthlyData[month] = { date: month }
        }
      } else {
        monthlyData?.[month] ? addOrSetObject(monthlyData[month], key, value) : (monthlyData[month] = { [key]: value })
      }
    }
  })

  return _.values(monthlyData)
}

export const OverviewChart = ({
  data,
  isUSD,
  isCumulativeData,
  showAggregated,
  height = 480,
}) => {
  let chartData = data
  if (isCumulativeData) {
    chartData = JSON.parse(JSON.stringify(data))
    for (let i = 1; i < chartData.length; i++) {
      for (let key in data[i]) {
        if (key !== 'date' && key !== '__typename') {
          chartData[i][key] += (chartData[i - 1]?.[key] ? chartData[i - 1][key] : 0)
        }

      }
    }
  }
  console.log(chartData)
  return (
    <ResponsiveContainer width={'99%'} height={height}>
      <BarChart width={0} height={480} data={chartData} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
        <XAxis hide dataKey="date" stroke="#374151" />
        <YAxis
        tick={{fontSize: "0.7rem"}}
        orientation="right"
          interval="preserveStart"
          width={20}
          stroke="#ffffff"
          tickCount={7}
          tickFormatter={(value) => isUSD ? "$" + formatTotalUsdVolumes(value) : formatTotalUsdVolumes(value)
          }
        />
        <Tooltip cursor={{fill: 'rgba(255, 255, 255, 0.1)'}} wrapperClassName="rounded-lg shadow-lg" content={isUSD ? CurrencyTooltip : NumericTooltip} />
        {showAggregated ? (
          <Bar isAnimationActive={false} dataKey="total" stackId="a" fill="#6a30b4" />
        ) : (
          <>
            <Bar isAnimationActive={false} dataKey="ethereum" stackId="a" fill="#637eea" />
            <Bar isAnimationActive={false} dataKey="avalanche" stackId="a" fill="#e74242" />
            <Bar isAnimationActive={false} dataKey="polygon" stackId="a" fill="#7b3fe4" />
            <Bar isAnimationActive={false} dataKey="bsc" stackId="a" fill="#efb90b" />
            <Bar isAnimationActive={false} dataKey="arbitrum" stackId="a" fill="#2d374b" />
            <Bar isAnimationActive={false} dataKey="fantom" stackId="a" fill="#1969ff" />
            <Bar isAnimationActive={false} dataKey="harmony" stackId="a" fill="#39cdd8" />
            <Bar isAnimationActive={false} dataKey="optimism" stackId="a" fill="#fe0621" />
            <Bar isAnimationActive={false} dataKey="moonriver" stackId="a" fill="#f2b707" />
            <Bar isAnimationActive={false} dataKey="boba" stackId="a" fill="#cbff00" />
            <Bar isAnimationActive={false} dataKey="aurora" stackId="a" fill="#78d64b" />
            <Bar isAnimationActive={false} dataKey="moonbeam" stackId="a" fill="#20223c" />
            <Bar isAnimationActive={false} dataKey="metis" stackId="a" fill="#22e5f2" />
            <Bar isAnimationActive={false} dataKey="cronos" stackId="a" fill="#1711a2" />
            <Bar isAnimationActive={false} dataKey="dfk" stackId="a" fill="#ffff83" />
            <Bar isAnimationActive={false} dataKey="klaytn" stackId="a" fill="#f9810b" />
            <Bar isAnimationActive={false} dataKey="canto" stackId="a" fill="#09fc99" />
            <Bar isAnimationActive={false} dataKey="dogechain" stackId="a" fill="#8168f7" />

          </>
        )}
      </BarChart>
    </ResponsiveContainer>
  )
}
