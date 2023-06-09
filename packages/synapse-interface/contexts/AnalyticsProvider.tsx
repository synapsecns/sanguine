import { AnalyticsBrowser } from '@segment/analytics-next'
import { createContext, useContext, useMemo } from 'react'

const AnalyticsContext = createContext(undefined)

export const AnalyticsProvider = ({ children, writeKey }) => {
  const analytics = useMemo(
    () => AnalyticsBrowser.load({ writeKey }, { initialPageview: true }),
    [writeKey]
  )

  return (
    <AnalyticsContext.Provider value={analytics}>
      {children}
    </AnalyticsContext.Provider>
  )
}

export const useAnalytics = () => {
  const result = useContext(AnalyticsContext)
  if (!result) {
    throw new Error('Context used outside of its Provider!')
  }
  return result
}
