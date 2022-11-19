/* eslint-disable react/jsx-key */
import { CHAIN_ID_NAMES_REVERSE, CHAIN_INFO } from 'constants/chains'
import { formatCurrency } from 'utils/formatCurrency'
import { formatNumber } from 'utils/formatNumber'

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

export const CurrencyTooltip = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
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
                <p className="mr-6 font-bold">Date</p>
              </div>
            </div>
            <div>
              <p className="">{label}</p>
            </div>
          </div>
          <div className="flex justify-between">
            <div className="flex justify-start">
              <div>
                <p className="mr-6 font-bold">Total</p>
              </div>
            </div>
            <div>
              <p className="">{formatCurrency.format(total)}</p>
            </div>
          </div>
        </div>
        <div className="px-4 py-2 bg-gray-700 rounded-b-lg ">
          {fills.map((fill, index) => {
            const name = names[index]
            const value = values[index]
            return (
              <div className="flex justify-between">
                <div className="flex justify-start">
                  <div>
                    <div style={{ backgroundColor: fill }} className="w-2 h-2 mt-2 mr-2 rounded-full" />
                  </div>
                  <div>
                    <p className="mr-6 font-bold">
                      {CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[name]]?.chainName
                        .replace('Binance Smart Chain', 'BSC')
                        .replace('Boba Network', 'Boba') || formatTooltipName(name)}
                    </p>
                  </div>
                </div>
                <div>
                  <p className="">{formatCurrency.format(value)}</p>
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

export const NumericTooltip = ({ active, payload, label }) => {
  if (active && payload && payload.length) {
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
                <p className="mr-6 font-bold">Date</p>
              </div>
            </div>
            <div>
              <p className="">{label}</p>
            </div>
          </div>
          <div className="flex justify-between">
            <div className="flex justify-start">
              <div>
                <p className="mr-6 font-bold">Total</p>
              </div>
            </div>
            <div>
              <p className="">{formatNumber.format(total)}</p>
            </div>
          </div>
        </div>
        <div className="px-4 py-2 bg-gray-700 rounded-b-lg ">
          {fills.map((fill, index) => {
            const name = names[index]
            const value = values[index]
            return (
              <div className="flex justify-between">
                <div className="flex justify-start">
                  <div>
                    <div style={{ backgroundColor: fill }} className="w-2 h-2 mt-2 mr-2 rounded-full" />
                  </div>
                  <div>
                    <p className="mr-6 font-bold">
                      {CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[name]]?.chainName
                        .replace('Binance Smart Chain', 'BSC')
                        .replace('Boba Network', 'Boba') || formatTooltipName(name)}
                    </p>
                  </div>
                </div>
                <div>
                  <p className="">{formatNumber.format(value)}</p>
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
