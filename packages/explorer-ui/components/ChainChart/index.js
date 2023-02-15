import _ from 'lodash'
import { useRef, useState, useMemo, useEffect } from 'react'
import { SynapseLogoSvg } from "@components/layouts/MainLayout/SynapseLogoSvg";
import { formatDate } from '@utils/formatDate'
import { formatUSD } from '@utils/formatUSD'
import ReactDOM from "react-dom";

import { Bar, BarChart, Cell, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
import { CurrencyTooltip, NumericTooltip } from '@components/misc/ToolTip'
const formatShort = new Intl.NumberFormat('en-US', { notation: 'compact', maximumFractionDigits: 1 })
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
  loading,
  showAggregated,
  height = 480,
  dailyStatisticType,
  platform
}) => {
  const initialData = (getNames) => {
    if (chartData.length === 0) return []

    let payload = chartData[chartData.length - 1]
    console.log("SDSD", payload)
    // Create items array
    var items = Object.keys(payload).map((key) => {
      if (payload[key] > 0) {
        console.log("key", key, payload[key])

        return [key, payload[key]];
      } else {
        return [key, 0];
      }
    });

    // Sort the array based on the second element
    items.sort(function(first, second) {
      return second[1] - first[1];
    });

    if (getNames) {
    console.log("items", items)
      let names= items.map((items)=>items[0])
      console.log("names", names)

      return names
    }
    let values= items.map((items)=>items[1])
    console.log("values", values)

    return values

  }
  const [toolTipNames, setToolTipNames] = useState(initialData(true))
  const [toolTipValues, setToolTipValues] = useState(initialData(false))
  const [toolTipLabel, setToolTipLabel] = useState("")

  useEffect(() => {
    setToolTipNames(initialData(true))
    setToolTipValues(initialData(false))
    console.log("DATSTTSTS",chartData[chartData.length - 1]?.date)
    setToolTipLabel(chartData[chartData.length - 1]?.date)
  }, [chartData])
  if (loading) {
    return <div className="flex justify-center align-center w-full my-[240px]"><div className='animate-spin'><SynapseLogoSvg /></div></div>
  }

  const getToolTip = ({ active, payload, label,isUSD: isUSD}) => {
    const domElement = document.getElementById("tooltip-subtitle"); //Reference to div#modals for create portal
console.log("domElement", domElement)
    payload.sort((a, b) => b.value - a.value);
    const names = _.map(payload, 'name')
    const values = _.map(payload, 'value')
    if (active) {
      setToolTipNames(names)
      setToolTipValues(values)

      return <CurrencyTooltip label={label} names={names} values={values} isUSD={isUSD} dailyStatisticType={dailyStatisticType}  platform={platform} />

    }
    // ReactDOM.createPortal(
    //   <p className="pl-2 text-md font-medium text-default mt-2 text-white">{formatDate(toolTipLabel)}sadhjksahdjsahdkjsadjk </p>, domElement)
    return <CurrencyTooltip label={toolTipLabel} names={toolTipNames} values={toolTipValues} isUSD={isUSD} dailyStatisticType={dailyStatisticType} platform={platform} />

  }

  return (
    <>
      <div id="tooltip-sidebar" />
      <ResponsiveContainer width={'99%'} height={height}>
        <BarChart
          width={0} height={480} data={chartData} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
          <XAxis hide dataKey="date" stroke="#374151" />
          <YAxis
            tick={{ fontSize: "0.7rem" }}
            orientation="right"
            interval="preserveStart"
            width={20}
            stroke="#ffffff"
            tickCount={7}
            tickFormatter={(value) => isUSD ? "$" + formatUSD(value) : formatUSD(value)
            }
          />
          {/* <Tooltip   position={{
            y: height
          }}   active={true} cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }}  content={<CustomTooltip />} /> */}

          {/* <Tooltip  cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }} content={<CustomTooltip jim ={"jom"}/>} /> */}
          {loading ? null :
            <Tooltip wrapperStyle={{ visibility: "visible" }} cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }} wrapperClassName="rounded-lg shadow-lg" active={true} content={getToolTip} />}
          {showAggregated ? (
            <Bar isAnimationActive={false} dataKey="total" stackId="a" fill="#6a30b4" />
          ) : (
            <>
              <Bar isAnimationActive={false} dataKey="ethereum" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#637eea"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="avalanche" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#e74242"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="polygon" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#7b3fe4"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="bsc" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#efb90b"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="arbitrum" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#2d374b"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="fantom" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#1969ff"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="harmony" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#39cdd8"}>

              </Bar>
              <Bar isAnimationActive={false} dataKey="optimism" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#fe0621"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="moonriver" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#f2b707"}>

              </Bar>
              <Bar isAnimationActive={false} dataKey="boba" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#cbff00"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="aurora" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#78d64b"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="moonbeam" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#20223c"}>

              </Bar>
              <Bar isAnimationActive={false} dataKey="metis" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#22e5f2"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="cronos" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#1711a2"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="dfk" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#ffff83"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="klaytn" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#f9810b"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="canto" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#09fc99"} >

              </Bar>
              <Bar isAnimationActive={false} dataKey="dogechain" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#8168f7"} >

              </Bar>

            </>
          )}
        </BarChart>
      </ResponsiveContainer></>
  )
}
