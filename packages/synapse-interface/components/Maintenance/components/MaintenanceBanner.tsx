import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { AnnouncementBanner } from './AnnouncementBanner'

/**
 * Component for creating and managing a maintenance announcement banner. This banner automatically
 * appears and disappears based on specified start and end times, providing users with timely information
 * about maintenance events.
 *
 * @param {string} id - A unique identifier for the banner instance. This is used to track the banner's state
 *                      in the browser and avoid conflicts with other instances.
 * @param {Date} startDate - The starting date and time when the banner should become visible to users.
 *                           This is the point at which the maintenance is considered to begin.
 * @param {Date} endDate - The ending date and time when the banner should be removed or hidden from view.
 *                         This corresponds to the end of the maintenance period.
 * @param {any} bannerMessage - The content to be displayed within the banner. This parameter allows for
 *                              flexibility in the message's structure and content, which can include text,
 *                              links, or even React components.
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
