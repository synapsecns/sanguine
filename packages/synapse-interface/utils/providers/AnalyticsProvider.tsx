import React, { createContext, useContext, useEffect } from 'react'
import * as amplitude from '@amplitude/analytics-browser'
import { useRouter } from 'next/router'
import packageJson from '../../package.json'

const AMPLITUDE_API_KEY: string | undefined =
  process.env.NEXT_PUBLIC_AMPLITUDE_KEY

const APP_VERSION: string = packageJson.version

const AMPLITUDE_USER_ID: string | null =
  process.env.NODE_ENV === 'development' && 'test'

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
        amplitude.init(AMPLITUDE_API_KEY, AMPLITUDE_USER_ID, {
          defaultTracking: {
            sessions: true,
            pageViews: true,
            formInteractions: true,
            fileDownloads: true,
          },
          logLevel:
            process.env.NODE_ENV === 'development'
              ? amplitude.Types.LogLevel.Debug
              : amplitude.Types.LogLevel.None,
          appVersion: APP_VERSION,
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
      amplitude.logEvent('Page Viewed', {
        path: router.pathname,
        // Add any additional data you want to track
      })
      console.log(
        'Logged PageViewed on RouteChange in Amplitude: ',
        router.pathname
      )
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
