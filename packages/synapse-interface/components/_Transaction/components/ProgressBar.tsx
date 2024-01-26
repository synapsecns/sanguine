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
  // const currentTime = useIntervalTimer(500)
  const currentTime = getTimeMinutesBeforeNow(0)

  const elapsedTime = currentTime - initialTime

  const elapsedProgressInFraction = elapsedTime / estimatedDurationTime
  const elapsedProgressInPercent = elapsedProgressInFraction * 100

  const remainingTime = estimatedDurationTime - elapsedTime

  const currentProgressInFraction =
    (currentTime - initialTime) / (targetTime - initialTime)

  const currentProgressInPercent = currentProgressInFraction * 100

  return (
    <>
      {!isComplete ? (
        <div
          id="progress-bar"
          className="flex w-full h-1 overflow-hidden bg-white"
        >
          <div
            style={{
              width: `${elapsedProgressInPercent}%`,
            }}
            className="h-full bg-green-500"
          ></div>
          <div
            style={{
              width: `${100 - elapsedProgressInPercent}%`,
              animationName: 'fillAnimation',
              animationDuration: `${remainingTime}s`,
              animationTimingFunction: 'ease-in',
            }}
            className="h-full bg-green-500"
          />
        </div>
      ) : (
        <div
          id="progress-bar"
          className="flex w-full h-1 overflow-hidden bg-white"
        >
          <div
            style={{
              width: `100%`,
            }}
          >
            <CompletedProgress startingPercentage={elapsedProgressInPercent} />
          </div>
        </div>
      )}
    </>
  )
}

const CompletedProgress = ({ startingPercentage }) => {
  return (
    <svg
      height="100%"
      width="100%"
      viewBox="0 0 100% 100%"
      preserveAspectRatio="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect width="100%" height="100%" fill="green">
        <animate
          attributeName="width"
          from={startingPercentage + '%'}
          to="100%"
          dur="2s"
          calcMode="spline"
          keySplines=".3 0 1 1"
        />
      </rect>
    </svg>
  )
}
