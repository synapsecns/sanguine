import { BridgeQuote } from '@/utils/types'
import { useState, useEffect, useMemo } from 'react'

export const BridgeQuoteResetTimer = ({
  bridgeQuote,
  isActive,
  duration, // in ms
}: {
  bridgeQuote: BridgeQuote
  isActive: boolean
  duration: number
}) => {
  const memoizedTimer = useMemo(() => {
    if (isActive) {
      return (
        <AnimatedProgressCircle
          animateKey={bridgeQuote.id}
          duration={duration}
        />
      )
    }
    return null
  }, [bridgeQuote, duration, isActive])

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
      stroke-opacity=".33"
      fill="none"
      className="absolute -rotate-90"
    >
      <circle r="8" />
      <circle r="8" stroke-dasharray="1" pathLength="1">
        <animate
          attributeName="stroke-dashoffset"
          values="2; 1"
          dur={`${convertMsToSeconds(duration)}s`}
          fill="freeze"
        />
      </circle>
    </svg>
  )
}

const convertMsToSeconds = (ms: number) => {
  return Math.ceil(ms / 1000)
}
