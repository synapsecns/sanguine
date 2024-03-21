import { useBridgeState } from '@/slices/bridge/hooks'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { WarningMessage } from '../Warning'
import { isChainIncluded } from './MaintenanceEvent'

/**
 * This component displays a warning message during a specified maintenance window for selected blockchain chains.
 * It checks if the current chain selected by the user is within the paused chains and if the current time is within
 * the maintenance window. If so, it displays a custom warning message.
 *
 * @param {Date} startDate - The starting date and time when the warning message should begin appearing.
 * @param {Date} endDate - The ending date and time when the warning message should stop appearing.
 * @param {number[]} pausedChains - An array of chain IDs that the warning message applies to. The message will
 *                                  only appear if the user's current from or to chain is in this list.
 * @param {any} warningMessage - The content of the warning message to be displayed. This allows for flexibility
 *                               in the message's structure and content.
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
