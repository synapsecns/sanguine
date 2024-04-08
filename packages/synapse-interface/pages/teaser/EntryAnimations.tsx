export function InstructionsSvg() {
  return (
    <PlatformSvg id="instructions">
      <path
        d="m-25 -15 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="instructionsEnter.begin + .3s"
          dur="1s"
          path="m0 0 0 -10"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="instructionsLeave.begin + .3s"
          dur="1s"
          path="m0 -10 0 10"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
      <path
        d="m-25 -25 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="instructionsEnter.begin + .2s"
          dur="1s"
          path="m0 0 0 -15"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="instructionsLeave.begin + .2s"
          dur="1s"
          path="m0 -15 0 15"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
      <path
        d="m-25 -35 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="instructionsEnter.begin + .1s"
          dur="1s"
          path="m0 0 0 -20"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="instructionsLeave.begin + .1s"
          dur="1s"
          path="m0 -20 0 20"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
      <path
        d="m-25 -45 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="instructionsEnter.begin + .0s"
          dur="1s"
          path="m0 0 0 -25"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="instructionsLeave.begin + .0s"
          dur="1s"
          path="m0 -25 0 25"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
    </PlatformSvg>
  )
}

export function BatterySvg() {
  return (
    <PlatformSvg id="batteries">
      <ellipse rx="25" ry="12.5" className="stroke-zinc-700 fill-zinc-400/5" />
      <ellipse
        rx="25"
        ry="12.5"
        cy="-50"
        className="stroke-zinc-700 fill-zinc-400/5"
      />
      <ellipse
        rx="6.25"
        ry="3.125"
        cy="-50"
        className="stroke-zinc-700 fill-zinc-400/5"
      />
      <ellipse
        rx="6.25"
        ry="3.125"
        cy="-56.125"
        className="stroke-zinc-700 fill-zinc-400/5"
      />
      <path d="m-25 0 v-50 m50 0 v50" className="stroke-zinc-700" />
      <ellipse
        rx="25"
        ry="12.5"
        cy="-7.5"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="batteriesEnter.begin + .3s"
          dur=".7s"
          path="m0 0 0 -42.5"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines="0 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="batteriesLeave.begin + .3s"
          dur=".7s"
          path="m0 -42.5 0 42.5"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </ellipse>
    </PlatformSvg>
  )
}

export function AssemblySvg() {
  return (
    <PlatformSvg id="assembly">
      {/* Bottom */}
      <path
        d="m-25 0 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="assemblyEnter.begin"
          dur="1s"
          path="m0 0 0 -5.5125"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="assemblyLeave.begin"
          dur="1s"
          path="m0 -5.5125 0 5.5125"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>

      {/* Front-left */}
      <path
        d="m-25 0 v-27.95 l25 12.5 v27.95z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="assemblyEnter.begin"
          dur="1s"
          path="m0 0 -25 -1.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="assemblyLeave.begin"
          dur="1s"
          path="m-25 -1.475 25 1.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>

      {/* Rear-left */}
      <path
        d="m-25 0 v-27.95 l25 -12.5 v27.95z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="assemblyEnter.begin"
          dur="1s"
          path="m0 0 -25 -26.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="assemblyLeave.begin"
          dur="1s"
          path="m-25 -26.475 25 26.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>

      {/* Front-right */}
      <path
        d="m0 12.5 v-27.95 l25 -12.5 v27.95z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="assemblyEnter.begin"
          dur="1s"
          path="m0 0 25 -1.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="assemblyLeave.begin"
          dur="1s"
          path="m25 -1.475 -25 1.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>

      {/* Rear-right */}
      <path
        d="m0 -12.5 v-27.95 l25 12.5 v27.95z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="assemblyEnter.begin"
          dur="1s"
          path="m0 0 25 -26.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="assemblyLeave.begin"
          dur="1s"
          path="m25 -26.475 -25 26.475"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>

      {/* Top */}
      <path
        d="m-25 -27.95 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5"
      >
        <animateMotion
          begin="assemblyEnter.begin"
          dur="1s"
          path="m0 0 0 -34.9375"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          begin="assemblyLeave.begin"
          dur="1s"
          path="m0 -34.9375 0 34.9375"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
    </PlatformSvg>
  )
}

export function IntegratedSvg() {
  return <PlatformSvg id="integrated" />
}

function PlatformSvg({
  id,
  children,
}: {
  id: string
  children?: React.ReactNode
}) {
  return (
    <svg
      viewBox="-100 -50 200 100"
      className="stroke-zinc-800 hover:stroke-zinc-700 fill-zinc-500/5 hover:fill-zinc-500/10 transition-all duration-500 p-8 pb-0 cursor-pointer"
      overflow="visible"
    >
      <animate id={`${id}Enter`} begin="mouseenter" />
      <animate id={`${id}Leave`} begin="mouseleave" />
      {/* Platform */}
      <path d="m-100 0 100 -50 100 50 -100 50z" />
      {children}
    </svg>
  )
}

export default function ({ id, children }) {
  return <PlatformSvg id="id">{children}</PlatformSvg>
}