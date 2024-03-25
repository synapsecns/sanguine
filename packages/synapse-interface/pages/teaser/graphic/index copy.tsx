import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import Hero from '../Hero'
import ValueProps from '../ValueProps'

import Wrapper from '@/components/WipWrapperComponents/Wrapper'

import styles from './keyframes.module.css'

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
      <svg
        width="1200"
        height="675"
        viewBox="-600 -337.5 1200 675"
        className={`border border-zinc-900 mx-auto my-8`}
        fill="none"
        stroke="#0ff"
        stroke-width="1"
      >
        <defs>
          <polygon id="platform" points="0,-100 200,0 0,100 -200,0" />
        </defs>
        <style>
          {`@keyframes circlePulse { from { r: 50; } to { r: 100; } }`}
          {`@keyframes platformBob { from { transform: translate(0, -1rem); } to { transform: translate(0, 0); } }`}
        </style>
        {/* <use
          href="#platform"
          style={{
            animation: '2s ease-in-out 0s infinite alternate platformBob',
          }}
        />
        <use
          href="#platform"
          style={{
            animation: '2s ease-in-out 0.25s infinite alternate platformBob',
          }}
        />
        <use
          href="#platform"
          style={{
            animation: '2s ease-in-out 0.5s infinite alternate platformBob',
          }}
        />
        <use
          href="#platform"
          style={{
            animation: '2s ease-in-out 1s infinite alternate platformBob',
          }}
        /> */}
        <g transform="translate(400,0)">
          <g
            style={{
              animation: '2s ease-in-out 0s infinite alternate platformBob',
            }}
          >
            <use href="#platform" />
            {/* <circle cx="0" cy="0" r="50" /> */}
          </g>
        </g>
        <g transform="translate(-400,0)">
          <g
            style={{
              animation: '2s ease-in-out 0.5s infinite alternate platformBob',
            }}
          >
            <use href="#platform" />
            {/* <circle cx="0" cy="0" r="50" /> */}
          </g>
        </g>
        <g transform="translate(0,200)">
          <g
            style={{
              animation: '2s ease-in-out 1s infinite alternate platformBob',
            }}
          >
            <use href="#platform" />
            {/* <circle cx="0" cy="0" r="50" /> */}
          </g>
        </g>
        <g transform="translate(0,-200)">
          <g
            style={{
              animation: '2s ease-in-out 1.5s infinite alternate platformBob',
            }}
          >
            <use href="#platform" />
            {/* <circle cx="0" cy="0" r="50" /> */}
          </g>
        </g>
        {/* <path d="m0,-50 100,50 -100,50 -100,-50 100,-50" /> */}
      </svg>
      <Hero />
      <ValueProps />
    </Wrapper>
  )
}

export default LandingPage
