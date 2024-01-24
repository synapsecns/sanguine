import { Banner } from '@/components/Banner'
import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import ReactGA from 'react-ga'
import useSyncQueryParamsWithBridgeState from '@/utils/hooks/useSyncQueryParamsWithBridgeState'

import { useEffect } from 'react'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'
import { arbitrum } from 'viem/chains'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = 'G-BBC13LQXBD'
ReactGA.initialize(TRACKING_ID)

const Home = () => {
  useSyncQueryParamsWithBridgeState()

  useEffect(() => {
    ;(async () => {
      const data = await getErc20TokenTransfers(
        '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
        '0xF080B794AbF6BB905F2330d25DF545914e6027F8',
        '0x81EF4608B796265F1e3695cE00FdCfC8aA5933Dd',
        173545720n,
        arbitrum
      )

      console.log('data:', data)
    })()
  }, [])

  return (
    <LandingPageWrapper>
      <main
        data-test-id="bridge-page"
        className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
      >
        <Banner />
        <div className="flex flex-col-reverse justify-center gap-16 px-4 py-20 mx-auto lg:flex-row 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <Portfolio />
          <StateManagedBridge />
        </div>
      </main>
    </LandingPageWrapper>
  )
}

export default Home
