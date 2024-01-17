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
      <div className="flex items-center gap-3 px-3 mb-2">
        <PoolTokenIcons pool={pool} />
        <div className="flex flex-grow">
          <span className="font-medium text-xxl">
            {poolData && numeral(poolData.totalLockedUSD).format(format)}&nbsp;
          </span>
          <span className="text-base text-zinc-400 mt-1">
            {pool.priceUnits}
          </span>
        </div>
        <ApyDisplay pool={pool} poolApyData={poolApyData} />
      </div>
    )
  }
)

const PoolTokenIcons = memo(({ pool }: { pool: Token }) => {
  return (
    <div className="grid grid-cols-2 h-10 place-content-center">
      {pool.poolTokens.map((token, i) => (
        <img
          alt={token.symbol}
          className={`
            w-5 h-5 justify-self-center
            ${pool.poolTokens.length === 3 && 'first:col-span-2'}
          `}
          src={token.icon.src}
          key={i}
        />
      ))}
    </div>
  )
})

const ApyDisplay = ({ pool, poolApyData }) => {
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
    <div className="text-right">
      <div className="font-medium text-xxl">
        {numeral(poolApyData.fullCompoundedAPY / 100).format('0.0%')}
      </div>
      <div className="text-zinc-400">APY</div>
    </div>
  )
}
