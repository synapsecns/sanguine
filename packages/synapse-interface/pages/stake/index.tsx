import { useMemo } from 'react'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { useNetwork } from 'wagmi'
import { Token } from '@/utils/types'
import { Chain } from '@/utils/types'
import { getNetworkTextColor } from '@/styles/chains'
import {
  STAKABLE_TOKENS,
  STAKING_MAP_TOKENS,
  POOLS_BY_CHAIN,
} from '@/constants/tokens'
import { CHAINS_BY_ID } from '@/constants/chains'
import { PageHeader } from '@/components/PageHeader'
import Card from '@/components/ui/tailwind/Card'
import Grid from '@/components/ui/tailwind/Grid'

function NoStakeCard({ chain }: { chain?: Chain }) {
  const chainName = chain?.chainName ?? 'current network'
  const networkColor = chain?.color
  return (
    <Card
      divider={false}
      className={`
        transform transition-all duration-100
        rounded-xl max-w-[420px]
      `}
    >
      <div className="w-full pt-4 text-center text-gray-400">
        No stakes available on{' '}
        <span className={`${getNetworkTextColor(networkColor)} font-medium`}>
          {chainName}
        </span>
      </div>
    </Card>
  )
}

const StakePage = () => {
  const { chain: connectedChain } = useNetwork()

  const connectedChainId: number | undefined = connectedChain
    ? Number(connectedChain.id)
    : undefined

  const connectedChainInfo: Chain | undefined = useMemo(() => {
    if (connectedChainId) {
      const chainMapping = CHAINS_BY_ID
      return chainMapping[connectedChainId]
    } else {
      return undefined
    }
  }, [connectedChainId])

  const availableStakingTokens: Token[] | [] =
    POOLS_BY_CHAIN[connectedChainId] ?? []

  // const availableStakingTokens = []
  // console.log('CHAINS_BY_ID: ', CHAINS_BY_ID)
  // console.log('connectedChainInfo: ', connectedChainInfo)

  const gridColumns: number = availableStakingTokens.length > 1 ? 2 : 1

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
        <Grid cols={gridColumns} className="mt-8">
          {availableStakingTokens.length > 0 ? (
            <></>
          ) : (
            <NoStakeCard chain={connectedChainInfo} />
          )}
        </Grid>
      </main>
    </LandingPageWrapper>
  )
}

export default StakePage
