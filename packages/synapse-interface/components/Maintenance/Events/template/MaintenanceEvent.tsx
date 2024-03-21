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
  const { isComplete } = getCountdownTimeStatus(
    // MAINTENANCE_BANNERS_START, // Banner will automatically appear after start time
    // MAINTENANCE_END_DATE // Banner will automatically disappear when end time is reached
    startDate,
    endDate
  )

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

export const MaintenanceWarningMessage = ({
  startDate,
  endDate,
  pausedChains,
  warningMessage,
}: {
  startDate: Date
  endDate: Date
  pausedChains: number[]
  warningMessage: any
}) => {
  const { fromChainId, toChainId } = useBridgeState()

  const isWarningChain = isChainIncluded(
    [fromChainId, toChainId],
    // [OPTIMISM.id, BASE.id] // Update for Chains to show warning on
    pausedChains
  )

  const { isComplete } = getCountdownTimeStatus(
    // MAINTENANCE_BANNERS_START, // Banner will automatically appear after start time
    // MAINTENANCE_END_DATE // Banner will automatically disappear when end time is reached
    startDate,
    endDate
  )

  if (isComplete) return null

  if (isWarningChain) {
    return <WarningMessage message={warningMessage} />
  }

  return null
}

export const useMaintenanceCountdownProgress = ({
  startDate,
  endDate,
  pausedChains,
  progressBarMessage,
}: {
  startDate: Date
  endDate: Date
  pausedChains: number[]
  progressBarMessage: any
}) => {
  const { fromChainId, toChainId } = useBridgeState()

  const isCurrentChain = isChainIncluded(
    [fromChainId, toChainId],
    // [OPTIMISM.id, BASE.id] // Update for Chains to show maintenance on
    pausedChains
  )

  const {
    isPending: isMaintenancePending,
    EventCountdownProgressBar: MaintenanceCountdownProgressBar,
  } = useEventCountdownProgressBar(
    progressBarMessage,
    // MAINTENANCE_START_DATE, // Countdown Bar will automatically appear after start time
    // MAINTENANCE_END_DATE // Countdown Bar will automatically disappear when end time is reached
    startDate,
    endDate
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

/** Aggregators */

interface ChainPause {
  id: string
  pausedChains: number[]
  startTime: Date
  endTime: Date
  bannerStartTime: Date
  bannerEndTime: Date
  warningMessage: any
  bannerMessage: any
  progressBarMessage: any
}

const PAUSED_CHAINS: ChainPause[] = [
  {
    id: 'optimism-chain-pause',
    pausedChains: [OPTIMISM.id],
    startTime: new Date(Date.UTC(2024, 2, 21, 17, 0, 0)),
    endTime: new Date(Date.UTC(2024, 2, 21, 17, 40, 0)),
    bannerStartTime: new Date(Date.UTC(2024, 2, 21, 17, 0, 0)),
    bannerEndTime: new Date(Date.UTC(2024, 2, 21, 17, 41, 0)),
    warningMessage: (
      <p> Optimism bridging is paused until maintenance is complete. </p>
    ),
    bannerMessage: (
      <p> Optimism bridging is paused until maintenance is complete. </p>
    ),
    progressBarMessage: <p> Optimism maintenance in progress </p>,
  },
  {
    id: 'optimism-chain-pause-2',
    pausedChains: [BASE.id],
    startTime: new Date(Date.UTC(2024, 2, 21, 17, 41, 0)),
    endTime: new Date(Date.UTC(2024, 2, 21, 17, 42, 0)),
    bannerStartTime: new Date(Date.UTC(2024, 2, 21, 17, 40, 0)),
    bannerEndTime: new Date(Date.UTC(2024, 2, 21, 17, 43, 0)),
    warningMessage: (
      <p> Base bridging is paused until maintenance is complete. </p>
    ),
    bannerMessage: (
      <p> Base bridging is paused until maintenance is complete. </p>
    ),
    progressBarMessage: <p> Base maintenance in progress </p>,
  },
]

export const MaintenanceBanners = () => {
  return (
    <>
      {PAUSED_CHAINS.map((event) => {
        return (
          <MaintenanceBanner
            id={event.id}
            bannerMessage={event.bannerMessage}
            startDate={event.bannerStartTime}
            endDate={event.bannerEndTime}
          />
        )
      })}
    </>
  )
}

export const MaintenanceWarningMessages = () => {
  return (
    <>
      {PAUSED_CHAINS.map((event) => {
        return (
          <MaintenanceWarningMessage
            startDate={event.startTime}
            endDate={event.endTime}
            pausedChains={event.pausedChains}
            warningMessage={event.warningMessage}
          />
        )
      })}
    </>
  )
}

/**
 * Hook that maps through PAUSED_CHAINS to apply the single chain countdown progress logic to each.
 * @returns Array of objects containing maintenance status and components for each paused chain.
 */
export const useMaintenanceCountdownProgresses = () => {
  return PAUSED_CHAINS.map((event) => {
    return useMaintenanceCountdownProgress({
      startDate: event.startTime,
      endDate: event.endTime,
      pausedChains: event.pausedChains,
      progressBarMessage: event.progressBarMessage,
    })
  })
}
