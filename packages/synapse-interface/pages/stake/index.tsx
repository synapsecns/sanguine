import { useMemo, useEffect, useState } from 'react'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useNetwork } from 'wagmi'
import { Token } from '@/utils/types'
import { Chain } from '@/utils/types'
import { getNetworkTextColor } from '@/styles/chains'
import { STAKABLE_TOKENS } from '@/constants/tokens'
import { CHAINS_BY_ID, ChainsByChainID } from '@/constants/chains'
import { PageHeader } from '@/components/PageHeader'
import Grid from '@/components/ui/tailwind/Grid'
import StakeCard from './StakeCard'
import NoStakeCard from './NoStakeCard'

const StakePage = () => {
  const [isClient, setIsClient] = useState<boolean>(false)
  const { chain: connectedChain } = useNetwork()
  const [columns, setColumns] = useState<number>(1)
  const [connectedChainId, setConnectedChainId] = useState<number>(undefined)

  const connectedChainInfo: Chain | undefined = useMemo(() => {
    if (connectedChainId) {
      const chainMapping: ChainsByChainID = CHAINS_BY_ID
      return chainMapping[connectedChainId]
    } else {
      return undefined
    }
  }, [connectedChainId])

  const availableStakingTokens: Token[] | [] =
    STAKABLE_TOKENS[connectedChainId] ?? []

  useEffect(() => {
    const isSingle = availableStakingTokens.length < 2
    setColumns(isSingle ? 1 : 2)
  }, [availableStakingTokens])

  useEffect(() => {
    setConnectedChainId(connectedChain && connectedChain.id)
  }, [connectedChain])

  useEffect(() => {
    setIsClient(true)
  }, [])

  return (
    <LandingPageWrapper>
      <main
        data-test-id="stake-page"
        className={`
          flex flex-col justify-between
          px-20 m-14 space-x-2
        `}
      >
        <PageHeader title="Stake" subtitle="Stake your LP Tokens." />

        <Grid cols={{ xs: 1, sm: 1, md: columns }} gap={6} className="mt-8">
          {isClient && availableStakingTokens.length > 0 ? (
            availableStakingTokens.map((token, key) => (
              <StakeCard key={key} chainId={connectedChainId} token={token} />
            ))
          ) : (
            <NoStakeCard chain={connectedChainInfo} />
          )}
        </Grid>
      </main>
    </LandingPageWrapper>
  )
}

export default StakePage
