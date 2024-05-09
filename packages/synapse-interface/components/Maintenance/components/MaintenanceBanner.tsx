import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { AnnouncementBanner } from './AnnouncementBanner'
import { isNull } from 'lodash'

/**
 * Handles constructing and updating instances of Announcement Banners based on specific start / end time.
 * If end date is null, indefinitely display banner.
 *
 * @param id Unique ID to prevent conflicts with other banner instances.
 *           Assign ID $MMDDYYYY-$BANNER_NAME format (e.g 03132024-ETH-DENCUN)
 * @param startDate Start time to display banner
 * @param endDate End time to remove banner
 * @param bannerMessage - Message to display in banner
 */
export const MaintenanceBanner = ({
  id,
  startDate,
  endDate,
  bannerMessage,
  disabled = false,
}: {
  id: string
  startDate: Date
  endDate: Date | null
  bannerMessage: any
  disabled?: boolean
}) => {
  const { isComplete } = getCountdownTimeStatus(startDate, endDate)

  const isIndefinite = isNull(endDate)

  useIntervalTimer(60000, isComplete || isIndefinite || disabled)

  if (disabled) {
    return null
  } else {
    return (
      <AnnouncementBanner
        bannerId={id}
        bannerContents={bannerMessage}
        startDate={startDate}
        endDate={endDate}
      />
    )
  }
}
