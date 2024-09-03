import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import ReactGA from 'react-ga'
import useSyncQueryParamsWithBridgeState from '@/utils/hooks/useSyncQueryParamsWithBridgeState'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = 'G-BBC13LQXBD'
ReactGA.initialize(TRACKING_ID)

export async function getStaticProps(context) {
  return {
    props: {
      // You can get the messages from anywhere you like. The recommended
      // pattern is to put them in JSON files separated by locale and read
      // the desired one based on the `locale` received from Next.js.
      messages: (await import(`../messages/${context.locale}.json`)).default,
    },
  }
}

const Home = () => {
  useSyncQueryParamsWithBridgeState()

  return (
    <>
      <LandingPageWrapper>
        <main
          data-test-id="bridge-page"
          className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
        >
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
