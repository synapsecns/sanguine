import { Chain } from 'viem'
import { useSynapseContext } from './providers/SynapseProvider'

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
}
