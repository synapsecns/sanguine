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
          <g id="box" transform="scale(.33,.33)">
            <path
              d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
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
              // d="m100,0 -100,-50 -100,50 m0,-111.8 100,50 100,-50 -100,-50 0,211.8"
              d="m-100,-111.9 100,50 100,-50 m-100,50 0,111.8"
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
          <g id="boxes" transform="scale(.33,.33)">
            <path
              d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
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
              // d="m100,0 -100,-50 -100,50 m0,-111.8 100,50 100,-50 -100,-50 0,211.8 m-100,-105.9 100,50 100,-50 -100,-50 -100,50 m50,-80.9 100 50 0,111.8 -100,-50 0,-111.8 m100,0 0,111.8 -100,50 0,-111.8 100,-50"
              d="m100,-111.8 -100,50 -100,-50 m0,55.9 100,50 100,-50 m-50,80.9 0,-111.8 -100,-50 m100,0 -100,50 0,111.8 m50,27.95 0,-111.8"
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
          {/* {`@keyframes circlePulse { from { r: 50; } to { r: 100; } }`} */}
          {`@keyframes platformBob { from { transform: translate(0, -.25rem); } to { transform: translate(0, 0); } }`}
        </style>
        <path
          id="simple-bridge"
          d="M-120,40 80,-60 120,-40 -80,60z"
          fill="hsl(300deg 100% 5%)"
          stroke="hsl(300deg 100% 25%)"
        />
        <path
          id="platform-yellow"
          d="m-200,0 200,100 -200,100 -200,-100z"
          stroke="hsl(60deg 100% 50%)"
        />
        <path
          id="platform-blue"
          d="m200,-200 200,100 -200,100 -200,-100z"
          stroke="hsl(195deg 100% 50%)"
        />
        <path
          id="platform-green"
          d="m500,-50 200,100 -200,100 -200,-100z"
          stroke="hsl(135deg 100% 50%)"
        />
        <g
          id="barge"
          style={
            {
              // animation: '2s ease-in-out 0s infinite alternate platformBob',
            }
          }
        >
          <path
            id="barge"
            d="m50,-75 100,50 -200,100 -100,-50z"
            stroke="hsl(300deg 100% 25%)"
            fill="hsl(300deg 100% 5%)"
          />
          <animateMotion
            id="bargeIn"
            dur="2s"
            begin="0s; bargeEnd.end + 2s"
            path="M400,-400 0,-200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeMid"
            dur="2s"
            begin="bargeIn.end + 2s"
            path="M0,-200 -400,0"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeEnd"
            dur="2s"
            begin="bargeMid.end + 2s"
            path="M-400,0 -800,200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
        </g>
        <g
          id="bargeBoxes"
          style={
            {
              // animation: '2s ease-in-out 0s infinite alternate platformBob',
            }
          }
        >
          <use
            href="#boxes"
            stroke="hsl(195deg 100% 50%)"
            fill="hsl(195deg 100% 5%)"
          />
          <animateMotion
            path="M150,-125"
            fill="freeze"
            begin="0s; bargeIn.begin"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            dur="2s"
            begin="bargeIn.begin"
          />
          <animateMotion
            dur="2s"
            begin="bargeIn.end"
            path="M150,-125 0,-200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            dur="2s"
            begin="bargeMid.begin"
            path="M0,-200 -400,0"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeBoxEnd"
            dur="2s"
            begin="bargeMid.end"
            path="M-400,0 -250,75"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="1; 0"
            dur="5s"
            begin="bargeBoxEnd.end"
          />
        </g>

        <use href="#box" fill="hsl(60deg 100% 5%)" stroke="hsl(60deg 100% 50%)">
          <animateMotion
            id="start1"
            dur="3s"
            begin="0s; mid1.end + 1s"
            path="M-200,100 200,-100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="mid1"
            dur="3s"
            begin="start1.end + 1s"
            path="M200,-100 -200,100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <set
            attributeName="stroke"
            to="hsl(195deg 100% 50%)"
            dur="4s"
            begin="start1.begin + 1.2s"
          />
          <set
            attributeName="fill"
            to="hsl(195deg 100% 5%)"
            dur="4s"
            begin="start1.begin + 1.2s"
          />
        </use>
        <use
          href="#boxes"
          stroke="hsl(195deg 100% 50%)"
          fill="hsl(195deg 100% 5%)"
        >
          <animateMotion
            id="start2"
            dur="3s"
            begin="mid1.begin"
            path="M-200,100 200,-100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="mid2"
            dur="3s"
            begin="0s; start1.begin"
            path="M200,-100 -200,100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <set
            attributeName="stroke"
            to="hsl(60deg 100% 50%)"
            dur="4s"
            begin="mid2.begin + 1.2s"
          />
          <set
            attributeName="fill"
            to="hsl(60deg 100% 5%)"
            dur="4s"
            begin="mid2.begin + 1.2s"
          />
        </use>

        {/* <use href="#box">
          <animateMotion
            dur="5s"
            begin="-2.5s"
            repeatCount="indefinite"
            path="M-200,100 200,-100z"
          />
        </use> */}
        {/* <use href="#boxes" stroke="pink" /> */}
      </svg>
      <Hero />
      <ValueProps />
    </Wrapper>
  )
}

export default LandingPage
