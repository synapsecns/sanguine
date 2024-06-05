import { useSynapseContext } from './providers/SynapseProvider'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { Chain } from './types'

enum SynapseBridgeModule {
  BRIDGE = 'SynapseBridge',
  CCTP = 'SynapseCCTP',
}

/**
 * Fetches estimated duration of Bridge Transaction from Synapse SDK
 *
 * @param bridgeOriginChain - The selected origin chain.
 * @param bridgeModuleName - The name of the bridge module. e.g 'Bridge' or 'CCTP'.
 * @param formattedEventType - The name of the bridge event.
 * @returns - The estimated time for a bridge operation, in seconds.
 */
export const getEstimatedBridgeTime = ({
  bridgeOriginChain,
  bridgeModuleName,
  formattedEventType,
}: {
  bridgeOriginChain: Chain
  bridgeModuleName?: string
  formattedEventType?: string
}) => {
  const { synapseSDK } = useSynapseContext()

  if (!bridgeOriginChain) return null

  if (bridgeModuleName) {
    return synapseSDK.getEstimatedTime(bridgeOriginChain.id, bridgeModuleName)
  }

  if (formattedEventType) {
    const fetchedBridgeModuleName: string =
      synapseSDK.getBridgeModuleName(formattedEventType)

    return synapseSDK.getEstimatedTime(
      bridgeOriginChain?.id,
      fetchedBridgeModuleName
    )
  }

  // Fallback estimated time when inputs invalid
  return synapseSDK.getEstimatedTime(
    bridgeOriginChain.id,
    SynapseBridgeModule.BRIDGE
  )
}

export const getEstimatedBridgeTimeInMinutes = ({
  bridgeOriginChain,
  bridgeModuleName,
  formattedEventType,
}: {
  bridgeOriginChain: Chain
  bridgeModuleName?: string
  formattedEventType?: string
}) => {
  const estimatedBridgeTime = getEstimatedBridgeTime({
    bridgeOriginChain,
    bridgeModuleName,
    formattedEventType,
  })

  return estimatedBridgeTime ? Math.ceil(estimatedBridgeTime / 60) : null
}
