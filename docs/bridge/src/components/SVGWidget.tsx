export default () => {
  return (
    <svg
      width="400"
      height="382"
      viewBox="0 0 400 382"
      fill="currentcolor"
      style={{ pointerEvents: 'none' }}
    >
      <defs>
        <linearGradient id="synGrad">
          <stop offset="0%" stopColor="hsl(300deg 60% 50%)" />
          <stop offset="100%" stopColor="hsl(270deg 85% 75%)" />
        </linearGradient>
        <mask id="switcherMask">
          <rect width="100%" height="100%" fill="white" />
          <rect x="182" y="122" width="36" height="36" rx="4" fill="black" />
        </mask>
      </defs>
      <style>
        {`
        .bridgeSurface {
            fill: white;
            fill-opacity: 1;
            stroke: currentcolor;
            stroke-opacity: .1;
        }
        :root[data-theme='dark'] .bridgeSurface {
            fill: currentcolor;
            fill-opacity: .05;
            stroke: none;
        }
        `}
      </style>
      <rect width="100%" height="100%" rx="8" fillOpacity=".05" />
      <rect
        x="8"
        y="8"
        width="384"
        height="128"
        rx="4"
        className="bridgeSurface"
        mask="url(#switcherMask)"
      />
      <rect x="16" y="20" width="116" height="36" rx="4" fillOpacity=".1" />
      <text x="28" y="44" style={{ fontSize: '1.1rem', fontWeight: 500 }}>
        Network
      </text>
      <path stroke="currentcolor" d="m108 34 4.5 8 4.5 -8z" />
      <rect x="304" y="62" width="80" height="36" rx="4" fillOpacity=".1" />
      <text
        x="316"
        y="82"
        dominantBaseline="middle"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        Out
      </text>
      <path stroke="currentcolor" d="m361 77 4.5 8 4.5 -8z" />
      <text
        x="380"
        y="124"
        textAnchor="end"
        opacity=".67"
        style={{ fontSize: '.9rem' }}
      >
        Available 0.0000
      </text>
      <text
        x="24"
        y="102"
        dominantBaseline="middle"
        style={{ fontSize: '2rem', fontWeight: 500 }}
      >
        0.0000
      </text>
      <rect
        x="8"
        y="144"
        width="384"
        height="128"
        rx="4"
        className="bridgeSurface"
        mask="url(#switcherMask)"
      />
      <rect x="16" y="154" width="116" height="36" rx="4" fillOpacity=".1" />
      <text x="28" y="178" style={{ fontSize: '1.1rem', fontWeight: 500 }}>
        Network
      </text>
      <path stroke="currentcolor" d="m108 168 4.5 8 4.5 -8z" />
      <text
        x="24"
        y="230"
        dominantBaseline="middle"
        style={{ fontSize: '2rem', fontWeight: 500 }}
      >
        0.0000
      </text>
      <rect x="320" y="210" width="64" height="36" rx="4" fillOpacity=".1" />
      <text
        x="332"
        y="230"
        dominantBaseline="middle"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        In
      </text>
      <path stroke="currentcolor" d="m361 225 4.5 8 4.5 -8z" />
      <rect
        x="186"
        y="126"
        width="28"
        height="28"
        rx="2"
        className="bridgeSurface"
      />
      <path
        stroke="currentcolor"
        strokeWidth="2"
        opacity=".7"
        fill="none"
        d="m200 131.5 v15 l-6 -5.5 m6 5.5 l6 -5.5"
      />
      <text
        x="364"
        y="300"
        textAnchor="end"
        style={{ fontSize: '.9rem', fontWeight: 500 }}
      >
        25 seconds via SynapseRFQ
      </text>
      <path
        stroke="currentcolor"
        strokeWidth="2"
        fill="none"
        opacity=".67"
        d="m372 289 5 5 5 -5 m-10 7 5 5 5 -5"
      />
      <rect
        x="8"
        y="320"
        width="384"
        height="52"
        rx="6"
        fill="white"
        fillOpacity=".025"
        stroke="url(#synGrad)"
      />
      <text
        x="200"
        y="348"
        textAnchor="middle"
        dominantBaseline="middle"
        style={{ fontSize: '1.5rem', fontWeight: 500 }}
      >
        Bridge
      </text>
    </svg>
  )
}
