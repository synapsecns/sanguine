import numeral from 'numeral'
import Link from 'next/link'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { POOLS_PATH } from '@urls'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './poolManagement'
import { zeroAddress } from 'viem'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'
import { Address } from '@wagmi/core'
import { useAccount, useSwitchNetwork } from 'wagmi'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useEffect, useState } from 'react'
import { PoolActionOptions } from '@/components/Pools/PoolActionOptions'
import PoolTitle from './components/PoolTitle'
import { DisplayBalances } from '../pools/PoolCard'
import { getStakedBalance } from '@/utils/actions/getStakedBalance'

const PoolBody = ({
  address,
  connectedChainId,
}: {
  address?: Address
  connectedChainId?: number
}) => {
  const [isClient, setIsClient] = useState(false)
  const { chains, switchNetwork } = useSwitchNetwork()
  const { openConnectModal } = useConnectModal()

  const { isConnected } = useAccount()

  const { pool, poolAPYData } = useSelector(
    (state: RootState) => state.poolData
  )
  const [stakedBalance, setStakedBalance] = useState({
    amount: 0n,
    reward: 0n,
  })

  useEffect(() => {
    setIsClient(true)
  }, [])

  useEffect(() => {
    if (pool && isClient) {
      segmentAnalyticsEvent(`[Pool] arrives`, {
        poolName: pool?.poolName,
      })
    }
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
  }, [isClient, address, pool])

  if (!pool) return null

  return (
    <>
      <div className="">
        <Link href={POOLS_PATH}>
          <div className="inline-flex items-center mb-3 text-sm hover:text-opacity-100">
            <ChevronLeftIcon className="w-4 h-4" />
            Back to Pools
          </div>
        </Link>
        <div className="flex justify-between">
          <PoolTitle pool={pool} />
          <div className="flex items-center space-x-4">
            <div className="hidden lg:flex">
              <DisplayBalances
                pool={pool}
                address={address}
                stakedBalance={stakedBalance}
                showIcon={false}
              />
            </div>
            <PoolActionOptions
              pool={pool}
              options={['Stake', 'Unstake', 'Claim']}
            />
            <div className="flex space-x-4">
              <div className="text-right">
                <div className="text-xl font-medium">
                  {poolAPYData && Object.keys(poolAPYData).length > 0
                    ? `${numeral(poolAPYData.fullCompoundedAPY / 100).format(
                        '0.0%'
                      )}`
                    : '-'}
                </div>
                <div className="text-sm">APY</div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="">
        <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} gap={8}>
          <Card
            className="rounded-md bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-transparent p-0"
            divider={false}
          >
            {!isConnected && (
              <div className="flex flex-col justify-center h-full p-10">
                <TransactionButton
                  label="Connect wallet"
                  pendingLabel="Connecting"
                  onClick={() =>
                    new Promise((resolve, reject) => {
                      try {
                        openConnectModal()
                        resolve(true)
                      } catch (e) {
                        reject(e)
                      }
                    })
                  }
                />
              </div>
            )}
            {isConnected && connectedChainId !== pool.chainId && (
              <div className="flex flex-col justify-center h-full">
                <TransactionButton
                  label={`Switch to ${
                    chains.find((c) => c.id === pool.chainId).name
                  }`}
                  pendingLabel="Switching chains"
                  onClick={() =>
                    new Promise((resolve, reject) => {
                      try {
                        switchNetwork(pool.chainId)
                        resolve(true)
                      } catch (e) {
                        reject(e)
                      }
                    })
                  }
                />
              </div>
            )}
            {isConnected && connectedChainId === pool.chainId && (
              <PoolManagement
                address={address ?? zeroAddress}
                chainId={connectedChainId}
              />
            )}
          </Card>
          <PoolInfoSection />
        </Grid>
      </div>
    </>
  )
}

export default PoolBody
