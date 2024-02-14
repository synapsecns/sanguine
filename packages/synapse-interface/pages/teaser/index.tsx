import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SynapseIcon } from './SynapseIcon'

const LandingPage = () => {
  const { address: currentAddress } = useAccount()
  const router = useRouter()

  useEffect(() => {
    segmentAnalyticsEvent(`[Teaser] arrives`, {
      address: currentAddress,
      query: router.query,
      pathname: router.pathname,
    })
  }, [])

  return (
    <div className="dark">
      <div className="w-screen h-screen bg-white dark:bg-black overflow-scroll text-black dark:text-white tracking-wide leading-normal">
        <div>
          Ticker
        </div>
        <header className="p-4 flex items-center">
          <div className="text-3xl font-medium flex gap-2 items-center absolute">
            <SynapseIcon width={40} height={40} /><span className="-mt-2">Synapse</span>
          </div>
          <ul className="flex text-lg w-full justify-center">
            <li><a className="px-3 py-2" href="#">About</a></li>
            <li><a className="px-3 py-2" href="#">Bridge</a></li>
            <li><a className="px-3 py-2" href="#">Community</a></li>
            <li><a className="px-3 py-2" href="#">Developers</a></li>
            <li><a className="px-3 py-2" href="#">Explorer</a></li>
          </ul>
        </header>
        <main className="mx-auto w-full max-w-7xl">
          <header className="p-8 text-center max-w-2xl mx-auto">
            <div className="text-6xl font-semibold leading-[1.1] m-4">
              Modular Interchain Messages
            </div>
            <h1 className="m-4 text-2xl font-medium">Synapse 2.0: The Modular Interchain Network</h1>
            <div className="m-8">
              <a className="px-3 py-2 m-2 border border-black dark:border-white rounded" href="#">Bridge</a>
              <a className="px-3 py-2 m-2 border border-black dark:border-white rounded" href="#">Build</a>
            </div>
            <p className="leading-relaxed">
              Say goodbye to centralized resource pools are required for cross-chain communication. Synapse allows you to customize literally every aspect of your interchain communcations.
            </p>
          </header>

        </main>
        <footer>
          Footer
        </footer>
      </div>
    </div>
  )
}

export default LandingPage
