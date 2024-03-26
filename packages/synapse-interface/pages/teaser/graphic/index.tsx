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
      <Hero />
      {/* <div className="hidden sm:block text-center text-3xl sm:text-6xl font-semibold mt-16 cursor-default">
        Secure cross-chain
        <br />
        communication
      </div> */}
      <svg
        id="hero-graphic"
        width="1200"
        height="675"
        viewBox="-700 -437.5 1400 875"
        // className={`border border-zinc-900 mx-auto my-8`}
        stroke="#0ff"
        stroke-width="1"
        fill="none"
      >
        <style>
          {/* {`@keyframes circlePulse { from { r: 50; } to { r: 100; } }`} */}
          {`@keyframes platformBob { from { transform: translate(0, -.25rem); } to { transform: translate(0, 0); } }`}
          {`#hero-graphic {
            --fill-synapse: hsl(300deg 100% 5%);
            --stroke-synapse: hsl(300deg 100% 25%);
            --fill-yellow: hsl(60deg 30% 3%);
            --stroke-yellow: hsl(60deg 80% 60%);
            --fill-orange: hsl(25deg 30% 3%);
            --stroke-orange: hsl(25deg 80% 60%);
            --fill-blue: hsl(195deg 30% 3%);
            --stroke-blue: hsl(195deg 80% 60%);
            --fill-green: hsl(135deg 30% 3%);
            --stroke-green: hsl(135deg 80% 60%);
          }`}
        </style>
        <defs>
          <g id="box" transform="scale(.25,.25)">
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
          <g id="boxes" transform="scale(.25,.25)">
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
          id="simple-bridge-ne"
          d="m-120,-160 40,20 -200,100 -40,-20z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />
        <path
          id="simple-bridge-sw"
          d="m280,40 40,20 -200,100 -40,-20z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />
        {/* <path
          id="simple-bridge-ew"
          d="m-240,-20 h480 v40 h-480z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        /> */}
        <path
          id="dock1"
          d="m-80,-300 40,20 -120,60 -40,-20z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />
        <path
          id="dock2"
          d="m-480,-100 40,20 -120,60 -40,-20z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />
        <path
          id="airpad1"
          d="m200,-200 60,30 -60,30 -60,-30z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />
        <path
          id="airpad2"
          d="m200,-60 60,30 -60,30 -60,-30z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />

        {/* <ellipse
          cx="0"
          cy="-200"
          rx="180"
          ry="90"
          stroke="var(--stroke-blue)"
          fill="var(--fill-blue)"
        />

        <ellipse
          cx="-400"
          cy="0"
          rx="180"
          ry="90"
          stroke="var(--stroke-yellow)"
          fill="var(--fill-yellow)"
        />
        <ellipse
          cx="400"
          cy="0"
          rx="180"
          ry="90"
          stroke="var(--stroke-green)"
          fill="var(--fill-green)"
        />
        <ellipse
          cx="0"
          cy="200"
          rx="200"
          ry="100"
          stroke="var(--stroke-orange)"
          fill="var(--fill-orange)"
        /> */}
        <path
          id="platform-blue"
          d="m0,-300 200,100 -200,100 -200,-100z"
          stroke="var(--stroke-blue)"
          fill="hsl(300deg 20% 5%)"
        />
        <path
          id="platform-yellow"
          d="m-400,-100 200,100 -200,100 -200,-100z"
          stroke="var(--stroke-yellow)"
          fill="hsl(300deg 20% 5%)"
        />
        <path
          id="platform-green"
          d="m400,-100 200,100 -200,100 -200,-100z"
          stroke="var(--stroke-green)"
          fill="hsl(300deg 20% 5%)"
        />
        <path
          id="platform-orange"
          d="m0,100 200,100 -200,100 -200,-100z"
          stroke="var(--stroke-orange)"
          fill="hsl(300deg 20% 5%)"
        />
        <g
          id="stack-green"
          fill="var(--fill-green)"
          stroke="var(--stroke-green)"
          transform="translate(400,0)"
        >
          <use href="#box" transform="translate(0,-27.95)" />
          <use href="#box" transform="translate(25,12.5)" />
          <use href="#box" transform="translate(-25,12.5)" />
          <use href="#box" transform="translate(0,25)" />
        </g>

        <g
          id="stack-orange"
          fill="var(--fill-orange)"
          stroke="var(--stroke-orange)"
          transform="translate(0,150)"
        >
          <use href="#box" transform="translate(0,-27.95)" />
          <use href="#box" transform="translate(25,12.5)" />
          <use href="#box" transform="translate(-25,12.5)" />
          <use href="#box" transform="translate(0,25)" />
        </g>
        <g
          id="barge"
          style={
            {
              // animation: '2s ease-in-out 0s infinite alternate platformBob',
            }
          }
        >
          <path
            d="m70,-75 100,50 -200,100 -100,-50z"
            fill="var(--fill-synapse)"
            stroke="var(--stroke-synapse)"
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
          id="bargeBox"
          href="#box"
          stroke="hsl(195deg 80% 60%)"
          fill="hsl(195deg 100% 5%)"
        >
          <animate
            attributeName="opacity"
            values="0; 1"
            dur="2s"
            begin="bargeIn.begin"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%)"
            begin="bargeIn.begin"
          />
          <animateMotion
            path="m-50,-225"
            fill="freeze"
            begin="0s; bargeIn.begin"
          />
          <animateMotion
            dur="2s"
            begin="bargeIn.end"
            path="m-50,-225 -140,-70"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            dur="2s"
            begin="bargeMid.begin"
            path="m-190,-295 -400,200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%); hsl(300deg 100% 40%); hsl(60deg 80% 60%)"
            begin="bargeMid.begin + .5s"
            dur=".5s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".33 0 .33 1; .33 0 .33 1"
            fill="freeze"
          />
          <animateMotion
            id="bargeBoxEnd"
            dur="2s"
            begin="bargeMid.end"
            path="m-590,-95 162.5,81.25"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="1; 0"
            dur="4s"
            begin="bargeBoxEnd.end"
          />
        </use>
        <g
          id="stack-blue"
          fill="var(--fill-blue)"
          stroke="var(--stroke-blue)"
          transform="translate(0,-225)"
        >
          <use href="#box" transform="translate(0,-27.95)" />
          <use href="#box" transform="translate(25,12.5)" />
          <use href="#box" transform="translate(-25,12.5)" />
          {/* <use href="#box" transform="translate(0,25)" /> */}
        </g>
        <use href="#box" id="simpleBridgeBoxFromBlue">
          <animateMotion
            id="boxOut"
            dur="3s"
            begin="0s; boxIn.end + 1s"
            path="M-375,-12.5 0,-200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="boxIn"
            dur="3s"
            begin="boxOut.end + 1s"
            path="M0,-200 -375,-12.5"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(60deg 80% 60%); hsl(300deg 100% 40%); hsl(195deg 80% 60%)"
            begin="boxOut.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="fill"
            values="hsl(60deg 80% 5%); hsl(300deg 100% 5%); hsl(195deg 80% 5%)"
            begin="boxOut.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%); hsl(300deg 100% 40%); hsl(60deg 80% 60%)"
            begin="boxIn.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="fill"
            values="hsl(195deg 80% 5%); hsl(300deg 100% 5%); hsl(60deg 80% 5%)"
            begin="boxIn.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </use>
        <use href="#box" id="simpleBridgeBoxFromYellow">
          <animateMotion
            id="boxesOut"
            dur="3s"
            begin="0s; boxOut.begin"
            path="M0,-200 -375,-12.5"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            id="boxesIn"
            dur="3s"
            begin="boxIn.begin"
            path="M-375,-12.5 0,-200"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(60deg 80% 60%); hsl(300deg 100% 40%); hsl(195deg 80% 60%)"
            begin="boxesIn.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="fill"
            values="hsl(60deg 80% 5%); hsl(300deg 100% 5%); hsl(195deg 80% 5%)"
            begin="boxesIn.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%); hsl(300deg 100% 40%); hsl(60deg 80% 60%)"
            begin="boxesOut.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="fill"
            values="hsl(195deg 80% 5%); hsl(300deg 100% 5%); hsl(60deg 80% 5%)"
            begin="boxesOut.begin + 1s"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </use>
        <g id="teleporter1">
          <ellipse
            rx="30"
            ry="15"
            cy="260"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          />
          <g transform="scale(.25,.25)" stroke="hsl(25deg 80% 60%)">
            <animate
              attributeName="stroke"
              values="hsl(25deg 80% 60%); hsl(300deg 100% 40%); hsl(135deg 80% 60%)"
              dur="3s"
              begin="teleporterBeamsOut.begin + 1s"
              calcMode="spline"
              keyTimes="0; .5; 1"
              keySplines=".5 0 .2 1; .5 0 .2 1"
              fill="freeze"
            />
            <animate
              attributeName="stroke"
              values="hsl(135deg 80% 60%); hsl(300deg 100% 40%); hsl(25deg 80% 60%)"
              dur="3s"
              begin="teleporterBeamsIn.begin + 1s"
              calcMode="spline"
              keyTimes="0; .5; 1"
              keySplines=".5 0 .2 1; .5 0 .2 1"
              fill="freeze"
            />
            <path
              d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
              vectorEffect="non-scaling-stroke"
            />
            <path
              d="m-100,-111.9 100,50 100,-50 m-100,50 0,111.8"
              vectorEffect="non-scaling-stroke"
              fill="none"
            />
            <animateMotion
              id="teleport1BobOut"
              dur="1s"
              begin="0s; teleport1BobIn.end"
              path="m0,240 0,12.5"
              calcMode="spline"
              keyTimes="0; 1"
              keySplines=".33 0 .67 1"
              fill="freeze"
            />
            <animateMotion
              id="teleport1BobIn"
              dur="1s"
              begin="teleport1BobOut.end"
              path="m0,252.5 0,-12.5"
              calcMode="spline"
              keyTimes="0; 1"
              keySplines=".33 0 .67 1"
              fill="freeze"
            />
          </g>
          <g id="teleporter1Beams" opacity="0" strokeWidth="3">
            <animate
              id="teleporterBeamsOut"
              attributeName="opacity"
              values="0; 1; 0"
              begin="2s; teleporterBeamsIn.end + 2s"
              dur="5s"
              fill="freeze"
            />
            <animate
              id="teleporterBeamsIn"
              attributeName="opacity"
              values="0; 1; 0"
              begin="teleporterBeamsOut.end + 2s"
              dur="5s"
            />
            <path
              d="m-20,207.5 0,50"
              stroke="hsl(300deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="0; 28"
                dur="1s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m-10,212.5 0,50"
              stroke="hsl(285deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="28; 0"
                dur=".67s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m0,207.5 0,50"
              stroke="hsl(300deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="0; 28"
                dur="1s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m10,212.5 0,50"
              stroke="hsl(285deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="28; 0"
                dur=".67s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m20,207.5 0,50"
              stroke="hsl(300deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="28; 0"
                dur="1s"
                repeatCount="indefinite"
              />
            </path>
          </g>
          <ellipse
            rx="30"
            ry="15"
            cy="190"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          />
        </g>

        <g id="teleporter2" transform="translate(520,-260)">
          <ellipse
            rx="30"
            ry="15"
            cy="260"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          />
          <g transform="scale(.25,.25)" stroke="hsl(25deg 80% 60%)">
            <animate
              attributeName="stroke"
              values="hsl(135deg 80% 60%); hsl(300deg 100% 40%); hsl(25deg 80% 60%)"
              dur="3s"
              begin="teleporterBeamsOut.begin + 1s"
              calcMode="spline"
              keyTimes="0; .5; 1"
              keySplines=".5 0 .2 1; .5 0 .2 1"
              fill="freeze"
            />
            <animate
              attributeName="stroke"
              values="hsl(25deg 80% 60%); hsl(300deg 100% 40%); hsl(135deg 80% 60%)"
              dur="3s"
              begin="teleporterBeamsIn.begin + 1s"
              calcMode="spline"
              keyTimes="0; .5; 1"
              keySplines=".5 0 .2 1; .5 0 .2 1"
              fill="freeze"
            />
            <path
              d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
              vectorEffect="non-scaling-stroke"
            />
            <path
              d="m-100,-111.9 100,50 100,-50 m-100,50 0,111.8"
              vectorEffect="non-scaling-stroke"
              fill="none"
            />
            <animateMotion
              dur="1s"
              begin="teleport1BobOut.end"
              path="m0,240 0,12.5"
              calcMode="spline"
              keyTimes="0; 1"
              keySplines=".33 0 .67 1"
              fill="freeze"
            />
            <animateMotion
              dur="1s"
              begin="0s; teleport1BobIn.end"
              path="m0,252.5 0,-12.5"
              calcMode="spline"
              keyTimes="0; 1"
              keySplines=".33 0 .67 1"
              fill="freeze"
            />
          </g>
          <g id="teleporter2Beams" opacity="0" strokeWidth="3">
            <animate
              attributeName="opacity"
              values="0; 1; 0"
              begin="2s; teleporterBeamsIn.end + 2s"
              dur="5s"
              fill="freeze"
            />
            <animate
              attributeName="opacity"
              values="0; 1; 0"
              begin="teleporterBeamsOut.end + 2s"
              dur="5s"
            />
            <path
              d="m-20,207.5 0,50"
              stroke="hsl(300deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="0; 28"
                dur="1s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m-10,212.5 0,50"
              stroke="hsl(285deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="28; 0"
                dur=".67s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m0,207.5 0,50"
              stroke="hsl(300deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="0; 28"
                dur="1s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m10,212.5 0,50"
              stroke="hsl(285deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="28; 0"
                dur=".67s"
                repeatCount="indefinite"
              />
            </path>
            <path
              d="m20,207.5 0,50"
              stroke="hsl(300deg 80% 60%)"
              strokeDasharray="8 6 6 8"
            >
              <animate
                attributeName="stroke-dashoffset"
                values="28; 0"
                dur="1s"
                repeatCount="indefinite"
              />
            </path>
          </g>
          <ellipse
            rx="30"
            ry="15"
            cy="190"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          />
        </g>
        <g transform="scale(.25,.25)" stroke="var(--stroke-blue)">
          <animateMotion
            id="airlift"
            dur="2s"
            begin="0s; airdrop.end + 5s"
            path="m200,-180 v-50"
            fill="freeze"
          />
          <animateMotion
            id="airpath"
            dur="5s"
            begin="airlift.end + 1s"
            path="m200,-230 v150"
            fill="freeze"
          />
          <animateMotion
            id="airdrop"
            dur=".5s"
            begin="airpath.end + 1s"
            path="m200,-80 v50"
            fill="freeze"
          />
          <path
            d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
            vectorEffect="non-scaling-stroke"
            fill="var(--fill-orange)"
          />
          <path
            d="m-100,-111.9 100,50 100,-50 m-100,50 0,111.8"
            vectorEffect="non-scaling-stroke"
            fill="none"
          />
          <g id="balloon">
            <animateMotion begin="airlift.begin" path="m0,0" />
            <animateMotion
              dur="1s"
              begin="airdrop.begin"
              path="m0,0 v-2000"
              fill="freeze"
            />
            <path
              d="m0,-111.8 v-111.8"
              vectorEffect="non-scaling-stroke"
              stroke="var(--stroke-synapse)"
            />
            <circle
              cy="-370"
              r="150"
              vectorEffect="non-scaling-stroke"
              stroke="var(--stroke-synapse)"
              fill="var(--fill-synapse)"
            />
          </g>
        </g>
        <g
          id="stack-yellow"
          fill="var(--fill-yellow)"
          stroke="var(--stroke-yellow)"
          transform="translate(-400,0)"
        >
          <use href="#box" transform="translate(0,-27.95)" />
          <use href="#box" transform="translate(25,12.5)" />
          <use href="#box" transform="translate(-25,12.5)" />
          <use href="#box" transform="translate(0,25)" />
        </g>
      </svg>
      <p className="text-center">Reference image</p>
      <img src={exampleImg.src} className="visible" />
      {/* <Hero /> */}
      <section>
        <ul className="w-fit md:w-max grid md:flex text-xl md:text-lg text-center items-center place-center bg-gradient-to-b from-white to-slate-100 dark:from-zinc-900 dark:to-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md px-6 gap-x-8 -mt-8 shadow-sm mx-auto mb-16 cursor-default">
          <li className="-mt-1 p-3">
            $<data className="mx-0.5">45.3B</data> Bridge volume
          </li>
          <li className="-mt-1 p-3">
            <data className="mx-0.5">10.6M</data> transactions
          </li>
          <li className="-mt-1 p-3">
            $<data className="mx-0.5">116.7M</data> Total value locked
          </li>
        </ul>
      </section>
      <ValueProps />
    </Wrapper>
  )
}

export default LandingPage
