import numeral from 'numeral'
import _ from 'lodash'
import { useEffect, useState, useMemo } from 'react'
import { LoaderIcon } from 'react-hot-toast'
import { Token } from '@/utils/types'
import { getPoolApyData } from '@/utils/actions/getPoolApyData'
import ApyTooltip from '@/components/ApyTooltip'

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
  prices: any
  lpTokenBalance: bigint
}

const StakeCardTitle = ({
  address,
  connectedChainId,
  token,
  poolTokens,
  poolLabel,
  prices,
  lpTokenBalance,
}: StakeCardTitleProps) => {
  const [poolApyData, setPoolApyData] = useState<any>(null)
  const [baseApyData, setBaseApyData] = useState<any>(null)

  useEffect(() => {
    if (connectedChainId && prices) {
      getPoolApyData(connectedChainId, token, prices)
        .then((res) => {
          setPoolApyData(res)
        })
        .catch((err) => {
          console.log('Could not get pool data', err)
        })
    }
  }, [connectedChainId, prices, lpTokenBalance])

  useEffect(() => {
    null
    setBaseApyData(null)
  }, [connectedChainId])

  const displayPoolApyData = useMemo(() => {
    if (!poolApyData) return null

    return poolApyData.fullCompoundedAPY
      ? `${numeral(poolApyData.fullCompoundedAPY / 100).format('0.0%')}`
      : `-%`
  }, [connectedChainId, prices, poolApyData])

  return (
    <div className="flex items-center justify-between mb-5">
      <div className="inline-flex items-center mt-2">
        <StakingPoolTokens poolTokens={poolTokens} />
        <h3 className="mr-2 text-xl font-medium text-white">{poolLabel}</h3>
      </div>

      <div className="text-lg font-normal text-white text-opacity-70">
        <div>
          {displayPoolApyData ? (
            <span className="text-white ">{displayPoolApyData}</span>
          ) : (
            <LoaderIcon />
          )}
        </div>
        <div className="flex">
          <div className="text-sm">APY</div>
          <ApyTooltip
            apyData={poolApyData}
            className="flex items-center ml-1"
          />
        </div>
      </div>
    </div>
  )
}

export default StakeCardTitle
