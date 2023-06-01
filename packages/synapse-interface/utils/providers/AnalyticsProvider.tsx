import React, { createContext, useContext, useEffect } from 'react'
import * as amplitude from '@amplitude/analytics-browser'
import { useRouter } from 'next/router'
import { logEvent } from '@amplitude/analytics-browser'

const AMPLITUDE_API_KEY = process.env.NEXT_PUBLIC_AMPLITUDE_KEY

const AmplitudeContext = createContext<any>(null)

export const AnalyticsProvider = ({
  children,
}: {
  children: React.ReactNode
}) => {
  const router = useRouter()

  useEffect(() => {
    if (router.isReady) {
      try {
        amplitude.init(AMPLITUDE_API_KEY, 'test@test.com', {
          defaultTracking: {
            sessions: true,
            pageViews: true,
            formInteractions: true,
            fileDownloads: true,
          },
        })
        console.log('amplitude initialized')
      } catch (error) {
        console.error('Error initializing Amplitude: ', error)
      }
    }
  }, [router.isReady])

  // Update Amplitude on route changes
  useEffect(() => {
    const handleRouteChange = () => {
      amplitude.logEvent('PageViewed', {
        path: router.pathname,
        // Add any additional data you want to track
      })
    }

    // Listen for route changes
    router.events.on('routeChangeComplete', handleRouteChange)

    // Clean up the listener when the component unmounts
    return () => {
      router.events.off('routeChangeComplete', handleRouteChange)
    }
  }, [router.events])

  return (
    <AmplitudeContext.Provider value={null}>
      {children}
    </AmplitudeContext.Provider>
  )
}

export const useAmplitude = (): any => useContext(AmplitudeContext)
