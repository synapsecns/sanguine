import { useState, useEffect } from 'react'

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
  const [hasMounted, setHasMounted] = useState(false)
  const [showBanner, setShowBanner] = useState(false)

  const currentDate = new Date()
  const isStarted =
    Math.floor(currentDate.getTime()) - Math.floor(startDate.getTime()) > 0
  const isComplete =
    Math.floor(currentDate.getTime()) - Math.floor(endDate.getTime()) > 0

  console.log('currentDate:', currentDate)
  console.log('endDate:', endDate)

  console.log('isStarted: ', isStarted)
  console.log('isComplete:', isComplete)

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
      className="flex items-center justify-center mx-auto lg:flex-row lg:px-20"
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
        <div className="m-1 font-thin">
          <div className="container mx-auto">
            <p className="text-md">
              {/* <div>
                The Bridge + RFQ will be globally offline 15mins ahead of the
                Dencun upgrade (March 13, 13:55 UTC, 9:55 EST). Will be back
                online about 15 - 30 mins after Dencun.
              </div> */}
              {bannerContents}
            </p>
          </div>
        </div>
        <button
          type="button"
          className={`
            inline-flex items-center justify-center
            h-7 w-7
            ml-auto -mx-1.5 -my-1.5 p-1.5
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
