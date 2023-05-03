import Link from 'next/link'
import { STAKE_PATH, POOLS_PATH, POOL_PATH } from '@urls'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { useEffect, useState, memo } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { Token } from '@types'
import { getPoolData } from '@utils/actions/getPoolData'
import { getPoolApyData } from '@utils/actions/getPoolApyData'
import Button from '@tw/Button'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import PoolManagement from './poolManagement'
import PoolInfoSection from './PoolInfoSection'
const PoolBody = memo(
  ({
    pool,
    address,
    poolChainId,
    connectedChainId,
  }: {
    pool: Token
    address: string
    poolChainId: number
    connectedChainId: number
  }) => {
    const SynapseSDK = useSynapseContext()
    const [poolData, setPoolData] = useState(undefined)
    const [poolUserData, setPoolUserData] = useState(undefined)
    const [poolAPYData, setPoolAPYData] = useState(undefined)
    useEffect(() => {
      if (connectedChainId && pool && address && poolChainId) {
        // TODO - separate the apy and tvl so they load async.
        getPoolData(poolChainId, pool, address, false)
          .then((res) => {
            console.log('POOL BODY', '\nres:', res)
            setPoolData(res)
          })
          .catch((err) => {
            console.log('Could not get pool data', err)
          })
        getPoolData(poolChainId, pool, address, true)
          .then((res) => {
            console.log('POOL BODY', '\nres:', res)
            setPoolUserData(res)
          })
          .catch((err) => {
            console.log('Could not get pool data', err)
          })
        getPoolApyData(poolChainId, pool)
          .then((res) => {
            console.log('POOL BODY', '\nres:', res)
            setPoolAPYData(res)
          })
          .catch((err) => {
            console.log('Could not get pool data', err)
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
            <div className="mb-5">
              <div className="inline-flex items-center mt-2">
                <div className="items-center hidden mr-4 md:flex lg:flex">
                  {pool.poolTokens.map((token) => (
                    <img
                      key={token.symbol}
                      className="relative inline-block w-8 -mr-2 text-white shadow-solid"
                      src={token.icon.src}
                    />
                  ))}
                </div>
                <h3 className="ml-2 mr-2 text-lg font-medium text-white md:ml-0 md:text-2xl">
                  {pool.name}
                </h3>
              </div>
            </div>

            <div className="flex space-x-4">
              <div className="text-right">
                <div className="text-sm text-white text-opacity-60">APY</div>
                <div className="text-xl font-medium text-green-400">
                  {poolAPYData
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
            <Card className="bg-bgBase rounded-3xl" divider={false}>
              <PoolManagement
                pool={pool}
                address={address}
                chainId={connectedChainId}
                poolData={poolData}
                poolUserData={poolUserData}
                // poolStakingLink={STAKE_PATH}
                // poolStakingLinkText="Stake" // check this
              />
            </Card>
            <div>
              <PoolInfoSection
                pool={pool}
                poolData={poolData}
                chainId={connectedChainId}
              />
            </div>
          </Grid>
        </div>
      </>
    )
  }
)

export default PoolBody
