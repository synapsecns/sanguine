import { useState } from 'react'

import {
  DISPLAY_POOLS_BY_CHAIN,
  USD_POOLS_BY_CHAIN,
  ETH_POOLS_BY_CHAIN,
  LEGACY_POOLS_BY_CHAIN,
} from '@constants/tokens/poolsByChain'

import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'

import PoolTabs from './PoolTabs'

import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { HarmonyCheck } from '@pages/Bridge'

export default function PoolsPage() {
  const { chainId } = useActiveWeb3React()
  const [tabIndex, setTabIndex] = useState(0)

  return (
    <LandingPageWrapper>
      <StandardPageContainer>
        <HarmonyCheck fromChainId={chainId} toChainId={chainId} />
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
