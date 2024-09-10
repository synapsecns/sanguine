import numeral from 'numeral'
import Link from 'next/link'
import { Address } from 'viem'
import { useEffect, useState } from 'react'
import { useAccount, useSwitchChain } from 'wagmi'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { useTranslations } from 'next-intl'

import Card from '@tw/Card'
import Grid from '@tw/Grid'
import { zeroAddress } from 'viem'
import { PoolActionOptions } from '@/components/Pools/PoolActionOptions'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { DisplayBalances } from '../pools/PoolCard'
import { usePoolDataState, usePoolUserDataState } from '@/slices/pools/hooks'
import { POOLS_PATH } from '@urls'
import PoolTitle from './components/PoolTitle'
import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './poolManagement'

const PoolBody = ({
  address,
  connectedChainId,
}: {
  address?: Address
  connectedChainId?: number
}) => {
  const [isClient, setIsClient] = useState(false)
  const { chains, switchChain } = useSwitchChain()
  const { openConnectModal } = useConnectModal()
  const { isConnected } = useAccount()

  const { poolUserData } = usePoolUserDataState()
  const { pool, poolAPYData } = usePoolDataState()

  const t = useTranslations('Pools')

  useEffect(() => {
    setIsClient(true)
  }, [])

  useEffect(() => {
    if (pool && isClient) {
      segmentAnalyticsEvent(`[Pool] arrives`, {
        poolName: pool?.poolName,
      })
    }
  }, [isClient, address, pool])

  if (!pool) return null

  return (
    <>
      <div id="pool-body">
        <Link href={POOLS_PATH}>
          <div className="inline-flex items-center mb-3 text-sm font-light text-white hover:text-opacity-100">
            <ChevronLeftIcon className="w-4 h-4" />
            {t('Back to Pools')}
          </div>
        </Link>
        <div className="flex justify-between">
          <PoolTitle pool={pool} />
          <div className="flex items-center space-x-4">
            <div className="hidden lg:flex">
              <DisplayBalances
                pool={pool}
                address={address}
                stakedBalance={poolUserData?.stakedBalance}
                showIcon={false}
              />
            </div>
            <PoolActionOptions
              pool={pool}
              options={[t('Stake'), t('Unstake'), t('Claim')]}
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
      <div className="">
        <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} gap={8}>
          <Card
            className="!pt-0 pb-0 pl-0 pr-0 rounded-md bg-bgBase"
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
                    chains.find((c) => c.id === pool.chainId)?.name
                  }`}
                  pendingLabel="Switching chains"
                  onClick={() =>
                    new Promise((resolve, reject) => {
                      try {
                        switchChain({ chainId: pool.chainId })
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
