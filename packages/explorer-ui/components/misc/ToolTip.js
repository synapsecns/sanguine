/* eslint-disable react/jsx-key */
import { CHAIN_ID_NAMES_REVERSE, CHAIN_INFO_MAP } from '@constants/networks'
import ReactDOM from "react-dom";
import { formatDate } from '@utils/formatDate'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { ChainInfo } from '@components/misc/ChainInfo'
import { formatUSD } from '@utils/formatUSD'

const titles = {
  VOLUME: 'Volume',
  FEE: 'Fees',
  ADDRESSES: 'Addrs',
  TRANSACTIONS: 'TXs',
}
const platformTitles = {
  BRIDGE: 'Bridge',
  SWAP: 'Swap',
  MESSAGE_BUS: 'Message Bus',
}
const formatCurrency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
})
const formatNumber = new Intl.NumberFormat('en-US')

const formatTooltipName = (name) => {
  if (name === 'total') {
    return 'Total'
  } else if (name === 'txCount') {
    return 'Transaction Count'
  } else if (name === 'volume') {
    return 'Volume'
  } else if (name === 'usdVolume') {
    return 'USD Volume'
  } else if (name === 'coinVolume') {
    return 'Coin Volume'
  }

  return name
}

export const CurrencyTooltip = ({ names: names, values:values, label, dailyStatisticType: dailyStatisticType, isUSD: isUSD, platform: platform }) => {

  const domElement = document.getElementById("tooltip-sidebar"); //Reference to div#modals for create portal

  // if (active && payload && payload.length) {
    // payload.sort((a, b) => b.value - a.value);
    // const names = _.map(payload, 'name')
    // const values = _.map(payload, 'value')
    // const fills = _.map(payload, 'fill')
    // const total = _.sum(values)


    return (
      ReactDOM.createPortal(
        <div className='bg-synapse h-full'>
            <p
                className="text-4xl font-medium text-default
              text-white pl-2"
              >
               {platform !== "ALL" ? platformTitles[platform] + " " : ""} Activity
                {/* {formatDate(chartData[0].date)} to{' '}
                {formatDate(chartData[chartData.length - 1].date)} */}
              </p>
          <p className="text-lg font-bold text-white pl-2 pt-4 ">{titles[dailyStatisticType]} for {formatDate(label)}</p>
          <table className='min-w-full'>

            <TableHeader headers={['Chain', titles[dailyStatisticType]]} />
            <tbody>

              {names.map((u, index) => {
                const name = names[index]
                const value = values[index]
                if (value === 0) {
                  return null
                }
                return (<tr
                  key={index}
                  className="w-full opacity-[1]"
                >
                  <td className='w-[64%]'>
                    {name === "total" ?
                    <p className="pl-2 whitespace-nowrap text-sm text-white">All Chains</p> :
                      <ChainInfo
                        useExplorerLink={true}
                        chainId={CHAIN_ID_NAMES_REVERSE[name]}
                        imgClassName="w-4 h-4 ml-2"
                        textClassName="whitespace-nowrap px-2 text-sm text-white w-full"
                      />}
                  </td>
                  <td className='w-fit '>
                    <div className="ml-1 mr-2 self-center">
                      <p className='whitespace-nowrap px-2  text-sm  text-white'>{isUSD ? "$": ""}{formatUSD(value)}</p>
                    </div>
                  </td>
                </tr>)
              })}
            </tbody>
          </table>
        </div>

        , domElement)
    )
  // }

  // return <p className='text-white bg-white'>penis</p>
}

export const NumericTooltip = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
    payload.sort((a, b) => b.value - a.value);

    const names = _.map(payload, 'name')
    const values = _.map(payload, 'value')
    const fills = _.map(payload, 'fill')
    const total = _.sum(values)
    return (
      <div className="bg-gray-600 rounded-lg shadow-lg">
        <div className="px-4 py-2">
          <div className="flex justify-between">
            <div className="flex justify-start">
              <div>
                <p className="mr-6 font-bold text-white">Date</p>
              </div>
            </div>
            <div>
              <p className="text-white">{label}</p>
            </div>
          </div>
          <div className="flex justify-between">
            <div className="flex justify-start">
              <div>
                <p className="mr-6 font-bold text-white">Total</p>
              </div>
            </div>
            <div>
              <p className="text-white">{formatNumber.format(total)}</p>
            </div>
          </div>
        </div>
        <div className="px-4 py-2 bg-gray-700 rounded-b-lg ">
          {fills.map((fill, index) => {
            const name = names[index]
            const value = values[index]
            if (value === 0) {
              return null
            }
            return (
              <div className="flex justify-between">
                <div className="flex justify-start">
                  <div>
                    <div style={{ backgroundColor: fill }} className="w-2 h-2 mt-2 mr-2 rounded-full" />
                  </div>
                  <div>
                    <p className="mr-6 font-bold text-white">
                      {CHAIN_INFO_MAP[CHAIN_ID_NAMES_REVERSE[name]]?.chainName
                        .replace('Binance Smart Chain', 'BSC')
                        .replace('Boba Network', 'Boba') || formatTooltipName(name)}
                    </p>
                  </div>
                </div>
                <div>
                  <p className="text-white">{formatNumber.format(value)}</p>
                </div>
              </div>
            )
          })}
        </div>
      </div>
    )
  }

  return <div />
}
