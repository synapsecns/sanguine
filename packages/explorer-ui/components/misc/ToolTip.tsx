/* eslint-disable react/jsx-key */
import { CHAINS } from 'synapse-constants/dist'
import ReactDOM from 'react-dom'
import { formatDate } from '@utils/formatDate'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { ChainInfo } from '@components/misc/ChainInfo'
import { formatUSD } from '@utils/formatUSD'

const CHAIN_ID_NAMES_REVERSE = CHAINS.CHAIN_ID_NAMES_REVERSE

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

export const CurrencyTooltip = ({
  names: names,
  values: values,
  label,
  dailyStatisticType: dailyStatisticType,
  isUSD: isUSD,
  singleChain,
  platform: platform,
  noTooltipLink: noTooltipLink,
}) => {
  const domElement = document.getElementById('tooltip-sidebar')
  return ReactDOM.createPortal(
    <div className="bg-synapse h-full">
      <p
        className="text-4xl font-medium text-default
              text-white pl-2 mt-4 pb-1"
      >
        {platform !== 'ALL' ? platformTitles[platform] + ' ' : ''} Activity
        {/* {formatDate(chartData[0].date)} to{' '}
                {formatDate(chartData[chartData.length - 1].date)} */}
      </p>
      <p className="text-lg text-white pl-2 ">{formatDate(label)}</p>
      <table className="min-w-full">
        <TableHeader headers={['Chain', titles[dailyStatisticType]]} />
        <tbody>
          {singleChain ? null : (
            <tr key={-1} className="w-full">
              <td className="w-[64%]">
                <p className="pl-2 whitespace-nowrap text-sm text-white">
                  All Chains
                </p>
              </td>
              <td className="w-fit ">
                <div className="ml-1 mr-2 self-center">
                  <p className="whitespace-nowrap px-2  text-sm  text-white">
                    {isUSD ? '$' : ''}
                    {formatUSD(values.reduce((a, b) => a + b, 0))}
                  </p>
                </div>
              </td>
            </tr>
          )}
          {names.map((u, index) => {
            const name = names[index]
            const value = values[index]
            if (name === 'total') {
              return null
            }
            return (
              <tr key={index} className="w-full opacity-[1]">
                <td className="w-[64%]">
                  {name === 'total' ? (
                    <p className="pl-2 whitespace-nowrap text-sm text-white">
                      All Chains
                    </p>
                  ) : (
                    <ChainInfo
                      noLink={noTooltipLink}
                      useExplorerLink={true}
                      chainId={CHAIN_ID_NAMES_REVERSE[name]}
                      imgClassName="w-4 h-4 ml-2"
                      textClassName="whitespace-nowrap px-2 text-sm text-white w-full"
                    />
                  )}
                </td>
                <td className="w-fit ">
                  <div className="ml-1 mr-2 self-center">
                    <p className="whitespace-nowrap px-2  text-sm  text-white">
                      {isUSD ? '$' : ''}
                      {formatUSD(value)}
                    </p>
                  </div>
                </td>
              </tr>
            )
          })}
        </tbody>
      </table>
    </div>,

    domElement
  )
  // }

  // return <p className='text-white bg-white'>penis</p>
}
