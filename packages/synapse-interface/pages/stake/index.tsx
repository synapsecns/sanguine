import { useMemo, useEffect, useState } from 'react'
import { useNetwork, useAccount } from 'wagmi'
import { Token } from '@/utils/types'
import { Chain } from '@/utils/types'
import { getNetworkTextColor } from '@/styles/chains'
import { STAKABLE_TOKENS } from '@/constants/tokens'
import { CHAINS_BY_ID, ChainsByChainID } from '@/constants/chains'
import Grid from '@/components/ui/tailwind/Grid'
import { PageHeader } from '@/components/PageHeader'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import StakeCard from './StakeCard'
import NoStakeCard from './NoStakeCard'

const StakePage = () => {
  const { chain: connectedChain } = useNetwork()
  const { address: currentAddress } = useAccount()
  const [connectedChainId, setConnectedChainId] = useState<number>(undefined)
  const [address, setAddress] = useState(undefined)
  const [isClient, setIsClient] = useState<boolean>(false)
  const [columns, setColumns] = useState<number>(1)

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
    setAddress(currentAddress)
  }, [currentAddress])

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
          px-4 py-16
          md:px-20 md:py-3 md:m-14
        `}
      >
        <div className="flex flex-col justify-center max-w-[1300px] m-auto">
          <PageHeader title="Stake" subtitle="Stake your LP Tokens." />

          <Grid cols={{ xs: 1, sm: 1, md: columns }} gap={6} className="mt-8">
            {isClient && availableStakingTokens.length > 0 ? (
              availableStakingTokens.map((token, key) => {
                if (token.notStake) {
                  return null
                }
                return (
                  <StakeCard
                    key={key}
                    address={currentAddress}
                    chainId={connectedChainId}
                    pool={token}
                  />
                )
              })
            ) : (
              <NoStakeCard chain={connectedChainInfo} />
            )}
          </Grid>
        </div>
      </main>
    </LandingPageWrapper>
  )
}

export default StakePage
