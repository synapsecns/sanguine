import { useState, useEffect } from 'react'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'

/**
 * Reusable Annoucement Banner with custom Start/End Time
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
  const [showBanner, setShowBanner] = useState(false)

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

  if (!showBanner || !hasMounted || isComplete) return null

  return (
    <div
      className="flex items-center justify-center mx-auto mb-3 lg:flex-row"
      style={{
        background:
          'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
      }}
    >
      <div
        id="banner-default"
        className="flex items-center px-8 pt-1 pb-1 rounded-md text-primaryTextColor"
        role="alert"
      >
        <div className="container p-1 mx-auto">
          <p className="text-md">{bannerContents}</p>
        </div>
        <button
          type="button"
          className={`
            inline-flex items-center justify-center
            h-7 w-7 p-1.5 ml-3
            text-primaryTextColor
          `}
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
