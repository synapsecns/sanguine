import { Zero } from '@ethersproject/constants'

import { ChainId } from '@constants/networks'

import { useSingleCallResult } from "@hooks/multicall"
import { useGenericSynapseContract } from "@hooks/contracts/useSynapseContract"


export function useGasDropAmount(chainId) {
  const synapseContract = useGenericSynapseContract(chainId)

  const gasDropAmountResult = useSingleCallResult(
    chainId,
    synapseContract,
    'chainGasAmount',
    [],
    { resultOnly: true }
  )

  if (chainId == ChainId.ETH) {
    return Zero
  } else {
    return gasDropAmountResult?.[0] ?? Zero
  }
}


