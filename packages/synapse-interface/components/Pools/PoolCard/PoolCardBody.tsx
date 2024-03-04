
import _ from 'lodash'
import numeral from 'numeral'
import { memo } from 'react'

import type { Token } from '@/utils/types'
import { LoadingHelix } from '@tw/LoadingHelix'

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
    const format = (poolData?.totalLockedUSD > 1000000) ? '$0,0' : '$0,0.00'

    return (
      <div className="flex items-center justify-between px-3 pt-1 pb-2 h-[65px]">
        <div className="flex items-center space-x-3">
          <PoolTokenIcons pool={pool} />
          {
            poolData
              ?
                <div className="text-white">
                  <div className="flex items-center space-x-2 font-medium text-2xl leading-7">
                    <div>
                      {numeral(poolData.totalLockedUSD).format(format)}
                    </div>
                    <span className="text-base text-[#BFBCC2]">
                      {pool.priceUnits}
                    </span>
                  </div>
                </div>
              :
                <LoadingHelix className="w-20 h-20" />
          }

        </div>
        <ApyDisplay pool={pool} poolApyData={poolApyData} />
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
  const { fullCompoundedAPY } = poolApyData ?? {}
  const apy =
    fullCompoundedAPY > 10000
      ? 10000
      : fullCompoundedAPY

  if (!pool.incentivized) {
    return ''
  }

  return (
    <div>
      <div className="font-medium text-white text-xxl">
        {apy
          ? `${numeral(apy / 100).format('0,0%')}${
              fullCompoundedAPY > 10000 ? '+' : ''
            }`
          : `-%`}
      </div>
      <div className="text-[#BFBCC2] text-right">APY</div>
    </div>
  )
}
