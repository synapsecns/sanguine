import { Token } from '@/utils/types'
import _ from 'lodash'
import numeral from 'numeral'
import { memo } from 'react'
import { LoaderIcon } from 'react-hot-toast'

export const PoolCardBody = memo(
  ({
    pool,
    poolData,
    poolApyData,
  }: {
    pool: Token
    poolData: any
    poolApyData: any
  }) => {
    const format = poolData.totalLockedUSD > 1000000 ? '$0,0.0' : '$0,0'
    return (
      <div className="flex items-center justify-between px-3 pt-1 pb-2 h-[65px]">
        <div className="flex items-center space-x-3">
          <PoolTokenIcons pool={pool} />
          <div className="text-white">
            <div className="flex items-center space-x-2 font-medium text-xxl">
              <div className="">
                {poolData && numeral(poolData.totalLockedUSD).format(format)}
              </div>
              <span className="text-base text-[#BFBCC2]">
                {pool.priceUnits}
              </span>
            </div>
          </div>
        </div>

        <div className="">
          <ApyDisplay pool={pool} poolApyData={poolApyData} />
        </div>
      </div>
    )
  }
)

const PoolTokenIcons = memo(({ pool }: { pool: Token }) => {
  if (pool.poolTokens.length === 3) {
    return (
      <div className="flex flex-col items-center">
        {pool.poolTokens.length === 3 && (
          <div>
            <img
              alt={pool.poolTokens[0].symbol}
              className="w-[1.5rem] h-[1.5rem] rounded-full"
              src={pool.poolTokens[0].icon.src}
            />
          </div>
        )}

        <div className="flex justify-center">
          {pool.poolTokens.slice(1).map((token, i) => (
            <img
              alt={token.symbol}
              className="w-[1.5rem] h-[1.5rem] rounded-full"
              src={token.icon.src}
              key={i}
            />
          ))}
        </div>
      </div>
    )
  } else {
    return (
      <div className="flex flex-wrap max-w-[40px]">
        {pool.poolTokens.map((token, i) => (
          <div className="flex items-center justify-between" key={i}>
            <img
              alt={token.symbol}
              className="w-5 h-5 rounded-full"
              src={token.icon.src}
              key={i}
            />
          </div>
        ))}
      </div>
    )
  }
})

const ApyDisplay = ({ pool, poolApyData }) => {
  const apy =
    poolApyData.fullCompoundedAPY > 10000
      ? 10000
      : poolApyData.fullCompoundedAPY

  if (!pool.incentivized) {
    return ''
  }

  if (
    isNaN(Number(poolApyData.fullCompoundedAPYStr)) ||
    poolApyData.fullCompoundedAPYStr === '0.00'
  ) {
    return <LoaderIcon />
  }

  return (
    <div>
      <div className="font-medium text-white text-xxl">
        {apy
          ? `${numeral(apy / 100).format('0,0%')}${
              poolApyData.fullCompoundedAPY > 10000 ? '+' : ''
            }`
          : `-%`}
      </div>
      <div className=" text-[#BFBCC2] text-right">APY</div>
    </div>
  )
}
