import { BridgeQuote } from '@/utils/types'
import { useState, useEffect } from 'react'

export const BridgeQuoteResetTimer = ({
  bridgeQuote,
  hasValidQuote,
}: {
  bridgeQuote: BridgeQuote
  hasValidQuote: boolean
}) => {
  if (hasValidQuote) {
    return <AnimatedProgressCircle animateKey={bridgeQuote.id} />
  }
}

const AnimatedProgressCircle = ({ animateKey }) => {
  const [animationKey, setAnimationKey] = useState(0)

  useEffect(() => {
    setAnimationKey((prevKey) => prevKey + 1)
  }, [animateKey])

  return (
    <svg
      key={animationKey}
      width="24"
      height="24"
      viewBox="-12 -12 24 24"
      stroke="currentcolor"
      strokeOpacity=".33"
      fill="none"
      className="absolute -rotate-90 -scale-y-100 right-4"
    >
      <circle r="8" />
      <circle r="8" strokeDasharray="1" pathLength="1">
        <animate attributeName="stroke-dashoffset" values="1; 2" dur="15s" />
      </circle>
    </svg>
  )
}
