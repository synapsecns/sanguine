import { useEffect } from 'react'
import { Banner } from '@/components/Banner'
import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import ReactGA from 'react-ga'
import useSyncQueryParamsWithBridgeState from '@/utils/hooks/useSyncQueryParamsWithBridgeState'
import { hotjar } from 'react-hotjar'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = 'G-BBC13LQXBD'
ReactGA.initialize(TRACKING_ID)

const Home = () => {
  useSyncQueryParamsWithBridgeState()

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
