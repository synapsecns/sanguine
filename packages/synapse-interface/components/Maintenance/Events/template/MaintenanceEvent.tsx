import { useBridgeState } from '@/slices/bridge/hooks'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { OPTIMISM, BASE } from '@/constants/chains/master'
import {
  useEventCountdownProgressBar,
  getCountdownTimeStatus,
} from '../../EventCountdownProgressBar'
import { AnnouncementBanner } from '../../AnnouncementBanner'
import { WarningMessage } from '../../../Warning'

/**
 * Edit this file for Website Maintenance, components already placed on Bridge page
 *
 * If require multiple maintenance events, create another file using this file as a template
 * and add another instance of components on relevant pages
 */

/** Banner start time */
const MAINTENANCE_BANNERS_START = new Date(Date.UTC(2024, 2, 20, 20, 20, 0))
/** Countdown Progress Bar, Bridge Warning Message + Bridge Pause start time */
const MAINTENANCE_START_DATE = new Date(Date.UTC(2024, 2, 20, 20, 20, 0))
/** Ends Banner, Countdown Progress Bar, Bridge Warning Message, Bridge Pause */
const MAINTENANCE_END_DATE = new Date(Date.UTC(2024, 2, 20, 22, 0, 0))

export const MaintenanceBanner = () => {
  const { isComplete } = getCountdownTimeStatus(
    MAINTENANCE_BANNERS_START, // Banner will automatically appear after start time
    MAINTENANCE_END_DATE // Banner will automatically disappear when end time is reached
  )

  useIntervalTimer(60000, isComplete)

  return (
    <AnnouncementBanner
      bannerId="03202024-maintenance-banner"
      bannerContents={
        <>
          <p className="m-auto">
            Bridging is paused until maintenance is complete.
          </p>
        </>
      }
      startDate={MAINTENANCE_BANNERS_START}
      endDate={MAINTENANCE_END_DATE}
    />
  )
}

export const MaintenanceWarningMessage = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isWarningChain = isChainIncluded(
    [fromChainId, toChainId],
    [OPTIMISM.id, BASE.id] // Update for Chains to show warning on
  )

  const { isComplete } = getCountdownTimeStatus(
    MAINTENANCE_BANNERS_START, // Banner will automatically appear after start time
    MAINTENANCE_END_DATE // Banner will automatically disappear when end time is reached
  )

  if (isComplete) return null

  if (isWarningChain) {
    return (
      <WarningMessage
        message={
          <>
            <p>Bridging is paused until maintenance is complete.</p>
          </>
        }
      />
    )
  }

  return null
}

export const useMaintenanceCountdownProgress = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isCurrentChain = isChainIncluded(
    [fromChainId, toChainId],
    [OPTIMISM.id, BASE.id] // Update for Chains to show maintenance on
  )

  const {
    isPending: isMaintenancePending,
    EventCountdownProgressBar: MaintenanceCountdownProgressBar,
  } = useEventCountdownProgressBar(
    'Maintenance in progress',
    MAINTENANCE_START_DATE, // Countdown Bar will automatically appear after start time
    MAINTENANCE_END_DATE // Countdown Bar will automatically disappear when end time is reached
  )

  return {
    isMaintenancePending,
    isCurrentChainDisabled: isCurrentChain && isMaintenancePending, // Used to pause Bridge
    MaintenanceCountdownProgressBar: isCurrentChain
      ? MaintenanceCountdownProgressBar
      : null,
  }
}

/**
 * Checks if any of the chain IDs in `hasChains` are found within the `chainList` array.
 *
 * @param {number[]} chainList - The array of chain IDs to check against.
 * @param {number[]} hasChains - The array of chain IDs to find within `checkChains`.
 * @returns {boolean} - True if any chain ID from `hasChains` is found in `checkChains`, otherwise false.
 */
const isChainIncluded = (chainList: number[], hasChains: number[]) => {
  return hasChains.some((chainId) => chainList.includes(chainId))
}
