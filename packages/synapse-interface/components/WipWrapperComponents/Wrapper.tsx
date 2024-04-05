import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SynapseAnchor } from './SynapseLogo'
import Footer from './Footer'
import NavMenu from './NavMenu'
import Ticker from './Ticker'

const Wrapper = ({ children }) => {
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

  /* TODO: Mobile Support */
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
    <div
      className={`${
        prefersDark ? 'dark text-zinc-200 bg-black' : 'text-zinc-800 bg-white'
      } tracking-wide`}
    >
      <Ticker />
      <div className="bg-white bg-gradient-to-b from-white to-[hsl(235deg_75%_96%)] dark:from-black dark:to-[hsl(265deg_25%_7.5%)]">
        <NavMenu />
        <main className="overflow-hidden">{children}</main>
        <Footer />
        <select
          className="sticky bottom-0 right-0 bg-white dark:bg-black text-sm text-inherit cursor-pointer rounded border-zinc-200 dark:border-zinc-800 justify-self-end w-min hover:border-zinc-300 hover:dark:bg-zinc-950 hover:dark:border-zinc-700 col-end-4"
          onChange={selectPrefersDark}
        >
          <option
            defaultValue={prefersColorScheme === 'dark' ? 'true' : 'false'}
          >
            Dark mode
          </option>
          <option
            defaultValue={prefersColorScheme === 'light' ? 'true' : 'false'}
          >
            Light mode
          </option>
          <option defaultValue={!prefersColorScheme ? 'true' : 'false'}>
            System {windowPrefersDark.matches ? 'dark' : 'light'}
          </option>
        </select>
      </div>
    </div>
  )
}

export default Wrapper
