import { useBridgeState } from '@/slices/bridge/hooks'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { WarningMessage } from '../Warning'
import { isChainIncluded } from './MaintenanceEvent'

/**
 * Creates a Warning Message based on that automatically appears/disappears based on start/end times.
 *
 * @param {Date} startDate - Date that automatically triggers displaying banner
 * @param {Date} endDate - Date that automatically triggers removing banner
 * @param {number[]} pausedChains - List of chain ids to display warning messages for, based on User selected from/to chains
 * @param {any} warningMessage - Allow for flexibility when constructing warning message
 */
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

  const isWarningChain = isChainIncluded([fromChainId, toChainId], pausedChains)

  const { isComplete } = getCountdownTimeStatus(startDate, endDate)

  if (isComplete) return null

  if (isWarningChain) {
    return <WarningMessage message={warningMessage} />
  }

  return null
}
