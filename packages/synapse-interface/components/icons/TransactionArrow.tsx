export default function TransactionArrow(props) {
  return (
    <svg
      width="16px"
      viewBox="0 0 16 16"
      preserveAspectRatio="none"
      style={{ alignSelf: 'stretch', overflow: 'visible' }}
      {...props}
    >
      <path
        d="M 0 0 L 16 8 L 0 16"

        strokeWidth="1"
        vectorEffect="non-scaling-stroke"
      />
    </svg>
  )
}
