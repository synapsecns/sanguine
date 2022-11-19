import { FlowChart } from 'components/Charts'
import { Header } from 'components/Header'
import { Card, GradientCard, PrimaryContainer } from 'components/Sections'
import { LabelSwitch } from 'components/Switch'
import { CHAIN_ID_NAMES_REVERSE, CHAIN_INFO } from 'constants/chains'
import _ from 'lodash'
import { useState } from 'react'
import { useParams } from 'react-router'
import useSWR from 'swr'
import { fetcher } from 'utils/fetcher'
import { formatNumber } from 'utils/formatNumber'
import { formatTotalUsdVolumes } from 'utils/formatTotal'

const Page = () => {
  const { chain } = useParams()
  const [showUSDVolume, setShowUSDVolume] = useState(false)
  const [showCumulativeVolume, setShowCumulativeVolume] = useState(false)
  const [showCumulativeTxCount, setShowCumulativeTxCount] = useState(false)

  let { data: nusdData, error: nusdError } = useSWR(
    `https://analytics-api.bridgesyn.com/api/v1/analytics/pools/volume/${chain}/nusd`,
    fetcher,
  )

  let { data: nethData, error: nethError } = useSWR(
    `https://analytics-api.bridgesyn.com/api/v1/analytics/pools/volume/${chain}/neth`,
    fetcher,
  )

  let previousCoinVolumes = {}
  let previousUsdVolumes = {}
  let previousTxCounts = {}
  let totalCoinVolume = 0
  let totalUsdVolume = 0
  let totalTxCount = 0

  let nusdDataFormatted = _.orderBy(
    _.map(nusdData, (value, key) => {
      let volume_usd = 0
      let tx_count = 0
      let volume = 0

      ;['add_remove', 'swap_base', 'swap_nexus'].forEach((method) => {
        const data = value[method] || {}

        volume_usd += data?.volume_usd ?? 0
        tx_count += data?.tx_count ?? 0
        volume += data?.volume ?? 0
      })

      totalUsdVolume += volume_usd
      totalCoinVolume += volume
      totalTxCount += tx_count

      previousCoinVolumes[key] = totalCoinVolume
      previousUsdVolumes[key] = totalUsdVolume
      previousTxCounts[key] = totalTxCount

      return {
        name,
        date: key,
        txCount: showCumulativeTxCount ? previousTxCounts[key] : tx_count,
        coinVolume: showCumulativeVolume ? previousCoinVolumes[key] : volume,
        usdVolume: showCumulativeVolume ? previousUsdVolumes[key] : volume_usd,
      }
    }),
    'date',
  )

  let previousNethCoinVolumes = {}
  let previousNethUsdVolumes = {}
  let previousNethTxCounts = {}
  let totalNethCoinVolume = 0
  let totalNethUsdCoinVolume = 0
  let totalNethTxCount = 0

  let nethDataFormatted = _.orderBy(
    _.map(nethData, (value, key) => {
      let volume_usd = 0
      let tx_count = 0
      let volume = 0

      ;['add_remove', 'swap_base', 'swap_nexus'].forEach((method) => {
        const data = value[method] || {}

        volume_usd += data?.volume_usd ?? 0
        tx_count += data?.tx_count ?? 0
        volume += data?.volume ?? 0
      })

      totalNethUsdCoinVolume += volume_usd
      totalNethCoinVolume += volume
      totalNethTxCount += tx_count

      previousNethUsdVolumes[key] = totalNethUsdCoinVolume
      previousNethCoinVolumes[key] = totalNethCoinVolume
      previousNethTxCounts[key] = totalNethTxCount

      return {
        name,
        date: key,
        txCount: showCumulativeTxCount ? previousNethTxCounts[key] : tx_count,
        coinVolume: showCumulativeVolume ? previousNethCoinVolumes[key] : volume,
        usdVolume: showCumulativeVolume ? previousNethUsdVolumes[key] : volume_usd,
      }
    }),
    'date',
  )

  totalUsdVolume += totalNethUsdCoinVolume
  totalTxCount += totalNethTxCount

  return (
    <div>
      <Header />
      <PrimaryContainer>
        <div className="flex items-center">
          <h1 className="mt-8 mb-8 text-4xl font-bold text-gray-200">
            {'Pool Statistics | '}
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
            <h2 className="mb-2 text-xl font-bold text-gray-400">Trading Volume</h2>
            <p className="text-4xl font-bold text-gray-200">{`$${
              totalUsdVolume ? formatTotalUsdVolumes(totalUsdVolume) : '0M'
            }`}</p>
          </Card>
          <Card>
            <h2 className="mb-2 text-xl font-bold text-gray-400">Transaction Count</h2>
            <p className="text-4xl font-bold text-gray-200">{formatNumber.format(totalTxCount)}</p>
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
        <div className="grid gap-4 mb-8 text-gray-200 ">
          {nusdDataFormatted?.length > 1 && (
            <div>
              <div className="flex mb-4">
                <h2 className="text-2xl font-bold">Stableswap Pool</h2>
              </div>
              <div className="grid gap-4 mb-8 text-gray-200 md:grid-cols-2 ">
                <GradientCard>
                  <div className="flex justify-between">
                    <h2 className="mb-2 text-xl font-bold">Volume</h2>
                  </div>
                  <FlowChart data={nusdDataFormatted} showUSDVolume={showUSDVolume} volume height={320} />
                </GradientCard>
                <GradientCard>
                  <div className="flex justify-between">
                    <h2 className="mb-2 text-xl font-bold">Transaction Count</h2>
                    <div className="space-y-2"></div>
                  </div>
                  <FlowChart data={nusdDataFormatted} txCount height={320} />
                </GradientCard>
              </div>
            </div>
          )}
          {nethDataFormatted?.length > 1 && (
            <div>
              <div className="flex mb-4">
                <h2 className="text-2xl font-bold">ETH Pool</h2>
              </div>
              <div className="grid gap-4 mb-8 text-gray-200 md:grid-cols-2 ">
                <GradientCard>
                  <div className="flex justify-between">
                    <h2 className="mb-2 text-xl font-bold">Volume</h2>
                  </div>
                  <FlowChart data={nethDataFormatted} showUSDVolume={showUSDVolume} volume height={320} />
                </GradientCard>
                <GradientCard>
                  <div className="flex justify-between">
                    <h2 className="mb-2 text-xl font-bold">Transaction Count</h2>
                    <div className="space-y-2"></div>
                  </div>
                  <FlowChart data={nethDataFormatted} txCount height={320} />
                </GradientCard>
              </div>
            </div>
          )}
        </div>
      </PrimaryContainer>
    </div>
  )
}

export default Page
