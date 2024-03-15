import { AnnouncementBanner } from '../../AnnouncementBanner'
import { WarningMessage } from '../../../Warning'
import { useBridgeState } from '@/slices/bridge/hooks'
import { OPTIMISM, BASE } from '@/constants/chains/master'
import {
  useEventCountdownProgressBar,
  getCountdownTimeStatus,
} from '../../EventCountdownProgressBar'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

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
  Date.UTC(2024, 2, 13, 23, 20, 0)
)
export const ECOTONE_FORK_START_DATE = new Date(
  Date.UTC(2024, 2, 13, 23, 35, 0)
)
export const ECOTONE_FORK_END_DATE = new Date(Date.UTC(2024, 2, 14, 0, 25, 0))

export const EcotoneForkUpgradeBanner = () => {
  const { isComplete } = getCountdownTimeStatus(
    ECOTONE_FORK_BANNERS_START,
    ECOTONE_FORK_END_DATE
  )

  useIntervalTimer(60000, isComplete)

  return (
    <AnnouncementBanner
      bannerId="03142024-ecotone-fork"
      bannerContents={
        <div className="flex flex-col justify-center space-y-1 text-center">
          <div>
            Optimism + Base Bridging will be paused 10 minutes ahead of Ecotone
            (March 14, 00:00 UTC, 20:00 EST).
          </div>
          <div>Will be back online shortly following the network upgrade.</div>
        </div>
      }
      startDate={ECOTONE_FORK_BANNERS_START}
      endDate={ECOTONE_FORK_END_DATE}
    />
  )
}

export const EcotoneForkWarningMessage = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isChainOptimism = [fromChainId, toChainId].includes(OPTIMISM.id)
  const isChainBase = [fromChainId, toChainId].includes(BASE.id)

  if (isChainOptimism || isChainBase) {
    return (
      <WarningMessage
        message={
          <>
            <p>
              Optimism Chain and Base Chain bridging are paused until the
              Ecotone Fork upgrade completes.
            </p>
          </>
        }
      />
    )
  } else return null
}

/**
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

  const isChainOptimism = [fromChainId, toChainId].includes(OPTIMISM.id)
  const isChainBase = [fromChainId, toChainId].includes(BASE.id)

  const {
    isPending: isEcotoneForkUpgradePending,
    EventCountdownProgressBar: EcotoneForkCountdownProgressBar,
  } = useEventCountdownProgressBar(
    'Ecotone Fork upgrade in progress',
    ECOTONE_FORK_START_DATE,
    ECOTONE_FORK_END_DATE
  )

  return {
    isEcotoneForkUpgradePending,
    isCurrentChainDisabled:
      (isChainOptimism || isChainBase) && isEcotoneForkUpgradePending,
    EcotoneForkCountdownProgressBar:
      isChainOptimism || isChainBase ? EcotoneForkCountdownProgressBar : null,
  }
}
