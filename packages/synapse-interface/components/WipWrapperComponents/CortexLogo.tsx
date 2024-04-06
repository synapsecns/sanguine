export function CortexAnchor({ id, width = 20, height = 20 }) {
  return (
    <a href="#" className="text-base flex gap-0.5 items-center">
      <CortexIcon id={id} width={width} height={height} />
      <span className="hidden xs:block -mt-px">
        Cor<span className="opacity-50 -ml-px">/</span>tex
      </span>
    </a>
  )
}

export default function CortexIcon({ id, width, height }) {
  return (
    <svg
      width={width}
      height={height}
      viewBox="-10 -10 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <animate id={id} begin="1s; mouseenter; mouseleave" />
      <circle
        r="9"
        stroke="white"
        strokeWidth="2"
        opacity=".5"
        strokeDasharray="1 1"
        pathLength="2"
        strokeDashoffset="-.5"
        strokeLinecap="square"
      >
        <animate
          id={`${id}Rotate`}
          attributeName="stroke-dashoffset"
          by="2"
          dur="1s"
          begin={`${id}.begin`}
          calcMode="spline"
          keyTimes="0; 1"
          keySplines=".5 0 .2 1"
          restart="whenNotActive"
        />
      </circle>
      <circle r="6" fill="white">
        <animate
          attributeName="opacity"
          values=".5; 1"
          repeatCount="3"
          dur=".1s"
          begin={`${id}Rotate.begin`}
        />
        <animate
          attributeName="r"
          values="3; 6"
          dur=".5s"
          calcMode="spline"
          keyTimes="0; 1"
          keySplines="0 0 .2 1"
          begin={`${id}Rotate.begin`}
        />
      </circle>
    </svg>
  )
}
