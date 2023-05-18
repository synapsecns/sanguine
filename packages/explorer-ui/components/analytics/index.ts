import { useEffect } from 'react'
import { UaEventOptions } from 'react-ga4/types/ga4'
import { getCLS, getFCP, getFID, getLCP, Metric } from 'web-vitals'

import GoogleAnalyticsProvider from './GoogleAnalyticsProvider'
import { useRouter } from "next/router";

const GOOGLE_ANALYTICS_CLIENT_ID_STORAGE_KEY = 'ga_client_id'
const GOOGLE_ANALYTICS_ID: string | undefined = process.env.NEXT_GOOGLE_ANALYTICS_ID

const googleAnalytics = new GoogleAnalyticsProvider()

export function sendEvent(event: string | UaEventOptions, params?: any) {
  return googleAnalytics.sendEvent(event, params)
}

export function outboundLink(
  {
    label,
  }: {
    label: string
  },
  hitCallback: () => unknown
) {
  return googleAnalytics.outboundLink({ label }, hitCallback)
}

export function sendTiming(timingCategory: any, timingVar: any, timingValue: any, timingLabel: any) {
  return googleAnalytics.gaCommandSendTiming(timingCategory, timingVar, timingValue, timingLabel)
}

if (typeof GOOGLE_ANALYTICS_ID === 'string') {
  googleAnalytics.initialize(GOOGLE_ANALYTICS_ID, {
    gaOptions: {
      storage: 'none',
      storeGac: false,
      clientId: undefined,
    },
  })
  if (window != undefined) {
    googleAnalytics.set({
      anonymizeIp: true,
      customBrowserType: !isMobile
        ? 'desktop'
        : 'web3' in window || 'ethereum' in window
          ? 'mobileWeb3'
          : 'mobileRegular',
    })
  }
} else {
  googleAnalytics.initialize('test', { gtagOptions: { debug_mode: true } })
}

function reportWebVitals({ name, delta, id }: Metric) {
  sendTiming('Web Vitals', name, Math.round(name === 'CLS' ? delta * 1000 : delta), id)
}

// tracks web vitals and pageviews
export function useAnalyticsReporter() {
  const { pathname, search } = useRouter()
  useEffect(() => {
    getFCP(reportWebVitals)
    getFID(reportWebVitals)
    getLCP(reportWebVitals)
    getCLS(reportWebVitals)
  }, [])

  useEffect(() => {
    googleAnalytics.pageview(`${pathname}${search}`)
  }, [pathname, search])

}
