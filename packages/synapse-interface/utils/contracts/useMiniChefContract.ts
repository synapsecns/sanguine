import { useContract, useNetwork } from 'wagmi'
import { Contract } from 'ethers'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import MINI_CHEF_ABI from '@/constants/abis/miniChef.json'

export const useMiniChefContract = (): Contract => {
  const { chain: connectedChain } = useNetwork()
  const miniChefContract = useContract({
    address: MINICHEF_ADDRESSES[connectedChain.id],
    abi: MINI_CHEF_ABI,
  })
  return miniChefContract
}
