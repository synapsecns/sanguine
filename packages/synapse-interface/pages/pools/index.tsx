import { useEffect, useState, useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'

import {
  DISPLAY_POOLS_BY_CHAIN,
  USD_POOLS_BY_CHAIN,
  ETH_POOLS_BY_CHAIN,
  LEGACY_POOLS_BY_CHAIN,
} from '@constants/tokens'

import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'

import PoolCards from './PoolCards'

const PoolsPage = () => {
  const { address: currentAddress } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)
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
        {/* <HarmonyCheck fromChainId={chainId} toChainId={chainId} /> */}
        <PoolCards
          address={address}
          arr={[
            {
              index: 0,
              label: 'All',
              poolsByChain: DISPLAY_POOLS_BY_CHAIN,
            },
            {
              index: 1,
              label: 'USD',
              poolsByChain: USD_POOLS_BY_CHAIN,
            },
            {
              index: 2,
              label: 'ETH',
              poolsByChain: ETH_POOLS_BY_CHAIN,
            },
            {
              index: 3,
              label: 'Legacy',
              poolsByChain: LEGACY_POOLS_BY_CHAIN,
            },
          ]}
          connectedChainId={connectedChainId}
        />
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default PoolsPage
