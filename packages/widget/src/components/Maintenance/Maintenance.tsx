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

export const getMaintenanceData = (pausedChains: any, pausedModules: any) => {
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

  return {
    pausedChainsList,
    pausedModulesList,
  }
}
