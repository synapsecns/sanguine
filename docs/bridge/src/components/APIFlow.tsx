export const APIFlow = () => {
  return (
    <svg
      width="100%"
      viewBox="-240 0 480 120"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className="flowAnimation"
    >
      <defs>
        <path id="synFigDiamond" d="M4 0 8 4 4 8 0 4z" />
        <path id="synFigArrow" d="M8 4 0 0 0 8z" />
        <marker
          id="synFigDiamondPurple"
          viewBox="0 0 8 8"
          refX="4"
          refY="4"
          markerWidth="7"
          markerHeight="7"
          fill="var(--syn-fig-purple)"
        >
          <use href="#synFigDiamond" />
        </marker>
        <marker
          id="synFigDiamondMagenta"
          viewBox="0 0 8 8"
          refX="4"
          refY="4"
          markerWidth="7"
          markerHeight="7"
          fill="var(--syn-fig-magenta)"
        >
          <use href="#synFigDiamond" />
        </marker>
        <marker
          id="synFigArrowMagenta"
          viewBox="0 0 8 8"
          refX="4"
          refY="4"
          markerWidth="7"
          markerHeight="7"
          orient="auto"
          fill="var(--syn-fig-magenta)"
        >
          <use href="#synFigArrow" />
        </marker>
        <marker
          id="synFigArrowPurple"
          viewBox="0 0 8 8"
          refX="4"
          refY="4"
          markerWidth="7"
          markerHeight="7"
          orient="auto"
          fill="var(--syn-fig-purple)"
        >
          <use href="#synFigArrow" />
        </marker>
        <marker
          id="synFigArrowGreen"
          viewBox="0 0 8 8"
          refX="4"
          refY="4"
          markerWidth="7"
          markerHeight="7"
          orient="auto-start-reverse"
          fill="var(--syn-fig-green)"
        >
          <use href="#synFigArrow" />
        </marker>
        <marker
          id="synFigArrowBlue"
          viewBox="0 0 8 8"
          refX="4"
          refY="4"
          markerWidth="7"
          markerHeight="7"
          orient="auto-start-reverse"
          fill="var(--syn-fig-blue)"
        >
          <use href="#synFigArrow" />
        </marker>
      </defs>
      <g fill="currentcolor" fillOpacity=".05">
        <rect x="-50%" rx="4" y="0" width="100%" height="56" />
        <rect x="-50%" rx="4" y="64" width="100%" height="56" />
        <rect x="112" rx="4" width="128" height="100%" />

        <rect x="-128" rx="4" y="12" width="72" height="32" />
        <rect x="-32" rx="4" y="12" width="128" height="32" />
        <rect x="132" rx="4" y="12" width="88" height="32" />
      </g>
      <path
        stroke="var(--syn-fig-green)"
        fill="none"
        d="M-180 84 H-92 V44"
        markerStart="url(#synFigArrowGreen)"
        markerEnd="url(#synFigArrowGreen)"
      />
      <path
        stroke="var(--syn-fig-blue)"
        fill="none"
        d="M-180 92 H32 V44"
        markerStart="url(#synFigArrowBlue)"
        markerEnd="url(#synFigArrowBlue)"
      />
      <path
        stroke="var(--syn-fig-purple)"
        fill="none"
        d="M-180 100 H104 V28 H132"
        markerStart="url(#synFigDiamondPurple)"
        markerEnd="url(#synFigArrowPurple)"
      />
      <path
        stroke="var(--syn-fig-magenta)"
        fill="none"
        d="M176 44 V77"
        markerStart="url(#synFigDiamondMagenta)"
        markerEnd="url(#synFigArrowMagenta)"
      />
      <g
        fill="currentcolor"
        dominantBaseline="middle"
        style={{ fontSize: '.9em' }}
      >
        <text x="-216" y="29">
          REST API
        </text>
        <g textAnchor="middle" fill="var(--syn-method)">
          <text x="-92" y="29">
            /bridge
          </text>
          <text x="32" y="29">
            /bridgeTxInfo
          </text>
          <text x="176" y="29">
            Contract
          </text>
        </g>
        <text x="-216" y="93">
          App
        </text>
        <text x="176" y="93" textAnchor="middle">
          Chain
        </text>
      </g>

      <g strokeWidth=".5" fill="var(--ifm-background-surface-color)">
        <rect
          x="-120"
          y="51"
          width="56"
          height="17"
          rx="2"
          stroke="var(--syn-fig-green)"
        />
        <rect
          x="8"
          y="51"
          width="48"
          height="17"
          rx="2"
          stroke="var(--syn-fig-blue)"
        />
        <rect
          x="82"
          y="51"
          width="44"
          height="17"
          rx="2"
          stroke="var(--syn-fig-purple)"
        />
        <rect
          x="148"
          y="51"
          width="56"
          height="17"
          rx="2"
          stroke="var(--syn-fig-magenta)"
        />
      </g>
      <g
        className="temp"
        textAnchor="middle"
        dominantBaseline="middle"
        style={{ fontSize: '.75em' }}
      >
        <text x="-92" y="60" fill="var(--syn-fig-green)">
          Quote
        </text>
        <text x="32" y="60" fill="var(--syn-fig-blue)">
          Data
        </text>
        <text x="104" y="60" fill="var(--syn-fig-purple)">
          Sign
        </text>
        <text x="176" y="60" fill="var(--syn-fig-magenta)">
          Submit
        </text>
      </g>
    </svg>
  )
}
