import { useState, useEffect } from 'react'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'

/**
 * Reusable automated Announcement Banner with custom Start/End Time
 * Will automatically appear after Start time
 * Will automatically disappear after End time
 * @param bannerId: store in $MMDDYYYY-$BANNER_NAME format (e.g 03132024-ETH-DENCUN)
 * @param bannerContents: contents to display in banner
 * @param startDate: start date to show banner
 * @param endDate: end date to remove banner
 */
export const AnnouncementBanner = ({
  bannerId,
  bannerContents,
  startDate,
  endDate,
}: {
  bannerId: string
  bannerContents: any
  startDate: Date
  endDate: Date
}) => {
  const { isStarted, isComplete } = getCountdownTimeStatus(startDate, endDate)
  const [hasMounted, setHasMounted] = useState(false)
  const [showBanner, setShowBanner] = useState(true)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    if (hasMounted && isStarted && !isComplete) {
      const storedShowBanner = localStorage.getItem('showAnnoucementBanner')
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
      localStorage.setItem('showAnnoucementBanner', showBanner.toString())
      localStorage.setItem('bannerId', bannerId)
    }
  }, [showBanner, hasMounted])

  if (!showBanner || !hasMounted || !isStarted || isComplete) return null

  return (
    <div className="flex items-center justify-center mx-auto text-sm text-left lg:flex-row bg-gradient-to-r from-fuchsia-600/25 to-purple-600/25">
      <div
        id="banner-default"
        className="flex items-center gap-4 px-4 py-2 w-full max-w-[1111px] text-primaryTextColor justify-between leading-normal"
        role="alert"
      >
        {bannerContents}
        <button
          type="button"
          className="inline-flex items-center justify-center text-primaryTextColor"
          data-dismiss-target="#banner-default"
          aria-label="Close"
          onClick={() => setShowBanner(false)}
        >
          <svg
            className="w-[9px] h-[9px]"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 14 14"
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
