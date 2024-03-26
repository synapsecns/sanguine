import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import ReactGA from 'react-ga'
import useSyncQueryParamsWithBridgeState from '@/utils/hooks/useSyncQueryParamsWithBridgeState'
import { MaintenanceBanner } from '@/components/Maintenance/Events/template/MaintenanceEvent'
import { AnnouncementBanner } from '@/components/Maintenance/AnnouncementBanner'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = 'G-BBC13LQXBD'
ReactGA.initialize(TRACKING_ID)

const Home = () => {
  useSyncQueryParamsWithBridgeState()

  return (
    <>
      <MaintenanceBanner />
      <LandingPageWrapper>
        <main
          data-test-id="bridge-page"
          className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
        >
          {/* <AnnouncementBanner
            bannerId="2024-03-26-blast-bridge-pause"
            bannerContents="Bridging USDB and WETH on Blast paused."
            startDate={new Date(Date.UTC(2024, 2, 20, 20, 20, 0))}
            endDate={new Date(Date.UTC(2026, 2, 20, 22, 0, 0))}
          /> */}
          <div className="flex flex-col-reverse justify-center gap-16 px-4 py-20 mx-auto lg:flex-row 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
            <Portfolio />
            <StateManagedBridge />
          </div>
        </main>
      </LandingPageWrapper>
    </>
  )
}

export default Home
