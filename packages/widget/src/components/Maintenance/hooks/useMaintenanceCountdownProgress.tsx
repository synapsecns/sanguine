import { useEventCountdownProgressBar } from './useEventCountdownProgressBar'
import { isChainIncluded } from '@/utils/isChainIncluded'

export const useMaintenanceCountdownProgress = ({
  originChainId,
  destinationChainId,
  startDate,
  endDate,
  pausedFromChains,
  pausedToChains,
  progressBarMessage,
  disabled = false,
}: {
  originChainId: number
  destinationChainId: number
  startDate: Date
  endDate: Date | null
  pausedFromChains: number[]
  pausedToChains: number[]
  progressBarMessage: any
  disabled?: boolean
}) => {
  const isCurrentChain =
    isChainIncluded([originChainId], pausedFromChains) ||
    isChainIncluded([destinationChainId], pausedToChains)

  const {
    isPending: isMaintenancePending,
    EventCountdownProgressBar: MaintenanceCountdownProgressBar,
  } = useEventCountdownProgressBar(progressBarMessage, startDate, endDate)

  return {
    isMaintenancePending,
    isCurrentChainDisabled: isCurrentChain && isMaintenancePending,
    MaintenanceCountdownProgressBar:
      isCurrentChain && !disabled ? MaintenanceCountdownProgressBar : null,
  }
}
