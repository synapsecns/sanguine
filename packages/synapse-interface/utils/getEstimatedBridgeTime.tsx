import { useSynapseContext } from './providers/SynapseProvider'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { Chain } from './types'

export const getEstimatedBridgeTime = ({
  bridgeOriginChain,
  bridgeModuleName,
  formattedEventType,
}: {
  bridgeOriginChain: Chain
  bridgeModuleName: string
  formattedEventType: string
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
  return (
    (BRIDGE_REQUIRED_CONFIRMATIONS[bridgeOriginChain.id] *
      bridgeOriginChain.blockTime) /
    1000
  )
}
