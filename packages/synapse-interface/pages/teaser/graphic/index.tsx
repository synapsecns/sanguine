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

  const paint = (color = 'synapse') => {
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

  const Platform = ({
    id,
    color,
    translate,
    begin = '0s',
  }: {
    id: string
    color?: string
    translate?: string
    begin?: string
  }) => {
    return (
      <path
        transform={translate ? `translate(${translate})` : null}
        {...paint(color)}
      >
        <animate
          id={id}
          attributeName="d"
          values="m0 0 0 0 0 0 0 0z; m0 -100 200 100 -200 100 -200 -100z"
          dur=".25s"
          begin={begin}
          {...animAttrs()}
        />
        <animateMotion
          path="M0 50 0 0"
          dur=".5s"
          begin={`${id}.begin`}
          {...animAttrs()}
        />
        <animate
          attributeName="opacity"
          values="0;1"
          repeatCount="3"
          dur=".1s"
          begin={`${id}.begin + .1s`}
        />
      </path>
    )
  }

  const Cube = ({
    color = 'synapse',
    translate,
    begin = 0,
    restart,
    children,
  }: {
    color?: string
    translate?: string
    begin?: number | string
    restart?: string
    children?: React.ReactNode
  }) => {
    if (typeof begin === 'number') begin = begin + 's'
    return (
      <g
        transform={translate ? `translate(${translate})` : undefined}
        stroke={stroke[color]}
      >
        <animate
          attributeName="opacity"
          begin={begin}
          dur=".375s"
          values=".25; 1"
          fill="freeze"
          restart={restart}
        />
        <animateMotion
          path="m0 0 v-12.5"
          additive="sum"
          begin={begin}
          dur=".5s"
          calcMode="spline"
          keyPoints="0; 1; 0"
          keyTimes="0; .5; 1"
          keySplines="0 0 .5 1; .8 0 .5 1"
          fill="freeze"
          restart={restart}
        />
        {children}
        <path fill={fill[color]} vectorEffect="non-scaling-stroke">
          <animate
            attributeName="d"
            values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
            begin={begin}
            dur=".25s"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
            restart={restart}
          />
        </path>
        <path vectorEffect="non-scaling-stroke">
          <animate
            attributeName="d"
            values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
            begin={begin}
            dur=".25s"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
            restart={restart}
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
        <defs> </defs>

        <path {...paint('synapse')}>
          <animate
            id="bridgeNw"
            attributeName="d"
            values="m-200,-100 0,0 0,0 0,0z; m-100,-150 0,0 -200,100 0,0z; m-120,-160 40,20 -200,100 -40,-20z"
            dur=".5s"
            begin="platformYellow.end + 1s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </path>
        <path transform="translate(400 200)" {...paint('synapse')}>
          <animate
            id="bridgeSe"
            attributeName="d"
            values="m-200,-100 0,0 0,0 0,0z; m-100,-150 0,0 -200,100 0,0z; m-120,-160 40,20 -200,100 -40,-20z"
            dur=".5s"
            begin="airDropOut.end + 1s"
            restart="never"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </path>
        <path transform="translate(-40 -240)" {...paint('synapse')}>
          <animate
            id="dockN"
            attributeName="d"
            values="m0 0 0 0 0 0 0 0z; m-40 -60 40 20 -120 60 -40 -20z"
            dur=".5s"
            begin="bridgeNw.end + 3.5s"
            {...animAttrs()}
          />
        </path>
        <path transform="translate(-440 -40)" {...paint('synapse')}>
          <animate
            id="dockW"
            attributeName="d"
            values="m0 0 0 0 0 0 0 0z; m-40 -60 40 20 -120 60 -40 -20z"
            dur=".5s"
            begin="dockN.begin + 1s"
            {...animAttrs()}
          />
        </path>

        <path transform="translate(200 -200)" {...paint('synapse')}>
          <animate
            id="airpadN"
            attributeName="d"
            dur=".5s"
            values="m0 0 0 0 0 0 0 0z; m0 -30 60 30 -60 30 -60 -30z"
            begin="bargeCross.end"
            restart="never"
            {...animAttrs()}
          />
          <animateMotion
            path="m0 30 0 0"
            dur=".5s"
            begin="airpadN.begin"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="airpadN.begin + .1s"
          />
        </path>
        <path transform="translate(200 -60)" {...paint('synapse')}>
          <animate
            id="airpadE"
            attributeName="d"
            dur=".5s"
            values="m0 0 0 0 0 0 0 0z; m0 -30 60 30 -60 30 -60 -30z"
            begin="airpadN.begin + .5s"
            {...animAttrs()}
          />

          <animateMotion
            path="m0 30 0 0"
            dur=".5s"
            begin="airpadE.begin"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="airpadE.begin + .1s"
          />
        </path>

        <Platform
          id="platformBlue"
          translate="0 -200"
          color="blue"
          begin="0s"
        />
        <Platform
          id="platformGreen"
          translate="400 0"
          color="green"
          begin=".1s"
        />
        <Platform
          id="platformOrange"
          translate="0 200"
          color="orange"
          begin=".2s"
        />
        <Platform
          id="platformYellow"
          translate="-400 0"
          color="yellow"
          begin=".3s"
        />

        <path
          id="barge"
          d="m50,-75 100,50 -200,100 -100,-50z"
          transform="translate(250 -515)"
          {...paint('synapse')}
        >
          <animateMotion
            id="bargeOut"
            dur="2s"
            begin="dockN.end; bargeIn.end"
            path="m0 0 -440 220"
            {...animAttrs()}
          />
          <animateMotion
            id="bargeCross"
            dur="2s"
            begin="bargeOut.end + 2s"
            path="m0 0 -400 200"
            additive="sum"
            {...animAttrs()}
          />
          <animateMotion
            id="bargeIn"
            dur="2s"
            begin="bargeCross.end + 2s"
            path="m0 0 -400 200"
            additive="sum"
            {...animAttrs()}
          />
        </path>

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

        <Cube color="green" translate="400 0" begin=".6s" />
        <Cube
          color="green"
          translate="375 12.5"
          begin="airDropOut.end + .5s"
          restart="never"
        />

        <Cube color="orange" translate="0 150" begin=".7s" />
        <Cube color="orange" translate="25 162.5" begin={20} />
        <Cube color="orange" translate="-25 162.5" begin={21} />
        <Cube color="orange" translate="0 122.05" begin={22} />

        <Cube
          translate="100 -225"
          color="blue"
          begin="airpadN.end; airDropOut.end + 1ms"
        >
          <animateMotion
            id="airCubeOut"
            path="m0 0 100 50"
            begin="airpadE.end + 2s; airDropOut.end + 2s"
            dur="1s"
            {...animAttrs()}
          />
          <animateMotion
            id="airLift"
            additive="sum"
            path="m0 0 v-50 150"
            begin="airCubeOut.end"
            keyPoints="0; .25; 1"
            keyTimes="0; .33; 1"
            calcMode="spline"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            dur="6s"
            fill="freeze"
          />
          <animateMotion
            id="airDrop"
            additive="sum"
            path="m0 0 v40"
            begin="airLift.end"
            dur=".25s"
            {...animAttrs(0.33, 1)}
          />
          <animate
            attributeName="stroke"
            begin="airDrop.begin"
            dur=".33s"
            {...flashAttrs('inherit', 'green')}
          />
          <animateMotion
            id="airDropOut"
            additive="sum"
            path="m0 0 137.5 68.75 12.5 -8.75"
            begin="airDrop.end"
            dur=".75s"
            {...animAttrs()}
          />
          <set attributeName="opacity" to="0" begin="airDropOut.end" />
          <set attributeName="stroke" to={stroke.blue} begin="airDropOut.end" />
          <animateMotion path="m0 0" begin="airDropOut.end" />
        </Cube>

        <g id="balloon" transform="translate(200 -200)">
          <animateMotion
            dur="6s"
            begin="airLift.begin"
            path="m0 0 v-50 150"
            keyPoints="0; .25; 1"
            keyTimes="0; .33; 1"
            calcMode="spline"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            dur="1s"
            begin="airDrop.begin"
            additive="sum"
            path="m0 0 v-400"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines="1 0 1 1"
            fill="freeze"
          />
          <path stroke={stroke.synapse}>
            <animate
              attributeName="d"
              values="m0 -2 v0; m0 -50 v48"
              begin="airLift.begin"
              dur="2s"
              {...animAttrs()}
            />
          </path>
          <circle {...paint('synapse')}>
            <animate
              attributeName="r"
              values="0; 36"
              begin="airLift.begin"
              dur="2s"
              {...animAttrs()}
            />
            <animateMotion
              path="m0 0 v-86"
              begin="airLift.begin"
              dur="2s"
              {...animAttrs()}
            />
          </circle>
        </g>

        <g transform="translate(350 25)" stroke={stroke.green} opacity="0">
          <set
            id="airBridge"
            attributeName="stroke"
            to={stroke.green}
            begin="airDropOut.end"
          />
          <set
            id="airBridge"
            attributeName="opacity"
            to="1"
            begin="airDropOut.end"
          />
          <animateMotion begin="airBridge.begin" path="m0 0" />
          <animateMotion
            dur="1s"
            begin="airBridge.begin + 2s"
            path="m0 0 -325 162.5 -25 -12.5"
            {...animAttrs()}
          />
          <AnimateFlash
            from="green"
            to="orange"
            begin="airBridge.begin + 2.3s"
          />
          <animate
            attributeName="opacity"
            values="1; 0"
            dur="4s"
            begin="airBridge.begin + 4s"
            {...animAttrs()}
          />
          <path
            fill={fill.green}
            d="m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
            vectorEffect="non-scaling-stroke"
          />
          <path
            d="m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
            vectorEffect="non-scaling-stroke"
          />
        </g>

        <Cube color="blue" translate="-25 -212.5" begin="bargeOut.begin">
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

        <Cube color="blue" translate="0 -200" begin=".5s" />

        {/* Simple Bridge Blue/Yellow Swap */}

        <Cube color="blue" translate="-25 -187.5" begin="bridgeNw.end + .5s">
          <animateMotion
            id="bridgeCubeOut"
            dur="4s"
            begin="bridgeNw.end + 2.5s; bridgeCubeOut.end + 2s"
            path="M0 0 -350 175"
            calcMode="spline"
            keyPoints="0; 1; 1; 0"
            keyTimes="0; .25; .75; 1"
            keySplines=".5 0 .2 1; 0 0 1 1; .5 0 .2 1"
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${stroke.blue}; ${stroke.synapse}; ${stroke.yellow}; ${stroke.yellow}; ${stroke.synapse}; ${stroke.blue}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
          <animate
            attributeName="fill"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${fill.blue}; ${fill.synapse}; ${fill.yellow}; ${fill.yellow}; ${fill.synapse}; ${fill.blue}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
        </Cube>

        <Cube color="yellow" translate="-375 -12.5" begin="bridgeNw.end + 1s">
          <animateMotion
            dur="4s"
            begin="bridgeNw.end + 2.5s; bridgeCubeOut.end + 2s"
            path="M0 0 350 -175"
            calcMode="spline"
            keyPoints="0; 1; 1; 0"
            keyTimes="0; .25; .75; 1"
            keySplines=".5 0 .2 1; 0 0 1 1; .5 0 .2 1"
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${stroke.yellow}; ${stroke.synapse}; ${stroke.blue}; ${stroke.blue}; ${stroke.synapse}; ${stroke.yellow}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
          <animate
            attributeName="fill"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${fill.yellow}; ${fill.synapse}; ${fill.blue}; ${fill.blue}; ${fill.synapse}; ${fill.yellow}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
        </Cube>

        <Cube color="yellow" translate="-400 0" begin=".8s" />

        <Cube color="yellow" translate="-450 0" begin={15} />
        <Cube color="yellow" translate="-425 12.5" begin={16} />
        <Cube color="yellow" translate="-425 -27.95" begin={18} />

        <Cube color="blue" translate="25 -187.5" begin={17} />
        <Cube color="blue" translate="0 -175" begin={19} />
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
