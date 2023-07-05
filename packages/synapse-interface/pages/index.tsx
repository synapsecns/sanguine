import StateManagedBridge from './state-managed-bridge'
import BridgePage from './bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

const Home = () => {
  return (
    <LandingPageWrapper>
      <main
        data-test-id="bridge-page"
        className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
      >
        <div className="flex flex-col justify-center px-4 py-20 mx-auto mt-4 lg:flex-row 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <Portfolio />
          <StateManagedBridge />
        </div>
      </main>
    </LandingPageWrapper>
  )
}

export default Home
