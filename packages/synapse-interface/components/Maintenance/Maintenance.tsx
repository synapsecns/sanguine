import { isEmpty } from 'lodash'
import { MaintenanceBanner } from './components/MaintenanceBanner'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
import { useMaintenanceCountdownProgress } from './components/useMaintenanceCountdownProgress'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useSwapState } from '@/slices/swap/hooks'
import { useMaintanceState } from '@/slices/maintenance/hooks'
import pausedChains from '@/public/pauses/v1/paused-chains.json'
import pausedBridgeModules from '@/public/pauses/v1/paused-bridge-modules.json'
import { isChainIncluded } from '@/utils/isChainIncluded'
import { useEventCountdownProgressBar } from './components/EventCountdownProgressBar'

/** Pause Chain Activity */
interface ChainPause {
  id: string
  pausedFromChains: number[]
  pausedToChains: number[]
  pauseBridge: boolean
  pauseSwap: boolean
  startTimePauseChain: Date
  endTimePauseChain: Date | null // If null, pause indefinitely
  startTimeBanner: Date
  endTimeBanner: Date | null // If null, pause indefinitely
  inputWarningMessage: string
  bannerMessage: JSX.Element
  progressBarMessage: string
  disableBanner: boolean
  disableWarning: boolean
  disableCountdown: boolean
}

// const PAUSED_CHAINS: ChainPause[] = pausedChains.map((pause) => {
//   return {
//     ...pause,
//     startTimePauseChain: new Date(pause.startTimePauseChain),
//     endTimePauseChain: pause.endTimePauseChain
//       ? new Date(pause.endTimePauseChain)
//       : null,
//     startTimeBanner: new Date(pause.startTimeBanner),
//     endTimeBanner: pause.endTimeBanner ? new Date(pause.endTimeBanner) : null,
//     inputWarningMessage: <p>{pause.inputWarningMessage}</p>,
//     bannerMessage: <p className="text-left">{pause.bannerMessage}</p>,
//     progressBarMessage: <p>{pause.progressBarMessage}</p>,
//   }
// })

const useMaintenanceData = () => {
  const { pausedChainsData, pausedModulesData } = useMaintanceState()

  const pausedChainsList: ChainPause[] = pausedChainsData
    ? pausedChainsData?.map((pause: ChainPause) => {
        return {
          ...pause,
          startTimeBanner: new Date(pause.startTimeBanner),
          endTimeBanner: pause.endTimeBanner
            ? new Date(pause.endTimeBanner)
            : null,
          startTimePauseChain: new Date(pause.startTimePauseChain),
          endTimePauseChain: pause.endTimePauseChain
            ? new Date(pause.endTimePauseChain)
            : null,
          bannerMessage: <p className="text-left">{pause.bannerMessage}</p>,
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

export const MaintenanceBanners = () => {
  const { pausedChainsList } = useMaintenanceData()
  const { fromChainId: bridgeFromChainId, toChainId: bridgeToChainId } =
    useBridgeState()
  const { swapChainId } = useSwapState()

  const activeBanner = pausedChainsList.find(
    (pausedChain) =>
      isChainIncluded(pausedChain?.pausedFromChains, [bridgeFromChainId]) ||
      isChainIncluded(pausedChain?.pausedToChains, [bridgeToChainId]) ||
      isChainIncluded(pausedChain?.pausedFromChains, [swapChainId]) ||
      isChainIncluded(pausedChain?.pausedToChains, [swapChainId])
  )

  if (activeBanner) {
    return (
      <MaintenanceBanner
        id={activeBanner?.id}
        bannerMessage={activeBanner?.bannerMessage}
        startDate={activeBanner?.startTimeBanner}
        endDate={activeBanner?.endTimeBanner}
        disabled={activeBanner?.disableBanner}
      />
    )
  }
}

export const useMaintenance = () => {
  const { pausedChainsList, pausedModulesList } = useMaintenanceData()
  const { fromChainId: bridgeFromChainId, toChainId: bridgeToChainId } =
    useBridgeState()
  const { swapChainId } = useSwapState()

  const activeBridgePause = pausedChainsList.find(
    (pausedChain) =>
      isChainIncluded(pausedChain?.pausedFromChains, [bridgeFromChainId]) ||
      isChainIncluded(pausedChain?.pausedToChains, [bridgeToChainId])
  )

  const activeSwapPause = pausedChainsList.find(
    (pausedChain) =>
      isChainIncluded(pausedChain?.pausedFromChains, [swapChainId]) ||
      isChainIncluded(pausedChain?.pausedToChains, [swapChainId])
  )

  const {
    isPending: isBridgePaused,
    EventCountdownProgressBar: BridgeEventCountdownProgressBar,
  } = useEventCountdownProgressBar(
    activeBridgePause?.progressBarMessage,
    activeBridgePause?.startTimePauseChain,
    activeBridgePause?.endTimePauseChain,
    activeBridgePause?.disableCountdown
  )

  const {
    isPending: isSwapPaused,
    EventCountdownProgressBar: SwapEventCountdownProgressBar,
  } = useEventCountdownProgressBar(
    activeSwapPause?.progressBarMessage,
    activeSwapPause?.startTimePauseChain,
    activeSwapPause?.endTimePauseChain,
    activeSwapPause?.disableCountdown
  )

  const BridgeMaintenanceProgressBar = () => BridgeEventCountdownProgressBar
  const SwapMaintenanceProgressBar = () => SwapEventCountdownProgressBar

  const BridgeMaintenanceWarningMessage = () => (
    <MaintenanceWarningMessage
      fromChainId={bridgeFromChainId}
      toChainId={bridgeToChainId}
      startDate={activeBridgePause?.startTimePauseChain}
      endDate={activeBridgePause?.endTimePauseChain}
      pausedFromChains={activeBridgePause?.pausedFromChains}
      pausedToChains={activeBridgePause?.pausedToChains}
      warningMessage={activeBridgePause?.inputWarningMessage}
      disabled={
        activeBridgePause?.disableWarning || !activeBridgePause?.pauseBridge
      }
    />
  )

  const SwapMaintenanceWarningMessage = () => (
    <MaintenanceWarningMessage
      fromChainId={swapChainId}
      toChainId={null}
      startDate={activeSwapPause?.startTimePauseChain}
      endDate={activeSwapPause?.endTimePauseChain}
      pausedFromChains={activeSwapPause?.pausedFromChains}
      pausedToChains={activeSwapPause?.pausedToChains}
      warningMessage={activeSwapPause?.inputWarningMessage}
      disabled={activeSwapPause?.disableWarning || !activeSwapPause?.pauseSwap}
    />
  )

  return {
    isBridgePaused,
    isSwapPaused,
    pausedChainsList,
    pausedModulesList,
    BridgeMaintenanceProgressBar,
    BridgeMaintenanceWarningMessage,
    SwapMaintenanceProgressBar,
    SwapMaintenanceWarningMessage,
  }
}

/** Pause Bridge Modules */
interface BridgeModulePause {
  chainId?: number // Will pause for all chains if undefined
  bridgeModuleName: 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL'
}

function isValidBridgeModule(
  module: any
): module is 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL' {
  return ['SynapseBridge', 'SynapseRFQ', 'SynapseCCTP', 'ALL'].includes(module)
}

export function getBridgeModuleNames(module) {
  if (module.bridgeModuleName === 'ALL') {
    return ['SynapseRFQ', 'SynapseCCTP', 'SynapseBridge']
  }
  return [module.bridgeModuleName]
}

export const PAUSED_MODULES: BridgeModulePause[] = pausedBridgeModules.map(
  (route) => {
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
