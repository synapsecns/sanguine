import { useEffect, useState } from 'react'
import { Token } from '@/utils/types'
import { getPoolApyData } from '@/utils/actions/getPoolApyData'
import {
  getSynPrices,
  getEthPrice,
  getAvaxPrice,
} from '@/utils/actions/getPrices'
import ApyTooltip from '@/components/ApyTooltip'
import _ from 'lodash'

const StakingPoolTokens = ({ poolTokens }: { poolTokens: Token[] }) => {
  if (poolTokens)
    return (
      <div className="items-center hidden mr-4 md:flex lg:flex">
        {poolTokens.map((token: Token) => (
          <img
            key={token.symbol}
            className="relative inline-block w-6 -mr-1 text-white shadow-solid"
            src={token.icon.src}
          />
        ))}
      </div>
    )
}

interface StakeCardTitleProps {
  address: string
  connectedChainId: number
  token: Token
  poolTokens: Token[]
  poolLabel: string
}

const StakeCardTitle = ({
  address,
  connectedChainId,
  token,
  poolTokens,
  poolLabel,
}: StakeCardTitleProps) => {
  const [synPrices, setSynPrices] = useState(undefined)
  const [ethPrice, setEthPrice] = useState(undefined)
  const [avaxPrice, setAvaxPrice] = useState(undefined)
  const [poolApyData, setPoolApyData] = useState<any>()
  const [baseApyData, setBaseApyData] = useState<any>(null)

  // Prices to reduce number of calls
  useEffect(() => {
    getSynPrices()
      .then((res) => {
        setSynPrices(res)
      })
      .catch((err) => console.log('Could not get syn prices', err))
    getEthPrice()
      .then((res) => {
        setEthPrice(res)
      })
      .catch((err) => console.log('Could not get eth prices', err))
    getAvaxPrice()
      .then((res) => {
        setAvaxPrice(res)
      })
      .catch((err) => console.log('Could not get avax prices', err))
  }, [])

  useEffect(() => {
    if (connectedChainId && address && synPrices && ethPrice && avaxPrice) {
      getPoolApyData(connectedChainId, token, {
        synPrices,
        ethPrice,
        avaxPrice,
      })
        .then((res) => {
          setPoolApyData(res)
        })
        .catch((err) => {
          console.log('Could not get pool data', err)
        })
    }
  }, [])

  let fullyCompoundedApyLabel
  if (poolApyData && _.isFinite(poolApyData.fullCompoundedAPY)) {
    fullyCompoundedApyLabel = _.round(
      poolApyData.fullCompoundedAPY + (baseApyData?.yearlyCompoundedApy ?? 0),
      2
    ).toFixed(2)
  } else {
    fullyCompoundedApyLabel = <i className="opacity-50"> - </i>
  }

  return (
    <div className="px-2 mb-5">
      <div className="inline-flex items-center mt-2">
        <StakingPoolTokens poolTokens={poolTokens} />
        <h3 className="mr-2 text-xl font-medium text-white">{poolLabel}</h3>
      </div>

      <div className="text-lg font-normal text-white text-opacity-70">
        <span className="text-green-400">{fullyCompoundedApyLabel}% </span>
        APY
        <ApyTooltip
          apyData={poolApyData}
          // baseApyData={baseApyData ??}
          className="ml-1"
        />
      </div>
    </div>
  )
}

export default StakeCardTitle
