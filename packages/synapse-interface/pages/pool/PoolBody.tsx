import numeral from 'numeral'
import Link from 'next/link'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { POOLS_PATH } from '@urls'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './poolManagement'
import { zeroAddress } from 'viem'

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
import { usePoolDataState } from '@/slices/pool/hooks'

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

  const { pool, poolAPYData } = usePoolDataState()

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
      <div>
        <Link href={POOLS_PATH}>
          <div className="inline-flex items-center mb-3 text-sm  text-white/70 hover:text-white/100 group">
            <ChevronLeftIcon className="w-4 h-4 mr-1 mt-0.25 group-hover:stroke-2" />
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
                <div className="text-xl font-medium text-white">
                  {poolAPYData && Object.keys(poolAPYData).length > 0
                    ? `${numeral(poolAPYData.fullCompoundedAPY / 100).format(
                        '0.0%'
                      )}`
                    : '-'}
                </div>
                <div className="text-sm text-white text-opacity-60">APY</div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div>
        <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} gap={8}>
          <Card
            className="!p-0 rounded-md bg-bgBase"
            divider={false}
          >
            {!isConnected && (
              <div className="flex flex-col justify-center h-full p-10">
                <TransactionButton
                  style={{
                    background:
                      'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
                    border: '1px solid #9B6DD7',
                    borderRadius: '4px',
                  }}
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
              <div className="flex flex-col justify-center h-full p-10">
                <TransactionButton
                  style={{
                    background:
                      'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
                    border: '1px solid #9B6DD7',
                    borderRadius: '4px',
                  }}
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
          <div>
            <PoolInfoSection />
          </div>
        </Grid>
      </div>
    </>
  )
}

export default PoolBody
