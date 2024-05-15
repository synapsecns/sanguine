import { MaintenanceWarningMessage } from './MaintenanceWarningMessage'
import { useMaintenanceCountdownProgress } from './useMaintenanceCountdownProgress'
import { useBridgeState } from '@/state/slices/bridge/hooks'

interface ChainPause {
  id: string
  pausedFromChains: number[]
  pausedToChains: number[]
  pauseBridge: boolean
  startTimePauseChain: Date
  endTimePauseChain: Date | null // If null, pause indefinitely
  inputWarningMessage: JSX.Element
  progressBarMessage: JSX.Element
  disableWarning: boolean
  disableCountdown: boolean
}

interface BridgeModulePause {
  chainId?: number // Will pause for all chains if undefined
  bridgeModuleName: 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL'
}

export const useMaintenanceComponents = (
  pausedChains: any,
  pausedModules: any
) => {
  const pausedChainsList: ChainPause[] = pausedChains.map((pause) => {
    return {
      ...pause,
      startTimePauseChain: new Date(pause.startTimePauseChain),
      endTimePauseChain: pause.endTimePauseChain
        ? new Date(pause.endTimePauseChain)
        : null,
      inputWarningMessage: <div>{pause.inputWarningMessage}</div>,
      progressBarMessage: <div>{pause.progressBarMessage}</div>,
    }
  })

  console.log('pausedChainsList: ', pausedChainsList)

  const pausedModulesList: BridgeModulePause[] = pausedModules.map((route) => {
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

  const MaintenanceWarningMessages = () => {
    const { originChainId, destinationChainId } = useBridgeState()

    return (
      <>
        {pausedChainsList.map((event) => {
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

  /**
   * Hook that maps through pausedChainsList to apply the single event countdown progress logic to each.
   *
   * @returns A list of objects containing maintenance status and components for each paused chain.
   */
  const useMaintenanceCountdownProgresses = () => {
    const {
      originChainId: bridgeFromChainId,
      destinationChainId: bridgeToChainId,
    } = useBridgeState()

    return pausedChainsList?.map((event) => {
      return useMaintenanceCountdownProgress({
        originChainId: bridgeFromChainId,
        destinationChainId: bridgeToChainId,
        startDate: event.startTimePauseChain,
        endDate: event.endTimePauseChain,
        pausedFromChains: event.pausedFromChains,
        pausedToChains: event.pausedToChains,
        progressBarMessage: event.progressBarMessage,
        disabled: event.disableCountdown || !event.pauseBridge,
      })
    })
  }

  return {
    pausedModules: pausedModulesList,
    MaintenanceWarningMessages,
    useMaintenanceCountdownProgresses,
  }
}

const isValidBridgeModule = (
  module: any
): module is 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL' => {
  return ['SynapseBridge', 'SynapseRFQ', 'SynapseCCTP', 'ALL'].includes(module)
}

export const getBridgeModuleNames = (module) => {
  if (module.bridgeModuleName === 'ALL') {
    return ['SynapseRFQ', 'SynapseCCTP', 'SynapseBridge']
  }
  return [module.bridgeModuleName]
}
