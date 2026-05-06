import { getBridgeModuleNames } from '@/utils/getBridgeModuleNames'

type BridgeModulePauseLike = {
  chainId?: number
  toChainId?: number
  bridgeModuleName: string
}

type BridgeQuoteLike = {
  moduleNames: string[]
}

export const selectBridgeQuote = <T extends BridgeQuoteLike>({
  quotes,
  originChainId,
  destinationChainId,
  pausedModules,
}: {
  quotes: T[]
  originChainId: number
  destinationChainId: number
  pausedModules: BridgeModulePauseLike[]
}): T | null => {
  const pausedBridgeModules = new Set(
    pausedModules
      .filter(
        (module) =>
          (module.chainId ? module.chainId === originChainId : true) &&
          (module.toChainId ? module.toChainId === destinationChainId : true)
      )
      .flatMap(getBridgeModuleNames)
  )

  const activeQuotes = quotes.filter(
    (quote) =>
      !quote.moduleNames.some((moduleName) =>
        pausedBridgeModules.has(moduleName)
      )
  )

  const rfqQuote = activeQuotes.find((quote) =>
    quote.moduleNames.includes('SynapseRFQ')
  )

  return rfqQuote ?? activeQuotes[0] ?? null
}
