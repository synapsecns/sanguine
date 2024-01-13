import { useEffect, useState } from 'react'

const BANNER_VERSION = '3'
const bannerClassName = "flex items-center p-3 border border-purple-400 rounded max-w-lg mx-auto"

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
      <div
        id="banner-default"
        className={bannerClassName}
        role="alert"
        style={{
          background:
            'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
          borderRadius: '10px',
        }}
      >
        {/* TODO: replace w/ tweetlink directly */}
        <a
          href={
            'https://twitter.com/SynapseProtocol/status/1726659540551352387'
          }
          target={'_blank'}
        >
          Synapse now supports Solana at solana.synapseprotocol.com
        </a>
        <button
          type="button"
          className={`
            inline-flex items-center justify-center
            h-7 w-7
            ml-auto -mx-1.5 -my-1.5 p-1.5
          `}
          data-dismiss-target="#banner-default"
          aria-label="Close"
          onClick={() => setShowBanner(false)}
        >
          <svg
            width="9px"
            height="9px"
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
  )
}

export const CustomBanner = ({ text, link }: { text: string; link?: any }) => {
  const [hasMounted, setHasMounted] = useState(false)
  const [showBanner, setShowBanner] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    if (hasMounted) {
      const storedBannerVersion = localStorage.getItem('customBannerVersion')
      const storedShowBanner = localStorage.getItem('customShowBanner')

      setShowBanner(
        storedBannerVersion !== BANNER_VERSION ||
          storedShowBanner === null ||
          storedShowBanner === 'true'
      )
    }
  }, [hasMounted])

  useEffect(() => {
    if (hasMounted) {
      localStorage.setItem('customShowBanner', showBanner.toString())
      localStorage.setItem('customBannerVersion', BANNER_VERSION)
    }
  }, [showBanner, hasMounted])

  if (!showBanner || !hasMounted) return null

  return (
    <div
      id="banner-default"
      className={bannerClassName}
      role="alert"
      style={{
        background:
          'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
        borderRadius: '10px',
      }}
    >
      <a href={link} target={'_blank'}>
        {text}
      </a>
      <button
        type="button"
        className={`
          inline-flex items-center justify-center
          h-7 w-7
          ml-auto -mx-1.5 -my-1.5 p-1.5
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
  )
}

export const InterruptedServiceBanner = () => {
  const [hasMounted, setHasMounted] = useState(false)
  const [showBanner, setShowBanner] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  useEffect(() => {
    if (hasMounted) {
      const storedBannerVersion = localStorage.getItem('bannerVersion')
      const storedShowBanner = localStorage.getItem('showInterruptedBanner')

      setShowBanner(
        storedBannerVersion !== BANNER_VERSION ||
          storedShowBanner === null ||
          storedShowBanner === 'true'
      )
    }
  }, [hasMounted])

  useEffect(() => {
    if (hasMounted) {
      localStorage.setItem('showInterruptedBanner', showBanner.toString())
      localStorage.setItem('bannerVersion', BANNER_VERSION)
    }
  }, [showBanner, hasMounted])

  if (!showBanner || !hasMounted) return null

  return (
    <div
      id="banner-default"
      className={bannerClassName}
      role="alert"
      style={{
        background:
          'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
      }}
    >
      <p>
        Synapse Explorer and the transaction watcher may not appear
        during planned maintenance on 2023-11-01 from 5am - 7pm UTC.
        <br className="block lg:hidden" />
        <br className="block lg:hidden" />
        <div className="hidden lg:inline"> </div>
        Transactions will still go through as expected. Please confirm
        transactions using the native explorer for your destination
        chain during this time.
      </p>
      <button
        type="button"
        className={`
          inline-flex items-center justify-center
          h-7 w-7
          ml-auto -mx-1.5 -my-1.5 p-1.5
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
  )
}
