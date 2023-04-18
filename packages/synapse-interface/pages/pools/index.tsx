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

import PoolTabs from './PoolTabs'

export default function PoolsPage() {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [tabIndex, setTabIndex] = useState(0)

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])
  console.log('DISPLAY_POOLS_BY_CHAIN', DISPLAY_POOLS_BY_CHAIN)
  console.log('USD_POOLS_BY_CHAIN', USD_POOLS_BY_CHAIN)
  console.log('ETH_POOLS_BY_CHAIN', ETH_POOLS_BY_CHAIN)
  console.log('LEGACY_POOLS_BY_CHAIN', LEGACY_POOLS_BY_CHAIN)
  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        {/* <HarmonyCheck fromChainId={chainId} toChainId={chainId} /> */}
        <PoolTabs
          arr={[
            {
              index: 0,
              label: 'All',
              textLabel: '',
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
          tabIndex={tabIndex}
          setTabIndex={setTabIndex}
        />
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}
