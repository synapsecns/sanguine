import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useNetwork } from 'wagmi'
import {
  STAKABLE_TOKENS,
  STAKING_MAP_TOKENS,
  POOLS_BY_CHAIN,
} from '@/constants/tokens'

const StakePage = () => {
  const { chain: connectedChain } = useNetwork()
  const connectedChainId = connectedChain ? connectedChain.id : undefined

  const availableStakingTokens = POOLS_BY_CHAIN[connectedChainId]

  return <LandingPageWrapper>lorem ipsum</LandingPageWrapper>
}

export default StakePage
