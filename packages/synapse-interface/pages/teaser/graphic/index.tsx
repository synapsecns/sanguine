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

  function animAttrs(x1 = 0.5, x2 = 0.2, y1 = 0, y2 = 1) {
    return {
      calcMode: 'spline',
      keyTimes: '0; 1',
      keySplines: `${x1} ${y1} ${x2} ${y2}`,
      fill: 'freeze',
    }
  }

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
          <path id="simpleBridgePath" d="m-25 -187.5 -350 175" />
          {/* <g id="box" transform="scale(.25,.25)">
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
          </g> */}
          {/* <g id="boxes" transform="scale(.25,.25)">
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
          </g> */}
        </defs>
        <path fill="var(--fill-synapse" stroke="var(--stroke-synapse)">
          <animate
            id="simpleBridgeNe"
            attributeName="d"
            values="m-200,-100 0,0 0,0 0,0z; m-100,-150 0,0 -200,100 0,0z; m-120,-160 40,20 -200,100 -40,-20z"
            dur=".5s"
            begin="platformYellow.end + 2s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </path>
        <path
          id="simple-bridge-sw"
          d="m280,40 40,20 -200,100 -40,-20z"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        />
        <path
          id="dock1"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        >
          <animate
            id="dockN"
            attributeName="d"
            values="m-40,-240 0,0 0,0 0,0z; m-80,-300 40,20 -120,60 -40,-20z"
            dur=".5s"
            begin="simpleBridgeNe.end + 2s"
            {...animAttrs()}
          />
        </path>
        <path
          id="dock1"
          fill="var(--fill-synapse"
          stroke="var(--stroke-synapse)"
        >
          <animate
            id="dockE"
            attributeName="d"
            values="m-440,-40 0,0 0,0 0,0z; m-480,-100 40,20 -120,60 -40,-20z"
            dur=".5s"
            begin="simpleBridgeNe.end + 2.5s"
            {...animAttrs()}
          />
        </path>
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
        ></path>

        <path stroke="var(--stroke-blue)" fill="var(--fill-blue)">
          <animate
            id="platformBlue"
            attributeName="d"
            values="m0,1 2,1 -2,1 -2,-1z; m0,-100 200,100 -200,100 -200,-100z"
            dur=".25s"
            begin="0s; click"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="platformBlue.begin + .1s"
          />
          <animateTransform
            attributeName="transform"
            attributeType="XML"
            type="translate"
            values="0 -150; 0 -200"
            dur=".5s"
            begin="platformBlue.begin"
            {...animAttrs()}
          />
        </path>
        <path stroke="var(--stroke-green)" fill="var(--fill-green)">
          <animate
            id="platformGreen"
            attributeName="d"
            values="m0,1 2,1 -2,1 -2,-1z; m0,-100 200,100 -200,100 -200,-100z"
            dur=".25s"
            begin="platformBlue.begin + .1s"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="platformGreen.begin + .1s"
          />
          <animateTransform
            attributeName="transform"
            attributeType="XML"
            type="translate"
            values="400 50; 400 0"
            dur=".5s"
            begin="platformGreen.begin"
            {...animAttrs()}
          />
        </path>
        <path stroke="var(--stroke-orange)" fill="var(--fill-orange)">
          <animate
            id="platformOrange"
            attributeName="d"
            values="m0,1 2,1 -2,1 -2,-1z; m0,-100 200,100 -200,100 -200,-100z"
            dur=".25s"
            begin="platformGreen.begin + .1s"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="platformOrange.begin + .1s"
          />
          <animateTransform
            attributeName="transform"
            attributeType="XML"
            type="translate"
            values="0 250; 0 200"
            dur=".5s"
            begin="platformOrange.begin"
            {...animAttrs()}
          />
        </path>
        <path stroke="var(--stroke-yellow)" fill="var(--fill-yellow)">
          <animate
            id="platformYellow"
            attributeName="d"
            values="m0,1 2,1 -2,1 -2,-1z; m0,-100 200,100 -200,100 -200,-100z"
            dur=".25s"
            begin="platformOrange.begin + .1s"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="platformYellow.begin + .1s"
          />
          <animateTransform
            attributeName="transform"
            attributeType="XML"
            type="translate"
            values="-400 50; -400 0"
            dur=".5s"
            begin="platformYellow.begin"
            {...animAttrs()}
          />
        </path>

        {/* <g
          id="stack-green"
          fill="var(--fill-green)"
          stroke="var(--stroke-green)"
          transform="translate(400,0)"
          visibility="hidden"
        >
          <set
            attributeName="visibility"
            to="visible"
            begin="platformYellow.end"
          />
          <set
            attributeName="visibility"
            to="hidden"
            begin="platformBlue.begin"
          />
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
        </g> */}

        <g id="barge">
          <animateMotion path="M200,-500" />
          <path
            d="m70,-75 100,50 -200,100 -100,-50z"
            fill="var(--fill-synapse)"
            stroke="var(--stroke-synapse)"
          />
          <animateMotion
            id="bargeOut"
            dur="2s"
            begin="dockN.end; bargeIn.end"
            path="M200,-500 -200,-300"
            {...animAttrs()}
          />
          <animateMotion
            id="bargeCross"
            dur="2s"
            begin="bargeOut.end + 2s"
            path="M-200,-300 -600,-100"
            {...animAttrs()}
          />
          <animateMotion
            id="bargeIn"
            dur="2s"
            begin="bargeCross.end + 2s"
            path="M-600,-100 -1000,100"
            {...animAttrs()}
          />
        </g>

        <g id="teleportFromOrange" stroke="hsl(25deg 80% 60%)">
          <ellipse
            rx="30"
            ry="15"
            cy="260"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          />
          <g transform="scale(.25,.25)">
            <animate
              attributeName="stroke"
              values="inherit; hsl(300deg 100% 40%); hsl(135deg 80% 60%)"
              dur="3s"
              begin="teleporterBeamsOut.begin + 1s"
              calcMode="spline"
              keyTimes="0; .5; 1"
              keySplines=".5 0 .2 1; .5 0 .2 1"
              fill="freeze"
            />
            <animate
              attributeName="stroke"
              values="hsl(135deg 80% 60%); hsl(300deg 100% 40%); inherit"
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

        <g
          id="teleportFromGreen"
          transform="translate(520,-260)"
          stroke="hsl(135deg 80% 60%)"
        >
          <ellipse
            rx="30"
            ry="15"
            cy="260"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          />
          <g transform="scale(.25,.25)">
            <animate
              attributeName="stroke"
              values="inherit; hsl(300deg 100% 40%); hsl(25deg 80% 60%)"
              dur="3s"
              begin="teleporterBeamsOut.begin + 1s"
              calcMode="spline"
              keyTimes="0; .5; 1"
              keySplines=".5 0 .2 1; .5 0 .2 1"
              fill="freeze"
            />
            <animate
              attributeName="stroke"
              values="hsl(25deg 80% 60%); hsl(300deg 100% 40%); inherit"
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
        <g
          id="balloonBox"
          transform="scale(.25,.25)"
          // stroke="var(--stroke-blue)"
        >
          <animateMotion
            id="stackOut"
            dur="1s"
            begin="0s; stackIn.end + 5s"
            path="M25,-227.95 200,-170"
            {...animAttrs()}
          />
          <animateMotion
            id="airlift"
            dur="2s"
            begin="stackOut.end + 1s"
            path="m200,-170 v-50"
            {...animAttrs()}
          />
          <animateMotion
            id="airpath"
            dur="4s"
            begin="airlift.end"
            path="m200,-220 v150"
            {...animAttrs()}
          />
          <animateMotion
            id="airdrop"
            dur=".25s"
            begin="airpath.end"
            path="m200,-70 v40"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".33 0 1 1"
            fill="freeze"
          />
          <animateMotion
            id="stackIn"
            dur="1s"
            begin="airdrop.end + .5s"
            path="m200,-30 174,68"
            {...animAttrs()}
          />
          <path
            d="m0,50 100,-50 0,-111.8 -100,-50 -100,50 0,111.8 100,50"
            vectorEffect="non-scaling-stroke"
          />
          <path
            d="m-100,-111.9 100,50 100,-50 m-100,50 0,111.8"
            vectorEffect="non-scaling-stroke"
          />
          <set
            attributeName="stroke"
            to="hsl(195deg 80% 60%)"
            begin="stackOut.begin"
          />
          <set
            attributeName="fill"
            to="hsl(195deg 80% 5%)"
            begin="stackOut.begin"
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%); hsl(300deg 100% 40%); hsl(135deg 80% 60%)"
            begin="airdrop.begin"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animate
            attributeName="fill"
            values="hsl(195deg 80% 5%); hsl(300deg 100% 5%); hsl(135deg 80% 5%)"
            begin="airpath.end"
            dur=".33s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </g>
        <g id="balloon">
          <animateMotion begin="stackOut.begin" path="m200,-97.5" />
          <animateMotion
            dur="2s"
            begin="airlift.begin"
            path="m200,-87.5 v-50"
            {...animAttrs()}
          />
          <animateMotion
            dur="4s"
            begin="airpath.begin"
            path="m200,-137.5 v150"
            {...animAttrs()}
          />
          <animateMotion
            dur="1s"
            begin="airdrop.begin"
            path="m200,12.5 v-875"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines="1 0 1 1"
            fill="freeze"
          />
          <path
            id="balloonString"
            vectorEffect="non-scaling-stroke"
            stroke="var(--stroke-synapse)"
          >
            <set attributeName="d" to="m0,-111.8 v0" begin="stackOut.begin" />
            <animate
              attributeName="d"
              values="m0,-101.8 v0; m0,-111.8 v-37.5"
              begin="airlift.begin"
              dur="2s"
              {...animAttrs()}
            />
          </path>
          <circle
            vectorEffect="non-scaling-stroke"
            stroke="var(--stroke-synapse)"
            fill="var(--fill-synapse)"
          >
            <animate attributeName="r" values="0" begin="stackOut.begin" />
            <animate
              attributeName="r"
              values="0; 36"
              begin="airlift.begin"
              dur="2s"
              {...animAttrs()}
            />
            <animate
              attributeName="cy"
              values="-111.8; -186.8"
              begin="airlift.begin"
              dur="2s"
              {...animAttrs()}
            />
          </circle>
        </g>
        {/* <g
          id="stack-blue"
          fill="var(--fill-blue)"
          stroke="var(--stroke-blue)"
          transform="translate(0,-225)"
        >
          <use href="#box" transform="translate(0,-27.95)" />
          <use href="#box" transform="translate(25,12.5)" />
          <use href="#box" transform="translate(-25,12.5)" />
        </g> */}
        {/* <g
          id="stack-yellow"
          fill="var(--fill-yellow)"
          stroke="var(--stroke-yellow)"
          transform="translate(-400,0)"
        >
          <use href="#box" transform="translate(0,-27.95)" />
          <use href="#box" transform="translate(25,12.5)" />
          <use href="#box" transform="translate(-50,0)" />
          <use href="#box" transform="translate(-25,12.5)" />
          <use href="#box" transform="translate(0,25)" />
        </g> */}
        <g id="bargeCubeGroup" stroke="var(--stroke-blue)">
          <animateMotion path="m-25 -212.5" />
          <animate
            id="bargeCube"
            attributeName="opacity"
            begin="dockN.end"
            dur=".5s"
            values="0; 1"
            fill="freeze"
          />
          <animateMotion
            path="m-25 -212.5 0 -12.5"
            begin="dockN.end + .125s"
            dur=".5s"
            calcMode="spline"
            keyPoints="0; 1; 0"
            keyTimes="0; .5; 1"
            keySplines="0 0 .5 1; .8 0 .5 1"
            fill="freeze"
          />
          <animateMotion
            begin="bargeOut.end"
            dur="1s"
            path="m-25 -212.5 -150 -75"
            {...animAttrs()}
          />
          <animateMotion
            begin="bargeCross.begin"
            dur="2s"
            path="m-175 -287.5 -400 200"
            {...animAttrs()}
          />
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%); hsl(300deg 100% 40%); hsl(60deg 80% 60%)"
            begin="bargeCross.begin + .5s"
            dur=".5s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".33 0 .33 1; .33 0 .33 1"
            fill="freeze"
          />
          <animateMotion
            begin="bargeCross.end"
            dur="1s"
            path="m-575 -87.5 150 75"
            {...animAttrs()}
          />
          <path fill="var(--fill-blue)" vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
              begin="dockN.end"
              dur=".25s"
              {...animAttrs()}
            />
          </path>

          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
              begin="dockN.end"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
        </g>
        <g id="platformCubeBlueGroup" stroke="var(--stroke-blue)">
          <animateMotion path="m0 -200" />
          <animate
            id="platformCubeBlue"
            attributeName="opacity"
            begin="platformBlue.end"
            dur=".5s"
            values="0; 1"
          />
          <animateMotion
            path="m0 -200 0 -12.5"
            begin="platformCubeBlue.begin + .125s"
            dur=".5s"
            calcMode="spline"
            keyPoints="0; 1; 0"
            keyTimes="0; .5; 1"
            keySplines="0 0 .5 1; .8 0 .5 1"
          />
          <path fill="var(--fill-blue)" vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
              begin="platformCubeBlue.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
              begin="platformCubeBlue.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
        </g>

        <g id="bridgeCubeBlueGroup" stroke="var(--stroke-blue)">
          <animateMotion path="m-25 -187.5" />
          <animate
            id="bridgeCubeBlue"
            attributeName="opacity"
            begin="simpleBridgeNe.begin + .5s"
            dur=".5s"
            values="0; 1"
          />
          <animateMotion
            path="m-25 -187.5 0 -12.5"
            begin="bridgeCubeBlue.begin + .125s"
            dur=".5s"
            calcMode="spline"
            keyPoints="0; 1; 0"
            keyTimes="0; .5; 1"
            keySplines="0 0 .5 1; .8 0 .5 1"
          />
          <animateMotion
            id="bridgeCubeBlueOut"
            begin="bridgeCubeBlue.begin + 2s; bridgeCubeBlueIn.end + 2s"
            dur="1s"
            calcMode="spline"
            keyPoints="0; 1"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
          >
            <mpath href="#simpleBridgePath" />
          </animateMotion>
          <animate
            attributeName="stroke"
            values="hsl(195deg 80% 60%); hsl(300deg 100% 40%); hsl(60deg 80% 60%)"
            begin="bridgeCubeBlueOut.begin + .5s; bridgeCubeBlueIn.begin + .5s"
            dur=".5s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".33 0 .33 1; .33 0 .33 1"
            fill="freeze"
          />
          <animateMotion
            id="bridgeCubeBlueIn"
            begin="bridgeCubeBlueOut.end + 2s"
            dur="1s"
            calcMode="spline"
            keyPoints="1; 0"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
          >
            <mpath href="#simpleBridgePath" />
          </animateMotion>
          <path fill="var(--fill-blue)" vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
              begin="bridgeCubeBlue.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
              begin="bridgeCubeBlue.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
        </g>
        <g id="bridgeCubeYellowGroup" stroke="var(--stroke-yellow)">
          <animateMotion path="m-375 -12.5" />
          <animate
            id="bridgeCubeYellow"
            attributeName="opacity"
            begin="bridgeCubeBlue.begin + .5s"
            dur=".5s"
            values="0; 1"
          />
          <animateMotion
            path="m-375 -12.5 0 -12.5"
            begin="bridgeCubeYellow.begin + .125s"
            dur=".5s"
            calcMode="spline"
            keyPoints="0; 1; 0"
            keyTimes="0; .5; 1"
            keySplines="0 0 .5 1; .8 0 .5 1"
          />
          <animateMotion
            begin="bridgeCubeBlueOut.begin"
            dur="1s"
            calcMode="spline"
            keyPoints="1; 0"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
          >
            <mpath href="#simpleBridgePath" />
          </animateMotion>
          <animate
            attributeName="stroke"
            values="hsl(60deg 80% 60%); hsl(300deg 100% 40%); hsl(195deg 80% 60%)"
            begin="bridgeCubeBlueOut.begin + .5s; bridgeCubeBlueIn.begin + .5s"
            dur=".5s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".33 0 .33 1; .33 0 .33 1"
          />
          <animateMotion
            begin="bridgeCubeBlueIn.begin"
            dur="1s"
            calcMode="spline"
            keyPoints="0; 1"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
          >
            <mpath href="#simpleBridgePath" />
          </animateMotion>
          <path fill="var(--fill-yellow)" vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
              begin="bridgeCubeYellow.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
              begin="bridgeCubeYellow.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
        </g>
        <g id="platformCubeYellowGroup" stroke="var(--stroke-yellow)">
          <animateMotion path="m-400 0" />
          <animate
            id="platformCubeYellow"
            attributeName="opacity"
            begin="platformYellow.end"
            dur=".5s"
            values="0; 1"
          />
          <animateMotion
            path="m-400 0 0 -12.5"
            begin="platformCubeYellow.begin + .125s"
            dur=".5s"
            calcMode="spline"
            keyPoints="0; 1; 0"
            keyTimes="0; .5; 1"
            keySplines="0 0 .5 1; .8 0 .5 1"
          />
          <path fill="var(--fill-blue)" vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
              begin="platformCubeYellow.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
              begin="platformCubeYellow.begin"
              dur=".25s"
              {...animAttrs()}
            />
          </path>
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
