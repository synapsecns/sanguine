import React, { useEffect, useState } from 'react'

type ProgressBarProps = {
  elapsedTime: number
  estimatedTotalTime: number
}

export const ProgressBar = ({
  elapsedTime,
  estimatedTotalTime,
}: ProgressBarProps) => {
  const [progress, setProgress] = useState(0)

  useEffect(() => {
    const interval = setInterval(() => {
      setProgress((elapsedTime / estimatedTotalTime) * 100)
    }, 1000)

    return () => clearInterval(interval)
  }, [elapsedTime, estimatedTotalTime])

  return (
    <div id="progress-bar" className="w-full h-5 bg-gray-200">
      <div
        style={{ width: `${progress}%`, transition: 'width 1s ease-in-out' }}
        className="h-full bg-green-300"
      ></div>
    </div>
  )
}
