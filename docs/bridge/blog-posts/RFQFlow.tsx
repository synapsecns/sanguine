export const RFQFlow = () => {
  return (
    <svg
      width="100%"
      viewBox="-240 0 480 224"
      xmlns="http://www.w3.org/2000/svg"
      className="flowAnimation"
    >
      <set
        id="rfqFlowTimer"
        attributeName="x"
        begin="0s; rfqFlowTimerOut.end + 2s"
      />
      <g fill="currentcolor" fillOpacity=".05">
        <rect x="-50%" rx="4" y="0" width="100%" height="48" />
        <rect x="-50%" rx="4" y="56" width="100%" height="48" />
        <rect x="-50%" rx="4" y="112" width="100%" height="48" />
        <rect x="-50%" rx="4" y="168" width="100%" height="48" />
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
        <set attributeName="x1" to="-50%" begin="rfqFlowTimer.begin" />
        <set attributeName="x2" to="-50%" begin="rfqFlowTimer.begin" />
        <animate
          attributeName="x2"
          values="-50%; 50%"
          begin="rfqFlowSend.begin"
          dur="1.85s"
          calcMode="linear"
          keyTimes="0; 1"
          keySplines=".5 0 1 1"
          fill="freeze"
        />
        <animate
          id="rfqFlowTimerOut"
          attributeName="x1"
          values="-50%; 50%"
          begin="rfqFlowRepay.end + 1s"
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
        <text y="80">User</text>
        <text y="136">Relayer</text>
        <text y="192">Bridge</text>
      </g>
      <circle
        cx="-33%"
        cy="80"
        r="12"
        fill="hsl(211deg 67% 50%)"
        stroke="hsl(211deg 67% 50%)"
      >
        <set attributeName="cy" to="80" begin="rfqFlowTimer.begin" />
        <animate
          id="rfqFlowSign"
          attributeName="opacity"
          values="0; 1"
          dur=".1s"
          repeatCount="3"
          begin="rfqFlowTimer.begin + 1s"
          fill="freeze"
        />
        <animate
          id="rfqFlowSend"
          attributeName="cy"
          to="192"
          dur=".5s"
          begin="rfqFlowSign.end + 2s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
        <animate
          id="rfqFlowRepay"
          attributeName="cy"
          to="136"
          dur=".5s"
          begin="rfqFlowReceive.end + 1.1s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </circle>
      <circle
        cx="-33%"
        cy="136"
        r="12"
        stroke="hsl(211deg 67% 50%)"
        fill="none"
        opacity="0"
        strokeDasharray="2.5"
      >
        <set attributeName="opacity" to="0" begin="rfqFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; .3; 0; .5; 0; .7; 0; 1"
          dur=".4s"
          begin="rfqFlowSign.begin"
          fill="freeze"
        />
        <animate
          attributeName="stroke-dashoffset"
          by="5"
          dur="1s"
          repeatCount="indefinite"
        />
      </circle>
      <circle
        r="11"
        cx="33%"
        cy="80"
        stroke="hsl(164deg 37% 50%)"
        fill="none"
        opacity="0"
        strokeDasharray="2.5"
      >
        <set attributeName="opacity" to="0" begin="rfqFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; .3; 0; .5; 0; .7; 0; 1"
          dur=".4s"
          begin="rfqFlowSign.begin"
          fill="freeze"
        />
        <animate
          attributeName="stroke-dashoffset"
          by="5"
          dur="1s"
          repeatCount="indefinite"
        />
      </circle>
      <circle
        r="12"
        cx="33%"
        cy="136"
        fill="hsl(164deg 37% 50%)"
        stroke="hsl(164deg 37% 50%)"
      >
        <set attributeName="cy" to="136" begin="rfqFlowTimer.begin" />
        <animate
          attributeName="opacity"
          values="0; 1"
          dur=".1s"
          repeatCount="3"
          begin="rfqFlowSign.begin"
          fill="freeze"
        />
        <animate
          id="rfqFlowReceive"
          attributeName="cy"
          to="80"
          dur=".5s"
          begin="rfqFlowSend.begin + 2s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          fill="freeze"
        />
      </circle>
    </svg>
  )
}
