export const SYNAPSE_EXPLORER_KAPPA = 'https://explorer.synapseprotocol.com/tx/'
export const SYNAPSE_EXPLORER = 'https://explorer.synapseprotocol.com/'

export const getTxSynapseExplorerLink = ({
  kappa,
  txHash,
  originChainId,
  destinationChainId,
}: {
  kappa?: string
  txHash?: string
  originChainId: number
  destinationChainId?: number
}): string => {
  if (kappa) {
    if (typeof originChainId === 'number') {
      return `${SYNAPSE_EXPLORER_KAPPA}${kappa}?chainIdFrom=${originChainId}&chainIdTo=${destinationChainId}`
    } else {
      return `${SYNAPSE_EXPLORER_KAPPA}${kappa}?chainIdFrom=${originChainId}`
    }
  }

  if (txHash) {
    return `${SYNAPSE_EXPLORER}txs?hash=${txHash}`
  }

  console.error('Transaction Hash and Synapse Transaction ID missing')
  return null
}
