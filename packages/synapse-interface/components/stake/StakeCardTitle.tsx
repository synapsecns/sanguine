import numeral from 'numeral'
import _ from 'lodash'
import { useEffect, useState, useMemo } from 'react'
import { LoaderIcon } from 'react-hot-toast'
import type { Token } from '@/utils/types'
import { getPoolApyData } from '@/utils/actions/getPoolApyData'
import ApyTooltip from '@/components/ApyTooltip'
import { hasAllPrices } from '@/utils/hasAllPrices'
import { useAppSelector } from '@/store/hooks'

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
  token: Token
  poolTokens: Token[]
  poolLabel: string
  lpTokenBalance: bigint
}

const StakeCardTitle = ({
  token,
  poolTokens,
  poolLabel,
  lpTokenBalance,
}: StakeCardTitleProps) => {
  const [poolApyData, setPoolApyData] = useState<any>(null)
  const { synPrices, ethPrice, avaxPrice, metisPrice } = useAppSelector(
    (state) => state.priceData
  )

  const prices = { synPrices, ethPrice, avaxPrice, metisPrice }

  useEffect(() => {
    if (hasAllPrices(prices)) {
      getPoolApyData(token.chainId, token, {
        synPrices,
        ethPrice,
        avaxPrice,
        metisPrice,
      })
        .then((res) => {
          setPoolApyData(res)
        })
        .catch((err) => {
          console.log('Could not get pool data', err)
        })
    }
  }, [token, hasAllPrices(prices), lpTokenBalance])

  const displayPoolApyData = useMemo(() => {
    if (!poolApyData) return null

    return poolApyData.fullCompoundedAPY
      ? `${numeral(poolApyData.fullCompoundedAPY / 100).format('0.0%')}`
      : `-%`
  }, [prices, poolApyData])

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
