export const BridgeFlow = () => {
  return (
    <svg
      width="100%"
      viewBox="-240 0 480 164"
      xmlns="http://www.w3.org/2000/svg"
    >
      <g fill="currentcolor" fillOpacity=".05">
        <rect x="-50%" rx="4" y="0" width="100%" height="48" />
        <rect x="-50%" rx="4" y="56" width="100%" height="48" />
        <rect x="-50%" rx="4" y="112" width="100%" height="48" />

        <rect x="-50%" rx="4" y="0%" width="33.3%" height="100%" />
        <rect x="16.7%" rx="4" y="0%" width="33.3%" height="100%" />
      </g>
      <set
        id="bridgeFlowTimer"
        attributeName="visibility"
        begin="0s; bridgeFlowTimerOut.end + 2s"
      />
      <line
        x1="-50%"
        y1="100%"
        x2="-50%"
        y2="100%"
        stroke="currentcolor"
        strokeWidth="4"
        strokeOpacity=".25"
      >
        <set attributeName="x1" to="-50%" begin="bridgeFlowTimer.begin" />
        <set attributeName="x2" to="-50%" begin="bridgeFlowTimer.begin" />
        <animate
          attributeName="x2"
          values="-50%; 50%"
          begin="bridgeFlowSend.begin"
          dur="4s"
          calcMode="linear"
          keyTimes="0; 1"
          keySplines=".5 0 1 1"
          fill="freeze"
        />
        <animate
          id="bridgeFlowTimerOut"
          attributeName="x1"
          values="-50%; 50%"
          begin="bridgeFlowReceive.end + 1s"
          dur=".75s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </line>
      {/* <line
        x1="-16.7%"
        y1="47"
        x2="16.7%"
        y2="47"
        stroke="var(--synapse-green-secondary)"
        strokeWidth="2"
      >
        <animate
          attributeName="x2"
          values="-16.7%; 16.7%"
          begin="0s; bridgeFlowSign.begin"
          dur="8s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </line>
      <line
        x1="-16.7%"
        y1="103"
        x2="16.7%"
        y2="103"
        stroke="var(--synapse-green-secondary)"
        strokeWidth="2"
      >
        <set attributeName="x2" to="-16.7%" begin="bridgeFlowTimer.begin" />
        <animate
          attributeName="x2"
          values="-16.7%; 0%"
          begin="bridgeFlowSend.begin"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          attributeName="x2"
          values="0%; 16.7%"
          begin="bridgeFlowReceive.begin"
          dur="1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </line>
      <line
        x1="-16.7%"
        y1="159"
        x2="16.7%"
        y2="159"
        stroke="var(--synapse-green-secondary)"
        strokeWidth="2"
      >
        <set attributeName="x2" to="-16.7%" begin="bridgeFlowTimer.begin" />
        <animate
          attributeName="x2"
          values="-16.7%; 16.7%"
          begin="bridgeFlowMint.begin"
          dur="2.5s"
          calcMode="linear"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </line> */}
      <g fill="currentcolor" textAnchor="middle" dominantBaseline="middle">
        <text x="-33%" y="24">
          originChain
        </text>
        <text x="33%" y="24">
          destChain
        </text>
        <text y="24">App / SDK</text>
        <text y="80">Wallet</text>
        <text y="136">Bridge</text>
      </g>
      <circle r="12" cx="-21%" cy="136" fill="hsl(285deg 100% 50%)">
        <set attributeName="opacity" to="1" begin="bridgeFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="1; 0"
          dur=".1s"
          begin="bridgeFlowBurn.begin"
          repeatCount="3"
          fill="freeze"
        />
      </circle>
      <circle r="12" fill="hsl(211deg 67% 50%)">
        <set attributeName="cy" to="80" begin="bridgeFlowTimer.begin" />
        <set attributeName="cx" to="-33%" begin="bridgeFlowTimer.begin" />
        <animate
          id="bridgeFlowSign"
          attributeName="opacity"
          values=".5; 1"
          dur=".1s"
          repeatCount="3"
          begin="bridgeFlowTimer.begin + 1s"
          fill="freeze"
        />
        <animate
          id="bridgeFlowSend"
          attributeName="cy"
          to="136"
          dur=".5s"
          begin="bridgeFlowSign.end + 2s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          attributeName="cx"
          to="-24.5%"
          dur=".5s"
          begin="bridgeFlowSend.begin + .5s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          id="bridgeFlowBurn"
          attributeName="cx"
          to="-21%"
          dur=".5s"
          begin="bridgeFlowSend.begin + 2s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </circle>
      <circle r="12" cx="21%" cy="136" fill="hsl(285deg 100% 50%)">
        <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; 1"
          dur=".1s"
          begin="bridgeFlowBurn.begin"
          repeatCount="3"
          fill="freeze"
        />
      </circle>
      <circle r="12" cx="30%" cy="136" fill="hsl(164deg 37% 50%)">
        <set attributeName="cy" to="136" begin="bridgeFlowTimer.begin" />
        <set attributeName="cx" to="21%" begin="bridgeFlowTimer.begin" />
        <animate
          id="bridgeFlowMint"
          attributeName="cx"
          to="24%"
          dur=".5s"
          begin="bridgeFlowBurn.begin + .1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          id="bridgeFlowReceive"
          attributeName="cx"
          to="33%"
          dur=".5s"
          begin="bridgeFlowMint.end + 1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          attributeName="cy"
          to="80"
          dur=".5s"
          begin="bridgeFlowReceive.begin + .3s"
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
        stroke="hsl(164deg 37% 50%)"
        fill="none"
        opacity="0"
        strokeDasharray="2 3"
      >
        <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; .3; 0; .5; 0; .7; 0; 1"
          dur=".4s"
          begin="bridgeFlowSign.begin"
          fill="freeze"
        />
        <animate
          attributeName="stroke-dashoffset"
          by="5"
          dur="1s"
          repeatCount="indefinite"
        />
      </circle>
      {/* <g fill="var(--synapse-green-secondary)">
        <circle r="3" cx="-14%" cy="24">
          <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
          <set
            attributeName="opacity"
            to="1"
            begin="bridgeFlowTimer.begin + .25s"
          />
        </circle>
        <circle r="3" cx="-14%" cy="80">
          <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
          <set attributeName="opacity" to="1" begin="bridgeFlowSend.end" />
        </circle>
        <circle r="3" cx="-14%" cy="136">
          <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
          <set
            attributeName="opacity"
            to="1"
            begin="bridgeFlowSend.end + .5s"
          />
        </circle>
        <circle r="3" cx="14%" cy="136">
          <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
          <set
            attributeName="opacity"
            to="1"
            begin="bridgeFlowBurn.end + .2s"
          />
        </circle>
        <circle r="3" cx="14%" cy="80">
          <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
          <set attributeName="opacity" to="1" begin="bridgeFlowReceive.end" />
        </circle>
        <circle r="3" cx="14%" cy="24">
          <set attributeName="opacity" to="0" begin="bridgeFlowTimer.begin" />
          <set
            attributeName="opacity"
            to="1"
            begin="bridgeFlowReceive.end + .5s"
          />
        </circle>
      </g> */}
    </svg>
  )
}
