import React from 'react'

interface ArrowProps {
  className?: string
  color?: string
}

const Arrow: React.FC<ArrowProps> = ({ className, color = 'currentColor' }) => (
  <svg
    width="40"
    height="42"
    viewBox="0 0 40 42"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    className={className}
  >
    <path
      d="M16 12L24 21L16 30"
      stroke={color}
      strokeOpacity="0.4"
      strokeWidth="2"
    />
  </svg>
)

export default Arrow
