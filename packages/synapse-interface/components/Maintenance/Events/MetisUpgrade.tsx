import { AnnouncementBanner } from '../AnnouncementBanner'
import { WarningMessage } from '../../Warning'
import { useBridgeState } from '@/slices/bridge/hooks'
import { METIS } from '@/constants/chains/master'
import { useEventCountdownProgressBar } from '../EventCountdownProgressBar'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { getCountdownTimeStatus } from '../EventCountdownProgressBar'

/**
 * Start: 30 min prior to Metis Chain Downtime @ (March 14, 02:00 UTC)
 * End: 12 hours after start of Metis Chain Downtime
 */
// export const METIS_DOWNTIME_BANNERS_START = new Date(
//   Date.UTC(2024, 2, 14, 1, 0, 0)
// )
// export const METIS_DOWNTIME_START_DATE = new Date(
//   Date.UTC(2024, 2, 14, 1, 30, 0)
// )
// export const METIS_DOWNTIME_END_DATE = new Date(
//   Date.UTC(2024, 2, 14, 13, 30, 0)
// )

/** Remove after test */
export const METIS_DOWNTIME_BANNERS_START = new Date(
  Date.UTC(2024, 2, 13, 19, 36, 0)
)
export const METIS_DOWNTIME_START_DATE = new Date(
  Date.UTC(2024, 2, 13, 19, 36, 0)
)
export const METIS_DOWNTIME_END_DATE = new Date(
  Date.UTC(2024, 2, 13, 19, 38, 0)
)
/** Remove after test */

export const MetisDowntimeBanner = () => {
  const { isComplete } = getCountdownTimeStatus(
    METIS_DOWNTIME_BANNERS_START,
    METIS_DOWNTIME_END_DATE
  )

  useIntervalTimer(60000, isComplete)

  console.log('isComplete: ', isComplete)

  return (
    <AnnouncementBanner
      bannerId="03142024-metis-downtime"
      bannerContents={
        <div className="flex flex-col justify-center space-y-1 text-center">
          <div>
            Metis Chain bridging will be paused 30 min ahead of the Metis
            Upgrade (March 14, 02:00 UTC, 22:00 EST)
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
              Metis Chain bridging is paused until the Metis upgrade completes.
            </p>
          </>
        }
      />
    )
  } else return null
}

export const useMetisDowntimeCountdownProgress = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isChainMetis = [fromChainId, toChainId].includes(METIS.id)

  const {
    isPending: isMetisUpgradePending,
    EventCountdownProgressBar: MetisUpgradeCountdownProgressBar,
  } = useEventCountdownProgressBar(
    'Metis upgrade in progress',
    METIS_DOWNTIME_START_DATE,
    METIS_DOWNTIME_END_DATE
  )

  return {
    isMetisUpgradePending,
    isCurrentChainDisabled: isChainMetis && isMetisUpgradePending,
    MetisUpgradeCountdownProgressBar: isChainMetis
      ? MetisUpgradeCountdownProgressBar
      : null,
  }
}
