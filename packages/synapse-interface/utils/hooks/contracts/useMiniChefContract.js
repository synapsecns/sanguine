import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useContract, useGenericContract } from '@hooks/contracts/useContract'

import { MINICHEF_ADDRESSES } from '@constants/minichef'
import MINICHEF_ABI from '@abis/miniChef.json'

export function useMiniChefContract() {
  const { chainId } = useActiveWeb3React()
  const miniChefContract = useContract(MINICHEF_ADDRESSES[chainId], MINICHEF_ABI)
  return miniChefContract
}

export function useGenericMiniChefContract(chainId) {
  const miniChefContract = useGenericContract(chainId, MINICHEF_ADDRESSES[chainId], MINICHEF_ABI)
  return miniChefContract
}