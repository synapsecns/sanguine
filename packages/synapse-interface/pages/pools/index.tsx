import _ from 'lodash'
import { useEffect, useMemo, useState } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { useRouter } from 'next/router'


import { DISPLAY_POOLS_BY_CHAIN } from '@/constants/tokens'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import {
  METIS_POOL_SWAP_TOKEN_MIGRATED,
  METIS_WETH_SWAP_TOKEN_MIGRATED,
} from '@/constants/tokens/poolMaster'


import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'



import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'
import { PageHeader } from '@/components/PageHeader'
import Grid from '@tw/Grid'
import PoolCards from '@/components/Pools/PoolCards'

const PoolsPage = () => {
  const { address: currentAddress } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)

  const router = useRouter()

  const migratedPools = {
    1088: [METIS_POOL_SWAP_TOKEN_MIGRATED, METIS_WETH_SWAP_TOKEN_MIGRATED],
  }

  const incentivizedPools = useMemo(
    () => filterPoolsByIncentivization(true),
    [DISPLAY_POOLS_BY_CHAIN]
  )
  let unIncentivizedPools = useMemo(
    () => filterPoolsByIncentivization(false),
    [DISPLAY_POOLS_BY_CHAIN]
  )
  unIncentivizedPools[1088] = unIncentivizedPools[1088].filter(
    (pool) =>
      pool !== METIS_POOL_SWAP_TOKEN_MIGRATED &&
      pool !== METIS_WETH_SWAP_TOKEN_MIGRATED
  )

  useEffect(() => {
    segmentAnalyticsEvent(`[Pools page] arrives`, {
      fromChainId: chain?.id,
      query: router.query,
      pathname: router.pathname,
    })
  }, [])

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])
  useEffect(() => {
    setAddress(currentAddress)
  }, [currentAddress])

  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        <div className="flex-wrap justify-between mb-4 md:flex">
          <PageHeader
            title="Incentivized Pools"
            subtitle="Contributors are rewarded for balancing asset pools."
          />
        </div>
        <Grid cols={{ xs: 1, sm: 1, md: 2 }} gap={4} className="mb-5">
          <PoolCards address={address} pools={incentivizedPools} />
        </Grid>
        <div className="flex-wrap justify-between mt-8 mb-4 md:flex">
          <PageHeader
            title="Unincentivized Pools"
            subtitle="Pools without contributor rewards."
          />
        </div>
        <Grid cols={{ xs: 1, sm: 1, md: 2 }} gap={4} className="mb-5">
          <PoolCards address={address} pools={unIncentivizedPools} />
        </Grid>
        <div className="flex-wrap justify-between mt-8 mb-4 md:flex">
          <PageHeader
            title="Migrated Pools"
            subtitle="Pools migrated to new reward contracts."
          />
        </div>
        <Grid cols={{ xs: 1, sm: 1, md: 2 }} gap={4} className="mb-5">
          <PoolCards address={address} pools={migratedPools} />
        </Grid>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

const filterPoolsByIncentivization = (incentivized) =>
  _.pickBy(
    _.mapValues(DISPLAY_POOLS_BY_CHAIN, (tokens) =>
      tokens.filter((token) => token.incentivized === incentivized)
    ),
    (tokens) => tokens.length > 0
  )

export default PoolsPage
