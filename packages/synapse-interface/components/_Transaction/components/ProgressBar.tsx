import React, { useEffect, useState } from 'react'
import { useIntervalTimer } from '../helpers/useIntervalTimer'
import { getTimeMinutesBeforeNow } from '@/utils/time'

type ProgressBarProps = {
  startTime: number
  targetTime: number
  estDuration: number
  isComplete: boolean
}

export const ProgressBar = ({
  startTime,
  targetTime,
  estDuration,
  isComplete,
}: ProgressBarProps) => {
  const currentTime = getTimeMinutesBeforeNow(0)

  const timer = useIntervalTimer(500)

  const elapsedTime = currentTime - startTime
  const remainingTime = estDuration - elapsedTime

  const percentElapsed = (elapsedTime / estDuration) * 100

  return (
    <div id="progress-bar" className="w-full h-1 overflow-hidden bg-white">
      <div
        style={{
          width: `${isComplete ? 100 : percentElapsed}%`,
          transition: `${isComplete ? 'width 1s ease-in' : 'width 1s linear'}`,
        }}
        className="h-full bg-green-500"
      ></div>
    </div>
  )
}
