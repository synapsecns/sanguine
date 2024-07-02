import { useEventCountdownProgressBar } from './EventCountdownProgressBar'
import { isChainIncluded } from '@/utils/isChainIncluded'

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
    isCurrentChainDisabled: isCurrentChain && isMaintenancePending,
    MaintenanceCountdownProgressBar:
      isCurrentChain && !disabled ? MaintenanceCountdownProgressBar : null,
  }
}
