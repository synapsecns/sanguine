import toast from 'react-hot-toast'
import { useEffect, useState } from 'react'
import { useNetwork, useAccount } from 'wagmi'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { ChevronLeftIcon } from '@heroicons/react/outline'
import { Token } from '@/utils/types'
import { STAKABLE_TOKENS } from '@/constants/tokens'
import Grid from '@tw/Grid'
import { PageHeader } from '@/components/PageHeader'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import StakeCard from './StakeCard'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { POOLS_PATH } from '@/constants/urls'
import { useHasMounted } from '@/utils/hooks/useHasMounted'



const StakePage = () => {
  const { chain: connectedChain } = useNetwork()
  const { address: currentAddress } = useAccount()

  const isClient = useHasMounted()
  const [columns, setColumns] = useState<number>(1)

  const router = useRouter()
  const { query, pathname } = router

  const availableStakingTokens: Token[] | [] =
    connectedChain && STAKABLE_TOKENS[connectedChain.id]
      ? STAKABLE_TOKENS[connectedChain.id]
      : []

  const routerIndices = availableStakingTokens.map((token) => token.routerIndex)
  useEffect(() => {
    segmentAnalyticsEvent(`[Stake page] arrives`, {
      connectedChainId: connectedChain ? connectedChain.id : null,
      query,
      pathname,
      routerIndices,
    })
  }, [])



  useEffect(() => {
    const isSingle = availableStakingTokens.length < 2
    setColumns(isSingle ? 1 : 2)

    if (availableStakingTokens.length === 0) {
      router.push('/pools')
    }
  }, [availableStakingTokens, router])


  if (!connectedChain) {
    toast.error('Please connect to see stakes', {
      id: 'approve-in-progress-popup',
      duration: 5000,
    })

    return <LandingPageWrapper> </LandingPageWrapper>
  }

  if (connectedChain && availableStakingTokens.length === 0) {
    toast.error(`No stakes available for ${connectedChain.name} network`, {
      id: 'approve-in-progress-popup',
      duration: 5000,
    })

    return <LandingPageWrapper> </LandingPageWrapper>
  }

  return (
    <LandingPageWrapper>
      <main
        data-test-id="stake-page"
        className={`
          flex flex-col justify-between
          px-4 py-16
          md:px-20 md:py-3 md:m-14
        `}
      >
        <div className="flex flex-col justify-center max-w-[1300px] m-auto">
          <div>
            <Link href={POOLS_PATH}>
              <div className="inline-flex items-center mb-3 text-sm  text-white/80 hover:text-white/100">
                <ChevronLeftIcon className="w-4 h-4" />
                Back to Pools
              </div>
            </Link>
          </div>
          <PageHeader title="Stake" subtitle="Stake your LP Tokens." />

          <Grid cols={{ xs: 1, sm: 1, md: columns }} gap={6} className="mt-8">
            {isClient &&
              availableStakingTokens.map((token, key) => {
                if (token.notStake) {
                  return null
                }
                return (
                  <StakeCard
                    key={key}
                    address={currentAddress}
                    chainId={connectedChain.id}
                    pool={token}
                  />
                )
              })}
          </Grid>
        </div>
      </main>
    </LandingPageWrapper>
  )
}

export default StakePage
