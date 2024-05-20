import { getCountdownTimeStatus } from '@/utils/getCountdownTimeStatus'
import { isChainIncluded } from '@/utils/isChainIncluded'

/**
 * Displays warning message triggered by start and end time.
 * Renders for selected origin and/or destination chain IDs.
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
    <div className="flex flex-col bg-[--synapse-surface] text-[--synapse-text] p-3 rounded-md">
      {header && <div className="mb-2 font-bold">{header}</div>}
      {message && <div>{message}</div>}
    </div>
  )
}
