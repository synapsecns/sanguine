import { useBridgeState } from '@/state/slices/bridge/hooks'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
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

export const useMaintenanceComponents = (
  pausedChains: any,
  pausedModules: any
) => {
  const pausedChainsList: ChainPause[] = pausedChains
    ? pausedChains?.map((pause: ChainPause) => {
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

  const pausedModulesList: BridgeModulePause[] = pausedModules
    ? pausedModules?.map((route: BridgeModulePause) => {
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

  const MaintenanceWarningMessages = () => {
    const { originChainId, destinationChainId } = useBridgeState()

    return (
      <>
        {pausedChainsList?.map((event) => {
          return (
            <MaintenanceWarningMessage
              key={event.id}
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

  return {
    pausedChainsList,
    pausedModulesList,
    MaintenanceWarningMessages,
  }
}
