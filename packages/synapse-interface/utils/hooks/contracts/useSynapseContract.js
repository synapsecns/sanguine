import {
  SYNAPSE_BRIDGE_ADDRESSES
} from '@constants/bridge'

import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge'

import { ChainId } from '@constants/networks'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useContract, useGenericContract } from '@hooks/contracts/useContract'




export function useSynapseContract() {
  const { chainId } = useActiveWeb3React()

  const synapseContract = useContract(
    SYNAPSE_BRIDGE_ADDRESSES[chainId],
    SYNAPSE_BRIDGE_ABI
  )
  return synapseContract
}




export function useGenericSynapseContract(chainId) {
  const synapseContract = useGenericContract(
    chainId,
    SYNAPSE_BRIDGE_ADDRESSES[chainId],
    SYNAPSE_BRIDGE_ABI,
    false
  )

  return synapseContract
}