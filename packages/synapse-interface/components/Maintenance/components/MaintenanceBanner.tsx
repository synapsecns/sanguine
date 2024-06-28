import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { AnnouncementBanner } from './AnnouncementBanner'
import { isNull } from 'lodash'

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
        bannerContent={bannerMessage}
        startDate={startDate}
        endDate={endDate}
      />
    )
  }
}
