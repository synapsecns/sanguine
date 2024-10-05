export default () => {
  return (
    <svg
      id="synapseAnimatedLogo"
      width="64"
      height="96"
      viewBox="-24 -24 48 48"
      overflow="visible"
    >
      <defs>
        <linearGradient id="synGrad" fx="0%">
          <stop offset="0%" stopColor="hsl(285deg 100% 65%)" />
          <stop offset="100%" stopColor="hsl(265deg 100% 75%)" />
        </linearGradient>
      </defs>
      <path
        d="M0 18 18 0 -18 0 0 -18"
        stroke="url(#synGrad)"
        strokeWidth="4"
        strokeLinejoin="bevel"
        fill="none"
        opacity=".5"
        pathLength="1"
        strokeDasharray="1"
        strokeDashoffset="1"
      >
        <animate
          id="synapseLogoPath"
          attributeName="stroke-dashoffset"
          to="0"
          dur=".9s"
          fill="freeze"
          begin="0s; synapseAnimatedLogo.mouseenter; synapseAnimatedLogo.mousedown"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".67 0 .8 1"
        />
        <animate
          attributeName="stroke-width"
          to="5.5"
          dur=".9s"
          begin="synapseLogoPath.begin"
          fill="freeze"
          calcMode="linear"
          keyTimes="0; 1"
          keySplines=".67 0 .8 1"
        />
        <animate
          attributeName="opacity"
          values="0; 1"
          repeatCount="5"
          dur=".1s"
          begin="synapseLogoPath.begin"
        />
      </path>
      <g stroke="url(#synGrad)" fill="none" opacity="0">
        <animate
          attributeName="opacity"
          to=".33"
          dur="1s"
          begin="synapseLogoPath.begin"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
        />
        <circle cy="18" r="6" />
        <circle cx="18" r="6" />
        <circle cx="-18" r="6" />
        <circle cy="-18" r="6" />
      </g>
      <g fill="url(#synGrad)">
        <circle cy="18" r="0">
          <set attributeName="r" to="0" begin="synapseLogoPath.begin" />
          <animate
            attributeName="r"
            values="0; 6"
            begin="synapseLogoPath.begin + .05s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="3"
            begin="synapseLogoPath.begin + .05s"
            dur=".1s"
          />
        </circle>
        <circle cx="18" r="0">
          <set attributeName="r" to="0" begin="synapseLogoPath.begin" />
          <animate
            attributeName="r"
            values="0; 6"
            begin="synapseLogoPath.begin + .3s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="4"
            begin="synapseLogoPath.begin + .3s"
            dur=".1s"
          />
        </circle>
        <circle cx="-18" r="0">
          <set attributeName="r" to="0" begin="synapseLogoPath.begin" />
          <animate
            attributeName="r"
            values="0; 6"
            begin="synapseLogoPath.begin + .55s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="4"
            begin="synapseLogoPath.begin + .55s"
            dur=".1s"
          />
        </circle>
        <circle cy="-18" r="0">
          <set attributeName="r" to="0" begin="synapseLogoPath.begin" />
          <animate
            attributeName="r"
            values="0; 6"
            begin="synapseLogoPath.begin + .8s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="3"
            begin="synapseLogoPath.begin + .8s"
            dur=".1s"
          />
        </circle>
      </g>
    </svg>
  )
}
