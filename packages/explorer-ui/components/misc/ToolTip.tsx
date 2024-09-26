import ReactDOM from 'react-dom'
import { CHAINS } from 'synapse-constants'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { ChainInfo } from '@components/misc/ChainInfo'
import { formatUSD } from '@utils/formatUSD'
import { formatDate } from '@utils/formatDate'

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
    <div className="h-full bg-synapse">
      <p className="pb-1 pl-2 mt-4 text-4xl font-medium text-white text-default">
        {platform !== 'ALL' ? platformTitles[platform] + ' ' : ''} Activity
      </p>
      <p className="pl-2 text-lg text-white ">{formatDate(label)}</p>
      <table className="min-w-full">
        <TableHeader headers={['Chain', titles[dailyStatisticType]]} />
        <tbody>
          {singleChain ? null : (
            <tr key={-1} className="w-full">
              <td className="w-[64%]">
                <p className="pl-2 text-sm text-white whitespace-nowrap">
                  All Chains
                </p>
              </td>
              <td className="w-fit ">
                <div className="self-center ml-1 mr-2">
                  <p className="px-2 text-sm text-white whitespace-nowrap">
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
                    <p className="pl-2 text-sm text-white whitespace-nowrap">
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
                  <div className="self-center ml-1 mr-2">
                    <p className="px-2 text-sm text-white whitespace-nowrap">
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
}
