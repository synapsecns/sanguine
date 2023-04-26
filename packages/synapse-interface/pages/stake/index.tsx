import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useNetwork } from 'wagmi'
import {
  STAKABLE_TOKENS,
  STAKING_MAP_TOKENS,
  POOLS_BY_CHAIN,
} from '@/constants/tokens'
import { PageHeader } from '@/components/PageHeader'

const StakePage = () => {
  const { chain: connectedChain } = useNetwork()
  const connectedChainId = connectedChain ? connectedChain.id : undefined

  const availableStakingTokens = POOLS_BY_CHAIN[connectedChainId]
  console.log('availableStakingTokens: ', availableStakingTokens)

  return (
    <LandingPageWrapper>
      <main
        data-test-id="stake-page"
        className="flex items-center justify-between px-20 mx-10 space-x-2"
      >
        <PageHeader title="Stake" subtitle="Stake your LP Tokens." />
      </main>
    </LandingPageWrapper>
  )
}

export default StakePage
