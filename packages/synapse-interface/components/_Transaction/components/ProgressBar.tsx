import React, { useEffect, useState } from 'react'
import { useIntervalTimer } from '../helpers/useIntervalTimer'
import { getTimeMinutesBeforeNow } from '@/utils/time'

type ProgressBarProps = {
  initialTime: number
  targetTime: number
  estimatedDurationTime: number
  isComplete: boolean
}

export const ProgressBar = ({
  initialTime,
  targetTime,
  estimatedDurationTime,
  isComplete,
}: ProgressBarProps) => {
  const currentTime = useIntervalTimer(1000)

  const currentProgressInFraction =
    (currentTime - initialTime) / (targetTime - initialTime)

  const currentProgressInPercent = currentProgressInFraction * 100

  // console.log('initialTime:', initialTime)
  // console.log('targetTime:', targetTime)

  return (
    <div id="progress-bar" className="w-full h-1 overflow-hidden bg-white">
      <div
        style={{
          width: `${isComplete ? 200 : currentProgressInPercent}%`,
          transition: 'width 1s ease-in-out',
        }}
        className="h-full bg-green-500"
      ></div>
    </div>
  )
}
