import { useEffect, useState } from 'react'

const BANNER_VERSION = '1'

export const Banner = () => {
  const [hasMounted, setHasMounted] = useState(false)
  const [showBanner, setShowBanner] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    if (hasMounted) {
      const storedBannerVersion = localStorage.getItem('bannerVersion')
      const storedShowBanner = localStorage.getItem('showBanner')

      setShowBanner(
        storedBannerVersion !== BANNER_VERSION ||
          storedShowBanner === null ||
          storedShowBanner === 'true'
      )
    }
  }, [hasMounted])

  useEffect(() => {
    if (hasMounted) {
      localStorage.setItem('showBanner', showBanner.toString())
      localStorage.setItem('bannerVersion', BANNER_VERSION)
    }
  }, [showBanner, hasMounted])

  if (!showBanner || !hasMounted) return null

  return (
    <div className="flex items-center justify-center px-4 mx-auto mt-4 lg:flex-row 2xl:w-3/4 sm:px-8 md:px-12">
      <div
        id="banner-default"
        className="flex items-center pl-3 pr-3 pt-1 pb-1 border border-[#AC8FFF] rounded-md text-primaryTextColor"
        role="alert"
        style={{
          background:
            'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
          borderRadius: '10px',
        }}
      >
        <div className="m-1 font-thin">
          <div className="container mx-auto">
            <p className="text-md">
              New! Chain and token selections now list routable options first.
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
