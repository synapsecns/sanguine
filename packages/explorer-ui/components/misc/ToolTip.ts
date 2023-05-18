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
  MESSAGE_BUS: 'Msg Bus',
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

export const CurrencyTooltip = ({ names: names, values: values, label, dailyStatisticType: dailyStatisticType, isUSD: isUSD, singleChain, platform: platform, noTooltipLink:noTooltipLink }) => {

  const domElement = document.getElementById("tooltip-sidebar");
  return (
    ReactDOM.createPortal(
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className='bg-synapse h-full'>
        // @ts-expect-error TS(2304): Cannot find name 'p'.
        <p
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="text-4xl font-medium text-default
              // @ts-expect-error TS(2304): Cannot find name 'text'.
              text-white pl-2 mt-4 pb-1"
        >
          {platform !== "ALL" ? platformTitles[platform] + " " : ""} Activity
          {/* {formatDate(chartData[0].date)} to{' '}
                {formatDate(chartData[chartData.length - 1].date)} */}
        </p>
        // @ts-expect-error TS(2304): Cannot find name 'p'.
        <p className="text-lg text-white pl-2 ">{formatDate(label)}</p>
        // @ts-expect-error TS(2304): Cannot find name 'table'.
        <table className='min-w-full'>

          // @ts-expect-error TS(2749): 'TableHeader' refers to a value, but is being used... Remove this comment to see the full error message
          <TableHeader headers={['Chain', titles[dailyStatisticType]]} />
          // @ts-expect-error TS(2304): Cannot find name 'tbody'.
          <tbody>
            // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
            {singleChain ? null : <tr
              // @ts-expect-error TS(2304): Cannot find name 'key'.
              key={-1}
              // @ts-expect-error TS(2304): Cannot find name 'className'.
              className="w-full"
            >
              // @ts-expect-error TS(2304): Cannot find name 'td'.
              <td className='w-[64%]'>
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="pl-2 whitespace-nowrap text-sm text-white">All Chains</p>
              </td>
              // @ts-expect-error TS(2304): Cannot find name 'td'.
              <td className='w-fit '>
                // @ts-expect-error TS(2304): Cannot find name 'div'.
                <div className="ml-1 mr-2 self-center">
                  // @ts-expect-error TS(2304): Cannot find name 'p'.
                  <p className='whitespace-nowrap px-2  text-sm  text-white'>{isUSD ? "$" : ""}{formatUSD(values.reduce((a, b) => a + b, 0))}</p>
                </div>
              </td>
            </tr>}
            // @ts-expect-error TS(2304): Cannot find name 'names'.
            {names.map((u, index) => {
              // @ts-expect-error TS(2304): Cannot find name 'names'.
              const name = names[index]
              // @ts-expect-error TS(2304): Cannot find name 'values'.
              const value = values[index]
              if (value === 0 || name === "total") {
                return null
              }
              return (
                // @ts-expect-error TS(2304): Cannot find name 'tr'.
                <tr
                  // @ts-expect-error TS(2304): Cannot find name 'key'.
                  key={index}
                  // @ts-expect-error TS(2304): Cannot find name 'className'.
                  className="w-full opacity-[1]"
                >
                  // @ts-expect-error TS(2304): Cannot find name 'td'.
                  <td className='w-[64%]'>
                    {name === "total" ?
                      // @ts-expect-error TS(2304): Cannot find name 'p'.
                      <p className="pl-2 whitespace-nowrap text-sm text-white">All Chains</p> :
                      <ChainInfo
                        // @ts-expect-error TS(2304): Cannot find name 'noLink'.
                        noLink={noTooltipLink}
                        // @ts-expect-error TS(2304): Cannot find name 'useExplorerLink'.
                        useExplorerLink={true}
                        // @ts-expect-error TS(2304): Cannot find name 'chainId'.
                        chainId={CHAIN_ID_NAMES_REVERSE[name]}
                        // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
                        imgClassName="w-4 h-4 ml-2"
                        // @ts-expect-error TS(2304): Cannot find name 'textClassName'.
                        textClassName="whitespace-nowrap px-2 text-sm text-white w-full"
                      />}
                  </td>
                  // @ts-expect-error TS(2304): Cannot find name 'td'.
                  <td className='w-fit '>
                    // @ts-expect-error TS(2304): Cannot find name 'div'.
                    <div className="ml-1 mr-2 self-center">
                      // @ts-expect-error TS(2304): Cannot find name 'p'.
                      <p className='whitespace-nowrap px-2  text-sm  text-white'>{isUSD ? "$" : ""}{formatUSD(value)}</p>
                    </div>
                  </td>
                </tr>)
            })}
          // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
          </tbody>
        </table>
      </div>

      // @ts-expect-error TS(2304): Cannot find name 'domElement'.
      , domElement)
  )
  // }

  // return <p className='text-white bg-white'>penis</p>
}
