import _ from 'lodash'
// import { useEffect } from 'react'
// // import toast from 'react-hot-toast'
import { ChevronLeftIcon } from '@heroicons/react/outline'

import { STAKE_PATH, POOLS_PATH, POOL_PATH } from '@urls'
import { useSynapseContext } from '@/utils/SynapseProvider'
import { useGenericPoolData } from '@hooks/pools/useGenericPoolData'
import { Token } from '@types'
// import { CHAIN_INFO_MAP } from '@constants/networks'
// import { CHAINS_BY_POOL_NAME } from '@constants/tokens'
// import { POOL_ROUTER_INDEX } from '@constants/poolRouter'

// import { getNetworkTextColor } from '@styles/networks'

// import { usePoolData } from '@hooks/pools/usePoolData'

import Card from '@tw/Card'
import Grid from '@tw/Grid'
import Button from '@tw/Button'
import Link from 'next/link'

// import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'

import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './PoolManagement'
import PoolTitle from './PoolTitle'
import { useRouter } from 'next/router'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'

import { useAccount, useNetwork } from 'wagmi'
// import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { useEffect, useState, useMemo } from 'react'
import { POOL_BY_ROUTER_INDEX, POOL_CHAINS_BY_NAME } from '@constants/tokens'
console.log('POOL_CHAINS_BY_NAME', POOL_CHAINS_BY_NAME)
const PoolPage = () => {
  const router = useRouter()
  const { poolId } = router.query
  const { address: currentAddress } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)
  const [pool, setPool] = useState(undefined)
  // const poolName = POOL_ROUTER_INDEX[poolId]
  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])
  useEffect(() => {
    setAddress(currentAddress)
  }, [currentAddress])
  useEffect(() => {
    const poolFromName = POOL_BY_ROUTER_INDEX[String(poolId)]
    setPool(poolFromName)
  }, [poolId])

  // const poolChainId = CHAINS_BY_POOL_NAME[poolName]

  // const isPoolOnChain = chainId == poolChainId

  // useEffect(() => {
  //   if (!isPoolOnChain) {
  //     const { chainName } = CHAIN_INFO_MAP[chainId] ?? {}
  //     if (chainName) {
  //       // toast(`Viewing pools on ${chainName}`)
  //     }
  //   }
  // }, [isPoolOnChain, chainId])
  const poolChainId = useMemo(
    () => (pool?.addresses ? Number(Object.keys(pool?.addresses)[0]) : 0),
    [pool]
  )
  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        {connectedChainId === poolChainId && (
          <div className="px-0 md:px-32">
            <Link href={POOLS_PATH}>
              <div className="flex items-center mb-3 text-sm font-light text-white text-opacity-50 hover:text-opacity-100">
                <ChevronLeftIcon className="w-4 h-4" />
                Back to Pools
              </div>
            </Link>
            <PoolHeader
              pool={pool}
              poolChainId={poolChainId}
              address={address}
              connectedChainId={connectedChainId}
            />
          </div>
        )}

        {connectedChainId === poolChainId && (
          <div className="px-0 md:px-24">
            <PoolPageContents pool={pool} poolChainId={poolChainId} />
          </div>
        )}
        {/*
        {!isPoolOnChain && (
          <Grid cols={{ xs: 1 }} gap={2}>
            <Card
              title="Pool Info "
              className={`
                bg-bgBase
                my-8 transform transition-all duration-100 rounded-3xl place-self-center
                min-w-4/5 sm:min-w-3/4 md:min-w-3/5 lg:min-w-1/2
              `}
              divider={false}
            >
              <div
                className={`
                  pt-4 text-gray-400 w-full text-center
                `}
              >
                Switch to{' '}
                <span
                  className={`${getNetworkTextColor(poolChainId)} font-medium`}
                >
                  {CHAIN_INFO_MAP[poolChainId].chainName}
                </span>{' '}
                to interact with the {poolName}
              </div>
            </Card>
          </Grid>
        )}*/}
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

const PoolHeader = ({
  pool,
  poolChainId,
  connectedChainId,
  address,
}: {
  pool: Token
  poolChainId: number
  connectedChainId: number
  address: string
}) => {
  const [poolData, setPoolData] = useState(undefined)
  const SynapseSDK = useSynapseContext()

  console.log('PoolsListCard RERENDER')
  useEffect(() => {
    if (
      connectedChainId === undefined ||
      pool === undefined ||
      address === undefined ||
      poolChainId === undefined
    ) {
      return
    }
    // TODO - separate the apy and tvl so they load async.
    useGenericPoolData(poolChainId, pool, address, SynapseSDK)
      .then((res) => {
        setPoolData(res.poolDataObj)
      })
      .catch((err) => {
        console.log('ERROR useGenericPoolData: ', err)
      })
  }, [])
  const apyData = poolData?.apy ?? {}

  let fullyCompoundedApyLabel
  if (_.isFinite(apyData.fullCompoundedAPY)) {
    fullyCompoundedApyLabel = apyData.fullCompoundedAPY?.toFixed(2)
  } else {
    fullyCompoundedApyLabel = <i className="opacity-50"> - </i>
  }

  return (
    <div className="flex justify-between">
      <div className="mb-5">
        <PoolTitle pool={pool} />
      </div>
      {poolChainId === connectedChainId && (
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
      )}
    </div>
  )
}

const PoolPageContents = ({ poolName, isPoolOnChain }) => {
  const [poolData, userShareData] = usePoolData(poolName)

  return (
    <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 2 }} gap={8}>
      <Card className="bg-bgBase rounded-3xl" divider={false}>
        {isPoolOnChain && (
          <PoolManagement poolName={poolName} poolStakingLink={STAKE_PATH} />
        )}
      </Card>
      <div>
        {isPoolOnChain && (
          <PoolInfoSection
            poolName={poolName}
            data={poolData}
            userData={userShareData}
          />
        )}
      </div>
    </Grid>
  )
}
export default PoolPage
