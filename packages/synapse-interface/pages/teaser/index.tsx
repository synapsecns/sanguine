import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { CHAINS_ARR } from '@/constants/chains'

console.log(CHAINS_ARR)

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import Hero from './Hero'

import Wrapper from '@/components/WipWrapperComponents/Wrapper'

import FauxBridge from './FauxBridge'
import PlatformAnimation from './PlatformAnimation'
import { Chain } from '@/utils/types'

const chainTagClassName = 'pl-2 pr-4 py-1 border rounded h-min'

function ChainTag({ chain }: { chain: Chain }): React.ReactNode {
  return (
    <a
      href="#"
      className={`grid grid-cols-[auto_auto] gap-x-2 items-center border-separator ${chainTagClassName}`}
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
    <Wrapper>
      <Hero />
      <PlatformAnimation />
      <section id="chain-list" className="-mt-12 mb-12">
        <dl className="flex gap-2 text-sm whitespace-nowrap">
          <a href="#" className={`${chainTagClassName} border-fuchsia-500`}>
            <dt>6 month vol.</dt>
            <dd>$7.03B</dd>
          </a>
          {CHAINS_ARR.slice(4, 11).map((a) => (
            <ChainTag chain={a} />
          ))}
          <dt className={`${chainTagClassName} border-separator`}>
            DFK, Klaytn,
            <br />
            Cronos & more
          </dt>
        </dl>
      </section>
      <section id="entry-points" className="text-center w-full">
        <h2 className="text-5xl font-medium my-6">The future is interchain</h2>
        <p className="max-w-lg mx-auto">
          The interchain network brings the federated blockchain universe
          together. Access new technologies, trade, loan, stake, and participate
          with any project from any chain.
        </p>
        <h3 className="text-2xl font-medium my-6">Start</h3>
        <svg
          viewBox="0 -1 8 2"
          className="stroke-zinc-700 w-full h-16"
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
        <ul className="flex">
          <li className="p-4 flex-1">
            <h3 className="text-2xl p-2">Instructions provided</h3>
            <p>Everything you need to make your project interchain.</p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
          <li className="p-4 flex-1">
            <h3 className="text-2xl p-2">Batteries included</h3>
            <p>Production-ready customizable interchain contracts.</p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
          <li className="p-4 flex-1">
            <h3 className="text-2xl p-2">Pre-assembled</h3>
            <p>
              Add the white label interchain Bridge to your existing project.
            </p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
          <li className="p-4 flex-1">
            <h3 className="text-2xl p-2">Integrated</h3>
            <p>Launch the real-time interchain project explorer.</p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8 pb-0"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
        </ul>
      </section>
      <section id="trusted-by" className="text-center w-full my-24 grid gap-12">
        <h2 className="text-5xl font-medium">Trusted by</h2>
        <ul className="flex">
          <li className="max-w-lg mx-auto">[ GMX ]</li>
          <li className="max-w-lg mx-auto">[ DeFi Kingdoms ]</li>
          <li className="max-w-lg mx-auto">[ thirdweb ]</li>
          <li className="max-w-lg mx-auto">[ Hercules ]</li>
        </ul>
      </section>
      <section id="community" className="text-center w-full">
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
        <ul className="flex">
          <li className="p-4 flex-1">
            <p>[ Discord ]</p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
          <li className="p-4 flex-1">
            <p>[ Telegram ]</p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
          <li className="p-4 flex-1">
            <p>[ X ]</p>
            <svg
              viewBox="-100 -50 200 100"
              className="stroke-zinc-800 fill-zinc-500/5 p-8"
            >
              <path d="m-100 0 100 -50 100 50 -100 50z" />
            </svg>
          </li>
        </ul>
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
