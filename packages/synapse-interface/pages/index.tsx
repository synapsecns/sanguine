import { Banner, InterruptedServiceBanner } from '@/components/Banner'
import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import ReactGA from 'react-ga'
import useSyncQueryParamsWithBridgeState from '@/utils/hooks/useSyncQueryParamsWithBridgeState'
import { AnnouncementBanner } from '@/components/Maintenance/AnnouncementBanner'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = 'G-BBC13LQXBD'
ReactGA.initialize(TRACKING_ID)

const Home = () => {
  useSyncQueryParamsWithBridgeState()

  const startDate = new Date(Date.UTC(2024, 2, 13, 1, 50, 0))
  const endDate = new Date(Date.UTC(2024, 2, 13, 1, 60, 0))

  return (
    <LandingPageWrapper>
      <main
        data-test-id="bridge-page"
        className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
      >
        {/* <Banner /> */}
        {/* <InterruptedServiceBanner /> */}
        <AnnouncementBanner
          bannerId="03122024-eth-dencun"
          bannerContents={
            <div>
              The Bridge + RFQ will be globally offline 15mins ahead of the
              Dencun upgrade (March 13, 13:55 UTC, 9:55 EST). Will be back
              online about 15 - 30 mins after Dencun.
            </div>
          }
          startDate={startDate}
          endDate={endDate}
        />
        <div className="flex flex-col-reverse justify-center gap-16 px-4 py-20 mx-auto lg:flex-row 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <Portfolio />
          <StateManagedBridge />
        </div>
      </main>
    </LandingPageWrapper>
  )
}

export default Home
