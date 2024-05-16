import { memo } from 'react'
import { isNull } from 'lodash'

import { getCountdownTimeStatus } from '../helpers/getCountdownTimeStatus'

/**
 * Constructs animated progress bar visualizing time progress of an event.
 * If end date is not provided, progress bar animates indefinitely.
 *
 * @param id - A unique ID to distinguish instances.
 * @param startDate - The start date of the tracked event.
 * @param endDate - The end date of the tracked event.
 */
export const LinearAnimatedProgressBar = memo(
  ({
    id,
    startDate,
    endDate,
  }: {
    id: string
    startDate: Date
    endDate: Date | null
  }) => {
    const height = 3
    const maskId = `mask-${id}`
    const progressId = `progress-${id}`
    const greenColor = 'rgb(74 222 128)'
    const purpleColor = 'hsl(265deg 100% 75%)'

    let duration: string | number

    const isIndefinite = isNull(endDate)

    if (isIndefinite) {
      duration = 'infinite'
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
            <linearGradient
              id={progressId}
              spreadMethod="reflect"
              x1="0"
              x2="1"
            >
              <stop stopColor={purpleColor} />
              <stop stopColor={purpleColor} offset=".25" />
              <stop stopColor={purpleColor} stopOpacity=".67" offset=".75" />
              <stop stopColor={purpleColor} stopOpacity=".67" offset="1" />
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
                  values={`100%; 100%`}
                  dur="infinite"
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
          ></rect>
        </svg>
      )
    } else {
      const {
        totalTimeInSeconds,
        totalTimeElapsedInSeconds,
        totalTimeRemainingInSeconds,
        isComplete,
      } = getCountdownTimeStatus(startDate, endDate)

      const percentElapsed = Math.floor(
        (totalTimeElapsedInSeconds / totalTimeInSeconds) * 100
      )

      duration = isComplete ? 0.5 : totalTimeRemainingInSeconds

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
            <linearGradient
              id={progressId}
              spreadMethod="reflect"
              x1="0"
              x2="1"
            >
              <stop stopColor={purpleColor} />
              <stop stopColor={purpleColor} offset=".25" />
              <stop stopColor={purpleColor} stopOpacity=".67" offset=".75" />
              <stop stopColor={purpleColor} stopOpacity=".67" offset="1" />
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
                values={`${purpleColor}; hsl(185deg 100% 40%); ${greenColor}`}
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
  }
)
