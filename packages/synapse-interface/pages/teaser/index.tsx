import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SynapseIcon } from './SynapseIcon'
import { generateTx } from './strings'
import styles from './marquee.module.css'
import Footer from './footer'

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

  // Detect and set 'prefers-color-scheme: dark'
  // https://stackoverflow.com/questions/61117608/how-do-i-set-system-preference-dark-mode-in-a-react-app-but-also-allow-users-to
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
      <div className="w-screen h-screen bg-white bg-gradient-to-b from-white to-slate-100 dark:from-black dark:to-zinc-950 overflow-scroll text-black dark:text-white tracking-wide leading-normal">
        {/* Ticker – Easter egg: define custom <marquee> element */}
        <div className="absolute w-screen z-10 mb-12 bg-zinc-50 dark:bg-zinc-950 text-sm border-y border-zinc-200 dark:border-zinc-900 overflow-hidden">
          <button className="absolute bg-inherit px-4 py-1.5 border-r border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10">
            <svg width="8" height="8" viewBox="-4 -4 8 8" overflow="visible" className="fill-green-500 stroke-green-500" xmlns="http://www.w3.org/2000/svg">
              <circle r="4">
                <animate
                  attributeName="stroke-width"
                  values="0; 16"
                  dur="1.5s"
                  repeatCount="indefinite"
                />
                <animate
                  attributeName="stroke-opacity"
                  values=".5; 0"
                  dur="1.5s"
                  repeatCount="indefinite"
                />
              </circle>
            </svg>
            Live – All transactions <span className="text-xxs">▼</span>
          </button>
          <ul className={`flex whitespace-nowrap list-disc marker:text-zinc-500 ${styles.marquee}`}>
            {[...Array(20)].map((x, i) => {
              return <li className="pl-2 ml-7"><a href="#" className="text-zinc-500 hover:text-inherit hover:underline py-1.5 inline-block">{generateTx()}</a></li>
            })}
          </ul>
          <a href="#" className="absolute top-0 right-0 bg-inherit px-4 py-1.5 border-l border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10">
            Explorer {'->'}
          </a>
        </div>
        {/* Nav */}
        <nav className="mt-12 p-8 flex items-center max-w-7xl relative mx-auto">
          <a href="#" className="text-3xl font-medium flex gap-2 items-center absolute">
            <SynapseIcon width={40} height={40} />
            <span className="-mt-1.5">Synapse</span>
          </a>
          <ul className="flex text-lg w-full justify-center">
            <li className="group">
              <a
                className="px-3 py-1 hover:bg-zinc-50 hover:dark:bg-zinc-950 border border-transparent hover:border-fuchsia-500 rounded inline-block"
                href="#"
                onTouchStartCapture={touchStartHandler}
              >
                About
              </a>
              <div className="hidden group-hover:block absolute p-2">
                <ul className="bg-zinc-50 dark:bg-zinc-950 rounded text-base -ml-2 border border-zinc-200 dark:border-zinc-800 shadow-sm">
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Vision</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Philosophy</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Roadmap</a></li>
                </ul>
              </div>
            </li>
            <li className="group">
              <a className="px-3 py-1 hover:bg-zinc-50 hover:dark:bg-zinc-950 border border-transparent hover:border-fuchsia-500 rounded inline-block" href="#">Bridge</a>
              <div className="hidden group-hover:block absolute p-2">
                <ul className="bg-zinc-50 dark:bg-zinc-950 rounded text-base -ml-2 border border-zinc-200 dark:border-zinc-800 shadow-sm">
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Synapse Bridge</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">On-chain swap</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Solana Bridge</a></li>
                </ul>
              </div>
            </li>
            <li className="group">
              <a className="px-3 py-1 hover:bg-zinc-50 hover:dark:bg-zinc-950 border border-transparent hover:border-fuchsia-500 rounded inline-block" href="#">Community</a>
              <div className="hidden group-hover:block absolute p-2">
                <ul className="bg-zinc-50 dark:bg-zinc-950 rounded text-base -ml-2 border border-zinc-200 dark:border-zinc-800 shadow-sm">
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Discord</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Telegram</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Twitter</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Blog</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Forum</a></li>
                </ul>
              </div>
            </li>
            <li className="group">
              <a className="px-3 py-1 hover:bg-zinc-50 hover:dark:bg-zinc-950 border border-transparent hover:border-fuchsia-500 rounded inline-block" href="#">Developers</a>
              <div className="hidden group-hover:block absolute p-2">
                <ul className="bg-zinc-50 dark:bg-zinc-950 rounded text-base -ml-2 border border-zinc-200 dark:border-zinc-800 shadow-sm">
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Docs</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Github</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Synapse CNS</a></li>
                  <li><a className="px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded" href="#">Interchain Network</a></li>
                </ul>
              </div>
            </li>
            <li className="group">
              <a className="px-3 py-1 hover:bg-zinc-50 hover:dark:bg-zinc-950 border border-transparent hover:border-fuchsia-500 rounded inline-block" href="#">Explorer</a>
            </li>
          </ul>
          <form className="absolute right-0 flex gap-2 p-4 items-center">
            <select className="bg-white dark:bg-black text-sm text-inherit cursor-pointer rounded border-zinc-200 dark:border-zinc-800" onChange={selectPrefersDark}>
              <option selected={prefersColorScheme === 'dark'}>Dark mode</option>
              <option selected={prefersColorScheme === 'light'}>Light mode</option>
              <option selected={!prefersColorScheme}>System {windowPrefersDark.matches ? 'dark' : 'light'}</option>
            </select>
          </form>
        </nav>
        {/* Hero */}
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
            <p className="leading-relaxed text-md">
              Say goodbye to centralized resource pools for cross-chain communication. Synapse lets you customize literally every aspect of your interchain communcations.
            </p>
          </header>
          <ul className="w-fit flex text-lg justify-center bg-gradient-to-b from-white to-slate-100 dark:from-zinc-900 dark:to-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md px-6 py-2 gap-8 mx-auto shadow-sm">
            <li>50 blockchains</li>
            <li>50,000 validators</li>
            <li>10.2B messages</li>
            <li>$1.2B transferred</li>
          </ul>
          <article className="p-8 grid grid-cols-2 place-items-center gap-16">
            <section className="py-8">
              <h2 className="text-4xl font-medium my-4">
                Securely connect every blockchain
              </h2>
              <p className="text-lg leading-relaxed">
                Synapse is comprised of a cross-chain messaging framework and an economically secure method to reach consensus on the validity of cross-chain transactions, enabling developers to build truly native cross-chain apps.
              </p>
            </section>
            <svg width="8" height="8" viewBox="-4 -4 8 8" stroke-width="1.5" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
              <circle r="72"/>
            </svg>
            <svg width="8" height="8" viewBox="-4 -4 8 8" stroke-width="1.5" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
              <circle r="72"/>
            </svg>
            <section className="py-8">
              <h2 className="text-4xl font-medium my-4">
                Powering the most popular bridge
              </h2>
              <p className="text-lg">
                Synapse Bridge is built on top of the cross-chain infrastructure  enabling users to seamlessly transfer assets across all blockchains. The Bridge has become the most widely-used method to move assets  cross-chain, offering low cost, fast, and secure bridging.
              </p>
            </section>
            <section className="py-8">
              <h2 className="text-4xl font-medium my-4">
                Battle-tested infrastructure
              </h2>
              <p className="text-lg">
                Synapse has processed millions of transactions and tens of billions in bridged assets.
              </p>
            </section>
            <svg width="8" height="8" viewBox="-4 -4 8 8" stroke-width="1.5" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
              <circle r="72"/>
            </svg>
            <svg width="8" height="8" viewBox="-4 -4 8 8" stroke-width="1.5" fill="none" overflow="visible" className="stroke-fuchsia-500" xmlns="http://www.w3.org/2000/svg">
              <circle r="72"/>
            </svg>
            <section className="py-8">
              <h2 className="text-4xl font-medium my-4">
                Widely integrated
              </h2>
              <p className="text-lg">
                Synapse is widely integrated across the most-used Layer 1 and Layer 2 networks for a seamless cross-chain experience.
              </p>
            </section>
          </article>
        </main>
        <Footer />
      </div>
    </div>
  )
}

export default LandingPage
