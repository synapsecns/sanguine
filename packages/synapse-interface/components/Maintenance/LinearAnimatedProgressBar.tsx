import { memo } from 'react'
import { getTimeMinutesBeforeNow } from '@/utils/time'

/**
 * @param id unique identifier for progress bar instance
 * @param startTime start time in unix seconds
 * @param endTime end time in unix seconds
 * @param status progress status
 */
export const LinearAnimatedProgressBar = memo(
  ({
    id,
    startTime,
    endTime,
    status,
  }: {
    id: string
    startTime: number
    endTime: number
    status: 'idle' | 'pending' | 'complete'
  }) => {
    const currentTime = Math.floor(getTimeMinutesBeforeNow(0))
    const elapsedTimeInSeconds = currentTime - startTime
    const remainingTimeInSeconds = endTime - currentTime
    const totalTimeInSeconds = endTime - startTime

    const percentElapsed = Math.floor(
      (elapsedTimeInSeconds / totalTimeInSeconds) * 100
    )

    const isComplete = status === 'complete'

    let duration = isComplete ? 0.5 : remainingTimeInSeconds

    const synapsePurple = 'hsl(265deg 100% 75%)'
    const tailwindGreen400 = 'rgb(74 222 128)'
    const height = 3
    const progressId = `progress-${id}`
    const maskId = `mask-${id}`

    return (
      <svg
        id="linear-animated-progress-bar"
        key={Date.now()}
        width="100%"
        height={height}
        xmlns="http://www.w3.org/2000/svg"
        className="rounded-sm"
      >
        <defs>
          <linearGradient id={progressId} spreadMethod="reflect" x1="0" x2="1">
            <stop stopColor={synapsePurple} />
            <stop stopColor={synapsePurple} offset=".25" />
            <stop stopColor={synapsePurple} stopOpacity=".67" offset=".75" />
            <stop stopColor={synapsePurple} stopOpacity=".67" offset="1" />
            <animate
              attributeName="x1"
              values="0%; -6%"
              dur=".67s"
              repeatCount="indefinite"
            />
            <animate
              attributeName="x2"
              values="3%; -3%"
              dur=".67s"
              repeatCount="indefinite"
            />
          </linearGradient>
          <clipPath id={maskId}>
            <rect height="100%">
              <animate
                attributeName="width"
                values={`${isComplete ? 100 : percentElapsed}%; 100%`}
                dur={totalTimeInSeconds}
                fill="freeze"
                calcMode={'linear'}
              />
            </rect>
          </clipPath>
        </defs>
        <rect
          width="100%"
          height={height}
          fill={`url(#${progressId})`}
          clipPath={`url(#${maskId})`}
        >
          {isComplete && (
            <animate
              attributeName="fill"
              values={`${synapsePurple}; hsl(185deg 100% 40%); ${tailwindGreen400}`}
              keyTimes="0; .5; 1"
              dur={duration}
              fill="freeze"
            />
          )}
        </rect>
        {isComplete && (
          <animate
            attributeName="height"
            values={`${height}; ${height}; 0`}
            keyTimes="0; .5; 1"
            calcMode="spline"
            keySplines="0 0 1 1; .8 0 .2 1"
            dur={duration * 1.5}
            fill="freeze"
          />
        )}
      </svg>
    )
  }
)
