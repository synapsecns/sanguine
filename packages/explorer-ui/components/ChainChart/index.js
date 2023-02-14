import _ from 'lodash'
import { useRef,useState,useMemo } from 'react'
import { SynapseLogoSvg } from "@components/layouts/MainLayout/SynapseLogoSvg";
import { formatDate } from '@utils/formatDate'

import { Bar, BarChart, Cell, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts'
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
  chartData,
  isUSD,
  loading,
  isCumulativeData,
  showAggregated,
  height = 480,
  setCurrentTooltipIndex,
  currentIndex,
}) => {
  if (loading){
    return <div className="flex justify-center align-center w-full my-[240px]"><div className='animate-spin'><SynapseLogoSvg /></div></div>
  }
  // const [currentIndex, setToolTipI] = useState(0)
  const timerRef = useRef(null);
  // const onMouseEnter = useMemo(() => {
  //   const throttled = _.throttle(activeTooltipeIndex => {
  //     handleCurrentIndex(activeTooltipeIndex)
  //     setToolTipI(activeTooltipeIndex)
  //   }, 100);
  //   return e => {
  //     // e.persist()
  //     return throttled(e)
  //   };
  // }, []);


  // let chartData = data
  // if (isCumulativeData) {
  //   chartData = JSON.parse(JSON.stringify(data))
  //   for (let i = 1; i < chartData.length; i++) {
  //     for (let key in data[i]) {
  //       if (key !== 'date' && key !== '__typename') {
  //         chartData[i][key] += (chartData[i - 1]?.[key] ? chartData[i - 1][key] : 0)
  //       }

  //     }
  //   }
  // }
const CustomTooltip = ({label }) => {
  return(<p className='rounded-md text-white opacity-[0.5] text-sm'>{formatDate(label)}</p>)
}
  return (
    <ResponsiveContainer width={'99%'} height={height}>
      <BarChart
      onClick={(state) => {
        if (state?.activeTooltipIndex) {
          setCurrentTooltipIndex(state.activeTooltipIndex)
      }}}
      width={0} height={480} data={chartData} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
        <XAxis hide dataKey="date" stroke="#374151" />
        <YAxis
          tick={{ fontSize: "0.7rem" }}
          orientation="right"
          interval="preserveStart"
          width={20}
          stroke="#ffffff"
          tickCount={7}
          tickFormatter={(value) => isUSD ? "$" + formatTotalUsdVolumes(value) : formatTotalUsdVolumes(value)
          }
        />
        <Tooltip   position={{
            y: height
          }}   active={true} cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }}  content={<CustomTooltip />} />

        {/* <Tooltip  cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }} content={<CustomTooltip jim ={"jom"}/>} /> */}
        {/* {loading ? null :
        <Tooltip cursor={{ fill: 'rgba(255, 255, 255, 0.1)' }} wrapperClassName="rounded-lg shadow-lg" content={isUSD ? CurrencyTooltip : NumericTooltip} />} */}
        {showAggregated ? (
          <Bar isAnimationActive={false} dataKey="total" stackId="a" fill="#6a30b4" />
        ) : (
          <>
            <Bar isAnimationActive={false} dataKey="ethereum" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#637eea"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#637eea": "rgba(99,126,234, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="avalanche" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#e74242"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#e74242": "rgba(231,66,66, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="polygon" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#7b3fe4"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#7b3fe4": "rgba(123,63,228, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="bsc" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#efb90b"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#efb90b": "rgba(239,185,11, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="arbitrum" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#2d374b"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#2d374b": "rgba(45,55,75, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="fantom" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#1969ff"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#1969ff": "rgba(99,126,234, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="harmony" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#39cdd8"}>
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#39cdd8": "rgba(57,205,216, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="optimism" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#fe0621"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#fe0621": "rgba(254,6,33, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="moonriver" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#f2b707"}>
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#f2b707": "rgba(242,183,7, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="boba" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#cbff00"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#cbff00": "rgba(203,255,0, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="aurora" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#78d64b"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#78d64b": "rgba(120,214,75, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="moonbeam" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#20223c"}>
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#20223c": "rgba(32,34,60, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="metis" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#22e5f2"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#22e5f2": "rgba(34,229,242, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="cronos" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#1711a2"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#1711a2": "rgba(23,17,162, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="dfk" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#ffff83"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#ffff83": "rgba(255,255,131, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="klaytn" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#f9810b"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#f9810b": "rgba(249,129,11, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="canto" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#09fc99"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#09fc99": "rgba(9,252,153, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>
            <Bar isAnimationActive={false} dataKey="dogechain" stackId="a" fill={loading ? 'rgba(255, 255, 255, 0.1)' : "#8168f7"} >
              {chartData.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={index === currentIndex ?  "#8168f7": "rgba(129,104,247, 0.7)"}
                  key={index}
                />
              ))}
            </Bar>

          </>
        )}
      </BarChart>
    </ResponsiveContainer>
  )
}
