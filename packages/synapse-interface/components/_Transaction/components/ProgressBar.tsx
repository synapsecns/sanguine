import React, { useEffect, useState } from 'react'

type ProgressBarProps = {
  elapsedTime: number
  totalTime: number
  isComplete: boolean
}

export const ProgressBar = ({
  elapsedTime,
  totalTime,
  isComplete,
}: ProgressBarProps) => {
  const [progress, setProgress] = useState(0)

  const percentageReachedOfTotal = (elapsedTime / totalTime) * 100

  useEffect(() => {
    if (isComplete) {
      /** Set Progress to 200% to speed up animation */
      setProgress(200)
    } else {
      setProgress(percentageReachedOfTotal)
    }
  }, [percentageReachedOfTotal, isComplete])

  return (
    <div id="progress-bar" className="w-full h-2 overflow-hidden bg-gray-200">
      <div
        style={{ width: `${progress}%`, transition: 'width 1s ease-in-out' }}
        className="h-full bg-green-300"
      ></div>
    </div>
  )
}
