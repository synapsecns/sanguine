import { shortenAddress } from '@/utils/shortenAddress'
import { AnalyticsBrowser } from '@segment/analytics-next'
import { getAccount } from '@wagmi/core'
import { createContext, useContext, useMemo } from 'react'

const writeKey = process.env.NEXT_PUBLIC_SEGMENT_WRITE_KEY

const SegmentAnalyticsContext = createContext(undefined)

export const analytics = AnalyticsBrowser.load(
  { writeKey },
  { initialPageview: false }
)

export const segmentAnalyticsEvent = (eventTitle: string, eventData: {}) => {
  const defaultOptions = { context: { ip: '0.0.0.0' } }

  const account = getAccount()
  const { address } = account

  const enrichedEventData = {
    ...eventData,
    address,
    timestamp: Date.now(),
  }

  const showAddress = address ? shortenAddress(address) : 'No address'
  const enrichedEventTitle = `[${showAddress}] ${eventTitle}`

  analytics.track(enrichedEventTitle, enrichedEventData, defaultOptions)
}

export const SegmentAnalyticsProvider = ({ children }) => (
  <SegmentAnalyticsContext.Provider value={analytics}>
    {children}
  </SegmentAnalyticsContext.Provider>
)

export const useAnalytics = () => {
  const result = useContext(SegmentAnalyticsContext)
  if (!result) {
    throw new Error('Context used outside of its Provider!')
  }
  return result
}
