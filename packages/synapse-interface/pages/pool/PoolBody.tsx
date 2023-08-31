import Link from 'next/link'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { STAKE_PATH, POOLS_PATH } from '@urls'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import Button from '@tw/Button'
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
import { useEffect } from 'react'

const PoolBody = ({
  address,
  connectedChainId,
}: {
  address?: Address
  connectedChainId?: number
}) => {
  const { chains, switchNetwork } = useSwitchNetwork()
  const { openConnectModal } = useConnectModal()

  const { isConnected } = useAccount()

  const { pool, poolAPYData } = useSelector(
    (state: RootState) => state.poolData
  )

  useEffect(() => {
    if (pool) {
      segmentAnalyticsEvent(`[Pool] arrives at ${pool.name}`, {
        poolName: pool.poolName,
      })
    }
  }, [])

  return (
    <>
      <div className="px-0 md:px-32">
        <Link href={POOLS_PATH}>
          <div className="flex items-center mb-3 text-sm font-light text-white text-opacity-50 hover:text-opacity-100">
            <ChevronLeftIcon className="w-4 h-4" />
            Back to Pools
          </div>
        </Link>
        <div className="flex justify-between">
          <PoolTitle pool={pool} />
          <div className="flex space-x-4">
            <div className="text-right">
              <div className="text-sm text-white text-opacity-60">APY</div>
              <div className="text-xl font-medium text-green-400">
                {poolAPYData && Object.keys(poolAPYData).length > 0
                  ? `${String(poolAPYData.fullCompoundedAPYStr)}%`
                  : '-'}
              </div>
            </div>
            <Link href={STAKE_PATH}>
              <Button
                onClick={() => null}
                className="w-16 h-12 bg-bgLight hover:bg-bgLighter active:bg-bgLighter"
              >
                Stake
              </Button>
            </Link>
          </div>
        </div>
      </div>
      <div className="px-0 md:px-24">
        <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} gap={8}>
          <Card className="bg-bgBase rounded-lg" divider={false}>
            {!isConnected && (
              <div className="flex flex-col justify-center h-full">
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
          <div>
            <PoolInfoSection chainId={connectedChainId} />
          </div>
        </Grid>
      </div>
    </>
  )
}

const PoolTitle = ({ pool }) => {
  return (
    <div className="mb-5">
      <div className="inline-flex items-center mt-2">
        <div className="items-center hidden mr-4 md:flex lg:flex">
          {pool?.poolTokens &&
            pool.poolTokens.map((token) => (
              <img
                key={token.symbol}
                className="relative inline-block w-8 -mr-2 text-white shadow-solid"
                src={token.icon.src}
              />
            ))}
        </div>
        <h3 className="ml-2 mr-2 text-lg font-medium text-white md:ml-0 md:text-2xl">
          {pool?.name}
        </h3>
      </div>
    </div>
  )
}

export default PoolBody
