import { useMemo, useEffect, useState } from 'react'
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
import { CHAINS_BY_ID, ChainsByChainID } from '@/constants/chains'
import { PageHeader } from '@/components/PageHeader'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import StakeCard from './StakeCard'
import NoStakeCard from './NoStakeCard'
import {
  getSynPrices,
  getEthPrice,
  getAvaxPrice,
} from '@/utils/actions/getPrices'
const StakePage = () => {
  const [isClient, setIsClient] = useState<boolean>(false)
  const { chain: connectedChain } = useNetwork()
  const [columns, setColumns] = useState<number>(1)
  const [connectedChainId, setConnectedChainId] = useState<number>()

  const [synPrices, setSynPrices] = useState(undefined)
  const [ethPrice, setEthPrice] = useState(undefined)
  const [avaxPrice, setAvaxPrice] = useState(undefined)

  // Prices to reduce number of calls
  useEffect(() => {
    getSynPrices()
      .then((res) => {
        setSynPrices(res)
      })
      .catch((err) => console.log('Could not get syn prices', err))
    getEthPrice()
      .then((res) => {
        setEthPrice(res)
      })
      .catch((err) => console.log('Could not get eth prices', err))
    getAvaxPrice()
      .then((res) => {
        setAvaxPrice(res)
      })
      .catch((err) => console.log('Could not get avax prices', err))
  }, [connectedChainId])

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
    setConnectedChainId(connectedChain.id)
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
