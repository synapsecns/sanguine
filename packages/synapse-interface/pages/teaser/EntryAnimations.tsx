export function InstructionsSvg() {
  return (
    <PlatformSvg id="instructions">
      <path
        d="m-25 -15 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5 hover:stroke-blue-700 hover:fill-blue-400/5 transition-colors duration-250"
      >
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 -10"
          begin="instructionsEnter.begin + .3s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 0"
          begin="instructionsLeave.begin + .3s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
      <path
        d="m-25 -25 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5 hover:stroke-green-700 hover:fill-green-400/5 transition-colors duration-250"
      >
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 -15"
          begin="instructionsEnter.begin + .2s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 0"
          begin="instructionsLeave.begin + .2s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
      <path
        d="m-25 -35 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5 hover:stroke-yellow-700 hover:fill-yellow-400/5 transition-colors duration-250"
      >
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 -20"
          begin="instructionsEnter.begin + .1s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 0"
          begin="instructionsLeave.begin + .1s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
      <path
        d="m-25 -45 25 -12.5 25 12.5 -25 12.5z"
        className="stroke-zinc-700 fill-zinc-400/5 hover:stroke-red-700 hover:fill-red-400/5 transition-colors duration-250"
      >
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 -25"
          begin="instructionsEnter.begin + .0s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 0"
          begin="instructionsLeave.begin + .0s"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
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
        cy="-10"
        className="stroke-zinc-700 fill-emerald-400/10 group-hover:fill-green-400/10"
      >
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 -35"
          begin="batteriesEnter.begin"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateTransform
          attributeName="transform"
          type="translate"
          to="0 0"
          begin="batteriesLeave.begin"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animateMotion
          dur="3s"
          path="m0 3 0 -3"
          keyPoints="0; 1; 0"
          keyTimes="0; .5; 1"
          calcMode="spline"
          keySplines=".25 0 .75 1; .25 0 .75 1"
          repeatCount="indefinite"
        />
        <animate
          attributeName="ry"
          values="12; 13; 12"
          keyTimes="0; .5; 1"
          calcMode="spline"
          keySplines=".25 0 .75 1; .25 0 .75 1"
          dur="4s"
          repeatCount="indefinite"
        />
      </ellipse>
      <path
        d="m-25 0 v-0 a25 12.5 0 1 1 50 0 v0 a25 12.5 0 1 1 -50 0z"
        className="fill-emerald-400/5 group-hover:fill-green-400/10 stroke-none"
      >
        <animate
          attributeName="d"
          to="m-25 0 v-44 a25 12.5 0 1 1 50 0 v44 a25 12.5 0 1 1 -50 0z"
          dur="1s"
          begin="batteriesEnter.begin"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          attributeName="d"
          to="m-25 0 v-0 a25 12.5 0 1 1 50 0 v0 a25 12.5 0 1 1 -50 0z"
          dur="1s"
          begin="batteriesLeave.begin"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </path>
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
      className="group stroke-zinc-800 hover:stroke-zinc-700 fill-zinc-500/5 hover:fill-zinc-500/10 transition-all duration-500 p-8 pb-0 cursor-pointer"
      overflow="visible"
    >
      <animate id={`${id}Enter`} begin="mouseenter" />
      <animate id={`${id}Leave`} begin="mouseleave" />
      {/* Platform */}
      <path
        d="m-100 0 100 -50 100 50 -100 50z"
        className="group-hover:translate-y-4 transition-all duration-1000"
      />
      {children}
    </svg>
  )
}

export default function ({ id, children }) {
  return <PlatformSvg id="id">{children}</PlatformSvg>
}