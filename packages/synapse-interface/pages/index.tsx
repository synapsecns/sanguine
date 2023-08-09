import { useState } from 'react'

import StateManagedBridge from './state-managed-bridge'
import { Portfolio } from '@/components/Portfolio/Portfolio'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

const Home = () => {
  const [showBanner, setShowBanner] = useState(true)

  return (
    <LandingPageWrapper>
      <main
        data-test-id="bridge-page"
        className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
      >
        {showBanner && <Banner setShowBanner={setShowBanner} />}
        <div className="flex flex-col-reverse justify-center gap-16 px-4 py-20 mx-auto mt-4 lg:flex-row 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <Portfolio />
          <StateManagedBridge />
        </div>
      </main>
    </LandingPageWrapper>
  )
}

const Banner = ({ setShowBanner }) => {
  return (
    <div className="flex items-center justify-center px-4 mx-auto mt-4 lg:flex-row 2xl:w-3/4 sm:px-8 md:px-12">
      <div
        id="banner-default"
        className="flex items-center w-8/12 p-3 border border-[#AC8FFF] rounded-md text-primaryTextColor"
        role="alert"
        style={{
          background:
            'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
          borderRadius: '10px',
        }}
      >
        <div className="ml-3 text-lg font-thin ">
          New and improved bridge navigation! See all possible routes you can
          bridge assets to.
        </div>
        <button
          type="button"
          className={`
            inline-flex items-center justify-center
            h-7 w-7
            ml-auto -mx-1.5 -my-1.5 p-1.5
            text-primaryTextColor 
          `}
          data-dismiss-target="#banner-default"
          aria-label="Close"
          onClick={() => setShowBanner(false)}
        >
          <svg
            className="w-3 h-3"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 14 14"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
            />
          </svg>
        </button>
      </div>
    </div>
  )
}

export default Home
