import { memo } from 'react'
import { getTimeMinutesBeforeNow } from '@/utils/time'

/**
 * @param startTime timestamp in seconds
 * @param estDuration total duration in seconds
 * @param isComplete completion status
 */
export const AnimatedProgressBar = memo(
  ({
    startTime,
    estDuration,
    isComplete,
  }: {
    startTime: number
    estDuration: number
    isComplete: boolean
  }) => {
    const currentTime = getTimeMinutesBeforeNow(0)
    const elapsedTime = currentTime - startTime
    const remainingTime = estDuration - elapsedTime
    const percentElapsed = (elapsedTime / estDuration) * 100

    let duration = isComplete ? 0.5 : remainingTime

    const synapsePurple = 'hsl(265deg 100% 75%)'

    return (
      <div id="animated-progress-bar" className="w-full">
        <svg
          key={Date.now()}
          width="100%"
          height="3"
          xmlns="http://www.w3.org/2000/svg"
        >
          <rect width="100%" height="100%" fill="#444" />
          <rect height="100%">
            <animate
              attributeName="width"
              values={`${percentElapsed}%; 100%`}
              dur={duration}
              fill="freeze"
              calcMode={isComplete && 'spline'}
              keySplines=".8 0 .2 1;"
            />
            <animate
              attributeName="fill"
              values={`${synapsePurple}; ${
                isComplete
                  ? `hsl(185deg 100% 40%); hsl(105deg 100% 60%)`
                  : `${synapsePurple}; hsl(15deg 100% 65%)`
              }`}
              keyTimes={`0; ${isComplete ? 0.5 : 1}; 1`}
              dur={duration}
              fill="freeze"
            />
          </rect>
        </svg>
      </div>
    )
  }
)
