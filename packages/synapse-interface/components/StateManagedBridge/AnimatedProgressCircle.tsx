import { BridgeQuote } from '@/utils/types'
import { useState, useEffect, useMemo } from 'react'

export const BridgeQuoteResetTimer = ({
  bridgeQuote,
  hasValidQuote,
  duration, // in ms
}: {
  bridgeQuote: BridgeQuote
  hasValidQuote: boolean
  duration: number
}) => {
  const memoizedTimer = useMemo(() => {
    if (hasValidQuote) {
      return (
        <AnimatedProgressCircle
          animateKey={bridgeQuote.id}
          duration={duration}
        />
      )
    }
    return null
  }, [bridgeQuote, hasValidQuote, duration])

  return memoizedTimer
}

const AnimatedProgressCircle = ({
  animateKey,
  duration,
}: {
  animateKey: string
  duration: number
}) => {
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
      className="absolute -rotate-90 -scale-y-100"
    >
      <circle r="8" />
      <circle r="8" strokeDasharray="1" pathLength="1">
        <animate
          attributeName="stroke-dashoffset"
          values="1; 2"
          dur={`${convertMsToSeconds(duration)}s`}
        />
      </circle>
    </svg>
  )
}

const convertMsToSeconds = (ms: number) => {
  return Math.ceil(ms / 1000)
}
