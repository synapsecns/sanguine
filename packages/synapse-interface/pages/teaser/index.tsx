import { useRouter } from 'next/router'
import { useEffect, useRef, useState } from 'react'
import { useAccount } from 'wagmi'

import { CHAINS_ARR } from '@/constants/chains'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import Hero from './Hero'

import Wrapper from '@/components/WipWrapperComponents/Wrapper'

import FauxBridge from './FauxBridge'
import PlatformAnimation from './HeroAnimation'
import { Chain } from '@/utils/types'
import HeroAnimation from './HeroAnimation'

const chainTagClassName = 'pl-2 pr-6 py-2 border-l h-min'

const LandingPage = () => {
  const { address: currentAddress } = useAccount()
  const router = useRouter()

  // const [platforms, setPlatforms] = useState([])

  const heroRef = useRef(null)
  const chainRef = useRef(null)

  useEffect(() => {
    segmentAnalyticsEvent(`[Teaser] arrives`, {
      address: currentAddress,
      query: router.query,
      pathname: router.pathname,
    })
  }, [])

  useEffect(() => {
    handleResize()

    window.addEventListener('resize', handleResize)
    return () => window.removeEventListener('resize', handleResize)
  })

  function handleResize() {
    const chains = chainRef.current as HTMLElement
    const tags: HTMLAnchorElement[] = Array.from(
      chains.querySelectorAll<HTMLAnchorElement>('a')
    ).slice(1, 5)

    heroRef.current.querySelectorAll("[id^='platform']").forEach((a, i) => {
      const { x, y, width, height } = a.parentElement.getBoundingClientRect()
      const w = tags[i].offsetWidth
      const h = tags[i].offsetHeight
      // const { left, top } = tags[i].getBoundingClientRect()
      let left, top
      switch (i) {
        case 0:
          left = x + width * 0.5
          top = y - h
          break
        case 1:
          left = x + width * 0.5
          top = y - h
          break
        case 2:
          left = x + width * 0.5
          top = y + height
          break
        case 3:
          left = x + width / 2
          top = y + height
          // top = y + height * 0.625
          break
      }
      if (width) {
        tags[i].style.position = `absolute`
        tags[i].style.left = `${left + window.scrollX}px`
        tags[i].style.top = `${top + window.scrollY}px`
      } else {
        // tags[i].style.display = 'none'
        // tags[i].style.left = `${left}px`
        // tags[i].style.top = `${top}px`
      }
    })
  }

  function ChainTag({
    chain,
    index,
  }: {
    chain: Chain
    index: number
  }): React.ReactNode {
    return (
      <a
        href="#"
        className={`grid grid-cols-[auto_auto] gap-x-2 items-center border-zinc-800 bg-zinc-900/25 backdrop-blur-sm  ${chainTagClassName}`}
        key={chain.name} // TODO: Solve 'unique key' warning
      >
        <img
          src={chain.chainImg.src}
          width="24"
          height="24"
          className="float-left row-span-2"
          alt={chain.name}
        />
        <dt>{chain.name}</dt>
        <dd>$417.83M</dd>
      </a>
    )
  }

  return (
    <Wrapper>
      {/* <Hero /> */}
      <section className="w-full px-4">
        <div className="hidden xs:block text-5xl sm:text-6xl font-semibold text-center mt-8 lg:mt-12 mb-4">
          Reach every chain.
        </div>
        <h1 className="text-3xl sm:text-2xl font-medium text-center mt-6 mb-4">
          The Web connects devices. We connect blockchains.
        </h1>
        <div className="flex gap-4 text-base sm:text-lg whitespace-nowrap justify-center mt-6">
          <a href="#" className="px-6 py-2 border border-fuchsia-500 rounded">
            Build
          </a>
          <a href="#" className="px-6 py-2 border border-zinc-800 rounded">
            Bridge
          </a>
        </div>
      </section>
      <section className="grid justify-center">
        <HeroAnimation
          heroRef={heroRef}
          className="w-full min-w-[480px] xs:min-w-[640px]"
        />
      </section>
      <section id="chain-list" className="mb-12 justify-center">
        <dl
          ref={chainRef}
          className="flex text-sm border-y border-zinc-800 whitespace-nowrap justify-center"
        >
          <a href="#" className={`${chainTagClassName} border-fuchsia-500`}>
            <dt>6 month vol.</dt>
            <dd>$7.03B</dd>
          </a>
          {CHAINS_ARR.slice(0, 10).map((a, i) => (
            <ChainTag chain={a} index={i} />
          ))}
          <dt className={`${chainTagClassName} border-zinc-800`}>
            DFK, Klaytn,
            <br />
            Cronos & more
          </dt>
        </dl>
      </section>
      <section
        id="entry-points"
        className="text-center max-w-screen-xl mx-auto mt-24"
      >
        <h2 className="text-3xl xs:text-5xl md:text-6xl font-medium px-4">
          Start here
        </h2>
        <p className="max-w-lg mx-auto px-4 mt-6 mb-8">
          Cortex brings the federated blockchain universe together. Access new
          technologies, trade, loan, vote, and stake with any project from any
          chain.
        </p>
        {/* <h3 className="text-2xl font-medium my-6">The future is interchain.</h3> */}
        <svg
          viewBox="0 -1 4 2"
          className="stroke-zinc-700 w-full h-16 block md:hidden"
          preserveAspectRatio="none"
          overflow="visible"
          strokeLinecap="round"
        >
          <path d="M2 0 v-1" vectorEffect="non-scaling-stroke" />
          <path d="M1 0 h2" vectorEffect="non-scaling-stroke" />
          <path d="M1 0 v1" vectorEffect="non-scaling-stroke" />
          <path d="M3 0 v1" vectorEffect="non-scaling-stroke" />
          <g strokeWidth="6">
            <path d="m2 -1 v.0001" vectorEffect="non-scaling-stroke" />
            <path d="m1 1 v.0001" vectorEffect="non-scaling-stroke" />
            <path d="m3 1 v.0001" vectorEffect="non-scaling-stroke" />
          </g>
        </svg>
        <svg
          viewBox="0 -1 8 2"
          className="stroke-zinc-700 w-full h-16 hidden md:block"
          preserveAspectRatio="none"
          overflow="visible"
          strokeLinecap="round"
        >
          <path d="M1 0 h6" vectorEffect="non-scaling-stroke" />
          <path d="M1 0 v1" vectorEffect="non-scaling-stroke" />
          <path d="M3 0 v1" vectorEffect="non-scaling-stroke" />
          <path d="M5 0 v1" vectorEffect="non-scaling-stroke" />
          <path d="M7 0 v1" vectorEffect="non-scaling-stroke" />
          <path d="M4 0 v-1" vectorEffect="non-scaling-stroke" />
          <g strokeWidth="6">
            <path d="m4 -1 v.0001" vectorEffect="non-scaling-stroke" />
            <path d="m1 1 v.0001" vectorEffect="non-scaling-stroke" />
            <path d="m3 1 v.0001" vectorEffect="non-scaling-stroke" />
            <path d="m5 1 v.0001" vectorEffect="non-scaling-stroke" />
            <path d="m7 1 v.0001" vectorEffect="non-scaling-stroke" />
          </g>
        </svg>
        <div className="grid grid-cols-2 md:grid-cols-4 text-left md:text-center items-center md:items-start">
          <div className="pl-8 py-4 md:p-4 flex-1 col-start-1 md:col-start-auto md:row-start-1">
            <h3 className="text-2xl py-2">Instructions provided</h3>
            <p>Make your project interchain today.</p>
          </div>
          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>
          <div className="pl-8 md:p-4 flex-1 col-start-1 md:col-start-auto md:row-start-1">
            <h3 className="text-2xl py-2">Batteries included</h3>
            <p>Production-ready customizable interchain contracts.</p>
          </div>
          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>
          <div className="pl-8 md:p-4 flex-1 col-start-1 md:col-start-auto md:row-start-1">
            <h3 className="text-2xl py-2">Pre-assembled</h3>
            <p>
              Add the white label interchain Bridge to your existing project.
            </p>
          </div>
          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>
          <div className="pl-8 md:p-4 flex-1 col-start-1 md:col-start-auto md:row-start-1">
            <h3 className="text-2xl py-2">Integrated</h3>
            <p>Launch the real-time interchain project explorer.</p>
          </div>
          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>
        </div>
      </section>
      <section
        id="trusted-by"
        className="text-center w-full my-24 grid gap-12 max-w-screen-xl mx-auto"
      >
        <h2 className="text-5xl font-medium">Trusted by</h2>
        <ul className="flex">
          <li className="max-w-lg mx-auto">[ GMX ]</li>
          <li className="max-w-lg mx-auto">[ DeFi Kingdoms ]</li>
          <li className="max-w-lg mx-auto">[ thirdweb ]</li>
          <li className="max-w-lg mx-auto">[ Hercules ]</li>
        </ul>
      </section>
      <section
        id="community"
        className="text-center w-full max-w-screen-lg mx-auto"
      >
        <h2 className="text-5xl font-medium my-6">Community</h2>
        <a href="#" className="underline text-sky-300">
          Discord
        </a>
        ,{' '}
        <a href="#" className="underline text-sky-300">
          Telegram
        </a>
        , and{' '}
        <a href="#" className="underline text-sky-300">
          X
        </a>
        <div className="grid grid-cols-3 justify-items-center my-8">
          <div className="rounded-full bg-zinc-800 border border-zinc-700 w-12 h-12 grid place-content-center z-10">
            <DiscordSvg width={32} height={32} />
          </div>
          <div className="rounded-full bg-zinc-800 border border-zinc-700 w-12 h-12 grid place-content-center z-10">
            <TelegramSvg width={32} height={32} />
          </div>
          <div className="rounded-full bg-zinc-800 border border-zinc-700 w-12 h-12 grid place-content-center z-10">
            <TwitterSvg width={32} height={32} />
          </div>
          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-4 sm:p-8 -my-8 sm:-my-12"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>
          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-4 sm:p-8 -my-8 sm:-my-12"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>

          <svg
            viewBox="-100 -50 200 100"
            className="stroke-zinc-800 fill-zinc-500/5 p-4 sm:p-8 -my-8 sm:-my-12"
          >
            <path d="m-100 0 100 -50 100 50 -100 50z" />
          </svg>
        </div>
      </section>
      {/* <section className="flex odd:flex-col even:flex-col-reverse md:grid grid-cols-2 gap-12 items-center p-4 max-w-4xl">
        <div>
          <h2 className="text-4xl font-medium mb-4">Interchain apps</h2>
          <p className="text-lg leading-relaxed mb-4">
            Synapse Bridge is built on top of the cross-chain infrastructure
            enabling users to seamlessly transfer assets across all blockchains.
            The Bridge has become the most widely-used method to move assets
            cross-chain, offering low cost, fast, and secure bridging.
          </p>
        </div>
        <div className="grid justify-center">
          <FauxBridge />
        </div>
      </section> */}
    </Wrapper>
  )
}

export default LandingPage

function DiscordSvg({
  width,
  height,
  className = 'fill-black dark:fill-white',
}) {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 48 48"
      className={className}
    >
      <path d="M40.6606 9.01494C37.5069 7.58618 34.1777 6.57257 30.7578 6C30.2898 6.82785 29.8664 7.6796 29.4893 8.55168C25.8465 8.00847 22.142 8.00847 18.4992 8.55168C18.1219 7.67969 17.6985 6.82795 17.2307 6C13.8087 6.5774 10.4772 7.59342 7.32038 9.02241C1.05329 18.198 -0.645622 27.1457 0.203833 35.9663C3.874 38.6497 7.98197 40.6905 12.3492 41.9999C13.3325 40.6911 14.2027 39.3027 14.9504 37.8492C13.5302 37.3244 12.1595 36.6768 10.8541 35.914C11.1977 35.6674 11.5337 35.4134 11.8584 35.1668C15.6567 36.9345 19.8025 37.851 23.9999 37.851C28.1974 37.851 32.3431 36.9345 36.1415 35.1668C36.4699 35.4321 36.8059 35.6861 37.1457 35.914C35.8378 36.678 34.4646 37.3268 33.0419 37.853C33.7887 39.3057 34.6589 40.693 35.6431 41.9999C40.014 40.6957 44.1252 38.6559 47.796 35.97C48.7927 25.7409 46.0933 16.8754 40.6606 9.01494ZM16.0264 30.5417C13.6592 30.5417 11.7036 28.4159 11.7036 25.8007C11.7036 23.1855 13.5913 21.0411 16.0188 21.0411C18.4464 21.0411 20.3869 23.1855 20.3454 25.8007C20.3038 28.4159 18.4388 30.5417 16.0264 30.5417ZM31.9735 30.5417C29.6025 30.5417 27.6545 28.4159 27.6545 25.8007C27.6545 23.1855 29.5421 21.0411 31.9735 21.0411C34.4048 21.0411 36.3302 23.1855 36.2887 25.8007C36.2472 28.4159 34.3859 30.5417 31.9735 30.5417Z" />
    </svg>
  )
}

function TelegramSvg({
  width,
  height,
  className = 'fill-black dark:fill-white',
}) {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 48 48"
      overflow="visible"
      className={className}
    >
      <path d="M17.2652 30.1665L16.5374 40.4041C17.5787 40.4041 18.0298 39.9567 18.5706 39.4195L23.4529 34.7536L33.5695 42.1623C35.4248 43.1963 36.732 42.6518 37.2326 40.4554L43.873 9.33932L43.8749 9.33749C44.4634 6.59476 42.883 5.52223 41.0753 6.19508L2.04271 21.1389C-0.621189 22.173 -0.580855 23.658 1.58986 24.3309L11.5689 27.4348L34.7483 12.9309C35.8392 12.2086 36.831 12.6082 36.0152 13.3306L17.2652 30.1665Z" />
    </svg>
  )
}

function TwitterSvg({
  width,
  height,
  className = 'fill-black dark:fill-white',
}) {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 48 48"
      className={className}
    >
      <path d="M35.5023 6H41.6359L28.2359 21.3154L44 42.1563H31.6568L21.9892 29.5164L10.9272 42.1563H4.78988L19.1226 25.7747L4 6H16.6566L25.3953 17.5533L35.5023 6ZM33.3497 38.4851H36.7483L14.8098 9.47842H11.1627L33.3497 38.4851Z" />
    </svg>
  )
}