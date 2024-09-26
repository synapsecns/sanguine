import { memo } from 'react'
import { getUnixTimeMinutesBeforeNow } from '@/utils/time'

/**
 * @param {string} id - The unique ID of a rendered instance.
 * @param {number} startTime - The start time of an event as a timestamp.
 * @param {number} duration - The estimated total duration of an event.
 * @param {'pending' | 'completed' | 'reverted' | 'refunded'} status - The current status of an event.
 */
export const AnimatedProgressBar = memo(
  ({
    id,
    startTime,
    estDuration,
    status,
  }: {
    id: string
    startTime: number
    estDuration: number
    status: 'pending' | 'completed' | 'reverted' | 'refunded'
  }) => {
    const currentTime = getUnixTimeMinutesBeforeNow(0)
    const elapsedTime = currentTime - startTime
    const remainingTime = estDuration - elapsedTime
    const percentElapsed = (elapsedTime / estDuration) * 100

    const isComplete = status === 'completed'
    const isError = status === 'reverted' || status === 'refunded'

    let duration = isComplete ? 0.5 : remainingTime

    const synapsePurple = 'hsl(265deg 100% 75%)'
    const tailwindGreen400 = 'rgb(74 222 128)'
    const yellowText = '#FFE14D'

    const height = 3

    const progressId = `progress-${id}`
    const maskId = `mask-${id}`

    return (
      <svg
        id="animated-progress-bar"
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
                values={`${
                  isComplete || isError ? 100 : percentElapsed
                }%; 100%`}
                dur={duration}
                fill="freeze"
                calcMode={isComplete || isError ? 'spline' : null}
                keySplines=".8 0 .2 1"
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
          {isError && (
            <animate
              attributeName="fill"
              values={`${yellowText}; ${yellowText}; ${yellowText}`}
              keyTimes="0; .5; 1"
              dur={duration}
              fill="freeze"
            />
          )}
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
