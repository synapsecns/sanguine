import { useAccount, useNetwork } from 'wagmi'
import Grid from '@tw/Grid'
import SwapCard from './SwapCard'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { PageHeader } from '@components/PageHeader'
import { SWAPABLE_TOKENS } from '@constants/tokens'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import NoSwapCard from './NoSwapCard'
import { useEffect, useState, useMemo } from 'react'

const SwapPage = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const [connectedChain, setConnectedChain] = useState(0)
  useEffect(() => {
    setConnectedChain(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])
  return (
    <LandingPageWrapper>
      <div>
        <Grid
          cols={{ xs: 1 }}
          gap={6}
          className="justify-center px-2 py-16 sm:px-6 md:px-8"
        >
          <div className="pb-3 place-self-center">
            <div className="flex justify-between mb-5 ml-5 mr-5">
              <PageHeader
                title="Swap"
                subtitle="Exchange stablecoins on-chain."
              />
            </div>
            {SWAPABLE_TOKENS[connectedChain]?.length > 0 ? (
              <SwapCard address={address} connectedChainId={connectedChain} />
            ) : (
              <NoSwapCard chainId={connectedChain} />
            )}
          </div>
        </Grid>
      </div>
    </LandingPageWrapper>
  )
}

export default SwapPage
