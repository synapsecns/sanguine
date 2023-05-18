import _ from 'lodash'
import { useRef, useState, useMemo, useEffect } from 'react'
import { SynapseLogoSvg } from '@components/layouts/MainLayout/SynapseLogoSvg'
import { formatDate } from '@utils/formatDate'
import { formatUSD } from '@utils/formatUSD'
import ReactDOM from 'react-dom'
import {
  Bar,
  BarChart,
  Cell,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts'
import { CurrencyTooltip } from '@components/misc/ToolTip'
const formatShort = new Intl.NumberFormat('en-US', {
  notation: 'compact',
  maximumFractionDigits: 1,
})
const formatMonth = new Intl.DateTimeFormat('en-US', {
  year: 'numeric',
  month: 'long',
})

export const addOrSetObject = (obj, key, value) => {
  obj[key] ? (obj[key] += value) : (obj[key] = value)
}

export const OverviewChart = ({
  chartData,
  isUSD,
  loading = false,
  showAggregated,
  height = 480,
  dailyStatisticType,
  platform,
  singleChain = false,
  noTooltipLink = false,
}) => {
  if (loading) {
    return (
      <div className="flex justify-center align-center w-full my-[240px]">
        <div className="w-[39px] animate-spin">
          <SynapseLogoSvg />
        </div>
      </div>
    )
  }
  const initialData = (getNames) => {
    if (chartData.length === 0) {
      return []
    }
    const payload = chartData[chartData.length - 1]
    const items = Object.keys(payload).map((key) => {
      if (payload[key] > 0 && key !== 'total') {
        return [key, payload[key]]
      } else {
        return [key, 0]
      }
    })
    if (items.length === 0) {
      return []
    }

    // Sort the array based on the second element
    items.sort(function (first, second) {
      return second[1] - first[1]
    })

    if (getNames) {
      const names = items.map((items) => items[0])
      return names
    }
    const values = items.map((items) => items[1])
    return values
  }

  const toolTipNamesRef = useRef()
  const toolTipValuesRef = useRef()
  const toolTipLabelRef = useRef()

  useEffect(() => {
    toolTipNamesRef.current = initialData(true)
    toolTipValuesRef.current = initialData(false)
    toolTipLabelRef.current = chartData[chartData.length - 1]?.date
  }, [chartData])

  const getToolTip = ({ active, payload, label, isUSD: isUSD }) => {
    payload?.sort((a, b) => b.value - a.value)
    const names = _.map(payload, 'name')
    const values = _.map(payload, 'value')
    if (active) {
      if (toolTipNamesRef.current !== names) {
        toolTipNamesRef.current = names
      }
      if (toolTipValuesRef.current !== values) {
        toolTipValuesRef.current = values
      }
      if (toolTipLabelRef.current !== label) {
        toolTipLabelRef.current = label
      }
      return (
        <CurrencyTooltip
          label={label}
          names={names}
          values={values}
          isUSD={isUSD}
          dailyStatisticType={dailyStatisticType}
          platform={platform}
          singleChain={singleChain}
          noTooltipLink={noTooltipLink}
        />
      )
    }

    return (
      <CurrencyTooltip
        label={toolTipLabelRef.current}
        names={toolTipNamesRef.current}
        values={toolTipValuesRef.current}
        isUSD={isUSD}
        dailyStatisticType={dailyStatisticType}
        platform={platform}
        singleChain={singleChain}
        noTooltipLink={noTooltipLink}
      />
    )
  }

  return (
    <>
      <div id="tooltip-sidebar" />
      <ResponsiveContainer width={'99%'} height={height}>
        <BarChart
          width={0}
          height={480}
          data={chartData}
          margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
        >
          <XAxis hide dataKey="date" stroke="#374151" />
          <YAxis
            tick={{ fontSize: '0.7rem' }}
            orientation="right"
            interval="preserveStart"
            width={20}
            stroke="#ffffff"
            tickCount={7}
            tickFormatter={(value) =>
              isUSD ? '$' + formatUSD(value) : formatUSD(value)
            }
          />
          {/* <Tooltip   position={{
            y: height
          }}   active={true} cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }}  content={<CustomTooltip />} /> */}

          {/* <Tooltip  cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }} content={<CustomTooltip jim ={"jom"}/>} /> */}
          {loading ? null : (
            <Tooltip
              wrapperStyle={{ visibility: 'visible' }}
              cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }}
              wrapperClassName="rounded-lg shadow-lg"
              active={true}
              content={getToolTip}
            />
          )}
          {showAggregated ? (
            <Bar
              isAnimationActive={false}
              dataKey="total"
              stackId="a"
              fill="#6a30b4"
            />
          ) : (
            <>
              <Bar
                isAnimationActive={false}
                dataKey="ethereum"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#637eea'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="avalanche"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#e74242'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="polygon"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#7b3fe4'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="bsc"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#efb90b'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="arbitrum"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#2d374b'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="fantom"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#1969ff'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="harmony"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#39cdd8'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="optimism"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#fe0621'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="moonriver"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#f2b707'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="boba"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#cbff00'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="aurora"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#78d64b'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="moonbeam"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#20223c'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="metis"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#22e5f2'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="cronos"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#1711a2'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="dfk"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#ffff83'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="klaytn"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#f9810b'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="canto"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#09fc99'}
              ></Bar>
              <Bar
                isAnimationActive={false}
                dataKey="dogechain"
                stackId="a"
                fill={loading ? 'rgba(255, 255, 255, 0.1)' : '#8168f7'}
              ></Bar>
            </>
          )}
        </BarChart>
      </ResponsiveContainer>
    </>
  )
}
