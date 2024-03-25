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
        stroke="#0ff"
        stroke-width="1"
        fill="none"
      >
        <defs>
          <g id="box">
            <path
              d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
              transform="scale(.33,.33)"
              vectorEffect="non-scaling-stroke"
              pathLength="1"
              strokeDasharray="1"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="1; 0"
                dur="2s"
                repeatCount="1"
              />
            </path>
            <path
              d="m100,0 -100,-50 -100,50 m0,-111.8 100,50 100,-50 -100,-50 0,211.8"
              transform="scale(.33,.33)"
              vectorEffect="non-scaling-stroke"
              fill="none"
              pathLength="1"
              strokeDasharray="1"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="1; 0"
                dur="3s"
                repeatCount="1"
              />
            </path>
          </g>
          <polygon
            id="platform"
            points="0,-100 200,0 0,100 -200,0"
            vectorEffect="non-scaling-stroke"
            pathLength="1"
            fill="#111"
          >
            <animate
              attributeName="stroke-dasharray"
              values="0 1; 1 0"
              dur="0.5s"
              repeatCount="1"
            />
          </polygon>
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
        <g
          transform="translate(400,0)"
          stroke="hsl(150deg 100% 50%)"
          fill="hsl(150deg 100% 50% / 10%)"
        >
          <use href="#platform" />
          <use href="#box" />
          <use href="#box" transform="translate(33,16.5)" />
          <use href="#box" transform="translate(-33,16.5)" />
          <use href="#box" transform="translate(0,-37)"></use>
        </g>
        <g
          transform="translate(-400,0)"
          stroke="hsl(50deg 100% 50%)"
          fill="hsl(50deg 100% 50% / 10%)"
        >
          <use href="#platform" />
          <use href="#box" transform="translate(33,16.5)" />
          <use href="#box" transform="translate(-33,16.5)" />
          <use href="#box" transform="translate(0,-37)" />
        </g>
        <g
          transform="translate(0,200) scale(1.25,1.25)"
          stroke="hsl(35deg 100% 50%)"
          fill="hsl(35deg 100% 50% / 10%)"
        >
          <use href="#platform" />
          <use href="#box" transform="translate(33,16.5)" />
          <use href="#box" transform="translate(-33,16.5)" />
          <use href="#box" transform="translate(0,-37)" />
        </g>
        <g
          transform="translate(0,-200) scale(1.25,1.25)"
          stroke="hsl(195deg 100% 50%)"
          fill="hsl(195deg 100% 50% / 10%)"
        >
          <use href="#platform" />
          <use href="#box" transform="translate(33,16.5)" />
          <use href="#box" transform="translate(-33,16.5)" />
          <use href="#box" transform="translate(0,-37)" />
        </g>
        {/* <path d="m0,-50 100,50 -100,50 -100,-50 100,-50" /> */}
      </svg>
      <Hero />
      <ValueProps />
    </Wrapper>
  )
}

export default LandingPage
