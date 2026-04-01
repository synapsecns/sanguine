import { getBridgeModuleNames } from '@/utils/getBridgeModuleNames'

type BridgeModulePauseLike = {
  chainId?: number
  bridgeModuleName: string
}

type BridgeQuoteLike = {
  moduleNames: string[]
}

export const selectBridgeQuote = <T extends BridgeQuoteLike>({
  quotes,
  originChainId,
  pausedModules,
}: {
  quotes: T[]
  originChainId: number
  pausedModules: BridgeModulePauseLike[]
}): T | null => {
  const pausedBridgeModules = new Set(
    pausedModules
      .filter((module) =>
        module.chainId ? module.chainId === originChainId : true
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
