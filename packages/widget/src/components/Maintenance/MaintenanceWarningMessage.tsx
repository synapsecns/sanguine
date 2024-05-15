import { useBridgeState } from '@/state/slices/bridge/hooks'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
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
  twClassName,
}: {
  header?: string
  message?: React.ReactNode
  twClassName?: string
}) => {
  return (
    <div
      className={`flex flex-col bg-[#353038] text-white text-sm p-3 rounded-md ${twClassName}`}
    >
      {header && <div className="mb-2 font-bold">{header}</div>}
      {message && <div>{message}</div>}
    </div>
  )
}
