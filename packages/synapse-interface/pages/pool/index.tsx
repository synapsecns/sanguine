import _ from 'lodash'
import { useEffect } from 'react'
// import toast from 'react-hot-toast'
import { ChevronLeftIcon } from '@heroicons/react/outline'

import { STAKE_PATH, POOLS_PATH } from '@urls'

import { CHAIN_INFO_MAP } from '@constants/networks'
import { CHAINS_BY_POOL_NAME } from '@constants/tokens/poolsByChain'
import { POOL_ROUTER_INDEX } from '@constants/poolRouter'

import { getNetworkTextColor } from '@styles/networks'

import { usePoolData } from '@hooks/pools/usePoolData'

import Card from '@tw/Card'
import Grid from '@tw/Grid'
import Button from '@tw/Button'
import Link from 'next/link'

import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'

import PoolInfoSection from './PoolInfoSection'
import PoolManagement from './PoolManagement'
import PoolTitle from './PoolTitle'

import { useAccount, useNetwork } from 'wagmi'
import Grid from '@tw/Grid'
import SwapCard from './SwapCard'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { PageHeader } from '@components/PageHeader'
import { SWAPABLE_TOKENS } from '@constants/tokens'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import NoSwapCard from './NoSwapCard'
import { useEffect, useState, useMemo } from 'react'

export default function PoolPage({
  match: {
    params: { id },
  },
}) {
  const poolName = POOL_ROUTER_INDEX[id]

  const { chainId } = useActiveWeb3React()

  const poolChainId = CHAINS_BY_POOL_NAME[poolName]

  const isPoolOnChain = chainId == poolChainId

  useEffect(() => {
    if (!isPoolOnChain) {
      const { chainName } = CHAIN_INFO_MAP[chainId] ?? {}
      if (chainName) {
        // toast(`Viewing pools on ${chainName}`)
      }
    }
  }, [isPoolOnChain, chainId])

  return (
    <LandingPageWrapper>
      <StandardPageContainer>
        {isPoolOnChain && (
          <div className="px-0 md:px-32">
            <Link to={POOLS_PATH}>
              <div className="flex items-center mb-3 text-sm font-light text-white text-opacity-50 hover:text-opacity-100">
                <ChevronLeftIcon className="w-4 h-4" />
                Back to Pools
              </div>
            </Link>
            <PoolHeader poolName={poolName} isPoolOnChain={isPoolOnChain} />
          </div>
        )}
        {isPoolOnChain && (
          <div className="px-0 md:px-24">
            <PoolPageContents
              poolName={poolName}
              isPoolOnChain={isPoolOnChain}
            />
          </div>
        )}
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
        )}
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

function PoolHeader({ poolName, isPoolOnChain }) {
  const [poolData, userShareData] = usePoolData(poolName)
  const apyData = poolData?.apy ?? {}

  const poolChainId = CHAINS_BY_POOL_NAME[poolName]

  let fullyCompoundedApyLabel
  if (_.isFinite(apyData.fullCompoundedAPY)) {
    fullyCompoundedApyLabel = apyData.fullCompoundedAPY?.toFixed(2)
  } else {
    fullyCompoundedApyLabel = <i className="opacity-50"> - </i>
  }

  return (
    <div className="flex justify-between">
      <div className="mb-5">
        <PoolTitle poolName={poolName} poolChainId={poolChainId} />
      </div>
      {isPoolOnChain && (
        <div className="flex space-x-4">
          <div className="text-right">
            <div className="text-sm text-white text-opacity-60">APY</div>
            <div className="text-xl font-medium text-green-400">
              {fullyCompoundedApyLabel}%
            </div>
          </div>
          <Link to={STAKE_PATH}>
            <Button className="w-16 h-12 bg-bgLight hover:bg-bgLighter active:bg-bgLighter">
              Stake
            </Button>
          </Link>
        </div>
      )}
    </div>
  )
}

function PoolPageContents({ poolName, isPoolOnChain }) {
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
