import { getBridgeModuleNames } from '@/utils/getBridgeModuleNames'

type BridgeModulePauseLike = {
  chainId?: number
  toChainId?: number
  bridgeModuleName: string
}

export const getPausedBridgeModuleNamesForRoute = ({
  pausedModules,
  fromChainId,
  toChainId,
}: {
  pausedModules: BridgeModulePauseLike[]
  fromChainId: number
  toChainId: number
}) => {
  return Array.from(
    new Set(
      pausedModules
        .filter(
          (module) =>
            (module.chainId ? module.chainId === fromChainId : true) &&
            (module.toChainId ? module.toChainId === toChainId : true)
        )
        .flatMap(getBridgeModuleNames)
    )
  ).sort()
}
