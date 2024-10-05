export default () => {
  return (
    <svg
      width="400"
      height="406"
      viewBox="0 0 400 406"
      fill="currentcolor"
      //   style={{ pointerEvents: 'none' }}
    >
      <defs>
        <linearGradient id="synGrad">
          <stop offset="0%" stopColor="hsl(300deg 60% 50%)" />
          <stop offset="100%" stopColor="hsl(270deg 85% 75%)" />
        </linearGradient>
        <mask id="switcherMask">
          <rect width="100%" height="100%" fill="white" />
          <rect x="182" y="134" width="36" height="36" rx="4" fill="black" />
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
        height="140"
        rx="4"
        className="bridgeSurface"
        mask="url(#switcherMask)"
      />
      <circle
        cx="266"
        cy="37"
        r="4"
        stroke="hsl(285deg 72.5% 62.5%)"
        fill="none"
      />
      <text x="380" y="42" textAnchor="end" style={{ fontSize: '.9rem' }}>
        Connect Wallet
      </text>
      <text x="20" y="32" opacity=".67" style={{ fontSize: '.9rem' }}>
        From
      </text>
      <text x="20" y="56" style={{ fontWeight: 500 }}>
        Network
      </text>
      <path stroke="currentcolor" opacity=".5" d="m96 33 4.5 8 4.5 -8z" />
      <rect
        x="16"
        y="72"
        width="368"
        height="56"
        rx="6"
        stroke="#80808040"
        fill="none"
      />
      <rect x="24" y="80" width="76" height="40" rx="4" fillOpacity=".1" />
      <text
        x="32"
        y="101"
        dominantBaseline="middle"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        Out
      </text>
      <path stroke="currentcolor" opacity=".5" d="m77 96 4.5 8 4.5 -8z" />
      <text
        x="116"
        y="102"
        dominantBaseline="middle"
        style={{ fontSize: '1.5rem', fontWeight: 600 }}
      >
        0.0000
      </text>

      <rect
        x="8"
        y="156"
        width="384"
        height="140"
        rx="4"
        className="bridgeSurface"
        mask="url(#switcherMask)"
      />
      <rect
        x="16"
        y="220"
        width="368"
        height="56"
        rx="6"
        stroke="#80808040"
        fill="none"
      />
      <text x="20" y="180" opacity=".67" style={{ fontSize: '.9rem' }}>
        To
      </text>
      <text x="20" y="204" style={{ fontWeight: 500 }}>
        Network
      </text>
      <path stroke="currentcolor" opacity=".5" d="m96 180 4.5 8 4.5 -8z" />
      <rect x="24" y="228" width="64" height="40" rx="4" fillOpacity=".1" />
      <text
        x="32"
        y="250"
        dominantBaseline="middle"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        In
      </text>
      <path stroke="currentcolor" opacity=".5" d="m64 244 4.5 8 4.5 -8z" />
      <text
        x="104"
        y="250"
        dominantBaseline="middle"
        style={{ fontSize: '1.5rem', fontWeight: 500 }}
      >
        0.0000
      </text>
      <rect
        x="186"
        y="138"
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
        d="m200 143.5 v15 l-6 -5.5 m6 5.5 l6 -5.5"
      />

      <text x="364" y="324" textAnchor="end" style={{ fontSize: '.9rem' }}>
        Select origin token
      </text>
      <path
        stroke="currentcolor"
        strokeWidth="2"
        fill="none"
        opacity=".67"
        d="m372 313 5 5 5 -5 m-10 7 5 5 5 -5"
      />

      <rect
        x="8"
        y="344"
        width="384"
        height="52"
        rx="6"
        fill="white"
        fillOpacity=".025"
        stroke="url(#synGrad)"
      />
      <text
        x="200"
        y="372"
        textAnchor="middle"
        dominantBaseline="middle"
        // fill="white"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        Bridge
      </text>
    </svg>
  )
}
