import { useBridgeState } from '@/state/slices/bridge/hooks'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
import { useMaintenanceCountdownProgress } from './hooks/useMaintenanceCountdownProgress'
import { isValidBridgeModule } from './helpers/isValidBridgeModule'

interface ChainPause {
  id: string
  pausedFromChains: number[]
  pausedToChains: number[]
  pauseBridge: boolean
  startTimePauseChain: Date
  endTimePauseChain: Date | null // If null, pause chain indefinitely
  inputWarningMessage: JSX.Element
  progressBarMessage: JSX.Element
  disableWarning: boolean
  disableCountdown: boolean
}

interface BridgeModulePause {
  chainId?: number // If undefined, pause module for all chains
  bridgeModuleName: 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL'
}

export const useMaintenanceComponents = (
  pausedChains: any,
  pausedModules: any
) => {
  const pausedChainsList: ChainPause[] = pausedChains?.map(
    (pause: ChainPause) => {
      return {
        ...pause,
        startTimePauseChain: new Date(pause.startTimePauseChain),
        endTimePauseChain: pause.endTimePauseChain
          ? new Date(pause.endTimePauseChain)
          : null,
        inputWarningMessage: <div>{pause.inputWarningMessage}</div>,
        progressBarMessage: <div>{pause.progressBarMessage}</div>,
      }
    }
  )

  const pausedModulesList: BridgeModulePause[] = pausedModules?.map(
    (route: BridgeModulePause) => {
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
    }
  )

  const MaintenanceWarningMessages = () => {
    const { originChainId, destinationChainId } = useBridgeState()

    return (
      <>
        {pausedChainsList?.map((event) => {
          return (
            <MaintenanceWarningMessage
              originChainId={originChainId}
              destinationChainId={destinationChainId}
              startDate={event.startTimePauseChain}
              endDate={event.endTimePauseChain}
              pausedOriginChains={event.pausedFromChains}
              pausedDestinationChains={event.pausedToChains}
              warningMessage={event.inputWarningMessage}
              disabled={event.disableWarning || !event.pauseBridge}
            />
          )
        })}
      </>
    )
  }

  const useMaintenanceCountdownProgresses = () => {
    const { originChainId, destinationChainId } = useBridgeState()

    return pausedChainsList?.map((event) => {
      return useMaintenanceCountdownProgress({
        originChainId,
        destinationChainId,
        startDate: event.startTimePauseChain,
        endDate: event.endTimePauseChain,
        pausedFromChains: event.pausedFromChains,
        pausedToChains: event.pausedToChains,
        progressBarMessage: event.progressBarMessage,
        disabled: event.disableCountdown,
      })
    })
  }

  return {
    pausedModules: pausedModulesList,
    MaintenanceWarningMessages,
    useMaintenanceCountdownProgresses,
  }
}
