import { useBridgeState } from '@/slices/bridge/hooks'
import { getCountdownTimeStatus } from './EventCountdownProgressBar'
import { WarningMessage } from '../../Warning'
import { isChainIncluded } from '@/utils/isChainIncluded'

/**
 * Displays warning message triggered by start and end time.
 * Renders for selected origin and/or destination chain IDs.
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
