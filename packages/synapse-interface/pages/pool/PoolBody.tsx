import { useEffect, useState, useCallback } from 'react'
import { AddressZero } from '@ethersproject/constants'
import Link from 'next/link'
import { Token } from '@types'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { getPoolData } from '@utils/actions/getPoolData'
import { getPoolApyData } from '@utils/actions/getPoolApyData'
import { STAKE_PATH, POOLS_PATH, POOL_PATH } from '@urls'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import Button from '@tw/Button'
import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './poolManagement'

const PoolBody = ({
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
  const [poolData, setPoolData] = useState(undefined)
  const [poolUserData, setPoolUserData] = useState(undefined)
  const [poolAPYData, setPoolAPYData] = useState(undefined)

  const handleGetPoolData = useCallback(() => {
    getPoolData(poolChainId, pool, address ?? AddressZero, false)
      .then((res) => {
        return setPoolData(res)
      })
      .catch((err) => {
        console.log('Could not get pool data', err)
        return err
      })
  }, [poolChainId, pool, address])

  const handleGetUserPoolData = useCallback(() => {
    if (address) {
      getPoolData(poolChainId, pool, address, true)
        .then((res) => {
          return setPoolUserData(res)
        })
        .catch((err) => {
          console.log('Could not get pool data', err)
          return err
        })
    }
  }, [poolChainId, pool, address])

  useEffect(() => {
    if (connectedChainId && pool && poolChainId) {
      // TODO - separate the apy and tvl so they load async.
      handleGetPoolData()
      handleGetUserPoolData()
      getPoolApyData(poolChainId, pool)
        .then((res) => {
          if (Object.keys(res).length > 0) {
            setPoolAPYData(res)
          }
        })
        .catch((err) => {
          console.log('Could not get pool data', err)
        })
    }
  }, [connectedChainId, pool, poolChainId, address])

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
              address={address ?? AddressZero}
              chainId={connectedChainId}
              poolData={poolData}
              poolUserData={poolUserData}
              refetchCallback={handleGetUserPoolData}
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

export default PoolBody
