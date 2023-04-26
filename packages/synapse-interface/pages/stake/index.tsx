import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useNetwork } from 'wagmi'

const StakePage = () => {
  const { chain: connectedChain } = useNetwork()
  const connectedChainId = connectedChain ? connectedChain.id : undefined

  return <LandingPageWrapper>lorem ipsum</LandingPageWrapper>
}

export default StakePage
