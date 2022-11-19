/* eslint-disable react/jsx-key */
import { Header } from 'components/Header'
import { PrimaryContainer } from 'components/Sections'
import { StatisticsSection } from 'components/Statistics'
import { CHAIN_ID_NAMES_REVERSE, CHAIN_INFO } from 'constants/chains'
import _ from 'lodash'
import { Link } from 'react-router-dom'
import useSWR from 'swr'
import { fetcher } from 'utils/fetcher'
import { formatCurrency } from 'utils/formatCurrency'

const LinkCard = ({ href = '/pool', name, volume }) => {
  const chainColor = `hover:border-[${CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[name]].chainColor}] ?`

  return (
    <div className="shadow-lg">
      <Link to={`${href}/${name}`}>
        <div className={'p-4 bg-[#21283a] rounded-md hover:bg-gray-700 hover:text-gray-200 border border-transparent'}>
          <div className="flex">
            <div className="relative flex-shrink-0 w-6 h-6 mt-1 mr-1 rounded-full ">
              <img src={CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[name]].chainLogo} alt="" layout="fill" objectFit="contain" />
            </div>
            <div className="ml-4">
              <Link to={href}>
                <a className="text-xl font-semibold text-gray-400">
                  {CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[name]].chainName}
                </a>
              </Link>
              {<div className="mt-2 text-base font-medium text-gray-400">{formatCurrency.format(volume)}</div>}
            </div>
          </div>
        </div>
      </Link>
    </div>
  )
}

const Page = () => {
  const { data: volume } = useSWR('https://analytics-api.bridgesyn.com/api/v1/analytics/pools/volume/total', fetcher)

  const volumeTableData = _.orderBy(
    _.map(volume?.totals, (value, key) => {
      return { name: key, value: value }
    }),
    'value',
    'desc',
  )

  return (
    <div>
      <Header />
      <PrimaryContainer>
        <h1 className="mt-8 mb-8 text-4xl font-bold text-gray-200">Pool Statistics</h1>
        <StatisticsSection />
        <h1 className="mt-8 mb-8 text-4xl font-bold text-gray-200">Select Chain</h1>
        <div className="grid gap-4 md:grid-cols-2 xl:grid-cols-4">
          {volumeTableData.map((result, key) => {
            const { name, value } = result
            return <LinkCard name={name} volume={value} />
          })}
        </div>
      </PrimaryContainer>
    </div>
  )
}

export default Page
