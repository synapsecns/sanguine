import { useState, useEffect, useMemo } from 'react'

import { BridgeQuote } from '@/utils/types'
import { convertMsToSeconds } from '@/utils/time'

export const BridgeQuoteResetTimer = ({
  bridgeQuote,
  isLoading,
  isActive,
  duration, // in ms
}: {
  bridgeQuote: BridgeQuote
  isLoading: boolean
  isActive: boolean
  duration: number
}) => {
  const memoizedTimer = useMemo(() => {
    if (!isActive) return null

    if (isLoading) {
      return <AnimatedLoadingCircle />
    } else {
      return (
        <AnimatedProgressCircle
          animateKey={bridgeQuote.id}
          duration={duration}
        />
      )
    }
  }, [bridgeQuote, duration, isActive])

  return memoizedTimer
}

const AnimatedLoadingCircle = () => {
  return (
    <svg
      width="24"
      height="24"
      viewBox="-12 -12 24 24"
      stroke="currentcolor"
      fill="none"
      className="absolute block -rotate-90"
    >
      <circle r="8" pathLength="1" stroke-dashArray="0.05" stroke-opacity=".5">
        <animate
          attributeName="stroke-dashoffset"
          to="-1"
          dur="2.5s"
          repeatCount="indefinite"
        />
      </circle>
    </svg>
  )
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
      fill="none"
      className="absolute block -rotate-90"
    >
      <circle r="8" pathLength="1" stroke-opacity=".25">
        <animate
          attributeName="stroke-dashoffset"
          to="-1"
          dur="2.5s"
          repeatCount="indefinite"
        />
        <set
          attributeName="stroke-dasharray"
          to="0.05"
          begin={`${convertMsToSeconds(duration)}s`}
        />
        <set
          attributeName="stroke-opacity"
          to="0.5"
          begin={`${convertMsToSeconds(duration)}s`}
        />
      </circle>
      <circle r="8" stroke-dasharray="1" pathLength="1">
        <animate
          attributeName="stroke-dashoffset"
          values="2; 1"
          dur={`${convertMsToSeconds(duration)}s`}
          fill="freeze"
        />
        <animate
          attributeName="stroke-opacity"
          values="0; 1"
          dur={`${convertMsToSeconds(duration)}s`}
        />
      </circle>
    </svg>
  )
}
