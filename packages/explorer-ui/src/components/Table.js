/* eslint-disable react/jsx-key */
import { CHAIN_ID_NAMES_REVERSE, CHAIN_INFO } from 'constants/chains'
import { formatCurrency } from 'utils/formatCurrency'
import { formatNumber } from 'utils/formatNumber'

export const OverviewTable = ({ columns, data, currency }) => {
  return (
    <div className="flex flex-col">
      <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <div className="overflow-hidden border-b border-gray-900 shadow sm:rounded-lg">
            <table className="min-w-full divide-y divide-gray-900">
              <thead className="bg-gray-800">
                <tr>
                  {columns.map((column) => (
                    <th
                      scope="col"
                      className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-100 uppercase"
                    >
                      {column}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody className="bg-gray-700 divide-y divide-gray-900">
                {data.map((item, index) => (
                  <tr>
                    {/* <td className="px-6 py-4 whitespace-nowrap">
                      <div className="flex flex-shrink-0">
                        <div className="text-base text-gray-100 max-w-[16px]">{index + 1}</div>
                      </div>
                    </td> */}
                    <td className="px-6 py-4 whitespace-nowrap">
                      <div className="flex items-center">
                        <div className="flex-shrink-0 w-6 h-6">
                          <div className="relative w-6 h-6 rounded-full">
                            <img
                              src={CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[item.name]].chainLogo}
                              alt=""
                              layout="fill"
                              objectFit="contain"
                            />
                          </div>
                        </div>
                        <div className="ml-4">
                          <div className="text-base font-medium text-gray-100">
                            {CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[item.name]].chainName
                              .replace('Binance Smart Chain', 'BSC')
                              .replace('Boba Network', 'Boba')}
                          </div>
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <div className="text-base text-gray-100">
                        {currency ? formatCurrency.format(item.value) : formatNumber.format(item.value)}
                      </div>
                    </td>
                    {item?.averageTxSize && (
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-base text-gray-100">{formatCurrency.format(item.averageTxSize)}</div>
                      </td>
                    )}
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  )
}
