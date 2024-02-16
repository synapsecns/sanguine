import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SynapseIcon } from './SynapseIcon'
import Footer from './Footer'
import Header from './Header'
import Ticker from './Ticker'

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

  /*
  https://stackoverflow.com/questions/61117608/how-do-i-set-system-preference-dark-mode-in-a-react-app-but-also-allow-users-to
  */
  const prefersColorScheme = localStorage.getItem('prefers-color-scheme')

  const windowPrefersDark = window.matchMedia('(prefers-color-scheme: dark)')

  const [prefersDark, setPrefersDark] = useState(
    prefersColorScheme
      ? prefersColorScheme === 'dark'
      : windowPrefersDark.matches
  )

  function selectPrefersDark(e) {
    switch(e.target.value) {
      case 'Dark mode':
        localStorage.setItem('prefers-color-scheme', 'dark')
        setPrefersDark(true)
        break
      case 'Light mode':
        localStorage.setItem('prefers-color-scheme', 'light')
        setPrefersDark(false)
        break
      default:
        localStorage.removeItem('prefers-color-scheme')
        setPrefersDark(windowPrefersDark.matches)
    }
  }

  windowPrefersDark.addEventListener("change", (e) => !prefersColorScheme && setPrefersDark(e.matches))

  function touchStartHandler(e) {
    const style = e.target.nextSibling.style
    console.log('touch', style.display === 'block')

    if (style.display === 'block') return

    style.display = 'block'
    document.addEventListener('click', () => style.display = 'none', { once: true, capture: true })
  }

  return (
    <div className={`${prefersDark && 'dark'}`}>
      <div className="w-screen h-screen bg-white bg-gradient-to-b from-white to-[hsl(235deg_75%_96%)] dark:from-black dark:to-[hsl(265deg_25%_7.5%)] overflow-scroll text-black dark:text-white tracking-wide leading-normal">
        <Ticker />
        {/* Page Header */}
        <nav className="mt-12 md:mt-16 px-2 sm:px-4 md:px-8 grid grid-cols-[auto_auto] md:grid-cols-[1fr_auto_1fr] gap-y-4 md:gap-y-6 items-center max-w-7xl relative mx-auto">
          <a href="#" className="text-2xl md:text-3xl font-medium flex gap-2 items-center">
            <SynapseIcon width={40} height={40} />
            <span className="-mt-1.5">Synapse</span>
          </a>
          <Header />
          <select className="bg-white dark:bg-black text-sm text-inherit cursor-pointer rounded border-zinc-200 dark:border-zinc-800 justify-self-end w-min hover:border-zinc-300 hover:dark:bg-zinc-950 hover:dark:border-zinc-700 col-end-4" onChange={selectPrefersDark}>
            <option selected={prefersColorScheme === 'dark'}>Dark mode</option>
            <option selected={prefersColorScheme === 'light'}>Light mode</option>
            <option selected={!prefersColorScheme}>System {windowPrefersDark.matches ? 'dark' : 'light'}</option>
          </select>
        </nav>
        {/* Hero */}
        <main className="px-2 xs:px-8 grid place-items-center w-full max-w-5xl mx-auto">
          <header className="my-2 md:my-8 lg:my-16 text-center max-w-3xl grid place-items-center">
            <div className="hidden md:block text-3xl md:text-6xl font-semibold my-4">
              Modular Interchain Messages
            </div>
            <h1 className="my-4 max-w-xl text-3xl md:text-2xl font-medium">Synapse 2.0:<wbr/> The Modular Interchain Network</h1>
            <div className="m-2">
              <a className="px-4 pt-2 pb-2.5 m-2 border border-black dark:border-white rounded inline-block" href="#">Bridge</a>
              <a className="px-4 pt-2 pb-2.5 m-2 border border-black dark:border-white rounded inline-block" href="#">Build</a>
            </div>
            <p className="leading-relaxed max-w-xl m-2">
              Say goodbye to centralized resource pools for cross-chain communication. Synapse lets you customize literally every aspect of your interchain communcations.
            </p>
            <ul className="w-fit md:w-max grid grid-cols-2 md:flex text-base sm:text-lg text-center items-center place-center bg-gradient-to-b from-white to-slate-100 dark:from-zinc-900 dark:to-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md px-6 py-2 gap-x-8 gap-y-4 shadow-sm my-4">
              <li className="-mt-1 p-1">50 blockchains</li>
              <li className="-mt-1 p-1">50,000 validators</li>
              <li className="-mt-1 p-1">10.2B messages</li>
              <li className="-mt-1 p-1">$1.2B transferred</li>
            </ul>
          </header>

          <article className="grid gap-8 px-4">
            <section className="grid md:grid-cols-2 gap-x-12">
              <div>
                <h2 className="text-4xl font-medium my-4">
                  Securely connect every blockchain
                </h2>
                <p className="text-lg leading-relaxed">
                  Synapse is comprised of a cross-chain messaging framework and an economically secure method to reach consensus on the validity of cross-chain transactions, enabling developers to build truly native cross-chain apps.
                </p>
              </div>
              <div className="w-full h-full grid items-center justify-center">
                <svg width="100%" height="50%" viewBox="-4 -4 8 8" stroke-width=".1" vector-effect="non-scaling-size" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
                  <circle r="4"/>
                </svg>
              </div>
            </section>
            <section className="grid md:grid-cols-2 gap-x-12">
              <div>
                <h2 className="text-4xl font-medium my-4">
                  Powering the most popular bridge
                </h2>
                <p className="text-lg leading-relaxed">
                Synapse Bridge is built on top of the cross-chain infrastructure  enabling users to seamlessly transfer assets across all blockchains. The Bridge has become the most widely-used method to move assets  cross-chain, offering low cost, fast, and secure bridging.
                </p>
              </div>
              <div className="w-full h-full grid items-center justify-center md:row-start-1 col-start-1 row-start-2">
                <svg width="100%" height="50%" viewBox="-4 -4 8 8" stroke-width=".1" vector-effect="non-scaling-size" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
                  <circle r="4"/>
                </svg>
              </div>
            </section>
            <section className="grid md:grid-cols-2 gap-x-12">
              <div>
                <h2 className="text-4xl font-medium my-4">
                  Battle-tested infrastructure
                </h2>
                <p className="text-lg leading-relaxed">
                  Synapse has processed millions of transactions and tens of billions in bridged assets.
                </p>
              </div>
              <div className="w-full h-full grid items-center justify-center">
                <svg width="100%" height="50%" viewBox="-4 -4 8 8" stroke-width=".1" vector-effect="non-scaling-size" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
                  <circle r="4"/>
                </svg>
              </div>
            </section>
            <section className="grid md:grid-cols-2 gap-x-12">
              <div>
                <h2 className="text-4xl font-medium my-4">
                  Widely integrated
                </h2>
                <p className="text-lg leading-relaxed">
                  Synapse is widely integrated across the most-used Layer 1 and Layer 2 networks for a seamless cross-chain experience.
                </p>
              </div>
              <div className="w-full h-full grid items-center justify-center md:row-start-1 col-start-1 row-start-2">
                <svg width="100%" height="50%" viewBox="-4 -4 8 8" stroke-width=".1" vector-effect="non-scaling-size" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
                  <circle r="4"/>
                </svg>
              </div>
            </section>
          </article>
        </main>
        <Footer />
      </div>
    </div>
  )
}

export default LandingPage
