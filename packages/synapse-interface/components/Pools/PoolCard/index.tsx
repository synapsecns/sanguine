import _ from 'lodash'
import { useEffect, useMemo, useState, memo } from 'react'
import Link from 'next/link'
import { Address, useAccount } from 'wagmi'
import { LoaderIcon, toast } from 'react-hot-toast'

import { Token } from '@types'
import { getPoolUrl } from '@urls'
import { getSinglePoolData } from '@utils/actions/getPoolData'
import { getPoolApyData } from '@utils/actions/getPoolApyData'
import { usePortfolioState } from '@/slices/portfolio/hooks'

import { useAppSelector } from '@/store/hooks'
import { useHasMounted } from '@/utils/hooks/useHasMounted'

import { getStakedBalance } from '@/utils/actions/getStakedBalance'

import { PoolActionOptions } from './PoolActionOptions'
import { PoolHeader } from './PoolHeader'
import { PoolCardBody } from './PoolCardBody'

import Card from '@tw/Card'
import { DisplayBalances } from '@/components/Pools/PoolCard/DisplayBalances'

const PoolCard = memo(({ pool, address }: { pool: Token; address: string }) => {
  const isClient = useHasMounted()
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
    <Card
      className={`
          group
          p-0
          transition-all
          ${
            pool && pool.incentivized
              ? 'ring-1 ring-white/10 '
              : 'from-transparent to-transparent border-transparent'
          }
          hover:ring-white/30
          rounded-md items-center
          space-y-2
          whitespace-wrap
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
    </Card>
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
    <div className="border-t border-slate-400/10 pt-2 pb-2 ">
      <div className="flex items-center justify-between pt-2 pl-3 pr-3">
        <DisplayBalances
          pool={pool}
          address={address}
          stakedBalance={stakedBalance}
          showIcon={true}
        />
        <div className="flex items-center text-xs">
          <PoolActionOptions
            pool={pool}
            options={['Deposit', 'Withdraw', 'Stake', 'Unstake', 'Claim']}
          />
        </div>
      </div>
    </div>
  )
}

export default PoolCard
