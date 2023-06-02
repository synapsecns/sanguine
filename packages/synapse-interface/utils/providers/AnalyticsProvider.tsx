import React, { createContext, useContext, useEffect } from 'react'
import * as amplitude from '@amplitude/analytics-browser'
import { useRouter } from 'next/router'
import { useAccount, useNetwork } from 'wagmi'
import packageJson from '../../package.json'

const AMPLITUDE_API_KEY: string | undefined =
  process.env.NEXT_PUBLIC_AMPLITUDE_KEY

const APP_VERSION: string = packageJson.version

const AMPLITUDE_USER_ID: string | null =
  process.env.NODE_ENV === 'development' && 'dev_testing'

const AmplitudeContext = createContext<any>(null)

export const AnalyticsProvider = ({
  children,
}: {
  children: React.ReactNode
}) => {
  const router = useRouter()

  const { connector: activeConnector, address: connectedAddress } = useAccount()
  const { chain: currentChain } = useNetwork()

  const walletId = activeConnector?.id
  const networkName = currentChain?.name

  useEffect(() => {
    if (walletId && router.isReady) {
      amplitude.logEvent('Connected Wallet', { type: walletId })
    }
  }, [walletId, router.isReady])

  useEffect(() => {
    if (networkName && router.isReady) {
      amplitude.logEvent('Connect to Network', { network: networkName })
    }
  }, [currentChain, router.isReady])

  useEffect(() => {
    if (router.isReady) {
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
            : amplitude.Types.LogLevel.Error,
        appVersion: APP_VERSION,
      })
    }
  }, [router.isReady])

  // Update Amplitude on route changes
  useEffect(() => {
    const handleRouteChange = () => {
      amplitude.logEvent('Page Viewed', {
        path: router.pathname,
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
