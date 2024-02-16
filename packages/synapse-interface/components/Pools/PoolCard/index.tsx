import _ from 'lodash'
import { useEffect, useMemo, useState, memo } from 'react'
import Link from 'next/link'
import { Address, useAccount } from 'wagmi'
import { toast } from 'react-hot-toast'

import { Token } from '@types'
import { getPoolUrl } from '@urls'
import { getSinglePoolData } from '@utils/actions/getPoolData'
import { getPoolApyData } from '@utils/actions/getPoolApyData'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'

import { usePortfolioState } from '@/slices/portfolio/hooks'
import { useAppSelector } from '@/store/hooks'

import Card from '@tw/Card'
import { LoadingHelix } from '@tw/LoadingHelix'

import { DisplayBalances } from './DisplayBalances'
import { PoolActionOptions } from './PoolActionOptions'
import { PoolHeader } from './PoolHeader'
import { PoolCardBody } from './PoolCardBody'


const PoolCard = memo(({ pool, address }: { pool: Token; address: string }) => {
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
    if (pool) {
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
  }, [pool])

  useEffect(() => {
    if (address) {
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
  }, [address])

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
            pool?.incentivized
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
          {pool
            ?
              <PoolCardBody
                pool={pool}
                poolApyData={poolApyData}
                poolData={poolData}
              />
            :
              <div className="flex items-center justify-between px-3 pt-1 pb-2 h-[65px]">
                <LoadingHelix />
              </div>
          }
        </Link>
        {pool && (
            <ManageLp
              pool={pool}
              stakedBalance={stakedBalance}
              address={address}
            />
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
