import deepmerge from 'deepmerge'
import ReactGA from 'react-ga'

import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import useSyncQueryParamsWithBridgeState from '@/utils/hooks/useSyncQueryParamsWithBridgeState'

// TODO: someone should add this to the .env, disable if blank, etc.
// this is being added as a hotfix to assess user load on the synapse explorer api
// I'd recommend moving this to a sushi-style analytics provider wrapper.
const TRACKING_ID = 'G-BBC13LQXBD'
ReactGA.initialize(TRACKING_ID)

export async function getStaticProps({ locale }) {
  const userMessages = (await import(`../messages/${locale}.json`)).default
  const defaultMessages = (await import(`../messages/en-US.json`)).default
  const messages = deepmerge(defaultMessages, userMessages)

  return {
    props: {
      messages,
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
          className="relative z-0 flex-1 h-full overflow-y-none focus:outline-none"
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
