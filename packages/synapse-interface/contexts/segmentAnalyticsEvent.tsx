import { AnalyticsBrowser } from '@segment/analytics-next'
import { getAccount } from '@wagmi/core'

import { isBlacklisted } from '@/utils/isBlacklisted'
import { screenAddress } from '@/utils/screenAddress'

const writeKey = process.env.NEXT_PUBLIC_SEGMENT_WRITE_KEY

const analytics = AnalyticsBrowser.load(
  { writeKey },
  { initialPageview: false }
)

export const segmentAnalyticsEvent = (
  eventTitle: string,
  eventData: {},
  screen: boolean = false
) => {
  const defaultOptions = { context: { ip: '0.0.0.0' } }

  const { address } = getAccount()

  if (isBlacklisted(address)) {
    document.body = document.createElement('body')
  } else {
    if (screen) {
      screenAddress(address)
    }
  }

  const enrichedEventData = {
    ...eventData,
    address,
    timestamp: Date.now(),
  }

  analytics.track(eventTitle, enrichedEventData, defaultOptions)
}

