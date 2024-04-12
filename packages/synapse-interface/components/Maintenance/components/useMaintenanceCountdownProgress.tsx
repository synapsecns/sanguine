import { useBridgeState } from '@/slices/bridge/hooks'
import { useEventCountdownProgressBar } from './EventCountdownProgressBar'
import { isChainIncluded } from '@/utils/isChainIncluded'

/**
 * A custom hook that provides logic for showing a countdown progress bar and determining if the bridge
 * should be paused based on the current chain and a maintenance schedule.
 *
 * @param {Date} startDate - The start date and time for the maintenance event.
 * @param {Date} endDate - The end date and time for the maintenance event.
 * @param {number[]} pausedChains - An array of chain IDs for which the bridge should be paused during the maintenance event.
 * @param {any} progressBarMessage - The message or content to display in the countdown progress bar.
 * @returns An object containing:
 * - isMaintenancePending: A boolean indicating if the maintenance is currently pending (i.e., ongoing).
 * - isCurrentChainDisabled: A boolean indicating if the current chain selected by the user is affected by the maintenance and should therefore be considered "disabled" or paused.
 * - MaintenanceCountdownProgressBar: A component (or null) that renders the countdown progress bar if the current chain is affected by the maintenance.
 */
export const useMaintenanceCountdownProgress = ({
  fromChainId,
  toChainId,
  startDate,
  endDate,
  pausedFromChains,
  pausedToChains,
  progressBarMessage,
  disabled = false,
}: {
  fromChainId: number
  toChainId: number
  startDate: Date
  endDate: Date | null
  pausedFromChains: number[]
  pausedToChains: number[]
  progressBarMessage: any
  disabled?: boolean
}) => {
  // const { fromChainId, toChainId } = useBridgeState()

  const isCurrentChain =
    isChainIncluded([fromChainId], pausedFromChains) ||
    isChainIncluded([toChainId], pausedToChains)

  const {
    isPending: isMaintenancePending,
    EventCountdownProgressBar: MaintenanceCountdownProgressBar,
  } = useEventCountdownProgressBar(progressBarMessage, startDate, endDate)

  return {
    isMaintenancePending,
    isCurrentChainDisabled: isCurrentChain && isMaintenancePending, // Used to pause Bridge
    MaintenanceCountdownProgressBar:
      isCurrentChain && !disabled ? MaintenanceCountdownProgressBar : null,
  }
}
