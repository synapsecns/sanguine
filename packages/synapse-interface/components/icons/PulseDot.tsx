const PulseDot = ({ className }) => {
  return (
    <svg
      width="8"
      height="8"
      viewBox="-4 -4 8 8"
      overflow="visible"
      className={className}
      xmlns="http://www.w3.org/2000/svg"
    >
      <circle r="4">
        <animate
          attributeName="stroke-width"
          values="0; 16"
          dur="1.5s"
          repeatCount="indefinite"
        />
        <animate
          attributeName="stroke-opacity"
          values=".5; 0"
          dur="1.5s"
          repeatCount="indefinite"
        />
      </circle>
    </svg>
  )
}

export default PulseDot
