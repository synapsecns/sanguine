import _ from 'lodash'
import { useMemo } from 'react'
import Link from 'next/link'

import { STAKE_PATH, getPoolUrl } from '@urls'

import { usePortfolioState } from '@/slices/portfolio/hooks'


import { formatBigIntToString } from '@/utils/bigint/format'






export const DisplayBalances = ({ pool, stakedBalance, showIcon, address }) => {
  const { poolTokenBalances } = usePortfolioState()
  const { amount, reward } = stakedBalance

  const lpTokenBalance = useMemo(() => {
    if (!address) {
      return null
    }
    const token = _.pickBy(poolTokenBalances[pool.chainId],
        (value, _key) =>
          value.tokenAddress === pool.addresses[pool.chainId] &&
          value.balance > 0n
      )

    if (Object.keys(token).length === 0) {
      return null
    } else {
      return token[0]
    }
  }, [pool, poolTokenBalances, address])

  const sum = useMemo(() => {
    const b =
      lpTokenBalance && lpTokenBalance.balance ? lpTokenBalance.balance : 0n
    return amount + b
  }, [lpTokenBalance, amount])

  if (!lpTokenBalance && amount === 0n && reward === 0n) {
    return null
  }

  return (
    <div className="flex items-center space-x-2">
      {showIcon && (
        <img src={pool.icon.src} className="w-[20px] h-[20px] rounded-full" />
      )}
      <div>
        <div className="flex items-center space-x-1">
          <div className="text-white text-md">
            <Link href={`${STAKE_PATH}/${pool.routerIndex}`}>
              <span className="hover:underline">
                {formatBigIntToString(amount, pool.decimals[pool.chainId], 5)}
              </span>
            </Link>
            <span className="text-[#BFBCC2] text-sm">
              {' '}
              /{' '}
              <Link href={getPoolUrl(pool)}>
                <span className="hover:underline">
                  {formatBigIntToString(sum, pool.decimals[pool.chainId], 5)}
                </span>
              </Link>
            </span>
          </div>
          <div className="text-sm text-[#BFBCC2]">{pool.symbol}</div>
        </div>
        {reward > 0n && (
          <div className="text-sm">
            <span className="text-white">Earned: </span>
            <span className="text-green-400 hover:underline">
              <Link href={`${STAKE_PATH}/${pool.routerIndex}`}>
                {formatBigIntToString(reward, 18, 5)}{' '}
                {pool?.customRewardToken ?? 'SYN'}
              </Link>
            </span>
          </div>
        )}
      </div>
    </div>
  )
}

