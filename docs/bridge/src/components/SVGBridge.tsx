export default () => {
  return (
    <svg
      width="400"
      height="382"
      viewBox="0 0 400 382"
      fill="currentcolor"
      //   style={{ pointerEvents: 'none' }}
    >
      <defs>
        <linearGradient id="synGrad">
          <stop offset="0%" stop-color="hsl(300deg 60% 50%)" />
          <stop offset="100%" stop-color="hsl(270deg 85% 75%)" />
        </linearGradient>
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
        y="144"
        width="384"
        height="128"
        rx="4"
        className="bridgeSurface"
      />
      <rect
        x="16"
        y="208"
        width="368"
        height="56"
        rx="6"
        stroke="#80808040"
        fill="none"
      />
      <text x="20" y="168" opacity=".67" style={{ fontSize: '.9rem' }}>
        To
      </text>
      <text x="20" y="192" style={{ fontWeight: 500 }}>
        Network
      </text>
      <path stroke="currentcolor" opacity=".5" d="m96 168 4.5 8 4.5 -8z" />
      <rect x="24" y="216" width="64" height="40" rx="4" fillOpacity=".1" />
      <text
        x="32"
        y="238"
        dominantBaseline="middle"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        In
      </text>
      <path stroke="currentcolor" opacity=".5" d="m64 232 4.5 8 4.5 -8z" />
      <text
        x="104"
        y="238"
        dominantBaseline="middle"
        style={{ fontSize: '1.5rem', fontWeight: 500 }}
      >
        0.0000
      </text>

      <text x="388" y="300" textAnchor="end" style={{ fontSize: '.9rem' }}>
        Select origin token
      </text>

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
        // fill="white"
        style={{ fontSize: '1.2rem', fontWeight: 500 }}
      >
        Bridge
      </text>
    </svg>
  )
}
