import _ from 'lodash'
import { useRouter } from 'next/router'
import { useEffect, useMemo, useState } from 'react'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { DISPLAY_POOLS_BY_CHAIN } from '@constants/tokens'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import {
  METIS_POOL_SWAP_TOKEN_MIGRATED,
  METIS_WETH_SWAP_TOKEN_MIGRATED,
} from '@/constants/tokens/poolMaster'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { PageHeader } from '@/components/PageHeader'
import Grid from '@/components/ui/tailwind/Grid'
import PoolCards from './PoolCards'
import * as CHAINS from '@/constants/chains/master'

export async function getStaticProps({ locale }) {
  return {
    props: {
      messages: (await import(`../../messages/${locale}.json`)).default,
    },
  }
}

const PoolsPage = () => {
  const { address: currentAddress } = useAccount()
  const { chain } = useAccount()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)

  const router = useRouter()

  const t = useTranslations('Pools')

  const migratedPools = {
    1088: [METIS_POOL_SWAP_TOKEN_MIGRATED, METIS_WETH_SWAP_TOKEN_MIGRATED],
  }

  const blastPools = filterPoolsByBlast()

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
            title={t('Incentivized Pools')}
            subtitle={t('Contributors are rewarded')}
          />
        </div>
        <Grid cols={{ xs: 1, sm: 1, md: 2 }} gap={4} className="mb-5">
          <PoolCards address={address} pools={blastPools} />
          <PoolCards address={address} pools={incentivizedPools} />
        </Grid>
        <div className="flex-wrap justify-between mt-8 mb-4 md:flex">
          <PageHeader
            title={t('Unincentivized Pools')}
            subtitle={t('Pools without contributor rewards')}
          />
        </div>
        <Grid cols={{ xs: 1, sm: 1, md: 2 }} gap={4} className="mb-5">
          <PoolCards address={address} pools={unIncentivizedPools} />
        </Grid>
        <div className="flex-wrap justify-between mt-8 mb-4 md:flex">
          <PageHeader
            title={t('Migrated Pools')}
            subtitle={t('Pools migrated to new reward contracts')}
          />
        </div>
        <Grid cols={{ xs: 1, sm: 1, md: 2 }} gap={4} className="mb-5">
          <PoolCards address={address} pools={migratedPools} />
        </Grid>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

const filterPoolsByBlast = () =>
  _.pickBy(
    _.mapValues(DISPLAY_POOLS_BY_CHAIN, (tokens) =>
      tokens.filter((token) => token.chainId === CHAINS.BLAST.id)
    ),
    (tokens) => tokens.length > 0
  )

const filterPoolsByIncentivization = (incentivized) =>
  _.pickBy(
    _.mapValues(DISPLAY_POOLS_BY_CHAIN, (tokens) =>
      tokens.filter(
        (token) =>
          token.incentivized === incentivized &&
          token.chainId !== CHAINS.BLAST.id
      )
    ),
    (tokens) => tokens.length > 0
  )

export default PoolsPage
