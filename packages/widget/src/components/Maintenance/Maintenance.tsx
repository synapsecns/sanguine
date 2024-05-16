import { useBridgeState } from '@/state/slices/bridge/hooks'
import { isChainIncluded } from '@/utils/isChainIncluded'
import { useEventCountdownProgressBar } from './hooks/useEventCountdownProgressBar'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
import { getSynapsePauseData } from './helpers/getSynapsePauseData'
import { isValidBridgeModule } from './helpers/isValidBridgeModule'

interface ChainPause {
  id: string
  pausedFromChains: number[]
  pausedToChains: number[]
  pauseBridge: boolean
  startTimePauseChain: Date
  endTimePauseChain: Date | null
  inputWarningMessage: string
  progressBarMessage: string
  disableWarning: boolean
  disableCountdown: boolean
}

interface BridgeModulePause {
  chainId: number | undefined
  bridgeModuleName: 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL'
}

export const getMaintenanceData = () => {
  const { pausedChainsData, pausedModulesData } = getSynapsePauseData()

  const pausedChainsList: ChainPause[] = pausedChainsData
    ? pausedChainsData?.map((pause: ChainPause) => {
        return {
          ...pause,
          startTimePauseChain: new Date(pause.startTimePauseChain),
          endTimePauseChain: pause.endTimePauseChain
            ? new Date(pause.endTimePauseChain)
            : null,
          inputWarningMessage: pause.inputWarningMessage,
          progressBarMessage: pause.progressBarMessage,
        }
      })
    : []

  const pausedModulesList: BridgeModulePause[] = pausedModulesData
    ? pausedModulesData?.map((route: BridgeModulePause) => {
        if (!isValidBridgeModule(route.bridgeModuleName)) {
          throw new Error(`Invalid module type: ${route.bridgeModuleName}`)
        }

        return {
          ...route,
          bridgeModuleName: route.bridgeModuleName as
            | 'SynapseBridge'
            | 'SynapseRFQ'
            | 'SynapseCCTP'
            | 'ALL',
        }
      })
    : []

  return {
    pausedChainsList,
    pausedModulesList,
  }
}

export const useMaintenance = () => {
  const { originChainId, destinationChainId } = useBridgeState()
  const { pausedChainsList, pausedModulesList } = getMaintenanceData()

  const activePause = pausedChainsList.find(
    (pauseData) =>
      isChainIncluded(pauseData?.pausedFromChains, [originChainId]) ||
      isChainIncluded(pauseData?.pausedToChains, [destinationChainId])
  )

  const { isPending: isBridgePaused, EventCountdownProgressBar } =
    useEventCountdownProgressBar(
      activePause?.progressBarMessage,
      activePause?.startTimePauseChain,
      activePause?.endTimePauseChain,
      activePause?.disableCountdown
    )

  const BridgeMaintenanceProgressBar = () => EventCountdownProgressBar

  const BridgeMaintenanceWarningMessage = () => (
    <MaintenanceWarningMessage
      originChainId={originChainId}
      destinationChainId={destinationChainId}
      startDate={activePause?.startTimePauseChain}
      endDate={activePause?.endTimePauseChain}
      pausedOriginChains={activePause?.pausedFromChains}
      pausedDestinationChains={activePause?.pausedToChains}
      warningMessage={activePause?.inputWarningMessage}
      disabled={activePause?.disableWarning}
    />
  )

  return {
    isBridgePaused,
    pausedChainsList,
    pausedModulesList,
    BridgeMaintenanceProgressBar,
    BridgeMaintenanceWarningMessage,
  }
}
