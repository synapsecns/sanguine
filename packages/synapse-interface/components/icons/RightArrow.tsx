export const RightArrow = ({
  color,
  className,
}: {
  color: string
  className?: string
}) => {
  return (
    <svg
      className={className}
      width="7"
      height="9"
      viewBox="0 0 7 12"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M1 1L6 6L1 11"
        stroke={color}
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  )
}
