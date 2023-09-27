import { Address } from 'viem'
import { useSynapseContext } from '../providers/SynapseProvider'

interface BridgeQuoteRequest {
  originChainId: number
  destinationChainId: number
  originTokenAddress: Address
  destinationTokenAddress: Address
  amount: bigint
}

function fetchBridgeQuote(request: BridgeQuoteRequest) {
  const { synapseSDK } = useSynapseContext()
}
