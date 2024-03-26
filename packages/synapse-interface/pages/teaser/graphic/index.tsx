import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import exampleImg from '@assets/example.png'

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
        id="hero-graphic"
        width="1200"
        height="675"
        viewBox="-600 -337.5 1200 675"
        className={`border border-zinc-900 mx-auto my-8`}
        stroke="#0ff"
        stroke-width="1"
        fill="none"
      >
        <style>
          {/* {`@keyframes circlePulse { from { r: 50; } to { r: 100; } }`} */}
          {`@keyframes platformBob { from { transform: translate(0, -.25rem); } to { transform: translate(0, 0); } }`}
          {`#hero-graphic {
            --synapse-fill: hsl(300deg 100% 5%);
            --synapse-stroke: hsl(300deg 100% 25%);
            --yellow-fill: hsl(60deg 100% 5%);
            --yellow-stroke: hsl(60deg 100% 50%);
            --orange-fill: hsl(25deg 100% 5%);
            --orange-stroke: hsl(25deg 100% 50%);
            --blue-fill: hsl(195deg 100% 5%);
            --blue-stroke: hsl(195deg 100% 50%);
            --green-fill: hsl(135deg 100% 5%);
            --green-stroke: hsl(135deg 100% 50%);
          }`}
        </style>
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
        <path
          id="simple-bridge"
          // d="M-120,40 80,-60 120,-40 -80,60z"
          d="m-120,-160 40,20 -200,100 -40,-20z"
          fill="var(--synapse-fill"
          stroke="var(--synapse-stroke)"
        />
        <path
          id="platform-blue"
          d="m0,-300 200,100 -200,100 -200,-100z"
          stroke="var(--blue-stroke)"
        />
        <path
          id="platform-yellow"
          d="m-400,-100 200,100 -200,100 -200,-100z"
          stroke="var(--yellow-stroke)"
        />
        <path
          id="platform-green"
          d="m400,-100 200,100 -200,100 -200,-100z"
          stroke="var(--green-stroke)"
        />
        <path
          id="platform-orange"
          d="m0,100 200,100 -200,100 -200,-100z"
          stroke="var(--orange-stroke)"
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
            fill="var(--synapse-fill)"
            stroke="var(--synapse-stroke)"
          />
          <animateMotion
            id="bargeIn"
            dur="2s"
            begin="0s; bargeEnd.end + 2s"
            path="M200,-500 -200,-300"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeMid"
            dur="2s"
            begin="bargeIn.end + 2s"
            path="M-200,-300 -600,-100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeEnd"
            dur="2s"
            begin="bargeMid.end + 2s"
            path="M-600,-100 -1000,100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
        </g>
        <use
          href="#boxes"
          stroke="hsl(195deg 100% 50%)"
          fill="hsl(195deg 100% 5%)"
        >
          <animateMotion
            path="M-50,-225"
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
            path="M-50,-225 -200,-300"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            dur="2s"
            begin="bargeMid.begin"
            path="M-200,-300 -600,-100"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeBoxEnd"
            dur="2s"
            begin="bargeMid.end"
            path="M-600,-100 -450,-25"
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
        </use>

        <use href="#box">
          <animateMotion
            id="boxOut"
            dur="3s"
            begin="0s; boxIn.end + 1s"
            path="M-400,0 0,-200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="boxIn"
            dur="3s"
            begin="boxOut.end + 1s"
            path="M0,-200 -400,0"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(60deg 100% 50%); hsl(300deg 100% 40%); hsl(195deg 100% 50%)"
            begin="boxOut.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 100% 50%); hsl(300deg 100% 40%); hsl(60deg 100% 50%)"
            begin="boxIn.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </use>
        <use href="#boxes">
          <animateMotion
            id="boxesOut"
            dur="3s"
            begin="0s; boxOut.begin"
            path="M0,-200 -400,0"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="boxesIn"
            dur="3s"
            begin="boxIn.begin"
            path="M-400,0 0,-200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(60deg 100% 50%); hsl(300deg 100% 40%); hsl(195deg 100% 50%)"
            begin="boxesIn.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 100% 50%); hsl(300deg 100% 40%); hsl(60deg 100% 50%)"
            begin="boxesOut.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
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
      <img src={exampleImg.src} />
      <Hero />
      <ValueProps />
    </Wrapper>
  )
}

export default LandingPage
