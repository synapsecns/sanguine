import { useEffect, useState } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { PageHeader } from '@components/PageHeader'
import { SWAPABLE_TOKENS } from '@constants/tokens'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'
import Grid from '@tw/Grid'
import SwapCard from './SwapCard'
import NoSwapCard from './NoSwapCard'

const SwapPage = () => {
  const { address: currentAddress } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState<number>(0)
  const [address, setAddress] = useState<`0x${string}` | undefined>(undefined)

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
              {SWAPABLE_TOKENS[connectedChainId]?.length > 0 ? (
                <SwapCard
                  address={address}
                  connectedChainId={connectedChainId}
                />
              ) : (
                <NoSwapCard chainId={connectedChainId} />
              )}
            </div>
          </Grid>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default SwapPage
