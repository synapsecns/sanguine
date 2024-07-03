import { MaintenanceBanner } from './components/MaintenanceBanner'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
import { useMaintenanceCountdownProgress } from './components/useMaintenanceCountdownProgress'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useSwapState } from '@/slices/swap/hooks'
import { useMaintanceState } from '@/slices/maintenance/hooks'
import pausedChains from '@/public/pauses/v1/paused-chains.json'
import pausedBridgeModules from '@/public/pauses/v1/paused-bridge-modules.json'

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
  inputWarningMessage: JSX.Element
  bannerMessage: JSX.Element
  progressBarMessage: JSX.Element
  disableBanner: boolean
  disableWarning: boolean
  disableCountdown: boolean
}

const PAUSED_CHAINS: ChainPause[] = pausedChains.map((pause) => {
  return {
    ...pause,
    startTimePauseChain: new Date(pause.startTimePauseChain),
    endTimePauseChain: pause.endTimePauseChain
      ? new Date(pause.endTimePauseChain)
      : null,
    startTimeBanner: new Date(pause.startTimeBanner),
    endTimeBanner: pause.endTimeBanner ? new Date(pause.endTimeBanner) : null,
    inputWarningMessage: <p>{pause.inputWarningMessage}</p>,
    bannerMessage: <p className="text-left">{pause.bannerMessage}</p>,
    progressBarMessage: <p>{pause.progressBarMessage}</p>,
  }
})

const useMaintenanceData = () => {
  const { pausedChainsData, pausedModulesData } = useMaintanceState()

  const pausedChainsList: ChainPause[] = pausedChainsData
    ? pausedChainsData?.map((pause: ChainPause) => {
        return {
          ...pause,
          startTimeBanner: new Date(pause.startTimeBanner),
          startTimePauseChain: new Date(pause.startTimePauseChain),
          endTimePauseChain: pause.endTimePauseChain
            ? new Date(pause.endTimePauseChain)
            : null,
          bannerMessage: <p className="text-left">{pause.bannerMessage}</p>,
          inputWarningMessage: <p>{pause.inputWarningMessage}</p>,
          progressBarMessage: <p>{pause.progressBarMessage}</p>,
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

  return (
    <>
      {pausedChainsList.map((event) => {
        return (
          <MaintenanceBanner
            id={event.id}
            bannerMessage={event.bannerMessage}
            startDate={event.startTimeBanner}
            endDate={event.endTimeBanner}
            disabled={event.disableBanner}
          />
        )
      })}
    </>
  )
}

export const MaintenanceWarningMessages = ({
  type,
}: {
  type: 'Bridge' | 'Swap'
}) => {
  const { pausedChainsList } = useMaintenanceData()
  const { fromChainId: bridgeFromChainId, toChainId: bridgeToChainId } =
    useBridgeState()
  const { swapChainId } = useSwapState()

  if (type === 'Bridge') {
    return (
      <>
        {pausedChainsList.map((event) => {
          return (
            <MaintenanceWarningMessage
              fromChainId={bridgeFromChainId}
              toChainId={bridgeToChainId}
              startDate={event.startTimePauseChain}
              endDate={event.endTimePauseChain}
              pausedFromChains={event.pausedFromChains}
              pausedToChains={event.pausedToChains}
              warningMessage={event.inputWarningMessage}
              disabled={event.disableWarning || !event.pauseBridge}
            />
          )
        })}
      </>
    )
  } else if (type === 'Swap') {
    return (
      <>
        {pausedChainsList.map((event) => {
          return (
            <MaintenanceWarningMessage
              fromChainId={swapChainId}
              toChainId={null}
              startDate={event.startTimePauseChain}
              endDate={event.endTimePauseChain}
              pausedFromChains={event.pausedFromChains}
              pausedToChains={event.pausedToChains}
              warningMessage={event.inputWarningMessage}
              disabled={event.disableWarning || !event.pauseSwap}
            />
          )
        })}
      </>
    )
  } else {
    return null
  }
}

/**
 * Hook that maps through PAUSED_CHAINS to apply the single event countdown progress logic to each.
 * @returns A list of objects containing maintenance status and components for each paused chain.
 */
export const useMaintenanceCountdownProgresses = ({
  type,
}: {
  type: 'Bridge' | 'Swap'
}) => {
  const { pausedChainsList } = useMaintenanceData()
  const { fromChainId: bridgeFromChainId, toChainId: bridgeToChainId } =
    useBridgeState()
  const { swapChainId } = useSwapState()

  if (type === 'Bridge') {
    return pausedChainsList.map((event) => {
      return useMaintenanceCountdownProgress({
        fromChainId: bridgeFromChainId,
        toChainId: bridgeToChainId,
        startDate: event.startTimePauseChain,
        endDate: event.endTimePauseChain,
        pausedFromChains: event.pausedFromChains,
        pausedToChains: event.pausedToChains,
        progressBarMessage: event.progressBarMessage,
        disabled: event.disableCountdown || !event.pauseBridge,
      })
    })
  } else if (type === 'Swap') {
    return pausedChainsList.map((event) => {
      return useMaintenanceCountdownProgress({
        fromChainId: swapChainId,
        toChainId: null,
        startDate: event.startTimePauseChain,
        endDate: event.endTimePauseChain,
        pausedFromChains: event.pausedFromChains,
        pausedToChains: event.pausedToChains,
        progressBarMessage: event.progressBarMessage,
        disabled: event.disableCountdown || !event.pauseSwap,
      })
    })
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
