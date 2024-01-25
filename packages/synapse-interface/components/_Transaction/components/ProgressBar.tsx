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
  const currentProgress = (elapsedTime / totalTime) * 100
  const [progress, setProgress] = useState(currentProgress)

  console.log('elapsedTime:', elapsedTime)
  console.log('totalTime:', totalTime)
  console.log('progress: ', progress)

  useEffect(() => {
    if (isComplete) {
      setProgress(200)
    } else {
      setProgress(currentProgress)
    }
  }, [currentProgress, isComplete])

  return (
    <div id="progress-bar" className="w-full h-1 overflow-hidden bg-white">
      <div
        style={{ width: `${progress}%`, transition: 'width 1s linear' }}
        className="h-full bg-green-500"
      ></div>
    </div>
  )
}
