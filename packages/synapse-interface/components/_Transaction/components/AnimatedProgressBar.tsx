import { memo } from 'react'
import { getTimeMinutesBeforeNow } from '@/utils/time'

/**
 * @param id unique identifier for progress bar instance
 * @param startTime timestamp in seconds
 * @param estDuration total duration in seconds
 * @param isComplete completion status
 */
export const AnimatedProgressBar = memo(
  ({
    id,
    startTime,
    estDuration,
    isComplete,
  }: {
    id: string
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
    const height = 3

    const progressId = `progress-${id}`
    const maskId = `mask-${id}`

    return (
      <div id="animated-progress-bar" className="absolute right-1 left-1">
        <svg
          key={progressId}
          width="100%"
          height={height}
          xmlns="http://www.w3.org/2000/svg"
          className="rounded-sm"
          style={{ background: '#444' }}
        >
          <defs>
            <linearGradient
              id={progressId}
              spreadMethod="reflect"
              x1="0"
              x2="1"
            >
              <stop stop-color={synapsePurple} />
              <stop stop-color={synapsePurple} offset=".25" />
              <stop
                stop-color={synapsePurple}
                stop-opacity=".67"
                offset=".75"
              />
              <stop stop-color={synapsePurple} stop-opacity=".67" offset="1" />
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
                  values={`${percentElapsed}%; 100%`}
                  dur={duration}
                  fill="freeze"
                  calcMode={isComplete && 'spline'}
                  keySplines=".8 0 .2 1;"
                />
              </rect>
            </clipPath>
          </defs>
          <rect
            width="100%"
            height={height}
            fill={`url(#${progressId})`}
            clip-path={`url(#${maskId})`}
          >
            {isComplete && (
              <animate
                attributeName="fill"
                values={`${synapsePurple}; ${
                  isComplete
                    ? `hsl(185deg 100% 40%); hsl(105deg 100% 60%)`
                    : `${synapsePurple}; hsl(15deg 100% 65%)`
                }`}
                keyTimes={`0; .5; 1`}
                dur={duration}
                fill="freeze"
              />
            )}
          </rect>
        </svg>
      </div>
    )
  }
)
