import Link from 'next/link'
import { STAKE_PATH, POOLS_PATH, POOL_PATH } from '@urls'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { useEffect, useState, memo } from 'react'
import { useSynapseContext } from '@/utils/SynapseProvider'
import { Token } from '@types'
import { useGenericPoolData } from '@hooks/pools/useGenericPoolData'
import PoolTitle from './components/PoolTitle'
import Button from '@tw/Button'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import PoolManagement from './components/PoolManagement'
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
    useEffect(() => {
      if (connectedChainId && pool && address && poolChainId) {
        // TODO - separate the apy and tvl so they load async.
        useGenericPoolData(poolChainId, pool, address, SynapseSDK)
          .then((res) => {
            setPoolData(res.poolDataObj)
          })
          .catch((err) => {
            console.log('ERROR useGenericPoolData: ', err)
          })
      }
    }, [])

    const apyData = poolData?.apy ?? {}

    let fullyCompoundedApyLabel
    if (isFinite(apyData.fullCompoundedAPY)) {
      fullyCompoundedApyLabel = apyData.fullCompoundedAPY?.toFixed(2)
    } else {
      fullyCompoundedApyLabel = <i className="opacity-50"> - </i>
    }
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
              <PoolTitle pool={pool} />
            </div>

            <div className="flex space-x-4">
              <div className="text-right">
                <div className="text-sm text-white text-opacity-60">APY</div>
                <div className="text-xl font-medium text-green-400">
                  {fullyCompoundedApyLabel}%
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
                // poolStakingLink={STAKE_PATH}
                // poolStakingLinkText="Stake" // check this
              />
            </Card>
            <div>
              <PoolInfoSection
                pool={pool}
                data={poolData}
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
