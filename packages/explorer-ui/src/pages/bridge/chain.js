import { FlowChart } from 'components/Charts'
import { Header } from 'components/Header'
import { Card, GradientCard, PrimaryContainer } from 'components/Sections'
import { LabelSwitch } from 'components/Switch'
import { CHAIN_ID_NAMES_REVERSE, CHAIN_INFO } from 'constants/chains'
import { TOKEN_INFO } from 'constants/tokens'
import _ from 'lodash'
import { useState } from 'react'
import { useParams } from 'react-router-dom'
import useSWR from 'swr'
import { fetcher } from 'utils/fetcher'
import { formatNumber } from 'utils/formatNumber'
import { formatTotalUsdVolumes } from 'utils/formatTotal'

const Page = () => {
  const { chain } = useParams()
  const [showUSDVolume, setShowUSDVolume] = useState(false)
  const [showCumulativeVolume, setShowCumulativeVolume] = useState(false)
  const [showCumulativeTxCount, setShowCumulativeTxCount] = useState(false)

  let { data: inData, error: inError } = useSWR(
    `https://analytics-api.bridgesyn.com/api/v1/analytics/volume/${chain}/in`,
    fetcher,
  )
  let { data: outData, error: outError } = useSWR(
    `https://analytics-api.bridgesyn.com/api/v1/analytics/volume/${chain}/out`,
    fetcher,
  )

  const newData = _.orderBy(
    _.map(inData?.data, (value, key) => {
      const name = TOKEN_INFO[CHAIN_ID_NAMES_REVERSE[chain] || 1][key]?.name
      let order
      if (name === 'Wrapped ETH') {
        order = -3
      } else if (name === 'Synapse nUSD') {
        order = -2
      } else if (name === 'Synapse') {
        order = -1
      } else {
        order = 0
      }

      return { ...value, name, key, order }
    }),
    'order',
  )

  let totalUsdVolumes = 0
  let totalTxCounts = 0
  let inDataFormatted = _.map(newData, (asset, key) => {
    let previousCoinVolumes = {}
    let previousUsdVolumes = {}
    let previousTxCounts = {}
    let totalCoinVolume = 0
    let totalUsdVolume = 0
    let totalTxCount = 0

    let data = _.orderBy(
      _.map(asset?.data, (value, key) => {
        let { price_usd, tx_count, volume } = value
        totalCoinVolume += volume
        previousCoinVolumes[key] = totalCoinVolume
        totalUsdVolume += price_usd
        previousUsdVolumes[key] = totalUsdVolume
        totalTxCount += tx_count
        previousTxCounts[key] = totalTxCount

        return {
          name: asset.key,
          date: key,
          txCount: showCumulativeTxCount ? previousTxCounts[key] : tx_count,
          coinVolume: showCumulativeVolume ? previousCoinVolumes[key] : volume,
          usdVolume: showCumulativeVolume ? previousUsdVolumes[key] : price_usd,
        }
      }),
      'date',
    )

    totalUsdVolumes += totalUsdVolume
    totalTxCounts += totalTxCount

    return data
  })

  return (
    <div>
      <Header />
      <PrimaryContainer>
        <div className="flex items-center">
          <h1 className="mt-8 mb-8 text-4xl font-bold text-gray-200">
            {'Bridge Statistics | '}
            {CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[chain]]?.chainName || ''}
          </h1>
          {chain && CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[chain]]?.chainLogo && (
            <div className="relative w-10 h-10 ml-4 rounded-full">
              <img src={CHAIN_INFO[CHAIN_ID_NAMES_REVERSE[chain]].chainLogo} alt="" layout="fill" objectFit="contain" />
            </div>
          )}
        </div>
        <div className="grid gap-4 mb-8 md:grid-cols-2 xl:grid-cols-4">
          <Card>
            <h2 className="mb-2 text-xl font-bold text-gray-400">Bridge Volume</h2>
            <p className="text-4xl font-bold text-gray-200">{`$${
              totalUsdVolumes ? formatTotalUsdVolumes(totalUsdVolumes) : '0M'
            }`} </p>
          </Card>
          <Card>
            <h2 className="mb-2 text-xl font-bold text-gray-400">Transaction Count</h2>
            <p className="text-4xl font-bold text-gray-200">{formatNumber.format(totalTxCounts)}</p>
          </Card>
          <div className="md:col-span-2">
            <Card>
              <div className="flex justify-between">
                <h3 className="mb-2 text-xl font-bold text-gray-400">Options</h3>
                <div className="grid gap-2">
                  <LabelSwitch enabled={showUSDVolume} setEnabled={setShowUSDVolume} text="Show volume in USD" />
                  <LabelSwitch
                    enabled={showCumulativeVolume}
                    setEnabled={setShowCumulativeVolume}
                    text="Show cumulative volume"
                  />
                  <LabelSwitch
                    enabled={showCumulativeTxCount}
                    setEnabled={setShowCumulativeTxCount}
                    text="Show cumulative count"
                  />
                </div>
              </div>
            </Card>
          </div>
        </div>
        <div className="gap-4 mb-8 text-gray-200">
          {_.map(inDataFormatted, (value, key) => {
            return (
              <div>
                <div className="flex mb-4">
                  <h2 className="text-2xl font-bold">
                    {TOKEN_INFO[CHAIN_ID_NAMES_REVERSE[chain] || 1][value[0]?.name]?.name || value[0]?.name}
                  </h2>
                  <div className="relative flex items-start w-8 h-8 ml-4 rounded-full">
                    {TOKEN_INFO[CHAIN_ID_NAMES_REVERSE[chain] || 1][value[0]?.name]?.icon && (
                      <img
                        src={TOKEN_INFO[CHAIN_ID_NAMES_REVERSE[chain] || 1][value[0]?.name].icon}
                        alt=""
                        layout="fill"
                        objectFit="contain"
                      />
                    )}
                  </div>
                </div>
                <div className="grid gap-4 mb-8 text-gray-200 md:grid-cols-2 ">
                  <GradientCard>
                    <div className="flex justify-between">
                      <h2 className="mb-2 text-xl font-bold">Volume</h2>
                    </div>
                    <FlowChart data={value} showUSDVolume={showUSDVolume} volume height={320} />
                  </GradientCard>
                  <GradientCard>
                    <div className="flex justify-between">
                      <h2 className="mb-2 text-xl font-bold">Transaction Count</h2>
                      <div className="space-y-2"></div>
                    </div>
                    <FlowChart data={value} txCount height={320} />
                  </GradientCard>
                </div>
              </div>
            )
          })}
        </div>
      </PrimaryContainer>
    </div>
  )
}

export default Page
