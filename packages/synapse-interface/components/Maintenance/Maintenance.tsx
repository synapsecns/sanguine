import { OPTIMISM, BASE } from '@/constants/chains/master'
import { MaintenanceBanner } from './components/MaintenanceBanner'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
import { useMaintenanceCountdownProgress } from './components/useMaintenanceCountdownProgress'

/**
 * Checks if any of the chain IDs in `hasChains` are found within the `chainList` array.
 *
 * @param {number[]} chainList - The array of chain IDs to check against.
 * @param {number[]} hasChains - The array of chain IDs to find within `checkChains`.
 * @returns {boolean} - True if any chain ID from `hasChains` is found in `checkChains`, otherwise false.
 */
export const isChainIncluded = (chainList: number[], hasChains: number[]) => {
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
