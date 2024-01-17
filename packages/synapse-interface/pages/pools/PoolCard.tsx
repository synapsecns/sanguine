import _ from 'lodash'
import { useEffect, useMemo, useState, memo } from 'react'
import Link from 'next/link'
import { Address, useAccount } from 'wagmi'
import { LoaderIcon, toast } from 'react-hot-toast'

import { Token } from '@types'
import { STAKE_PATH, getPoolUrl } from '@urls'
import { getSinglePoolData } from '@utils/actions/getPoolData'
import { getPoolApyData } from '@utils/actions/getPoolApyData'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'
import { formatBigIntToString } from '@/utils/bigint/format'
import { PoolActionOptions } from '../../components/Pools/PoolActionOptions'
import { PoolHeader } from '../../components/Pools/PoolHeader'
import { PoolCardBody } from '../../components/Pools/PoolCardBody'
import { useAppSelector } from '@/store/hooks'

const PoolCard = memo(({ pool, address }: { pool: Token; address: string }) => {
  const [isClient, setIsClient] = useState(false)
  const [poolData, setPoolData] = useState(undefined)
  const [poolApyData, setPoolApyData] = useState(undefined)
  const [stakedBalance, setStakedBalance] = useState({
    amount: 0n,
    reward: 0n,
  })
  const { isDisconnected } = useAccount()
  const { synPrices, ethPrice, avaxPrice, metisPrice } = useAppSelector(
    (state) => state.priceData
  )

  const prices = { synPrices, ethPrice, avaxPrice, metisPrice }

  let popup: string
  useEffect(() => {
    setIsClient(true)
  }, [])

  useEffect(() => {
    if (pool && isClient) {
      // TODO - separate the apy and tvl so they load async.
      getSinglePoolData(pool.chainId, pool, prices)
        .then((res) => {
          setPoolData(res)
        })
        .catch((err) => {
          console.log('Could not get Pool Data: ', err)
        })
      getPoolApyData(pool.chainId, pool, prices)
        .then((res) => {
          setPoolApyData(res)
        })
        .catch((err) => {
          console.log('Could not get Pool APY Data: ', err)
        })
    }
  }, [pool, isClient])

  useEffect(() => {
    if (address && isClient) {
      getStakedBalance(
        address as Address,
        pool.chainId,
        pool.poolId[pool.chainId],
        pool
      )
        .then((res) => {
          setStakedBalance(res)
        })
        .catch((err) => {
          console.log('Could not get staked balances: ', err)
        })
    } else {
      setStakedBalance({ amount: 0n, reward: 0n })
    }
  }, [address, isClient])

  /*
  useEffect triggers: address, isDisconnected, popup
  - will dismiss toast asking user to connect wallet once wallet has been connected
  */
  useEffect(() => {
    if (address && !isDisconnected && popup) {
      toast.dismiss(popup)
    }
  }, [address, isDisconnected, popup])

  return (
    <div
      className={`
          border
          rounded-md h-max
          ${pool && pool.incentivized
            ? 'bg-zinc-100 dark:bg-zinc-800 border-zinc-300 dark:border-transparent'
            : 'bg-transparent border-zinc-300 dark:border-zinc-700'
          }
        `}
    >
      <div>
        <Link href={getPoolUrl(pool)}>
          {pool && <PoolHeader pool={pool} address={address as Address} />}
          {pool &&
          poolData &&
          poolApyData &&
          poolData.tokens &&
          poolData.totalLockedUSD ? (
            <PoolCardBody
              pool={pool}
              poolApyData={poolApyData}
              poolData={poolData}
            />
          ) : (
            <div className="flex items-center justify-between px-3 pt-1 pb-2 h-[65px]">
              <LoaderIcon />
            </div>
          )}
        </Link>
        {pool && (
          <>
            <ManageLp
              pool={pool}
              stakedBalance={stakedBalance}
              address={address}
            />
          </>
        )}
      </div>
    </div>
  )
})

const ManageLp = ({ pool, stakedBalance, address }) => {
  const { poolTokenBalances } = usePortfolioState()
  const { amount, reward } = stakedBalance

  const lpTokenBalance = useMemo(() => {
    if (!address) {
      return null
    }
    const token = _(poolTokenBalances[pool.chainId])
      .pickBy(
        (value, _key) =>
          value.tokenAddress === pool.addresses[pool.chainId] &&
          value.balance > 0n
      )
      .value()
    if (Object.keys(token).length === 0) {
      return null
    } else {
      return token[0]
    }
  }, [pool, poolTokenBalances, address])

  if (!lpTokenBalance && amount === 0n && reward === 0n) {
    return null
  }

  return (
    <DisplayBalances
      pool={pool}
      address={address}
      stakedBalance={stakedBalance}
      showIcon={true}
    />
  )
}

export const DisplayBalances = ({ pool, stakedBalance, showIcon, address }) => {
  const { poolTokenBalances } = usePortfolioState()
  const { amount, reward } = stakedBalance

  const lpTokenBalance = useMemo(() => {
    if (!address) {
      return null
    }
    const token = _(poolTokenBalances[pool.chainId])
      .pickBy(
        (value, _key) =>
          value.tokenAddress === pool.addresses[pool.chainId] &&
          value.balance > 0n
      )
      .value()
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
    <div className="flex items-center py-2 px-3 border-t border-zinc-200 dark:border-zinc-700 gap-2 w-full">
      {showIcon && (
        <img src={pool.icon.src} className="w-5 h-5" />
      )}
      <div className="flex-grow w-0">
        <div className="overflow-hidden">
          <Link href={`${STAKE_PATH}/${pool.routerIndex}`}>
            <span className="hover:underline">
              {formatBigIntToString(amount, pool.decimals[pool.chainId], 5)}
            </span>
          </Link>
          <span className="text-zinc-400 text-sm">
            &nbsp;/&nbsp;
            <Link href={getPoolUrl(pool)}>
              <span className="hover:underline">
                {formatBigIntToString(sum, pool.decimals[pool.chainId], 5)}
              </span>
            </Link>
            &nbsp;
            <span className="overflow-ellipsis">
              {pool.symbol}
            </span>
          </span>
        </div>
        {reward > 0n && (
          <div className="text-sm">
            Earned: <span className="text-green-500 hover:underline">
              <Link href={`${STAKE_PATH}/${pool.routerIndex}`}>
                {formatBigIntToString(reward, 18, 5)}{' '}
                {pool?.customRewardToken ?? 'SYN'}
              </Link>
            </span>
          </div>
        )}
      </div>
      <PoolActionOptions
        pool={pool}
        options={['Deposit', 'Withdraw', 'Stake', 'Unstake', 'Claim']}
      />
    </div>
  )
}
export default PoolCard
