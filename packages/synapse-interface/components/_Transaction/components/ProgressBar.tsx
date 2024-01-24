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

  useEffect(() => {
    if (isComplete) {
      setProgress(200)
    } else {
      const interval = setInterval(() => {
        setProgress((prevProgress) => prevProgress + 1)
      }, 1000)

      return () => clearInterval(interval)
      // setProgress(currentProgress)
    }
  }, [currentProgress, isComplete])

  console.log('progress: ', progress)

  return (
    <div id="progress-bar" className="w-full h-1 overflow-hidden bg-gray-200">
      <div
        style={{ width: `${progress}%`, transition: 'width 1s linear' }}
        className="h-full bg-green-300"
      ></div>
    </div>
  )
}
