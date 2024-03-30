import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import exampleImg from '@assets/example.png'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import Hero from '../Hero'
import ValueProps from '../ValueProps'

import Wrapper from '@/components/WipWrapperComponents/Wrapper'

import styles from './keyframes.module.css'

const begin = {
  blue: {
    platform: 0,
    anchor: 0.5,
    bridge: 1,
    barge: 4,
    balloon: 6,
  },
  green: {
    platform: 0.1,
    anchor: 0.6,
    teleporter: 5,
  },
  orange: {
    platform: 0.2,
    anchor: 0.7,
    teleporter: 5.1,
  },
  yellow: {
    platform: 0.3,
    anchor: 0.8,
    bridge: 1.1,
  },
  synapse: {
    bridgeNw: 2,
    dockN: 2,
    dockE: 2,
    padN: 2,
    padE: 2,
  },
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

  const animAttrs = (x1 = 0.5, x2 = 0.2, y1 = 0, y2 = 1) => {
    return {
      calcMode: 'spline',
      keyTimes: '0; 1',
      keySplines: `${x1} ${y1} ${x2} ${y2}`,
      fill: 'freeze',
    }
  }

  const flashAttrs = (from, to, dur = '.4s') => {
    return {
      values: `${stroke[from]}; ${stroke.flash}; ${stroke[to]}`,
      dur: dur,
      calcMode: 'spline',
      keyTimes: `0; .5; 1`,
      keySplines: '.5 0 .2 1; .5 0 .2 1',
      fill: 'freeze',
    }
  }

  const AnimateFlash = ({
    hasStroke = true,
    hasFill = false,
    from,
    to,
    begin,
    dur = '.4s',
  }) => {
    return (
      <>
        {hasStroke && (
          <animate
            attributeName="stroke"
            begin={begin}
            dur={dur}
            values={`${stroke[from]}; ${stroke.flash}; ${stroke[to]}`}
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        )}
        {hasFill && (
          <animate
            attributeName="fill"
            begin={begin}
            dur={dur}
            values={`${fill[from]}; ${fill.synapse}; ${fill[to]}`}
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        )}
      </>
    )
  }

  const paint = (color) => {
    return {
      stroke: stroke[color],
      fill: fill[color],
    }
  }

  const hslStr = (h, s, l, a = undefined) =>
    `hsl(${h}deg ${s}% ${l}%${a === undefined ? '' : ` / ${a}%`})`

  const stroke = {
    inherit: 'inherit',
    synapse: hslStr(300, 100, 25),
    flash: hslStr(300, 100, 40),
    yellow: hslStr(60, 80, 60),
    orange: hslStr(25, 80, 60),
    blue: hslStr(195, 80, 60),
    green: hslStr(135, 80, 60),
  }
  const fill = {
    inherit: 'inherit',
    synapse: hslStr(300, 100, 5),
    yellow: hslStr(60, 30, 3),
    orange: hslStr(25, 30, 3),
    blue: hslStr(195, 30, 3),
    green: hslStr(135, 30, 3),
  }

  const AnimateHop = ({ begin }) => {
    return (
      <animateMotion
        path="m0 0 0 -12.5"
        additive="sum"
        begin={begin}
        dur=".5s"
        calcMode="spline"
        keyPoints="0; 1; 0"
        keyTimes="0; .5; 1"
        keySplines="0 0 .5 1; .8 0 .5 1"
      />
    )
  }

  const Cube = ({
    color = 'synapse',
    translate,
    begin = 0,
    children,
  }: {
    color?: string
    translate?: string
    begin?: number
    children?: React.ReactNode
  }) => {
    return (
      <g
        transform={translate ? `translate(${translate})` : undefined}
        stroke={stroke[color]}
      >
        <animate
          attributeName="opacity"
          begin={begin + 's'}
          dur=".5s"
          values="0; 1"
          fill="freeze"
        />
        <animateMotion
          path="m0 0 v-12.5"
          additive="sum"
          begin={begin + 0.125 + 's'}
          dur=".5s"
          calcMode="spline"
          keyPoints="0; 1; 0"
          keyTimes="0; .5; 1"
          keySplines="0 0 .5 1; .8 0 .5 1"
          fill="freeze"
        />
        {children}
        <path fill={fill[color]} vectorEffect="non-scaling-stroke">
          <animate
            attributeName="d"
            values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
            begin={begin + 's'}
            dur=".25s"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
        </path>
        <path vectorEffect="non-scaling-stroke">
          <animate
            attributeName="d"
            values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
            begin={begin + 's'}
            dur=".25s"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
          />
        </path>
      </g>
    )
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
        </style>
        <defs></defs>

        <path {...paint('synapse')}>
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
          {...paint('synapse')}
        />
        <path {...paint('synapse')}>
          <animate
            id="dockN"
            attributeName="d"
            values="m-40,-240 0,0 0,0 0,0z; m-80,-300 40,20 -120,60 -40,-20z"
            dur=".5s"
            begin="simpleBridgeNe.end + 2s"
            {...animAttrs()}
          />
        </path>
        <path {...paint('synapse')}>
          <animate
            id="dockE"
            attributeName="d"
            values="m-440,-40 0,0 0,0 0,0z; m-480,-100 40,20 -120,60 -40,-20z"
            dur=".5s"
            begin="simpleBridgeNe.end + 2.5s"
            {...animAttrs()}
          />
        </path>

        <path {...paint('synapse')}>
          <animate
            id="airpad1"
            attributeName="d"
            dur=".5s"
            values="m200,-170 0,0 -0,0 -0,-0z; m200,-200 60,30 -60,30 -60,-30z"
            begin={begin.blue.balloon}
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="airpad1.begin"
          />
          <animateTransform
            attributeName="transform"
            additive="sum"
            type="translate"
            values="0 30; 0 0"
            dur=".5s"
            begin="airpad1.begin"
            {...animAttrs()}
          />
        </path>
        <path {...paint('synapse')}>
          <animate
            id="airpad2"
            attributeName="d"
            dur=".5s"
            values="m200,-30 0,0 -0,0 -0,-0z; m200,-60 60,30 -60,30 -60,-30z"
            begin={begin.blue.balloon}
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="airpad2.begin"
          />
          <animateTransform
            attributeName="transform"
            additive="sum"
            type="translate"
            values="0 30; 0 0"
            dur=".5s"
            begin="airpad2.begin"
            {...animAttrs()}
          />
        </path>

        <path {...paint('blue')}>
          <animate
            id="platformBlue"
            attributeName="d"
            values="m0,1 2,1 -2,1 -2,-1z; m0,-100 200,100 -200,100 -200,-100z"
            dur=".25s"
            begin={begin.blue.platform + 's'}
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
            type="translate"
            values="0 -150; 0 -200"
            dur=".5s"
            begin="platformBlue.begin"
            {...animAttrs()}
          />
        </path>
        <path {...paint('green')}>
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
            type="translate"
            values="400 50; 400 0"
            dur=".5s"
            begin="platformGreen.begin"
            {...animAttrs()}
          />
        </path>
        <path {...paint('orange')}>
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
            type="translate"
            values="0 250; 0 200"
            dur=".5s"
            begin="platformOrange.begin"
            {...animAttrs()}
          />
        </path>
        <path {...paint('yellow')}>
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
            type="translate"
            values="-400 50; -400 0"
            dur=".5s"
            begin="platformYellow.begin"
            {...animAttrs()}
          />
        </path>

        <g id="barge">
          <animateMotion path="M200,-500" />
          <path d="m70,-75 100,50 -200,100 -100,-50z" {...paint('synapse')} />
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

        <g stroke={stroke.orange}>
          <ellipse rx="30" ry="15" cy="260" {...paint('synapse')} />
          <Cube color="orange" begin={0}>
            <AnimateFlash
              from="orange"
              to="green"
              begin="teleporterBeamsOut.begin + 1s"
              dur="3s"
            />
            <AnimateFlash
              from="green"
              to="orange"
              begin="teleporterBeamsIn.begin + 1s"
              dur="3s"
            />
            <animateMotion
              dur="2s"
              path="m0,240 v12.5"
              calcMode="spline"
              keyPoints="0; 1; 0"
              keyTimes="0; 0.5; 1"
              keySplines=".33 0 .67 1; .33 0 .67 1"
              repeatCount="indefinite"
            />
          </Cube>
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
            stroke={stroke.synapse}
            fill={fill.synapse}
          />
        </g>

        <g stroke={stroke.green} transform="translate(520,-260)">
          <ellipse rx="30" ry="15" cy="260" {...paint('synapse')} />
          <Cube color="green">
            <animate
              attributeName="stroke"
              begin="teleporterBeamsOut.begin + 1s"
              {...flashAttrs('inherit', 'orange', '3s')}
            />
            <animate
              attributeName="stroke"
              begin="teleporterBeamsIn.begin + 1s"
              {...flashAttrs('orange', 'inherit', '3s')}
            />
            <animateMotion
              dur="2s"
              path="m0,240 v12.5"
              calcMode="spline"
              keyPoints="0; 1; 0"
              keyTimes="0; .5; 1"
              keySplines=".33 0 .67 1; .33 0 .67 1"
              repeatCount="indefinite"
            />
          </Cube>
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
            stroke={stroke.synapse}
            fill={fill.synapse}
          />
        </g>

        <Cube color="green" translate="400 0" begin={begin.green.anchor} />
        <Cube color="green" translate="375 12.5" begin={10.5} />

        <Cube color="orange" translate="0 150" begin={begin.orange.anchor} />
        <Cube color="orange" translate="25 162.5" begin={11.5} />
        <Cube color="orange" translate="-25 162.5" begin={13.5} />
        <Cube color="orange" translate="0 122.05" begin={20} />

        <Cube color="blue" begin={begin.blue.balloon}>
          <animateTransform
            id="balloonReset"
            attributeName="transform"
            type="translate"
            to="100 -225"
            begin="0s; balloonFadeOut.end"
            dur="1ms"
            {...animAttrs()}
          />
          <animateTransform
            id="stackOut"
            attributeName="transform"
            type="translate"
            by="100 50"
            begin={`${begin.blue.balloon + 2}s; balloonFadeIn.end`}
            dur="1s"
            {...animAttrs()}
          />
          <animateTransform
            id="airlift"
            attributeName="transform"
            type="translate"
            by="0 -50"
            begin="balloon.begin + .1s"
            dur="2s"
            {...animAttrs()}
          />
          <animateTransform
            id="airpath"
            attributeName="transform"
            type="translate"
            by="0 150"
            begin="airlift.end"
            dur="4s"
            {...animAttrs()}
          />
          <animateTransform
            id="airdrop"
            attributeName="transform"
            type="translate"
            by="0 40"
            begin="airpath.end"
            dur=".25s"
            {...animAttrs(0.33, 1)}
          />
          <animate
            attributeName="stroke"
            begin="airdrop.begin"
            dur=".33s"
            {...flashAttrs('inherit', 'green')}
          />
          <animateTransform
            id="airdropOut"
            attributeName="transform"
            type="translate"
            by="137.5 68.75"
            begin="airdrop.end + .1s"
            dur=".75s"
            {...animAttrs()}
          />
          <animateTransform
            id="ricochet"
            attributeName="transform"
            type="translate"
            by="-287.5 141"
            begin="airdropOut.end"
            dur=".75s"
            {...animAttrs()}
          />
          <animate
            attributeName="stroke"
            begin="ricochet.begin + .25s"
            dur=".33s"
            {...flashAttrs('green', 'orange')}
          />
          <animate
            id="balloonFadeOut"
            attributeName="opacity"
            to="0"
            dur="3s"
            begin="ricochet.end + 1s"
          />

          <animate
            id="balloonFadeIn"
            attributeName="opacity"
            values="0; 1"
            dur="3s"
            begin="balloonFadeOut.end"
          />
          <set
            attributeName="stroke"
            to={stroke.blue}
            begin="balloonFadeIn.begin"
          />
          <set
            attributeName="fill"
            to={fill.blue}
            begin="balloonFadeIn.begin"
          />

          {/* <animate
            attributeName="fill"
            begin="airpath.end"
            // {...flashAttrs('blue', 'green')} // TODO: 'fill' version
            dur=".33s"
            values="hsl(195deg 80% 5%); hsl(300deg 100% 5%); hsl(135deg 80% 5%)"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          /> */}
        </Cube>

        <g>
          <animateMotion begin="stackOut.begin" path="m200,-97.5" />
          <animateMotion
            id="balloon"
            dur="2s"
            begin="stackOut.end"
            path="m200 -87.5 v-50"
            {...animAttrs()}
          />
          <animateMotion
            dur="4s"
            additive="sum"
            begin="airpath.begin"
            path="m0 0 v150"
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
            stroke={stroke.synapse}
          >
            <set attributeName="d" to="m0,-111.8 v0" begin="stackOut.begin" />
            <animate
              attributeName="d"
              values="m0,-111.8 v0; m0,-111.8 v-37.5"
              begin="balloon.begin"
              dur="2s"
              {...animAttrs()}
            />
          </path>
          <circle
            vectorEffect="non-scaling-stroke"
            stroke={stroke.synapse}
            fill={fill.synapse}
          >
            <animate attributeName="r" values="0" begin="stackOut.begin" />
            <animate
              attributeName="r"
              values="0; 36"
              begin="balloon.begin"
              dur="2s"
              {...animAttrs()}
            />
            <animate
              attributeName="cy"
              values="-111.8; -186.8"
              begin="balloon.begin"
              dur="2s"
              {...animAttrs()}
            />
          </circle>
        </g>

        <Cube color="blue" translate="-25 -212.5" begin={begin.blue.barge}>
          <set attributeName="stroke" to={stroke.blue} begin="bargeOut.begin" />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeOut.end"
            dur="1s"
            by="-162.5 -81.25"
            {...animAttrs()}
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeCross.begin"
            dur="2s"
            by="-400 200"
            {...animAttrs()}
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeCross.end"
            dur="1s"
            by="137.5 68.75"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            begin="bargeCross.end + 1s"
            to="0"
            dur="2s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            begin="bargeOut.begin"
            to="1"
            dur="2s"
            fill="freeze"
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeOut.begin"
            dur="1ms"
            to="-25 -212.5"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            begin="bargeCross.begin + .5s"
            dur=".5s"
            {...flashAttrs('inherit', 'yellow')}
          />
        </Cube>

        <Cube color="blue" translate="0 -200" begin={begin.blue.anchor} />

        <Cube color="blue" translate="-25 -187.5" begin={begin.blue.bridge}>
          <animateTransform
            id="bridgeCubeOut"
            attributeName="transform"
            type="translate"
            from="-25 -187.5"
            by="-350 175"
            begin="3s; bridgeCubeIn.end + 2s"
            dur="1s"
            {...animAttrs()}
          />
          <animateTransform
            id="bridgeCubeIn"
            attributeName="transform"
            type="translate"
            by="350 -175"
            begin="bridgeCubeOut.end + 2s"
            dur="1s"
            {...animAttrs()}
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeOut.begin + .3s"
            dur=".5s"
            {...flashAttrs('blue', 'yellow')}
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeIn.begin + .3s"
            dur=".5s"
            {...flashAttrs('yellow', 'blue')}
          />
        </Cube>
        <Cube color="yellow" translate="-375 -12.5" begin={begin.yellow.bridge}>
          <animateTransform
            attributeName="transform"
            type="translate"
            from="-25 -187.5"
            by="-350 175"
            begin="bridgeCubeIn.begin"
            dur="1s"
            {...animAttrs()}
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            by="350 -175"
            begin="bridgeCubeOut.begin"
            dur="1s"
            {...animAttrs()}
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeOut.begin + .3s"
            dur=".5s"
            {...flashAttrs('yellow', 'blue')}
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeIn.begin + .3s"
            dur=".5s"
            {...flashAttrs('blue', 'yellow')}
          />
        </Cube>

        <Cube color="yellow" translate="-400 0" begin={begin.yellow.anchor} />

        <Cube color="yellow" translate="-450 0" begin={15} />
        <Cube color="yellow" translate="-425 12.5" begin={16} />
        <Cube color="yellow" translate="-425 -27.95" begin={18} />

        <Cube color="blue" translate="25 -187.5" begin={17} />
        <Cube color="blue" translate="0 -175" begin={19} />

        {/* <rect
          width="40"
          height="200"
          transform-origin="20 100"
          vectorEffect="non-scaling-stroke"
          transform="translate(-20 -100) matrix(1 .5 -1 .5 0 0) rotate(90)"
        /> */}
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
