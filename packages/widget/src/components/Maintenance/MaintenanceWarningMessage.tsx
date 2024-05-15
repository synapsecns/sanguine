import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { isChainIncluded } from '@/utils/isChainIncluded'

/**
 * Displays a warning message based on selected chains in Bridge Input
 * Able to specify which chains to display warning messages for.
 *
 * @param startDate Start time to display message
 * @param endDate End time to remove message
 * @param pausedChains A list of chainIds to display warning messages for
 * @param warningMessage Message to display if User is connected to paused chainIds
 */
export const MaintenanceWarningMessage = ({
  originChainId,
  destinationChainId,
  startDate,
  endDate,
  pausedOriginChains,
  pausedDestinationChains,
  warningMessage,
  disabled = false,
}: {
  originChainId: number
  destinationChainId: number
  startDate: Date
  endDate: Date
  pausedOriginChains: number[]
  pausedDestinationChains: number[]
  warningMessage: any
  disabled?: boolean
}) => {
  const isWarningChain =
    isChainIncluded([originChainId], pausedOriginChains) ||
    isChainIncluded([destinationChainId], pausedDestinationChains)

  const { isComplete } = getCountdownTimeStatus(startDate, endDate)

  if (isComplete || disabled) return null

  if (isWarningChain) {
    return <WarningMessage message={warningMessage} />
  }

  return null
}

export const WarningMessage = ({
  header,
  message,
}: {
  header?: string
  message?: React.ReactNode
}) => {
  return (
    <div className="flex flex-col bg-[--synapse-surface] text-[--synapse-text] text-sm p-3 rounded-md">
      {header && <div className="mb-2 font-bold">{header}</div>}
      {message && <div>{message}</div>}
    </div>
  )
}
