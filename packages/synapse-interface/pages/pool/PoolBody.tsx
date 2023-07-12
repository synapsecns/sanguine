import { useEffect } from 'react'
import Link from 'next/link'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { STAKE_PATH, POOLS_PATH, POOL_PATH } from '@urls'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import Button from '@tw/Button'
import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './poolManagement'
import { zeroAddress } from 'viem'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '@/store/store'
import {
  fetchPoolUserData,
  resetPoolUserData,
} from '@/slices/poolUserDataSlice'
import { Address } from '@wagmi/core'

const PoolBody = ({
  address,
  connectedChainId,
}: {
  address?: Address
  connectedChainId?: number
}) => {
  const { pool, poolAPYData } = useSelector(
    (state: RootState) => state.poolData
  )

  const dispatch: any = useDispatch()

  useEffect(() => {
    if (pool && address) {
      dispatch(resetPoolUserData())
      dispatch(fetchPoolUserData({ pool, address }))
    }
  }, [pool, address])

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
              address={address ?? zeroAddress}
              chainId={connectedChainId}
            />
          </Card>
          <div>
            <PoolInfoSection chainId={connectedChainId} />
          </div>
        </Grid>
      </div>
    </>
  )
}

export default PoolBody
