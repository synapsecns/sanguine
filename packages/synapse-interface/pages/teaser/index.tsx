import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SynapseAnchor } from './SynapseLogo'
import Footer from './LandingFooter'
import NavMenu from './NavMenu'
import Hero from './Hero'
import Ticker from './Ticker'
import ValueProps from './ValueProps'

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
    switch (e.target.value) {
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

  windowPrefersDark.addEventListener(
    'change',
    (e) => !prefersColorScheme && setPrefersDark(e.matches)
  )

  function touchStartHandler(e) {
    const style = e.target.nextSibling.style
    console.log('touch', style.display === 'block')

    if (style.display === 'block') return

    style.display = 'block'
    document.addEventListener('click', () => (style.display = 'none'), {
      once: true,
      capture: true,
    })
  }

  return (
    <div className={prefersDark && 'dark'}>
      <div className="w-screen h-screen bg-white bg-gradient-to-b from-white to-[hsl(235deg_75%_96%)] dark:from-black dark:to-[hsl(265deg_25%_7.5%)] overflow-scroll text-zinc-800 dark:text-zinc-200 tracking-wide">
        <Ticker />
        <nav className="mt-12 md:mt-16 px-2 sm:px-4 md:px-8 grid gap-y-4 md:gap-y-6 items-center max-w-7xl m-auto">
          <SynapseAnchor />
          <NavMenu />
          <select
            className="bg-white dark:bg-black text-sm text-inherit cursor-pointer rounded border-zinc-200 dark:border-zinc-800 justify-self-end w-min hover:border-zinc-300 hover:dark:bg-zinc-950 hover:dark:border-zinc-700 col-end-4"
            onChange={selectPrefersDark}
          >
            <option selected={prefersColorScheme === 'dark'}>Dark mode</option>
            <option selected={prefersColorScheme === 'light'}>
              Light mode
            </option>
            <option selected={!prefersColorScheme}>
              System {windowPrefersDark.matches ? 'dark' : 'light'}
            </option>
          </select>
        </nav>
        {/* Hero */}
        <main className="px-2 xs:px-8 grid place-items-center max-w-5xl mx-auto">
          <Hero />
          <ValueProps />
        </main>
        <Footer />
      </div>
    </div>
  )
}

export default LandingPage
