import { useBridgeState } from '@/slices/bridge/hooks'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { WarningMessage } from '../../Warning'
import { isChainIncluded } from '@/utils/isChainIncluded'

/**
 * Displays a warning message based on selected chains for Bridge / Swap.
 * Able to specify which chains to display warning messages for.
 *
 * @param startDate Start time to display message
 * @param endDate End time to remove message
 * @param pausedChains A list of chainIds to display warning messages for
 * @param warningMessage Message to display if User is connected to paused chainIds
 */
export const MaintenanceWarningMessage = ({
  fromChainId,
  toChainId,
  startDate,
  endDate,
  pausedFromChains,
  pausedToChains,
  warningMessage,
  disabled = false,
}: {
  fromChainId: number
  toChainId: number
  startDate: Date
  endDate: Date
  pausedFromChains: number[]
  pausedToChains: number[]
  warningMessage: any
  disabled?: boolean
}) => {
  const isWarningChain =
    isChainIncluded([fromChainId], pausedFromChains) ||
    isChainIncluded([toChainId], pausedToChains)

  const { isComplete } = getCountdownTimeStatus(startDate, endDate)

  if (isComplete || disabled) return null

  if (isWarningChain) {
    return <WarningMessage message={warningMessage} />
  }

  return null
}
