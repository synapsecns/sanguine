// import { Stats } from './Stats'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { ChainInfo } from '@components/misc/ChainInfo'
import { OverviewChart } from '@components/ChainChart'
import { CHAIN_ID_NAMES_REVERSE } from '@constants/networks'
import { useEffect } from 'react'
import {formatDate} from '@utils/formatDate'


export function ChartFull(loading,
  titles,
  height,
  dailyDataArr,
  isCumulativeData,
  isUSD,
  showAggregated,
  monthlyData,
  currency,
  dailyStatisticType,
  setCurrentTooltipIndex) {
  // const [currentTooltipIndex, setCurrentTooltipIndex] = useState(0)
  var currentTooltipIndexVar = 0
  const handleCurrentIndex = (index) => {
    console.log("indexindexindexindex", index)
    currentTooltipIndex = index
  }

  useEffect(() => {
    setCurrentTooltipIndex(index)
  }, [currentTooltipIndexVar])

  const returnChainData = () => {
    var items = Object.keys(dailyDataArr?.[currentTooltipIndex]).map((key) => { return [key, dailyDataArr?.[currentTooltipIndex][key]] })
    items.sort((first, second) => { return second[1] - first[1] })
    var keys = items.map((e) => { return e[0] })
    return keys
  }
  return (<div className="grid grid-cols-4 gap-4">
    <div className="col-span-1 w-[100%]">
      <p className="text-lg font-bold text-white pl-2 pt-4 ">Date: {formatDate(dailyDataArr?.[currentTooltipIndex]?.date)}</p>
      <table className='min-w-full'>

        <TableHeader headers={['Chain', titles[dailyStatisticType]]} />
        {loadingRankedChains ? <tbody> {Object.values(CHAIN_ID_NAMES_REVERSE).map((i) =>
          <tr
            key={i}

          ><td className='w-[70%]'> <div className="h-3 w-full mt-4 bg-slate-700 rounded animate-pulse"></div></td><td className='w-[30%]'><div className="h-3 w-full mt-4 bg-slate-700 rounded animate-pulse"></div></td></tr>)}</tbody> :
          (<tbody>

            {currentTooltipIndex >= 0 && dailyDataArr?.[currentTooltipIndex] ? returnChainData().map((key, i) => {
              return dailyDataArr[currentTooltipIndex][key] > 0 ? (<tr
                key={i}
                className=" rounded-md w-[100%]"
              // onClick={(event) => event.target.type !== "link" && setCurrentChainID(row.chainID)}
              >
                <td className='w-[70%]'>
                  {key === "total" ? <p className="pl-2 whitespace-nowrap text-sm text-white">All Chains</p> :
                    <ChainInfo

                      chainId={CHAIN_ID_NAMES_REVERSE[key]}
                      imgClassName="w-4 h-4 ml-2"
                      textClassName="whitespace-nowrap px-2  text-sm  text-white"
                    />}
                </td>
                <td className='w-fit '>
                  <div className="ml-1 mr-2 self-center">
                    <p className='whitespace-nowrap px-2  text-sm  text-white'>{formatUSD(dailyDataArr[currentTooltipIndex][key])}</p>
                  </div>
                </td>
              </tr>) : null
            }) : console.log("NOT REAL")}
          </tbody>)}
      </table>
    </div>
    <div className="col-span-3 ">
      {/* { loadingDailyData ?  <div className={"flex justify-center align-center w-full animate-spin mt-[" + (Object.values(CHAIN_ID_NAMES_REVERSE).length * 10).toString() + "px]"}><SynapseLogoSvg /></div> : */}
      {loadingDailyData ? <div className="flex justify-center align-center w-full my-[240px]"><div className='animate-spin'><SynapseLogoSvg /></div></div> :

        <OverviewChart
          currentTooltipIndex={currentTooltipIndex}
          handleCurrentIndex={handleCurrentIndex}
          loading={loading}
          height={height}
          data={dailyDataArr}
          isCumulativeData={isCumulativeData}
          isUSD={isUSD}
          showAggregated={showAggregated}
          monthlyData={monthlyData}
          currency={currency}
        />}
    </div>
  </div>)
}
