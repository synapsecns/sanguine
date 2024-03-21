import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { AnnouncementBanner } from './AnnouncementBanner'

/**
 * Creates an automatic Annoucement Banner based on start/end times.
 *
 * @param {string} id - Unique id that determines instance of banner to track in browser
 * @param {Date} startDate - Date that automatically triggers displaying banner
 * @param {Date} endDate - Date that automatically triggers removing banner
 * @param {any} bannerMessage - Allow for flexibility when constructing banner message
 */
export const MaintenanceBanner = ({
  id,
  startDate,
  endDate,
  bannerMessage,
}: {
  id: string
  startDate: Date
  endDate: Date
  bannerMessage: any
}) => {
  const { isComplete } = getCountdownTimeStatus(startDate, endDate)

  useIntervalTimer(60000, isComplete)

  return (
    <AnnouncementBanner
      bannerId={id}
      bannerContents={bannerMessage}
      startDate={startDate}
      endDate={endDate}
    />
  )
}
