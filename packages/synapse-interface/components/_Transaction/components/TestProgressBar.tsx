import { useState, useEffect } from 'react'
import { ProgressBar } from './ProgressBar'
import { getTimeMinutesFromNow } from '@/utils/time'

export const TestProgressBar = () => {
  const [startTime, setStartTime] = useState(null)
  const [estimatedTime, setEstimatedTime] = useState(0)
  const [completedTime, setCompletedTime] = useState(0)
  const [isComplete, setIsComplete] = useState(false)

  const handleEstimatedTime = (event) => {
    setEstimatedTime(event.target.value)
  }

  const handleCompletedTime = (event) => {
    setCompletedTime(event.target.value)
  }

  const startTimer = () => {
    setStartTime(getTimeMinutesFromNow(0))
  }

  const resetTimer = () => {
    setStartTime(null)
    setIsComplete(false)
  }

  useEffect(() => {
    if (startTime && !isComplete) {
      const completedTimeout = setTimeout(() => {
        setIsComplete(true)
      }, completedTime * 1000) // Convert completedTime to milliseconds
      return () => clearTimeout(completedTimeout)
    }
  }, [startTime, isComplete, completedTime])

  return (
    <div className="bg-white border border-purple-500 max-w-[600px] mx-auto my-5">
      <div className="flex items-center justify-between">
        <div>estimated time (in seconds)</div>
        <input
          type="number"
          onChange={handleEstimatedTime}
          value={estimatedTime}
        />
      </div>
      <div className="flex items-center justify-between">
        <div>completed time (in seconds)</div>
        <input
          type="number"
          onChange={handleCompletedTime}
          value={completedTime}
        />
      </div>
      <div className="flex space-x-3">
        <button className="border border-black" onClick={startTimer}>
          Start Timer
        </button>
        <button className="border border-black" onClick={resetTimer}>
          Reset Timer
        </button>
      </div>
      {startTime ? (
        <ProgressBar
          initialTime={startTime}
          targetTime={Number(startTime) + Number(estimatedTime)}
          estimatedDurationTime={estimatedTime}
          isComplete={isComplete}
        />
      ) : null}
    </div>
  )
}
