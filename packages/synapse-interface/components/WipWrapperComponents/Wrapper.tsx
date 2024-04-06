import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { SynapseAnchor } from './SynapseLogo'
import Footer from './Footer'
import NavMenu from './NavMenu'
import Ticker from './Ticker'
import CortexCli from './CortexCli'

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
        prefersDark
          ? 'dark text-zinc-200 from-black dark:to-[hsl(265deg_25%_7.5%)]'
          : 'text-zinc-800 from-white to-[hsl(235deg_75%_96%)]'
      } bg-gradient-to-b tracking-wide`}
    >
      <Ticker />
      <NavMenu />
      <main className="overflow-hidden">{children}</main>
      <CortexCli setPrefersDark={setPrefersDark} />
      <Footer />
    </div>
  )
}

export default Wrapper
