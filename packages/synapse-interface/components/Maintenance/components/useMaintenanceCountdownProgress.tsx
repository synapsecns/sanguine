import { useBridgeState } from '@/slices/bridge/hooks'
import { useEventCountdownProgressBar } from './EventCountdownProgressBar'
import { isChainIncluded } from '@/utils/isChainIncluded'

/**
 * A hook that will return a constructed progress bar instance
 * and Event status checks to use for pausing Bridge / Swap.
 *
 * @param startDate Start time of event to track
 * @param endDate End time of event to track
 * @param pausedChains A list of chain IDs that is paused for Bridge / Swap
 * @param progressBarMessage The message to be displayed alongside the Progress Bar
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
  const isCurrentChain =
    isChainIncluded([fromChainId], pausedFromChains) ||
    isChainIncluded([toChainId], pausedToChains)

  const {
    isPending: isMaintenancePending,
    EventCountdownProgressBar: MaintenanceCountdownProgressBar,
  } = useEventCountdownProgressBar(progressBarMessage, startDate, endDate)

  return {
    isMaintenancePending,
    isCurrentChainDisabled: isCurrentChain && isMaintenancePending && !disabled,
    MaintenanceCountdownProgressBar:
      isCurrentChain && !disabled ? MaintenanceCountdownProgressBar : null,
  }
}
