import { useState, useEffect, ReactNode } from 'react'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'

/**
 * Message banner that renders between defined start <> end dates.
 *
 * @param {string} bannerId - The unique ID assigned to banner instance to prevent collisions. ID Format: $MMDDYYYY-$BANNER_NAME format (e.g 03132024-ETH-DENCUN)
 * @param {any} bannerContent - The content to display in the banner.
 * @param {Date} startDate - The start date that initiates rendering banner.
 * @param {Date | null} endDate - The end date that removes banner. If null, the banner will render indefinitely.
 */
export const AnnouncementBanner = ({
  bannerId,
  bannerContent,
  startDate,
  endDate,
}: {
  bannerId: string
  bannerContent: any
  startDate: Date
  endDate: Date | null
}) => {
  const { isStarted, isComplete } = getCountdownTimeStatus(startDate, endDate)

  const [hasMounted, setHasMounted] = useState(false)
  const [showBanner, setShowBanner] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    if (hasMounted && isStarted && !isComplete) {
      const storedShowBanner = localStorage.getItem('showAnnouncementBanner')
      const storedBannerId = localStorage.getItem('bannerId')

      setShowBanner(
        Boolean(
          storedBannerId !== bannerId ||
            storedShowBanner === null ||
            storedShowBanner === 'true'
        )
      )
    }
  }, [hasMounted])

  useEffect(() => {
    if (hasMounted && isStarted && !isComplete) {
      localStorage.setItem('showAnnouncementBanner', showBanner.toString())
      localStorage.setItem('bannerId', bannerId)
    }
  }, [showBanner, hasMounted])

  if (!showBanner || !hasMounted || !isStarted || isComplete) return null

  return (
    <div
      className={`
        flex items-center justify-center mx-auto text-sm
        text-left lg:flex-row bg-gradient-to-r
        from-fuchsia-600/25 to-purple-600/25
      `}
    >
      <div
        id={bannerId}
        className={`
          flex gap-4 py-1 w-full
          justify-center leading-normal items-center
          max-w-[1111px] text-primaryTextColor
        `}
        role="alert"
      >
        {bannerContent}
        <button
          onClick={() => setShowBanner(false)}
          className="inline-flex items-center justify-center p-3 text-primaryTextColor hover:opacity-70"
          data-dismiss-target="#banner-default"
          aria-label="Close"
          type="button"
        >
          <svg
            className="m-auto"
            width={10}
            height={10}
            viewBox="0 0 14 14"
            xmlns="http://www.w3.org/2000/svg"
            aria-hidden="true"
            fill="none"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
            />
          </svg>
        </button>
      </div>
    </div>
  )
}
