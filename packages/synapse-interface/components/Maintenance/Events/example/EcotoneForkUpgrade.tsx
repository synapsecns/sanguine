import { useBridgeState } from '@/slices/bridge/hooks'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { METIS } from '@/constants/chains/master'
import {
  useEventCountdownProgressBar,
  getCountdownTimeStatus,
} from '../../EventCountdownProgressBar'
import { AnnouncementBanner } from '../../AnnouncementBanner'
import { WarningMessage } from '../../../Warning'

/**
 * Leaving this file to serve as an example for how to create
 * automated annoucement banners and event countdown timer bars
 * with the ability to pause Bridge by selected chain ids
 */

/**
 * Start: 15 min prior to Ecotone Fork Upgrade Time @ (March 14, 00:00 UTC)
 * End: 50 min after start of Ecotone Fork Upgrade Time
 */
export const ECOTONE_FORK_BANNERS_START = new Date(
  Date.UTC(2024, 2, 20, 16, 0, 0)
)
export const ECOTONE_FORK_START_DATE = new Date(Date.UTC(2024, 2, 20, 18, 0, 0))
export const ECOTONE_FORK_END_DATE = new Date(Date.UTC(2024, 2, 20, 21, 0, 0))

/** Previous implementation can be seen here: https://github.com/synapsecns/sanguine/pull/2294/files#diff-bbe6298d3dfbc80e46e2ff8b399a3e1822cede80f392b1af91875145ad4eeb19R19 */
export const EcotoneForkUpgradeBanner = () => {
  const { isComplete } = getCountdownTimeStatus(
    ECOTONE_FORK_BANNERS_START, // Banner will automatically appear after start time
    ECOTONE_FORK_END_DATE // Banner will automatically disappear when end time is reached
  )

  useIntervalTimer(60000, isComplete)

  return (
    <AnnouncementBanner
      bannerId="03202024-metis-pause"
      bannerContents={
        <>Metis Bridging is paused. Will be back online shortly.</>
      }
      startDate={ECOTONE_FORK_BANNERS_START}
      endDate={ECOTONE_FORK_END_DATE}
    />
  )
}

/**
 * Warning Message to place within the Bridge Card
 * Below example sets to show only when chains are selected (Optimism, Base)
 *
 * Example: https://github.com/synapsecns/sanguine/blob/f068eff5e86ec97e17fc8e703d7203c12fb7f733/packages/synapse-interface/pages/state-managed-bridge/index.tsx#L629
 */
export const EcotoneForkWarningMessage = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isToChainMetis = toChainId === METIS.id

  if (isToChainMetis) {
    return (
      <WarningMessage
        message={
          <>
            <p>Metis Chain bridging are paused.</p>
          </>
        }
      />
    )
  } else return null
}

/**
 * Countdown Bar with Progress Animation based on time remaining
 * Below example sets to show only when chains are selected (Optimism, Base)
 *
 * Previously used in this location: https://github.com/synapsecns/sanguine/blob/f068eff5e86ec97e17fc8e703d7203c12fb7f733/packages/synapse-interface/pages/state-managed-bridge/index.tsx#L588
 * Bridge pause implemented here: https://github.com/synapsecns/sanguine/blob/f068eff5e86ec97e17fc8e703d7203c12fb7f733/packages/synapse-interface/pages/state-managed-bridge/index.tsx#L652-L654
 *
 * Example of how to use hook:
 * import { useEcotoneForkCountdownProgress } = '@/components/Maintenance/Events/EcotoneForkUpgrade'
 *
 * const {
    isEcotoneForkUpgradePending,
    isCurrentChainDisabled: isEcotoneUpgradeChainsDisabled,
    EcotoneForkCountdownProgressBar,
  } = useEcotoneForkCountdownProgress()

  In JSX, render the component:

    <div>
      {EcotoneForkCountdownProgressBar}
    </div>
 */

export const useEcotoneForkCountdownProgress = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isToChainMetis = toChainId === METIS.id

  const {
    isPending: isEcotoneForkUpgradePending,
    EventCountdownProgressBar: EcotoneForkCountdownProgressBar,
  } = useEventCountdownProgressBar(
    'Metis Bridging is paused',
    ECOTONE_FORK_START_DATE, // Countdown Bar will automatically appear after start time
    ECOTONE_FORK_END_DATE // Countdown Bar will automatically disappear when end time is reached
  )

  return {
    isEcotoneForkUpgradePending,
    isCurrentChainDisabled: isToChainMetis, // Used to pause Bridge
    EcotoneForkCountdownProgressBar: isToChainMetis
      ? EcotoneForkCountdownProgressBar
      : null,
  }
}
