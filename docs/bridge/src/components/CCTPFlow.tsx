export const CCTPFlow = () => {
  return (
    <svg
      width="100%"
      viewBox="-240 0 480 164"
      xmlns="http://www.w3.org/2000/svg"
      className="flowAnimation"
    >
      <set
        id="cctpFlowTimer"
        attributeName="x"
        begin="0s; cctpFlowTimerOut.end + 2s"
      />
      <g fill="currentcolor" fillOpacity=".05">
        <rect x="-50%" rx="4" y="0" width="100%" height="48" />
        <rect x="-50%" rx="4" y="56" width="100%" height="48" />
        <rect x="-50%" rx="4" y="112" width="100%" height="48" />
        <rect x="-50%" rx="4" y="0%" width="33.3%" height="100%" />
        <rect x="16.7%" rx="4" y="0%" width="33.3%" height="100%" />
      </g>
      <line
        x1="-50%"
        y1="100%"
        x2="-50%"
        y2="100%"
        stroke="currentcolor"
        strokeWidth="4"
        strokeOpacity=".25"
      >
        <set attributeName="x1" to="-50%" begin="cctpFlowTimer.begin" />
        <set attributeName="x2" to="-50%" begin="cctpFlowTimer.begin" />
        <animate
          attributeName="x2"
          values="-50%; 50%"
          begin="cctpFlowSend.begin"
          dur="4s"
          calcMode="linear"
          keyTimes="0; 1"
          keySplines=".5 0 1 1"
          fill="freeze"
        />
        <animate
          id="cctpFlowTimerOut"
          attributeName="x1"
          values="-50%; 50%"
          begin="cctpFlowReceive.end + 1s"
          dur=".75s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </line>
      <g fill="currentcolor" textAnchor="middle" dominantBaseline="middle">
        <text x="-33%" y="24">
          originChain
        </text>
        <text x="33%" y="24">
          destChain
        </text>
        <text y="24">App / SDK</text>
        <text y="80">Wallet</text>
        <text y="136">Circle</text>
      </g>
      <circle
        cx="-33%"
        cy="80"
        r="12"
        stroke="hsl(211deg 67% 50%)"
        fill="none"
        opacity=".5"
        strokeDasharray="2.5"
      ></circle>
      <circle
        cx="-33%"
        cy="80"
        r="12"
        fill="hsl(211deg 67% 50%)"
        stroke="hsl(211deg 67% 50%)"
      >
        <set attributeName="opacity" to="1" begin="cctpFlowTimer.begin" />
        <set attributeName="cy" to="80" begin="cctpFlowTimer.begin" />
        <set attributeName="cx" to="-33%" begin="cctpFlowTimer.begin" />
        <animate
          id="cctpFlowSign"
          attributeName="opacity"
          values=".5; 1"
          dur=".1s"
          repeatCount="3"
          begin="cctpFlowTimer.begin + 1s"
          fill="freeze"
        />
        <animate
          id="cctpFlowSend"
          attributeName="cy"
          to="136"
          dur=".5s"
          begin="cctpFlowSign.end + 2s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          id="cctpFlowBurn"
          attributeName="cx"
          to="-31.5%"
          dur=".5s"
          begin="cctpFlowSend.begin + 2s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          attributeName="opacity"
          values="1; 0"
          dur=".1s"
          begin="cctpFlowBurn.begin"
          repeatCount="3"
          fill="freeze"
        />
      </circle>
      <circle
        r="12"
        cx="31.5%"
        cy="136"
        fill="hsl(211deg 67% 50%)"
        stroke="hsl(211deg 67% 50%)"
      >
        <set attributeName="cy" to="136" begin="cctpFlowTimer.begin" />
        <set attributeName="cx" to="31.5%" begin="cctpFlowTimer.begin" />
        <animate
          id="cctpFlowMint"
          attributeName="cx"
          to="33%"
          dur=".5s"
          begin="cctpFlowBurn.begin + .1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <set attributeName="opacity" to="0" begin="cctpFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; 1"
          begin="cctpFlowMint.begin"
          dur=".1s"
          repeatCount="5"
          fill="freeze"
        />
        <animate
          id="cctpFlowReceive"
          attributeName="cy"
          to="80"
          dur=".5s"
          begin="cctpFlowMint.end + 1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </circle>
      <circle
        r="11"
        cx="33%"
        cy="80"
        stroke="hsl(211deg 67% 50%)"
        fill="none"
        opacity="0"
        strokeDasharray="2.5"
      >
        <set attributeName="opacity" to="0" begin="cctpFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; .3; 0; .5; 0; .7; 0; 1"
          dur=".4s"
          begin="cctpFlowSign.begin"
          fill="freeze"
        />
        <animate
          attributeName="stroke-dashoffset"
          by="5"
          dur="1s"
          repeatCount="indefinite"
        />
      </circle>
    </svg>
  )
}
