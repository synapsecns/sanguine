import { AnnouncementBanner } from './AnnouncementBanner'
import { WarningMessage } from '../Warning'
import { useBridgeState } from '@/slices/bridge/hooks'
import { METIS } from '@/constants/chains/master'

/**
 * Start: 30 min prior to Metis Chain Downtime @ (March 14, 02:00 UTC)
 * End: 12 hours after start of Metis Chain Downtime
 */
export const METIS_DOWNTIME_BANNERS_START = new Date(
  Date.UTC(2024, 2, 14, 1, 0, 0)
)
export const METIS_DOWNTIME_START_DATE = new Date(
  Date.UTC(2024, 2, 14, 1, 30, 0)
)
export const METIS_DOWNTIME_END_DATE = new Date(
  Date.UTC(2024, 2, 14, 13, 30, 0)
)

export const MetisDowntimeBanner = () => {
  return (
    <AnnouncementBanner
      bannerId="03142024-metis-downtime"
      bannerContents={
        <div className="flex flex-col justify-center space-y-1 text-center">
          <div>
            Due to a Metis upgrade, bridging to and from Metis will pause 30
            minutes ahead of March 14, 02:00 UTC,
          </div>
          <div>and stay paused for ~12 hours.</div>
        </div>
      }
      startDate={METIS_DOWNTIME_BANNERS_START}
      endDate={METIS_DOWNTIME_END_DATE}
    />
  )
}

export const MetisDowntimeWarningMessage = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isChainMetis = [fromChainId, toChainId].includes(METIS.id)

  if (isChainMetis) {
    return (
      <WarningMessage
        message={
          <>
            <p>
              Metis Chain bridging is paused until a planned Metis upgrade
              completes.
            </p>
          </>
        }
      />
    )
  } else return null
}
