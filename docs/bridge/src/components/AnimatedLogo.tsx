export default () => {
  return (
    <svg width="64" height="64" viewBox="-24 -24 48 48">
      <defs>
        <linearGradient id="synGrad" fx="0%">
          <stop offset="0%" stop-color="hsl(285deg 100% 65%)" />
          <stop offset="100%" stop-color="hsl(265deg 100% 75%)" />
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
          attributeName="stroke-dashoffset"
          to="0"
          dur=".9s"
          fill="freeze"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".67 0 .8 1"
        />
        <animate
          attributeName="stroke-width"
          to="5.5"
          dur=".9s"
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
        />
      </path>
      <g fill="url(#synGrad)">
        <circle cy="18" r="0">
          <animate
            attributeName="r"
            values="0; 6"
            begin=".05s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="3"
            begin=".05s"
            dur=".1s"
          />
        </circle>
        <circle cx="18" r="0">
          <animate
            attributeName="r"
            values="0; 6"
            begin=".3s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="4"
            begin=".3s"
            dur=".1s"
          />
        </circle>
        <circle cx="-18" r="0">
          <animate
            attributeName="r"
            values="0; 6"
            begin=".55s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="4"
            begin=".55s"
            dur=".1s"
          />
        </circle>
        <circle cy="-18" r="0">
          <animate
            attributeName="r"
            values="0; 6"
            begin=".8s"
            dur=".25s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            values="0; 1"
            repeatCount="3"
            begin=".8s"
            dur=".1s"
          />
        </circle>
      </g>
    </svg>
  )
}
