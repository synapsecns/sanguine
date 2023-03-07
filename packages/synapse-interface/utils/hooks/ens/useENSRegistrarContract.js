import ENS_ABI from '@abis/ensRegistrar.json'

import { ChainId } from '@constants/networks'
import { ENS_REGISTRAR_ADDRESS } from '@constants/ens'

import { useGenericContract } from '@hooks/contracts/useContract'





export function useENSRegistrarContract(withSignerIfPossible) {

  return useGenericContract(ChainId.ETH, ENS_REGISTRAR_ADDRESS, ENS_ABI, withSignerIfPossible)
}
